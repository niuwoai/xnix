package appidentity

import (
	"errors"
	"fmt"
)

type KDEActionReviewPreview struct {
	SchemaVersion            string                           `json:"schema_version"`
	RequestType              string                           `json:"request_type"`
	ReviewType               string                           `json:"review_type"`
	Source                   string                           `json:"source"`
	Desktop                  string                           `json:"desktop"`
	RuntimeMethod            string                           `json:"runtime_method"`
	ReadMethod               string                           `json:"read_method"`
	ApplicationID            string                           `json:"application_id"`
	ApplicationName          string                           `json:"application_name"`
	Icon                     string                           `json:"icon"`
	DesktopFile              string                           `json:"desktop_file"`
	LauncherCommand          []string                         `json:"launcher_command"`
	Queue                    KDEActionReviewQueueSummary      `json:"queue"`
	Action                   KDEActionReviewActionSummary     `json:"action"`
	Decision                 KDEActionReviewDecisionSummary   `json:"decision"`
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
	ReviewReceiptRequired    bool                             `json:"review_receipt_required"`
	ReviewReceiptRecorded    bool                             `json:"review_receipt_recorded"`
	ActionQueuePersisted     bool                             `json:"action_queue_persisted"`
	QueueStateChanged        bool                             `json:"queue_state_changed"`
	SettingsPersisted        bool                             `json:"settings_persisted"`
	NotificationsSent        bool                             `json:"notifications_sent"`
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

type KDEActionReviewQueueSummary struct {
	QueueType             string `json:"queue_type"`
	SourceRequestType     string `json:"source_request_type"`
	ActionCount           int    `json:"action_count"`
	PendingActionCount    int    `json:"pending_action_count"`
	QueuePersisted        bool   `json:"queue_persisted"`
	ExecutionEnabled      bool   `json:"execution_enabled"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

type KDEActionReviewActionSummary struct {
	ID                    string `json:"id"`
	SourceType            string `json:"source_type"`
	EntryPointID          string `json:"entry_point_id"`
	EntryPointLabel       string `json:"entry_point_label"`
	KDEComponent          string `json:"kde_component"`
	Intent                string `json:"intent"`
	Status                string `json:"status"`
	Title                 string `json:"title"`
	UserReviewRequired    bool   `json:"user_review_required"`
	RuntimeGate           string `json:"runtime_gate"`
	NextStep              string `json:"next_step"`
	RequiresPortal        bool   `json:"requires_portal"`
	RequiresRuntimeGate   bool   `json:"requires_runtime_gate"`
	ExecutionEnabled      bool   `json:"execution_enabled"`
	RequestObjectCreated  bool   `json:"request_object_created"`
	BackendProcessStarted bool   `json:"backend_process_started"`
	HostRootModified      bool   `json:"host_root_modified"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

type KDEActionReviewDecisionSummary struct {
	Decision                  string `json:"decision"`
	DecisionAccepted          bool   `json:"decision_accepted"`
	UserIntentCaptured        bool   `json:"user_intent_captured"`
	DecisionRecorded          bool   `json:"decision_recorded"`
	RuntimeApprovalGranted    bool   `json:"runtime_approval_granted"`
	ExecutionAllowed          bool   `json:"execution_allowed"`
	ReviewReceiptCreated      bool   `json:"review_receipt_created"`
	QueueStateChanged         bool   `json:"queue_state_changed"`
	PermissionGrantCreated    bool   `json:"permission_grant_created"`
	SettingsPersisted         bool   `json:"settings_persisted"`
	DesktopNotificationIntent string `json:"desktop_notification_intent"`
	NextStep                  string `json:"next_step"`
}

func (plan Plan) KDEActionReviewPreview(actionID string, decision string, fileURIs []string) (KDEActionReviewPreview, error) {
	if !singleLine(actionID) {
		return KDEActionReviewPreview{}, errors.New("KDE action review preview requires a single-line action id")
	}
	if !singleLine(decision) {
		return KDEActionReviewPreview{}, errors.New("KDE action review preview requires a single-line decision")
	}
	if !supportedExecutionReviewDecision(decision) {
		return KDEActionReviewPreview{}, fmt.Errorf("unsupported KDE action review decision: %s", decision)
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return KDEActionReviewPreview{}, errors.New("KDE action review preview requires single-line identity fields")
		}
	}

	queue, err := plan.KDEActionQueuePreview(decision, fileURIs)
	if err != nil {
		return KDEActionReviewPreview{}, err
	}
	action, ok := findKDEActionQueuePreviewItem(queue.Actions, actionID)
	if !ok {
		return KDEActionReviewPreview{}, fmt.Errorf("unknown KDE action queue item: %s", actionID)
	}

	preview := KDEActionReviewPreview{
		SchemaVersion:   "xnix.runtime.kde_action_review.v1",
		RequestType:     "kde-action-review-preview",
		ReviewType:      "compatibility-center-kde-action-review",
		Source:          "kde-action-queue-preview",
		Desktop:         "KDE Plasma",
		RuntimeMethod:   "ReviewKDEAction",
		ReadMethod:      "GetKDEActionReviewPreview",
		ApplicationID:   queue.ApplicationID,
		ApplicationName: queue.ApplicationName,
		Icon:            queue.Icon,
		DesktopFile:     queue.DesktopFile,
		LauncherCommand: queue.LauncherCommand,
		Queue: KDEActionReviewQueueSummary{
			QueueType:             queue.QueueType,
			SourceRequestType:     queue.RequestType,
			ActionCount:           queue.ActionCount,
			PendingActionCount:    queue.PendingActionCount,
			QueuePersisted:        false,
			ExecutionEnabled:      false,
			BackendDetailsExposed: false,
		},
		Action: KDEActionReviewActionSummary{
			ID:                    action.ID,
			SourceType:            action.SourceType,
			EntryPointID:          action.EntryPointID,
			EntryPointLabel:       action.EntryPointLabel,
			KDEComponent:          action.KDEComponent,
			Intent:                action.Intent,
			Status:                action.Status,
			Title:                 action.Title,
			UserReviewRequired:    action.UserReviewRequired,
			RuntimeGate:           action.RuntimeGate,
			NextStep:              action.NextStep,
			RequiresPortal:        action.RequiresPortal,
			RequiresRuntimeGate:   action.RequiresRuntimeGate,
			ExecutionEnabled:      false,
			RequestObjectCreated:  false,
			BackendProcessStarted: false,
			HostRootModified:      false,
			BackendDetailsExposed: false,
		},
		Decision: KDEActionReviewDecisionSummary{
			Decision:                  decision,
			DecisionAccepted:          true,
			UserIntentCaptured:        true,
			DecisionRecorded:          false,
			RuntimeApprovalGranted:    false,
			ExecutionAllowed:          false,
			ReviewReceiptCreated:      false,
			QueueStateChanged:         false,
			PermissionGrantCreated:    false,
			SettingsPersisted:         false,
			DesktopNotificationIntent: kdeActionReviewNotificationIntent(decision),
			NextStep:                  kdeActionReviewNextStep(decision, action.RuntimeGate),
		},
		SessionStatus: KDEEntryPointActionSessionStatus{
			RequestType:           queue.SessionStatus.RequestType,
			StatusType:            queue.SessionStatus.StatusType,
			SessionState:          queue.SessionStatus.SessionState,
			GateCount:             queue.SessionStatus.GateCount,
			BlockedGateCount:      queue.SessionStatus.BlockedGateCount,
			DesktopSurfaceState:   queue.SessionStatus.DesktopSurfaceState,
			UserVisibleState:      queue.SessionStatus.UserVisibleState,
			RuntimeLaunchApproval: false,
			LaunchAllowed:         false,
			ExecutionStarted:      false,
		},
		FileCount:                queue.FileCount,
		FileURIs:                 queue.FileURIs,
		RuntimeOwned:             true,
		GoRuntimeBacked:          true,
		KDEPolicyOwner:           false,
		OfficialDesktopOnly:      true,
		CompatibilityCenterCard:  true,
		SafeForAIDiagnostics:     true,
		UserDecisionCaptured:     queue.UserDecisionCaptured,
		UserDecisionAllowsLaunch: queue.UserDecisionAllowsLaunch,
		ActionReviewCaptured:     true,
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
		BlockedActions:           []string{"record KDE action review from preview state", "treat KDE review intent as Runtime execution approval", "persist KDE action queue from review preview", "create Runtime request objects from review preview", "persist compatibility settings from review preview", "send desktop notifications from review preview", "start compatibility profile from review preview", "grant desktop resources from review preview", "mutate host root during KDE action review preview", "expose raw backend command to desktop shell"},
		UserFacingSettings:       queue.UserFacingSettings,
		DesktopSafeSummary:       "Compatibility Center can preview KDE action review intent, but the Runtime does not record receipts, mutate queue state, or start execution.",
	}
	if err := validateNoBackendTerms(preview, "KDE action review preview"); err != nil {
		return KDEActionReviewPreview{}, err
	}
	return preview, nil
}

func findKDEActionQueuePreviewItem(actions []KDEActionQueueItem, id string) (KDEActionQueueItem, bool) {
	for _, action := range actions {
		if action.ID == id {
			return action, true
		}
	}
	return KDEActionQueueItem{}, false
}

func kdeActionReviewNextStep(decision string, runtimeGate string) string {
	switch decision {
	case "rejected":
		return "Keep the queued KDE action blocked and show rejection guidance."
	case "deferred":
		return "Keep the queued KDE action pending for later review."
	case "approved":
		return fmt.Sprintf("Keep the queued KDE action waiting for %s before any execution.", runtimeGate)
	default:
		return fmt.Sprintf("Record review intent later only after %s is available.", runtimeGate)
	}
}

func kdeActionReviewNotificationIntent(decision string) string {
	switch decision {
	case "rejected":
		return "show-kde-action-rejected"
	case "deferred":
		return "show-kde-action-deferred"
	case "approved":
		return "show-kde-action-approval-pending"
	default:
		return "show-kde-action-reviewed"
	}
}
