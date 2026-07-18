package appidentity

import "path/filepath"

type KDETestLaunchMaterializationFanOutOwnerRoutePreview struct {
	Version                                 string                                   `json:"version"`
	SchemaVersion                           string                                   `json:"schema_version"`
	RequestType                             string                                   `json:"request_type"`
	RouteType                               string                                   `json:"route_type"`
	Source                                  string                                   `json:"source"`
	RuntimeMethod                           string                                   `json:"runtime_method"`
	ReadMethod                              string                                   `json:"read_method"`
	OpaqueMaterializationReceiptID          string                                   `json:"opaque_materialization_receipt_id"`
	SupportedOpaqueMaterializationReceiptID string                                   `json:"supported_opaque_materialization_receipt_id"`
	MaterializationPlanID                   string                                   `json:"materialization_plan_id"`
	ReceiptRelativePath                     string                                   `json:"receipt_relative_path"`
	ReceiptRecordType                       string                                   `json:"receipt_record_type"`
	ReceiptLookupState                      string                                   `json:"receipt_lookup_state"`
	ReceiptAvailable                        bool                                     `json:"receipt_available"`
	ReceiptConsumed                         bool                                     `json:"receipt_consumed"`
	MissingReceiptSafe                      bool                                     `json:"missing_receipt_safe"`
	OwnerManagedOpaqueReceiptLookupReady    bool                                     `json:"owner_managed_opaque_receipt_lookup_ready"`
	OpaqueMaterializationReceiptIDSupported bool                                     `json:"opaque_materialization_receipt_id_supported"`
	RequiresCallerRegistryPath              bool                                     `json:"requires_caller_registry_path"`
	RequiresCallerApplicationID             bool                                     `json:"requires_caller_application_id"`
	RequiresCallerStateRoot                 bool                                     `json:"requires_caller_state_root"`
	ReadOnlyFanOut                          bool                                     `json:"read_only_fan_out"`
	FanOutResultState                       string                                   `json:"fan_out_result_state"`
	OwnerLocalRouteCandidateReady           bool                                     `json:"owner_local_route_candidate_ready"`
	ProductionDBusExposureReady             bool                                     `json:"production_dbus_exposure_ready"`
	SurfaceCount                            int                                      `json:"surface_count"`
	Surfaces                                []KDETestLaunchMaterializationFanSurface `json:"surfaces"`
	CompatibilityCenter                     KDETestLaunchMaterializationFanSurface   `json:"compatibility_center"`
	TaskManager                             KDETestLaunchMaterializationFanSurface   `json:"task_manager"`
	Tray                                    KDETestLaunchMaterializationFanSurface   `json:"tray"`
	Notification                            KDETestLaunchMaterializationFanSurface   `json:"notification"`
	Checks                                  []KDEFakeExecutionCheck                  `json:"checks"`
	CheckIDs                                []string                                 `json:"check_ids"`
	CheckCount                              int                                      `json:"check_count"`
	PassedCheckCount                        int                                      `json:"passed_check_count"`
	AllChecksPassed                         bool                                     `json:"all_checks_passed"`
	RuntimeOwned                            bool                                     `json:"runtime_owned"`
	GoRuntimeBacked                         bool                                     `json:"go_runtime_backed"`
	KDEPolicyOwner                          bool                                     `json:"kde_policy_owner"`
	MaterializationReceiptConsumed          bool                                     `json:"materialization_receipt_consumed"`
	ExecutionSessionFanOutConsumed          bool                                     `json:"execution_session_fan_out_consumed"`
	StateRootPathExposed                    bool                                     `json:"state_root_path_exposed"`
	StateRootWritesEnabled                  bool                                     `json:"state_root_writes_enabled"`
	RuntimeWritesEnabled                    bool                                     `json:"runtime_writes_enabled"`
	FanOutWritesEnabled                     bool                                     `json:"fan_out_writes_enabled"`
	TestOnly                                bool                                     `json:"test_only"`
	ProductImageReady                       bool                                     `json:"product_image_ready"`
	ProductionTrustSatisfied                bool                                     `json:"production_trust_satisfied"`
	RuntimeWriteGateEnabled                 bool                                     `json:"runtime_write_gate_enabled"`
	LaunchPreflightPassed                   bool                                     `json:"launch_preflight_passed"`
	LaunchAuthorized                        bool                                     `json:"launch_authorized"`
	ExecutionApproved                       bool                                     `json:"execution_approved"`
	ProcessStartAuthorized                  bool                                     `json:"process_start_authorized"`
	CommandMaterialized                     bool                                     `json:"command_materialized"`
	ExecutablePathResolved                  bool                                     `json:"executable_path_resolved"`
	BackendSelectedForLaunch                bool                                     `json:"backend_selected_for_launch"`
	BackendLaunchEnabled                    bool                                     `json:"backend_launch_enabled"`
	BackendProcessStarted                   bool                                     `json:"backend_process_started"`
	TaskManagerEntryActive                  bool                                     `json:"task_manager_entry_active"`
	LiveTrayBridgeEnabled                   bool                                     `json:"live_tray_bridge_enabled"`
	NotificationSent                        bool                                     `json:"notification_sent"`
	NotificationDeliveryEnabled             bool                                     `json:"notification_delivery_enabled"`
	CompatibilityCenterActionsEnabled       bool                                     `json:"compatibility_center_actions_enabled"`
	RequestObjectsCreated                   bool                                     `json:"request_objects_created"`
	ProductionBusOwnership                  bool                                     `json:"production_bus_ownership"`
	NetworkRequired                         bool                                     `json:"network_required"`
	HostRootModified                        bool                                     `json:"host_root_modified"`
	PrivilegedContainerRequired             bool                                     `json:"privileged_container_required"`
	RawCommandExposed                       bool                                     `json:"raw_command_exposed"`
	RawExecutableExposed                    bool                                     `json:"raw_executable_exposed"`
	BackendDetailsExposed                   bool                                     `json:"backend_details_exposed"`
	BlockedActions                          []string                                 `json:"blocked_actions"`
	NextRequirements                        []string                                 `json:"next_requirements"`
	DesktopSafeSummary                      string                                   `json:"desktop_safe_summary"`
}

