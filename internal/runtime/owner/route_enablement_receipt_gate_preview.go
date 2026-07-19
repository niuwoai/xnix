package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGatePreview struct {
	Version                                      string                                                                                                                                                                                                             `json:"version"`
	SchemaVersion                                string                                                                                                                                                                                                             `json:"schema_version"`
	RequestType                                  string                                                                                                                                                                                                             `json:"request_type"`
	GateType                                     string                                                                                                                                                                                                             `json:"gate_type"`
	Source                                       string                                                                                                                                                                                                             `json:"source"`
	AuditDecision                                string                                                                                                                                                                                                             `json:"audit_decision"`
	CurrentMainlineConsumed                      bool                                                                                                                                                                                                               `json:"current_mainline_consumed"`
	RouteEnablementEvidenceGateConsumed          bool                                                                                                                                                                                                               `json:"route_enablement_evidence_gate_consumed"`
	RouteEnablementEvidenceGateReady             bool                                                                                                                                                                                                               `json:"route_enablement_evidence_gate_ready"`
	RouteEnablementReceiptGateRequired           bool                                                                                                                                                                                                               `json:"route_enablement_receipt_gate_required"`
	RouteEnablementReceiptGateModeled            bool                                                                                                                                                                                                               `json:"route_enablement_receipt_gate_modeled"`
	RouteEnablementReceiptGateReady              bool                                                                                                                                                                                                               `json:"route_enablement_receipt_gate_ready"`
	OperatorReviewedReceiptGateRequired          bool                                                                                                                                                                                                               `json:"operator_reviewed_receipt_gate_required"`
	OperatorReviewedReceiptModeled               bool                                                                                                                                                                                                               `json:"operator_reviewed_receipt_modeled"`
	CompatibilityCenterReceiptGateModeled        bool                                                                                                                                                                                                               `json:"compatibility_center_receipt_gate_modeled"`
	RuntimeDiagnosticsReceiptGateModeled         bool                                                                                                                                                                                                               `json:"runtime_diagnostics_receipt_gate_modeled"`
	RedactedProjectionReceiptGateModeled         bool                                                                                                                                                                                                               `json:"redacted_projection_receipt_gate_modeled"`
	UserVisibleReceiptGateModeled                bool                                                                                                                                                                                                               `json:"user_visible_receipt_gate_modeled"`
	FailureProjectionReceiptGateModeled          bool                                                                                                                                                                                                               `json:"failure_projection_receipt_gate_modeled"`
	KDESafeRedactedStatusOnly                    bool                                                                                                                                                                                                               `json:"kde_safe_redacted_status_only"`
	ReceiptPresent                               bool                                                                                                                                                                                                               `json:"receipt_present"`
	ReceiptAccepted                              bool                                                                                                                                                                                                               `json:"receipt_accepted"`
	ReceiptConsumed                              bool                                                                                                                                                                                                               `json:"receipt_consumed"`
	ReceiptCreated                               bool                                                                                                                                                                                                               `json:"receipt_created"`
	ReceiptPersisted                             bool                                                                                                                                                                                                               `json:"receipt_persisted"`
	ReceiptAcceptanceAuthorized                  bool                                                                                                                                                                                                               `json:"receipt_acceptance_authorized"`
	StorageRootPolicyGrantAuthorized             bool                                                                                                                                                                                                               `json:"storage_root_policy_grant_authorized"`
	RecordWriterCallAuthorized                   bool                                                                                                                                                                                                               `json:"record_writer_call_authorized"`
	WriterCallable                               bool                                                                                                                                                                                                               `json:"writer_callable"`
	StorageRootResolved                          bool                                                                                                                                                                                                               `json:"storage_root_resolved"`
	StorageWriteEnabled                          bool                                                                                                                                                                                                               `json:"storage_write_enabled"`
	StatusPersistenceWriteEnabled                bool                                                                                                                                                                                                               `json:"status_persistence_write_enabled"`
	EvidenceItemCount                            int                                                                                                                                                                                                                `json:"evidence_item_count"`
	ReadyEvidenceItemCount                       int                                                                                                                                                                                                                `json:"ready_evidence_item_count"`
	MissingEvidenceItemCount                     int                                                                                                                                                                                                                `json:"missing_evidence_item_count"`
	CompatibilityCenterEvidenceItemCount         int                                                                                                                                                                                                                `json:"compatibility_center_evidence_item_count"`
	RuntimeDiagnosticsEvidenceItemCount          int                                                                                                                                                                                                                `json:"runtime_diagnostics_evidence_item_count"`
	RawExposedEvidenceItemCount                  int                                                                                                                                                                                                                `json:"raw_exposed_evidence_item_count"`
	SideEffectEvidenceItemCount                  int                                                                                                                                                                                                                `json:"side_effect_evidence_item_count"`
	EvidenceItems                                []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateItem  `json:"evidence_items"`
	EvidenceItemIDs                              []string                                                                                                                                                                                                           `json:"evidence_item_ids"`
	Checks                                       []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateCheck `json:"checks"`
	CheckIDs                                     []string                                                                                                                                                                                                           `json:"check_ids"`
	Counts                                       ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateCounts  `json:"counts"`
	ConsumerDryRunLookupResultBoundaryAuthorized bool                                                                                                                                                                                                               `json:"consumer_dry_run_lookup_result_boundary_authorized"`
	ConsumerEnablementAuthorized                 bool                                                                                                                                                                                                               `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                           bool                                                                                                                                                                                                               `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                       bool                                                                                                                                                                                                               `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                        bool                                                                                                                                                                                                               `json:"lookup_route_authorized"`
	LookupRouteEnablementAuthorized              bool                                                                                                                                                                                                               `json:"lookup_route_enablement_authorized"`
	LookupRouteEnabled                           bool                                                                                                                                                                                                               `json:"lookup_route_enabled"`
	OpaqueLookupEnabled                          bool                                                                                                                                                                                                               `json:"opaque_lookup_enabled"`
	RedactedSummaryPersisted                     bool                                                                                                                                                                                                               `json:"redacted_summary_persisted"`
	KDEStatusPersisted                           bool                                                                                                                                                                                                               `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted                  bool                                                                                                                                                                                                               `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                        bool                                                                                                                                                                                                               `json:"dry_run_result_persisted"`
	RawResultExposed                             bool                                                                                                                                                                                                               `json:"raw_result_exposed"`
	DispatchDryRunExecuted                       bool                                                                                                                                                                                                               `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled                 bool                                                                                                                                                                                                               `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled                 bool                                                                                                                                                                                                               `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                         bool                                                                                                                                                                                                               `json:"portal_request_created"`
	NotificationActionEnabled                    bool                                                                                                                                                                                                               `json:"notification_action_enabled"`
	CompatibilityCenterOpened                    bool                                                                                                                                                                                                               `json:"compatibility_center_opened"`
	SupportBundleExported                        bool                                                                                                                                                                                                               `json:"support_bundle_exported"`
	SupportCaseCreated                           bool                                                                                                                                                                                                               `json:"support_case_created"`
	RuntimeOwned                                 bool                                                                                                                                                                                                               `json:"runtime_owned"`
	GoRuntimeBacked                              bool                                                                                                                                                                                                               `json:"go_runtime_backed"`
	KDEPolicyOwner                               bool                                                                                                                                                                                                               `json:"kde_policy_owner"`
	ProductionReadiness                          bool                                                                                                                                                                                                               `json:"production_readiness"`
	ProductionOwnershipReady                     bool                                                                                                                                                                                                               `json:"production_ownership_ready"`
	SystemServiceStarted                         bool                                                                                                                                                                                                               `json:"system_service_started"`
	SessionBusClaimed                            bool                                                                                                                                                                                                               `json:"session_bus_claimed"`
	ProductionBusClaimed                         bool                                                                                                                                                                                                               `json:"production_bus_claimed"`
	WriteMethodsEnabled                          bool                                                                                                                                                                                                               `json:"write_methods_enabled"`
	RuntimeWritesEnabled                         bool                                                                                                                                                                                                               `json:"runtime_writes_enabled"`
	DesktopFilesWritten                          bool                                                                                                                                                                                                               `json:"desktop_files_written"`
	KDEConfigurationWritten                      bool                                                                                                                                                                                                               `json:"kde_configuration_written"`
	PortalCallExecuted                           bool                                                                                                                                                                                                               `json:"portal_call_executed"`
	AdapterInvocationEnabled                     bool                                                                                                                                                                                                               `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                         bool                                                                                                                                                                                                               `json:"backend_launch_enabled"`
	BackendProcessStarted                        bool                                                                                                                                                                                                               `json:"backend_process_started"`
	NetworkRequired                              bool                                                                                                                                                                                                               `json:"network_required"`
	HostRootModified                             bool                                                                                                                                                                                                               `json:"host_root_modified"`
	PrivilegedContainerRequired                  bool                                                                                                                                                                                                               `json:"privileged_container_required"`
	StateRootPathExposed                         bool                                                                                                                                                                                                               `json:"state_root_path_exposed"`
	FilePathsExposed                             bool                                                                                                                                                                                                               `json:"file_paths_exposed"`
	FileContentRead                              bool                                                                                                                                                                                                               `json:"file_content_read"`
	RawCommandExposed                            bool                                                                                                                                                                                                               `json:"raw_command_exposed"`
	RawExecutableExposed                         bool                                                                                                                                                                                                               `json:"raw_executable_exposed"`
	BackendDetailsExposed                        bool                                                                                                                                                                                                               `json:"backend_details_exposed"`
	BlockedActions                               []string                                                                                                                                                                                                           `json:"blocked_actions"`
	NextRequirements                             []string                                                                                                                                                                                                           `json:"next_requirements"`
	DesktopSafeSummary                           string                                                                                                                                                                                                             `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateItem struct {
	ID                                   string `json:"id"`
	ActionKind                           string `json:"action_kind"`
	SurfaceKind                          string `json:"surface_kind"`
	EvidenceKind                         string `json:"evidence_kind"`
	EvidencePresent                      bool   `json:"evidence_present"`
	CurrentMainlineConsumed              bool   `json:"current_mainline_consumed"`
	RouteEnablementEvidenceGateConsumed  bool   `json:"route_enablement_evidence_gate_consumed"`
	RouteEnablementEvidenceGateReady     bool   `json:"route_enablement_evidence_gate_ready"`
	RouteEnablementReceiptGateModeled    bool   `json:"route_enablement_receipt_gate_modeled"`
	OperatorReviewedReceiptModeled       bool   `json:"operator_reviewed_receipt_modeled"`
	RedactedProjectionReceiptGateModeled bool   `json:"redacted_projection_receipt_gate_modeled"`
	UserVisibleReceiptGateModeled        bool   `json:"user_visible_receipt_gate_modeled"`
	FailureProjectionReceiptGateModeled  bool   `json:"failure_projection_receipt_gate_modeled"`
	KDESafeRedactedStatusOnly            bool   `json:"kde_safe_redacted_status_only"`
	ReceiptPresent                       bool   `json:"receipt_present"`
	ReceiptAccepted                      bool   `json:"receipt_accepted"`
	ReceiptConsumed                      bool   `json:"receipt_consumed"`
	ReceiptCreated                       bool   `json:"receipt_created"`
	ReceiptPersisted                     bool   `json:"receipt_persisted"`
	ReceiptAcceptanceAuthorized          bool   `json:"receipt_acceptance_authorized"`
	LookupRouteEnablementAuthorized      bool   `json:"lookup_route_enablement_authorized"`
	LookupRouteEnabled                   bool   `json:"lookup_route_enabled"`
	ConsumerEnablementAuthorized         bool   `json:"consumer_enablement_authorized"`
	StorageWriteEnabled                  bool   `json:"storage_write_enabled"`
	RawResultExposed                     bool   `json:"raw_result_exposed"`
	UserVisible                          bool   `json:"user_visible"`
	ReviewOnly                           bool   `json:"review_only"`
	RuntimeOwned                         bool   `json:"runtime_owned"`
	GoRuntimeBacked                      bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                       bool   `json:"kde_policy_owner"`
	SideEffectsDisabled                  bool   `json:"side_effects_disabled"`
	HostRootModified                     bool   `json:"host_root_modified"`
	InternalDetailsExposed               bool   `json:"internal_details_exposed"`
	RouteEnablementReceiptGateStatus     string `json:"route_enablement_receipt_gate_status"`
	NextRequirement                      string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

type productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateSourceSet struct {
	CurrentMainline             string
	RouteEnablementEvidenceGate string
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGatePreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGatePreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGatePreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateSources(root)
	mainlineReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateMainlineReady(sources.CurrentMainline)
	evidenceGateReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateEvidenceReady(sources.RouteEnablementEvidenceGate)
	ready := mainlineReady && evidenceGateReady
	items := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateItems(mainlineReady, evidenceGateReady)
	readyCount := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateReadyCount(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGatePreview{
		Version:                               version,
		SchemaVersion:                         "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_receipt_gate.v1",
		RequestType:                           "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-gate-preview",
		GateType:                              "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-gate",
		Source:                                "xnix-current-mainline+redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-evidence-gate",
		AuditDecision:                         "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-gate-blocked",
		CurrentMainlineConsumed:               mainlineReady,
		RouteEnablementEvidenceGateConsumed:   evidenceGateReady,
		RouteEnablementEvidenceGateReady:      evidenceGateReady,
		RouteEnablementReceiptGateRequired:    true,
		RouteEnablementReceiptGateModeled:     ready,
		RouteEnablementReceiptGateReady:       ready,
		OperatorReviewedReceiptGateRequired:   true,
		OperatorReviewedReceiptModeled:        ready,
		CompatibilityCenterReceiptGateModeled: ready,
		RuntimeDiagnosticsReceiptGateModeled:  ready,
		RedactedProjectionReceiptGateModeled:  ready,
		UserVisibleReceiptGateModeled:         ready,
		FailureProjectionReceiptGateModeled:   ready,
		KDESafeRedactedStatusOnly:             ready,
		EvidenceItemCount:                     len(items),
		ReadyEvidenceItemCount:                readyCount,
		MissingEvidenceItemCount:              len(items) - readyCount,
		CompatibilityCenterEvidenceItemCount:  productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateSurfaceCount(items, "compatibility-center"),
		RuntimeDiagnosticsEvidenceItemCount:   productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateSurfaceCount(items, "runtime-diagnostics"),
		EvidenceItems:                         items,
		EvidenceItemIDs:                       productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateItemIDs(items),
		RuntimeOwned:                          true,
		GoRuntimeBacked:                       true,
		KDEPolicyOwner:                        false,
		BlockedActions: []string{
			"receipt-create",
			"receipt-accept",
			"receipt-consume",
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
			"Add a route enablement receipt acceptance audit before any operator-reviewed receipt can be accepted.",
			"Keep route enablement receipts modeled, redacted, and unpersisted until production ownership is explicitly authorized.",
		},
		DesktopSafeSummary: "KDE-safe redacted route enablement receipt gates are modeled for Compatibility Center and Runtime diagnostics while receipts, lookup routes, consumers, dispatch, persistence, production, backend, and host side effects remain disabled.",
	}
	if ready {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-gate-ready-routes-disabled"
	}
	preview.Checks = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateChecks(preview)
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateCheckIDs(preview.Checks)
	preview.Counts = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateCountsFor(preview.Checks)
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status dry-run lookup result consumer projection route enablement receipt gate preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGatePreview{}, err
	}
	return preview, nil
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateSources(root string) productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateSourceSet {
	return productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateSourceSet{
		CurrentMainline: productionAuthorizationReadSources(root, []string{"docs/xnix-current-mainline.md"}),
		RouteEnablementEvidenceGate: productionAuthorizationReadSources(root, []string{
			"internal/runtime/owner/route_enablement_evidence_gate_preview.go",
			"internal/runtime/owner/route_enablement_evidence_gate_preview_test.go",
		}),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateMainlineReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"route enablement receipt gate preview",
		"route enablement evidence gate preview",
		"without accepting or consuming receipts",
	})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateEvidenceReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_evidence_gate.v1",
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-evidence-gate-ready-routes-disabled",
		"RouteEnablementEvidenceGateReady",
		"ResultConsumerProjectionRouteEnablementGated",
	})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateItems(mainlineReady, evidenceGateReady bool) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateItem {
	actionKinds := []string{"review", "renew", "open-compatibility-center", "dismiss", "support-info"}
	items := make([]ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateItem, 0, len(actionKinds)*2)
	for _, actionKind := range actionKinds {
		for _, surfaceKind := range []string{"compatibility-center", "runtime-diagnostics"} {
			ready := mainlineReady && evidenceGateReady
			status := "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-gate"
			if ready {
				status = "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-gated-routes-disabled"
			}
			items = append(items, ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateItem{
				ID:                                   "route-enablement-receipt-gate-" + surfaceKind + "-" + actionKind,
				ActionKind:                           actionKind,
				SurfaceKind:                          surfaceKind,
				EvidenceKind:                         "kde-safe-redacted-lookup-result-projection-route-enablement-receipt-gate",
				EvidencePresent:                      ready,
				CurrentMainlineConsumed:              mainlineReady,
				RouteEnablementEvidenceGateConsumed:  evidenceGateReady,
				RouteEnablementEvidenceGateReady:     evidenceGateReady,
				RouteEnablementReceiptGateModeled:    ready,
				OperatorReviewedReceiptModeled:       ready,
				RedactedProjectionReceiptGateModeled: ready,
				UserVisibleReceiptGateModeled:        ready,
				FailureProjectionReceiptGateModeled:  ready,
				KDESafeRedactedStatusOnly:            ready,
				UserVisible:                          ready,
				ReviewOnly:                           ready,
				RuntimeOwned:                         true,
				GoRuntimeBacked:                      true,
				KDEPolicyOwner:                       false,
				SideEffectsDisabled:                  true,
				RouteEnablementReceiptGateStatus:     status,
				NextRequirement:                      "Require a separate receipt acceptance audit before creating, accepting, consuming, or persisting route enablement receipts.",
			})
		}
	}
	return items
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent && item.RouteEnablementReceiptGateModeled && item.OperatorReviewedReceiptModeled && item.RedactedProjectionReceiptGateModeled && item.UserVisibleReceiptGateModeled && item.FailureProjectionReceiptGateModeled && item.KDESafeRedactedStatusOnly && !item.ReceiptPresent && !item.ReceiptAccepted && !item.ReceiptConsumed && !item.LookupRouteEnabled && !item.StorageWriteEnabled && !item.RawResultExposed && !item.HostRootModified && !item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateSurfaceCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateItem, surfaceKind string) int {
	count := 0
	for _, item := range items {
		if item.SurfaceKind == surfaceKind {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGatePreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateCheck{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateCheck("current-mainline-consumed", preview.CurrentMainlineConsumed, "Xnix current mainline names the result consumer projection route enablement receipt gate preview."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateCheck("route-enablement-evidence-gate-consumed", preview.RouteEnablementEvidenceGateConsumed && preview.RouteEnablementEvidenceGateReady, "The gate consumes the ready route enablement evidence gate."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateCheck("operator-reviewed-receipt-gate-modeled", preview.OperatorReviewedReceiptGateRequired && preview.OperatorReviewedReceiptModeled, "Operator-reviewed route enablement receipt gate is modeled without creating a receipt."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateCheck("compatibility-center-and-runtime-receipt-gates-modeled", preview.CompatibilityCenterReceiptGateModeled && preview.RuntimeDiagnosticsReceiptGateModeled, "Compatibility Center and Runtime diagnostics route enablement receipt gate boundaries are modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateCheck("ten-receipt-gate-items-ready-routes-disabled", preview.EvidenceItemCount == 10 && preview.ReadyEvidenceItemCount == 10 && preview.MissingEvidenceItemCount == 0, "Ten KDE-safe receipt gate items are ready while lookup routes remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateCheck("receipt-writes-and-persistence-disabled", !preview.ReceiptPresent && !preview.ReceiptCreated && !preview.ReceiptPersisted && !preview.ReceiptAccepted && !preview.ReceiptConsumed && !preview.StorageWriteEnabled && !preview.StatusPersistenceWriteEnabled && !preview.RawResultExposed, "Receipt creation, acceptance, consumption, storage, persistence, and raw result exposure remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateCheck("consumer-routes-dispatch-and-support-disabled", !preview.ConsumerEnablementAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled && !preview.LookupRouteEnabled && !preview.DispatchDryRunExecuted && !preview.SupportBundleExported && !preview.SupportCaseCreated, "Consumer enablement, lookup routes, dispatch, and support side effects remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateCheck("production-and-host-boundary-closed", !preview.ProductionBusClaimed && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.BackendLaunchEnabled && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired, "Production, backend, network, privileged container, and host mutation gates remain closed."),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateCheck(id string, ok bool, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateCheck {
	status := "blocked"
	if ok {
		status = "pass"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateCountsFor(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGateCounts{Total: len(checks)}
	for _, check := range checks {
		if check.Status == "pass" {
			counts.Passed++
		} else {
			counts.Blocked++
		}
	}
	return counts
}
