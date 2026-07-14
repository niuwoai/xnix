package appidentity

import (
	"errors"
	"fmt"
)

type KDEEntryPointActionPreview struct {
	SchemaVersion            string                           `json:"schema_version"`
	RequestType              string                           `json:"request_type"`
	ActionType               string                           `json:"action_type"`
	Source                   string                           `json:"source"`
	Desktop                  string                           `json:"desktop"`
	EntryPointID             string                           `json:"entry_point_id"`
	EntryPointLabel          string                           `json:"entry_point_label"`
	KDEComponent             string                           `json:"kde_component"`
	RuntimeSource            string                           `json:"runtime_source"`
	RuntimeMethod            string                           `json:"runtime_method"`
	ReadMethod               string                           `json:"read_method"`
	ApplicationID            string                           `json:"application_id"`
	ApplicationName          string                           `json:"application_name"`
	Icon                     string                           `json:"icon"`
	DesktopFile              string                           `json:"desktop_file"`
	LauncherCommand          []string                         `json:"launcher_command"`
	Action                   KDEEntryPointActionSummary       `json:"action"`
	EntryPoint               KDEEntryPointActionEntry         `json:"entry_point"`
	SessionStatus            KDEEntryPointActionSessionStatus `json:"session_status"`
	GateSummary              KDEEntryPointActionGateSummary   `json:"gate_summary"`
	FileCount                int                              `json:"file_count"`
	FileURIs                 []string                         `json:"file_uris"`
	RuntimeOwned             bool                             `json:"runtime_owned"`
	GoRuntimeBacked          bool                             `json:"go_runtime_backed"`
	KDEPolicyOwner           bool                             `json:"kde_policy_owner"`
	OfficialDesktopOnly      bool                             `json:"official_desktop_only"`
	StableDesktopContract    bool                             `json:"stable_desktop_contract"`
	NormalApplicationSurface bool                             `json:"normal_application_surface"`
	CompatibilityCenterCard  bool                             `json:"compatibility_center_card"`
	SafeForAIDiagnostics     bool                             `json:"safe_for_ai_diagnostics"`
	UserDecisionCaptured     bool                             `json:"user_decision_captured"`
	UserDecisionAllowsLaunch bool                             `json:"user_decision_allows_launch"`
	EntryPointActionCaptured bool                             `json:"entry_point_action_captured"`
	EntryPointPlanCreated    bool                             `json:"entry_point_plan_created"`
	DesktopFilesWritten      bool                             `json:"desktop_files_written"`
	MIMEAppsWritten          bool                             `json:"mimeapps_written"`
	SettingsPersisted        bool                             `json:"settings_persisted"`
	NotificationsSent        bool                             `json:"notifications_sent"`
	TaskManagerEntryActive   bool                             `json:"task_manager_entry_active"`
	KWinRuleApplied          bool                             `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled    bool                             `json:"live_tray_bridge_enabled"`
	RuntimeLaunchApproval    bool                             `json:"runtime_launch_approval"`
	LaunchAllowed            bool                             `json:"launch_allowed"`
	LaunchEnabled            bool                             `json:"launch_enabled"`
	ExecutionStarted         bool                             `json:"execution_started"`
	BackendProcessStarted    bool                             `json:"backend_process_started"`
	RequestObjectsCreated    bool                             `json:"request_objects_created"`
	HostRootModified         bool                             `json:"host_root_modified"`
	NetworkRequired          bool                             `json:"network_required"`
	BackendDetailsExposed    bool                             `json:"backend_details_exposed"`
	BlockedActions           []string                         `json:"blocked_actions"`
	UserFacingSettings       map[string]string                `json:"user_facing_settings"`
	DesktopSafeSummary       string                           `json:"desktop_safe_summary"`
}

type KDEEntryPointActionSummary struct {
	Intent                   string `json:"intent"`
	UserAction               string `json:"user_action"`
	SafeResult               string `json:"safe_result"`
	RequiresPortal           bool   `json:"requires_portal"`
	RequiresRuntimeGate      bool   `json:"requires_runtime_gate"`
	BlockedByRuntimeGate     bool   `json:"blocked_by_runtime_gate"`
	OpensCompatibilityCenter bool   `json:"opens_compatibility_center"`
	OpensSettings            bool   `json:"opens_settings"`
	CreatesRequestObject     bool   `json:"creates_request_object"`
	StartsBackend            bool   `json:"starts_backend"`
	MutatesHost              bool   `json:"mutates_host"`
	BackendDetailsExposed    bool   `json:"backend_details_exposed"`
}

type KDEEntryPointActionEntry struct {
	ID                    string `json:"id"`
	Label                 string `json:"label"`
	KDEComponent          string `json:"kde_component"`
	State                 string `json:"state"`
	Visible               bool   `json:"visible"`
	Planned               bool   `json:"planned"`
	Active                bool   `json:"active"`
	RequiresPortal        bool   `json:"requires_portal"`
	RequiresRuntimeGate   bool   `json:"requires_runtime_gate"`
	BlockedByRuntimeGate  bool   `json:"blocked_by_runtime_gate"`
	WritesHost            bool   `json:"writes_host"`
	StartsBackend         bool   `json:"starts_backend"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

type KDEEntryPointActionSessionStatus struct {
	RequestType           string `json:"request_type"`
	StatusType            string `json:"status_type"`
	SessionState          string `json:"session_state"`
	GateCount             int    `json:"gate_count"`
	BlockedGateCount      int    `json:"blocked_gate_count"`
	DesktopSurfaceState   string `json:"desktop_surface_state"`
	UserVisibleState      string `json:"user_visible_state"`
	RuntimeLaunchApproval bool   `json:"runtime_launch_approval"`
	LaunchAllowed         bool   `json:"launch_allowed"`
	ExecutionStarted      bool   `json:"execution_started"`
}

type KDEEntryPointActionGateSummary struct {
	GateCount             int  `json:"gate_count"`
	BlockedGateCount      int  `json:"blocked_gate_count"`
	RequiresRuntimeGate   bool `json:"requires_runtime_gate"`
	BlockedByRuntimeGate  bool `json:"blocked_by_runtime_gate"`
	RuntimeLaunchApproval bool `json:"runtime_launch_approval"`
	LaunchAllowed         bool `json:"launch_allowed"`
	ExecutionStarted      bool `json:"execution_started"`
	RequestObjectCreated  bool `json:"request_object_created"`
	BackendProcessStarted bool `json:"backend_process_started"`
	BackendDetailsExposed bool `json:"backend_details_exposed"`
}

func (plan Plan) KDEEntryPointActionPreview(entryPointID string, decision string, fileURIs []string) (KDEEntryPointActionPreview, error) {
	if !singleLine(entryPointID) {
		return KDEEntryPointActionPreview{}, errors.New("KDE entrypoint action preview requires a single-line entrypoint id")
	}
	if !singleLine(decision) {
		return KDEEntryPointActionPreview{}, errors.New("KDE entrypoint action preview requires a single-line decision")
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return KDEEntryPointActionPreview{}, errors.New("KDE entrypoint action preview requires single-line identity fields")
		}
	}

	entryPoints, err := plan.KDEEntryPointsPreview(decision, fileURIs)
	if err != nil {
		return KDEEntryPointActionPreview{}, err
	}
	entryPoint, ok := findKDEEntryPointPreview(entryPoints.EntryPoints, entryPointID)
	if !ok {
		return KDEEntryPointActionPreview{}, fmt.Errorf("unsupported KDE entrypoint: %s", entryPointID)
	}
	action := kdeEntryPointActionSummary(entryPoint)

	preview := KDEEntryPointActionPreview{
		SchemaVersion:   "xnix.runtime.kde_entrypoint_action.v1",
		RequestType:     "kde-entrypoint-action-preview",
		ActionType:      "kde-entrypoint-action",
		Source:          "kde-entrypoints-preview",
		Desktop:         "KDE Plasma",
		EntryPointID:    entryPoint.ID,
		EntryPointLabel: entryPoint.Label,
		KDEComponent:    entryPoint.KDEComponent,
		RuntimeSource:   entryPoint.RuntimeSource,
		RuntimeMethod:   entryPoint.RuntimeMethod,
		ReadMethod:      "GetKDEEntryPointActionPreview",
		ApplicationID:   entryPoints.ApplicationID,
		ApplicationName: entryPoints.ApplicationName,
		Icon:            entryPoints.Icon,
		DesktopFile:     entryPoints.DesktopFile,
		LauncherCommand: entryPoints.LauncherCommand,
		Action:          action,
		EntryPoint: KDEEntryPointActionEntry{
			ID:                    entryPoint.ID,
			Label:                 entryPoint.Label,
			KDEComponent:          entryPoint.KDEComponent,
			State:                 entryPoint.State,
			Visible:               entryPoint.Visible,
			Planned:               entryPoint.Planned,
			Active:                entryPoint.Active,
			RequiresPortal:        entryPoint.RequiresPortal,
			RequiresRuntimeGate:   entryPoint.RequiresRuntimeGate,
			BlockedByRuntimeGate:  entryPoint.BlockedByRuntimeGate,
			WritesHost:            entryPoint.WritesHost,
			StartsBackend:         entryPoint.StartsBackend,
			BackendDetailsExposed: entryPoint.BackendDetailsExposed,
		},
		SessionStatus: KDEEntryPointActionSessionStatus{
			RequestType:           entryPoints.SessionStatus.RequestType,
			StatusType:            entryPoints.SessionStatus.StatusType,
			SessionState:          entryPoints.SessionStatus.SessionState,
			GateCount:             entryPoints.SessionStatus.GateCount,
			BlockedGateCount:      entryPoints.SessionStatus.BlockedGateCount,
			DesktopSurfaceState:   entryPoints.SessionStatus.DesktopSurfaceState,
			UserVisibleState:      entryPoints.SessionStatus.UserVisibleState,
			RuntimeLaunchApproval: false,
			LaunchAllowed:         false,
			ExecutionStarted:      false,
		},
		GateSummary: KDEEntryPointActionGateSummary{
			GateCount:             entryPoints.SessionStatus.GateCount,
			BlockedGateCount:      entryPoints.SessionStatus.BlockedGateCount,
			RequiresRuntimeGate:   action.RequiresRuntimeGate,
			BlockedByRuntimeGate:  action.BlockedByRuntimeGate,
			RuntimeLaunchApproval: false,
			LaunchAllowed:         false,
			ExecutionStarted:      false,
			RequestObjectCreated:  false,
			BackendProcessStarted: false,
			BackendDetailsExposed: false,
		},
		FileCount:                entryPoints.FileCount,
		FileURIs:                 entryPoints.FileURIs,
		RuntimeOwned:             true,
		GoRuntimeBacked:          true,
		KDEPolicyOwner:           false,
		OfficialDesktopOnly:      true,
		StableDesktopContract:    true,
		NormalApplicationSurface: true,
		CompatibilityCenterCard:  true,
		SafeForAIDiagnostics:     true,
		UserDecisionCaptured:     entryPoints.UserDecisionCaptured,
		UserDecisionAllowsLaunch: entryPoints.UserDecisionAllowsLaunch,
		EntryPointActionCaptured: true,
		EntryPointPlanCreated:    true,
		DesktopFilesWritten:      false,
		MIMEAppsWritten:          false,
		SettingsPersisted:        false,
		NotificationsSent:        false,
		TaskManagerEntryActive:   false,
		KWinRuleApplied:          false,
		LiveTrayBridgeEnabled:    false,
		RuntimeLaunchApproval:    false,
		LaunchAllowed:            false,
		LaunchEnabled:            false,
		ExecutionStarted:         false,
		BackendProcessStarted:    false,
		RequestObjectsCreated:    false,
		HostRootModified:         false,
		NetworkRequired:          false,
		BackendDetailsExposed:    false,
		BlockedActions:           []string{"commit KDE entrypoint action from preview", "write KDE desktop files from entrypoint action preview", "persist MIME defaults from entrypoint action preview", "persist settings from entrypoint action preview", "send notifications from entrypoint action preview", "activate task manager entry from entrypoint action preview", "apply KWin rule from entrypoint action preview", "enable live tray bridge from entrypoint action preview", "grant Runtime launch approval from entrypoint action preview", "start compatibility profile from entrypoint action preview", "create Runtime request object from entrypoint action preview", "mutate host root during entrypoint action preview", "expose raw backend command to desktop shell"},
		UserFacingSettings:       entryPoints.UserFacingSettings,
		DesktopSafeSummary:       fmt.Sprintf("KDE can route the %s entry point to a Runtime-owned action preview, but the Runtime does not write desktop files, create request objects, or start execution.", entryPoint.Label),
	}
	if err := validateNoBackendTerms(preview, "KDE entrypoint action preview"); err != nil {
		return KDEEntryPointActionPreview{}, err
	}
	return preview, nil
}

