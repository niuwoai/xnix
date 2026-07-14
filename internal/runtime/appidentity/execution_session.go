package appidentity

import "errors"

type ExecutionSessionPreview struct {
	SchemaVersion             string                         `json:"schema_version"`
	RequestType               string                         `json:"request_type"`
	SessionType               string                         `json:"session_type"`
	RequestState              string                         `json:"request_state"`
	Source                    string                         `json:"source"`
	Desktop                   string                         `json:"desktop"`
	RuntimeMethod             string                         `json:"runtime_method"`
	ReadMethod                string                         `json:"read_method"`
	ApplicationID             string                         `json:"application_id"`
	ApplicationName           string                         `json:"application_name"`
	Icon                      string                         `json:"icon"`
	DesktopFile               string                         `json:"desktop_file"`
	LauncherCommand           []string                       `json:"launcher_command"`
	Transaction               ExecutionSessionTransaction    `json:"transaction"`
	WindowIdentity            ExecutionSessionWindowIdentity `json:"window_identity"`
	TaskManager               ExecutionSessionTaskManager    `json:"task_manager"`
	KWin                      ExecutionSessionKWin           `json:"kwin"`
	Tray                      ExecutionSessionTray           `json:"tray"`
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
	SessionPlanCreated        bool                           `json:"session_plan_created"`
	SessionCreated            bool                           `json:"session_created"`
	SessionRegistered         bool                           `json:"session_registered"`
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

type ExecutionSessionTransaction struct {
	RequestType             string `json:"request_type"`
	TransactionType         string `json:"transaction_type"`
	RequestState            string `json:"request_state"`
	StepCount               int    `json:"step_count"`
	BlockedStepCount        int    `json:"blocked_step_count"`
	TransactionCommitted    bool   `json:"transaction_committed"`
	RuntimeLaunchApproval   bool   `json:"runtime_launch_approval"`
	LaunchAllowed           bool   `json:"launch_allowed"`
	ExecutionStarted        bool   `json:"execution_started"`
	BackendBindingCommitted bool   `json:"backend_binding_committed"`
	SnapshotBaselineCreated bool   `json:"snapshot_baseline_created"`
	ResourceGrantsCommitted bool   `json:"resource_grants_committed"`
}

type ExecutionSessionWindowIdentity struct {
	SchemaVersion         string `json:"schema_version"`
	WindowKind            string `json:"window_kind"`
	ClassGroup            string `json:"class_group"`
	ResourceName          string `json:"resource_name"`
	LauncherURL           string `json:"launcher_url"`
	TitleHint             string `json:"title_hint"`
	WindowObserved        bool   `json:"window_observed"`
	WindowRegistration    string `json:"window_registration"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

type ExecutionSessionTaskManager struct {
	GroupingKey          string `json:"grouping_key"`
	PinningAllowed       bool   `json:"pinning_allowed"`
	RestoreAllowed       bool   `json:"restore_allowed"`
	SkipTaskbar          bool   `json:"skip_taskbar"`
	ShowInSwitcher       bool   `json:"show_in_switcher"`
	PreferExistingWindow bool   `json:"prefer_existing_window"`
	EntryPlanned         bool   `json:"entry_planned"`
	EntryActive          bool   `json:"entry_active"`
}

type ExecutionSessionKWin struct {
	ScriptRole              string `json:"script_role"`
	ResourceName            string `json:"resource_name"`
	ClassGroup              string `json:"class_group"`
	DesktopFile             string `json:"desktop_file"`
	TaskManagerGroupingKey  string `json:"task_manager_grouping_key"`
	LauncherURL             string `json:"launcher_url"`
	Placement               string `json:"placement"`
	WindowManagerPolicyOnly bool   `json:"window_manager_policy_only"`
	RulePlanned             bool   `json:"rule_planned"`
	RuleApplied             bool   `json:"rule_applied"`
}

type ExecutionSessionTray struct {
	StatusType                   string `json:"status_type"`
	RegisteredApplicationCount   int    `json:"registered_application_count"`
	ActiveApplicationCount       int    `json:"active_application_count"`
	AttentionRequiredCount       int    `json:"attention_required_count"`
	CompatibilityState           string `json:"compatibility_state"`
	TrayBridgeState              string `json:"tray_bridge_state"`
	EntryPlanned                 bool   `json:"entry_planned"`
	LiveBridgeEnabled            bool   `json:"live_bridge_enabled"`
	BridgeConfigurationPersisted bool   `json:"bridge_configuration_persisted"`
	BackendDetailsExposed        bool   `json:"backend_details_exposed"`
}

func (plan Plan) ExecutionSessionPreview(decision string, fileURIs []string) (ExecutionSessionPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return ExecutionSessionPreview{}, err
	}
	if !singleLine(decision) {
		return ExecutionSessionPreview{}, errors.New("execution session preview requires a single-line decision")
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return ExecutionSessionPreview{}, errors.New("execution session preview requires single-line identity fields")
		}
	}

	transaction, err := plan.ExecutionTransactionPreview(decision, fileURIs)
	if err != nil {
		return ExecutionSessionPreview{}, err
	}
	windowIdentity, err := plan.WindowIdentityPreview()
	if err != nil {
		return ExecutionSessionPreview{}, err
	}
	trayStatus, err := plan.TrayStatusPreview()
	if err != nil {
		return ExecutionSessionPreview{}, err
	}

	preview := ExecutionSessionPreview{
		SchemaVersion:   "xnix.runtime.session_identity.v1",
		RequestType:     "execution-session-preview",
		SessionType:     "compatibility-execution-session",
		RequestState:    transaction.RequestState,
		Source:          "execution-transaction-preview",
		Desktop:         "KDE Plasma",
		RuntimeMethod:   "Launch",
		ReadMethod:      "GetExecutionSessionPreview",
		ApplicationID:   plan.ApplicationID,
		ApplicationName: plan.DisplayName,
		Icon:            plan.Icon,
		DesktopFile:     plan.DesktopFile,
		LauncherCommand: plan.LaunchCommand,
		Transaction: ExecutionSessionTransaction{
			RequestType:             transaction.RequestType,
			TransactionType:         transaction.TransactionType,
			RequestState:            transaction.RequestState,
			StepCount:               transaction.StepCount,
			BlockedStepCount:        transaction.BlockedStepCount,
			TransactionCommitted:    transaction.TransactionCommitted,
			RuntimeLaunchApproval:   transaction.RuntimeLaunchApproval,
			LaunchAllowed:           transaction.LaunchAllowed,
			ExecutionStarted:        transaction.ExecutionStarted,
			BackendBindingCommitted: transaction.BackendBindingCommitted,
			SnapshotBaselineCreated: transaction.SnapshotBaselineCreated,
			ResourceGrantsCommitted: transaction.ResourceGrantsCommitted,
		},
		WindowIdentity: ExecutionSessionWindowIdentity{
			SchemaVersion:         windowIdentity.SchemaVersion,
			WindowKind:            windowIdentity.WindowKind,
			ClassGroup:            windowIdentity.ClassGroup,
			ResourceName:          windowIdentity.ResourceName,
			LauncherURL:           windowIdentity.LauncherURL,
			TitleHint:             windowIdentity.TitleHint,
			WindowObserved:        false,
			WindowRegistration:    "planned",
			BackendDetailsExposed: windowIdentity.BackendDetailsExposed,
		},
		TaskManager: ExecutionSessionTaskManager{
			GroupingKey:          windowIdentity.TaskManager.GroupingKey,
			PinningAllowed:       windowIdentity.TaskManager.PinningAllowed,
			RestoreAllowed:       windowIdentity.TaskManager.RestoreAllowed,
			SkipTaskbar:          windowIdentity.TaskManager.SkipTaskbar,
			ShowInSwitcher:       windowIdentity.TaskManager.ShowInSwitcher,
			PreferExistingWindow: windowIdentity.TaskManager.PreferExistingWindow,
			EntryPlanned:         true,
			EntryActive:          false,
		},
		KWin: ExecutionSessionKWin{
			ScriptRole:              windowIdentity.KWin.ScriptRole,
			ResourceName:            windowIdentity.KWin.ResourceName,
			ClassGroup:              windowIdentity.KWin.ClassGroup,
			DesktopFile:             windowIdentity.KWin.DesktopFile,
			TaskManagerGroupingKey:  windowIdentity.KWin.TaskManagerGroupingKey,
			LauncherURL:             windowIdentity.KWin.LauncherURL,
			Placement:               windowIdentity.KWin.Placement,
			WindowManagerPolicyOnly: windowIdentity.KWin.WindowManagerPolicyOnly,
			RulePlanned:             true,
			RuleApplied:             false,
		},
		Tray: ExecutionSessionTray{
			StatusType:                   trayStatus.StatusType,
			RegisteredApplicationCount:   trayStatus.RuntimeActivity.RegisteredApplicationCount,
			ActiveApplicationCount:       trayStatus.RuntimeActivity.ActiveApplicationCount,
			AttentionRequiredCount:       trayStatus.RuntimeActivity.AttentionRequiredCount,
			CompatibilityState:           trayStatus.ApplicationEntry.CompatibilityState,
			TrayBridgeState:              trayStatus.TrayBridge.State,
			EntryPlanned:                 true,
			LiveBridgeEnabled:            trayStatus.LiveBackendBridgeEnabled,
			BridgeConfigurationPersisted: trayStatus.BridgeConfigurationPersisted,
			BackendDetailsExposed:        trayStatus.BackendDetailsExposed,
		},
		FileCount:                 transaction.FileCount,
		FileURIs:                  transaction.FileURIs,
		RuntimeOwned:              true,
		GoRuntimeBacked:           true,
		KDEPolicyOwner:            false,
		CompatibilityCenterCard:   true,
		SafeForAIDiagnostics:      true,
		DesktopEntryLaunchVisible: true,
		LaunchIntentCaptured:      transaction.LaunchIntentCaptured,
		UserDecisionCaptured:      transaction.UserDecisionCaptured,
		UserDecisionAllowsLaunch:  transaction.UserDecisionAllowsLaunch,
		SessionPlanCreated:        true,
		SessionCreated:            false,
		SessionRegistered:         false,
		WindowObserved:            false,
		TaskManagerEntryPlanned:   true,
		TaskManagerEntryActive:    false,
		KWinRulePlanned:           true,
		KWinRuleApplied:           false,
		TrayEntryPlanned:          true,
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
		BlockedActions:            []string{"create execution session from preview", "register live window before Runtime launch", "activate task manager entry from session preview", "apply KWin runtime rule before session exists", "enable live tray bridge before session exists", "start compatibility profile from session preview", "mutate host root during execution session planning", "expose raw backend command to desktop shell"},
		UserFacingSettings:        plan.UserFacingSettings,
		DesktopSafeSummary:        "KDE can prepare normal application window, task manager, and tray identity, but the Runtime does not create a live execution session.",
	}
	if err := validateNoBackendTerms(preview, "execution session preview"); err != nil {
		return ExecutionSessionPreview{}, err
	}
	return preview, nil
}
