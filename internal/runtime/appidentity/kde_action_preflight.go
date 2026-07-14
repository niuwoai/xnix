package appidentity

import (
	"errors"
	"fmt"
)

type KDEActionPreflightPreview struct {
	SchemaVersion            string                             `json:"schema_version"`
	RequestType              string                             `json:"request_type"`
	PreflightType            string                             `json:"preflight_type"`
	Source                   string                             `json:"source"`
	Desktop                  string                             `json:"desktop"`
	RuntimeMethod            string                             `json:"runtime_method"`
	ReadMethod               string                             `json:"read_method"`
	ApplicationID            string                             `json:"application_id"`
	ApplicationName          string                             `json:"application_name"`
	Icon                     string                             `json:"icon"`
	DesktopFile              string                             `json:"desktop_file"`
	LauncherCommand          []string                           `json:"launcher_command"`
	Review                   KDEActionPreflightReviewSummary    `json:"review"`
	Action                   KDEActionReviewActionSummary       `json:"action"`
	ExecutionPreflight       KDEActionExecutionPreflightSummary `json:"execution_preflight"`
	PreflightChecks          []KDEActionPreflightGate           `json:"preflight_checks"`
	CheckCount               int                                `json:"check_count"`
	PassedCheckCount         int                                `json:"passed_check_count"`
	RequiredCheckCount       int                                `json:"required_check_count"`
	PendingCheckCount        int                                `json:"pending_check_count"`
	BlockedCheckCount        int                                `json:"blocked_check_count"`
	SessionStatus            KDEEntryPointActionSessionStatus   `json:"session_status"`
	FileCount                int                                `json:"file_count"`
	FileURIs                 []string                           `json:"file_uris"`
	RuntimeOwned             bool                               `json:"runtime_owned"`
	GoRuntimeBacked          bool                               `json:"go_runtime_backed"`
	KDEPolicyOwner           bool                               `json:"kde_policy_owner"`
	OfficialDesktopOnly      bool                               `json:"official_desktop_only"`
	CompatibilityCenterCard  bool                               `json:"compatibility_center_card"`
	SafeForAIDiagnostics     bool                               `json:"safe_for_ai_diagnostics"`
	UserDecisionCaptured     bool                               `json:"user_decision_captured"`
	UserDecisionAllowsLaunch bool                               `json:"user_decision_allows_launch"`
	ActionReviewCaptured     bool                               `json:"action_review_captured"`
	ActionPreflightCreated   bool                               `json:"action_preflight_created"`
	ReviewAllowsPreflight    bool                               `json:"review_allows_preflight"`
	PreflightComplete        bool                               `json:"preflight_complete"`
	PreflightPassed          bool                               `json:"preflight_passed"`
	PortalPreflightReady     bool                               `json:"portal_preflight_ready"`
	ReviewReceiptRequired    bool                               `json:"review_receipt_required"`
	ReviewReceiptRecorded    bool                               `json:"review_receipt_recorded"`
	ActionQueuePersisted     bool                               `json:"action_queue_persisted"`
	QueueStateChanged        bool                               `json:"queue_state_changed"`
	SettingsPersisted        bool                               `json:"settings_persisted"`
	NotificationsSent        bool                               `json:"notifications_sent"`
	RuntimeLaunchApproval    bool                               `json:"runtime_launch_approval"`
	LaunchAllowed            bool                               `json:"launch_allowed"`
	LaunchEnabled            bool                               `json:"launch_enabled"`
	ExecutionStarted         bool                               `json:"execution_started"`
	BackendProcessStarted    bool                               `json:"backend_process_started"`
	RequestObjectsCreated    bool                               `json:"request_objects_created"`
	PermissionGrantCreated   bool                               `json:"permission_grant_created"`
	HostRootModified         bool                               `json:"host_root_modified"`
	NetworkRequired          bool                               `json:"network_required"`
	BackendDetailsExposed    bool                               `json:"backend_details_exposed"`
	BlockedActions           []string                           `json:"blocked_actions"`
	UserFacingSettings       map[string]string                  `json:"user_facing_settings"`
	DesktopSafeSummary       string                             `json:"desktop_safe_summary"`
}

type KDEActionPreflightReviewSummary struct {
	RequestType               string `json:"request_type"`
	ReviewType                string `json:"review_type"`
	Decision                  string `json:"decision"`
	DecisionAccepted          bool   `json:"decision_accepted"`
	UserIntentCaptured        bool   `json:"user_intent_captured"`
	DecisionRecorded          bool   `json:"decision_recorded"`
	RuntimeApprovalGranted    bool   `json:"runtime_approval_granted"`
	ExecutionAllowed          bool   `json:"execution_allowed"`
	ReviewReceiptCreated      bool   `json:"review_receipt_created"`
	QueueStateChanged         bool   `json:"queue_state_changed"`
	PermissionGrantCreated    bool   `json:"permission_grant_created"`
	DesktopNotificationIntent string `json:"desktop_notification_intent"`
	NextStep                  string `json:"next_step"`
}

