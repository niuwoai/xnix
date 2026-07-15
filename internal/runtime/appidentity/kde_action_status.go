package appidentity

import (
	"errors"
	"fmt"
)

type KDEActionStatusPreview struct {
	SchemaVersion            string                           `json:"schema_version"`
	RequestType              string                           `json:"request_type"`
	StatusType               string                           `json:"status_type"`
	StatusState              string                           `json:"status_state"`
	Source                   string                           `json:"source"`
	Desktop                  string                           `json:"desktop"`
	RuntimeMethod            string                           `json:"runtime_method"`
	ReadMethod               string                           `json:"read_method"`
	ApplicationID            string                           `json:"application_id"`
	ApplicationName          string                           `json:"application_name"`
	Icon                     string                           `json:"icon"`
	DesktopFile              string                           `json:"desktop_file"`
	LauncherCommand          []string                         `json:"launcher_command"`
	Receipt                  KDEActionStatusReceiptSummary    `json:"receipt"`
	Action                   KDEActionReviewActionSummary     `json:"action"`
	Preflight                KDEActionReceiptPreflightSummary `json:"preflight"`
	UserVisibleState         KDEActionStatusUserVisibleState  `json:"user_visible_state"`
	CompatibilityCenterState string                           `json:"compatibility_center_state"`
	NotificationIntent       string                           `json:"notification_intent"`
	RequiredRuntimeGate      string                           `json:"required_runtime_gate"`
	NextStep                 string                           `json:"next_step"`
	SessionStatus            KDEEntryPointActionSessionStatus `json:"session_status"`
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
	ActionReviewCaptured     bool                             `json:"action_review_captured"`
	ActionPreflightCreated   bool                             `json:"action_preflight_created"`
	ReceiptPreviewCreated    bool                             `json:"receipt_preview_created"`
	StatusPreviewCreated     bool                             `json:"status_preview_created"`
	StatusPersisted          bool                             `json:"status_persisted"`
	DecisionRecorded         bool                             `json:"decision_recorded"`
	ReviewReceiptRecorded    bool                             `json:"review_receipt_recorded"`
	ActionQueuePersisted     bool                             `json:"action_queue_persisted"`
	QueueStateChanged        bool                             `json:"queue_state_changed"`
	SettingsPersisted        bool                             `json:"settings_persisted"`
	NotificationsSent        bool                             `json:"notifications_sent"`
	ResourceGrantCreated     bool                             `json:"resource_grant_created"`
	RuntimeLaunchApproval    bool                             `json:"runtime_launch_approval"`
	LaunchAllowed            bool                             `json:"launch_allowed"`
	LaunchEnabled            bool                             `json:"launch_enabled"`
	ExecutionStarted         bool                             `json:"execution_started"`
	BackendProcessStarted    bool                             `json:"backend_process_started"`
	RequestObjectsCreated    bool                             `json:"request_objects_created"`
	PermissionGrantCreated   bool                             `json:"permission_grant_created"`
	HostRootModified         bool                             `json:"host_root_modified"`
	NetworkRequired          bool                             `json:"network_required"`
	BackendDetailsExposed    bool                             `json:"backend_details_exposed"`
	BlockedActions           []string                         `json:"blocked_actions"`
	UserFacingSettings       map[string]string                `json:"user_facing_settings"`
	DesktopSafeSummary       string                           `json:"desktop_safe_summary"`
}

