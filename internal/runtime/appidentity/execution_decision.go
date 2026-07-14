package appidentity

import (
	"errors"
	"fmt"
)

type ExecutionDecisionPreview struct {
	SchemaVersion             string                  `json:"schema_version"`
	RequestType               string                  `json:"request_type"`
	DecisionType              string                  `json:"decision_type"`
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
	ExecutionReview           ExecutionReviewSummary  `json:"execution_review"`
	Decision                  ExecutionReviewDecision `json:"decision"`
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
	UserDecisionCaptured      bool                    `json:"user_decision_captured"`
	ReviewReceiptRequired     bool                    `json:"review_receipt_required"`
	ReviewReceiptRecorded     bool                    `json:"review_receipt_recorded"`
	RuntimeLaunchApproval     bool                    `json:"runtime_launch_approval"`
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

type ExecutionReviewSummary struct {
	SchemaVersion         string `json:"schema_version"`
	RequestType           string `json:"request_type"`
	ReviewType            string `json:"review_type"`
	RequestState          string `json:"request_state"`
	Source                string `json:"source"`
	ReadMethod            string `json:"read_method"`
	ActionQueueCandidate  bool   `json:"action_queue_candidate"`
	ReviewReceiptRequired bool   `json:"review_receipt_required"`
	ReviewReceiptRecorded bool   `json:"review_receipt_recorded"`
	ActionQueuePersisted  bool   `json:"action_queue_persisted"`
	ExecutionStarted      bool   `json:"execution_started"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

type ExecutionReviewDecision struct {
	Decision                  string `json:"decision"`
	DecisionAccepted          bool   `json:"decision_accepted"`
	DecisionRecorded          bool   `json:"decision_recorded"`
	UserIntentCaptured        bool   `json:"user_intent_captured"`
	RuntimeApprovalGranted    bool   `json:"runtime_approval_granted"`
	ExecutionAllowed          bool   `json:"execution_allowed"`
	ReviewReceiptCreated      bool   `json:"review_receipt_created"`
	QueueStateChanged         bool   `json:"queue_state_changed"`
	PermissionGrantCreated    bool   `json:"permission_grant_created"`
	NextStep                  string `json:"next_step"`
	DesktopNotificationIntent string `json:"desktop_notification_intent"`
}

func (plan Plan) ExecutionDecisionPreview(decision string, fileURIs []string) (ExecutionDecisionPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return ExecutionDecisionPreview{}, err
	}
	if !singleLine(decision) {
		return ExecutionDecisionPreview{}, errors.New("execution decision preview requires a single-line decision")
	}
	if !supportedExecutionReviewDecision(decision) {
		return ExecutionDecisionPreview{}, fmt.Errorf("unsupported execution review decision: %s", decision)
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return ExecutionDecisionPreview{}, errors.New("execution decision preview requires single-line identity fields")
		}
	}
	review, err := plan.ExecutionReviewPreview(fileURIs)
	if err != nil {
		return ExecutionDecisionPreview{}, err
	}

	preview := ExecutionDecisionPreview{
		SchemaVersion:   "xnix.runtime.request_decision.v1",
		RequestType:     "execution-decision-preview",
		DecisionType:    "compatibility-center-launch-decision",
		RequestState:    review.RequestState,
		Source:          "execution-review-preview",
		Desktop:         "KDE Plasma",
		RuntimeMethod:   "Launch",
		ReadMethod:      "GetExecutionDecisionPreview",
		ApplicationID:   plan.ApplicationID,
		ApplicationName: plan.DisplayName,
		Icon:            plan.Icon,
		DesktopFile:     plan.DesktopFile,
		LauncherCommand: plan.LaunchCommand,
		ExecutionReview: ExecutionReviewSummary{
			SchemaVersion:         review.SchemaVersion,
			RequestType:           review.RequestType,
			ReviewType:            review.ReviewType,
			RequestState:          review.RequestState,
			Source:                review.Source,
			ReadMethod:            review.ReadMethod,
			ActionQueueCandidate:  review.ActionQueueCandidate,
			ReviewReceiptRequired: review.ReviewReceiptRequired,
			ReviewReceiptRecorded: review.ReviewReceiptRecorded,
			ActionQueuePersisted:  review.ActionQueuePersisted,
			ExecutionStarted:      review.ExecutionStarted,
			BackendDetailsExposed: review.BackendDetailsExposed,
		},
		Decision: ExecutionReviewDecision{
			Decision:                  decision,
			DecisionAccepted:          true,
			DecisionRecorded:          false,
			UserIntentCaptured:        true,
			RuntimeApprovalGranted:    false,
			ExecutionAllowed:          false,
			ReviewReceiptCreated:      false,
			QueueStateChanged:         false,
			PermissionGrantCreated:    false,
			NextStep:                  executionReviewDecisionNextStep(decision),
			DesktopNotificationIntent: executionReviewDecisionNotificationIntent(decision),
		},
		ReviewCard:                review.ReviewCard,
		ActionQueue:               review.ActionQueue,
		CompatibilityProfile:      review.CompatibilityProfile,
		GateSummary:               review.GateSummary,
		ExecutionState:            review.ExecutionState,
		OverallStatus:             review.OverallStatus,
		WriteGateDecision:         review.WriteGateDecision,
		DenialErrorName:           review.DenialErrorName,
		PortalRequired:            review.PortalRequired,
		SnapshotRequired:          review.SnapshotRequired,
		FileCount:                 review.FileCount,
		FileURIs:                  review.FileURIs,
		RuntimeOwned:              true,
		GoRuntimeBacked:           true,
		KDEPolicyOwner:            false,
		CompatibilityCenterCard:   true,
		SafeForAIDiagnostics:      true,
		DesktopEntryLaunchVisible: true,
		LaunchIntentCaptured:      review.LaunchIntentCaptured,
		ActionQueueCandidate:      review.ActionQueueCandidate,
		UserDecisionCaptured:      true,
		ReviewReceiptRequired:     true,
		ReviewReceiptRecorded:     false,
		RuntimeLaunchApproval:     false,
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
		BlockedActions:            []string{"record launch decision from preview state", "treat user decision as Runtime launch approval", "persist launch review before Runtime gates pass", "start compatibility profile from decision preview", "grant desktop resources from decision preview", "mutate host root during launch decision planning", "expose raw backend command to desktop shell"},
		UserFacingSettings:        plan.UserFacingSettings,
		DesktopSafeSummary:        "Compatibility Center can preview the launch decision, but the Runtime does not record approval or run the application.",
	}
	if err := validateNoBackendTerms(preview, "execution decision preview"); err != nil {
		return ExecutionDecisionPreview{}, err
	}
	return preview, nil
}

func supportedExecutionReviewDecision(decision string) bool {
	switch decision {
	case "reviewed", "approved", "deferred", "rejected":
		return true
	default:
		return false
	}
}

func executionReviewDecisionNextStep(decision string) string {
	switch decision {
	case "rejected":
		return "Keep the launch request blocked and show rejection guidance."
	case "deferred":
		return "Keep the launch request pending for later review."
	case "approved":
		return "Require Runtime gates before recording approval or launching."
	default:
		return "Keep the launch request in review until Runtime gates pass."
	}
}

func executionReviewDecisionNotificationIntent(decision string) string {
	switch decision {
	case "rejected":
		return "show-launch-rejected"
	case "deferred":
		return "show-launch-deferred"
	case "approved":
		return "show-launch-approval-pending"
	default:
		return "show-launch-reviewed"
	}
}
