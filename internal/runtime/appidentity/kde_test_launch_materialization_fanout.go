package appidentity

import (
	"errors"
	"path/filepath"
)

type KDETestLaunchMaterializationFanOutPreview struct {
	SchemaVersion                      string                                   `json:"schema_version"`
	RequestType                        string                                   `json:"request_type"`
	Source                             string                                   `json:"source"`
	RuntimeMethod                      string                                   `json:"runtime_method"`
	ReadMethod                         string                                   `json:"read_method"`
	Mode                               string                                   `json:"mode"`
	Application                        KDEFakeExecutionApplication              `json:"application"`
	MaterializationPlanID              string                                   `json:"materialization_plan_id"`
	MaterializationReceiptRelativePath string                                   `json:"materialization_receipt_relative_path"`
	MaterializationReceiptSHA256       string                                   `json:"materialization_receipt_sha256"`
	MaterializationStatus              string                                   `json:"materialization_status"`
	MaterializationScope               string                                   `json:"materialization_scope"`
	ExecutionSessionRequestID          string                                   `json:"execution_session_request_id"`
	ExecutionSessionReceiptPath        string                                   `json:"execution_session_receipt_path"`
	SessionState                       string                                   `json:"session_state"`
	SurfaceCount                       int                                      `json:"surface_count"`
	Surfaces                           []KDETestLaunchMaterializationFanSurface `json:"surfaces"`
	CompatibilityCenter                KDETestLaunchMaterializationFanSurface   `json:"compatibility_center"`
	TaskManager                        KDETestLaunchMaterializationFanSurface   `json:"task_manager"`
	Tray                               KDETestLaunchMaterializationFanSurface   `json:"tray"`
	Notification                       KDETestLaunchMaterializationFanSurface   `json:"notification"`
	Checks                             []KDEFakeExecutionCheck                  `json:"checks"`
	CheckCount                         int                                      `json:"check_count"`
	PassedCheckCount                   int                                      `json:"passed_check_count"`
	AllChecksPassed                    bool                                     `json:"all_checks_passed"`
	RuntimeOwned                       bool                                     `json:"runtime_owned"`
	GoRuntimeBacked                    bool                                     `json:"go_runtime_backed"`
	KDEPolicyOwner                     bool                                     `json:"kde_policy_owner"`
	MaterializationReceiptConsumed     bool                                     `json:"materialization_receipt_consumed"`
	ExecutionSessionFanOutConsumed     bool                                     `json:"execution_session_fan_out_consumed"`
	StateRootPathExposed               bool                                     `json:"state_root_path_exposed"`
	StateRootWritesEnabled             bool                                     `json:"state_root_writes_enabled"`
	StateRootWriteScope                string                                   `json:"state_root_write_scope"`
	FanOutWritesEnabled                bool                                     `json:"fan_out_writes_enabled"`
	TestOnly                           bool                                     `json:"test_only"`
	ProductImageReady                  bool                                     `json:"product_image_ready"`
	ProductionTrustSatisfied           bool                                     `json:"production_trust_satisfied"`
	RuntimeWriteGateEnabled            bool                                     `json:"runtime_write_gate_enabled"`
	LaunchPreflightPassed              bool                                     `json:"launch_preflight_passed"`
	LaunchAuthorized                   bool                                     `json:"launch_authorized"`
	ExecutionApproved                  bool                                     `json:"execution_approved"`
	ProcessStartAuthorized             bool                                     `json:"process_start_authorized"`
	CommandMaterialized                bool                                     `json:"command_materialized"`
	ExecutablePathResolved             bool                                     `json:"executable_path_resolved"`
	BackendSelectedForLaunch           bool                                     `json:"backend_selected_for_launch"`
	BackendLaunchEnabled               bool                                     `json:"backend_launch_enabled"`
	BackendProcessStarted              bool                                     `json:"backend_process_started"`
	TaskManagerEntryActive             bool                                     `json:"task_manager_entry_active"`
	LiveTrayBridgeEnabled              bool                                     `json:"live_tray_bridge_enabled"`
	NotificationSent                   bool                                     `json:"notification_sent"`
	NotificationDeliveryEnabled        bool                                     `json:"notification_delivery_enabled"`
	CompatibilityCenterActionsEnabled  bool                                     `json:"compatibility_center_actions_enabled"`
	RequestObjectsCreated              bool                                     `json:"request_objects_created"`
	RuntimeWritesEnabled               bool                                     `json:"runtime_writes_enabled"`
	ProductionBusOwnership             bool                                     `json:"production_bus_ownership"`
	NetworkRequired                    bool                                     `json:"network_required"`
	HostRootModified                   bool                                     `json:"host_root_modified"`
	PrivilegedContainerRequired        bool                                     `json:"privileged_container_required"`
	RawCommandExposed                  bool                                     `json:"raw_command_exposed"`
	RawExecutableExposed               bool                                     `json:"raw_executable_exposed"`
	BackendDetailsExposed              bool                                     `json:"backend_details_exposed"`
	DesktopSafeSummary                 string                                   `json:"desktop_safe_summary"`
}

