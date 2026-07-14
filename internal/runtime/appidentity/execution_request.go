package appidentity

import "errors"

type ExecutionRequestPreview struct {
	SchemaVersion             string               `json:"schema_version"`
	RequestType               string               `json:"request_type"`
	IntentType                string               `json:"intent_type"`
	RequestState              string               `json:"request_state"`
	Source                    string               `json:"source"`
	Desktop                   string               `json:"desktop"`
	RuntimeMethod             string               `json:"runtime_method"`
	ReadMethod                string               `json:"read_method"`
	ApplicationID             string               `json:"application_id"`
	ApplicationName           string               `json:"application_name"`
	Icon                      string               `json:"icon"`
	DesktopFile               string               `json:"desktop_file"`
	LauncherCommand           []string             `json:"launcher_command"`
	LaunchIntent              LaunchIntentSummary  `json:"launch_intent"`
	CompatibilityProfile      ReadinessProfile     `json:"compatibility_profile"`
	GateSummary               ExecutionGateSummary `json:"gate_summary"`
	ExecutionState            string               `json:"execution_state"`
	OverallStatus             string               `json:"overall_status"`
	WriteGateDecision         string               `json:"write_gate_decision"`
	DenialErrorName           string               `json:"denial_error_name"`
	PortalRequired            bool                 `json:"portal_required"`
	SnapshotRequired          bool                 `json:"snapshot_required"`
	FileCount                 int                  `json:"file_count"`
	FileURIs                  []string             `json:"file_uris"`
	RuntimeOwned              bool                 `json:"runtime_owned"`
	GoRuntimeBacked           bool                 `json:"go_runtime_backed"`
	KDEPolicyOwner            bool                 `json:"kde_policy_owner"`
	CompatibilityCenterCard   bool                 `json:"compatibility_center_card"`
	SafeForAIDiagnostics      bool                 `json:"safe_for_ai_diagnostics"`
	DesktopEntryLaunchVisible bool                 `json:"desktop_entry_launch_visible"`
	LaunchIntentCaptured      bool                 `json:"launch_intent_captured"`
	LaunchAllowed             bool                 `json:"launch_allowed"`
	LaunchEnabled             bool                 `json:"launch_enabled"`
	ExecutionRequestCreated   bool                 `json:"execution_request_created"`
	ExecutionRequestPersisted bool                 `json:"execution_request_persisted"`
	ExecutionStarted          bool                 `json:"execution_started"`
	BackendBindingReady       bool                 `json:"backend_binding_ready"`
	RequestObjectCreated      bool                 `json:"request_object_created"`
	PermissionGranted         bool                 `json:"permission_granted"`
	HostRootModified          bool                 `json:"host_root_modified"`
	NetworkRequired           bool                 `json:"network_required"`
	BackendDetailsExposed     bool                 `json:"backend_details_exposed"`
	BlockedActions            []string             `json:"blocked_actions"`
	UserFacingSettings        map[string]string    `json:"user_facing_settings"`
	DesktopSafeSummary        string               `json:"desktop_safe_summary"`
}

type LaunchIntentSummary struct {
	Source                string   `json:"source"`
	IntentType            string   `json:"intent_type"`
	RuntimeMethod         string   `json:"runtime_method"`
	ReadMethod            string   `json:"read_method"`
	ExecutionState        string   `json:"execution_state"`
	OverallStatus         string   `json:"overall_status"`
	WriteGateDecision     string   `json:"write_gate_decision"`
	LaunchAllowed         bool     `json:"launch_allowed"`
	LaunchEnabled         bool     `json:"launch_enabled"`
	PortalRequired        bool     `json:"portal_required"`
	FileCount             int      `json:"file_count"`
	FileURIs              []string `json:"file_uris"`
	RequestObjectCreated  bool     `json:"request_object_created"`
	ExecutionStarted      bool     `json:"execution_started"`
	BackendDetailsExposed bool     `json:"backend_details_exposed"`
}

type ExecutionGateSummary struct {
	GateCount         int `json:"gate_count"`
	RequiredGateCount int `json:"required_gate_count"`
	PendingGateCount  int `json:"pending_gate_count"`
	BlockedGateCount  int `json:"blocked_gate_count"`
}

