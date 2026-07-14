package appidentity

import "errors"

type KDEActionQueuePreview struct {
	SchemaVersion            string                           `json:"schema_version"`
	RequestType              string                           `json:"request_type"`
	QueueType                string                           `json:"queue_type"`
	Source                   string                           `json:"source"`
	Desktop                  string                           `json:"desktop"`
	RuntimeMethod            string                           `json:"runtime_method"`
	ReadMethod               string                           `json:"read_method"`
	ApplicationID            string                           `json:"application_id"`
	ApplicationName          string                           `json:"application_name"`
	Icon                     string                           `json:"icon"`
	DesktopFile              string                           `json:"desktop_file"`
	LauncherCommand          []string                         `json:"launcher_command"`
	SessionStatus            KDEEntryPointActionSessionStatus `json:"session_status"`
	Actions                  []KDEActionQueueItem             `json:"actions"`
	ActionIDs                []string                         `json:"action_ids"`
	ActionCount              int                              `json:"action_count"`
	PendingActionCount       int                              `json:"pending_action_count"`
	UserReviewRequiredCount  int                              `json:"user_review_required_count"`
	PortalActionCount        int                              `json:"portal_action_count"`
	RuntimeGateActionCount   int                              `json:"runtime_gate_action_count"`
	FileCount                int                              `json:"file_count"`
	FileURIs                 []string                         `json:"file_uris"`
	RuntimeOwned             bool                             `json:"runtime_owned"`
	GoRuntimeBacked          bool                             `json:"go_runtime_backed"`
	KDEPolicyOwner           bool                             `json:"kde_policy_owner"`
	OfficialDesktopOnly      bool                             `json:"official_desktop_only"`
	CompatibilityCenterCard  bool                             `json:"compatibility_center_card"`
	SafeForAIDiagnostics     bool                             `json:"safe_for_ai_diagnostics"`
	UserDecisionCaptured     bool                             `json:"user_decision_captured"`
	UserDecisionAllowsLaunch bool                             `json:"user_decision_allows_launch"`
	ActionQueueCreated       bool                             `json:"action_queue_created"`
	ActionQueuePersisted     bool                             `json:"action_queue_persisted"`
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

type KDEActionQueueItem struct {
	ID                       string `json:"id"`
	SourceType               string `json:"source_type"`
	EntryPointID             string `json:"entry_point_id"`
	EntryPointLabel          string `json:"entry_point_label"`
	KDEComponent             string `json:"kde_component"`
	Intent                   string `json:"intent"`
	Status                   string `json:"status"`
	Priority                 string `json:"priority"`
	Title                    string `json:"title"`
	Summary                  string `json:"summary"`
	UserReviewRequired       bool   `json:"user_review_required"`
	RuntimeGate              string `json:"runtime_gate"`
	NextStep                 string `json:"next_step"`
	RequiresPortal           bool   `json:"requires_portal"`
	RequiresRuntimeGate      bool   `json:"requires_runtime_gate"`
	BlockedByRuntimeGate     bool   `json:"blocked_by_runtime_gate"`
	OpensCompatibilityCenter bool   `json:"opens_compatibility_center"`
	OpensSettings            bool   `json:"opens_settings"`
	ExecutionEnabled         bool   `json:"execution_enabled"`
	RequestObjectCreated     bool   `json:"request_object_created"`
	BackendProcessStarted    bool   `json:"backend_process_started"`
	HostRootModified         bool   `json:"host_root_modified"`
	BackendDetailsExposed    bool   `json:"backend_details_exposed"`
}

func (plan Plan) KDEActionQueuePreview(decision string, fileURIs []string) (KDEActionQueuePreview, error) {
	if !singleLine(decision) {
		return KDEActionQueuePreview{}, errors.New("KDE action queue preview requires a single-line decision")
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return KDEActionQueuePreview{}, errors.New("KDE action queue preview requires single-line identity fields")
		}
	}

	entryPoints, err := plan.KDEEntryPointsPreview(decision, fileURIs)
	if err != nil {
		return KDEActionQueuePreview{}, err
	}

	actions := make([]KDEActionQueueItem, 0, len(entryPoints.EntryPoints))
	for _, entryPoint := range entryPoints.EntryPoints {
		actions = append(actions, kdeActionQueueItem(entryPoint))
	}
	actionIDs, pendingCount, reviewCount, portalCount, runtimeGateCount := summarizeKDEActionQueue(actions)

	preview := KDEActionQueuePreview{
		SchemaVersion:   "xnix.runtime.kde_action_queue.v1",
		RequestType:     "kde-action-queue-preview",
		QueueType:       "compatibility-center-kde-action-queue",
		Source:          "kde-entrypoint-action-preview",
		Desktop:         "KDE Plasma",
		RuntimeMethod:   "Launch",
		ReadMethod:      "GetKDEActionQueuePreview",
		ApplicationID:   entryPoints.ApplicationID,
		ApplicationName: entryPoints.ApplicationName,
		Icon:            entryPoints.Icon,
		DesktopFile:     entryPoints.DesktopFile,
		LauncherCommand: entryPoints.LauncherCommand,
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
		Actions:                  actions,
		ActionIDs:                actionIDs,
		ActionCount:              len(actions),
		PendingActionCount:       pendingCount,
		UserReviewRequiredCount:  reviewCount,
		PortalActionCount:        portalCount,
		RuntimeGateActionCount:   runtimeGateCount,
		FileCount:                entryPoints.FileCount,
		FileURIs:                 entryPoints.FileURIs,
		RuntimeOwned:             true,
		GoRuntimeBacked:          true,
		KDEPolicyOwner:           false,
		OfficialDesktopOnly:      true,
		CompatibilityCenterCard:  true,
		SafeForAIDiagnostics:     true,
		UserDecisionCaptured:     entryPoints.UserDecisionCaptured,
		UserDecisionAllowsLaunch: entryPoints.UserDecisionAllowsLaunch,
		ActionQueueCreated:       true,
		ActionQueuePersisted:     false,
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
		BlockedActions:           []string{"persist KDE action queue from preview", "execute queued KDE actions without Runtime approval", "create Runtime request objects from queued KDE actions", "write KDE desktop files from queued actions", "persist MIME defaults from queued actions", "persist settings from queued actions", "send notifications from queued actions", "activate task manager entries from queued actions", "apply KWin rules from queued actions", "enable live tray bridge from queued actions", "start compatibility profile from queued actions", "mutate host root during KDE action queue preview", "expose raw backend command to desktop shell"},
		UserFacingSettings:       entryPoints.UserFacingSettings,
		DesktopSafeSummary:       "KDE can show a Compatibility Center queue for the seven first-release entrypoint actions, but the Runtime does not persist the queue, create request objects, or start execution.",
	}
	if err := validateNoBackendTerms(preview, "KDE action queue preview"); err != nil {
		return KDEActionQueuePreview{}, err
	}
	return preview, nil
}

