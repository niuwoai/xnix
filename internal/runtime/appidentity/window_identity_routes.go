package appidentity

import (
	"errors"
	"strings"
)

type TaskManagerIdentityPlanPreview struct {
	SchemaVersion            string           `json:"schema_version"`
	RequestType              string           `json:"request_type"`
	PlanType                 string           `json:"plan_type"`
	Source                   string           `json:"source"`
	RuntimeMethod            string           `json:"runtime_method"`
	ReadMethod               string           `json:"read_method"`
	Desktop                  string           `json:"desktop"`
	ApplicationID            string           `json:"application_id"`
	DisplayName              string           `json:"display_name"`
	DesktopFile              string           `json:"desktop_file"`
	LauncherURL              string           `json:"launcher_url"`
	WindowKind               string           `json:"window_kind"`
	GroupingKey              string           `json:"grouping_key"`
	TaskManager              TaskManagerHints `json:"task_manager"`
	Restore                  RestoreHints     `json:"restore"`
	RuntimeOwned             bool             `json:"runtime_owned"`
	GoRuntimeBacked          bool             `json:"go_runtime_backed"`
	KDEPolicyOwner           bool             `json:"kde_policy_owner"`
	PinningAllowed           bool             `json:"pinning_allowed"`
	RestoreAllowed           bool             `json:"restore_allowed"`
	PreferExistingWindow     bool             `json:"prefer_existing_window"`
	SkipTaskbar              bool             `json:"skip_taskbar"`
	ShowInSwitcher           bool             `json:"show_in_switcher"`
	ActivationReceiptRoot    bool             `json:"activation_receipt_root"`
	ActivationReceiptBacked  bool             `json:"activation_receipt_backed"`
	ActivationReceiptPath    string           `json:"activation_receipt_path,omitempty"`
	ExecutionSessionRoot     bool             `json:"execution_session_root"`
	ExecutionSessionBacked   bool             `json:"execution_session_backed"`
	ExecutionSessionPath     string           `json:"execution_session_path,omitempty"`
	ExecutionSessionState    string           `json:"execution_session_state,omitempty"`
	TaskManagerEntryActive   bool             `json:"task_manager_entry_active"`
	WindowObservationStarted bool             `json:"window_observation_started"`
	LaunchEnabled            bool             `json:"launch_enabled"`
	ExecutionStarted         bool             `json:"execution_started"`
	HostRootModified         bool             `json:"host_root_modified"`
	BackendDetailsExposed    bool             `json:"backend_details_exposed"`
	DesktopSafeSummary       string           `json:"desktop_safe_summary"`
}

type TaskManagerIdentityOptions struct {
	ActivationRoot            string
	ExecutionSessionRoot      string
	ExecutionSessionRequestID string
}

type KWinWindowRulePlanPreview struct {
	SchemaVersion            string        `json:"schema_version"`
	RequestType              string        `json:"request_type"`
	PlanType                 string        `json:"plan_type"`
	Source                   string        `json:"source"`
	RuntimeMethod            string        `json:"runtime_method"`
	ReadMethod               string        `json:"read_method"`
	Desktop                  string        `json:"desktop"`
	ApplicationID            string        `json:"application_id"`
	DisplayName              string        `json:"display_name"`
	DesktopFile              string        `json:"desktop_file"`
	LauncherURL              string        `json:"launcher_url"`
	WindowKind               string        `json:"window_kind"`
	Match                    KWinRuleMatch `json:"match"`
	Set                      KWinRuleSet   `json:"set"`
	Restore                  RestoreHints  `json:"restore"`
	RuntimeOwned             bool          `json:"runtime_owned"`
	GoRuntimeBacked          bool          `json:"go_runtime_backed"`
	KDEPolicyOwner           bool          `json:"kde_policy_owner"`
	WindowManagerPolicyOnly  bool          `json:"window_manager_policy_only"`
	RuntimeOwnsBackendPolicy bool          `json:"runtime_owns_backend_policy"`
	ActivationReceiptRoot    bool          `json:"activation_receipt_root"`
	ActivationReceiptBacked  bool          `json:"activation_receipt_backed"`
	ActivationReceiptPath    string        `json:"activation_receipt_path,omitempty"`
	ExecutionSessionRoot     bool          `json:"execution_session_root"`
	ExecutionSessionBacked   bool          `json:"execution_session_backed"`
	ExecutionSessionPath     string        `json:"execution_session_path,omitempty"`
	ExecutionSessionState    string        `json:"execution_session_state,omitempty"`
	KWinRuleApplied          bool          `json:"kwin_rule_applied"`
	TaskManagerEntryActive   bool          `json:"task_manager_entry_active"`
	WindowObservationStarted bool          `json:"window_observation_started"`
	LaunchEnabled            bool          `json:"launch_enabled"`
	ExecutionStarted         bool          `json:"execution_started"`
	HostRootModified         bool          `json:"host_root_modified"`
	BackendDetailsExposed    bool          `json:"backend_details_exposed"`
	DesktopSafeSummary       string        `json:"desktop_safe_summary"`
}