type KDETestLaunchMaterializationFanSurface struct {
	ID                        string `json:"id"`
	Consumer                  string `json:"consumer"`
	RuntimeMethod             string `json:"runtime_method"`
	ReadModel                 string `json:"read_model"`
	SessionState              string `json:"session_state"`
	MaterializationStatus     string `json:"materialization_status"`
	ReceiptRelativePath       string `json:"receipt_relative_path"`
	ReadOnly                  bool   `json:"read_only"`
	NavigationOnly            bool   `json:"navigation_only"`
	UserVisible               bool   `json:"user_visible"`
	MutatesRuntime            bool   `json:"mutates_runtime"`
	StartsProgram             bool   `json:"starts_program"`
	DeliversNotification      bool   `json:"delivers_notification"`
	ActivatesTaskManagerEntry bool   `json:"activates_task_manager_entry"`
	EnablesTrayBridge         bool   `json:"enables_tray_bridge"`
	EnablesCenterActions      bool   `json:"enables_center_actions"`
	ExposesStateRootPath      bool   `json:"exposes_state_root_path"`
	ExposesRawCommand         bool   `json:"exposes_raw_command"`
	ExposesRawExecutable      bool   `json:"exposes_raw_executable"`
	ExposesBackendDetails     bool   `json:"exposes_backend_details"`
	Summary                   string `json:"summary"`
}

