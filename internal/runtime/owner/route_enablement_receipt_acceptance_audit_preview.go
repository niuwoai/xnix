package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditPreview struct {
	Version                                      string                                                                                                                                                                                                                            `json:"version"`
	SchemaVersion                                string                                                                                                                                                                                                                            `json:"schema_version"`
	RequestType                                  string                                                                                                                                                                                                                            `json:"request_type"`
	AuditType                                    string                                                                                                                                                                                                                            `json:"audit_type"`
	Source                                       string                                                                                                                                                                                                                            `json:"source"`
	AuditDecision                                string                                                                                                                                                                                                                            `json:"audit_decision"`
	CurrentMainlineConsumed                      bool                                                                                                                                                                                                                              `json:"current_mainline_consumed"`
	RouteEnablementReceiptGateConsumed           bool                                                                                                                                                                                                                              `json:"route_enablement_receipt_gate_consumed"`
	RouteEnablementReceiptGateReady              bool                                                                                                                                                                                                                              `json:"route_enablement_receipt_gate_ready"`
	AcceptanceAuthorizationRequired              bool                                                                                                                                                                                                                              `json:"acceptance_authorization_required"`
	AcceptanceAuthorizationModeled               bool                                                                                                                                                                                                                              `json:"acceptance_authorization_modeled"`
	AcceptanceAuthorizationReady                 bool                                                                                                                                                                                                                              `json:"acceptance_authorization_ready"`
	RouteEnablementAcceptanceBoundaryReady       bool                                                                                                                                                                                                                              `json:"route_enablement_acceptance_boundary_ready"`
	OpaqueRouteEnablementReceipt                 bool                                                                                                                                                                                                                              `json:"opaque_route_enablement_receipt"`
	KDESafeRedactedStatusOnly                    bool                                                                                                                                                                                                                              `json:"kde_safe_redacted_status_only"`
	CompatibilityCenterAcceptanceModeled         bool                                                                                                                                                                                                                              `json:"compatibility_center_acceptance_modeled"`
	RuntimeDiagnosticsAcceptanceModeled          bool                                                                                                                                                                                                                              `json:"runtime_diagnostics_acceptance_modeled"`
	ReceiptPresent                               bool                                                                                                                                                                                                                              `json:"receipt_present"`
	ReceiptAccepted                              bool                                                                                                                                                                                                                              `json:"receipt_accepted"`
	ReceiptConsumed                              bool                                                                                                                                                                                                                              `json:"receipt_consumed"`
	AcceptanceAuthorized                         bool                                                                                                                                                                                                                              `json:"acceptance_authorized"`
	RouteEnablementAccepted                      bool                                                                                                                                                                                                                              `json:"route_enablement_accepted"`
	LookupRouteEnablementGranted                 bool                                                                                                                                                                                                                              `json:"lookup_route_enablement_granted"`
	LookupRouteEnabled                           bool                                                                                                                                                                                                                              `json:"lookup_route_enabled"`
	StorageRootPolicyGrantAuthorized             bool                                                                                                                                                                                                                              `json:"storage_root_policy_grant_authorized"`
	RecordWriterCallAuthorized                   bool                                                                                                                                                                                                                              `json:"record_writer_call_authorized"`
	WriterCallable                               bool                                                                                                                                                                                                                              `json:"writer_callable"`
	StorageRootResolved                          bool                                                                                                                                                                                                                              `json:"storage_root_resolved"`
	StorageWriteEnabled                          bool                                                                                                                                                                                                                              `json:"storage_write_enabled"`
	StatusPersistenceWriteEnabled                bool                                                                                                                                                                                                                              `json:"status_persistence_write_enabled"`
	AcceptanceItemCount                          int                                                                                                                                                                                                                               `json:"acceptance_item_count"`
	RequiredAcceptanceItemCount                  int                                                                                                                                                                                                                               `json:"required_acceptance_item_count"`
	ReadyAcceptanceItemCount                     int                                                                                                                                                                                                                               `json:"ready_acceptance_item_count"`
	MissingAcceptanceItemCount                   int                                                                                                                                                                                                                               `json:"missing_acceptance_item_count"`
	AcceptedReceiptItemCount                     int                                                                                                                                                                                                                               `json:"accepted_receipt_item_count"`
	EnabledRouteItemCount                        int                                                                                                                                                                                                                               `json:"enabled_route_item_count"`
	PersistedAcceptanceItemCount                 int                                                                                                                                                                                                                               `json:"persisted_acceptance_item_count"`
	RawExposedAcceptanceItemCount                int                                                                                                                                                                                                                               `json:"raw_exposed_acceptance_item_count"`
	SideEffectAcceptanceItemCount                int                                                                                                                                                                                                                               `json:"side_effect_acceptance_item_count"`
	CompatibilityCenterAcceptanceItemCount       int                                                                                                                                                                                                                               `json:"compatibility_center_acceptance_item_count"`
	RuntimeDiagnosticsAcceptanceItemCount        int                                                                                                                                                                                                                               `json:"runtime_diagnostics_acceptance_item_count"`
	AcceptanceItems                              []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditItem      `json:"acceptance_items"`
	AcceptanceItemIDs                            []string                                                                                                                                                                                                                          `json:"acceptance_item_ids"`
	Checks                                       []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditCheck     `json:"checks"`
	CheckIDs                                     []string                                                                                                                                                                                                                          `json:"check_ids"`
	Counts                                       ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditCheckCounts `json:"counts"`
	ConsumerDryRunLookupResultBoundaryAuthorized bool                                                                                                                                                                                                                              `json:"consumer_dry_run_lookup_result_boundary_authorized"`
	ConsumerEnablementAuthorized                 bool                                                                                                                                                                                                                              `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                           bool                                                                                                                                                                                                                              `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                       bool                                                                                                                                                                                                                              `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                        bool                                                                                                                                                                                                                              `json:"lookup_route_authorized"`
	LookupRouteEnablementAuthorized              bool                                                                                                                                                                                                                              `json:"lookup_route_enablement_authorized"`
	OpaqueLookupEnabled                          bool                                                                                                                                                                                                                              `json:"opaque_lookup_enabled"`
	RedactedSummaryPersisted                     bool                                                                                                                                                                                                                              `json:"redacted_summary_persisted"`
	KDEStatusPersisted                           bool                                                                                                                                                                                                                              `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted                  bool                                                                                                                                                                                                                              `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                        bool                                                                                                                                                                                                                              `json:"dry_run_result_persisted"`
	RawResultExposed                             bool                                                                                                                                                                                                                              `json:"raw_result_exposed"`
	DispatchDryRunExecuted                       bool                                                                                                                                                                                                                              `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled                 bool                                                                                                                                                                                                                              `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled                 bool                                                                                                                                                                                                                              `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                         bool                                                                                                                                                                                                                              `json:"portal_request_created"`
	NotificationActionEnabled                    bool                                                                                                                                                                                                                              `json:"notification_action_enabled"`
	CompatibilityCenterOpened                    bool                                                                                                                                                                                                                              `json:"compatibility_center_opened"`
	SupportBundleExported                        bool                                                                                                                                                                                                                              `json:"support_bundle_exported"`
	SupportCaseCreated                           bool                                                                                                                                                                                                                              `json:"support_case_created"`
	RuntimeOwned                                 bool                                                                                                                                                                                                                              `json:"runtime_owned"`
	GoRuntimeBacked                              bool                                                                                                                                                                                                                              `json:"go_runtime_backed"`
	KDEPolicyOwner                               bool                                                                                                                                                                                                                              `json:"kde_policy_owner"`
	ProductionReadiness                          bool                                                                                                                                                                                                                              `json:"production_readiness"`
	ProductionOwnershipReady                     bool                                                                                                                                                                                                                              `json:"production_ownership_ready"`
	SystemServiceStarted                         bool                                                                                                                                                                                                                              `json:"system_service_started"`
	SessionBusClaimed                            bool                                                                                                                                                                                                                              `json:"session_bus_claimed"`
	ProductionBusClaimed                         bool                                                                                                                                                                                                                              `json:"production_bus_claimed"`
	WriteMethodsEnabled                          bool                                                                                                                                                                                                                              `json:"write_methods_enabled"`
	RuntimeWritesEnabled                         bool                                                                                                                                                                                                                              `json:"runtime_writes_enabled"`
	DesktopFilesWritten                          bool                                                                                                                                                                                                                              `json:"desktop_files_written"`
	KDEConfigurationWritten                      bool                                                                                                                                                                                                                              `json:"kde_configuration_written"`
	PortalCallExecuted                           bool                                                                                                                                                                                                                              `json:"portal_call_executed"`
	AdapterInvocationEnabled                     bool                                                                                                                                                                                                                              `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                         bool                                                                                                                                                                                                                              `json:"backend_launch_enabled"`
	BackendProcessStarted                        bool                                                                                                                                                                                                                              `json:"backend_process_started"`
	NetworkRequired                              bool                                                                                                                                                                                                                              `json:"network_required"`
	HostRootModified                             bool                                                                                                                                                                                                                              `json:"host_root_modified"`
	PrivilegedContainerRequired                  bool                                                                                                                                                                                                                              `json:"privileged_container_required"`
	StateRootPathExposed                         bool                                                                                                                                                                                                                              `json:"state_root_path_exposed"`
	FilePathsExposed                             bool                                                                                                                                                                                                                              `json:"file_paths_exposed"`
	FileContentRead                              bool                                                                                                                                                                                                                              `json:"file_content_read"`
	RawCommandExposed                            bool                                                                                                                                                                                                                              `json:"raw_command_exposed"`
	RawExecutableExposed                         bool                                                                                                                                                                                                                              `json:"raw_executable_exposed"`
	BackendDetailsExposed                        bool                                                                                                                                                                                                                              `json:"backend_details_exposed"`
	BlockedActions                               []string                                                                                                                                                                                                                          `json:"blocked_actions"`
	NextRequirements                             []string                                                                                                                                                                                                                          `json:"next_requirements"`
	DesktopSafeSummary                           string                                                                                                                                                                                                                            `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditItem struct {
	ID                                     string `json:"id"`
	ActionKind                             string `json:"action_kind"`
	SurfaceKind                            string `json:"surface_kind"`
	StatusConsumerKind                     string `json:"status_consumer_kind"`
	OpaqueReceiptID                        string `json:"opaque_receipt_id"`
	EvidencePresent                        bool   `json:"evidence_present"`
	CurrentMainlineConsumed                bool   `json:"current_mainline_consumed"`
	RouteEnablementReceiptGateConsumed     bool   `json:"route_enablement_receipt_gate_consumed"`
	RouteEnablementReceiptGateReady        bool   `json:"route_enablement_receipt_gate_ready"`
	AcceptanceAuthorizationModeled         bool   `json:"acceptance_authorization_modeled"`
	RouteEnablementAcceptanceBoundaryReady bool   `json:"route_enablement_acceptance_boundary_ready"`
	OpaqueRouteEnablementReceipt           bool   `json:"opaque_route_enablement_receipt"`
	KDESafeRedactedStatusOnly              bool   `json:"kde_safe_redacted_status_only"`
	ReceiptPresent                         bool   `json:"receipt_present"`
	ReceiptAccepted                        bool   `json:"receipt_accepted"`
	ReceiptConsumed                        bool   `json:"receipt_consumed"`
	AcceptanceAuthorized                   bool   `json:"acceptance_authorized"`
	RouteEnablementAccepted                bool   `json:"route_enablement_accepted"`
	LookupRouteEnablementGranted           bool   `json:"lookup_route_enablement_granted"`
	LookupRouteEnabled                     bool   `json:"lookup_route_enabled"`
	StorageWriteEnabled                    bool   `json:"storage_write_enabled"`
	RedactedSummaryPersisted               bool   `json:"redacted_summary_persisted"`
	KDEStatusPersisted                     bool   `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted            bool   `json:"runtime_diagnostics_persisted"`
	RawResultExposed                       bool   `json:"raw_result_exposed"`
	UserVisible                            bool   `json:"user_visible"`
	ReviewOnly                             bool   `json:"review_only"`
	RuntimeOwned                           bool   `json:"runtime_owned"`
	GoRuntimeBacked                        bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                         bool   `json:"kde_policy_owner"`
	SideEffectsDisabled                    bool   `json:"side_effects_disabled"`
	HostRootModified                       bool   `json:"host_root_modified"`
	InternalDetailsExposed                 bool   `json:"internal_details_exposed"`
	AcceptanceAuthorizationStatus          string `json:"acceptance_authorization_status"`
	NextRequirement                        string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditCheckCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

type productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditSourceSet struct {
	CurrentMainline            string
	RouteEnablementReceiptGate string
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditSources(root)
	mainlineReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditMainlineReady(sources.CurrentMainline)
	receiptGateReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditGateReady(sources.RouteEnablementReceiptGate)
	items := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditItems(mainlineReady, receiptGateReady)
	readyItemCount := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditReadyCount(items)
	ready := mainlineReady && receiptGateReady && readyItemCount == len(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditPreview{
		Version:                                version,
		SchemaVersion:                          "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_receipt_acceptance_audit.v1",
		RequestType:                            "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-acceptance-audit-preview",
		AuditType:                              "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-acceptance-audit",
		Source:                                 "xnix-current-mainline+redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-gate",
		AuditDecision:                          "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-acceptance-audit-blocked",
		CurrentMainlineConsumed:                mainlineReady,
		RouteEnablementReceiptGateConsumed:     receiptGateReady,
		RouteEnablementReceiptGateReady:        receiptGateReady,
		AcceptanceAuthorizationRequired:        true,
		AcceptanceAuthorizationModeled:         true,
		AcceptanceAuthorizationReady:           ready,
		RouteEnablementAcceptanceBoundaryReady: ready,
		OpaqueRouteEnablementReceipt:           receiptGateReady,
		KDESafeRedactedStatusOnly:              ready,
		CompatibilityCenterAcceptanceModeled:   ready,
		RuntimeDiagnosticsAcceptanceModeled:    ready,
		AcceptanceItemCount:                    len(items),
		RequiredAcceptanceItemCount:            len(items),
		ReadyAcceptanceItemCount:               readyItemCount,
		MissingAcceptanceItemCount:             len(items) - readyItemCount,
		CompatibilityCenterAcceptanceItemCount: productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditSurfaceCount(items, "compatibility-center"),
		RuntimeDiagnosticsAcceptanceItemCount:  productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditSurfaceCount(items, "runtime-diagnostics"),
		AcceptanceItems:                        items,
		AcceptanceItemIDs:                      productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditItemIDs(items),
		RuntimeOwned:                           true,
		GoRuntimeBacked:                        true,
		KDEPolicyOwner:                         false,
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
			"Add a route enablement accepted receipt gate audit before any lookup route can be enabled.",
			"Keep receipt acceptance modeled, redacted, and disabled until production ownership is explicitly authorized.",
		},
		DesktopSafeSummary: "KDE-safe redacted route enablement receipt acceptance is audited for Compatibility Center and Runtime diagnostics while receipt acceptance, lookup routes, consumers, dispatch, persistence, production, backend, and host side effects remain disabled.",
	}
	if ready {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-acceptance-audit-ready-acceptance-disabled-routes-disabled"
	}
	preview.Checks = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditChecks(preview)
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditCheckIDs(preview.Checks)
	preview.Counts = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditCountsFor(preview.Checks)
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status dry-run lookup result consumer projection route enablement receipt acceptance audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditPreview{}, err
	}
	return preview, nil
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditSources(root string) productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditSourceSet {
	return productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditSourceSet{
		CurrentMainline: productionAuthorizationReadSources(root, []string{"docs/xnix-current-mainline.md"}),
		RouteEnablementReceiptGate: productionAuthorizationReadSources(root, []string{
			"internal/runtime/owner/route_enablement_receipt_gate_preview.go",
			"internal/runtime/owner/route_enablement_receipt_gate_preview_test.go",
		}),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditMainlineReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"route enablement receipt acceptance audit preview",
		"route enablement receipt gate preview",
		"without accepting or consuming receipts",
	})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditGateReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_receipt_gate.v1",
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-gate-ready-routes-disabled",
		"RouteEnablementReceiptGateReady",
		"OperatorReviewedReceiptModeled",
	})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditItems(mainlineReady, receiptGateReady bool) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditItem {
	actionKinds := []string{"review", "renew", "open-compatibility-center", "dismiss", "support-info"}
	items := make([]ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditItem, 0, len(actionKinds)*2)
	for _, actionKind := range actionKinds {
		for _, surfaceKind := range []string{"compatibility-center", "runtime-diagnostics"} {
			ready := mainlineReady && receiptGateReady
			status := "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-acceptance-evidence"
			if ready {
				status = "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-acceptance-modeled-acceptance-disabled-routes-disabled"
			}
			items = append(items, ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditItem{
				ID:                                     actionKind + "-route-enablement-receipt-acceptance-" + surfaceKind,
				ActionKind:                             actionKind,
				SurfaceKind:                            surfaceKind,
				StatusConsumerKind:                     "kde-safe-redacted-route-enablement-receipt-acceptance",
				OpaqueReceiptID:                        "opaque-route-enablement-receipt-" + surfaceKind + "-" + actionKind,
				EvidencePresent:                        ready,
				CurrentMainlineConsumed:                mainlineReady,
				RouteEnablementReceiptGateConsumed:     receiptGateReady,
				RouteEnablementReceiptGateReady:        receiptGateReady,
				AcceptanceAuthorizationModeled:         ready,
				RouteEnablementAcceptanceBoundaryReady: ready,
				OpaqueRouteEnablementReceipt:           receiptGateReady,
				KDESafeRedactedStatusOnly:              ready,
				UserVisible:                            ready,
				ReviewOnly:                             ready,
				RuntimeOwned:                           true,
				GoRuntimeBacked:                        true,
				KDEPolicyOwner:                         false,
				SideEffectsDisabled:                    true,
				AcceptanceAuthorizationStatus:          status,
				NextRequirement:                        "Require a separate accepted receipt gate audit before lookup route enablement can be granted.",
			})
		}
	}
	return items
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent && item.AcceptanceAuthorizationModeled && item.RouteEnablementAcceptanceBoundaryReady && item.KDESafeRedactedStatusOnly && !item.ReceiptPresent && !item.ReceiptAccepted && !item.ReceiptConsumed && !item.AcceptanceAuthorized && !item.RouteEnablementAccepted && !item.LookupRouteEnablementGranted && !item.LookupRouteEnabled && !item.StorageWriteEnabled && !item.RawResultExposed && !item.HostRootModified && !item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditSurfaceCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditItem, surfaceKind string) int {
	count := 0
	for _, item := range items {
		if item.SurfaceKind == surfaceKind {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditCheck{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditCheck("current-mainline-consumed", preview.CurrentMainlineConsumed, "Xnix current mainline names the result consumer projection route enablement receipt acceptance audit preview."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditCheck("route-enablement-receipt-gate-consumed", preview.RouteEnablementReceiptGateConsumed && preview.RouteEnablementReceiptGateReady, "The audit consumes the ready route enablement receipt gate."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditCheck("acceptance-authorization-boundary-modeled", preview.AcceptanceAuthorizationRequired && preview.AcceptanceAuthorizationModeled, "Route enablement receipt acceptance authorization is modeled without accepting a receipt."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditCheck("compatibility-center-and-runtime-acceptance-modeled", preview.CompatibilityCenterAcceptanceModeled && preview.RuntimeDiagnosticsAcceptanceModeled, "Compatibility Center and Runtime diagnostics acceptance boundaries are modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditCheck("ten-acceptance-items-ready-acceptance-disabled", preview.AcceptanceItemCount == 10 && preview.ReadyAcceptanceItemCount == 10 && preview.MissingAcceptanceItemCount == 0, "Ten KDE-safe acceptance items are ready while receipt acceptance remains disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditCheck("receipt-acceptance-routes-and-persistence-disabled", !preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.ReceiptConsumed && !preview.AcceptanceAuthorized && !preview.RouteEnablementAccepted && !preview.LookupRouteEnablementGranted && !preview.LookupRouteEnabled && !preview.StorageWriteEnabled && !preview.StatusPersistenceWriteEnabled && !preview.RawResultExposed, "Receipt acceptance, route enablement grants, lookup routes, storage, persistence, and raw result exposure remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditCheck("consumer-routes-dispatch-and-support-disabled", !preview.ConsumerEnablementAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled && !preview.LookupRouteEnabled && !preview.DispatchDryRunExecuted && !preview.SupportBundleExported && !preview.SupportCaseCreated, "Consumer enablement, lookup routes, dispatch, and support side effects remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditCheck("production-and-host-boundary-closed", !preview.ProductionBusClaimed && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.BackendLaunchEnabled && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired, "Production, backend, network, privileged container, and host mutation gates remain closed."),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditCheck(id string, ok bool, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditCheck {
	status := "blocked"
	if ok {
		status = "pass"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditCountsFor(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditCheckCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditCheckCounts{Total: len(checks)}
	for _, check := range checks {
		if check.Status == "pass" {
			counts.Passed++
		} else {
			counts.Blocked++
		}
	}
	return counts
}