func NewKDETestLaunchMaterializationFanOutOwnerRoutePreview(root string, opaqueReceiptID string) (KDETestLaunchMaterializationFanOutOwnerRoutePreview, error) {
	lookup, err := ResolveKDETestLaunchMaterializationReceipt(root, opaqueReceiptID)
	if err != nil {
		return KDETestLaunchMaterializationFanOutOwnerRoutePreview{}, err
	}

	center := kdeTestLaunchMaterializationFanOutOwnerRouteSurface("compatibility-center", "KDE Compatibility Center", "GetKDECenterPage", "kde-center-page-preview", lookup, "Compatibility Center can point to the owner-managed materialization receipt slot while missing receipts remain blocked.")
	taskManager := kdeTestLaunchMaterializationFanOutOwnerRouteSurface("task-manager", "Plasma task manager", "GetTaskManagerIdentityPlan", "task-manager-identity-preview", lookup, "Task manager can keep the test launch entry inactive until digest-verified materialization evidence exists.")
	tray := kdeTestLaunchMaterializationFanOutOwnerRouteSurface("tray", "Plasma system tray", "GetTrayStatus", "tray-status-preview", lookup, "Tray can show that materialization fan-out evidence is missing without enabling a live bridge.")
	notification := kdeTestLaunchMaterializationFanOutOwnerRouteSurface("notification", "KDE notification preview", "GetNotificationPlan", "notification-preview", lookup, "Notification preview can describe missing materialization evidence without sending a desktop notification.")
	surfaces := []KDETestLaunchMaterializationFanSurface{center, taskManager, tray, notification}

	preview := KDETestLaunchMaterializationFanOutOwnerRoutePreview{
		Version:                                 lookup.Version,
		SchemaVersion:                           "xnix.runtime.kde_test_launch_materialization_fanout_owner_route.v1",
		RequestType:                             "kde-test-launch-materialization-fanout-owner-route-preview",
		RouteType:                               "owner-local-kde-test-launch-materialization-fanout",
		Source:                                  "runtime-owner-dispatch+kde-test-launch-materialization-receipt-id+opaque-materialization-receipt-id-registry+kde-test-launch-materialization-fanout",
		RuntimeMethod:                           "GetKDETestLaunchMaterializationFanOut",
		ReadMethod:                              "GetKDETestLaunchMaterializationFanOutPreview",
		OpaqueMaterializationReceiptID:          lookup.OpaqueMaterializationReceiptID,
		SupportedOpaqueMaterializationReceiptID: lookup.SupportedOpaqueMaterializationReceiptID,
		MaterializationPlanID:                   lookup.MaterializationPlanID,
		ReceiptRelativePath:                     lookup.ReceiptRelativePath,
		ReceiptRecordType:                       lookup.ReceiptRecordType,
		ReceiptLookupState:                      lookup.ReceiptLookupState,
		ReceiptAvailable:                        lookup.ReceiptAvailable,
		ReceiptConsumed:                         false,
		MissingReceiptSafe:                      lookup.MissingReceiptSafe,
		OwnerManagedOpaqueReceiptLookupReady:    lookup.OwnerManagedOpaqueReceiptLookupReady,
		OpaqueMaterializationReceiptIDSupported: lookup.OpaqueMaterializationReceiptIDSupported,
		RequiresCallerRegistryPath:              false,
		RequiresCallerApplicationID:             false,
		RequiresCallerStateRoot:                 false,
		ReadOnlyFanOut:                          true,
		FanOutResultState:                       "missing-receipt-fail-closed",
		OwnerLocalRouteCandidateReady:           true,
		ProductionDBusExposureReady:             false,
		SurfaceCount:                            len(surfaces),
		Surfaces:                                surfaces,
		CompatibilityCenter:                     center,
		TaskManager:                             taskManager,
		Tray:                                    tray,
		Notification:                            notification,
		RuntimeOwned:                            true,
		GoRuntimeBacked:                         true,
		KDEPolicyOwner:                          false,
		MaterializationReceiptConsumed:          false,
		ExecutionSessionFanOutConsumed:          false,
		StateRootPathExposed:                    false,
		StateRootWritesEnabled:                  false,
		RuntimeWritesEnabled:                    false,
		FanOutWritesEnabled:                     false,
		TestOnly:                                true,
		ProductImageReady:                       false,
		ProductionTrustSatisfied:                false,
		RuntimeWriteGateEnabled:                 false,
		LaunchPreflightPassed:                   false,
		LaunchAuthorized:                        false,
		ExecutionApproved:                       false,
		ProcessStartAuthorized:                  false,
		CommandMaterialized:                     false,
		ExecutablePathResolved:                  false,
		BackendSelectedForLaunch:                false,
		BackendLaunchEnabled:                    false,
		BackendProcessStarted:                   false,
		TaskManagerEntryActive:                  false,
		LiveTrayBridgeEnabled:                   false,
		NotificationSent:                        false,
		NotificationDeliveryEnabled:             false,
		CompatibilityCenterActionsEnabled:       false,
		RequestObjectsCreated:                   false,
		ProductionBusOwnership:                  false,
		NetworkRequired:                         false,
		HostRootModified:                        false,
		PrivilegedContainerRequired:             false,
		RawCommandExposed:                       false,
		RawExecutableExposed:                    false,
		BackendDetailsExposed:                   false,
		BlockedActions: []string{
			"accept caller registry, application id, state-root, or authorization inputs through the owner fan-out route",
			"treat a missing materialization receipt as KDE fan-out evidence",
			"materialize commands, resolve executables, or start compatibility backends from owner-local fan-out",
			"claim production D-Bus ownership or enable Runtime writes from materialization fan-out",
			"send desktop notifications, activate task-manager entries, enable tray bridges, or mutate host root from fan-out",
		},
		NextRequirements: []string{
			"Attach a digest-verified materialization receipt to the owner-managed receipt slot.",
			"Keep owner-local route coverage before considering production D-Bus exposure.",
			"Promote KDE surfaces only after the materialization and session receipts are available and verified.",
			"Continue reporting missing materialization receipts as fail-closed evidence.",
		},
		DesktopSafeSummary: "Runtime owner can route KDE test launch materialization fan-out through an opaque receipt id without caller paths, but missing materialization evidence remains fail-closed and does not enable writes, launch, desktop side effects, production ownership, or host mutation.",
	}
	checks := kdeTestLaunchMaterializationFanOutOwnerRouteChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = kdeTestLaunchMaterializationFanOutOwnerRouteCheckIDs(checks)
	preview.CheckCount = len(checks)
	preview.PassedCheckCount = countKDETestLaunchMaterializationFanOutOwnerRoutePasses(checks)
	preview.AllChecksPassed = preview.PassedCheckCount == preview.CheckCount
	if err := validateNoBackendTerms(preview, "KDE test launch materialization fan-out owner route preview"); err != nil {
		return KDETestLaunchMaterializationFanOutOwnerRoutePreview{}, err
	}
	return preview, nil
}

