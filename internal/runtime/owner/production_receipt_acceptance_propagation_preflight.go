package owner

type ProductionReceiptAcceptancePropagationPreflightPreview struct {
	Version                         string                                                     `json:"version"`
	SchemaVersion                   string                                                     `json:"schema_version"`
	RequestType                     string                                                     `json:"request_type"`
	PreflightType                   string                                                     `json:"preflight_type"`
	Source                          string                                                     `json:"source"`
	PreflightDecision               string                                                     `json:"preflight_decision"`
	ReceiptSchema                   string                                                     `json:"receipt_schema"`
	OpaqueReceiptID                 string                                                     `json:"opaque_receipt_id"`
	ReceiptRequired                 bool                                                       `json:"receipt_required"`
	ReceiptPresent                  bool                                                       `json:"receipt_present"`
	ReceiptAccepted                 bool                                                       `json:"receipt_accepted"`
	FutureAcceptanceModeled         bool                                                       `json:"future_acceptance_modeled"`
	AcceptanceSimulationOnly        bool                                                       `json:"acceptance_simulation_only"`
	ConsumptionAuditConsumed        bool                                                       `json:"consumption_audit_consumed"`
	OwnerManagedOpaqueBoundaryReady bool                                                       `json:"owner_managed_opaque_boundary_ready"`
	CallerStateRootRequired         bool                                                       `json:"caller_state_root_required"`
	AuthorizationAccepted           bool                                                       `json:"authorization_accepted"`
	PropagationPreflightReady       bool                                                       `json:"propagation_preflight_ready"`
	ProductionReadiness             bool                                                       `json:"production_readiness"`
	ProductionOwnershipReady        bool                                                       `json:"production_ownership_ready"`
	TargetCount                     int                                                        `json:"target_count"`
	RequiredTargetCount             int                                                        `json:"required_target_count"`
	PropagationReadyTargetCount     int                                                        `json:"propagation_ready_target_count"`
	MissingTargetCount              int                                                        `json:"missing_target_count"`
	AcceptanceEnabledTargetCount    int                                                        `json:"acceptance_enabled_target_count"`
	ProductionReadyTargetCount      int                                                        `json:"production_ready_target_count"`
	SideEffectTargetCount           int                                                        `json:"side_effect_target_count"`
	Targets                         []ProductionReceiptAcceptancePropagationTarget             `json:"targets"`
	TargetIDs                       []string                                                   `json:"target_ids"`
	RequiredBeforeReceiptAcceptance []string                                                   `json:"required_before_receipt_acceptance"`
	Checks                          []ProductionReceiptAcceptancePropagationPreflightCheck     `json:"checks"`
	CheckIDs                        []string                                                   `json:"check_ids"`
	Counts                          ProductionReceiptAcceptancePropagationPreflightCheckCounts `json:"counts"`
	RuntimeOwned                    bool                                                       `json:"runtime_owned"`
	GoRuntimeBacked                 bool                                                       `json:"go_runtime_backed"`
	KDEPolicyOwner                  bool                                                       `json:"kde_policy_owner"`
	SystemServiceStarted            bool                                                       `json:"system_service_started"`
	SessionBusClaimed               bool                                                       `json:"session_bus_claimed"`
	ProductionBusClaimed            bool                                                       `json:"production_bus_claimed"`
	ProductionOwnerEnabled          bool                                                       `json:"production_owner_enabled"`
	ProductionActivationReady       bool                                                       `json:"production_activation_ready"`
	WriteMethodsEnabled             bool                                                       `json:"write_methods_enabled"`
	RuntimeWritesEnabled            bool                                                       `json:"runtime_writes_enabled"`
	ReceiptWriterEnabled            bool                                                       `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled       bool                                                       `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled      bool                                                       `json:"receipt_lookup_writes_enabled"`
	DesktopFilesWritten             bool                                                       `json:"desktop_files_written"`
	MIMEAppsWritten                 bool                                                       `json:"mimeapps_written"`
	ShellConfigurationWritten       bool                                                       `json:"shell_configuration_written"`
	SettingsPersisted               bool                                                       `json:"settings_persisted"`
	NotificationSent                bool                                                       `json:"notification_sent"`
	NotificationDeliveryEnabled     bool                                                       `json:"notification_delivery_enabled"`
	PortalRequestCreated            bool                                                       `json:"portal_request_created"`
	RequestObjectsCreated           bool                                                       `json:"request_objects_created"`
	AdapterInvocationEnabled        bool                                                       `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled            bool                                                       `json:"backend_launch_enabled"`
	BackendProcessStarted           bool                                                       `json:"backend_process_started"`
	SupportBundleExported           bool                                                       `json:"support_bundle_exported"`
	SupportCaseCreated              bool                                                       `json:"support_case_created"`
	SnapshotRestoreExecuted         bool                                                       `json:"snapshot_restore_executed"`
	StateCleanupExecuted            bool                                                       `json:"state_cleanup_executed"`
	FileContentRead                 bool                                                       `json:"file_content_read"`
	FilePathsExposed                bool                                                       `json:"file_paths_exposed"`
	NetworkRequired                 bool                                                       `json:"network_required"`
	HostRootModified                bool                                                       `json:"host_root_modified"`
	PrivilegedContainerRequired     bool                                                       `json:"privileged_container_required"`
	StateRootPathExposed            bool                                                       `json:"state_root_path_exposed"`
	RawCommandExposed               bool                                                       `json:"raw_command_exposed"`
	RawExecutableExposed            bool                                                       `json:"raw_executable_exposed"`
	BackendDetailsExposed           bool                                                       `json:"backend_details_exposed"`
	BlockedActions                  []string                                                   `json:"blocked_actions"`
	NextRequirements                []string                                                   `json:"next_requirements"`
	DesktopSafeSummary              string                                                     `json:"desktop_safe_summary"`
}

type ProductionReceiptAcceptancePropagationTarget struct {
	ID                         string `json:"id"`
	RequestType                string `json:"request_type"`
	SourceFile                 string `json:"source_file"`
	InputState                 string `json:"input_state"`
	OutputState                string `json:"output_state"`
	ConsumesAuditBoundary      bool   `json:"consumes_audit_boundary"`
	PropagatesFutureAcceptance bool   `json:"propagates_future_acceptance"`
	ReceiptBoundaryReady       bool   `json:"receipt_boundary_ready"`
	FutureAcceptanceModeled    bool   `json:"future_acceptance_modeled"`
	ReceiptAccepted            bool   `json:"receipt_accepted"`
	AuthorizationAccepted      bool   `json:"authorization_accepted"`
	ProductionReadiness        bool   `json:"production_readiness"`
	ProductionOwnershipReady   bool   `json:"production_ownership_ready"`
	RuntimeOwned               bool   `json:"runtime_owned"`
	GoRuntimeBacked            bool   `json:"go_runtime_backed"`
	KDEPolicyOwner             bool   `json:"kde_policy_owner"`
	ReviewOnly                 bool   `json:"review_only"`
	WriteMethodsEnabled        bool   `json:"write_methods_enabled"`
	RuntimeWritesEnabled       bool   `json:"runtime_writes_enabled"`
	DesktopSideEffectsEnabled  bool   `json:"desktop_side_effects_enabled"`
	SupportSideEffectsEnabled  bool   `json:"support_side_effects_enabled"`
	BackendLaunchEnabled       bool   `json:"backend_launch_enabled"`
	HostRootModified           bool   `json:"host_root_modified"`
	InternalDetailsExposed     bool   `json:"internal_details_exposed"`
	PropagationStatus          string `json:"propagation_status"`
	NextRequirement            string `json:"next_requirement"`
}

type ProductionReceiptAcceptancePropagationPreflightCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptAcceptancePropagationPreflightCheckCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptAcceptancePropagationPreflightPreview(root string) (ProductionReceiptAcceptancePropagationPreflightPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptAcceptancePropagationPreflightPreview{}, err
	}
	sources := productionReceiptAcceptancePropagationSources(root)
	targets := productionReceiptAcceptancePropagationTargets(sources)
	consumptionAuditReady := productionReceiptAcceptanceConsumptionAuditReady(sources.ConsumptionAudit)
	preview := ProductionReceiptAcceptancePropagationPreflightPreview{
		Version:                         version,
		SchemaVersion:                   "xnix.runtime.production_receipt_acceptance_propagation_preflight.v1",
		RequestType:                     "production-receipt-acceptance-propagation-preflight-preview",
		PreflightType:                   "future-authorization-receipt-acceptance-propagation-preflight",
		Source:                          "production-authorization-consumption-audit-preview+production-dbus-gate-review-preview+production-dbus-method-review-preview+runtime-service-activation-preflight-preview+runtime-write-gate-preview+production-rollback-diagnostics-review-preview+production-desktop-side-effect-review-preview",
		PreflightDecision:               "production-receipt-acceptance-propagation-preflight-blocked",
		ReceiptSchema:                   "xnix.runtime.production_dbus_human_authorization_receipt.v1",
		OpaqueReceiptID:                 ProductionDBusHumanAuthorizationReceiptID,
		ReceiptRequired:                 true,
		ReceiptPresent:                  false,
		ReceiptAccepted:                 false,
		FutureAcceptanceModeled:         true,
		AcceptanceSimulationOnly:        true,
		ConsumptionAuditConsumed:        consumptionAuditReady,
		OwnerManagedOpaqueBoundaryReady: consumptionAuditReady,
		CallerStateRootRequired:         false,
		AuthorizationAccepted:           false,
		PropagationPreflightReady:       consumptionAuditReady && productionReceiptAcceptanceAllTargetsReady(targets),
		ProductionReadiness:             false,
		ProductionOwnershipReady:        false,
		TargetCount:                     len(targets),
		RequiredTargetCount:             6,
		PropagationReadyTargetCount:     productionReceiptAcceptanceReadyTargetCount(targets),
		MissingTargetCount:              productionReceiptAcceptanceMissingTargetCount(targets),
		AcceptanceEnabledTargetCount:    0,
		ProductionReadyTargetCount:      0,
		SideEffectTargetCount:           0,
		Targets:                         targets,
		TargetIDs:                       productionReceiptAcceptanceTargetIDs(targets),
		RequiredBeforeReceiptAcceptance: []string{
			"production-authorization-consumption-audit-preview",
			"production-dbus-gate-review-preview",
			"production-dbus-method-review-preview",
			"runtime-service-activation-preflight-preview",
			"runtime-write-gate-preview",
			"production-rollback-diagnostics-review-preview",
			"production-desktop-side-effect-review-preview",
			"separate operator action authorizing receipt acceptance outside this preflight",
			"separate receipt writer and persistence review outside this preflight",
			"separate production service ownership proof outside this preflight",
		},
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		SystemServiceStarted:        false,
		SessionBusClaimed:           false,
		ProductionBusClaimed:        false,
		ProductionOwnerEnabled:      false,
		ProductionActivationReady:   false,
		WriteMethodsEnabled:         false,
		RuntimeWritesEnabled:        false,
		ReceiptWriterEnabled:        false,
		ReceiptPersistenceEnabled:   false,
		ReceiptLookupWritesEnabled:  false,
		DesktopFilesWritten:         false,
		MIMEAppsWritten:             false,
		ShellConfigurationWritten:   false,
		SettingsPersisted:           false,
		NotificationSent:            false,
		NotificationDeliveryEnabled: false,
		PortalRequestCreated:        false,
		RequestObjectsCreated:       false,
		AdapterInvocationEnabled:    false,
		BackendLaunchEnabled:        false,
		BackendProcessStarted:       false,
		SupportBundleExported:       false,
		SupportCaseCreated:          false,
		SnapshotRestoreExecuted:     false,
		StateCleanupExecuted:        false,
		FileContentRead:             false,
		FilePathsExposed:            false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		StateRootPathExposed:        false,
		RawCommandExposed:           false,
		RawExecutableExposed:        false,
		BackendDetailsExposed:       false,
		BlockedActions: []string{
			"treat future acceptance modeling as an accepted authorization receipt",
			"write, persist, accept, replay, or look up authorization receipts from this preflight",
			"claim production D-Bus ownership, start services, or enable Runtime writes from this preflight",
			"enable desktop writes, notifications, Portal requests, request objects, adapter invocation, or compatibility engine launch",
			"export support bundles, create support cases, restore snapshots, clean state, read file contents, or expose paths",
			"require network, require privileged containers, expose raw commands, expose internal engine details, or mutate host root",
		},
		NextRequirements: []string{
			"Keep the accepted receipt path simulated until a separate operator action authorizes real receipt acceptance.",
			"Route any future accepted opaque receipt through the same six production gates before ownership can be considered.",
			"Keep service start, bus claim, write dispatch, desktop side effects, support side effects, restore, cleanup, launch, and host mutation disabled.",
			"Add a separate acceptance writer review before any persistent receipt record exists.",
		},
		DesktopSafeSummary: "The production receipt acceptance propagation preflight models how a future accepted opaque authorization receipt would flow through production gates, but it does not accept authorization, write receipts, claim a bus, start services, enable writes, emit desktop side effects, launch engines, or mutate host state.",
	}
	checks := productionReceiptAcceptancePropagationChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptAcceptanceCheckIDs(checks)
	preview.Counts = countProductionReceiptAcceptanceChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.PreflightDecision = "production-receipt-acceptance-propagation-ready-acceptance-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt acceptance propagation preflight preview"); err != nil {
		return ProductionReceiptAcceptancePropagationPreflightPreview{}, err
	}
	return preview, nil
}

type productionReceiptAcceptancePropagationSourceSet struct {
	ConsumptionAudit    string
	ProductionDBusGate  string
	MethodReview        string
	ServiceActivation   string
	RuntimeWriteGate    string
	RollbackDiagnostics string
	DesktopSideEffects  string
}

func productionReceiptAcceptancePropagationSources(root string) productionReceiptAcceptancePropagationSourceSet {
	return productionReceiptAcceptancePropagationSourceSet{
		ConsumptionAudit:    productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_authorization_consumption_audit.go"}),
		ProductionDBusGate:  productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_dbus_gate_review.go"}),
		MethodReview:        productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_dbus_method_review.go"}),
		ServiceActivation:   productionAuthorizationReadSources(root, []string{"internal/runtime/appidentity/runtime_service_activation_preflight.go"}),
		RuntimeWriteGate:    productionAuthorizationReadSources(root, []string{"internal/runtime/appidentity/runtime_write_gate.go"}),
		RollbackDiagnostics: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_rollback_diagnostics_review.go"}),
		DesktopSideEffects:  productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_desktop_side_effect_review.go"}),
	}
}

func productionReceiptAcceptancePropagationTargets(sources productionReceiptAcceptancePropagationSourceSet) []ProductionReceiptAcceptancePropagationTarget {
	return []ProductionReceiptAcceptancePropagationTarget{
		productionReceiptAcceptanceTarget("production-dbus-gate-review", "production-dbus-gate-review-preview", "internal/runtime/owner/production_dbus_gate_review.go", sources.ProductionDBusGate, []string{"production-human-authorization-receipt-consolidation-preview", "AuthorizationReceiptAccepted", "ProductionReadiness"}, "feed future receipt acceptance into the production ownership decision only after a separate operator action"),
		productionReceiptAcceptanceTarget("production-dbus-method-review", "production-dbus-method-review-preview", "internal/runtime/owner/production_dbus_method_review.go", sources.MethodReview, []string{"production-human-authorization-receipt-consolidation-preview", "consolidated opaque authorization receipt boundary is accepted", "production-exposure-disabled"}, "keep method exposure disabled until receipt acceptance and production ownership are separately authorized"),
		productionReceiptAcceptanceTarget("runtime-service-activation-preflight", "runtime-service-activation-preflight-preview", "internal/runtime/appidentity/runtime_service_activation_preflight.go", sources.ServiceActivation, []string{"production-human-authorization-receipt-consolidation-preview", "AuthorizationReceiptAccepted", "ProductionActivationReady"}, "keep service activation disabled until the accepted receipt and service ownership proof exist"),
		productionReceiptAcceptanceTarget("runtime-write-gate", "runtime-write-gate-preview", "internal/runtime/appidentity/runtime_write_gate.go", sources.RuntimeWriteGate, []string{"production-human-authorization-receipt-consolidation-preview", "AuthorizationReceiptAccepted", "WriteMethodEnabled"}, "keep write dispatch disabled until every production gate accepts the receipt and ownership is proven"),
		productionReceiptAcceptanceTarget("rollback-diagnostics-review", "production-rollback-diagnostics-review-preview", "internal/runtime/owner/production_rollback_diagnostics_review.go", sources.RollbackDiagnostics, []string{"production-human-authorization-receipt-consolidation-preview", "consolidated opaque authorization receipt boundary", "ProductionOwnershipReady"}, "keep rollback and diagnostics side effects disabled until ownership commit exists"),
		productionReceiptAcceptanceTarget("desktop-side-effect-review", "production-desktop-side-effect-review-preview", "internal/runtime/owner/production_desktop_side_effect_review.go", sources.DesktopSideEffects, []string{"production-human-authorization-receipt-consolidation-preview", "consolidated human authorization receipt boundary", "ProductionOwnershipReady"}, "keep KDE desktop side effects disabled until production ownership commit exists"),
	}
}

func productionReceiptAcceptanceTarget(id string, requestType string, sourceFile string, source string, tokens []string, nextRequirement string) ProductionReceiptAcceptancePropagationTarget {
	ready := productionAuthorizationHasAll(source, tokens)
	status := "missing-propagation-source"
	if ready {
		status = "future-acceptance-modeled-side-effects-disabled"
	}
	return ProductionReceiptAcceptancePropagationTarget{
		ID:                         id,
		RequestType:                requestType,
		SourceFile:                 sourceFile,
		InputState:                 "future-accepted-opaque-authorization-receipt",
		OutputState:                "review-only-propagation-with-production-side-effects-disabled",
		ConsumesAuditBoundary:      ready,
		PropagatesFutureAcceptance: ready,
		ReceiptBoundaryReady:       ready,
		FutureAcceptanceModeled:    ready,
		ReceiptAccepted:            false,
		AuthorizationAccepted:      false,
		ProductionReadiness:        false,
		ProductionOwnershipReady:   false,
		RuntimeOwned:               true,
		GoRuntimeBacked:            true,
		KDEPolicyOwner:             false,
		ReviewOnly:                 true,
		WriteMethodsEnabled:        false,
		RuntimeWritesEnabled:       false,
		DesktopSideEffectsEnabled:  false,
		SupportSideEffectsEnabled:  false,
		BackendLaunchEnabled:       false,
		HostRootModified:           false,
		InternalDetailsExposed:     false,
		PropagationStatus:          status,
		NextRequirement:            nextRequirement,
	}
}

func productionReceiptAcceptancePropagationChecks(preview ProductionReceiptAcceptancePropagationPreflightPreview) []ProductionReceiptAcceptancePropagationPreflightCheck {
	return []ProductionReceiptAcceptancePropagationPreflightCheck{
		productionReceiptAcceptanceCheck("consumption-audit-consumed", productionAuthorizationPassBlocked(preview.ConsumptionAuditConsumed && preview.OwnerManagedOpaqueBoundaryReady && preview.ReceiptSchema == "xnix.runtime.production_dbus_human_authorization_receipt.v1" && preview.OpaqueReceiptID == ProductionDBusHumanAuthorizationReceiptID), "The preflight consumes the production authorization consumption audit and opaque receipt boundary."),
		productionReceiptAcceptanceCheck("future-acceptance-modeled-only", productionAuthorizationPassBlocked(preview.ReceiptRequired && preview.FutureAcceptanceModeled && preview.AcceptanceSimulationOnly && !preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.AuthorizationAccepted), "The preflight models future acceptance without accepting a receipt."),
		productionReceiptAcceptanceCheck("six-propagation-targets-present", productionAuthorizationPassBlocked(preview.TargetCount == 6 && preview.RequiredTargetCount == 6 && preview.MissingTargetCount == 0), "The preflight tracks all six production propagation targets."),
		productionReceiptAcceptanceCheck("targets-propagate-future-acceptance", productionAuthorizationPassBlocked(preview.PropagationReadyTargetCount == 6 && productionReceiptAcceptanceAllTargetsReady(preview.Targets)), "Every production target can model future receipt acceptance propagation while keeping actions disabled."),
		productionReceiptAcceptanceCheck("acceptance-and-production-disabled", productionAuthorizationPassBlocked(!preview.PropagationPreflightReady || (!preview.ReceiptAccepted && !preview.AuthorizationAccepted && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && preview.AcceptanceEnabledTargetCount == 0 && preview.ProductionReadyTargetCount == 0)), "Receipt acceptance and production readiness remain disabled even when propagation is modeled."),
		productionReceiptAcceptanceCheck("production-ownership-disabled", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady), "Service start, session bus claim, production bus claim, and production ownership remain disabled."),
		productionReceiptAcceptanceCheck("write-and-launch-disabled", productionAuthorizationPassBlocked(!preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.RequestObjectsCreated && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && productionReceiptAcceptanceTargetsKeepWritesLaunchDisabled(preview.Targets)), "Runtime writes, request creation, adapter invocation, and launch remain disabled."),
		productionReceiptAcceptanceCheck("desktop-support-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.DesktopFilesWritten && !preview.MIMEAppsWritten && !preview.ShellConfigurationWritten && !preview.SettingsPersisted && !preview.NotificationSent && !preview.NotificationDeliveryEnabled && !preview.PortalRequestCreated && !preview.SupportBundleExported && !preview.SupportCaseCreated && !preview.SnapshotRestoreExecuted && !preview.StateCleanupExecuted && !preview.FileContentRead && !preview.FilePathsExposed && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && preview.SideEffectTargetCount == 0 && productionReceiptAcceptanceTargetsKeepSideEffectsAndHostClosed(preview.Targets)), "KDE desktop, Portal, support, restore, cleanup, unsafe-data, network, privilege, internal detail exposure, and host mutation gates remain closed."),
	}
}

func productionReceiptAcceptanceCheck(id string, status string, summary string) ProductionReceiptAcceptancePropagationPreflightCheck {
	return ProductionReceiptAcceptancePropagationPreflightCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptAcceptanceCheckIDs(checks []ProductionReceiptAcceptancePropagationPreflightCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptAcceptanceChecks(checks []ProductionReceiptAcceptancePropagationPreflightCheck) ProductionReceiptAcceptancePropagationPreflightCheckCounts {
	counts := ProductionReceiptAcceptancePropagationPreflightCheckCounts{Total: len(checks)}
	for _, check := range checks {
		switch check.Status {
		case "pass":
			counts.Passed++
		case "pending":
			counts.Pending++
		case "blocked":
			counts.Blocked++
		}
	}
	return counts
}

func productionReceiptAcceptanceConsumptionAuditReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-authorization-consumption-audit-preview",
		"production-gate-consolidated-authorization-consumption-audit",
		"production-authorization-consumption-audit-ready-authorization-disabled",
		"consumers-use-consolidated-boundary",
		"authorization-not-accepted",
		"write-and-launch-disabled",
		"desktop-and-support-side-effects-disabled",
		"host-boundary-closed",
	})
}

func productionReceiptAcceptanceReadyTargetCount(targets []ProductionReceiptAcceptancePropagationTarget) int {
	count := 0
	for _, target := range targets {
		if target.PropagatesFutureAcceptance {
			count++
		}
	}
	return count
}

func productionReceiptAcceptanceMissingTargetCount(targets []ProductionReceiptAcceptancePropagationTarget) int {
	count := 0
	for _, target := range targets {
		if !target.PropagatesFutureAcceptance {
			count++
		}
	}
	return count
}

func productionReceiptAcceptanceTargetIDs(targets []ProductionReceiptAcceptancePropagationTarget) []string {
	ids := make([]string, 0, len(targets))
	for _, target := range targets {
		ids = append(ids, target.ID)
	}
	return ids
}

func productionReceiptAcceptanceAllTargetsReady(targets []ProductionReceiptAcceptancePropagationTarget) bool {
	if len(targets) != 6 {
		return false
	}
	for _, target := range targets {
		if !target.ConsumesAuditBoundary || !target.PropagatesFutureAcceptance || !target.ReceiptBoundaryReady || !target.FutureAcceptanceModeled || target.PropagationStatus != "future-acceptance-modeled-side-effects-disabled" {
			return false
		}
		if target.ReceiptAccepted || target.AuthorizationAccepted || target.ProductionReadiness || target.ProductionOwnershipReady || target.KDEPolicyOwner || !target.ReviewOnly {
			return false
		}
	}
	return true
}

func productionReceiptAcceptanceTargetsKeepWritesLaunchDisabled(targets []ProductionReceiptAcceptancePropagationTarget) bool {
	for _, target := range targets {
		if target.WriteMethodsEnabled || target.RuntimeWritesEnabled || target.BackendLaunchEnabled {
			return false
		}
	}
	return true
}

func productionReceiptAcceptanceTargetsKeepSideEffectsAndHostClosed(targets []ProductionReceiptAcceptancePropagationTarget) bool {
	for _, target := range targets {
		if target.DesktopSideEffectsEnabled || target.SupportSideEffectsEnabled || target.HostRootModified || target.InternalDetailsExposed {
			return false
		}
	}
	return true
}
