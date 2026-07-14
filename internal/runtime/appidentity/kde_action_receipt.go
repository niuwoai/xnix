package appidentity

import (
	"errors"
	"fmt"
	"strings"
)

type KDEActionReceiptPreview struct {
	SchemaVersion            string                           `json:"schema_version"`
	RequestType              string                           `json:"request_type"`
	ReceiptType              string                           `json:"receipt_type"`
	ReceiptID                string                           `json:"receipt_id"`
	ReceiptState             string                           `json:"receipt_state"`
	Source                   string                           `json:"source"`
	Desktop                  string                           `json:"desktop"`
	RuntimeMethod            string                           `json:"runtime_method"`
	ReadMethod               string                           `json:"read_method"`
	ApplicationID            string                           `json:"application_id"`
	ApplicationName          string                           `json:"application_name"`
	Icon                     string                           `json:"icon"`
	DesktopFile              string                           `json:"desktop_file"`
	LauncherCommand          []string                         `json:"launcher_command"`
	Review                   KDEActionReceiptReviewSummary    `json:"review"`
	Action                   KDEActionReviewActionSummary     `json:"action"`
	Preflight                KDEActionReceiptPreflightSummary `json:"preflight"`
	ReceiptFields            []KDEActionReceiptField          `json:"receipt_fields"`
	FieldCount               int                              `json:"field_count"`
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
	ReceiptRecordable        bool                             `json:"receipt_recordable"`
	DecisionRecorded         bool                             `json:"decision_recorded"`
	ReviewReceiptCreated     bool                             `json:"review_receipt_created"`
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

type KDEActionReceiptReviewSummary struct {
	RequestType            string `json:"request_type"`
	ReviewType             string `json:"review_type"`
	Decision               string `json:"decision"`
	DecisionAccepted       bool   `json:"decision_accepted"`
	UserIntentCaptured     bool   `json:"user_intent_captured"`
	DecisionRecorded       bool   `json:"decision_recorded"`
	RuntimeApprovalGranted bool   `json:"runtime_approval_granted"`
	ExecutionAllowed       bool   `json:"execution_allowed"`
	ReviewReceiptCreated   bool   `json:"review_receipt_created"`
	QueueStateChanged      bool   `json:"queue_state_changed"`
	PermissionGrantCreated bool   `json:"permission_grant_created"`
}

type KDEActionReceiptPreflightSummary struct {
	RequestType           string `json:"request_type"`
	PreflightType         string `json:"preflight_type"`
	ReceiptGateStatus     string `json:"receipt_gate_status"`
	CheckCount            int    `json:"check_count"`
	PassedCheckCount      int    `json:"passed_check_count"`
	RequiredCheckCount    int    `json:"required_check_count"`
	PendingCheckCount     int    `json:"pending_check_count"`
	BlockedCheckCount     int    `json:"blocked_check_count"`
	PreflightComplete     bool   `json:"preflight_complete"`
	PreflightPassed       bool   `json:"preflight_passed"`
	ReviewReceiptRequired bool   `json:"review_receipt_required"`
	ReviewReceiptRecorded bool   `json:"review_receipt_recorded"`
	RuntimeLaunchApproval bool   `json:"runtime_launch_approval"`
	LaunchAllowed         bool   `json:"launch_allowed"`
	ExecutionStarted      bool   `json:"execution_started"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

type KDEActionReceiptField struct {
	ID       string `json:"id"`
	Source   string `json:"source"`
	Required bool   `json:"required"`
	Recorded bool   `json:"recorded"`
	Summary  string `json:"summary"`
}

func (plan Plan) KDEActionReceiptPreview(actionID string, decision string, fileURIs []string) (KDEActionReceiptPreview, error) {
	if !singleLine(actionID) {
		return KDEActionReceiptPreview{}, errors.New("KDE action receipt preview requires a single-line action id")
	}
	if !singleLine(decision) {
		return KDEActionReceiptPreview{}, errors.New("KDE action receipt preview requires a single-line decision")
	}
	if !supportedExecutionReviewDecision(decision) {
		return KDEActionReceiptPreview{}, fmt.Errorf("unsupported KDE action receipt decision: %s", decision)
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return KDEActionReceiptPreview{}, errors.New("KDE action receipt preview requires single-line identity fields")
		}
	}

	preflight, err := plan.KDEActionPreflightPreview(actionID, decision, fileURIs)
	if err != nil {
		return KDEActionReceiptPreview{}, err
	}

	receiptFields := kdeActionReceiptFields(preflight)
	preview := KDEActionReceiptPreview{
		SchemaVersion:   "xnix.runtime.kde_action_receipt.v1",
		RequestType:     "kde-action-receipt-preview",
		ReceiptType:     "compatibility-center-kde-action-review-receipt",
		ReceiptID:       kdeActionReceiptID(preflight.ApplicationID, actionID, decision),
		ReceiptState:    "preview-only",
		Source:          "kde-action-preflight-preview",
		Desktop:         "KDE Plasma",
		RuntimeMethod:   "RecordKDEActionReviewReceipt",
		ReadMethod:      "GetKDEActionReceiptPreview",
		ApplicationID:   preflight.ApplicationID,
		ApplicationName: preflight.ApplicationName,
		Icon:            preflight.Icon,
		DesktopFile:     preflight.DesktopFile,
		LauncherCommand: preflight.LauncherCommand,
		Review: KDEActionReceiptReviewSummary{
			RequestType:            preflight.Review.RequestType,
			ReviewType:             preflight.Review.ReviewType,
			Decision:               preflight.Review.Decision,
			DecisionAccepted:       preflight.Review.DecisionAccepted,
			UserIntentCaptured:     preflight.Review.UserIntentCaptured,
			DecisionRecorded:       preflight.Review.DecisionRecorded,
			RuntimeApprovalGranted: preflight.Review.RuntimeApprovalGranted,
			ExecutionAllowed:       preflight.Review.ExecutionAllowed,
			ReviewReceiptCreated:   preflight.Review.ReviewReceiptCreated,
			QueueStateChanged:      preflight.Review.QueueStateChanged,
			PermissionGrantCreated: preflight.Review.PermissionGrantCreated,
		},
		Action: preflight.Action,
		Preflight: KDEActionReceiptPreflightSummary{
			RequestType:           preflight.RequestType,
			PreflightType:         preflight.PreflightType,
			ReceiptGateStatus:     kdeActionReceiptGateStatus(preflight),
			CheckCount:            preflight.CheckCount,
			PassedCheckCount:      preflight.PassedCheckCount,
			RequiredCheckCount:    preflight.RequiredCheckCount,
			PendingCheckCount:     preflight.PendingCheckCount,
			BlockedCheckCount:     preflight.BlockedCheckCount,
			PreflightComplete:     preflight.PreflightComplete,
			PreflightPassed:       preflight.PreflightPassed,
			ReviewReceiptRequired: preflight.ReviewReceiptRequired,
			ReviewReceiptRecorded: preflight.ReviewReceiptRecorded,
			RuntimeLaunchApproval: preflight.RuntimeLaunchApproval,
			LaunchAllowed:         preflight.LaunchAllowed,
			ExecutionStarted:      preflight.ExecutionStarted,
			BackendDetailsExposed: preflight.BackendDetailsExposed,
		},
		ReceiptFields:            receiptFields,
		FieldCount:               len(receiptFields),
		RequiredRuntimeGate:      preflight.Action.RuntimeGate,
		NextStep:                 kdeActionReceiptNextStep(decision, preflight.Action.RuntimeGate),
		SessionStatus:            preflight.SessionStatus,
		FileCount:                preflight.FileCount,
		FileURIs:                 preflight.FileURIs,
		RuntimeOwned:             true,
		GoRuntimeBacked:          true,
		KDEPolicyOwner:           false,
		OfficialDesktopOnly:      true,
		CompatibilityCenterCard:  true,
		SafeForAIDiagnostics:     true,
		UserDecisionCaptured:     preflight.UserDecisionCaptured,
		UserDecisionAllowsLaunch: preflight.UserDecisionAllowsLaunch,
		ActionReviewCaptured:     true,
		ActionPreflightCreated:   preflight.ActionPreflightCreated,
		ReceiptPreviewCreated:    true,
		ReceiptRecordable:        true,
		DecisionRecorded:         false,
		ReviewReceiptCreated:     false,
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
		BlockedActions:           []string{"record KDE action receipt from preview state", "treat KDE action receipt as Runtime execution approval", "persist KDE action queue from receipt preview", "create Runtime request objects from receipt preview", "grant desktop resources from receipt preview", "persist compatibility settings from receipt preview", "send desktop notifications from receipt preview", "start compatibility profile from receipt preview", "mutate host root during KDE action receipt preview", "expose raw backend command to desktop shell"},
		UserFacingSettings:       preflight.UserFacingSettings,
		DesktopSafeSummary:       "Compatibility Center can preview the Runtime receipt that would record KDE review intent, but the preview records nothing and still cannot start execution.",
	}
	if err := validateNoBackendTerms(preview, "KDE action receipt preview"); err != nil {
		return KDEActionReceiptPreview{}, err
	}
	return preview, nil
}

func kdeActionReceiptFields(preflight KDEActionPreflightPreview) []KDEActionReceiptField {
	return []KDEActionReceiptField{
		kdeActionReceiptField("application-id", preflight.ApplicationID, "kde-action-preflight-preview", "Runtime application identity must be part of the review receipt."),
		kdeActionReceiptField("action-id", preflight.Action.ID, "kde-action-preflight-preview", "Queued KDE action identity must be part of the review receipt."),
		kdeActionReceiptField("decision", preflight.Review.Decision, "kde-action-review-preview", "User review decision must be part of the review receipt."),
		kdeActionReceiptField("required-runtime-gate", preflight.Action.RuntimeGate, "kde-action-preflight-preview", "Required Runtime gate must remain visible after receipt recording."),
		kdeActionReceiptField("execution-disabled", "true", "kde-action-preflight-preview", "Receipt recording must not imply execution approval."),
	}
}

func kdeActionReceiptField(id string, value string, source string, summary string) KDEActionReceiptField {
	return KDEActionReceiptField{
		ID:       id,
		Source:   source,
		Required: true,
		Recorded: false,
		Summary:  fmt.Sprintf("%s Preview value: %s.", summary, value),
	}
}

func kdeActionReceiptID(applicationID string, actionID string, decision string) string {
	raw := strings.Join([]string{"compat-review", applicationID, actionID, decision}, "-")
	var builder strings.Builder
	lastDash := false
	for _, value := range raw {
		allowed := value >= 'a' && value <= 'z' ||
			value >= 'A' && value <= 'Z' ||
			value >= '0' && value <= '9' ||
			value == '.'
		if allowed {
			builder.WriteRune(value)
			lastDash = false
			continue
		}
		if !lastDash {
			builder.WriteRune('-')
			lastDash = true
		}
	}
	return strings.Trim(builder.String(), "-")
}

func kdeActionReceiptGateStatus(preflight KDEActionPreflightPreview) string {
	for _, gate := range preflight.PreflightChecks {
		if gate.ID == "review-receipt" {
			return gate.Status
		}
	}
	return "pending"
}

func kdeActionReceiptNextStep(decision string, runtimeGate string) string {
	switch decision {
	case "rejected":
		return "Runtime would record the rejection, keep the queue visible, and keep action execution blocked."
	case "deferred":
		return "Runtime would record the deferral and keep the queued KDE action waiting for later review."
	default:
		return fmt.Sprintf("Runtime would record review intent, then wait for %s before any execution.", runtimeGate)
	}
}