type KDEActionExecutionPreflightSummary struct {
	RequestType           string `json:"request_type"`
	PreflightType         string `json:"preflight_type"`
	RequestState          string `json:"request_state"`
	CheckCount            int    `json:"check_count"`
	PassedCheckCount      int    `json:"passed_check_count"`
	RequiredCheckCount    int    `json:"required_check_count"`
	PendingCheckCount     int    `json:"pending_check_count"`
	BlockedCheckCount     int    `json:"blocked_check_count"`
	WriteGateDecision     string `json:"write_gate_decision"`
	DenialErrorName       string `json:"denial_error_name"`
	PortalRequired        bool   `json:"portal_required"`
	SnapshotRequired      bool   `json:"snapshot_required"`
	PreflightComplete     bool   `json:"preflight_complete"`
	PreflightPassed       bool   `json:"preflight_passed"`
	RuntimeLaunchApproval bool   `json:"runtime_launch_approval"`
	LaunchAllowed         bool   `json:"launch_allowed"`
	ExecutionStarted      bool   `json:"execution_started"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

type KDEActionPreflightGate struct {
	ID       string `json:"id"`
	Status   string `json:"status"`
	Required bool   `json:"required"`
	Source   string `json:"source"`
	Summary  string `json:"summary"`
}

func (plan Plan) KDEActionPreflightPreview(actionID string, decision string, fileURIs []string) (KDEActionPreflightPreview, error) {
	if !singleLine(actionID) {
		return KDEActionPreflightPreview{}, errors.New("KDE action preflight preview requires a single-line action id")
	}
	if !singleLine(decision) {
		return KDEActionPreflightPreview{}, errors.New("KDE action preflight preview requires a single-line decision")
	}
	if !supportedExecutionReviewDecision(decision) {
		return KDEActionPreflightPreview{}, fmt.Errorf("unsupported KDE action preflight decision: %s", decision)
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return KDEActionPreflightPreview{}, errors.New("KDE action preflight preview requires single-line identity fields")
		}
	}

	review, err := plan.KDEActionReviewPreview(actionID, decision, fileURIs)
	if err != nil {
		return KDEActionPreflightPreview{}, err
	}
	executionPreflight, err := plan.ExecutionPreflightPreview(decision, fileURIs)
	if err != nil {
		return KDEActionPreflightPreview{}, err
	}

	preflightChecks := kdeActionPreflightChecks(review)
	passedCheckCount, requiredCheckCount, pendingCheckCount, blockedCheckCount := countKDEActionPreflightChecks(preflightChecks)
	reviewAllowsPreflight := decision == "approved" || decision == "reviewed"

	preview := KDEActionPreflightPreview{
		SchemaVersion:   "xnix.runtime.kde_action_preflight.v1",
		RequestType:     "kde-action-preflight-preview",
		PreflightType:   "compatibility-center-kde-action-preflight",
		Source:          "kde-action-review-preview",
		Desktop:         "KDE Plasma",
		RuntimeMethod:   "PreflightKDEAction",
		ReadMethod:      "GetKDEActionPreflightPreview",
		ApplicationID:   review.ApplicationID,
		ApplicationName: review.ApplicationName,
		Icon:            review.Icon,
		DesktopFile:     review.DesktopFile,
		LauncherCommand: review.LauncherCommand,
		Review: KDEActionPreflightReviewSummary{
			RequestType:               review.RequestType,
			ReviewType:                review.ReviewType,
			Decision:                  review.Decision.Decision,
			DecisionAccepted:          review.Decision.DecisionAccepted,
			UserIntentCaptured:        review.Decision.UserIntentCaptured,
			DecisionRecorded:          review.Decision.DecisionRecorded,
			RuntimeApprovalGranted:    review.Decision.RuntimeApprovalGranted,
			ExecutionAllowed:          review.Decision.ExecutionAllowed,
			ReviewReceiptCreated:      review.Decision.ReviewReceiptCreated,
			QueueStateChanged:         review.Decision.QueueStateChanged,
			PermissionGrantCreated:    review.Decision.PermissionGrantCreated,
			DesktopNotificationIntent: review.Decision.DesktopNotificationIntent,
			NextStep:                  review.Decision.NextStep,
		},
		Action: review.Action,
		ExecutionPreflight: KDEActionExecutionPreflightSummary{
			RequestType:           executionPreflight.RequestType,
			PreflightType:         executionPreflight.PreflightType,
			RequestState:          executionPreflight.RequestState,
			CheckCount:            executionPreflight.CheckCount,
			PassedCheckCount:      executionPreflight.PassedCheckCount,
			RequiredCheckCount:    executionPreflight.RequiredCheckCount,
			PendingCheckCount:     executionPreflight.PendingCheckCount,
			BlockedCheckCount:     executionPreflight.BlockedCheckCount,
			WriteGateDecision:     executionPreflight.WriteGateDecision,
			DenialErrorName:       executionPreflight.DenialErrorName,
			PortalRequired:        executionPreflight.PortalRequired,
			SnapshotRequired:      executionPreflight.SnapshotRequired,
			PreflightComplete:     executionPreflight.PreflightComplete,
			PreflightPassed:       executionPreflight.PreflightPassed,
			RuntimeLaunchApproval: executionPreflight.RuntimeLaunchApproval,
			LaunchAllowed:         executionPreflight.LaunchAllowed,
			ExecutionStarted:      executionPreflight.ExecutionStarted,
			BackendDetailsExposed: executionPreflight.BackendDetailsExposed,
		},
		PreflightChecks:          preflightChecks,
		CheckCount:               len(preflightChecks),
		PassedCheckCount:         passedCheckCount,
		RequiredCheckCount:       requiredCheckCount,
		PendingCheckCount:        pendingCheckCount,
		BlockedCheckCount:        blockedCheckCount,
		SessionStatus:            review.SessionStatus,
		FileCount:                review.FileCount,
		FileURIs:                 review.FileURIs,
		RuntimeOwned:             true,
		GoRuntimeBacked:          true,
		KDEPolicyOwner:           false,
		OfficialDesktopOnly:      true,
		CompatibilityCenterCard:  true,
		SafeForAIDiagnostics:     true,
		UserDecisionCaptured:     review.UserDecisionCaptured,
		UserDecisionAllowsLaunch: review.UserDecisionAllowsLaunch,
		ActionReviewCaptured:     true,
		ActionPreflightCreated:   true,
		ReviewAllowsPreflight:    reviewAllowsPreflight,
		PreflightComplete:        false,
		PreflightPassed:          false,
		PortalPreflightReady:     false,
		ReviewReceiptRequired:    true,
		ReviewReceiptRecorded:    false,
		ActionQueuePersisted:     false,
		QueueStateChanged:        false,
		SettingsPersisted:        false,
		NotificationsSent:        false,
		RuntimeLaunchApproval:    false,
		LaunchAllowed:            false,
		LaunchEnabled:            false,
		ExecutionStarted:         false,
		BackendProcessStarted:    false,
		RequestObjectsCreated:    false,
		PermissionGrantCreated:   false,
		HostRootModified:         false,
		NetworkRequired:          false,
		BackendDetailsExposed:    false,
		BlockedActions:           []string{"treat KDE action preflight as Runtime execution approval", "record review receipt from action preflight preview", "persist KDE action queue from action preflight preview", "create Runtime request objects from action preflight preview", "grant desktop resources from action preflight preview", "send desktop notifications from action preflight preview", "start compatibility profile from action preflight preview", "mutate host root during KDE action preflight preview", "expose raw backend command to desktop shell"},
		UserFacingSettings:       review.UserFacingSettings,
		DesktopSafeSummary:       "Compatibility Center can preview the preflight gates for one reviewed KDE action, but the Runtime still records no receipt, creates no request object, and starts no execution.",
	}
	if err := validateNoBackendTerms(preview, "KDE action preflight preview"); err != nil {
		return KDEActionPreflightPreview{}, err
	}
	return preview, nil
}

func kdeActionPreflightChecks(review KDEActionReviewPreview) []KDEActionPreflightGate {
	reviewStatus := "pass"
	reviewSummary := "KDE action review intent is captured for Runtime preflight."
	switch review.Decision.Decision {
	case "rejected":
		reviewStatus = "blocked"
		reviewSummary = "User rejected the KDE action; Runtime must keep action execution blocked."
	case "deferred":
		reviewStatus = "pending"
		reviewSummary = "User deferred the KDE action; Runtime must wait for a later review."
	}

	portalStatus := "pass"
	if review.Action.RequiresPortal {
		portalStatus = "required"
	}

	return []KDEActionPreflightGate{
		kdeActionPreflightGate("kde-action-review", reviewStatus, true, review.RequestType, reviewSummary),
		kdeActionPreflightGate("portal-policy-review", portalStatus, true, "kde-action-review-preview", "Portal-mediated desktop resource access must be reviewed before action execution."),
		kdeActionPreflightGate("review-receipt", "pending", true, "kde-action-review-preview", "Runtime-owned review receipt must be recorded before queue state changes."),
		kdeActionPreflightGate("request-object", "blocked", true, "kde-action-review-preview", "Runtime request objects remain disabled during KDE action preflight preview."),
		kdeActionPreflightGate("runtime-launch-write-gate", "blocked", true, "execution-preflight-preview", "Runtime Launch remains disabled until production compatibility ownership is ready."),
	}
}

func kdeActionPreflightGate(id string, status string, required bool, source string, summary string) KDEActionPreflightGate {
	return KDEActionPreflightGate{
		ID:       id,
		Status:   status,
		Required: required,
		Source:   source,
		Summary:  summary,
	}
}

func countKDEActionPreflightChecks(checks []KDEActionPreflightGate) (int, int, int, int) {
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