func NewKDETestLaunchMaterializationFanOutPreview(recipeRecord Recipe, provenance Provenance, options KDERestrictedLaunchAuthorizationOptions) (KDETestLaunchMaterializationFanOutPreview, error) {
	materializationRecord, err := NewKDETestLaunchMaterializationRecord(recipeRecord, provenance, options)
	if err != nil {
		return KDETestLaunchMaterializationFanOutPreview{}, err
	}
	plan, err := NewPlanWithProvenance(recipeRecord, provenance)
	if err != nil {
		return KDETestLaunchMaterializationFanOutPreview{}, err
	}
	fanOut, err := plan.ExecutionSessionFanOutEvidence(options.StateRoot, materializationRecord.Execution.RequestID)
	if err != nil {
		return KDETestLaunchMaterializationFanOutPreview{}, err
	}
	if err := validateKDETestLaunchMaterializationFanOutInputs(materializationRecord, fanOut); err != nil {
		return KDETestLaunchMaterializationFanOutPreview{}, err
	}

	center := kdeTestLaunchMaterializationFanSurface("compatibility-center", "KDE Compatibility Center", "GetKDECenterPage", "kde-center-page-preview", fanOut.CompatibilityCenter.State, materializationRecord.Materialization, "Compatibility Center can show the test-only review plan and blocked gates without approving execution.")
	taskManager := kdeTestLaunchMaterializationFanSurface("task-manager", "Plasma task manager", "GetTaskManagerIdentityPlan", "task-manager-identity-preview", fanOut.TaskManager.State, materializationRecord.Materialization, "Task manager can show blocked review-plan state without activating entries or observing windows.")
	tray := kdeTestLaunchMaterializationFanSurface("tray", "Plasma system tray", "GetTrayStatus", "tray-status-preview", fanOut.Tray.State, materializationRecord.Materialization, "Tray can show Runtime-gated compatibility status without enabling a live bridge.")
	notification := kdeTestLaunchMaterializationFanSurface("notification", "KDE notification preview", "GetNotificationPlan", "notification-preview", "review-plan-available", materializationRecord.Materialization, "Notification preview can describe that approval is required without sending a desktop notification.")
	surfaces := []KDETestLaunchMaterializationFanSurface{center, taskManager, tray, notification}
	checks := kdeTestLaunchMaterializationFanOutChecks(materializationRecord, fanOut, surfaces)
	passed := 0
	for _, check := range checks {
		if check.Status == "pass" {
			passed++
		}
	}
	if passed != len(checks) {
		return KDETestLaunchMaterializationFanOutPreview{}, errors.New("test launch materialization fan-out checks did not all pass")
	}

	preview := KDETestLaunchMaterializationFanOutPreview{
		SchemaVersion:                      "xnix.runtime.kde_test_launch_materialization_fanout.v1",
		RequestType:                        "kde-test-launch-materialization-fanout-preview",
		Source:                             "kde-test-launch-materialization-record+execution-session-fanout-evidence",
		RuntimeMethod:                      "GetKDETestLaunchMaterializationFanOut",
		ReadMethod:                         "GetKDETestLaunchMaterializationFanOutPreview",
		Mode:                               options.Mode,
		Application:                        materializationRecord.Application,
		MaterializationPlanID:              materializationRecord.Materialization.PlanID,
		MaterializationReceiptRelativePath: materializationRecord.Materialization.ReceiptRelativePath,
		MaterializationReceiptSHA256:       materializationRecord.Materialization.ReceiptSHA256,
		MaterializationStatus:              materializationRecord.Materialization.Status,
		MaterializationScope:               materializationRecord.Materialization.MaterializationScope,
		ExecutionSessionRequestID:          fanOut.RequestID,
		ExecutionSessionReceiptPath:        fanOut.ReceiptRelativePath,
		SessionState:                       fanOut.SessionState,
		SurfaceCount:                       len(surfaces),
		Surfaces:                           surfaces,
		CompatibilityCenter:                center,
		TaskManager:                        taskManager,
		Tray:                               tray,
		Notification:                       notification,
		Checks:                             checks,
		CheckCount:                         len(checks),
		PassedCheckCount:                   passed,
		AllChecksPassed:                    true,
		RuntimeOwned:                       true,
		GoRuntimeBacked:                    true,
		KDEPolicyOwner:                     false,
		MaterializationReceiptConsumed:     true,
		ExecutionSessionFanOutConsumed:     true,
		StateRootPathExposed:               false,
		StateRootWritesEnabled:             true,
		StateRootWriteScope:                "explicit-test-root-materialization-only",
		FanOutWritesEnabled:                false,
		TestOnly:                           true,
		ProductImageReady:                  false,
		ProductionTrustSatisfied:           false,
		RuntimeWriteGateEnabled:            false,
		LaunchPreflightPassed:              false,
		LaunchAuthorized:                   false,
		ExecutionApproved:                  false,
		ProcessStartAuthorized:             false,
		CommandMaterialized:                false,
		ExecutablePathResolved:             false,
		BackendSelectedForLaunch:           false,
		BackendLaunchEnabled:               false,
		BackendProcessStarted:              false,
		TaskManagerEntryActive:             false,
		LiveTrayBridgeEnabled:              false,
		NotificationSent:                   false,
		NotificationDeliveryEnabled:        false,
		CompatibilityCenterActionsEnabled:  false,
		RequestObjectsCreated:              false,
		RuntimeWritesEnabled:               false,
		ProductionBusOwnership:             false,
		NetworkRequired:                    false,
		HostRootModified:                   false,
		PrivilegedContainerRequired:        false,
		RawCommandExposed:                  false,
		RawExecutableExposed:               false,
		BackendDetailsExposed:              false,
		DesktopSafeSummary:                 "KDE can consume a test-only materialized review plan across Compatibility Center, task manager, tray, and notification previews while launch, execution, write methods, desktop notification delivery, live tray bridging, and host mutation remain disabled.",
	}
	if err := validateNoBackendTerms(preview, "KDE test launch materialization fan-out preview"); err != nil {
		return KDETestLaunchMaterializationFanOutPreview{}, err
	}
	return preview, nil
}

func validateKDETestLaunchMaterializationFanOutInputs(record KDETestLaunchMaterializationRecord, fanOut ExecutionSessionFanOutEvidence) error {
	if !record.AllChecksPassed || !record.PlanMaterialized || !record.TestOnly || !record.Materialization.ReceiptReadBack {
		return errors.New("materialization fan-out requires a verified test-only materialization record")
	}
	if record.Execution.RequestID == "" || record.Execution.RequestID != fanOut.RequestID {
		return errors.New("materialization fan-out requires matching execution session evidence")
	}
	if !fanOut.SafeForKDE || fanOut.SurfaceCount != 4 {
		return errors.New("materialization fan-out requires safe execution session fan-out evidence")
	}
	if filepath.IsAbs(record.Materialization.ReceiptRelativePath) || filepath.IsAbs(fanOut.ReceiptRelativePath) {
		return errors.New("materialization fan-out requires relative receipt paths")
	}
	if record.Materialization.Status != "blocked-plan-materialized" || record.Materialization.MaterializationScope != "test-only-review-plan" {
		return errors.New("materialization fan-out requires a blocked test-only review plan")
	}
	if record.CommandMaterialized || record.ExecutablePathResolved || record.BackendSelectedForLaunch || record.BackendLaunchEnabled || record.BackendProcessStarted {
		return errors.New("materialization fan-out received unsafe launch materialization")
	}
	return nil
}

