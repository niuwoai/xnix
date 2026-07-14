package appidentity

import "errors"

type ExecutionSessionStatusPreview struct {
	SchemaVersion             string                         `json:"schema_version"`
	RequestType               string                         `json:"request_type"`
	StatusType                string                         `json:"status_type"`
	SessionType               string                         `json:"session_type"`
	RequestState              string                         `json:"request_state"`
	SessionState              string                         `json:"session_state"`
	Source                    string                         `json:"source"`
	Desktop                   string                         `json:"desktop"`
	RuntimeMethod             string                         `json:"runtime_method"`
	ReadMethod                string                         `json:"read_method"`
	ApplicationID             string                         `json:"application_id"`
	ApplicationName           string                         `json:"application_name"`
	Icon                      string                         `json:"icon"`
	DesktopFile               string                         `json:"desktop_file"`
	LauncherCommand           []string                       `json:"launcher_command"`
	Session                   ExecutionSessionStatusSession  `json:"session"`
	DesktopSurface            ExecutionSessionDesktopSurface `json:"desktop_surface"`
	UserVisibleState          ExecutionSessionUserState      `json:"user_visible_state"`
	Gates                     []ExecutionSessionStatusGate   `json:"gates"`
	GateCount                 int                            `json:"gate_count"`
	PassedGateCount           int                            `json:"passed_gate_count"`
	RequiredGateCount         int                            `json:"required_gate_count"`
	PendingGateCount          int                            `json:"pending_gate_count"`
	BlockedGateCount          int                            `json:"blocked_gate_count"`
	FileCount                 int                            `json:"file_count"`
	FileURIs                  []string                       `json:"file_uris"`
	RuntimeOwned              bool                           `json:"runtime_owned"`
	GoRuntimeBacked           bool                           `json:"go_runtime_backed"`
	KDEPolicyOwner            bool                           `json:"kde_policy_owner"`
	CompatibilityCenterCard   bool                           `json:"compatibility_center_card"`
	SafeForAIDiagnostics      bool                           `json:"safe_for_ai_diagnostics"`
	DesktopEntryLaunchVisible bool                           `json:"desktop_entry_launch_visible"`
	LaunchIntentCaptured      bool                           `json:"launch_intent_captured"`
	UserDecisionCaptured      bool                           `json:"user_decision_captured"`
	UserDecisionAllowsLaunch  bool                           `json:"user_decision_allows_launch"`
	StatusReadModelCreated    bool                           `json:"status_read_model_created"`
	SessionPlanCreated        bool                           `json:"session_plan_created"`
	SessionCreated            bool                           `json:"session_created"`
	SessionRegistered         bool                           `json:"session_registered"`
	SessionActive             bool                           `json:"session_active"`
	LiveStateObserved         bool                           `json:"live_state_observed"`
	StatusPersisted           bool                           `json:"status_persisted"`
	WindowObserved            bool                           `json:"window_observed"`
	TaskManagerEntryPlanned   bool                           `json:"task_manager_entry_planned"`
	TaskManagerEntryActive    bool                           `json:"task_manager_entry_active"`
	KWinRulePlanned           bool                           `json:"kwin_rule_planned"`
	KWinRuleApplied           bool                           `json:"kwin_rule_applied"`
	TrayEntryPlanned          bool                           `json:"tray_entry_planned"`
	LiveTrayBridgeEnabled     bool                           `json:"live_tray_bridge_enabled"`
	TransactionCommitted      bool                           `json:"transaction_committed"`
	RuntimeLaunchApproval     bool                           `json:"runtime_launch_approval"`
	LaunchAllowed             bool                           `json:"launch_allowed"`
	LaunchEnabled             bool                           `json:"launch_enabled"`
	ExecutionStarted          bool                           `json:"execution_started"`
	BackendProcessStarted     bool                           `json:"backend_process_started"`
	RequestObjectsCreated     bool                           `json:"request_objects_created"`
	HostRootModified          bool                           `json:"host_root_modified"`
	NetworkRequired           bool                           `json:"network_required"`
	BackendDetailsExposed     bool                           `json:"backend_details_exposed"`
	BlockedActions            []string                       `json:"blocked_actions"`
	UserFacingSettings        map[string]string              `json:"user_facing_settings"`
	DesktopSafeSummary        string                         `json:"desktop_safe_summary"`
}