func kdeTestLaunchMaterializationFanOutOwnerRouteSurface(id string, consumer string, runtimeMethod string, readModel string, lookup KDETestLaunchMaterializationReceiptLookupPreview, summary string) KDETestLaunchMaterializationFanSurface {
	return KDETestLaunchMaterializationFanSurface{
		ID:                        id,
		Consumer:                  consumer,
		RuntimeMethod:             runtimeMethod,
		ReadModel:                 readModel,
		SessionState:              "missing-receipt",
		MaterializationStatus:     lookup.ReceiptLookupState,
		ReceiptRelativePath:       lookup.ReceiptRelativePath,
		ReadOnly:                  true,
		NavigationOnly:            true,
		UserVisible:               true,
		MutatesRuntime:            false,
		StartsProgram:             false,
		DeliversNotification:      false,
		ActivatesTaskManagerEntry: false,
		EnablesTrayBridge:         false,
		EnablesCenterActions:      false,
		ExposesStateRootPath:      false,
		ExposesRawCommand:         false,
		ExposesRawExecutable:      false,
		ExposesBackendDetails:     false,
		Summary:                   summary,
	}
}

func kdeTestLaunchMaterializationFanOutOwnerRouteChecks(preview KDETestLaunchMaterializationFanOutOwnerRoutePreview) []KDEFakeExecutionCheck {
	return []KDEFakeExecutionCheck{
		fakeExecutionCheck("opaque-lookup-consumed", preview.OwnerManagedOpaqueReceiptLookupReady && preview.OpaqueMaterializationReceiptIDSupported && preview.OpaqueMaterializationReceiptID == KDETestLaunchMaterializationOpaqueReceiptID, "Owner-local fan-out consumes the opaque materialization receipt lookup result instead of caller paths."),
		fakeExecutionCheck("caller-paths-hidden", !preview.RequiresCallerRegistryPath && !preview.RequiresCallerApplicationID && !preview.RequiresCallerStateRoot && !preview.StateRootPathExposed && !filepath.IsAbs(preview.ReceiptRelativePath), "Owner-local fan-out does not accept or expose caller registry, application, or state-root paths."),
		fakeExecutionCheck("missing-receipt-fails-closed", !preview.ReceiptAvailable && !preview.ReceiptConsumed && !preview.MaterializationReceiptConsumed && !preview.ExecutionSessionFanOutConsumed && preview.MissingReceiptSafe && preview.FanOutResultState == "missing-receipt-fail-closed", "Missing materialization receipt evidence stays blocked instead of satisfying KDE fan-out surfaces."),
		fakeExecutionCheck("surface-fanout-deferred", kdeTestLaunchMaterializationFanOutOwnerRouteSurfacesDeferred(preview.Surfaces), "All KDE fan-out surfaces are present but do not consume a missing receipt."),
		fakeExecutionCheck("desktop-side-effects-disabled", !preview.TaskManagerEntryActive && !preview.LiveTrayBridgeEnabled && !preview.NotificationSent && !preview.NotificationDeliveryEnabled && !preview.CompatibilityCenterActionsEnabled, "Owner-local fan-out does not activate task-manager entries, enable tray bridges, send notifications, or enable Compatibility Center actions."),
		fakeExecutionCheck("owner-route-ready-production-dbus-blocked", preview.OwnerLocalRouteCandidateReady && !preview.ProductionDBusExposureReady && !preview.ProductionBusOwnership, "The owner-local read route is ready while production D-Bus exposure remains blocked."),
		fakeExecutionCheck("unsafe-gates-closed", !preview.StateRootWritesEnabled && !preview.RuntimeWritesEnabled && !preview.FanOutWritesEnabled && !preview.ProductImageReady && !preview.ProductionTrustSatisfied && !preview.RuntimeWriteGateEnabled && !preview.LaunchPreflightPassed && !preview.LaunchAuthorized && !preview.ExecutionApproved && !preview.ProcessStartAuthorized && !preview.CommandMaterialized && !preview.ExecutablePathResolved && !preview.BackendSelectedForLaunch && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed, "Owner-local fan-out keeps writes, launch, command materialization, network, privilege, unsafe data, and host mutation disabled."),
	}
}

func kdeTestLaunchMaterializationFanOutOwnerRouteSurfacesDeferred(surfaces []KDETestLaunchMaterializationFanSurface) bool {
	if len(surfaces) != 4 {
		return false
	}
	for _, surface := range surfaces {
		if !surface.ReadOnly || !surface.NavigationOnly || !surface.UserVisible || surface.MutatesRuntime || surface.StartsProgram || surface.DeliversNotification || surface.ActivatesTaskManagerEntry || surface.EnablesTrayBridge || surface.EnablesCenterActions || surface.ExposesStateRootPath || surface.ExposesRawCommand || surface.ExposesRawExecutable || surface.ExposesBackendDetails || filepath.IsAbs(surface.ReceiptRelativePath) {
			return false
		}
		if surface.MaterializationStatus != "missing-receipt" || surface.SessionState != "missing-receipt" {
			return false
		}
	}
	return true
}

func kdeTestLaunchMaterializationFanOutOwnerRouteCheckIDs(checks []KDEFakeExecutionCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countKDETestLaunchMaterializationFanOutOwnerRoutePasses(checks []KDEFakeExecutionCheck) int {
	passed := 0
	for _, check := range checks {
		if check.Status == "pass" {
			passed++
		}
	}
	return passed
}
