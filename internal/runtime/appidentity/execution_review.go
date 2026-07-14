package appidentity

import "errors"

type ExecutionReviewPreview struct {
	SchemaVersion             string                  `json:"schema_version"`
	RequestType               string                  `json:"request_type"`
	ReviewType                string                  `json:"review_type"`
	RequestState              string                  `json:"request_state"`
	Source                    string                  `json:"source"`
	Desktop                   string                  `json:"desktop"`
	RuntimeMethod             string                  `json:"runtime_method"`
	ReadMethod                string                  `json:"read_method"`
	ApplicationID             string                  `json:"application_id"`
	ApplicationName           string                  `json:"application_name"`
	Icon                      string                  `json:"icon"`
	DesktopFile               string                  `json:"desktop_file"`
	LauncherCommand           []string                `json:"launcher_command"`
	ExecutionRequest          ExecutionRequestSummary `json:"execution_request"`
	ReviewCard                ExecutionReviewCard     `json:"review_card"`
	ActionQueue               ExecutionReviewQueue    `json:"action_queue"`
	CompatibilityProfile      ReadinessProfile        `json:"compatibility_profile"`
	GateSummary               ExecutionGateSummary    `json:"gate_summary"`
	ExecutionState            string                  `json:"execution_state"`
	OverallStatus             string                  `json:"overall_status"`
	WriteGateDecision         string                  `json:"write_gate_decision"`
	DenialErrorName           string                  `json:"denial_error_name"`
	PortalRequired            bool                    `json:"portal_required"`
	SnapshotRequired          bool                    `json:"snapshot_required"`
	FileCount                 int                     `json:"file_count"`
	FileURIs                  []string                `json:"file_uris"`
	RuntimeOwned              bool                    `json:"runtime_owned"`
	GoRuntimeBacked           bool                    `json:"go_runtime_backed"`
	KDEPolicyOwner            bool                    `json:"kde_policy_owner"`
	CompatibilityCenterCard   bool                    `json:"compatibility_center_card"`
	SafeForAIDiagnostics      bool                    `json:"safe_for_ai_diagnostics"`
	DesktopEntryLaunchVisible bool                    `json:"desktop_entry_launch_visible"`
	LaunchIntentCaptured      bool                    `json:"launch_intent_captured"`
	ActionQueueCandidate      bool                    `json:"action_queue_candidate"`
	ReviewReceiptRequired     bool                    `json:"review_receipt_required"`
	ReviewReceiptRecorded     bool                    `json:"review_receipt_recorded"`
	LaunchAllowed             bool                    `json:"launch_allowed"`
	LaunchEnabled             bool                    `json:"launch_enabled"`
	ExecutionRequestCreated   bool                    `json:"execution_request_created"`
	ExecutionRequestPersisted bool                    `json:"execution_request_persisted"`
	ActionQueuePersisted      bool                    `json:"action_queue_persisted"`
	ExecutionStarted          bool                    `json:"execution_started"`
	BackendBindingReady       bool                    `json:"backend_binding_ready"`
	RequestObjectCreated      bool                    `json:"request_object_created"`
	PermissionGranted         bool                    `json:"permission_granted"`
	HostRootModified          bool                    `json:"host_root_modified"`
	NetworkRequired           bool                    `json:"network_required"`
	BackendDetailsExposed     bool                    `json:"backend_details_exposed"`
	BlockedActions            []string                `json:"blocked_actions"`
	UserFacingSettings        map[string]string       `json:"user_facing_settings"`
	DesktopSafeSummary        string                  `json:"desktop_safe_summary"`
}

type ExecutionRequestSummary struct {
	SchemaVersion             string `json:"schema_version"`
	RequestType               string `json:"request_type"`
	RequestState              string `json:"request_state"`
	Source                    string `json:"source"`
	ReadMethod                string `json:"read_method"`
	WriteGateDecision         string `json:"write_gate_decision"`
	PortalRequired            bool   `json:"portal_required"`
	SnapshotRequired          bool   `json:"snapshot_required"`
	FileCount                 int    `json:"file_count"`
	LaunchIntentCaptured      bool   `json:"launch_intent_captured"`
	ExecutionRequestCreated   bool   `json:"execution_request_created"`
	ExecutionRequestPersisted bool   `json:"execution_request_persisted"`
	ExecutionStarted          bool   `json:"execution_started"`
	BackendDetailsExposed     bool   `json:"backend_details_exposed"`
}

type ExecutionReviewCard struct {
	ID                    string   `json:"id"`
	CardType              string   `json:"card_type"`
	Title                 string   `json:"title"`
	Status                string   `json:"status"`
	Severity              string   `json:"severity"`
	UserReviewRequired    bool     `json:"user_review_required"`
	RuntimeApprovalNeeded bool     `json:"runtime_approval_needed"`
	PrimaryAction         string   `json:"primary_action"`
	SecondaryActions      []string `json:"secondary_actions"`
	Summary               string   `json:"summary"`
}

