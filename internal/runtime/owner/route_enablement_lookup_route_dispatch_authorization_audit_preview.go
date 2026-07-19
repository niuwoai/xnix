package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditPreview struct {
	Version                                                      string                                                                                                                                                                                                                                           `json:"version"`
	SchemaVersion                                                string                                                                                                                                                                                                                                           `json:"schema_version"`
	RequestType                                                  string                                                                                                                                                                                                                                           `json:"request_type"`
	AuditType                                                    string                                                                                                                                                                                                                                           `json:"audit_type"`
	Source                                                       string                                                                                                                                                                                                                                           `json:"source"`
	AuditDecision                                                string                                                                                                                                                                                                                                           `json:"audit_decision"`
	CurrentMainlineConsumed                                      bool                                                                                                                                                                                                                                             `json:"current_mainline_consumed"`
	RouteEnablementLookupRouteDispatchGateConsumed               bool                                                                                                                                                                                                                                             `json:"route_enablement_lookup_route_dispatch_gate_consumed"`
	RouteEnablementLookupRouteDispatchGateReady                  bool                                                                                                                                                                                                                                             `json:"route_enablement_lookup_route_dispatch_gate_ready"`
	LookupRouteDispatchAuthorizationRequired                     bool                                                                                                                                                                                                                                             `json:"lookup_route_dispatch_authorization_required"`
	LookupRouteDispatchAuthorizationModeled                      bool                                                                                                                                                                                                                                             `json:"lookup_route_dispatch_authorization_modeled"`
	LookupRouteDispatchAuthorizationReady                        bool                                                                                                                                                                                                                                             `json:"lookup_route_dispatch_authorization_ready"`
	RouteEnablementLookupRouteDispatchAuthorizationBoundaryReady bool                                                                                                                                                                                                                                             `json:"route_enablement_lookup_route_dispatch_authorization_boundary_ready"`
	OpaqueRouteEnablementReceipt                                 bool                                                                                                                                                                                                                                             `json:"opaque_route_enablement_receipt"`
	KDESafeRedactedStatusOnly                                    bool                                                                                                                                                                                                                                             `json:"kde_safe_redacted_status_only"`
	CompatibilityCenterLookupRouteDispatchAuthorizationModeled   bool                                                                                                                                                                                                                                             `json:"compatibility_center_lookup_route_dispatch_authorization_modeled"`
	RuntimeDiagnosticsLookupRouteDispatchAuthorizationModeled    bool                                                                                                                                                                                                                                             `json:"runtime_diagnostics_lookup_route_dispatch_authorization_modeled"`
	ReceiptPresent                                               bool                                                                                                                                                                                                                                             `json:"receipt_present"`
	ReceiptAccepted                                              bool                                                                                                                                                                                                                                             `json:"receipt_accepted"`
	ReceiptConsumed                                              bool                                                                                                                                                                                                                                             `json:"receipt_consumed"`
	AcceptanceAuthorized                                         bool                                                                                                                                                                                                                                             `json:"acceptance_authorized"`
	RouteEnablementAccepted                                      bool                                                                                                                                                                                                                                             `json:"route_enablement_accepted"`
	LookupRouteEnablementGranted                                 bool                                                                                                                                                                                                                                             `json:"lookup_route_enablement_granted"`
	LookupRouteGrantAuthorized                                   bool                                                                                                                                                                                                                                             `json:"lookup_route_grant_authorized"`
	LookupRouteEnabled                                           bool                                                                                                                                                                                                                                             `json:"lookup_route_enabled"`
	LookupRouteDispatchAuthorized                                bool                                                                                                                                                                                                                                             `json:"lookup_route_dispatch_authorized"`
	LookupRouteDispatchAuthorizationPassed                       bool                                                                                                                                                                                                                                             `json:"lookup_route_dispatch_authorization_passed"`
	LookupRouteDispatchCallable                                  bool                                                                                                                                                                                                                                             `json:"lookup_route_dispatch_callable"`
	StorageRootPolicyGrantAuthorized                             bool                                                                                                                                                                                                                                             `json:"storage_root_policy_grant_authorized"`
	RecordWriterCallAuthorized                                   bool                                                                                                                                                                                                                                             `json:"record_writer_call_authorized"`
	WriterCallable                                               bool                                                                                                                                                                                                                                             `json:"writer_callable"`
	StorageRootResolved                                          bool                                                                                                                                                                                                                                             `json:"storage_root_resolved"`
	StorageWriteEnabled                                          bool                                                                                                                                                                                                                                             `json:"storage_write_enabled"`
	StatusPersistenceWriteEnabled                                bool                                                                                                                                                                                                                                             `json:"status_persistence_write_enabled"`
	DispatchAuthorizationItemCount                               int                                                                                                                                                                                                                                              `json:"dispatch_authorization_item_count"`
	RequiredDispatchAuthorizationItemCount                       int                                                                                                                                                                                                                                              `json:"required_dispatch_authorization_item_count"`
	ReadyDispatchAuthorizationItemCount                          int                                                                                                                                                                                                                                              `json:"ready_dispatch_authorization_item_count"`
	MissingDispatchAuthorizationItemCount                        int                                                                                                                                                                                                                                              `json:"missing_dispatch_authorization_item_count"`
	PassedDispatchAuthorizationItemCount                         int                                                                                                                                                                                                                                              `json:"passed_dispatch_authorization_item_count"`
	EnabledRouteItemCount                                        int                                                                                                                                                                                                                                              `json:"enabled_route_item_count"`
	PersistedDispatchAuthorizationItemCount                      int                                                                                                                                                                                                                                              `json:"persisted_dispatch_authorization_item_count"`
	RawExposedDispatchAuthorizationItemCount                     int                                                                                                                                                                                                                                              `json:"raw_exposed_dispatch_authorization_item_count"`
	SideEffectDispatchAuthorizationItemCount                     int                                                                                                                                                                                                                                              `json:"side_effect_dispatch_authorization_item_count"`
	CompatibilityCenterDispatchAuthorizationItemCount            int                                                                                                                                                                                                                                              `json:"compatibility_center_dispatch_authorization_item_count"`
	RuntimeDiagnosticsDispatchAuthorizationItemCount             int                                                                                                                                                                                                                                              `json:"runtime_diagnostics_dispatch_authorization_item_count"`
	DispatchAuthorizationItems                                   []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditItem      `json:"dispatch_authorization_items"`
	DispatchAuthorizationItemIDs                                 []string                                                                                                                                                                                                                                         `json:"dispatch_authorization_item_ids"`
	Checks                                                       []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditCheck     `json:"checks"`
	CheckIDs                                                     []string                                                                                                                                                                                                                                         `json:"check_ids"`
	Counts                                                       ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditCheckCounts `json:"counts"`
	ConsumerDryRunLookupResultBoundaryAuthorized                 bool                                                                                                                                                                                                                                             `json:"consumer_dry_run_lookup_result_boundary_authorized"`
	ConsumerEnablementAuthorized                                 bool                                                                                                                                                                                                                                             `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                                           bool                                                                                                                                                                                                                                             `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                                       bool                                                                                                                                                                                                                                             `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                                        bool                                                                                                                                                                                                                                             `json:"lookup_route_authorized"`
	LookupRouteEnablementAuthorized                              bool                                                                                                                                                                                                                                             `json:"lookup_route_enablement_authorized"`
	OpaqueLookupEnabled                                          bool                                                                                                                                                                                                                                             `json:"opaque_lookup_enabled"`
	RedactedSummaryPersisted                                     bool                                                                                                                                                                                                                                             `json:"redacted_summary_persisted"`
	KDEStatusPersisted                                           bool                                                                                                                                                                                                                                             `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted                                  bool                                                                                                                                                                                                                                             `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                                        bool                                                                                                                                                                                                                                             `json:"dry_run_result_persisted"`
	RawResultExposed                                             bool                                                                                                                                                                                                                                             `json:"raw_result_exposed"`
	DispatchDryRunExecuted                                       bool                                                                                                                                                                                                                                             `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled                                 bool                                                                                                                                                                                                                                             `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled                                 bool                                                                                                                                                                                                                                             `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                                         bool                                                                                                                                                                                                                                             `json:"portal_request_created"`
	NotificationActionEnabled                                    bool                                                                                                                                                                                                                                             `json:"notification_action_enabled"`
	CompatibilityCenterOpened                                    bool                                                                                                                                                                                                                                             `json:"compatibility_center_opened"`
	SupportBundleExported                                        bool                                                                                                                                                                                                                                             `json:"support_bundle_exported"`
	SupportCaseCreated                                           bool                                                                                                                                                                                                                                             `json:"support_case_created"`
	RuntimeOwned                                                 bool                                                                                                                                                                                                                                             `json:"runtime_owned"`
	GoRuntimeBacked                                              bool                                                                                                                                                                                                                                             `json:"go_runtime_backed"`
	KDEPolicyOwner                                               bool                                                                                                                                                                                                                                             `json:"kde_policy_owner"`
	ProductionReadiness                                          bool                                                                                                                                                                                                                                             `json:"production_readiness"`
	ProductionOwnershipReady                                     bool                                                                                                                                                                                                                                             `json:"production_ownership_ready"`
	SystemServiceStarted                                         bool                                                                                                                                                                                                                                             `json:"system_service_started"`
	SessionBusClaimed                                            bool                                                                                                                                                                                                                                             `json:"session_bus_claimed"`
	ProductionBusClaimed                                         bool                                                                                                                                                                                                                                             `json:"production_bus_claimed"`
	WriteMethodsEnabled                                          bool                                                                                                                                                                                                                                             `json:"write_methods_enabled"`
	RuntimeWritesEnabled                                         bool                                                                                                                                                                                                                                             `json:"runtime_writes_enabled"`
	DesktopFilesWritten                                          bool                                                                                                                                                                                                                                             `json:"desktop_files_written"`
	KDEConfigurationWritten                                      bool                                                                                                                                                                                                                                             `json:"kde_configuration_written"`
	PortalCallExecuted                                           bool                                                                                                                                                                                                                                             `json:"portal_call_executed"`
	AdapterInvocationEnabled                                     bool                                                                                                                                                                                                                                             `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                                         bool                                                                                                                                                                                                                                             `json:"backend_launch_enabled"`
	BackendProcessStarted                                        bool                                                                                                                                                                                                                                             `json:"backend_process_started"`
	NetworkRequired                                              bool                                                                                                                                                                                                                                             `json:"network_required"`
	HostRootModified                                             bool                                                                                                                                                                                                                                             `json:"host_root_modified"`
	PrivilegedContainerRequired                                  bool                                                                                                                                                                                                                                             `json:"privileged_container_required"`
	StateRootPathExposed                                         bool                                                                                                                                                                                                                                             `json:"state_root_path_exposed"`
	FilePathsExposed                                             bool                                                                                                                                                                                                                                             `json:"file_paths_exposed"`
	FileContentRead                                              bool                                                                                                                                                                                                                                             `json:"file_content_read"`
	RawCommandExposed                                            bool                                                                                                                                                                                                                                             `json:"raw_command_exposed"`
	RawExecutableExposed                                         bool                                                                                                                                                                                                                                             `json:"raw_executable_exposed"`
	BackendDetailsExposed                                        bool                                                                                                                                                                                                                                             `json:"backend_details_exposed"`
	BlockedActions                                               []string                                                                                                                                                                                                                                         `json:"blocked_actions"`
	NextRequirements                                             []string                                                                                                                                                                                                                                         `json:"next_requirements"`
	DesktopSafeSummary                                           string                                                                                                                                                                                                                                           `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditItem struct {
	ID                                                           string `json:"id"`
	ActionKind                                                   string `json:"action_kind"`
	SurfaceKind                                                  string `json:"surface_kind"`
	StatusConsumerKind                                           string `json:"status_consumer_kind"`
	OpaqueReceiptID                                              string `json:"opaque_receipt_id"`
	EvidencePresent                                              bool   `json:"evidence_present"`
	CurrentMainlineConsumed                                      bool   `json:"current_mainline_consumed"`
	RouteEnablementLookupRouteDispatchGateConsumed               bool   `json:"route_enablement_lookup_route_dispatch_gate_consumed"`
	RouteEnablementLookupRouteDispatchGateReady                  bool   `json:"route_enablement_lookup_route_dispatch_gate_ready"`
	LookupRouteDispatchAuthorizationModeled                      bool   `json:"lookup_route_dispatch_authorization_modeled"`
	RouteEnablementLookupRouteDispatchAuthorizationBoundaryReady bool   `json:"route_enablement_lookup_route_dispatch_authorization_boundary_ready"`
	OpaqueRouteEnablementReceipt                                 bool   `json:"opaque_route_enablement_receipt"`
	KDESafeRedactedStatusOnly                                    bool   `json:"kde_safe_redacted_status_only"`
	ReceiptPresent                                               bool   `json:"receipt_present"`
	ReceiptAccepted                                              bool   `json:"receipt_accepted"`
	ReceiptConsumed                                              bool   `json:"receipt_consumed"`
	AcceptanceAuthorized                                         bool   `json:"acceptance_authorized"`
	RouteEnablementAccepted                                      bool   `json:"route_enablement_accepted"`
	LookupRouteEnablementGranted                                 bool   `json:"lookup_route_enablement_granted"`
	LookupRouteGrantAuthorized                                   bool   `json:"lookup_route_grant_authorized"`
	LookupRouteEnabled                                           bool   `json:"lookup_route_enabled"`
	LookupRouteDispatchAuthorized                                bool   `json:"lookup_route_dispatch_authorized"`
	LookupRouteDispatchAuthorizationPassed                       bool   `json:"lookup_route_dispatch_authorization_passed"`
	LookupRouteDispatchCallable                                  bool   `json:"lookup_route_dispatch_callable"`
	StorageWriteEnabled                                          bool   `json:"storage_write_enabled"`
	RedactedSummaryPersisted                                     bool   `json:"redacted_summary_persisted"`
	KDEStatusPersisted                                           bool   `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted                                  bool   `json:"runtime_diagnostics_persisted"`
	RawResultExposed                                             bool   `json:"raw_result_exposed"`
	UserVisible                                                  bool   `json:"user_visible"`
	ReviewOnly                                                   bool   `json:"review_only"`
	RuntimeOwned                                                 bool   `json:"runtime_owned"`
	GoRuntimeBacked                                              bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                                               bool   `json:"kde_policy_owner"`
	SideEffectsDisabled                                          bool   `json:"side_effects_disabled"`
	HostRootModified                                             bool   `json:"host_root_modified"`
	InternalDetailsExposed                                       bool   `json:"internal_details_exposed"`
	LookupRouteDispatchAuthorizationStatus                       string `json:"lookup_route_dispatch_authorization_status"`
	NextRequirement                                              string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditCheckCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

type routeEnablementLookupRouteDispatchAuthorizationAuditSourceSet struct {
	CurrentMainline                             string
	RouteEnablementLookupRouteDispatchGateAudit string
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditPreview{}, err
	}
	sources := routeEnablementLookupRouteDispatchAuthorizationAuditSources(root)
	mainlineReady := routeEnablementLookupRouteDispatchAuthorizationAuditMainlineReady(sources.CurrentMainline)
	dispatchGateReady := routeEnablementLookupRouteDispatchAuthorizationAuditDispatchGateReady(sources.RouteEnablementLookupRouteDispatchGateAudit)
	items := routeEnablementLookupRouteDispatchAuthorizationAuditItems(mainlineReady, dispatchGateReady)
	readyItemCount := routeEnablementLookupRouteDispatchAuthorizationAuditReadyCount(items)
	ready := mainlineReady && dispatchGateReady && readyItemCount == len(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditPreview{
		Version:                 version,
		SchemaVersion:           "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_authorization_audit.v1",
		RequestType:             "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-authorization-audit-preview",
		AuditType:               "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-authorization-audit",
		Source:                  "docs/xnix-current-mainline.md+route-enable-lookup-route-dispatch-gate-audit",
		AuditDecision:           "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-authorization-audit-blocked",
		CurrentMainlineConsumed: mainlineReady,
		RouteEnablementLookupRouteDispatchGateConsumed:               dispatchGateReady,
		RouteEnablementLookupRouteDispatchGateReady:                  dispatchGateReady,
		LookupRouteDispatchAuthorizationRequired:                     true,
		LookupRouteDispatchAuthorizationModeled:                      true,
		LookupRouteDispatchAuthorizationReady:                        ready,
		RouteEnablementLookupRouteDispatchAuthorizationBoundaryReady: ready,
		OpaqueRouteEnablementReceipt:                                 dispatchGateReady,
		KDESafeRedactedStatusOnly:                                    ready,
		CompatibilityCenterLookupRouteDispatchAuthorizationModeled:   ready,
		RuntimeDiagnosticsLookupRouteDispatchAuthorizationModeled:    ready,
		DispatchAuthorizationItemCount:                               len(items),
		RequiredDispatchAuthorizationItemCount:                       len(items),
		ReadyDispatchAuthorizationItemCount:                          readyItemCount,
		MissingDispatchAuthorizationItemCount:                        len(items) - readyItemCount,
		CompatibilityCenterDispatchAuthorizationItemCount:            routeEnablementLookupRouteDispatchAuthorizationAuditSurfaceCount(items, "compatibility-center"),
		RuntimeDiagnosticsDispatchAuthorizationItemCount:             routeEnablementLookupRouteDispatchAuthorizationAuditSurfaceCount(items, "runtime-diagnostics"),
		DispatchAuthorizationItems:                                   items,
		DispatchAuthorizationItemIDs:                                 routeEnablementLookupRouteDispatchAuthorizationAuditItemIDs(items),
		RuntimeOwned:                                                 true,
		GoRuntimeBacked:                                              true,
		KDEPolicyOwner:                                               false,
		BlockedActions: []string{
			"receipt-present",
			"receipt-accept",
			"receipt-consume",
			"lookup-route-dispatch-authorization",
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
			"Add a route enablement lookup route dispatch dry-run execution gate audit before any authorized dispatch can execute a dry-run.",
			"Keep lookup route dispatch authorization modeled, redacted, and non-authorizing until production ownership is explicitly authorized.",
		},
		DesktopSafeSummary: "KDE-safe redacted route enablement lookup route dispatch authorization is audited for Compatibility Center and Runtime diagnostics while receipts, route grants, lookup routes, consumers, dispatch, persistence, production, backend, and host side effects remain disabled.",
	}
	if ready {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-authorization-audit-ready-route-disabled-dispatch-disabled"
	}
	preview.Checks = routeEnablementLookupRouteDispatchAuthorizationAuditChecks(preview)
	preview.CheckIDs = routeEnablementLookupRouteDispatchAuthorizationAuditCheckIDs(preview.Checks)
	preview.Counts = routeEnablementLookupRouteDispatchAuthorizationAuditCountsFor(preview.Checks)
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch authorization audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditPreview{}, err
	}
	return preview, nil
}

func routeEnablementLookupRouteDispatchAuthorizationAuditSources(root string) routeEnablementLookupRouteDispatchAuthorizationAuditSourceSet {
	return routeEnablementLookupRouteDispatchAuthorizationAuditSourceSet{
		CurrentMainline: productionAuthorizationReadSources(root, []string{"docs/xnix-current-mainline.md"}),
		RouteEnablementLookupRouteDispatchGateAudit: productionAuthorizationReadSources(root, []string{
			"internal/runtime/owner/route_enablement_lookup_route_dispatch_gate_audit_preview.go",
			"internal/runtime/owner/route_enablement_lookup_route_dispatch_gate_audit_preview_test.go",
		}),
	}
}

func routeEnablementLookupRouteDispatchAuthorizationAuditMainlineReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"route enablement lookup route dispatch authorization audit preview",
		"route enablement lookup route dispatch gate audit preview",
		"without accepting or consuming receipts",
	})
}

func routeEnablementLookupRouteDispatchAuthorizationAuditDispatchGateReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_gate_audit.v1",
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-gate-audit-ready-route-disabled-dispatch-disabled",
		"LookupRouteDispatchGateReady",
		"RouteEnablementLookupRouteDispatchGateBoundaryReady",
	})
}

func routeEnablementLookupRouteDispatchAuthorizationAuditItems(mainlineReady, dispatchGateReady bool) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditItem {
	actionKinds := []string{"review", "renew", "open-compatibility-center", "dismiss", "support-info"}
	items := make([]ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditItem, 0, len(actionKinds)*2)
	for _, actionKind := range actionKinds {
		for _, surfaceKind := range []string{"compatibility-center", "runtime-diagnostics"} {
			ready := mainlineReady && dispatchGateReady
			status := "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-authorization-evidence"
			if ready {
				status = "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-authorization-modeled-dispatch-disabled-route-disabled"
			}
			items = append(items, ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditItem{
				ID:                      actionKind + "-route-enablement-lookup-route-dispatch-authorization-" + surfaceKind,
				ActionKind:              actionKind,
				SurfaceKind:             surfaceKind,
				StatusConsumerKind:      "kde-safe-redacted-route-enablement-lookup-route-dispatch-authorization",
				OpaqueReceiptID:         "opaque-route-enablement-lookup-route-dispatch-authorization-" + surfaceKind + "-" + actionKind,
				EvidencePresent:         ready,
				CurrentMainlineConsumed: mainlineReady,
				RouteEnablementLookupRouteDispatchGateConsumed:               dispatchGateReady,
				RouteEnablementLookupRouteDispatchGateReady:                  dispatchGateReady,
				LookupRouteDispatchAuthorizationModeled:                      ready,
				RouteEnablementLookupRouteDispatchAuthorizationBoundaryReady: ready,
				OpaqueRouteEnablementReceipt:                                 dispatchGateReady,
				KDESafeRedactedStatusOnly:                                    ready,
				UserVisible:                                                  ready,
				ReviewOnly:                                                   ready,
				RuntimeOwned:                                                 true,
				GoRuntimeBacked:                                              true,
				KDEPolicyOwner:                                               false,
				SideEffectsDisabled:                                          true,
				LookupRouteDispatchAuthorizationStatus:                       status,
				NextRequirement:                                              "Require a separate lookup route dispatch dry-run execution gate audit before any authorized dispatch can execute a dry-run.",
			})
		}
	}
	return items
}

func routeEnablementLookupRouteDispatchAuthorizationAuditReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent && item.LookupRouteDispatchAuthorizationModeled && item.RouteEnablementLookupRouteDispatchAuthorizationBoundaryReady && item.KDESafeRedactedStatusOnly && !item.ReceiptPresent && !item.ReceiptAccepted && !item.ReceiptConsumed && !item.AcceptanceAuthorized && !item.RouteEnablementAccepted && !item.LookupRouteEnablementGranted && !item.LookupRouteGrantAuthorized && !item.LookupRouteEnabled && !item.LookupRouteDispatchAuthorized && !item.LookupRouteDispatchAuthorizationPassed && !item.LookupRouteDispatchCallable && !item.StorageWriteEnabled && !item.RawResultExposed && !item.HostRootModified && !item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func routeEnablementLookupRouteDispatchAuthorizationAuditSurfaceCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditItem, surfaceKind string) int {
	count := 0
	for _, item := range items {
		if item.SurfaceKind == surfaceKind {
			count++
		}
	}
	return count
}

func routeEnablementLookupRouteDispatchAuthorizationAuditItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func routeEnablementLookupRouteDispatchAuthorizationAuditChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditCheck{
		routeEnablementLookupRouteDispatchAuthorizationAuditCheck("current-mainline-consumed", preview.CurrentMainlineConsumed, "Xnix current mainline names the result consumer projection route enablement lookup route dispatch authorization audit preview."),
		routeEnablementLookupRouteDispatchAuthorizationAuditCheck("route-enablement-lookup-route-dispatch-gate-consumed", preview.RouteEnablementLookupRouteDispatchGateConsumed && preview.RouteEnablementLookupRouteDispatchGateReady, "The lookup route dispatch authorization audit consumes the ready route enablement lookup route dispatch gate audit."),
		routeEnablementLookupRouteDispatchAuthorizationAuditCheck("lookup-route-dispatch-authorization-modeled", preview.LookupRouteDispatchAuthorizationRequired && preview.LookupRouteDispatchAuthorizationModeled, "Lookup route dispatch authorization is modeled without enabling a route."),
		routeEnablementLookupRouteDispatchAuthorizationAuditCheck("compatibility-center-and-runtime-lookup-route-dispatch-authorizations-modeled", preview.CompatibilityCenterLookupRouteDispatchAuthorizationModeled && preview.RuntimeDiagnosticsLookupRouteDispatchAuthorizationModeled, "Compatibility Center and Runtime diagnostics lookup route dispatch authorization boundaries are modeled."),
		routeEnablementLookupRouteDispatchAuthorizationAuditCheck("ten-dispatch-authorization-items-ready-dispatch-disabled", preview.DispatchAuthorizationItemCount == 10 && preview.ReadyDispatchAuthorizationItemCount == 10 && preview.MissingDispatchAuthorizationItemCount == 0, "Ten KDE-safe lookup route dispatch authorization items are ready while route and dispatch authorizations remain disabled."),
		routeEnablementLookupRouteDispatchAuthorizationAuditCheck("lookup-route-dispatch-authorization-and-persistence-disabled", !preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.ReceiptConsumed && !preview.AcceptanceAuthorized && !preview.RouteEnablementAccepted && !preview.LookupRouteEnablementGranted && !preview.LookupRouteGrantAuthorized && !preview.LookupRouteEnabled && !preview.LookupRouteDispatchAuthorized && !preview.LookupRouteDispatchAuthorizationPassed && !preview.LookupRouteDispatchCallable && !preview.StorageWriteEnabled && !preview.StatusPersistenceWriteEnabled && !preview.RawResultExposed, "Lookup route dispatch authorization, lookup grants, lookup routes, storage, persistence, and raw result exposure remain disabled."),
		routeEnablementLookupRouteDispatchAuthorizationAuditCheck("consumer-routes-dispatch-and-support-disabled", !preview.ConsumerEnablementAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled && !preview.LookupRouteEnabled && !preview.LookupRouteDispatchAuthorized && !preview.DispatchDryRunExecuted && !preview.SupportBundleExported && !preview.SupportCaseCreated, "Consumer enablement, lookup routes, dispatch, and support side effects remain disabled."),
		routeEnablementLookupRouteDispatchAuthorizationAuditCheck("production-and-host-boundary-closed", !preview.ProductionBusClaimed && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.BackendLaunchEnabled && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired, "Production, backend, network, privileged container, and host mutation gates remain closed."),
	}
}

func routeEnablementLookupRouteDispatchAuthorizationAuditCheck(id string, ok bool, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditCheck {
	status := "blocked"
	if ok {
		status = "pass"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditCheck{ID: id, Status: status, Summary: summary}
}

func routeEnablementLookupRouteDispatchAuthorizationAuditCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func routeEnablementLookupRouteDispatchAuthorizationAuditCountsFor(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditCheckCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditCheckCounts{Total: len(checks)}
	for _, check := range checks {
		if check.Status == "pass" {
			counts.Passed++
		} else {
			counts.Blocked++
		}
	}
	return counts
}