func kdeActionQueueItem(entryPoint KDEEntryPointPreview) KDEActionQueueItem {
	action := kdeEntryPointActionSummary(entryPoint)
	return KDEActionQueueItem{
		ID:                       kdeActionQueueItemID(entryPoint.ID),
		SourceType:               "kde-entrypoint-action-preview",
		EntryPointID:             entryPoint.ID,
		EntryPointLabel:          entryPoint.Label,
		KDEComponent:             entryPoint.KDEComponent,
		Intent:                   action.Intent,
		Status:                   kdeActionQueueStatus(action),
		Priority:                 kdeActionQueuePriority(entryPoint.ID),
		Title:                    kdeActionQueueTitle(entryPoint.ID),
		Summary:                  action.SafeResult,
		UserReviewRequired:       kdeActionQueueUserReviewRequired(entryPoint.ID),
		RuntimeGate:              kdeActionQueueRuntimeGate(entryPoint.ID),
		NextStep:                 kdeActionQueueNextStep(entryPoint.ID),
		RequiresPortal:           action.RequiresPortal,
		RequiresRuntimeGate:      action.RequiresRuntimeGate,
		BlockedByRuntimeGate:     action.BlockedByRuntimeGate,
		OpensCompatibilityCenter: action.OpensCompatibilityCenter,
		OpensSettings:            action.OpensSettings,
		ExecutionEnabled:         false,
		RequestObjectCreated:     false,
		BackendProcessStarted:    false,
		HostRootModified:         false,
		BackendDetailsExposed:    false,
	}
}