type KWinWindowRuleOptions struct {
	ActivationRoot            string
	ExecutionSessionRoot      string
	ExecutionSessionRequestID string
}

type KWinRuleMatch struct {
	ResourceName string `json:"resource_name"`
	ClassGroup   string `json:"class_group"`
	TitleHint    string `json:"title_hint"`
}

type KWinRuleSet struct {
	DesktopFile            string `json:"desktop_file"`
	ApplicationID          string `json:"application_id"`
	TaskManagerGroupingKey string `json:"task_manager_grouping_key"`
	LauncherURL            string `json:"launcher_url"`
	SkipTaskbar            bool   `json:"skip_taskbar"`
	ShowInSwitcher         bool   `json:"show_in_switcher"`
	Placement              string `json:"placement"`
}

func (plan Plan) TaskManagerIdentityPlanPreview() (TaskManagerIdentityPlanPreview, error) {
	return plan.TaskManagerIdentityPlanPreviewWithOptions(TaskManagerIdentityOptions{})
}

func (plan Plan) TaskManagerIdentityPlanPreviewWithOptions(options TaskManagerIdentityOptions) (TaskManagerIdentityPlanPreview, error) {
	windowIdentity, err := plan.WindowIdentityPreview()
	if err != nil {
		return TaskManagerIdentityPlanPreview{}, err
	}
	if !singleLine(windowIdentity.ApplicationID) ||
		!singleLine(windowIdentity.DisplayName) ||
		!singleLine(windowIdentity.DesktopFile) ||
		!singleLine(windowIdentity.LauncherURL) {
		return TaskManagerIdentityPlanPreview{}, errors.New("task manager identity preview requires single-line identity fields")
	}

	source := "window-identity-preview"
	activationReceiptRoot := strings.TrimSpace(options.ActivationRoot) != ""
	activationReceiptBacked := false
	activationReceiptPath := ""
	if activationReceiptRoot {
		evidence, err := plan.DesktopActivationReceiptEvidence(options.ActivationRoot)
		if err != nil {
			return TaskManagerIdentityPlanPreview{}, err
		}
		activationReceiptBacked = evidence.SafeForKDE
		activationReceiptPath = evidence.ReceiptRelativePath
		source = "window-identity-preview+desktop-activation-receipt"
	}
	executionSessionRoot := strings.TrimSpace(options.ExecutionSessionRoot) != ""
	executionSessionBacked := false
	executionSessionPath := ""
	executionSessionState := ""
	if executionSessionRoot {
		evidence, err := plan.ExecutionSessionFanOutEvidence(options.ExecutionSessionRoot, options.ExecutionSessionRequestID)
		if err != nil {
			return TaskManagerIdentityPlanPreview{}, err
		}
		executionSessionBacked = evidence.SafeForKDE
		executionSessionPath = evidence.ReceiptRelativePath
		executionSessionState = evidence.TaskManager.State
		source = source + "+execution-session-record"
	}

	preview := TaskManagerIdentityPlanPreview{
		SchemaVersion:            "xnix.runtime.task_manager_identity.v1",
		RequestType:              "task-manager-identity-preview",
		PlanType:                 "task-manager-identity-plan",
		Source:                   source,
		RuntimeMethod:            "GetTaskManagerIdentityPlan",
		ReadMethod:               "GetTaskManagerIdentityPlanPreview",
		Desktop:                  windowIdentity.Desktop,
		ApplicationID:            windowIdentity.ApplicationID,
		DisplayName:              windowIdentity.DisplayName,
		DesktopFile:              windowIdentity.DesktopFile,
		LauncherURL:              windowIdentity.LauncherURL,
		WindowKind:               windowIdentity.WindowKind,
		GroupingKey:              windowIdentity.TaskManager.GroupingKey,
		TaskManager:              windowIdentity.TaskManager,
		Restore:                  windowIdentity.Restore,
		RuntimeOwned:             true,
		GoRuntimeBacked:          true,
		KDEPolicyOwner:           false,
		PinningAllowed:           windowIdentity.TaskManager.PinningAllowed,
		RestoreAllowed:           windowIdentity.TaskManager.RestoreAllowed,
		PreferExistingWindow:     windowIdentity.TaskManager.PreferExistingWindow,
		SkipTaskbar:              windowIdentity.TaskManager.SkipTaskbar,
		ShowInSwitcher:           windowIdentity.TaskManager.ShowInSwitcher,
		ActivationReceiptRoot:    activationReceiptRoot,
		ActivationReceiptBacked:  activationReceiptBacked,
		ActivationReceiptPath:    activationReceiptPath,
		ExecutionSessionRoot:     executionSessionRoot,
		ExecutionSessionBacked:   executionSessionBacked,
		ExecutionSessionPath:     executionSessionPath,
		ExecutionSessionState:    executionSessionState,
		TaskManagerEntryActive:   false,
		WindowObservationStarted: false,
		LaunchEnabled:            false,
		ExecutionStarted:         false,
		HostRootModified:         false,
		BackendDetailsExposed:    false,
		DesktopSafeSummary:       "Task manager identity lets KDE group, pin, switch, and restore the application as a normal desktop window while execution remains blocked.",
	}
	if err := validateNoBackendTerms(preview, "task manager identity preview"); err != nil {
		return TaskManagerIdentityPlanPreview{}, err
	}
	return preview, nil
}