type ExecutionSessionStatusSession struct {
	SessionType              string `json:"session_type"`
	SessionState             string `json:"session_state"`
	RequestState             string `json:"request_state"`
	TransactionState         string `json:"transaction_state"`
	TransactionStepCount     int    `json:"transaction_step_count"`
	BlockedTransactionSteps  int    `json:"blocked_transaction_steps"`
	WindowRegistration       string `json:"window_registration"`
	DesktopSurfaceState      string `json:"desktop_surface_state"`
	CompatibilityCenterState string `json:"compatibility_center_state"`
	TaskManagerState         string `json:"task_manager_state"`
	TrayState                string `json:"tray_state"`
	SessionPlanCreated       bool   `json:"session_plan_created"`
	SessionCreated           bool   `json:"session_created"`
	SessionRegistered        bool   `json:"session_registered"`
	SessionActive            bool   `json:"session_active"`
	LiveStateObserved        bool   `json:"live_state_observed"`
	StatusPersisted          bool   `json:"status_persisted"`
	RuntimeLaunchApproval    bool   `json:"runtime_launch_approval"`
	LaunchAllowed            bool   `json:"launch_allowed"`
	ExecutionStarted         bool   `json:"execution_started"`
	BackendProcessStarted    bool   `json:"backend_process_started"`
	BackendDetailsExposed    bool   `json:"backend_details_exposed"`
}

type ExecutionSessionDesktopSurface struct {
	WindowKind               string `json:"window_kind"`
	ClassGroup               string `json:"class_group"`
	ResourceName             string `json:"resource_name"`
	LauncherURL              string `json:"launcher_url"`
	WindowState              string `json:"window_state"`
	TaskManagerGroupingKey   string `json:"task_manager_grouping_key"`
	TaskManagerState         string `json:"task_manager_state"`
	KWinState                string `json:"kwin_state"`
	TrayState                string `json:"tray_state"`
	CompatibilityCenterState string `json:"compatibility_center_state"`
	WindowObserved           bool   `json:"window_observed"`
	TaskManagerEntryPlanned  bool   `json:"task_manager_entry_planned"`
	TaskManagerEntryActive   bool   `json:"task_manager_entry_active"`
	KWinRulePlanned          bool   `json:"kwin_rule_planned"`
	KWinRuleApplied          bool   `json:"kwin_rule_applied"`
	TrayEntryPlanned         bool   `json:"tray_entry_planned"`
	LiveTrayBridgeEnabled    bool   `json:"live_tray_bridge_enabled"`
}

type ExecutionSessionUserState struct {
	PrimaryLabel              string `json:"primary_label"`
	SecondaryLabel            string `json:"secondary_label"`
	TaskbarBadge              string `json:"taskbar_badge"`
	TrayLabel                 string `json:"tray_label"`
	CompatibilityCenterStatus string `json:"compatibility_center_status"`
	NotificationHint          string `json:"notification_hint"`
	NextUserAction            string `json:"next_user_action"`
	UserFacingMode            string `json:"user_facing_mode"`
	UserFacingAccess          string `json:"user_facing_access"`
	BackendDetailsExposed     bool   `json:"backend_details_exposed"`
}

type ExecutionSessionStatusGate struct {
	ID                  string `json:"id"`
	Status              string `json:"status"`
	Required            bool   `json:"required"`
	BlocksLiveSession   bool   `json:"blocks_live_session"`
	DesktopVisible      bool   `json:"desktop_visible"`
	RuntimeGateRequired bool   `json:"runtime_gate_required"`
	Summary             string `json:"summary"`
}