func kdeTestLaunchMaterializationFanSurface(id string, consumer string, runtimeMethod string, readModel string, sessionState string, materialization KDETestLaunchMaterializationEvidence, summary string) KDETestLaunchMaterializationFanSurface {
	return KDETestLaunchMaterializationFanSurface{
		ID:                        id,
		Consumer:                  consumer,
		RuntimeMethod:             runtimeMethod,
		ReadModel:                 readModel,
		SessionState:              sessionState,
		MaterializationStatus:     materialization.Status,
		ReceiptRelativePath:       materialization.ReceiptRelativePath,
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

func kdeTestLaunchMaterializationFanOutChecks(record KDETestLaunchMaterializationRecord, fanOut ExecutionSessionFanOutEvidence, surfaces []KDETestLaunchMaterializationFanSurface) []KDEFakeExecutionCheck {
	return []KDEFakeExecutionCheck{
		fakeExecutionCheck("materialization-consumed", record.Materialization.ReceiptReadBack && record.Materialization.PlanMaterialized && record.Materialization.Status == "blocked-plan-materialized" && !filepath.IsAbs(record.Materialization.ReceiptRelativePath), "The fan-out consumed a digest-verified test-only materialization receipt."),
		fakeExecutionCheck("session-fanout-consumed", fanOut.SafeForKDE && fanOut.SurfaceCount == 4 && fanOut.RequestID == record.Execution.RequestID && fanOut.SessionState == "blocked", "The fan-out consumed safe blocked execution session evidence."),
		fakeExecutionCheck("surface-coverage", len(surfaces) == 4 && kdeTestLaunchFanOutHasSurface(surfaces, "compatibility-center") && kdeTestLaunchFanOutHasSurface(surfaces, "task-manager") && kdeTestLaunchFanOutHasSurface(surfaces, "tray") && kdeTestLaunchFanOutHasSurface(surfaces, "notification"), "Compatibility Center, task manager, tray, and notification previews are covered."),
		fakeExecutionCheck("surfaces-read-only", kdeTestLaunchFanOutSurfacesReadOnly(surfaces), "All KDE surfaces are read-only, navigation-only projections."),
		fakeExecutionCheck("desktop-side-effects-disabled", !fanOut.TaskManagerEntryActive && !fanOut.LiveTrayBridgeEnabled && !record.BackendProcessStarted, "Task-manager activation, live tray bridge, and process start remain disabled."),
		fakeExecutionCheck("notification-not-delivered", kdeTestLaunchFanOutNoDeliveredNotification(surfaces), "Notification is a preview only and is not delivered."),
		fakeExecutionCheck("launch-boundary-closed", !record.LaunchAuthorized && !record.ExecutionApproved && !record.ProcessStartAuthorized && !record.CommandMaterialized && !record.ExecutablePathResolved && !record.BackendSelectedForLaunch && !record.BackendLaunchEnabled, "Launch, execution approval, command materialization, executable resolution, and launch selection stay disabled."),
		fakeExecutionCheck("unsafe-data-hidden", !record.StateRootPathExposed && !record.RawCommandExposed && !record.RawExecutableExposed && !record.BackendDetailsExposed && !fanOut.StateRootPathExposed && !fanOut.BackendDetailsExposed, "State-root paths, raw commands, raw executable details, and implementation details stay hidden."),
		fakeExecutionCheck("host-boundary-closed", !record.NetworkRequired && !record.PrivilegedContainerRequired && !record.HostRootModified && !fanOut.NetworkRequired && !fanOut.HostRootModified, "Network, privileged-container, and host-root mutation gates stay closed."),
	}
}

func kdeTestLaunchFanOutHasSurface(surfaces []KDETestLaunchMaterializationFanSurface, id string) bool {
	for _, surface := range surfaces {
		if surface.ID == id {
			return true
		}
	}
	return false
}

func kdeTestLaunchFanOutSurfacesReadOnly(surfaces []KDETestLaunchMaterializationFanSurface) bool {
	for _, surface := range surfaces {
		if !surface.ReadOnly || !surface.NavigationOnly || !surface.UserVisible || surface.MutatesRuntime || surface.StartsProgram || surface.ExposesStateRootPath || surface.ExposesRawCommand || surface.ExposesRawExecutable || surface.ExposesBackendDetails {
			return false
		}
	}
	return len(surfaces) > 0
}

func kdeTestLaunchFanOutNoDeliveredNotification(surfaces []KDETestLaunchMaterializationFanSurface) bool {
	for _, surface := range surfaces {
		if surface.ID == "notification" {
			return !surface.DeliversNotification
		}
	}
	return false
}