type ExecutionReviewQueue struct {
	QueueType               string   `json:"queue_type"`
	ActionCount             int      `json:"action_count"`
	PendingActionCount      int      `json:"pending_action_count"`
	UserReviewRequiredCount int      `json:"user_review_required_count"`
	ExecutionEnabled        bool     `json:"execution_enabled"`
	QueuePersisted          bool     `json:"queue_persisted"`
	ReviewReceiptRecorded   bool     `json:"review_receipt_recorded"`
	Actions                 []string `json:"actions"`
}

func (plan Plan) ExecutionReviewPreview(fileURIs []string) (ExecutionReviewPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return ExecutionReviewPreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return ExecutionReviewPreview{}, errors.New("execution review preview requires single-line identity fields")
		}
	}
	request, err := plan.ExecutionRequestPreview(fileURIs)
	if err != nil {
		return ExecutionReviewPreview{}, err
	}

	preview := ExecutionReviewPreview{
		SchemaVersion:   "xnix.runtime.request_review.v1",
		RequestType:     "execution-review-preview",
		ReviewType:      "compatibility-center-launch-review",
		RequestState:    request.RequestState,
		Source:          "execution-request-preview",
		Desktop:         "KDE Plasma",
		RuntimeMethod:   "Launch",
		ReadMethod:      "GetExecutionReviewPreview",
		ApplicationID:   plan.ApplicationID,
		ApplicationName: plan.DisplayName,
		Icon:            plan.Icon,
		DesktopFile:     plan.DesktopFile,
		LauncherCommand: plan.LaunchCommand,
		ExecutionRequest: ExecutionRequestSummary{
			SchemaVersion:             request.SchemaVersion,
			RequestType:               request.RequestType,
			RequestState:              request.RequestState,
			Source:                    request.Source,
			ReadMethod:                request.ReadMethod,
			WriteGateDecision:         request.WriteGateDecision,
			PortalRequired:            request.PortalRequired,
			SnapshotRequired:          request.SnapshotRequired,
			FileCount:                 request.FileCount,
			LaunchIntentCaptured:      request.LaunchIntentCaptured,
			ExecutionRequestCreated:   request.ExecutionRequestCreated,
			ExecutionRequestPersisted: request.ExecutionRequestPersisted,
			ExecutionStarted:          request.ExecutionStarted,
			BackendDetailsExposed:     request.BackendDetailsExposed,
		},
		ReviewCard: ExecutionReviewCard{
			ID:                    plan.ApplicationID + ":launch-review",
			CardType:              "compatibility-center-launch-review",
			Title:                 "Review launch request",
			Status:                "blocked",
			Severity:              "requires-runtime-gates",
			UserReviewRequired:    true,
			RuntimeApprovalNeeded: true,
			PrimaryAction:         "Open Compatibility Center",
			SecondaryActions:      []string{"Review resource access", "Review restore point", "View diagnostics"},
			Summary:               "This launch request is ready for review but cannot run until Runtime gates pass.",
		},
		ActionQueue: ExecutionReviewQueue{
			QueueType:               "compatibility-center-request-review-queue",
			ActionCount:             1,
			PendingActionCount:      1,
			UserReviewRequiredCount: 1,
			ExecutionEnabled:        false,
			QueuePersisted:          false,
			ReviewReceiptRecorded:   false,
			Actions:                 []string{"review-launch-request"},
		},
		CompatibilityProfile:      request.CompatibilityProfile,
		GateSummary:               request.GateSummary,
		ExecutionState:            request.ExecutionState,
		OverallStatus:             request.OverallStatus,
		WriteGateDecision:         request.WriteGateDecision,
		DenialErrorName:           request.DenialErrorName,
		PortalRequired:            request.PortalRequired,
		SnapshotRequired:          request.SnapshotRequired,
		FileCount:                 request.FileCount,
		FileURIs:                  request.FileURIs,
		RuntimeOwned:              true,
		GoRuntimeBacked:           true,
		KDEPolicyOwner:            false,
		CompatibilityCenterCard:   true,
		SafeForAIDiagnostics:      true,
		DesktopEntryLaunchVisible: true,
		LaunchIntentCaptured:      request.LaunchIntentCaptured,
		ActionQueueCandidate:      true,
		ReviewReceiptRequired:     true,
		ReviewReceiptRecorded:     false,
		LaunchAllowed:             false,
		LaunchEnabled:             false,
		ExecutionRequestCreated:   false,
		ExecutionRequestPersisted: false,
		ActionQueuePersisted:      false,
		ExecutionStarted:          false,
		BackendBindingReady:       false,
		RequestObjectCreated:      false,
		PermissionGranted:         false,
		HostRootModified:          false,
		NetworkRequired:           false,
		BackendDetailsExposed:     false,
		BlockedActions:            []string{"persist launch review before Runtime gates pass", "record launch approval from preview state", "start compatibility profile from review card", "grant desktop resources from review card", "mutate host root during launch review planning", "expose raw backend command to desktop shell"},
		UserFacingSettings:        plan.UserFacingSettings,
		DesktopSafeSummary:        "Compatibility Center can show a launch review card, but the Runtime does not record approval or run the application.",
	}
	if err := validateNoBackendTerms(preview, "execution review preview"); err != nil {
		return ExecutionReviewPreview{}, err
	}
	return preview, nil
}