func (plan Plan) ExecutionSessionStatusPreview(decision string, fileURIs []string) (ExecutionSessionStatusPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return ExecutionSessionStatusPreview{}, err
	}
	if !singleLine(decision) {
		return ExecutionSessionStatusPreview{}, errors.New("execution session status preview requires a single-line decision")
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return ExecutionSessionStatusPreview{}, errors.New("execution session status preview requires single-line identity fields")
		}
	}

	session, err := plan.ExecutionSessionPreview(decision, fileURIs)
	if err != nil {
		return ExecutionSessionStatusPreview{}, err
	}

	gates := executionSessionStatusGates(session)
	passedGateCount, requiredGateCount, pendingGateCount, blockedGateCount := countExecutionSessionStatusGates(gates)
	sessionState := "planned-blocked"
	if session.UserDecisionCaptured && !session.UserDecisionAllowsLaunch {
		sessionState = "review-declined"
	}

	preview := ExecutionSessionStatusPreview{
		SchemaVersion:   "xnix.runtime.session_status.v1",
		RequestType:     "execution-session-status-preview",
		StatusType:      "compatibility-session-status",
		SessionType:     session.SessionType,
		RequestState:    session.RequestState,
		SessionState:    sessionState,
		Source:          "execution-session-preview",
		Desktop:         "KDE Plasma",
		RuntimeMethod:   "Launch",
		ReadMethod:      "GetExecutionSessionStatusPreview",
		ApplicationID:   session.ApplicationID,
		ApplicationName: session.ApplicationName,
		Icon:            session.Icon,
		DesktopFile:     session.DesktopFile,
		LauncherCommand: session.LauncherCommand,
		Session: ExecutionSessionStatusSession{
			SessionType:              session.SessionType,
			SessionState:             sessionState,
			RequestState:             session.RequestState,
			TransactionState:         session.Transaction.RequestState,
			TransactionStepCount:     session.Transaction.StepCount,
			BlockedTransactionSteps:  session.Transaction.BlockedStepCount,
			WindowRegistration:       session.WindowIdentity.WindowRegistration,
			DesktopSurfaceState:      "planned",
			CompatibilityCenterState: "waiting-for-runtime-gates",
			TaskManagerState:         "planned",
			TrayState:                "planned",
			SessionPlanCreated:       session.SessionPlanCreated,
			SessionCreated:           false,
			SessionRegistered:        false,
			SessionActive:            false,
			LiveStateObserved:        false,
			StatusPersisted:          false,
			RuntimeLaunchApproval:    false,
			LaunchAllowed:            false,
			ExecutionStarted:         false,
			BackendProcessStarted:    false,
			BackendDetailsExposed:    false,
		},
		DesktopSurface: ExecutionSessionDesktopSurface{
			WindowKind:               session.WindowIdentity.WindowKind,
			ClassGroup:               session.WindowIdentity.ClassGroup,
			ResourceName:             session.WindowIdentity.ResourceName,
			LauncherURL:              session.WindowIdentity.LauncherURL,
			WindowState:              "not-observed",
			TaskManagerGroupingKey:   session.TaskManager.GroupingKey,
			TaskManagerState:         "planned",
			KWinState:                "planned",
			TrayState:                "planned",
			CompatibilityCenterState: "waiting-for-runtime-gates",
			WindowObserved:           false,
			TaskManagerEntryPlanned:  session.TaskManagerEntryPlanned,
			TaskManagerEntryActive:   false,
			KWinRulePlanned:          session.KWinRulePlanned,
			KWinRuleApplied:          false,
			TrayEntryPlanned:         session.TrayEntryPlanned,
			LiveTrayBridgeEnabled:    false,
		},
		UserVisibleState: ExecutionSessionUserState{
			PrimaryLabel:              session.ApplicationName,
			SecondaryLabel:            "Waiting for Runtime gates",
			TaskbarBadge:              "Planned",
			TrayLabel:                 "Ready for review",
			CompatibilityCenterStatus: "Runtime gates required",
			NotificationHint:          "No live application session has started.",
			NextUserAction:            "Open Compatibility Center",
			UserFacingMode:            modeLabel(session.UserFacingSettings["run_mode"]),
			UserFacingAccess:          "Review required",
			BackendDetailsExposed:     false,
		},
		Gates:                     gates,
		GateCount:                 len(gates),
		PassedGateCount:           passedGateCount,
		RequiredGateCount:         requiredGateCount,
		PendingGateCount:          pendingGateCount,
		BlockedGateCount:          blockedGateCount,
		FileCount:                 session.FileCount,
		FileURIs:                  session.FileURIs,
		RuntimeOwned:              true,
		GoRuntimeBacked:           true,
		KDEPolicyOwner:            false,
		CompatibilityCenterCard:   true,
		SafeForAIDiagnostics:      true,
		DesktopEntryLaunchVisible: true,
		LaunchIntentCaptured:      session.LaunchIntentCaptured,
		UserDecisionCaptured:      session.UserDecisionCaptured,
		UserDecisionAllowsLaunch:  session.UserDecisionAllowsLaunch,
		StatusReadModelCreated:    true,
		SessionPlanCreated:        session.SessionPlanCreated,
		SessionCreated:            false,
		SessionRegistered:         false,
		SessionActive:             false,
		LiveStateObserved:         false,
		StatusPersisted:           false,
		WindowObserved:            false,
		TaskManagerEntryPlanned:   session.TaskManagerEntryPlanned,
		TaskManagerEntryActive:    false,
		KWinRulePlanned:           session.KWinRulePlanned,
		KWinRuleApplied:           false,
		TrayEntryPlanned:          session.TrayEntryPlanned,
		LiveTrayBridgeEnabled:     false,
		TransactionCommitted:      false,
		RuntimeLaunchApproval:     false,
		LaunchAllowed:             false,
		LaunchEnabled:             false,
		ExecutionStarted:          false,
		BackendProcessStarted:     false,
		RequestObjectsCreated:     false,
		HostRootModified:          false,
		NetworkRequired:           false,
		BackendDetailsExposed:     false,
		BlockedActions:            []string{"persist live session status from preview", "mark task manager entry active from status preview", "observe windows before Runtime launch", "apply KWin runtime rule before session exists", "enable live tray bridge before session exists", "start compatibility profile from status preview", "mutate host root during session status planning", "expose raw backend command to desktop shell"},
		UserFacingSettings:        session.UserFacingSettings,
		DesktopSafeSummary:        "KDE can show the application as a planned compatibility session, but no live session state has been observed or persisted.",
	}
	if err := validateNoBackendTerms(preview, "execution session status preview"); err != nil {
		return ExecutionSessionStatusPreview{}, err
	}
	return preview, nil
}