type KDEActionStatusReceiptSummary struct {
	RequestType           string `json:"request_type"`
	ReceiptType           string `json:"receipt_type"`
	ReceiptID             string `json:"receipt_id"`
	ReceiptState          string `json:"receipt_state"`
	Decision              string `json:"decision"`
	ReceiptPreviewCreated bool   `json:"receipt_preview_created"`
	ReceiptRecordable     bool   `json:"receipt_recordable"`
	DecisionRecorded      bool   `json:"decision_recorded"`
	ReviewReceiptRecorded bool   `json:"review_receipt_recorded"`
	RuntimeLaunchApproval bool   `json:"runtime_launch_approval"`
	LaunchAllowed         bool   `json:"launch_allowed"`
	ExecutionStarted      bool   `json:"execution_started"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

type KDEActionStatusUserVisibleState struct {
	PrimaryLabel              string `json:"primary_label"`
	SecondaryLabel            string `json:"secondary_label"`
	Badge                     string `json:"badge"`
	ActionStateLabel          string `json:"action_state_label"`
	CompatibilityCenterStatus string `json:"compatibility_center_status"`
	NextUserAction            string `json:"next_user_action"`
	UserFacingMode            string `json:"user_facing_mode"`
	UserFacingAccess          string `json:"user_facing_access"`
	BackendDetailsExposed     bool   `json:"backend_details_exposed"`
}

func (plan Plan) KDEActionStatusPreview(actionID string, decision string, fileURIs []string) (KDEActionStatusPreview, error) {
	if !singleLine(actionID) {
		return KDEActionStatusPreview{}, errors.New("KDE action status preview requires a single-line action id")
	}
	if !singleLine(decision) {
		return KDEActionStatusPreview{}, errors.New("KDE action status preview requires a single-line decision")
	}
	if !supportedExecutionReviewDecision(decision) {
		return KDEActionStatusPreview{}, fmt.Errorf("unsupported KDE action status decision: %s", decision)
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return KDEActionStatusPreview{}, errors.New("KDE action status preview requires single-line identity fields")
		}
	}

	receipt, err := plan.KDEActionReceiptPreview(actionID, decision, fileURIs)
	if err != nil {
		return KDEActionStatusPreview{}, err
	}
	state := kdeActionStatusState(decision)

	preview := KDEActionStatusPreview{
		SchemaVersion:   "xnix.runtime.kde_action_status.v1",
		RequestType:     "kde-action-status-preview",
		StatusType:      "compatibility-center-kde-action-status",
		StatusState:     state,
		Source:          "kde-action-receipt-preview",
		Desktop:         "KDE Plasma",
		RuntimeMethod:   "GetKDEActionStatus",
		ReadMethod:      "GetKDEActionStatusPreview",
		ApplicationID:   receipt.ApplicationID,
		ApplicationName: receipt.ApplicationName,
		Icon:            receipt.Icon,
		DesktopFile:     receipt.DesktopFile,
		LauncherCommand: receipt.LauncherCommand,
		Receipt: KDEActionStatusReceiptSummary{
			RequestType:           receipt.RequestType,
			ReceiptType:           receipt.ReceiptType,
			ReceiptID:             receipt.ReceiptID,
			ReceiptState:          receipt.ReceiptState,
			Decision:              receipt.Review.Decision,
			ReceiptPreviewCreated: receipt.ReceiptPreviewCreated,
			ReceiptRecordable:     receipt.ReceiptRecordable,
			DecisionRecorded:      receipt.DecisionRecorded,
			ReviewReceiptRecorded: receipt.ReviewReceiptRecorded,
			RuntimeLaunchApproval: receipt.RuntimeLaunchApproval,
			LaunchAllowed:         receipt.LaunchAllowed,
			ExecutionStarted:      receipt.ExecutionStarted,
			BackendDetailsExposed: receipt.BackendDetailsExposed,
		},
		Action:                   receipt.Action,
		Preflight:                receipt.Preflight,
		UserVisibleState:         kdeActionStatusUserVisibleState(receipt, state),
		CompatibilityCenterState: kdeActionCompatibilityCenterState(state),
		NotificationIntent:       kdeActionStatusNotificationIntent(state),
		RequiredRuntimeGate:      receipt.RequiredRuntimeGate,
		NextStep:                 kdeActionStatusNextStep(decision, receipt.RequiredRuntimeGate),
		SessionStatus:            receipt.SessionStatus,
		FileCount:                receipt.FileCount,
		FileURIs:                 receipt.FileURIs,
		RuntimeOwned:             true,
		GoRuntimeBacked:          true,
		KDEPolicyOwner:           false,
		OfficialDesktopOnly:      true,
		CompatibilityCenterCard:  true,
		SafeForAIDiagnostics:     true,
		UserDecisionCaptured:     receipt.UserDecisionCaptured,
		UserDecisionAllowsLaunch: receipt.UserDecisionAllowsLaunch,
		ActionReviewCaptured:     receipt.ActionReviewCaptured,
		ActionPreflightCreated:   receipt.ActionPreflightCreated,
		ReceiptPreviewCreated:    receipt.ReceiptPreviewCreated,
		StatusPreviewCreated:     true,
		StatusPersisted:          false,
		DecisionRecorded:         false,
		ReviewReceiptRecorded:    false,
		ActionQueuePersisted:     false,
		QueueStateChanged:        false,
		SettingsPersisted:        false,
		NotificationsSent:        false,
		ResourceGrantCreated:     false,
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
		BlockedActions:           []string{"persist KDE action status from preview state", "treat KDE action status as Runtime execution approval", "persist KDE action queue from status preview", "create Runtime request objects from status preview", "grant desktop resources from status preview", "persist compatibility settings from status preview", "send desktop notifications from status preview", "start compatibility profile from status preview", "mutate host root during KDE action status preview", "expose raw backend command to desktop shell"},
		UserFacingSettings:       receipt.UserFacingSettings,
		DesktopSafeSummary:       "Compatibility Center can preview the user-visible status for a reviewed KDE action, but status is not persisted and execution remains blocked.",
	}
	if err := validateNoBackendTerms(preview, "KDE action status preview"); err != nil {
		return KDEActionStatusPreview{}, err
	}
	return preview, nil
}

func kdeActionStatusState(decision string) string {
	switch decision {
	case "rejected":
		return "review-rejected"
	case "deferred":
		return "review-deferred"
	default:
		return "waiting-for-runtime-gates"
	}
}

func kdeActionStatusUserVisibleState(receipt KDEActionReceiptPreview, state string) KDEActionStatusUserVisibleState {
	return KDEActionStatusUserVisibleState{
		PrimaryLabel:              receipt.ApplicationName,
		SecondaryLabel:            kdeActionStatusSecondaryLabel(state, receipt.RequiredRuntimeGate),
		Badge:                     kdeActionStatusBadge(state),
		ActionStateLabel:          kdeActionStatusActionLabel(state),
		CompatibilityCenterStatus: kdeActionCompatibilityCenterState(state),
		NextUserAction:            kdeActionStatusNextUserAction(state),
		UserFacingMode:            receipt.UserFacingSettings["run_mode"],
		UserFacingAccess:          receipt.UserFacingSettings["resource_access"],
		BackendDetailsExposed:     false,
	}
}

func kdeActionStatusSecondaryLabel(state string, runtimeGate string) string {
	switch state {
	case "review-rejected":
		return "Action rejected"
	case "review-deferred":
		return "Action deferred"
	default:
		return fmt.Sprintf("Waiting for %s", runtimeGate)
	}
}

func kdeActionStatusBadge(state string) string {
	switch state {
	case "review-rejected":
		return "Rejected"
	case "review-deferred":
		return "Deferred"
	default:
		return "Waiting"
	}
}

func kdeActionStatusActionLabel(state string) string {
	switch state {
	case "review-rejected":
		return "Rejected by user"
	case "review-deferred":
		return "Deferred for later review"
	default:
		return "Runtime gates required"
	}
}

func kdeActionCompatibilityCenterState(state string) string {
	switch state {
	case "review-rejected":
		return "action-rejected"
	case "review-deferred":
		return "action-deferred"
	default:
		return "waiting-for-runtime-gates"
	}
}

func kdeActionStatusNextUserAction(state string) string {
	switch state {
	case "review-rejected":
		return "Review rejection guidance"
	case "review-deferred":
		return "Review later"
	default:
		return "Wait for Runtime gates"
	}
}

func kdeActionStatusNotificationIntent(state string) string {
	switch state {
	case "review-rejected":
		return "show-kde-action-status-rejected"
	case "review-deferred":
		return "show-kde-action-status-deferred"
	default:
		return "show-kde-action-status-waiting"
	}
}

func kdeActionStatusNextStep(decision string, runtimeGate string) string {
	switch decision {
	case "rejected":
		return "Keep the KDE action rejected and do not start compatibility execution."
	case "deferred":
		return "Keep the KDE action deferred until the user returns to Compatibility Center."
	default:
		return fmt.Sprintf("Keep the KDE action visible while waiting for %s before any execution.", runtimeGate)
	}
}