func findKDEEntryPointPreview(entryPoints []KDEEntryPointPreview, id string) (KDEEntryPointPreview, bool) {
	for _, entryPoint := range entryPoints {
		if entryPoint.ID == id {
			return entryPoint, true
		}
	}
	return KDEEntryPointPreview{}, false
}

func kdeEntryPointActionSummary(entryPoint KDEEntryPointPreview) KDEEntryPointActionSummary {
	action := KDEEntryPointActionSummary{
		Intent:                   "open-application",
		UserAction:               entryPoint.UserAction,
		SafeResult:               "show Compatibility Center launch review",
		RequiresPortal:           entryPoint.RequiresPortal,
		RequiresRuntimeGate:      entryPoint.RequiresRuntimeGate,
		BlockedByRuntimeGate:     entryPoint.BlockedByRuntimeGate,
		OpensCompatibilityCenter: true,
		OpensSettings:            false,
		CreatesRequestObject:     false,
		StartsBackend:            false,
		MutatesHost:              false,
		BackendDetailsExposed:    false,
	}

	switch entryPoint.ID {
	case "launcher":
		action.Intent = "open-application"
		action.SafeResult = "show Compatibility Center launch review"
	case "task-manager":
		action.Intent = "restore-application"
		action.SafeResult = "show planned session status"
	case "file-manager":
		action.Intent = "open-files"
		action.SafeResult = "show file access review"
	case "system-tray":
		action.Intent = "open-status"
		action.SafeResult = "show compatibility status"
	case "notifications":
		action.Intent = "open-review-card"
		action.SafeResult = "show review details"
	case "compatibility-center":
		action.Intent = "open-compatibility-center"
		action.SafeResult = "show application compatibility card"
	case "settings":
		action.Intent = "open-settings"
		action.SafeResult = "show unified settings"
		action.OpensSettings = true
	}

	return action
}