func (plan Plan) ExecutionRequestPreview(fileURIs []string) (ExecutionRequestPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return ExecutionRequestPreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return ExecutionRequestPreview{}, errors.New("execution request preview requires single-line identity fields")
		}
	}
	launchIntent, err := plan.LaunchIntentPreview(fileURIs)
	if err != nil {
		return ExecutionRequestPreview{}, err
	}
	readiness, err := plan.ExecutionReadinessPreview()
	if err != nil {
		return ExecutionRequestPreview{}, err
	}

	preview := ExecutionRequestPreview{
		SchemaVersion:   "xnix.runtime.request_intake.v1",
		RequestType:     "execution-request-preview",
		IntentType:      launchIntent.IntentType,
		RequestState:    "blocked",
		Source:          "runtime-launch-intent",
		Desktop:         "KDE Plasma",
		RuntimeMethod:   "Launch",
		ReadMethod:      "GetExecutionRequestPreview",
		ApplicationID:   plan.ApplicationID,
		ApplicationName: plan.DisplayName,
		Icon:            plan.Icon,
		DesktopFile:     plan.DesktopFile,
		LauncherCommand: plan.LaunchCommand,
		LaunchIntent: LaunchIntentSummary{
			Source:                launchIntent.Source,
			IntentType:            launchIntent.IntentType,
			RuntimeMethod:         launchIntent.RuntimeMethod,
			ReadMethod:            launchIntent.ReadMethod,
			ExecutionState:        launchIntent.ExecutionState,
			OverallStatus:         launchIntent.OverallStatus,
			WriteGateDecision:     launchIntent.WriteGateDecision,
			LaunchAllowed:         launchIntent.LaunchAllowed,
			LaunchEnabled:         launchIntent.LaunchEnabled,
			PortalRequired:        launchIntent.PortalRequired,
			FileCount:             launchIntent.FileCount,
			FileURIs:              launchIntent.FileURIs,
			RequestObjectCreated:  launchIntent.RequestObjectCreated,
			ExecutionStarted:      launchIntent.ExecutionStarted,
			BackendDetailsExposed: launchIntent.BackendDetailsExposed,
		},
		CompatibilityProfile: readiness.CompatibilityProfile,
		GateSummary: ExecutionGateSummary{
			GateCount:         readiness.GateCount,
			RequiredGateCount: readiness.RequiredGateCount,
			PendingGateCount:  readiness.PendingGateCount,
			BlockedGateCount:  readiness.BlockedGateCount,
		},
		ExecutionState:            readiness.ExecutionState,
		OverallStatus:             readiness.OverallStatus,
		WriteGateDecision:         launchIntent.WriteGateDecision,
		DenialErrorName:           launchIntent.DenialErrorName,
		PortalRequired:            launchIntent.PortalRequired,
		SnapshotRequired:          readiness.SnapshotRequired,
		FileCount:                 launchIntent.FileCount,
		FileURIs:                  launchIntent.FileURIs,
		RuntimeOwned:              true,
		GoRuntimeBacked:           true,
		KDEPolicyOwner:            false,
		CompatibilityCenterCard:   true,
		SafeForAIDiagnostics:      true,
		DesktopEntryLaunchVisible: true,
		LaunchIntentCaptured:      true,
		LaunchAllowed:             false,
		LaunchEnabled:             false,
		ExecutionRequestCreated:   false,
		ExecutionRequestPersisted: false,
		ExecutionStarted:          false,
		BackendBindingReady:       false,
		RequestObjectCreated:      false,
		PermissionGranted:         false,
		HostRootModified:          false,
		NetworkRequired:           false,
		BackendDetailsExposed:     false,
		BlockedActions:            []string{"persist execution request before Runtime gates pass", "start compatibility profile from execution request preview", "grant desktop resources without Portal review", "mutate host root during execution request planning", "expose raw backend command to desktop shell"},
		UserFacingSettings:        plan.UserFacingSettings,
		DesktopSafeSummary:        "KDE launch intent is converted into a blocked Runtime request preview without starting the application.",
	}
	if err := validateNoBackendTerms(preview, "execution request preview"); err != nil {
		return ExecutionRequestPreview{}, err
	}
	return preview, nil
}