func executionSessionStatusGates(session ExecutionSessionPreview) []ExecutionSessionStatusGate {
	userDecisionStatus := "pass"
	userDecisionRequired := false
	userDecisionBlocks := false
	if session.UserDecisionCaptured && !session.UserDecisionAllowsLaunch {
		userDecisionStatus = "blocked"
		userDecisionRequired = true
		userDecisionBlocks = true
	}

	return []ExecutionSessionStatusGate{
		{
			ID:                  "session-identity",
			Status:              "pass",
			Required:            false,
			BlocksLiveSession:   false,
			DesktopVisible:      true,
			RuntimeGateRequired: false,
			Summary:             "Normal desktop identity is available for the application.",
		},
		{
			ID:                  "user-decision",
			Status:              userDecisionStatus,
			Required:            userDecisionRequired,
			BlocksLiveSession:   userDecisionBlocks,
			DesktopVisible:      true,
			RuntimeGateRequired: userDecisionRequired,
			Summary:             "Compatibility Center decision is captured before live launch.",
		},
		{
			ID:                  "runtime-launch-approval",
			Status:              "blocked",
			Required:            true,
			BlocksLiveSession:   true,
			DesktopVisible:      true,
			RuntimeGateRequired: true,
			Summary:             "Runtime launch approval is disabled until production launch support exists.",
		},
		{
			ID:                  "live-window-observation",
			Status:              "pending",
			Required:            false,
			BlocksLiveSession:   false,
			DesktopVisible:      true,
			RuntimeGateRequired: false,
			Summary:             "Window observation waits for a live session.",
		},
		{
			ID:                  "desktop-surface-activation",
			Status:              "pending",
			Required:            false,
			BlocksLiveSession:   false,
			DesktopVisible:      true,
			RuntimeGateRequired: false,
			Summary:             "Task manager, KWin, and tray activation wait for a live session.",
		},
	}
}

func countExecutionSessionStatusGates(gates []ExecutionSessionStatusGate) (int, int, int, int) {
	passed := 0
	required := 0
	pending := 0
	blocked := 0
	for _, gate := range gates {
		switch gate.Status {
		case "pass":
			passed++
		case "pending":
			pending++
		case "blocked":
			blocked++
		}
		if gate.Required {
			required++
		}
	}
	return passed, required, pending, blocked
}
