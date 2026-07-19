package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditPreview struct {
	Version                                      string                                                                                                                                                                                                                           `json:"version"`
	SchemaVersion                                string                                                                                                                                                                                                                           `json:"schema_version"`
	RequestType                                  string                                                                                                                                                                                                                           `json:"request_type"`
	AuditType                                    string                                                                                                                                                                                                                           `json:"audit_type"`
	Source                                       string                                                                                                                                                                                                                           `json:"source"`
	AuditDecision                                string                                                                                                                                                                                                                           `json:"audit_decision"`
	CurrentMainlineConsumed                      bool                                                                                                                                                                                                                             `json:"current_mainline_consumed"`
	RouteEnablementAcceptedReceiptGateConsumed   bool                                                                                                                                                                                                                             `json:"route_enablement_accepted_receipt_gate_consumed"`
	RouteEnablementAcceptedReceiptGateReady      bool                                                                                                                                                                                                                             `json:"route_enablement_accepted_receipt_gate_ready"`
	LookupRouteGrantRequired                     bool                                                                                                                                                                                                                             `json:"lookup_route_grant_required"`
	LookupRouteGrantModeled                      bool                                                                                                                                                                                                                             `json:"lookup_route_grant_modeled"`
	LookupRouteGrantReady                        bool                                                                                                                                                                                                                             `json:"lookup_route_grant_ready"`
	RouteEnablementLookupRouteGrantBoundaryReady bool                                                                                                                                                                                                                             `json:"route_enablement_lookup_route_grant_boundary_ready"`
	OpaqueRouteEnablementReceipt                 bool                                                                                                                                                                                                                             `json:"opaque_route_enablement_receipt"`
	KDESafeRedactedStatusOnly                    bool                                                                                                                                                                                                                             `json:"kde_safe_redacted_status_only"`
	CompatibilityCenterLookupRouteGrantModeled   bool                                                                                                                                                                                                                             `json:"compatibility_center_lookup_route_grant_modeled"`
	RuntimeDiagnosticsLookupRouteGrantModeled    bool                                                                                                                                                                                                                             `json:"runtime_diagnostics_lookup_route_grant_modeled"`
	ReceiptPresent                               bool                                                                                                                                                                                                                             `json:"receipt_present"`
	ReceiptAccepted                              bool                                                                                                                                                                                                                             `json:"receipt_accepted"`
	ReceiptConsumed                              bool                                                                                                                                                                                                                             `json:"receipt_consumed"`
	AcceptanceAuthorized                         bool                                                                                                                                                                                                                             `json:"acceptance_authorized"`
	RouteEnablementAccepted                      bool                                                                                                                                                                                                                             `json:"route_enablement_accepted"`
	LookupRouteEnablementGranted                 bool                                                                                                                                                                                                                             `json:"lookup_route_enablement_granted"`
	LookupRouteGrantAuthorized                   bool                                                                                                                                                                                                                             `json:"lookup_route_grant_authorized"`
	LookupRouteEnabled                           bool                                                                                                                                                                                                                             `json:"lookup_route_enabled"`
	StorageRootPolicyGrantAuthorized             bool                                                                                                                                                                                                                             `json:"storage_root_policy_grant_authorized"`
	RecordWriterCallAuthorized                   bool                                                                                                                                                                                                                             `json:"record_writer_call_authorized"`
	WriterCallable                               bool                                                                                                                                                                                                                             `json:"writer_callable"`
	StorageRootResolved                          bool                                                                                                                                                                                                                             `json:"storage_root_resolved"`
	StorageWriteEnabled                          bool                                                                                                                                                                                                                             `json:"storage_write_enabled"`
	StatusPersistenceWriteEnabled                bool                                                                                                                                                                                                                             `json:"status_persistence_write_enabled"`
	GrantItemCount                               int                                                                                                                                                                                                                              `json:"grant_item_count"`
	RequiredGrantItemCount                       int                                                                                                                                                                                                                              `json:"required_grant_item_count"`
	ReadyGrantItemCount                          int                                                                                                                                                                                                                              `json:"ready_grant_item_count"`
	MissingGrantItemCount                        int                                                                                                                                                                                                                              `json:"missing_grant_item_count"`
	GrantedRouteItemCount                        int                                                                                                                                                                                                                              `json:"granted_route_item_count"`
	EnabledRouteItemCount                        int                                                                                                                                                                                                                              `json:"enabled_route_item_count"`
	PersistedGrantItemCount                      int                                                                                                                                                                                                                              `json:"persisted_grant_item_count"`
	RawExposedGrantItemCount                     int                                                                                                                                                                                                                              `json:"raw_exposed_grant_item_count"`
	SideEffectGrantItemCount                     int                                                                                                                                                                                                                              `json:"side_effect_grant_item_count"`
	CompatibilityCenterGrantItemCount            int                                                                                                                                                                                                                              `json:"compatibility_center_grant_item_count"`
	RuntimeDiagnosticsGrantItemCount             int                                                                                                                                                                                                                              `json:"runtime_diagnostics_grant_item_count"`
	GrantItems                                   []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditItem      `json:"grant_items"`
	GrantItemIDs                                 []string                                                                                                                                                                                                                         `json:"grant_item_ids"`
	Checks                                       []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditCheck     `json:"checks"`
	CheckIDs                                     []string                                                                                                                                                                                                                         `json:"check_ids"`
	Counts                                       ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditCheckCounts `json:"counts"`
	ConsumerDryRunLookupResultBoundaryAuthorized bool                                                                                                                                                                                                                             `json:"consumer_dry_run_lookup_result_boundary_authorized"`
	ConsumerEnablementAuthorized                 bool                                                                                                                                                                                                                             `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                           bool                                                                                                                                                                                                                             `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                       bool                                                                                                                                                                                                                             `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                        bool                                                                                                                                                                                                                             `json:"lookup_route_authorized"`
	LookupRouteEnablementAuthorized              bool                                                                                                                                                                                                                             `json:"lookup_route_enablement_authorized"`
	OpaqueLookupEnabled                          bool                                                                                                                                                                                                                             `json:"opaque_lookup_enabled"`
	RedactedSummaryPersisted                     bool                                                                                                                                                                                                                             `json:"redacted_summary_persisted"`
	KDEStatusPersisted                           bool                                                                                                                                                                                                                             `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted                  bool                                                                                                                                                                                                                             `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                        bool                                                                                                                                                                                                                             `json:"dry_run_result_persisted"`
	RawResultExposed                             bool                                                                                                                                                                                                                             `json:"raw_result_exposed"`
	DispatchDryRunExecuted                       bool                                                                                                                                                                                                                             `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled                 bool                                                                                                                                                                                                                             `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled                 bool                                                                                                                                                                                                                             `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                         bool                                                                                                                                                                                                                             `json:"portal_request_created"`
	NotificationActionEnabled                    bool                                                                                                                                                                                                                             `json:"notification_action_enabled"`
	CompatibilityCenterOpened                    bool                                                                                                                                                                                                                             `json:"compatibility_center_opened"`
	SupportBundleExported                        bool                                                                                                                                                                                                                             `json:"support_bundle_exported"`
	SupportCaseCreated                           bool                                                                                                                                                                                                                             `json:"support_case_created"`
	RuntimeOwned                                 bool                                                                                                                                                                                                                             `json:"runtime_owned"`
	GoRuntimeBacked                              bool                                                                                                                                                                                                                             `json:"go_runtime_backed"`
	KDEPolicyOwner                               bool                                                                                                                                                                                                                             `json:"kde_policy_owner"`
	ProductionReadiness                          bool                                                                                                                                                                                                                             `json:"production_readiness"`
	ProductionOwnershipReady                     bool                                                                                                                                                                                                                             `json:"production_ownership_ready"`
	SystemServiceStarted                         bool                                                                                                                                                                                                                             `json:"system_service_started"`
	SessionBusClaimed                            bool                                                                                                                                                                                                                             `json:"session_bus_claimed"`
	ProductionBusClaimed                         bool                                                                                                                                                                                                                             `json:"production_bus_claimed"`
	WriteMethodsEnabled                          bool                                                                                                                                                                                                                             `json:"write_methods_enabled"`
	RuntimeWritesEnabled                         bool                                                                                                                                                                                                                             `json:"runtime_writes_enabled"`
	DesktopFilesWritten                          bool                                                                                                                                                                                                                             `json:"desktop_files_written"`
	KDEConfigurationWritten                      bool                                                                                                                                                                                                                             `json:"kde_configuration_written"`
	PortalCallExecuted                           bool                                                                                                                                                                                                                             `json:"portal_call_executed"`
	AdapterInvocationEnabled                     bool                                                                                                                                                                                                                             `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                         bool                                                                                                                                                                                                                             `json:"backend_launch_enabled"`
	BackendProcessStarted                        bool                                                                                                                                                                                                                             `json:"backend_process_started"`
	NetworkRequired                              bool                                                                                                                                                                                                                             `json:"network_required"`
	HostRootModified                             bool                                                                                                                                                                                                                             `json:"host_root_modified"`
	PrivilegedContainerRequired                  bool                                                                                                                                                                                                                             `json:"privileged_container_required"`
	StateRootPathExposed                         bool                                                                                                                                                                                                                             `json:"state_root_path_exposed"`
	FilePathsExposed                             bool                                                                                                                                                                                                                             `json:"file_paths_exposed"`
	FileContentRead                              bool                                                                                                                                                                                                                             `json:"file_content_read"`
	RawCommandExposed                            bool                                                                                                                                                                                                                             `json:"raw_command_exposed"`
	RawExecutableExposed                         bool                                                                                                                                                                                                                             `json:"raw_executable_exposed"`
	BackendDetailsExposed                        bool                                                                                                                                                                                                                             `json:"backend_details_exposed"`
	BlockedActions                               []string                                                                                                                                                                                                                         `json:"blocked_actions"`
	NextRequirements                             []string                                                                                                                                                                                                                         `json:"next_requirements"`
	DesktopSafeSummary                           string                                                                                                                                                                                                                           `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditItem struct {
	ID                                           string `json:"id"`
	ActionKind                                   string `json:"action_kind"`
	SurfaceKind                                  string `json:"surface_kind"`
	StatusConsumerKind                           string `json:"status_consumer_kind"`
	OpaqueReceiptID                              string `json:"opaque_receipt_id"`
	EvidencePresent                              bool   `json:"evidence_present"`
	CurrentMainlineConsumed                      bool   `json:"current_mainline_consumed"`
	RouteEnablementAcceptedReceiptGateConsumed   bool   `json:"route_enablement_accepted_receipt_gate_consumed"`
	RouteEnablementAcceptedReceiptGateReady      bool   `json:"route_enablement_accepted_receipt_gate_ready"`
	LookupRouteGrantModeled                      bool   `json:"lookup_route_grant_modeled"`
	RouteEnablementLookupRouteGrantBoundaryReady bool   `json:"route_enablement_lookup_route_grant_boundary_ready"`
	OpaqueRouteEnablementReceipt                 bool   `json:"opaque_route_enablement_receipt"`
	KDESafeRedactedStatusOnly                    bool   `json:"kde_safe_redacted_status_only"`
	ReceiptPresent                               bool   `json:"receipt_present"`
	ReceiptAccepted                              bool   `json:"receipt_accepted"`
	ReceiptConsumed                              bool   `json:"receipt_consumed"`
	AcceptanceAuthorized                         bool   `json:"acceptance_authorized"`
	RouteEnablementAccepted                      bool   `json:"route_enablement_accepted"`
	LookupRouteEnablementGranted                 bool   `json:"lookup_route_enablement_granted"`
	LookupRouteGrantAuthorized                   bool   `json:"lookup_route_grant_authorized"`
	LookupRouteEnabled                           bool   `json:"lookup_route_enabled"`
	StorageWriteEnabled                          bool   `json:"storage_write_enabled"`
	RedactedSummaryPersisted                     bool   `json:"redacted_summary_persisted"`
	KDEStatusPersisted                           bool   `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted                  bool   `json:"runtime_diagnostics_persisted"`
	RawResultExposed                             bool   `json:"raw_result_exposed"`
	UserVisible                                  bool   `json:"user_visible"`
	ReviewOnly                                   bool   `json:"review_only"`
	RuntimeOwned                                 bool   `json:"runtime_owned"`
	GoRuntimeBacked                              bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                               bool   `json:"kde_policy_owner"`
	SideEffectsDisabled                          bool   `json:"side_effects_disabled"`
	HostRootModified                             bool   `json:"host_root_modified"`
	InternalDetailsExposed                       bool   `json:"internal_details_exposed"`
	LookupRouteGrantStatus                       string `json:"lookup_route_grant_status"`
	NextRequirement                              string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditCheckCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

type routeEnablementLookupRouteGrantAuditSourceSet struct {
	CurrentMainline                         string
	RouteEnablementAcceptedReceiptGateAudit string
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditPreview{}, err
	}
	sources := routeEnablementLookupRouteGrantAuditSources(root)
	mainlineReady := routeEnablementLookupRouteGrantAuditMainlineReady(sources.CurrentMainline)
	acceptedGateReady := routeEnablementLookupRouteGrantAuditAcceptedGateReady(sources.RouteEnablementAcceptedReceiptGateAudit)
	items := routeEnablementLookupRouteGrantAuditItems(mainlineReady, acceptedGateReady)
	readyItemCount := routeEnablementLookupRouteGrantAuditReadyCount(items)
	ready := mainlineReady && acceptedGateReady && readyItemCount == len(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditPreview{
		Version:                 version,
		SchemaVersion:           "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_grant_audit.v1",
		RequestType:             "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-grant-audit-preview",
		AuditType:               "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-grant-audit",
		Source:                  "docs/xnix-current-mainline.md+route-enable-accepted-receipt-gate-audit",
		AuditDecision:           "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-grant-audit-blocked",
		CurrentMainlineConsumed: mainlineReady,
		RouteEnablementAcceptedReceiptGateConsumed:   acceptedGateReady,
		RouteEnablementAcceptedReceiptGateReady:      acceptedGateReady,
		LookupRouteGrantRequired:                     true,
		LookupRouteGrantModeled:                      true,
		LookupRouteGrantReady:                        ready,
		RouteEnablementLookupRouteGrantBoundaryReady: ready,
		OpaqueRouteEnablementReceipt:                 acceptedGateReady,
		KDESafeRedactedStatusOnly:                    ready,
		CompatibilityCenterLookupRouteGrantModeled:   ready,
		RuntimeDiagnosticsLookupRouteGrantModeled:    ready,
		GrantItemCount:                               len(items),
		RequiredGrantItemCount:                       len(items),
		ReadyGrantItemCount:                          readyItemCount,
		MissingGrantItemCount:                        len(items) - readyItemCount,
		CompatibilityCenterGrantItemCount:            routeEnablementLookupRouteGrantAuditSurfaceCount(items, "compatibility-center"),
		RuntimeDiagnosticsGrantItemCount:             routeEnablementLookupRouteGrantAuditSurfaceCount(items, "runtime-diagnostics"),
		GrantItems:                                   items,
		GrantItemIDs:                                 routeEnablementLookupRouteGrantAuditItemIDs(items),
		RuntimeOwned:                                 true,
		GoRuntimeBacked:                              true,
		KDEPolicyOwner:                               false,
		BlockedActions: []string{
			"receipt-present",
			"receipt-accept",
			"receipt-consume",
			"lookup-route-grant",
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
			"Add a route enablement lookup route enablement audit before any lookup route can be enabled.",
			"Keep lookup route grants modeled, redacted, and disabled until production ownership is explicitly authorized.",
		},
		DesktopSafeSummary: "KDE-safe redacted route enablement lookup route grants are audited for Compatibility Center and Runtime diagnostics while receipts, route grants, lookup routes, consumers, dispatch, persistence, production, backend, and host side effects remain disabled.",
	}
	if ready {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-grant-audit-ready-grant-disabled-routes-disabled"
	}
	preview.Checks = routeEnablementLookupRouteGrantAuditChecks(preview)
	preview.CheckIDs = routeEnablementLookupRouteGrantAuditCheckIDs(preview.Checks)
	preview.Counts = routeEnablementLookupRouteGrantAuditCountsFor(preview.Checks)
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route grant audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditPreview{}, err
	}
	return preview, nil
}

func routeEnablementLookupRouteGrantAuditSources(root string) routeEnablementLookupRouteGrantAuditSourceSet {
	return routeEnablementLookupRouteGrantAuditSourceSet{
		CurrentMainline: productionAuthorizationReadSources(root, []string{"docs/xnix-current-mainline.md"}),
		RouteEnablementAcceptedReceiptGateAudit: productionAuthorizationReadSources(root, []string{
			"internal/runtime/owner/route_enablement_accepted_receipt_gate_audit_preview.go",
			"internal/runtime/owner/route_enablement_accepted_receipt_gate_audit_preview_test.go",
		}),
	}
}

func routeEnablementLookupRouteGrantAuditMainlineReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"route enablement lookup route grant audit preview",
		"route enablement accepted receipt gate audit preview",
		"without accepting or consuming receipts",
	})
}

func routeEnablementLookupRouteGrantAuditAcceptedGateReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_accepted_receipt_gate_audit.v1",
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-accepted-receipt-gate-audit-ready-gate-disabled-routes-disabled",
		"AcceptedReceiptGateReady",
		"RouteEnablementAcceptedReceiptBoundaryReady",
	})
}

func routeEnablementLookupRouteGrantAuditItems(mainlineReady, acceptedGateReady bool) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditItem {
	actionKinds := []string{"review", "renew", "open-compatibility-center", "dismiss", "support-info"}
	items := make([]ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditItem, 0, len(actionKinds)*2)
	for _, actionKind := range actionKinds {
		for _, surfaceKind := range []string{"compatibility-center", "runtime-diagnostics"} {
			ready := mainlineReady && acceptedGateReady
			status := "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-grant-evidence"
			if ready {
				status = "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-grant-modeled-grant-disabled-routes-disabled"
			}
			items = append(items, ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditItem{
				ID:                      actionKind + "-route-enablement-lookup-route-grant-" + surfaceKind,
				ActionKind:              actionKind,
				SurfaceKind:             surfaceKind,
				StatusConsumerKind:      "kde-safe-redacted-route-enablement-lookup-route-grant",
				OpaqueReceiptID:         "opaque-route-enablement-lookup-route-grant-" + surfaceKind + "-" + actionKind,
				EvidencePresent:         ready,
				CurrentMainlineConsumed: mainlineReady,
				RouteEnablementAcceptedReceiptGateConsumed:   acceptedGateReady,
				RouteEnablementAcceptedReceiptGateReady:      acceptedGateReady,
				LookupRouteGrantModeled:                      ready,
				RouteEnablementLookupRouteGrantBoundaryReady: ready,
				OpaqueRouteEnablementReceipt:                 acceptedGateReady,
				KDESafeRedactedStatusOnly:                    ready,
				UserVisible:                                  ready,
				ReviewOnly:                                   ready,
				RuntimeOwned:                                 true,
				GoRuntimeBacked:                              true,
				KDEPolicyOwner:                               false,
				SideEffectsDisabled:                          true,
				LookupRouteGrantStatus:                       status,
				NextRequirement:                              "Require a separate lookup route enablement audit before the modeled grant can enable any route.",
			})
		}
	}
	return items
}

func routeEnablementLookupRouteGrantAuditReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent && item.LookupRouteGrantModeled && item.RouteEnablementLookupRouteGrantBoundaryReady && item.KDESafeRedactedStatusOnly && !item.ReceiptPresent && !item.ReceiptAccepted && !item.ReceiptConsumed && !item.AcceptanceAuthorized && !item.RouteEnablementAccepted && !item.LookupRouteEnablementGranted && !item.LookupRouteGrantAuthorized && !item.LookupRouteEnabled && !item.StorageWriteEnabled && !item.RawResultExposed && !item.HostRootModified && !item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func routeEnablementLookupRouteGrantAuditSurfaceCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditItem, surfaceKind string) int {
	count := 0
	for _, item := range items {
		if item.SurfaceKind == surfaceKind {
			count++
		}
	}
	return count
}

func routeEnablementLookupRouteGrantAuditItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func routeEnablementLookupRouteGrantAuditChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditCheck{
		routeEnablementLookupRouteGrantAuditCheck("current-mainline-consumed", preview.CurrentMainlineConsumed, "Xnix current mainline names the result consumer projection route enablement lookup route grant audit preview."),
		routeEnablementLookupRouteGrantAuditCheck("route-enablement-accepted-receipt-gate-consumed", preview.RouteEnablementAcceptedReceiptGateConsumed && preview.RouteEnablementAcceptedReceiptGateReady, "The lookup route grant audit consumes the ready route enablement accepted receipt gate audit."),
		routeEnablementLookupRouteGrantAuditCheck("lookup-route-grant-modeled", preview.LookupRouteGrantRequired && preview.LookupRouteGrantModeled, "Lookup route grant is modeled without granting a route."),
		routeEnablementLookupRouteGrantAuditCheck("compatibility-center-and-runtime-lookup-route-grants-modeled", preview.CompatibilityCenterLookupRouteGrantModeled && preview.RuntimeDiagnosticsLookupRouteGrantModeled, "Compatibility Center and Runtime diagnostics lookup route grant boundaries are modeled."),
		routeEnablementLookupRouteGrantAuditCheck("ten-grant-items-ready-grant-disabled", preview.GrantItemCount == 10 && preview.ReadyGrantItemCount == 10 && preview.MissingGrantItemCount == 0, "Ten KDE-safe lookup route grant items are ready while grant and route gates remain disabled."),
		routeEnablementLookupRouteGrantAuditCheck("lookup-route-grants-and-persistence-disabled", !preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.ReceiptConsumed && !preview.AcceptanceAuthorized && !preview.RouteEnablementAccepted && !preview.LookupRouteEnablementGranted && !preview.LookupRouteGrantAuthorized && !preview.LookupRouteEnabled && !preview.StorageWriteEnabled && !preview.StatusPersistenceWriteEnabled && !preview.RawResultExposed, "Lookup route grants, route enablement, lookup routes, storage, persistence, and raw result exposure remain disabled."),
		routeEnablementLookupRouteGrantAuditCheck("consumer-routes-dispatch-and-support-disabled", !preview.ConsumerEnablementAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled && !preview.LookupRouteEnabled && !preview.DispatchDryRunExecuted && !preview.SupportBundleExported && !preview.SupportCaseCreated, "Consumer enablement, lookup routes, dispatch, and support side effects remain disabled."),
		routeEnablementLookupRouteGrantAuditCheck("production-and-host-boundary-closed", !preview.ProductionBusClaimed && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.BackendLaunchEnabled && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired, "Production, backend, network, privileged container, and host mutation gates remain closed."),
	}
}

func routeEnablementLookupRouteGrantAuditCheck(id string, ok bool, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditCheck {
	status := "blocked"
	if ok {
		status = "pass"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditCheck{ID: id, Status: status, Summary: summary}
}

func routeEnablementLookupRouteGrantAuditCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func routeEnablementLookupRouteGrantAuditCountsFor(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditCheckCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditCheckCounts{Total: len(checks)}
	for _, check := range checks {
		if check.Status == "pass" {
			counts.Passed++
		} else {
			counts.Blocked++
		}
	}
	return counts
}
