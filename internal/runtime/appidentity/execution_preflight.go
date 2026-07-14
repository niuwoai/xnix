package appidentity

import "errors"

type ExecutionPreflightPreview struct {
	SchemaVersion             string                   `json:"schema_version"`
	RequestType               string                   `json:"request_type"`
	PreflightType             string                   `json:"preflight_type"`
	RequestState              string                   `json:"request_state"`
	Source                    string                   `json:"source"`
	Desktop                   string                   `json:"desktop"`
	RuntimeMethod             string                   `json:"runtime_method"`
	ReadMethod                string                   `json:"read_method"`
	ApplicationID             string                   `json:"application_id"`
	ApplicationName           string                   `json:"application_name"`
	Icon                      string                   `json:"icon"`
	DesktopFile               string                   `json:"desktop_file"`
	LauncherCommand           []string                 `json:"launcher_command"`
	ExecutionDecision         ExecutionDecisionSummary `json:"execution_decision"`
	PreflightChecks           []ExecutionPreflightGate `json:"preflight_checks"`
	CheckCount                int                      `json:"check_count"`
	PassedCheckCount          int                      `json:"passed_check_count"`
	RequiredCheckCount        int                      `json:"required_check_count"`
	PendingCheckCount         int                      `json:"pending_check_count"`
	BlockedCheckCount         int                      `json:"blocked_check_count"`
	CompatibilityProfile      ReadinessProfile         `json:"compatibility_profile"`
	GateSummary               ExecutionGateSummary     `json:"gate_summary"`
	ExecutionState            string                   `json:"execution_state"`
	OverallStatus             string                   `json:"overall_status"`
	WriteGateDecision         string                   `json:"write_gate_decision"`
	DenialErrorName           string                   `json:"denial_error_name"`
	PortalRequired            bool                     `json:"portal_required"`
	SnapshotRequired          bool                     `json:"snapshot_required"`
	FileCount                 int                      `json:"file_count"`
	FileURIs                  []string                 `json:"file_uris"`
	RuntimeOwned              bool                     `json:"runtime_owned"`
	GoRuntimeBacked           bool                     `json:"go_runtime_backed"`
	KDEPolicyOwner            bool                     `json:"kde_policy_owner"`
	CompatibilityCenterCard   bool                     `json:"compatibility_center_card"`
	SafeForAIDiagnostics      bool                     `json:"safe_for_ai_diagnostics"`
	DesktopEntryLaunchVisible bool                     `json:"desktop_entry_launch_visible"`
	LaunchIntentCaptured      bool                     `json:"launch_intent_captured"`
	UserDecisionCaptured      bool                     `json:"user_decision_captured"`
	UserDecisionAllowsLaunch  bool                     `json:"user_decision_allows_launch"`
	PreflightComplete         bool                     `json:"preflight_complete"`
	PreflightPassed           bool                     `json:"preflight_passed"`
	PortalPreflightReady      bool                     `json:"portal_preflight_ready"`
	SnapshotPreflightReady    bool                     `json:"snapshot_preflight_ready"`
	BackendPreflightReady     bool                     `json:"backend_preflight_ready"`
	WriteGateOpen             bool                     `json:"write_gate_open"`
	RuntimeLaunchApproval     bool                     `json:"runtime_launch_approval"`
	LaunchAllowed             bool                     `json:"launch_allowed"`
	LaunchEnabled             bool                     `json:"launch_enabled"`
	ExecutionRequestCreated   bool                     `json:"execution_request_created"`
	ExecutionRequestPersisted bool                     `json:"execution_request_persisted"`
	ActionQueuePersisted      bool                     `json:"action_queue_persisted"`
	ReviewReceiptRecorded     bool                     `json:"review_receipt_recorded"`
	ExecutionStarted          bool                     `json:"execution_started"`
	BackendBindingReady       bool                     `json:"backend_binding_ready"`
	RequestObjectCreated      bool                     `json:"request_object_created"`
	PermissionGranted         bool                     `json:"permission_granted"`
	HostRootModified          bool                     `json:"host_root_modified"`
	NetworkRequired           bool                     `json:"network_required"`
	BackendDetailsExposed     bool                     `json:"backend_details_exposed"`
	BlockedActions            []string                 `json:"blocked_actions"`
	UserFacingSettings        map[string]string        `json:"user_facing_settings"`
	DesktopSafeSummary        string                   `json:"desktop_safe_summary"`
}