func kdeActionQueueItemID(entryPointID string) string {
	switch entryPointID {
	case "launcher":
		return "review-launcher-action"
	case "task-manager":
		return "review-task-manager-action"
	case "file-manager":
		return "review-file-manager-action"
	case "system-tray":
		return "review-tray-status-action"
	case "notifications":
		return "review-notification-action"
	case "compatibility-center":
		return "review-compatibility-center-action"
	case "settings":
		return "review-settings-action"
	default:
		return "review-kde-entrypoint-action"
	}
}

func kdeActionQueueStatus(action KDEEntryPointActionSummary) string {
	if action.RequiresPortal {
		return "portal-review-required"
	}
	if action.OpensSettings {
		return "settings-review-required"
	}
	return "runtime-gate-required"
}

func kdeActionQueuePriority(entryPointID string) string {
	switch entryPointID {
	case "launcher", "file-manager":
		return "high"
	case "notifications", "settings":
		return "medium"
	default:
		return "normal"
	}
}

func kdeActionQueueTitle(entryPointID string) string {
	switch entryPointID {
	case "launcher":
		return "Review launcher action"
	case "task-manager":
		return "Review task manager action"
	case "file-manager":
		return "Review file manager action"
	case "system-tray":
		return "Review tray status action"
	case "notifications":
		return "Review notification action"
	case "compatibility-center":
		return "Review Compatibility Center action"
	case "settings":
		return "Review settings action"
	default:
		return "Review KDE entrypoint action"
	}
}

func kdeActionQueueUserReviewRequired(entryPointID string) bool {
	switch entryPointID {
	case "launcher", "file-manager", "notifications", "settings":
		return true
	default:
		return false
	}
}

func kdeActionQueueRuntimeGate(entryPointID string) string {
	switch entryPointID {
	case "launcher":
		return "runtime-launch-review"
	case "task-manager":
		return "runtime-session-status"
	case "file-manager":
		return "portal-file-open-review"
	case "system-tray":
		return "runtime-status-review"
	case "notifications":
		return "runtime-notification-review"
	case "compatibility-center":
		return "compatibility-center-navigation"
	case "settings":
		return "runtime-settings-review"
	default:
		return "runtime-entrypoint-review"
	}
}

func kdeActionQueueNextStep(entryPointID string) string {
	switch entryPointID {
	case "launcher":
		return "Show launch review until Runtime gates pass."
	case "task-manager":
		return "Show planned session status until a live session exists."
	case "file-manager":
		return "Request file access review before creating any launch request."
	case "system-tray":
		return "Show compatibility status without enabling a live tray bridge."
	case "notifications":
		return "Show review details without sending desktop notifications from the preview."
	case "compatibility-center":
		return "Open the application card without executing queued actions."
	case "settings":
		return "Show unified settings without persisting changes."
	default:
		return "Keep the KDE action queued until Runtime gates pass."
	}
}

func summarizeKDEActionQueue(actions []KDEActionQueueItem) ([]string, int, int, int, int) {
	actionIDs := make([]string, 0, len(actions))
	pendingCount := 0
	reviewCount := 0
	portalCount := 0
	runtimeGateCount := 0
	for _, action := range actions {
		actionIDs = append(actionIDs, action.ID)
		if action.Status != "pass" {
			pendingCount++
		}
		if action.UserReviewRequired {
			reviewCount++
		}
		if action.RequiresPortal {
			portalCount++
		}
		if action.RequiresRuntimeGate {
			runtimeGateCount++
		}
	}
	return actionIDs, pendingCount, reviewCount, portalCount, runtimeGateCount
}
