package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGatePreview struct {
	Version                                      string                                                                                                                                                                                                              `json:"version"`
	SchemaVersion                                string                                                                                                                                                                                                              `json:"schema_version"`
	RequestType                                  string                                                                                                                                                                                                              `json:"request_type"`
	GateType                                     string                                                                                                                                                                                                              `json:"gate_type"`
	Source                                       string                                                                                                                                                                                                              `json:"source"`
	AuditDecision                                string                                                                                                                                                                                                              `json:"audit_decision"`
	CurrentMainlineConsumed                      bool                                                                                                                                                                                                                `json:"current_mainline_consumed"`
	RouteAuthorizationEvidenceGateConsumed       bool                                                                                                                                                                                                                `json:"route_authorization_evidence_gate_consumed"`
	RouteAuthorizationEvidenceGateReady          bool                                                                                                                                                                                                                `json:"route_authorization_evidence_gate_ready"`
	RouteEnablementEvidenceGateRequired          bool                                                                                                                                                                                                                `json:"route_enablement_evidence_gate_required"`
	ResultConsumerProjectionRouteEnablementGated bool                                                                                                                                                                                                                `json:"result_consumer_projection_route_enablement_gated"`
	RouteEnablementEvidenceGateReady             bool                                                                                                                                                                                                                `json:"route_enablement_evidence_gate_ready"`
	CompatibilityCenterRouteEnablementGated      bool                                                                                                                                                                                                                `json:"compatibility_center_route_enablement_gated"`
	RuntimeDiagnosticsRouteEnablementGated       bool                                                                                                                                                                                                                `json:"runtime_diagnostics_route_enablement_gated"`
	RedactedProjectionRouteEnablementGated       bool                                                                                                                                                                                                                `json:"redacted_projection_route_enablement_gated"`
	UserVisibleRouteEnablementGated              bool                                                                                                                                                                                                                `json:"user_visible_route_enablement_gated"`
	FailureProjectionRouteEnablementGated        bool                                                                                                                                                                                                                `json:"failure_projection_route_enablement_gated"`
	KDESafeRedactedStatusOnly                    bool                                                                                                                                                                                                                `json:"kde_safe_redacted_status_only"`
	ReceiptPresent                               bool                                                                                                                                                                                                                `json:"receipt_present"`
	ReceiptAccepted                              bool                                                                                                                                                                                                                `json:"receipt_accepted"`
	ReceiptConsumed                              bool                                                                                                                                                                                                                `json:"receipt_consumed"`
	StorageRootPolicyGrantAuthorized             bool                                                                                                                                                                                                                `json:"storage_root_policy_grant_authorized"`
	RecordWriterCallAuthorized                   bool                                                                                                                                                                                                                `json:"record_writer_call_authorized"`
	WriterCallable                               bool                                                                                                                                                                                                                `json:"writer_callable"`
	StorageRootResolved                          bool                                                                                                                                                                                                                `json:"storage_root_resolved"`
	StorageWriteEnabled                          bool                                                                                                                                                                                                                `json:"storage_write_enabled"`
	StatusPersistenceWriteEnabled                bool                                                                                                                                                                                                                `json:"status_persistence_write_enabled"`
	EvidenceItemCount                            int                                                                                                                                                                                                                 `json:"evidence_item_count"`
	ReadyEvidenceItemCount                       int                                                                                                                                                                                                                 `json:"ready_evidence_item_count"`
	MissingEvidenceItemCount                     int                                                                                                                                                                                                                 `json:"missing_evidence_item_count"`
	CompatibilityCenterEvidenceItemCount         int                                                                                                                                                                                                                 `json:"compatibility_center_evidence_item_count"`
	RuntimeDiagnosticsEvidenceItemCount          int                                                                                                                                                                                                                 `json:"runtime_diagnostics_evidence_item_count"`
	RawExposedEvidenceItemCount                  int                                                                                                                                                                                                                 `json:"raw_exposed_evidence_item_count"`
	SideEffectEvidenceItemCount                  int                                                                                                                                                                                                                 `json:"side_effect_evidence_item_count"`
	EvidenceItems                                []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateItem  `json:"evidence_items"`
	EvidenceItemIDs                              []string                                                                                                                                                                                                            `json:"evidence_item_ids"`
	Checks                                       []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateCheck `json:"checks"`
	CheckIDs                                     []string                                                                                                                                                                                                            `json:"check_ids"`
	Counts                                       ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateCounts  `json:"counts"`
	ConsumerDryRunLookupResultBoundaryAuthorized bool                                                                                                                                                                                                                `json:"consumer_dry_run_lookup_result_boundary_authorized"`
	ConsumerEnablementAuthorized                 bool                                                                                                                                                                                                                `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                           bool                                                                                                                                                                                                                `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                       bool                                                                                                                                                                                                                `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                        bool                                                                                                                                                                                                                `json:"lookup_route_authorized"`
	LookupRouteEnablementAuthorized              bool                                                                                                                                                                                                                `json:"lookup_route_enablement_authorized"`
	LookupRouteEnabled                           bool                                                                                                                                                                                                                `json:"lookup_route_enabled"`
	OpaqueLookupEnabled                          bool                                                                                                                                                                                                                `json:"opaque_lookup_enabled"`
	RedactedSummaryPersisted                     bool                                                                                                                                                                                                                `json:"redacted_summary_persisted"`
	KDEStatusPersisted                           bool                                                                                                                                                                                                                `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted                  bool                                                                                                                                                                                                                `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                        bool                                                                                                                                                                                                                `json:"dry_run_result_persisted"`
	RawResultExposed                             bool                                                                                                                                                                                                                `json:"raw_result_exposed"`
	DispatchDryRunExecuted                       bool                                                                                                                                                                                                                `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled                 bool                                                                                                                                                                                                                `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled                 bool                                                                                                                                                                                                                `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                         bool                                                                                                                                                                                                                `json:"portal_request_created"`
	NotificationActionEnabled                    bool                                                                                                                                                                                                                `json:"notification_action_enabled"`
	CompatibilityCenterOpened                    bool                                                                                                                                                                                                                `json:"compatibility_center_opened"`
	SupportBundleExported                        bool                                                                                                                                                                                                                `json:"support_bundle_exported"`
	SupportCaseCreated                           bool                                                                                                                                                                                                                `json:"support_case_created"`
	RuntimeOwned                                 bool                                                                                                                                                                                                                `json:"runtime_owned"`
	GoRuntimeBacked                              bool                                                                                                                                                                                                                `json:"go_runtime_backed"`
	KDEPolicyOwner                               bool                                                                                                                                                                                                                `json:"kde_policy_owner"`
	ProductionReadiness                          bool                                                                                                                                                                                                                `json:"production_readiness"`
	ProductionOwnershipReady                     bool                                                                                                                                                                                                                `json:"production_ownership_ready"`
	SystemServiceStarted                         bool                                                                                                                                                                                                                `json:"system_service_started"`
	SessionBusClaimed                            bool                                                                                                                                                                                                                `json:"session_bus_claimed"`
	ProductionBusClaimed                         bool                                                                                                                                                                                                                `json:"production_bus_claimed"`
	WriteMethodsEnabled                          bool                                                                                                                                                                                                                `json:"write_methods_enabled"`
	RuntimeWritesEnabled                         bool                                                                                                                                                                                                                `json:"runtime_writes_enabled"`
	DesktopFilesWritten                          bool                                                                                                                                                                                                                `json:"desktop_files_written"`
	KDEConfigurationWritten                      bool                                                                                                                                                                                                                `json:"kde_configuration_written"`
	PortalCallExecuted                           bool                                                                                                                                                                                                                `json:"portal_call_executed"`
	AdapterInvocationEnabled                     bool                                                                                                                                                                                                                `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                         bool                                                                                                                                                                                                                `json:"backend_launch_enabled"`
	BackendProcessStarted                        bool                                                                                                                                                                                                                `json:"backend_process_started"`
	NetworkRequired                              bool                                                                                                                                                                                                                `json:"network_required"`
	HostRootModified                             bool                                                                                                                                                                                                                `json:"host_root_modified"`
	PrivilegedContainerRequired                  bool                                                                                                                                                                                                                `json:"privileged_container_required"`
	StateRootPathExposed                         bool                                                                                                                                                                                                                `json:"state_root_path_exposed"`
	FilePathsExposed                             bool                                                                                                                                                                                                                `json:"file_paths_exposed"`
	FileContentRead                              bool                                                                                                                                                                                                                `json:"file_content_read"`
	RawCommandExposed                            bool                                                                                                                                                                                                                `json:"raw_command_exposed"`
	RawExecutableExposed                         bool                                                                                                                                                                                                                `json:"raw_executable_exposed"`
	BackendDetailsExposed                        bool                                                                                                                                                                                                                `json:"backend_details_exposed"`
	BlockedActions                               []string                                                                                                                                                                                                            `json:"blocked_actions"`
	NextRequirements                             []string                                                                                                                                                                                                            `json:"next_requirements"`
	DesktopSafeSummary                           string                                                                                                                                                                                                              `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateItem struct {
	ID                                           string `json:"id"`
	ActionKind                                   string `json:"action_kind"`
	SurfaceKind                                  string `json:"surface_kind"`
	EvidenceKind                                 string `json:"evidence_kind"`
	EvidencePresent                              bool   `json:"evidence_present"`
	CurrentMainlineConsumed                      bool   `json:"current_mainline_consumed"`
	RouteAuthorizationEvidenceGateConsumed       bool   `json:"route_authorization_evidence_gate_consumed"`
	RouteAuthorizationEvidenceGateReady          bool   `json:"route_authorization_evidence_gate_ready"`
	ResultConsumerProjectionRouteEnablementGated bool   `json:"result_consumer_projection_route_enablement_gated"`
	RedactedProjectionRouteEnablementGated       bool   `json:"redacted_projection_route_enablement_gated"`
	UserVisibleRouteEnablementGated              bool   `json:"user_visible_route_enablement_gated"`
	FailureProjectionRouteEnablementGated        bool   `json:"failure_projection_route_enablement_gated"`
	KDESafeRedactedStatusOnly                    bool   `json:"kde_safe_redacted_status_only"`
	LookupRouteEnablementAuthorized              bool   `json:"lookup_route_enablement_authorized"`
	LookupRouteEnabled                           bool   `json:"lookup_route_enabled"`
	ConsumerEnablementAuthorized                 bool   `json:"consumer_enablement_authorized"`
	StorageWriteEnabled                          bool   `json:"storage_write_enabled"`
	RawResultExposed                             bool   `json:"raw_result_exposed"`
	UserVisible                                  bool   `json:"user_visible"`
	ReviewOnly                                   bool   `json:"review_only"`
	RuntimeOwned                                 bool   `json:"runtime_owned"`
	GoRuntimeBacked                              bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                               bool   `json:"kde_policy_owner"`
	SideEffectsDisabled                          bool   `json:"side_effects_disabled"`
	HostRootModified                             bool   `json:"host_root_modified"`
	InternalDetailsExposed                       bool   `json:"internal_details_exposed"`
	RouteEnablementGateStatus                    string `json:"route_enablement_gate_status"`
	NextRequirement                              string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

type productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateSourceSet struct {
	CurrentMainline                string
	RouteAuthorizationEvidenceGate string
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGatePreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGatePreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGatePreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateSources(root)
	mainlineReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateMainlineReady(sources.CurrentMainline)
	authorizationGateReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateAuthorizationReady(sources.RouteAuthorizationEvidenceGate)
	ready := mainlineReady && authorizationGateReady
	items := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateItems(mainlineReady, authorizationGateReady)
	readyCount := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateReadyCount(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGatePreview{
		Version:                                version,
		SchemaVersion:                          "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_evidence_gate.v1",
		RequestType:                            "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-evidence-gate-preview",
		GateType:                               "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-evidence-gate",
		Source:                                 "xnix-current-mainline+redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-authorization-evidence-gate",
		AuditDecision:                          "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-evidence-gate-blocked",
		CurrentMainlineConsumed:                mainlineReady,
		RouteAuthorizationEvidenceGateConsumed: authorizationGateReady,
		RouteAuthorizationEvidenceGateReady:    authorizationGateReady,
		RouteEnablementEvidenceGateRequired:    true,
		ResultConsumerProjectionRouteEnablementGated: ready,
		RouteEnablementEvidenceGateReady:             ready,
		CompatibilityCenterRouteEnablementGated:      ready,
		RuntimeDiagnosticsRouteEnablementGated:       ready,
		RedactedProjectionRouteEnablementGated:       ready,
		UserVisibleRouteEnablementGated:              ready,
		FailureProjectionRouteEnablementGated:        ready,
		KDESafeRedactedStatusOnly:                    ready,
		EvidenceItemCount:                            len(items),
		ReadyEvidenceItemCount:                       readyCount,
		MissingEvidenceItemCount:                     len(items) - readyCount,
		CompatibilityCenterEvidenceItemCount:         productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateSurfaceCount(items, "compatibility-center"),
		RuntimeDiagnosticsEvidenceItemCount:          productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateSurfaceCount(items, "runtime-diagnostics"),
		EvidenceItems:                                items,
		EvidenceItemIDs:                              productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateItemIDs(items),
		RuntimeOwned:                                 true,
		GoRuntimeBacked:                              true,
		KDEPolicyOwner:                               false,
		BlockedActions: []string{
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
			"Add an operator-reviewed lookup route enablement receipt before lookup routes can become callable.",
			"Keep route enablement evidence redacted and review-only until production ownership is explicitly authorized.",
		},
		DesktopSafeSummary: "KDE-safe redacted dry-run lookup result consumer projection route enablement evidence is gated for Compatibility Center and Runtime diagnostics while the lookup route itself and every consumer, dispatch, persistence, production, backend, and host side effect remain disabled.",
	}
	if ready {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-evidence-gate-ready-routes-disabled"
	}
	preview.Checks = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateChecks(preview)
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateCheckIDs(preview.Checks)
	preview.Counts = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateCountsFor(preview.Checks)
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status dry-run lookup result consumer projection route enablement evidence gate preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGatePreview{}, err
	}
	return preview, nil
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateSources(root string) productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateSourceSet {
	return productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateSourceSet{
		CurrentMainline: productionAuthorizationReadSources(root, []string{"docs/xnix-current-mainline.md"}),
		RouteAuthorizationEvidenceGate: productionAuthorizationReadSources(root, []string{
			"internal/runtime/owner/route_authorization_evidence_gate_preview.go",
			"internal/runtime/owner/route_authorization_evidence_gate_preview_test.go",
		}),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateMainlineReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"route enablement evidence gate preview",
		"route authorization evidence gate preview",
		"without accepting or consuming receipts",
	})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateAuthorizationReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_authorization_evidence_gate.v1",
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-authorization-evidence-gate-ready-routes-disabled",
		"RouteAuthorizationEvidenceGateReady",
		"ResultConsumerProjectionRouteAuthorizationEvidenceGated",
	})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateItems(mainlineReady, authorizationGateReady bool) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateItem {
	actionKinds := []string{"review", "renew", "open-compatibility-center", "dismiss", "support-info"}
	items := make([]ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateItem, 0, len(actionKinds)*2)
	for _, actionKind := range actionKinds {
		for _, surfaceKind := range []string{"compatibility-center", "runtime-diagnostics"} {
			ready := mainlineReady && authorizationGateReady
			status := "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-evidence-gate"
			if ready {
				status = "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-evidence-gated-routes-disabled"
			}
			items = append(items, ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateItem{
				ID:                                     "result-consumer-route-enablement-evidence-" + surfaceKind + "-" + actionKind,
				ActionKind:                             actionKind,
				SurfaceKind:                            surfaceKind,
				EvidenceKind:                           "kde-safe-redacted-lookup-result-projection-route-enablement-evidence",
				EvidencePresent:                        ready,
				CurrentMainlineConsumed:                mainlineReady,
				RouteAuthorizationEvidenceGateConsumed: authorizationGateReady,
				RouteAuthorizationEvidenceGateReady:    authorizationGateReady,
				ResultConsumerProjectionRouteEnablementGated: ready,
				RedactedProjectionRouteEnablementGated:       ready,
				UserVisibleRouteEnablementGated:              ready,
				FailureProjectionRouteEnablementGated:        ready,
				KDESafeRedactedStatusOnly:                    ready,
				UserVisible:                                  ready,
				ReviewOnly:                                   ready,
				RuntimeOwned:                                 true,
				GoRuntimeBacked:                              true,
				KDEPolicyOwner:                               false,
				SideEffectsDisabled:                          true,
				RouteEnablementGateStatus:                    status,
				NextRequirement:                              "Require a separate operator-reviewed route enablement receipt before enabling lookup routes, consumers, persistence, dry-run dispatch, or production side effects.",
			})
		}
	}
	return items
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent && item.ResultConsumerProjectionRouteEnablementGated && item.RedactedProjectionRouteEnablementGated && item.UserVisibleRouteEnablementGated && item.FailureProjectionRouteEnablementGated && item.KDESafeRedactedStatusOnly && !item.LookupRouteEnabled && !item.StorageWriteEnabled && !item.RawResultExposed && !item.HostRootModified && !item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateSurfaceCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateItem, surfaceKind string) int {
	count := 0
	for _, item := range items {
		if item.SurfaceKind == surfaceKind {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGatePreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateCheck{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateCheck("current-mainline-consumed", preview.CurrentMainlineConsumed, "Xnix current mainline names the result consumer projection route enablement evidence gate preview."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateCheck("route-authorization-evidence-gate-consumed", preview.RouteAuthorizationEvidenceGateConsumed && preview.RouteAuthorizationEvidenceGateReady, "The gate consumes the ready route authorization evidence gate."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateCheck("result-consumer-projection-route-enablement-gated", preview.ResultConsumerProjectionRouteEnablementGated, "Projection route enablement identity, user-visible redaction, and failure evidence are gated."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateCheck("compatibility-center-and-runtime-route-enablement-gated", preview.CompatibilityCenterRouteEnablementGated && preview.RuntimeDiagnosticsRouteEnablementGated, "Compatibility Center and Runtime diagnostics route enablement boundaries are gated."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateCheck("ten-route-enablement-evidence-items-ready-routes-disabled", preview.EvidenceItemCount == 10 && preview.ReadyEvidenceItemCount == 10 && preview.MissingEvidenceItemCount == 0, "Ten KDE-safe route enablement evidence items are ready while lookup routes remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateCheck("route-enablement-evidence-and-writes-disabled", !preview.StorageWriteEnabled && !preview.StatusPersistenceWriteEnabled && !preview.DryRunResultPersisted && !preview.RawResultExposed, "Route enablement evidence gate does not enable storage, status persistence, dry-run result persistence, or raw result exposure."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateCheck("consumer-routes-dispatch-and-support-disabled", !preview.ConsumerEnablementAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled && !preview.LookupRouteEnabled && !preview.DispatchDryRunExecuted && !preview.SupportBundleExported && !preview.SupportCaseCreated, "Consumer enablement, lookup routes, dispatch, and support side effects remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateCheck("production-and-host-boundary-closed", !preview.ProductionBusClaimed && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.BackendLaunchEnabled && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired, "Production, backend, network, privileged container, and host mutation gates remain closed."),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateCheck(id string, ok bool, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateCheck {
	status := "blocked"
	if ok {
		status = "pass"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateCountsFor(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGateCounts{Total: len(checks)}
	for _, check := range checks {
		if check.Status == "pass" {
			counts.Passed++
		} else {
			counts.Blocked++
		}
	}
	return counts
}