type ExecutionDecisionSummary struct {
	SchemaVersion          string `json:"schema_version"`
	RequestType            string `json:"request_type"`
	DecisionType           string `json:"decision_type"`
	Decision               string `json:"decision"`
	DecisionAccepted       bool   `json:"decision_accepted"`
	UserIntentCaptured     bool   `json:"user_intent_captured"`
	DecisionRecorded       bool   `json:"decision_recorded"`
	RuntimeApprovalGranted bool   `json:"runtime_approval_granted"`
	ExecutionAllowed       bool   `json:"execution_allowed"`
	ReviewReceiptCreated   bool   `json:"review_receipt_created"`
	QueueStateChanged      bool   `json:"queue_state_changed"`
}

type ExecutionPreflightGate struct {
	ID       string `json:"id"`
	Status   string `json:"status"`
	Required bool   `json:"required"`
	Summary  string `json:"summary"`
}

func (plan Plan) ExecutionPreflightPreview(decision string, fileURIs []string) (ExecutionPreflightPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return ExecutionPreflightPreview{}, err
	}
	if !singleLine(decision) {
		return ExecutionPreflightPreview{}, errors.New("execution preflight preview requires a single-line decision")
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return ExecutionPreflightPreview{}, errors.New("execution preflight preview requires single-line identity fields")
		}
	}
	decisionPreview, err := plan.ExecutionDecisionPreview(decision, fileURIs)
	if err != nil {
		return ExecutionPreflightPreview{}, err
	}

	preflightChecks := executionPreflightChecks(decisionPreview)
	passedCheckCount, requiredCheckCount, pendingCheckCount, blockedCheckCount := countExecutionPreflightChecks(preflightChecks)
	userDecisionAllowsLaunch := decision == "approved" || decision == "reviewed"

	preview := ExecutionPreflightPreview{
		SchemaVersion:   "xnix.runtime.launch_preflight.v1",
		RequestType:     "execution-preflight-preview",
		PreflightType:   "compatibility-launch-preflight",
		RequestState:    decisionPreview.RequestState,
		Source:          "execution-decision-preview",
		Desktop:         "KDE Plasma",
		RuntimeMethod:   "Launch",
		ReadMethod:      "GetExecutionPreflightPreview",
		ApplicationID:   plan.ApplicationID,
		ApplicationName: plan.DisplayName,
		Icon:            plan.Icon,
		DesktopFile:     plan.DesktopFile,
		LauncherCommand: plan.LaunchCommand,
		ExecutionDecision: ExecutionDecisionSummary{
			SchemaVersion:          decisionPreview.SchemaVersion,
			RequestType:            decisionPreview.RequestType,
			DecisionType:           decisionPreview.DecisionType,
			Decision:               decisionPreview.Decision.Decision,
			DecisionAccepted:       decisionPreview.Decision.DecisionAccepted,
			UserIntentCaptured:     decisionPreview.Decision.UserIntentCaptured,
			DecisionRecorded:       decisionPreview.Decision.DecisionRecorded,
			RuntimeApprovalGranted: decisionPreview.Decision.RuntimeApprovalGranted,
			ExecutionAllowed:       decisionPreview.Decision.ExecutionAllowed,
			ReviewReceiptCreated:   decisionPreview.Decision.ReviewReceiptCreated,
			QueueStateChanged:      decisionPreview.Decision.QueueStateChanged,
		},
		PreflightChecks:           preflightChecks,
		CheckCount:                len(preflightChecks),
		PassedCheckCount:          passedCheckCount,
		RequiredCheckCount:        requiredCheckCount,
		PendingCheckCount:         pendingCheckCount,
		BlockedCheckCount:         blockedCheckCount,
		CompatibilityProfile:      decisionPreview.CompatibilityProfile,
		GateSummary:               decisionPreview.GateSummary,
		ExecutionState:            decisionPreview.ExecutionState,
		OverallStatus:             decisionPreview.OverallStatus,
		WriteGateDecision:         decisionPreview.WriteGateDecision,
		DenialErrorName:           decisionPreview.DenialErrorName,
		PortalRequired:            decisionPreview.PortalRequired,
		SnapshotRequired:          decisionPreview.SnapshotRequired,
		FileCount:                 decisionPreview.FileCount,
		FileURIs:                  decisionPreview.FileURIs,
		RuntimeOwned:              true,
		GoRuntimeBacked:           true,
		KDEPolicyOwner:            false,
		CompatibilityCenterCard:   true,
		SafeForAIDiagnostics:      true,
		DesktopEntryLaunchVisible: true,
		LaunchIntentCaptured:      decisionPreview.LaunchIntentCaptured,
		UserDecisionCaptured:      true,
		UserDecisionAllowsLaunch:  userDecisionAllowsLaunch,
		PreflightComplete:         false,
		PreflightPassed:           false,
		PortalPreflightReady:      false,
		SnapshotPreflightReady:    false,
		BackendPreflightReady:     false,
		WriteGateOpen:             false,
		RuntimeLaunchApproval:     false,
		LaunchAllowed:             false,
		LaunchEnabled:             false,
		ExecutionRequestCreated:   false,
		ExecutionRequestPersisted: false,
		ActionQueuePersisted:      false,
		ReviewReceiptRecorded:     false,
		ExecutionStarted:          false,
		BackendBindingReady:       false,
		RequestObjectCreated:      false,
		PermissionGranted:         false,
		HostRootModified:          false,
		NetworkRequired:           false,
		BackendDetailsExposed:     false,
		BlockedActions:            []string{"start compatibility profile before preflight passes", "record launch approval from preflight preview", "grant desktop resources before Portal review", "create restore point from preflight preview", "bind compatibility profile from preflight preview", "mutate host root during launch preflight planning", "expose raw backend command to desktop shell"},
		UserFacingSettings:        plan.UserFacingSettings,
		DesktopSafeSummary:        "Runtime launch preflight is visible to KDE, but launch remains blocked until every Runtime gate passes.",
	}
	if err := validateNoBackendTerms(preview, "execution preflight preview"); err != nil {
		return ExecutionPreflightPreview{}, err
	}
	return preview, nil
}

