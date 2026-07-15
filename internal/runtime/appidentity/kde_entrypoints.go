package appidentity

import "errors"

type KDEEntryPointsPreview struct {
	SchemaVersion              string                      `json:"schema_version"`
	RequestType                string                      `json:"request_type"`
	SurfaceType                string                      `json:"surface_type"`
	Source                     string                      `json:"source"`
	Desktop                    string                      `json:"desktop"`
	RuntimeMethod              string                      `json:"runtime_method"`
	ReadMethod                 string                      `json:"read_method"`
	ApplicationID              string                      `json:"application_id"`
	ApplicationName            string                      `json:"application_name"`
	Icon                       string                      `json:"icon"`
	DesktopFile                string                      `json:"desktop_file"`
	LauncherCommand            []string                    `json:"launcher_command"`
	SessionStatus              KDEEntryPointsSessionStatus `json:"session_status"`
	EntryPoints                []KDEEntryPointPreview      `json:"entry_points"`
	EntryPointIDs              []string                    `json:"entry_point_ids"`
	EntryPointCount            int                         `json:"entry_point_count"`
	VisibleEntryPointCount     int                         `json:"visible_entry_point_count"`
	PlannedEntryPointCount     int                         `json:"planned_entry_point_count"`
	ActiveEntryPointCount      int                         `json:"active_entry_point_count"`
	PortalEntryPointCount      int                         `json:"portal_entry_point_count"`
	RuntimeGateEntryPointCount int                         `json:"runtime_gate_entry_point_count"`
	FileCount                  int                         `json:"file_count"`
	FileURIs                   []string                    `json:"file_uris"`
	RuntimeOwned               bool                        `json:"runtime_owned"`
	GoRuntimeBacked            bool                        `json:"go_runtime_backed"`
	KDEPolicyOwner             bool                        `json:"kde_policy_owner"`
	OfficialDesktopOnly        bool                        `json:"official_desktop_only"`
	StableDesktopContract      bool                        `json:"stable_desktop_contract"`
	NormalApplicationSurface   bool                        `json:"normal_application_surface"`
	CompatibilityCenterCard    bool                        `json:"compatibility_center_card"`
	SafeForAIDiagnostics       bool                        `json:"safe_for_ai_diagnostics"`
	DesktopEntryLaunchVisible  bool                        `json:"desktop_entry_launch_visible"`
	UserDecisionCaptured       bool                        `json:"user_decision_captured"`
	UserDecisionAllowsLaunch   bool                        `json:"user_decision_allows_launch"`
	EntryPointPlanCreated      bool                        `json:"entry_point_plan_created"`
	DesktopFilesWritten        bool                        `json:"desktop_files_written"`
	MIMEAppsWritten            bool                        `json:"mimeapps_written"`
	SettingsPersisted          bool                        `json:"settings_persisted"`
	NotificationsSent          bool                        `json:"notifications_sent"`
	TaskManagerEntryActive     bool                        `json:"task_manager_entry_active"`
	KWinRuleApplied            bool                        `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled      bool                        `json:"live_tray_bridge_enabled"`
	RuntimeLaunchApproval      bool                        `json:"runtime_launch_approval"`
	LaunchAllowed              bool                        `json:"launch_allowed"`
	LaunchEnabled              bool                        `json:"launch_enabled"`
	ExecutionStarted           bool                        `json:"execution_started"`
	BackendProcessStarted      bool                        `json:"backend_process_started"`
	RequestObjectsCreated      bool                        `json:"request_objects_created"`
	HostRootModified           bool                        `json:"host_root_modified"`
	NetworkRequired            bool                        `json:"network_required"`
	BackendDetailsExposed      bool                        `json:"backend_details_exposed"`
	BlockedActions             []string                    `json:"blocked_actions"`
	UserFacingSettings         map[string]string           `json:"user_facing_settings"`
	DesktopSafeSummary         string                      `json:"desktop_safe_summary"`
}

type KDEEntryPointsSessionStatus struct {
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

type KDEEntryPointPreview struct {
	ID                    string             `json:"id"`
	Label                 string             `json:"label"`
	KDEComponent          string             `json:"kde_component"`
	RuntimeSource         string             `json:"runtime_source"`
	RuntimeMethod         string             `json:"runtime_method"`
	State                 string             `json:"state"`
	UserAction            string             `json:"user_action"`
	AIAnalysis            *KDEAIAnalysisLink `json:"ai_analysis,omitempty"`
	Visible               bool               `json:"visible"`
	Planned               bool               `json:"planned"`
	Active                bool               `json:"active"`
	RequiresPortal        bool               `json:"requires_portal"`
	RequiresRuntimeGate   bool               `json:"requires_runtime_gate"`
	BlockedByRuntimeGate  bool               `json:"blocked_by_runtime_gate"`
	WritesHost            bool               `json:"writes_host"`
	StartsBackend         bool               `json:"starts_backend"`
	BackendDetailsExposed bool               `json:"backend_details_exposed"`
}

func (plan Plan) KDEEntryPointsPreview(decision string, fileURIs []string) (KDEEntryPointsPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return KDEEntryPointsPreview{}, err
	}
	if !singleLine(decision) {
		return KDEEntryPointsPreview{}, errors.New("KDE entrypoints preview requires a single-line decision")
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return KDEEntryPointsPreview{}, errors.New("KDE entrypoints preview requires single-line identity fields")
		}
	}

	sessionStatus, err := plan.ExecutionSessionStatusPreview(decision, fileURIs)
	if err != nil {
		return KDEEntryPointsPreview{}, err
	}

	entryPoints := kdeEntryPointPreviews()
	ids, visibleCount, plannedCount, activeCount, portalCount, runtimeGateCount := summarizeKDEEntryPoints(entryPoints)
	preview := KDEEntryPointsPreview{
		SchemaVersion:   "xnix.runtime.kde_entrypoints.v1",
		RequestType:     "kde-entrypoints-preview",
		SurfaceType:     "kde-first-release-entrypoints",
		Source:          "execution-session-status-preview",
		Desktop:         "KDE Plasma",
		RuntimeMethod:   "Launch",
		ReadMethod:      "GetKDEEntryPointsPreview",
		ApplicationID:   plan.ApplicationID,
		ApplicationName: plan.DisplayName,
		Icon:            plan.Icon,
		DesktopFile:     plan.DesktopFile,
		LauncherCommand: plan.LaunchCommand,
		SessionStatus: KDEEntryPointsSessionStatus{
			RequestType:           sessionStatus.RequestType,
			StatusType:            sessionStatus.StatusType,
			SessionState:          sessionStatus.SessionState,
			GateCount:             sessionStatus.GateCount,
			BlockedGateCount:      sessionStatus.BlockedGateCount,
			DesktopSurfaceState:   sessionStatus.Session.DesktopSurfaceState,
			UserVisibleState:      sessionStatus.UserVisibleState.CompatibilityCenterStatus,
			RuntimeLaunchApproval: false,
			LaunchAllowed:         false,
			ExecutionStarted:      false,
		},
		EntryPoints:                entryPoints,
		EntryPointIDs:              ids,
		EntryPointCount:            len(entryPoints),
		VisibleEntryPointCount:     visibleCount,
		PlannedEntryPointCount:     plannedCount,
		ActiveEntryPointCount:      activeCount,
		PortalEntryPointCount:      portalCount,
		RuntimeGateEntryPointCount: runtimeGateCount,
		FileCount:                  sessionStatus.FileCount,
		FileURIs:                   sessionStatus.FileURIs,
		RuntimeOwned:               true,
		GoRuntimeBacked:            true,
		KDEPolicyOwner:             false,
		OfficialDesktopOnly:        true,
		StableDesktopContract:      true,
		NormalApplicationSurface:   true,
		CompatibilityCenterCard:    true,
		SafeForAIDiagnostics:       true,
		DesktopEntryLaunchVisible:  true,
		UserDecisionCaptured:       sessionStatus.UserDecisionCaptured,
		UserDecisionAllowsLaunch:   sessionStatus.UserDecisionAllowsLaunch,
		EntryPointPlanCreated:      true,
		DesktopFilesWritten:        false,
		MIMEAppsWritten:            false,
		SettingsPersisted:          false,
		NotificationsSent:          false,
		TaskManagerEntryActive:     false,
		KWinRuleApplied:            false,
		LiveTrayBridgeEnabled:      false,
		RuntimeLaunchApproval:      false,
		LaunchAllowed:              false,
		LaunchEnabled:              false,
		ExecutionStarted:           false,
		BackendProcessStarted:      false,
		RequestObjectsCreated:      false,
		HostRootModified:           false,
		NetworkRequired:            false,
		BackendDetailsExposed:      false,
		BlockedActions:             []string{"write KDE entrypoint files from preview", "persist MIME defaults from preview", "activate task manager entry from entrypoint preview", "apply KWin runtime rule from entrypoint preview", "send notifications from entrypoint preview", "enable live tray bridge from entrypoint preview", "start compatibility profile from entrypoint preview", "mutate host root during KDE entrypoint planning", "expose raw backend command to desktop shell"},
		UserFacingSettings:         plan.UserFacingSettings,
		DesktopSafeSummary:         "KDE can show the seven first-release entry points for the application, but the Runtime does not write desktop files or start a live session.",
	}
	if err := validateNoBackendTerms(preview, "KDE entrypoints preview"); err != nil {
		return KDEEntryPointsPreview{}, err
	}
	return preview, nil
}

func kdeEntryPointPreviews() []KDEEntryPointPreview {
	return []KDEEntryPointPreview{
		kdeEntryPointPreview("launcher", "Start menu", "Plasma application launcher", "desktop-entry-preview", "Launch", "visible-planned", "Launch through Runtime", false, true),
		kdeEntryPointPreview("task-manager", "Task manager", "Plasma task manager", "execution-session-status-preview", "GetExecutionSessionStatusPreview", "planned", "Pin or restore when live", false, true),
		kdeEntryPointPreview("file-manager", "File manager", "Dolphin", "file-open-preview", "Launch", "visible-planned", "Open selected files through Runtime", true, true),
		kdeEntryPointPreview("system-tray", "System tray", "Plasma system tray", "execution-session-status-preview", "GetExecutionSessionStatusPreview", "planned", "Show compatibility status", false, true),
		kdeEntryPointPreview("notifications", "Notifications", "Plasma notification center", "notification-preview", "GetNotificationPlan", "planned", "Show review and repair updates", false, true),
		kdeEntryPointPreview("compatibility-center", "Compatibility Center", "Plasma widget", "compatibility-center-preview", "GetCompatibilityCenterSummary", "visible-planned", "Open Compatibility Center", false, true),
		kdeEntryPointPreview("settings", "Unified settings", "KDE system settings", "settings-preview", "GetCompatibilitySettings", "visible-planned", "Open application settings", false, true),
	}
}

func kdeEntryPointPreview(id string, label string, component string, source string, method string, state string, action string, requiresPortal bool, requiresRuntimeGate bool) KDEEntryPointPreview {
	return KDEEntryPointPreview{
		ID:                    id,
		Label:                 label,
		KDEComponent:          component,
		RuntimeSource:         source,
		RuntimeMethod:         method,
		State:                 state,
		UserAction:            action,
		AIAnalysis:            kdeAIAnalysisLinkForEntryPoint(id),
		Visible:               true,
		Planned:               true,
		Active:                false,
		RequiresPortal:        requiresPortal,
		RequiresRuntimeGate:   requiresRuntimeGate,
		BlockedByRuntimeGate:  requiresRuntimeGate,
		WritesHost:            false,
		StartsBackend:         false,
		BackendDetailsExposed: false,
	}
}

func summarizeKDEEntryPoints(entryPoints []KDEEntryPointPreview) ([]string, int, int, int, int, int) {
	ids := make([]string, 0, len(entryPoints))
	visibleCount := 0
	plannedCount := 0
	activeCount := 0
	portalCount := 0
	runtimeGateCount := 0
	for _, entryPoint := range entryPoints {
		ids = append(ids, entryPoint.ID)
		if entryPoint.Visible {
			visibleCount++
		}
		if entryPoint.Planned {
			plannedCount++
		}
		if entryPoint.Active {
			activeCount++
		}
		if entryPoint.RequiresPortal {
			portalCount++
		}
		if entryPoint.RequiresRuntimeGate {
			runtimeGateCount++
		}
	}
	return ids, visibleCount, plannedCount, activeCount, portalCount, runtimeGateCount
}