func (plan Plan) KWinWindowRulePlanPreview() (KWinWindowRulePlanPreview, error) {
	return plan.KWinWindowRulePlanPreviewWithOptions(KWinWindowRuleOptions{})
}

func (plan Plan) KWinWindowRulePlanPreviewWithOptions(options KWinWindowRuleOptions) (KWinWindowRulePlanPreview, error) {
	windowIdentity, err := plan.WindowIdentityPreview()
	if err != nil {
		return KWinWindowRulePlanPreview{}, err
	}
	if !singleLine(windowIdentity.ApplicationID) ||
		!singleLine(windowIdentity.DisplayName) ||
		!singleLine(windowIdentity.DesktopFile) ||
		!singleLine(windowIdentity.LauncherURL) {
		return KWinWindowRulePlanPreview{}, errors.New("KWin window rule preview requires single-line identity fields")
	}

	source := "window-identity-preview"
	activationReceiptRoot := strings.TrimSpace(options.ActivationRoot) != ""
	activationReceiptBacked := false
	activationReceiptPath := ""
	if activationReceiptRoot {
		evidence, err := plan.DesktopActivationReceiptEvidence(options.ActivationRoot)
		if err != nil {
			return KWinWindowRulePlanPreview{}, err
		}
		activationReceiptBacked = evidence.SafeForKDE
		activationReceiptPath = evidence.ReceiptRelativePath
		source = "window-identity-preview+desktop-activation-receipt"
	}
	executionSessionRoot := strings.TrimSpace(options.ExecutionSessionRoot) != ""
	executionSessionBacked := false
	executionSessionPath := ""
	executionSessionState := ""
	if executionSessionRoot {
		evidence, err := plan.ExecutionSessionFanOutEvidence(options.ExecutionSessionRoot, options.ExecutionSessionRequestID)
		if err != nil {
			return KWinWindowRulePlanPreview{}, err
		}
		executionSessionBacked = evidence.SafeForKDE
		executionSessionPath = evidence.ReceiptRelativePath
		executionSessionState = evidence.KWin.State
		source = source + "+execution-session-record"
	}

	preview := KWinWindowRulePlanPreview{
		SchemaVersion: "xnix.runtime.kwin_window_rule.v1",
		RequestType:   "kwin-window-rule-preview",
		PlanType:      "kwin-window-rule-plan",
		Source:        source,
		RuntimeMethod: "GetKWinWindowRulePlan",
		ReadMethod:    "GetKWinWindowRulePlanPreview",
		Desktop:       windowIdentity.Desktop,
		ApplicationID: windowIdentity.ApplicationID,
		DisplayName:   windowIdentity.DisplayName,
		DesktopFile:   windowIdentity.DesktopFile,
		LauncherURL:   windowIdentity.LauncherURL,
		WindowKind:    windowIdentity.WindowKind,
		Match: KWinRuleMatch{
			ResourceName: windowIdentity.KWin.ResourceName,
			ClassGroup:   windowIdentity.KWin.ClassGroup,
			TitleHint:    windowIdentity.TitleHint,
		},
		Set: KWinRuleSet{
			DesktopFile:            windowIdentity.KWin.DesktopFile,
			ApplicationID:          windowIdentity.ApplicationID,
			TaskManagerGroupingKey: windowIdentity.KWin.TaskManagerGroupingKey,
			LauncherURL:            windowIdentity.KWin.LauncherURL,
			SkipTaskbar:            false,
			ShowInSwitcher:         true,
			Placement:              windowIdentity.KWin.Placement,
		},
		Restore:                  windowIdentity.Restore,
		RuntimeOwned:             true,
		GoRuntimeBacked:          true,
		KDEPolicyOwner:           false,
		WindowManagerPolicyOnly:  windowIdentity.KWin.WindowManagerPolicyOnly,
		RuntimeOwnsBackendPolicy: windowIdentity.KWin.RuntimeOwnsBackendPolicy,
		ActivationReceiptRoot:    activationReceiptRoot,
		ActivationReceiptBacked:  activationReceiptBacked,
		ActivationReceiptPath:    activationReceiptPath,
		ExecutionSessionRoot:     executionSessionRoot,
		ExecutionSessionBacked:   executionSessionBacked,
		ExecutionSessionPath:     executionSessionPath,
		ExecutionSessionState:    executionSessionState,
		KWinRuleApplied:          false,
		TaskManagerEntryActive:   false,
		WindowObservationStarted: false,
		LaunchEnabled:            false,
		ExecutionStarted:         false,
		HostRootModified:         false,
		BackendDetailsExposed:    false,
		DesktopSafeSummary:       "KWin receives Runtime-owned identity and layout hints while rule application and execution stay blocked.",
	}
	if err := validateNoBackendTerms(preview, "KWin window rule preview"); err != nil {
		return KWinWindowRulePlanPreview{}, err
	}
	return preview, nil
}
