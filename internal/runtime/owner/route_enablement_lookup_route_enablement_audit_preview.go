package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditPreview struct {
	Version                                           string                                                                                                                                                                                                                                `json:"version"`
	SchemaVersion                                     string                                                                                                                                                                                                                                `json:"schema_version"`
	RequestType                                       string                                                                                                                                                                                                                                `json:"request_type"`
	AuditType                                         string                                                                                                                                                                                                                                `json:"audit_type"`
	Source                                            string                                                                                                                                                                                                                                `json:"source"`
	AuditDecision                                     string                                                                                                                                                                                                                                `json:"audit_decision"`
	CurrentMainlineConsumed                           bool                                                                                                                                                                                                                                  `json:"current_mainline_consumed"`
	RouteEnablementLookupRouteGrantConsumed           bool                                                                                                                                                                                                                                  `json:"route_enablement_lookup_route_grant_consumed"`
	RouteEnablementLookupRouteGrantReady              bool                                                                                                                                                                                                                                  `json:"route_enablement_lookup_route_grant_ready"`
	LookupRouteEnablementRequired                     bool                                                                                                                                                                                                                                  `json:"lookup_route_enablement_required"`
	LookupRouteEnablementModeled                      bool                                                                                                                                                                                                                                  `json:"lookup_route_enablement_modeled"`
	LookupRouteEnablementReady                        bool                                                                                                                                                                                                                                  `json:"lookup_route_enablement_ready"`
	RouteEnablementLookupRouteEnablementBoundaryReady bool                                                                                                                                                                                                                                  `json:"route_enablement_lookup_route_enablement_boundary_ready"`
	OpaqueRouteEnablementReceipt                      bool                                                                                                                                                                                                                                  `json:"opaque_route_enablement_receipt"`
	KDESafeRedactedStatusOnly                         bool                                                                                                                                                                                                                                  `json:"kde_safe_redacted_status_only"`
	CompatibilityCenterLookupRouteEnablementModeled   bool                                                                                                                                                                                                                                  `json:"compatibility_center_lookup_route_enablement_modeled"`
	RuntimeDiagnosticsLookupRouteEnablementModeled    bool                                                                                                                                                                                                                                  `json:"runtime_diagnostics_lookup_route_enablement_modeled"`
	ReceiptPresent                                    bool                                                                                                                                                                                                                                  `json:"receipt_present"`
	ReceiptAccepted                                   bool                                                                                                                                                                                                                                  `json:"receipt_accepted"`
	ReceiptConsumed                                   bool                                                                                                                                                                                                                                  `json:"receipt_consumed"`
	AcceptanceAuthorized                              bool                                                                                                                                                                                                                                  `json:"acceptance_authorized"`
	RouteEnablementAccepted                           bool                                                                                                                                                                                                                                  `json:"route_enablement_accepted"`
	LookupRouteEnablementGranted                      bool                                                                                                                                                                                                                                  `json:"lookup_route_enablement_granted"`
	LookupRouteGrantAuthorized                        bool                                                                                                                                                                                                                                  `json:"lookup_route_grant_authorized"`
	LookupRouteEnabled                                bool                                                                                                                                                                                                                                  `json:"lookup_route_enabled"`
	StorageRootPolicyGrantAuthorized                  bool                                                                                                                                                                                                                                  `json:"storage_root_policy_grant_authorized"`
	RecordWriterCallAuthorized                        bool                                                                                                                                                                                                                                  `json:"record_writer_call_authorized"`
	WriterCallable                                    bool                                                                                                                                                                                                                                  `json:"writer_callable"`
	StorageRootResolved                               bool                                                                                                                                                                                                                                  `json:"storage_root_resolved"`
	StorageWriteEnabled                               bool                                                                                                                                                                                                                                  `json:"storage_write_enabled"`
	StatusPersistenceWriteEnabled                     bool                                                                                                                                                                                                                                  `json:"status_persistence_write_enabled"`
	EnablementItemCount                               int                                                                                                                                                                                                                                   `json:"enablement_item_count"`
	RequiredEnablementItemCount                       int                                                                                                                                                                                                                                   `json:"required_enablement_item_count"`
	ReadyEnablementItemCount                          int                                                                                                                                                                                                                                   `json:"ready_enablement_item_count"`
	MissingEnablementItemCount                        int                                                                                                                                                                                                                                   `json:"missing_enablement_item_count"`
	GrantedRouteItemCount                             int                                                                                                                                                                                                                                   `json:"granted_route_item_count"`
	EnabledRouteItemCount                             int                                                                                                                                                                                                                                   `json:"enabled_route_item_count"`
	PersistedEnablementItemCount                      int                                                                                                                                                                                                                                   `json:"persisted_enablement_item_count"`
	RawExposedEnablementItemCount                     int                                                                                                                                                                                                                                   `json:"raw_exposed_enablement_item_count"`
	SideEffectEnablementItemCount                     int                                                                                                                                                                                                                                   `json:"side_effect_enablement_item_count"`
	CompatibilityCenterEnablementItemCount            int                                                                                                                                                                                                                                   `json:"compatibility_center_enablement_item_count"`
	RuntimeDiagnosticsEnablementItemCount             int                                                                                                                                                                                                                                   `json:"runtime_diagnostics_enablement_item_count"`
	EnablementItems                                   []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditItem      `json:"enablement_items"`
	EnablementItemIDs                                 []string                                                                                                                                                                                                                              `json:"enablement_item_ids"`
	Checks                                            []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditCheck     `json:"checks"`
	CheckIDs                                          []string                                                                                                                                                                                                                              `json:"check_ids"`
	Counts                                            ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditCheckCounts `json:"counts"`
	ConsumerDryRunLookupResultBoundaryAuthorized      bool                                                                                                                                                                                                                                  `json:"consumer_dry_run_lookup_result_boundary_authorized"`
	ConsumerEnablementAuthorized                      bool                                                                                                                                                                                                                                  `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                                bool                                                                                                                                                                                                                                  `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                            bool                                                                                                                                                                                                                                  `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                             bool                                                                                                                                                                                                                                  `json:"lookup_route_authorized"`
	LookupRouteEnablementAuthorized                   bool                                                                                                                                                                                                                                  `json:"lookup_route_enablement_authorized"`
	OpaqueLookupEnabled                               bool                                                                                                                                                                                                                                  `json:"opaque_lookup_enabled"`
	RedactedSummaryPersisted                          bool                                                                                                                                                                                                                                  `json:"redacted_summary_persisted"`
	KDEStatusPersisted                                bool                                                                                                                                                                                                                                  `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted                       bool                                                                                                                                                                                                                                  `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                             bool                                                                                                                                                                                                                                  `json:"dry_run_result_persisted"`
	RawResultExposed                                  bool                                                                                                                                                                                                                                  `json:"raw_result_exposed"`
	DispatchDryRunExecuted                            bool                                                                                                                                                                                                                                  `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled                      bool                                                                                                                                                                                                                                  `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled                      bool                                                                                                                                                                                                                                  `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                              bool                                                                                                                                                                                                                                  `json:"portal_request_created"`
	NotificationActionEnabled                         bool                                                                                                                                                                                                                                  `json:"notification_action_enabled"`
	CompatibilityCenterOpened                         bool                                                                                                                                                                                                                                  `json:"compatibility_center_opened"`
	SupportBundleExported                             bool                                                                                                                                                                                                                                  `json:"support_bundle_exported"`
	SupportCaseCreated                                bool                                                                                                                                                                                                                                  `json:"support_case_created"`
	RuntimeOwned                                      bool                                                                                                                                                                                                                                  `json:"runtime_owned"`
	GoRuntimeBacked                                   bool                                                                                                                                                                                                                                  `json:"go_runtime_backed"`
	KDEPolicyOwner                                    bool                                                                                                                                                                                                                                  `json:"kde_policy_owner"`
	ProductionReadiness                               bool                                                                                                                                                                                                                                  `json:"production_readiness"`
	ProductionOwnershipReady                          bool                                                                                                                                                                                                                                  `json:"production_ownership_ready"`
	SystemServiceStarted                              bool                                                                                                                                                                                                                                  `json:"system_service_started"`
	SessionBusClaimed                                 bool                                                                                                                                                                                                                                  `json:"session_bus_claimed"`
	ProductionBusClaimed                              bool                                                                                                                                                                                                                                  `json:"production_bus_claimed"`
	WriteMethodsEnabled                               bool                                                                                                                                                                                                                                  `json:"write_methods_enabled"`
	RuntimeWritesEnabled                              bool                                                                                                                                                                                                                                  `json:"runtime_writes_enabled"`
	DesktopFilesWritten                               bool                                                                                                                                                                                                                                  `json:"desktop_files_written"`
	KDEConfigurationWritten                           bool                                                                                                                                                                                                                                  `json:"kde_configuration_written"`
	PortalCallExecuted                                bool                                                                                                                                                                                                                                  `json:"portal_call_executed"`
	AdapterInvocationEnabled                          bool                                                                                                                                                                                                                                  `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                              bool                                                                                                                                                                                                                                  `json:"backend_launch_enabled"`
	BackendProcessStarted                             bool                                                                                                                                                                                                                                  `json:"backend_process_started"`
	NetworkRequired                                   bool                                                                                                                                                                                                                                  `json:"network_required"`
	HostRootModified                                  bool                                                                                                                                                                                                                                  `json:"host_root_modified"`
	PrivilegedContainerRequired                       bool                                                                                                                                                                                                                                  `json:"privileged_container_required"`
	StateRootPathExposed                              bool                                                                                                                                                                                                                                  `json:"state_root_path_exposed"`
	FilePathsExposed                                  bool                                                                                                                                                                                                                                  `json:"file_paths_exposed"`
	FileContentRead                                   bool                                                                                                                                                                                                                                  `json:"file_content_read"`
	RawCommandExposed                                 bool                                                                                                                                                                                                                                  `json:"raw_command_exposed"`
	RawExecutableExposed                              bool                                                                                                                                                                                                                                  `json:"raw_executable_exposed"`
	BackendDetailsExposed                             bool                                                                                                                                                                                                                                  `json:"backend_details_exposed"`
	BlockedActions                                    []string                                                                                                                                                                                                                              `json:"blocked_actions"`
	NextRequirements                                  []string                                                                                                                                                                                                                              `json:"next_requirements"`
	DesktopSafeSummary                                string                                                                                                                                                                                                                                `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditItem struct {
	ID                                                string `json:"id"`
	ActionKind                                        string `json:"action_kind"`
	SurfaceKind                                       string `json:"surface_kind"`
	StatusConsumerKind                                string `json:"status_consumer_kind"`
	OpaqueReceiptID                                   string `json:"opaque_receipt_id"`
	EvidencePresent                                   bool   `json:"evidence_present"`
	CurrentMainlineConsumed                           bool   `json:"current_mainline_consumed"`
	RouteEnablementLookupRouteGrantConsumed           bool   `json:"route_enablement_lookup_route_grant_consumed"`
	RouteEnablementLookupRouteGrantReady              bool   `json:"route_enablement_lookup_route_grant_ready"`
	LookupRouteEnablementModeled                      bool   `json:"lookup_route_enablement_modeled"`
	RouteEnablementLookupRouteEnablementBoundaryReady bool   `json:"route_enablement_lookup_route_enablement_boundary_ready"`
	OpaqueRouteEnablementReceipt                      bool   `json:"opaque_route_enablement_receipt"`
	KDESafeRedactedStatusOnly                         bool   `json:"kde_safe_redacted_status_only"`
	ReceiptPresent                                    bool   `json:"receipt_present"`
	ReceiptAccepted                                   bool   `json:"receipt_accepted"`
	ReceiptConsumed                                   bool   `json:"receipt_consumed"`
	AcceptanceAuthorized                              bool   `json:"acceptance_authorized"`
	RouteEnablementAccepted                           bool   `json:"route_enablement_accepted"`
	LookupRouteEnablementGranted                      bool   `json:"lookup_route_enablement_granted"`
	LookupRouteGrantAuthorized                        bool   `json:"lookup_route_grant_authorized"`
	LookupRouteEnabled                                bool   `json:"lookup_route_enabled"`
	StorageWriteEnabled                               bool   `json:"storage_write_enabled"`
	RedactedSummaryPersisted                          bool   `json:"redacted_summary_persisted"`
	KDEStatusPersisted                                bool   `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted                       bool   `json:"runtime_diagnostics_persisted"`
	RawResultExposed                                  bool   `json:"raw_result_exposed"`
	UserVisible                                       bool   `json:"user_visible"`
	ReviewOnly                                        bool   `json:"review_only"`
	RuntimeOwned                                      bool   `json:"runtime_owned"`
	GoRuntimeBacked                                   bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                                    bool   `json:"kde_policy_owner"`
	SideEffectsDisabled                               bool   `json:"side_effects_disabled"`
	HostRootModified                                  bool   `json:"host_root_modified"`
	InternalDetailsExposed                            bool   `json:"internal_details_exposed"`
	LookupRouteEnablementStatus                       string `json:"lookup_route_enablement_status"`
	NextRequirement                                   string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditCheckCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

type routeEnablementLookupRouteEnablementAuditSourceSet struct {
	CurrentMainline                      string
	RouteEnablementLookupRouteGrantAudit string
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditPreview{}, err
	}
	sources := routeEnablementLookupRouteEnablementAuditSources(root)
	mainlineReady := routeEnablementLookupRouteEnablementAuditMainlineReady(sources.CurrentMainline)
	grantReady := routeEnablementLookupRouteEnablementAuditGrantReady(sources.RouteEnablementLookupRouteGrantAudit)
	items := routeEnablementLookupRouteEnablementAuditItems(mainlineReady, grantReady)
	readyItemCount := routeEnablementLookupRouteEnablementAuditReadyCount(items)
	ready := mainlineReady && grantReady && readyItemCount == len(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditPreview{
		Version:                                 version,
		SchemaVersion:                           "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_enablement_audit.v1",
		RequestType:                             "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-enablement-audit-preview",
		AuditType:                               "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-enablement-audit",
		Source:                                  "docs/xnix-current-mainline.md+route-enable-lookup-route-grant-audit",
		AuditDecision:                           "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-enablement-audit-blocked",
		CurrentMainlineConsumed:                 mainlineReady,
		RouteEnablementLookupRouteGrantConsumed: grantReady,
		RouteEnablementLookupRouteGrantReady:    grantReady,
		LookupRouteEnablementRequired:           true,
		LookupRouteEnablementModeled:            true,
		LookupRouteEnablementReady:              ready,
		RouteEnablementLookupRouteEnablementBoundaryReady: ready,
		OpaqueRouteEnablementReceipt:                      grantReady,
		KDESafeRedactedStatusOnly:                         ready,
		CompatibilityCenterLookupRouteEnablementModeled:   ready,
		RuntimeDiagnosticsLookupRouteEnablementModeled:    ready,
		EnablementItemCount:                               len(items),
		RequiredEnablementItemCount:                       len(items),
		ReadyEnablementItemCount:                          readyItemCount,
		MissingEnablementItemCount:                        len(items) - readyItemCount,
		CompatibilityCenterEnablementItemCount:            routeEnablementLookupRouteEnablementAuditSurfaceCount(items, "compatibility-center"),
		RuntimeDiagnosticsEnablementItemCount:             routeEnablementLookupRouteEnablementAuditSurfaceCount(items, "runtime-diagnostics"),
		EnablementItems:                                   items,
		EnablementItemIDs:                                 routeEnablementLookupRouteEnablementAuditItemIDs(items),
		RuntimeOwned:                                      true,
		GoRuntimeBacked:                                   true,
		KDEPolicyOwner:                                    false,
		BlockedActions: []string{
			"receipt-present",
			"receipt-accept",
			"receipt-consume",
			"lookup-route-enablement",
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
			"Add a route enablement lookup route dispatch gate audit before any lookup route can be called.",
			"Keep lookup route enablement modeled, redacted, and disabled until production ownership is explicitly authorized.",
		},
		DesktopSafeSummary: "KDE-safe redacted route enablement lookup route enablement is audited for Compatibility Center and Runtime diagnostics while receipts, route grants, lookup routes, consumers, dispatch, persistence, production, backend, and host side effects remain disabled.",
	}
	if ready {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-enablement-audit-ready-route-disabled-dispatch-disabled"
	}
	preview.Checks = routeEnablementLookupRouteEnablementAuditChecks(preview)
	preview.CheckIDs = routeEnablementLookupRouteEnablementAuditCheckIDs(preview.Checks)
	preview.Counts = routeEnablementLookupRouteEnablementAuditCountsFor(preview.Checks)
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route enablement audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditPreview{}, err
	}
	return preview, nil
}

func routeEnablementLookupRouteEnablementAuditSources(root string) routeEnablementLookupRouteEnablementAuditSourceSet {
	return routeEnablementLookupRouteEnablementAuditSourceSet{
		CurrentMainline: productionAuthorizationReadSources(root, []string{"docs/xnix-current-mainline.md"}),
		RouteEnablementLookupRouteGrantAudit: productionAuthorizationReadSources(root, []string{
			"internal/runtime/owner/route_enablement_lookup_route_grant_audit_preview.go",
			"internal/runtime/owner/route_enablement_lookup_route_grant_audit_preview_test.go",
		}),
	}
}

func routeEnablementLookupRouteEnablementAuditMainlineReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"route enablement lookup route enablement audit preview",
		"route enablement lookup route grant audit preview",
		"without accepting or consuming receipts",
	})
}

func routeEnablementLookupRouteEnablementAuditGrantReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_grant_audit.v1",
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-grant-audit-ready-grant-disabled-routes-disabled",
		"LookupRouteGrantReady",
		"RouteEnablementLookupRouteGrantBoundaryReady",
	})
}

func routeEnablementLookupRouteEnablementAuditItems(mainlineReady, grantReady bool) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditItem {
	actionKinds := []string{"review", "renew", "open-compatibility-center", "dismiss", "support-info"}
	items := make([]ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditItem, 0, len(actionKinds)*2)
	for _, actionKind := range actionKinds {
		for _, surfaceKind := range []string{"compatibility-center", "runtime-diagnostics"} {
			ready := mainlineReady && grantReady
			status := "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-enablement-evidence"
			if ready {
				status = "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-enablement-modeled-route-disabled-dispatch-disabled"
			}
			items = append(items, ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditItem{
				ID:                                      actionKind + "-route-enablement-lookup-route-enablement-" + surfaceKind,
				ActionKind:                              actionKind,
				SurfaceKind:                             surfaceKind,
				StatusConsumerKind:                      "kde-safe-redacted-route-enablement-lookup-route-enablement",
				OpaqueReceiptID:                         "opaque-route-enablement-lookup-route-enablement-" + surfaceKind + "-" + actionKind,
				EvidencePresent:                         ready,
				CurrentMainlineConsumed:                 mainlineReady,
				RouteEnablementLookupRouteGrantConsumed: grantReady,
				RouteEnablementLookupRouteGrantReady:    grantReady,
				LookupRouteEnablementModeled:            ready,
				RouteEnablementLookupRouteEnablementBoundaryReady: ready,
				OpaqueRouteEnablementReceipt:                      grantReady,
				KDESafeRedactedStatusOnly:                         ready,
				UserVisible:                                       ready,
				ReviewOnly:                                        ready,
				RuntimeOwned:                                      true,
				GoRuntimeBacked:                                   true,
				KDEPolicyOwner:                                    false,
				SideEffectsDisabled:                               true,
				LookupRouteEnablementStatus:                       status,
				NextRequirement:                                   "Require a separate lookup route dispatch gate audit before the modeled enablement can call any route.",
			})
		}
	}
	return items
}

func routeEnablementLookupRouteEnablementAuditReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent && item.LookupRouteEnablementModeled && item.RouteEnablementLookupRouteEnablementBoundaryReady && item.KDESafeRedactedStatusOnly && !item.ReceiptPresent && !item.ReceiptAccepted && !item.ReceiptConsumed && !item.AcceptanceAuthorized && !item.RouteEnablementAccepted && !item.LookupRouteEnablementGranted && !item.LookupRouteGrantAuthorized && !item.LookupRouteEnabled && !item.StorageWriteEnabled && !item.RawResultExposed && !item.HostRootModified && !item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func routeEnablementLookupRouteEnablementAuditSurfaceCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditItem, surfaceKind string) int {
	count := 0
	for _, item := range items {
		if item.SurfaceKind == surfaceKind {
			count++
		}
	}
	return count
}

func routeEnablementLookupRouteEnablementAuditItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func routeEnablementLookupRouteEnablementAuditChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditCheck{
		routeEnablementLookupRouteEnablementAuditCheck("current-mainline-consumed", preview.CurrentMainlineConsumed, "Xnix current mainline names the result consumer projection route enablement lookup route enablement audit preview."),
		routeEnablementLookupRouteEnablementAuditCheck("route-enablement-lookup-route-grant-consumed", preview.RouteEnablementLookupRouteGrantConsumed && preview.RouteEnablementLookupRouteGrantReady, "The lookup route enablement audit consumes the ready route enablement lookup route grant audit."),
		routeEnablementLookupRouteEnablementAuditCheck("lookup-route-enablement-modeled", preview.LookupRouteEnablementRequired && preview.LookupRouteEnablementModeled, "Lookup route enablement is modeled without enabling a route."),
		routeEnablementLookupRouteEnablementAuditCheck("compatibility-center-and-runtime-lookup-route-enablements-modeled", preview.CompatibilityCenterLookupRouteEnablementModeled && preview.RuntimeDiagnosticsLookupRouteEnablementModeled, "Compatibility Center and Runtime diagnostics lookup route enablement boundaries are modeled."),
		routeEnablementLookupRouteEnablementAuditCheck("ten-enablement-items-ready-route-disabled", preview.EnablementItemCount == 10 && preview.ReadyEnablementItemCount == 10 && preview.MissingEnablementItemCount == 0, "Ten KDE-safe lookup route enablement items are ready while route and dispatch gates remain disabled."),
		routeEnablementLookupRouteEnablementAuditCheck("lookup-route-enablement-and-persistence-disabled", !preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.ReceiptConsumed && !preview.AcceptanceAuthorized && !preview.RouteEnablementAccepted && !preview.LookupRouteEnablementGranted && !preview.LookupRouteGrantAuthorized && !preview.LookupRouteEnabled && !preview.StorageWriteEnabled && !preview.StatusPersistenceWriteEnabled && !preview.RawResultExposed, "Lookup route enablement, lookup grants, lookup routes, storage, persistence, and raw result exposure remain disabled."),
		routeEnablementLookupRouteEnablementAuditCheck("consumer-routes-dispatch-and-support-disabled", !preview.ConsumerEnablementAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled && !preview.LookupRouteEnabled && !preview.DispatchDryRunExecuted && !preview.SupportBundleExported && !preview.SupportCaseCreated, "Consumer enablement, lookup routes, dispatch, and support side effects remain disabled."),
		routeEnablementLookupRouteEnablementAuditCheck("production-and-host-boundary-closed", !preview.ProductionBusClaimed && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.BackendLaunchEnabled && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired, "Production, backend, network, privileged container, and host mutation gates remain closed."),
	}
}

func routeEnablementLookupRouteEnablementAuditCheck(id string, ok bool, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditCheck {
	status := "blocked"
	if ok {
		status = "pass"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditCheck{ID: id, Status: status, Summary: summary}
}

func routeEnablementLookupRouteEnablementAuditCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func routeEnablementLookupRouteEnablementAuditCountsFor(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditCheckCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditCheckCounts{Total: len(checks)}
	for _, check := range checks {
		if check.Status == "pass" {
			counts.Passed++
		} else {
			counts.Blocked++
		}
	}
	return counts
}