func executionPreflightChecks(decisionPreview ExecutionDecisionPreview) []ExecutionPreflightGate {
	decisionStatus := "pass"
	decisionSummary := "User launch-review intent is captured for Runtime evaluation."
	if decisionPreview.Decision.Decision == "rejected" {
		decisionStatus = "blocked"
		decisionSummary = "User rejected the launch request; Runtime must keep launch blocked."
	} else if decisionPreview.Decision.Decision == "deferred" {
		decisionStatus = "pending"
		decisionSummary = "User deferred the launch request; Runtime must wait for a later review."
	}

	portalStatus := "required"
	if !decisionPreview.PortalRequired {
		portalStatus = "pass"
	}

	return []ExecutionPreflightGate{
		executionPreflightGate("user-decision", decisionStatus, true, decisionSummary),
		executionPreflightGate("portal-policy-review", portalStatus, true, "Portal-mediated desktop resource access must be reviewed before launch."),
		executionPreflightGate("snapshot-baseline", "required", true, "Runtime-managed restore point must exist before launch."),
		executionPreflightGate("backend-binding", "pending", true, "A managed compatibility profile binding is required before launch."),
		executionPreflightGate("runtime-launch-write-gate", "blocked", true, "Runtime Launch remains disabled until production compatibility ownership is ready."),
	}
}

func executionPreflightGate(id string, status string, required bool, summary string) ExecutionPreflightGate {
	return ExecutionPreflightGate{
		ID:       id,
		Status:   status,
		Required: required,
		Summary:  summary,
	}
}

func countExecutionPreflightChecks(checks []ExecutionPreflightGate) (int, int, int, int) {
	passedCheckCount := 0
	requiredCheckCount := 0
	pendingCheckCount := 0
	blockedCheckCount := 0
	for _, check := range checks {
		switch check.Status {
		case "pass":
			passedCheckCount++
		case "required":
			requiredCheckCount++
		case "pending":
			pendingCheckCount++
		case "blocked":
			blockedCheckCount++
		}
	}
	return passedCheckCount, requiredCheckCount, pendingCheckCount, blockedCheckCount
}
