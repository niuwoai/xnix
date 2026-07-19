package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditPreview struct {
	Version                                 string                                                                                                                                           `json:"version"`
	SchemaVersion                           string                                                                                                                                           `json:"schema_version"`
	RequestType                             string                                                                                                                                           `json:"request_type"`
	AuditType                               string                                                                                                                                           `json:"audit_type"`
	Source                                  string                                                                                                                                           `json:"source"`
	AuditDecision                           string                                                                                                                                           `json:"audit_decision"`
	CurrentMainlineConsumed                 bool                                                                                                                                             `json:"current_mainline_consumed"`
	WriterAuthorizationReceiptAuditConsumed bool                                                                                                                                             `json:"writer_authorization_receipt_audit_consumed"`
	WriterAuthorizationReceiptReady         bool                                                                                                                                             `json:"writer_authorization_receipt_ready"`
	AcceptanceAuthorizationRequired         bool                                                                                                                                             `json:"acceptance_authorization_required"`
	AcceptanceAuthorizationModeled          bool                                                                                                                                             `json:"acceptance_authorization_modeled"`
	AcceptanceAuthorizationReady            bool                                                                                                                                             `json:"acceptance_authorization_ready"`
	ReceiptAcceptanceBoundaryReady          bool                                                                                                                                             `json:"receipt_acceptance_boundary_ready"`
	OpaqueWriterAuthorizationReceipt        bool                                                                                                                                             `json:"opaque_writer_authorization_receipt"`
	KDESafeRedactedStatusOnly               bool                                                                                                                                             `json:"kde_safe_redacted_status_only"`
	CompatibilityCenterAcceptanceModeled    bool                                                                                                                                             `json:"compatibility_center_acceptance_modeled"`
	RuntimeDiagnosticsAcceptanceModeled     bool                                                                                                                                             `json:"runtime_diagnostics_acceptance_modeled"`
	ReceiptPresent                          bool                                                                                                                                             `json:"receipt_present"`
	ReceiptAccepted                         bool                                                                                                                                             `json:"receipt_accepted"`
	AuthorizationAccepted                   bool                                                                                                                                             `json:"authorization_accepted"`
	WriterAuthorizationGranted              bool                                                                                                                                             `json:"writer_authorization_granted"`
	StatusWriterEnabled                     bool                                                                                                                                             `json:"status_writer_enabled"`
	StatusPersistenceWriteEnabled           bool                                                                                                                                             `json:"status_persistence_write_enabled"`
	KDEStatusWriteEnabled                   bool                                                                                                                                             `json:"kde_status_write_enabled"`
	RuntimeDiagnosticsWriteEnabled          bool                                                                                                                                             `json:"runtime_diagnostics_write_enabled"`
	AcceptanceItemCount                     int                                                                                                                                              `json:"acceptance_item_count"`
	RequiredAcceptanceItemCount             int                                                                                                                                              `json:"required_acceptance_item_count"`
	ReadyAcceptanceItemCount                int                                                                                                                                              `json:"ready_acceptance_item_count"`
	MissingAcceptanceItemCount              int                                                                                                                                              `json:"missing_acceptance_item_count"`
	AcceptedReceiptItemCount                int                                                                                                                                              `json:"accepted_receipt_item_count"`
	GrantedWriterItemCount                  int                                                                                                                                              `json:"granted_writer_item_count"`
	EnabledWriterItemCount                  int                                                                                                                                              `json:"enabled_writer_item_count"`
	PersistedWriterItemCount                int                                                                                                                                              `json:"persisted_writer_item_count"`
	RawExposedAcceptanceItemCount           int                                                                                                                                              `json:"raw_exposed_acceptance_item_count"`
	SideEffectAcceptanceItemCount           int                                                                                                                                              `json:"side_effect_acceptance_item_count"`
	CompatibilityCenterAcceptanceItemCount  int                                                                                                                                              `json:"compatibility_center_acceptance_item_count"`
	RuntimeDiagnosticsAcceptanceItemCount   int                                                                                                                                              `json:"runtime_diagnostics_acceptance_item_count"`
	AcceptanceItems                         []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditItem      `json:"acceptance_items"`
	AcceptanceItemIDs                       []string                                                                                                                                         `json:"acceptance_item_ids"`
	Checks                                  []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditCheck     `json:"checks"`
	CheckIDs                                []string                                                                                                                                         `json:"check_ids"`
	Counts                                  ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditCheckCounts `json:"counts"`
	ConsumerConsumptionAuthorized           bool                                                                                                                                             `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized            bool                                                                                                                                             `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                      bool                                                                                                                                             `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                  bool                                                                                                                                             `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                   bool                                                                                                                                             `json:"lookup_route_authorized"`
	LookupRouteEnabled                      bool                                                                                                                                             `json:"lookup_route_enabled"`
	OpaqueLookupEnabled                     bool                                                                                                                                             `json:"opaque_lookup_enabled"`
	RedactedSummaryPersisted                bool                                                                                                                                             `json:"redacted_summary_persisted"`
	KDEStatusPersisted                      bool                                                                                                                                             `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted             bool                                                                                                                                             `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                   bool                                                                                                                                             `json:"dry_run_result_persisted"`
	RawResultExposed                        bool                                                                                                                                             `json:"raw_result_exposed"`
	DispatchDryRunExecuted                  bool                                                                                                                                             `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled            bool                                                                                                                                             `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled            bool                                                                                                                                             `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                    bool                                                                                                                                             `json:"portal_request_created"`
	NotificationActionEnabled               bool                                                                                                                                             `json:"notification_action_enabled"`
	CompatibilityCenterOpened               bool                                                                                                                                             `json:"compatibility_center_opened"`
	SupportBundleExported                   bool                                                                                                                                             `json:"support_bundle_exported"`
	SupportCaseCreated                      bool                                                                                                                                             `json:"support_case_created"`
	RuntimeOwned                            bool                                                                                                                                             `json:"runtime_owned"`
	GoRuntimeBacked                         bool                                                                                                                                             `json:"go_runtime_backed"`
	KDEPolicyOwner                          bool                                                                                                                                             `json:"kde_policy_owner"`
	ProductionReadiness                     bool                                                                                                                                             `json:"production_readiness"`
	ProductionOwnershipReady                bool                                                                                                                                             `json:"production_ownership_ready"`
	SystemServiceStarted                    bool                                                                                                                                             `json:"system_service_started"`
	SessionBusClaimed                       bool                                                                                                                                             `json:"session_bus_claimed"`
	ProductionBusClaimed                    bool                                                                                                                                             `json:"production_bus_claimed"`
	ProductionOwnerEnabled                  bool                                                                                                                                             `json:"production_owner_enabled"`
	WriteMethodsEnabled                     bool                                                                                                                                             `json:"write_methods_enabled"`
	RuntimeWritesEnabled                    bool                                                                                                                                             `json:"runtime_writes_enabled"`
	DesktopFilesWritten                     bool                                                                                                                                             `json:"desktop_files_written"`
	SettingsPersisted                       bool                                                                                                                                             `json:"settings_persisted"`
	AdapterInvocationEnabled                bool                                                                                                                                             `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                    bool                                                                                                                                             `json:"backend_launch_enabled"`
	BackendProcessStarted                   bool                                                                                                                                             `json:"backend_process_started"`
	SnapshotRestoreExecuted                 bool                                                                                                                                             `json:"snapshot_restore_executed"`
	StateCleanupExecuted                    bool                                                                                                                                             `json:"state_cleanup_executed"`
	NetworkRequired                         bool                                                                                                                                             `json:"network_required"`
	HostRootModified                        bool                                                                                                                                             `json:"host_root_modified"`
	PrivilegedContainerRequired             bool                                                                                                                                             `json:"privileged_container_required"`
	CallerStateRootRequired                 bool                                                                                                                                             `json:"caller_state_root_required"`
	StateRootPathExposed                    bool                                                                                                                                             `json:"state_root_path_exposed"`
	FilePathsExposed                        bool                                                                                                                                             `json:"file_paths_exposed"`
	FileContentRead                         bool                                                                                                                                             `json:"file_content_read"`
	RawCommandExposed                       bool                                                                                                                                             `json:"raw_command_exposed"`
	RawExecutableExposed                    bool                                                                                                                                             `json:"raw_executable_exposed"`
	BackendDetailsExposed                   bool                                                                                                                                             `json:"backend_details_exposed"`
	BlockedActions                          []string                                                                                                                                         `json:"blocked_actions"`
	NextRequirements                        []string                                                                                                                                         `json:"next_requirements"`
	DesktopSafeSummary                      string                                                                                                                                           `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditItem struct {
	ID                                 string `json:"id"`
	ActionKind                         string `json:"action_kind"`
	SurfaceKind                        string `json:"surface_kind"`
	StatusConsumerKind                 string `json:"status_consumer_kind"`
	OpaqueReceiptID                    string `json:"opaque_receipt_id"`
	EvidencePresent                    bool   `json:"evidence_present"`
	CurrentMainlineConsumed            bool   `json:"current_mainline_consumed"`
	WriterAuthorizationReceiptConsumed bool   `json:"writer_authorization_receipt_consumed"`
	AcceptanceAuthorizationModeled     bool   `json:"acceptance_authorization_modeled"`
	ReceiptAcceptanceBoundaryReady     bool   `json:"receipt_acceptance_boundary_ready"`
	OpaqueWriterAuthorizationReceipt   bool   `json:"opaque_writer_authorization_receipt"`
	KDESafeRedactedStatusOnly          bool   `json:"kde_safe_redacted_status_only"`
	ReceiptPresent                     bool   `json:"receipt_present"`
	ReceiptAccepted                    bool   `json:"receipt_accepted"`
	AuthorizationAccepted              bool   `json:"authorization_accepted"`
	WriterAuthorizationGranted         bool   `json:"writer_authorization_granted"`
	StatusWriterEnabled                bool   `json:"status_writer_enabled"`
	StatusPersistenceWriteEnabled      bool   `json:"status_persistence_write_enabled"`
	KDEStatusWriteEnabled              bool   `json:"kde_status_write_enabled"`
	RuntimeDiagnosticsWriteEnabled     bool   `json:"runtime_diagnostics_write_enabled"`
	RedactedSummaryPersisted           bool   `json:"redacted_summary_persisted"`
	KDEStatusPersisted                 bool   `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted        bool   `json:"runtime_diagnostics_persisted"`
	RawResultExposed                   bool   `json:"raw_result_exposed"`
	UserVisible                        bool   `json:"user_visible"`
	ReviewOnly                         bool   `json:"review_only"`
	RuntimeOwned                       bool   `json:"runtime_owned"`
	GoRuntimeBacked                    bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                     bool   `json:"kde_policy_owner"`
	SideEffectsDisabled                bool   `json:"side_effects_disabled"`
	HostRootModified                   bool   `json:"host_root_modified"`
	InternalDetailsExposed             bool   `json:"internal_details_exposed"`
	AcceptanceAuthorizationStatus      string `json:"acceptance_authorization_status"`
	NextRequirement                    string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditCheckCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceSources(root)
	mainlineReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceMainlineReady(sources.CurrentMainline)
	receiptReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceReceiptReady(sources.ReceiptAudit)
	items := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceItems(mainlineReady, receiptReady)
	readyItemCount := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceReadyCount(items)
	acceptanceReady := mainlineReady && receiptReady && readyItemCount == len(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditPreview{
		Version:                                 version,
		SchemaVersion:                           "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_authorization_receipt_acceptance_audit.v1",
		RequestType:                             "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-acceptance-audit-preview",
		AuditType:                               "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-acceptance-audit",
		Source:                                  "xnix-current-mainline+redacted-status-writer-authorization-receipt",
		AuditDecision:                           "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-acceptance-audit-blocked",
		CurrentMainlineConsumed:                 mainlineReady,
		WriterAuthorizationReceiptAuditConsumed: receiptReady,
		WriterAuthorizationReceiptReady:         receiptReady,
		AcceptanceAuthorizationRequired:         true,
		AcceptanceAuthorizationModeled:          true,
		AcceptanceAuthorizationReady:            acceptanceReady,
		ReceiptAcceptanceBoundaryReady:          acceptanceReady,
		OpaqueWriterAuthorizationReceipt:        receiptReady,
		KDESafeRedactedStatusOnly:               acceptanceReady,
		CompatibilityCenterAcceptanceModeled:    acceptanceReady,
		RuntimeDiagnosticsAcceptanceModeled:     acceptanceReady,
		ReceiptPresent:                          false,
		ReceiptAccepted:                         false,
		AuthorizationAccepted:                   false,
		WriterAuthorizationGranted:              false,
		StatusWriterEnabled:                     false,
		StatusPersistenceWriteEnabled:           false,
		KDEStatusWriteEnabled:                   false,
		RuntimeDiagnosticsWriteEnabled:          false,
		AcceptanceItemCount:                     len(items),
		RequiredAcceptanceItemCount:             len(items),
		ReadyAcceptanceItemCount:                readyItemCount,
		MissingAcceptanceItemCount:              len(items) - readyItemCount,
		AcceptedReceiptItemCount:                0,
		GrantedWriterItemCount:                  0,
		EnabledWriterItemCount:                  0,
		PersistedWriterItemCount:                0,
		RawExposedAcceptanceItemCount:           0,
		SideEffectAcceptanceItemCount:           0,
		CompatibilityCenterAcceptanceItemCount:  productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceSurfaceCount(items, "compatibility-center"),
		RuntimeDiagnosticsAcceptanceItemCount:   productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceSurfaceCount(items, "runtime-diagnostics"),
		AcceptanceItems:                         items,
		AcceptanceItemIDs:                       productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceItemIDs(items),
		ConsumerConsumptionAuthorized:           false,
		ConsumerEnablementAuthorized:            false,
		KDEConsumerEnabled:                      false,
		RuntimeConsumerEnabled:                  false,
		LookupRouteAuthorized:                   false,
		LookupRouteEnabled:                      false,
		OpaqueLookupEnabled:                     false,
		RedactedSummaryPersisted:                false,
		KDEStatusPersisted:                      false,
		RuntimeDiagnosticsPersisted:             false,
		DryRunResultPersisted:                   false,
		RawResultExposed:                        false,
		DispatchDryRunExecuted:                  false,
		RequestObjectCreationEnabled:            false,
		RequestObjectDispatchEnabled:            false,
		PortalRequestCreated:                    false,
		NotificationActionEnabled:               false,
		CompatibilityCenterOpened:               false,
		SupportBundleExported:                   false,
		SupportCaseCreated:                      false,
		RuntimeOwned:                            true,
		GoRuntimeBacked:                         true,
		KDEPolicyOwner:                          false,
		ProductionReadiness:                     false,
		ProductionOwnershipReady:                false,
		SystemServiceStarted:                    false,
		SessionBusClaimed:                       false,
		ProductionBusClaimed:                    false,
		ProductionOwnerEnabled:                  false,
		WriteMethodsEnabled:                     false,
		RuntimeWritesEnabled:                    false,
		DesktopFilesWritten:                     false,
		SettingsPersisted:                       false,
		AdapterInvocationEnabled:                false,
		BackendLaunchEnabled:                    false,
		BackendProcessStarted:                   false,
		SnapshotRestoreExecuted:                 false,
		StateCleanupExecuted:                    false,
		NetworkRequired:                         false,
		HostRootModified:                        false,
		PrivilegedContainerRequired:             false,
		CallerStateRootRequired:                 false,
		StateRootPathExposed:                    false,
		FilePathsExposed:                        false,
		FileContentRead:                         false,
		RawCommandExposed:                       false,
		RawExecutableExposed:                    false,
		BackendDetailsExposed:                   false,
		BlockedActions: []string{
			"treat this audit as permission to accept KDE-safe redacted status writer authorization receipts",
			"grant writer authorization or enable Compatibility Center or Runtime diagnostics status writers",
			"persist redacted status summaries, KDE status, Runtime diagnostics, dry-run results, or lookup records",
			"create request objects, dispatch actions, create Portal requests, send notifications, open Compatibility Center, export support bundles, or create support cases",
			"claim production ownership, start services, enable Runtime writes, launch compatibility engines, expose raw results or paths, or mutate host root",
		},
		NextRequirements: []string{
			"Implement a separate accepted writer authorization receipt gate before any redacted status writer can be enabled.",
			"Keep writer authorization grants, status writers, and status persistence writes disabled until that gate consumes accepted evidence.",
			"Keep consumer enablement, lookup routes, dry-run execution, result persistence, request creation, dispatch, Portal requests, notifications, Compatibility Center navigation, and support writes disabled.",
			"Keep production ownership, Runtime writes, desktop side effects, engine launch, raw result exposure, and host mutation disabled.",
		},
		DesktopSafeSummary: "The KDE-safe redacted status writer authorization receipt acceptance audit consumes the Xnix current mainline and the writer authorization receipt audit, models the future acceptance authorization boundary required before opaque writer receipts can be accepted, and keeps receipt acceptance, writer authorization grants, status writes, desktop side effects, unsafe data exposure, and host mutation disabled.",
	}
	checks := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-acceptance-audit-ready-acceptance-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dry-run result lookup consumer enablement kde safe redacted status writer authorization receipt acceptance audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceSourceSet struct {
	CurrentMainline string
	ReceiptAudit    string
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceSources(root string) productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceSourceSet {
	return productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceSourceSet{
		CurrentMainline: productionAuthorizationReadSources(root, []string{"docs/xnix-current-mainline.md"}),
		ReceiptAudit:    productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_authorization_receipt_audit.go", "internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_authorization_receipt_audit_test.go"}),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceMainlineReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"Xnix Current Mainline", "redacted status writer authorization receipt acceptance audit", "redacted status writer authorization receipt audit", "No Runtime write methods", "No KDE configuration writes"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceReceiptReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-audit-preview", "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-audit-ready-receipt-disabled", "WriterAuthorizationReceiptReady", "OpaqueWriterAuthorizationReceipt", "ReceiptAccepted", "WriterAuthorizationGranted", "StatusWriterEnabled", "CompatibilityCenterReceiptModeled", "RuntimeDiagnosticsReceiptModeled"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceItems(mainlineReady bool, receiptReady bool) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditItem {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditItem{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceItem("review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-authorization-receipt-acceptance", "review", "compatibility-center", "review-result-compatibility-center-status", "review-compatibility-center-redacted-status-writer-authorization-receipt-acceptance-id", mainlineReady, receiptReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceItem("review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-authorization-receipt-acceptance", "review", "runtime-diagnostics", "review-result-runtime-diagnostics-status", "review-runtime-diagnostics-redacted-status-writer-authorization-receipt-acceptance-id", mainlineReady, receiptReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceItem("renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-authorization-receipt-acceptance", "renew", "compatibility-center", "renewal-result-compatibility-center-status", "renewal-compatibility-center-redacted-status-writer-authorization-receipt-acceptance-id", mainlineReady, receiptReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceItem("renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-authorization-receipt-acceptance", "renew", "runtime-diagnostics", "renewal-result-runtime-diagnostics-status", "renewal-runtime-diagnostics-redacted-status-writer-authorization-receipt-acceptance-id", mainlineReady, receiptReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-authorization-receipt-acceptance", "open-compatibility-center", "compatibility-center", "navigation-result-compatibility-center-status", "navigation-compatibility-center-redacted-status-writer-authorization-receipt-acceptance-id", mainlineReady, receiptReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-authorization-receipt-acceptance", "open-compatibility-center", "runtime-diagnostics", "navigation-result-runtime-diagnostics-status", "navigation-runtime-diagnostics-redacted-status-writer-authorization-receipt-acceptance-id", mainlineReady, receiptReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-authorization-receipt-acceptance", "dismiss", "compatibility-center", "dismissal-result-compatibility-center-status", "dismissal-compatibility-center-redacted-status-writer-authorization-receipt-acceptance-id", mainlineReady, receiptReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-authorization-receipt-acceptance", "dismiss", "runtime-diagnostics", "dismissal-result-runtime-diagnostics-status", "dismissal-runtime-diagnostics-redacted-status-writer-authorization-receipt-acceptance-id", mainlineReady, receiptReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceItem("support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-authorization-receipt-acceptance", "support-info", "compatibility-center", "support-info-result-compatibility-center-status", "support-info-compatibility-center-redacted-status-writer-authorization-receipt-acceptance-id", mainlineReady, receiptReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceItem("support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-authorization-receipt-acceptance", "support-info", "runtime-diagnostics", "support-info-result-runtime-diagnostics-status", "support-info-runtime-diagnostics-redacted-status-writer-authorization-receipt-acceptance-id", mainlineReady, receiptReady),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceItem(id string, actionKind string, surfaceKind string, statusConsumerKind string, opaqueReceiptID string, mainlineReady bool, receiptReady bool) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditItem {
	ready := mainlineReady && receiptReady
	status := "missing-redacted-status-writer-authorization-receipt-acceptance-evidence"
	if ready {
		status = "redacted-status-writer-authorization-receipt-acceptance-modeled-acceptance-disabled"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditItem{
		ID:                                 id,
		ActionKind:                         actionKind,
		SurfaceKind:                        surfaceKind,
		StatusConsumerKind:                 statusConsumerKind,
		OpaqueReceiptID:                    opaqueReceiptID,
		EvidencePresent:                    ready,
		CurrentMainlineConsumed:            mainlineReady,
		WriterAuthorizationReceiptConsumed: receiptReady,
		AcceptanceAuthorizationModeled:     ready,
		ReceiptAcceptanceBoundaryReady:     ready,
		OpaqueWriterAuthorizationReceipt:   ready,
		KDESafeRedactedStatusOnly:          ready,
		ReceiptPresent:                     false,
		ReceiptAccepted:                    false,
		AuthorizationAccepted:              false,
		WriterAuthorizationGranted:         false,
		StatusWriterEnabled:                false,
		StatusPersistenceWriteEnabled:      false,
		KDEStatusWriteEnabled:              false,
		RuntimeDiagnosticsWriteEnabled:     false,
		RedactedSummaryPersisted:           false,
		KDEStatusPersisted:                 false,
		RuntimeDiagnosticsPersisted:        false,
		RawResultExposed:                   false,
		UserVisible:                        ready,
		ReviewOnly:                         true,
		RuntimeOwned:                       true,
		GoRuntimeBacked:                    true,
		KDEPolicyOwner:                     false,
		SideEffectsDisabled:                true,
		HostRootModified:                   false,
		InternalDetailsExposed:             false,
		AcceptanceAuthorizationStatus:      status,
		NextRequirement:                    "require separate accepted receipt gate before accepting this writer authorization receipt",
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditCheck{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceCheck("current-mainline-consumed", productionAuthorizationPassBlocked(preview.CurrentMainlineConsumed), "The receipt acceptance audit consumes the Xnix current mainline."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceCheck("writer-authorization-receipt-audit-consumed", productionAuthorizationPassBlocked(preview.WriterAuthorizationReceiptAuditConsumed && preview.WriterAuthorizationReceiptReady), "The receipt acceptance audit consumes the writer authorization receipt audit."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceCheck("acceptance-authorization-boundary-modeled", productionAuthorizationPassBlocked(preview.AcceptanceAuthorizationRequired && preview.AcceptanceAuthorizationModeled && preview.AcceptanceAuthorizationReady && preview.ReceiptAcceptanceBoundaryReady && preview.OpaqueWriterAuthorizationReceipt), "The receipt acceptance authorization boundary is modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceCheck("compatibility-center-and-runtime-acceptance-modeled", productionAuthorizationPassBlocked(preview.CompatibilityCenterAcceptanceModeled && preview.RuntimeDiagnosticsAcceptanceModeled && preview.CompatibilityCenterAcceptanceItemCount == 5 && preview.RuntimeDiagnosticsAcceptanceItemCount == 5), "Compatibility Center and Runtime diagnostics receipt acceptance candidates are both modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceCheck("ten-acceptance-items-ready-acceptance-disabled", productionAuthorizationPassBlocked(preview.AcceptanceItemCount == 10 && preview.RequiredAcceptanceItemCount == 10 && preview.ReadyAcceptanceItemCount == 10 && preview.MissingAcceptanceItemCount == 0 && preview.AcceptedReceiptItemCount == 0 && preview.GrantedWriterItemCount == 0 && preview.EnabledWriterItemCount == 0), "All redacted status writer authorization receipt acceptance candidates are present while acceptance and writers remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceCheck("receipt-acceptance-writers-and-persistence-disabled", productionAuthorizationPassBlocked(!preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.AuthorizationAccepted && !preview.WriterAuthorizationGranted && !preview.StatusWriterEnabled && !preview.StatusPersistenceWriteEnabled && !preview.KDEStatusWriteEnabled && !preview.RuntimeDiagnosticsWriteEnabled && !preview.RedactedSummaryPersisted && !preview.KDEStatusPersisted && !preview.RuntimeDiagnosticsPersisted && !preview.DryRunResultPersisted), "Receipt acceptance, writer authorization grants, status writers, and status persistence remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceCheck("consumer-lookup-request-and-support-disabled", productionAuthorizationPassBlocked(!preview.ConsumerConsumptionAuthorized && !preview.ConsumerEnablementAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled && !preview.LookupRouteAuthorized && !preview.LookupRouteEnabled && !preview.OpaqueLookupEnabled && !preview.RequestObjectCreationEnabled && !preview.RequestObjectDispatchEnabled && !preview.PortalRequestCreated && !preview.NotificationActionEnabled && !preview.CompatibilityCenterOpened && !preview.SupportBundleExported && !preview.SupportCaseCreated), "Consumers, lookup, request objects, desktop actions, and support writes remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.CallerStateRootRequired && !preview.StateRootPathExposed && !preview.FilePathsExposed && !preview.FileContentRead && !preview.RawResultExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && preview.RawExposedAcceptanceItemCount == 0 && preview.SideEffectAcceptanceItemCount == 0 && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceItemsKeepClosed(preview.AcceptanceItems)), "Production ownership, writes, unsafe data exposure, network, privilege, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceCheck(id string, status string, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditCheck {
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceSurfaceCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditItem, surfaceKind string) int {
	count := 0
	for _, item := range items {
		if item.SurfaceKind == surfaceKind {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceItemsKeepClosed(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditItem) bool {
	for _, item := range items {
		if item.ReceiptPresent || item.ReceiptAccepted || item.AuthorizationAccepted || item.WriterAuthorizationGranted || item.StatusWriterEnabled || item.StatusPersistenceWriteEnabled || item.KDEStatusWriteEnabled || item.RuntimeDiagnosticsWriteEnabled || item.RedactedSummaryPersisted || item.KDEStatusPersisted || item.RuntimeDiagnosticsPersisted || item.RawResultExposed || !item.SideEffectsDisabled || item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceChecks(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditCheckCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditCheckCounts{Total: len(checks)}
	for _, check := range checks {
		switch check.Status {
		case "pass":
			counts.Passed++
		case "pending":
			counts.Pending++
		default:
			counts.Blocked++
		}
	}
	return counts
}
