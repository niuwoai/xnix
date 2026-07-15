package appidentity

import "errors"

type KDEActionCardPreview struct {
	SchemaVersion            string                           `json:"schema_version"`
	RequestType              string                           `json:"request_type"`
	CardType                 string                           `json:"card_type"`
	CardState                string                           `json:"card_state"`
	Source                   string                           `json:"source"`
	Desktop                  string                           `json:"desktop"`
	RuntimeMethod            string                           `json:"runtime_method"`
	ReadMethod               string                           `json:"read_method"`
	AIAnalysis               *KDEAIAnalysisLink               `json:"ai_analysis,omitempty"`
	ApplicationID            string                           `json:"application_id"`
	ApplicationName          string                           `json:"application_name"`
	Icon                     string                           `json:"icon"`
	DesktopFile              string                           `json:"desktop_file"`
	LauncherCommand          []string                         `json:"launcher_command"`
	Status                   KDEActionCardStatusSummary       `json:"status"`
	Card                     KDEActionCardVisualState         `json:"card"`
	Action                   KDEActionReviewActionSummary     `json:"action"`
	Preflight                KDEActionReceiptPreflightSummary `json:"preflight"`
	Receipt                  KDEActionStatusReceiptSummary    `json:"receipt"`
	RequiredRuntimeGate      string                           `json:"required_runtime_gate"`
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
	CardPreviewCreated       bool                             `json:"card_preview_created"`
	CardPersisted            bool                             `json:"card_persisted"`
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

type KDEActionCardStatusSummary struct {
	RequestType              string `json:"request_type"`
	StatusType               string `json:"status_type"`
	StatusState              string `json:"status_state"`
	CompatibilityCenterState string `json:"compatibility_center_state"`
	NotificationIntent       string `json:"notification_intent"`
	NextStep                 string `json:"next_step"`
	StatusPreviewCreated     bool   `json:"status_preview_created"`
	StatusPersisted          bool   `json:"status_persisted"`
	RuntimeLaunchApproval    bool   `json:"runtime_launch_approval"`
	LaunchAllowed            bool   `json:"launch_allowed"`
	ExecutionStarted         bool   `json:"execution_started"`
	BackendDetailsExposed    bool   `json:"backend_details_exposed"`
}

type KDEActionCardVisualState struct {
	Title                 string                `json:"title"`
	Subtitle              string                `json:"subtitle"`
	Badge                 string                `json:"badge"`
	BadgeTone             string                `json:"badge_tone"`
	PrimaryAction         KDEActionCardAction   `json:"primary_action"`
	AIAnalysisAction      *KDEActionCardAction  `json:"ai_analysis_action,omitempty"`
	SecondaryActions      []KDEActionCardAction `json:"secondary_actions"`
	DisabledActions       []KDEActionCardAction `json:"disabled_actions"`
	DetailRows            []KDEActionCardDetail `json:"detail_rows"`
	Footer                string                `json:"footer"`
	UserFacingMode        string                `json:"user_facing_mode"`
	UserFacingAccess      string                `json:"user_facing_access"`
	BackendDetailsExposed bool                  `json:"backend_details_exposed"`
}

type KDEActionCardAction struct {
	ID             string `json:"id"`
	Label          string `json:"label"`
	Target         string `json:"target"`
	Enabled        bool   `json:"enabled"`
	NavigationOnly bool   `json:"navigation_only"`
	MutatesRuntime bool   `json:"mutates_runtime"`
	StartsProgram  bool   `json:"starts_program"`
}

type KDEActionCardDetail struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func (plan Plan) KDEActionCardPreview(actionID string, decision string, fileURIs []string) (KDEActionCardPreview, error) {
	if !singleLine(actionID) {
		return KDEActionCardPreview{}, errors.New("KDE action card preview requires a single-line action id")
	}
	if !singleLine(decision) {
		return KDEActionCardPreview{}, errors.New("KDE action card preview requires a single-line decision")
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return KDEActionCardPreview{}, errors.New("KDE action card preview requires single-line identity fields")
		}
	}

	status, err := plan.KDEActionStatusPreview(actionID, decision, fileURIs)
	if err != nil {
		return KDEActionCardPreview{}, err
	}

	preview := KDEActionCardPreview{
		SchemaVersion:   "xnix.runtime.kde_action_card.v1",
		RequestType:     "kde-action-card-preview",
		CardType:        "compatibility-center-kde-action-card",
		CardState:       status.CompatibilityCenterState,
		Source:          "kde-action-status-preview",
		Desktop:         "KDE Plasma",
		RuntimeMethod:   "GetKDEActionCard",
		ReadMethod:      "GetKDEActionCardPreview",
		AIAnalysis:      kdeAIAnalysisLinkForEntryPoint(status.Action.EntryPointID),
		ApplicationID:   status.ApplicationID,
		ApplicationName: status.ApplicationName,
		Icon:            status.Icon,
		DesktopFile:     status.DesktopFile,
		LauncherCommand: status.LauncherCommand,
		Status: KDEActionCardStatusSummary{
			RequestType:              status.RequestType,
			StatusType:               status.StatusType,
			StatusState:              status.StatusState,
			CompatibilityCenterState: status.CompatibilityCenterState,
			NotificationIntent:       status.NotificationIntent,
			NextStep:                 status.NextStep,
			StatusPreviewCreated:     status.StatusPreviewCreated,
			StatusPersisted:          status.StatusPersisted,
			RuntimeLaunchApproval:    status.RuntimeLaunchApproval,
			LaunchAllowed:            status.LaunchAllowed,
			ExecutionStarted:         status.ExecutionStarted,
			BackendDetailsExposed:    status.BackendDetailsExposed,
		},
		Card:                     kdeActionCardVisualState(status),
		Action:                   status.Action,
		Preflight:                status.Preflight,
		Receipt:                  status.Receipt,
		RequiredRuntimeGate:      status.RequiredRuntimeGate,
		FileCount:                status.FileCount,
		FileURIs:                 status.FileURIs,
		RuntimeOwned:             true,
		GoRuntimeBacked:          true,
		KDEPolicyOwner:           false,
		OfficialDesktopOnly:      true,
		CompatibilityCenterCard:  true,
		SafeForAIDiagnostics:     true,
		UserDecisionCaptured:     status.UserDecisionCaptured,
		UserDecisionAllowsLaunch: status.UserDecisionAllowsLaunch,
		ActionReviewCaptured:     status.ActionReviewCaptured,
		ActionPreflightCreated:   status.ActionPreflightCreated,
		ReceiptPreviewCreated:    status.ReceiptPreviewCreated,
		StatusPreviewCreated:     status.StatusPreviewCreated,
		CardPreviewCreated:       true,
		CardPersisted:            false,
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
		BlockedActions:           []string{"persist KDE action card from preview state", "treat KDE action card as Runtime execution approval", "record KDE action receipt from card preview", "persist KDE action queue from card preview", "create Runtime request objects from card preview", "grant desktop resources from card preview", "persist compatibility settings from card preview", "send desktop notifications from card preview", "start compatibility profile from card preview", "mutate host root during KDE action card preview", "expose raw backend command to desktop shell"},
		UserFacingSettings:       status.UserFacingSettings,
		DesktopSafeSummary:       "Compatibility Center can render a KDE action card from Runtime status, but the card is read-only and cannot approve, grant, persist, notify, or start execution.",
	}
	if err := validateNoBackendTerms(preview, "KDE action card preview"); err != nil {
		return KDEActionCardPreview{}, err
	}
	return preview, nil
}

func kdeActionCardVisualState(status KDEActionStatusPreview) KDEActionCardVisualState {
	return KDEActionCardVisualState{
		Title:                 status.ApplicationName,
		Subtitle:              kdeActionCardSubtitle(status.StatusState),
		Badge:                 status.UserVisibleState.Badge,
		BadgeTone:             kdeActionCardBadgeTone(status.StatusState),
		PrimaryAction:         kdeActionCardPrimaryAction(status.StatusState),
		AIAnalysisAction:      kdeActionCardAIAnalysisAction(status.Action.EntryPointID),
		SecondaryActions:      kdeActionCardSecondaryActions(status.StatusState),
		DisabledActions:       kdeActionCardDisabledActions(),
		DetailRows:            kdeActionCardDetailRows(status),
		Footer:                kdeActionCardFooter(status.StatusState),
		UserFacingMode:        status.UserVisibleState.UserFacingMode,
		UserFacingAccess:      status.UserVisibleState.UserFacingAccess,
		BackendDetailsExposed: false,
	}
}

func kdeActionCardAIAnalysisAction(entryPointID string) *KDEActionCardAction {
	if entryPointID != "file-manager" {
		return nil
	}
	action := kdeActionCardAction("preview-dolphin-ai-analysis", "Preview AI-safe file review", "dolphin-ai-analysis-preview", true)
	return &action
}

func kdeActionCardSubtitle(state string) string {
	switch state {
	case "review-rejected":
		return "Action rejected by user review"
	case "review-deferred":
		return "Action deferred for later review"
	default:
		return "Waiting for desktop access review"
	}
}

func kdeActionCardBadgeTone(state string) string {
	switch state {
	case "review-rejected":
		return "critical"
	case "review-deferred":
		return "neutral"
	default:
		return "warning"
	}
}

func kdeActionCardPrimaryAction(state string) KDEActionCardAction {
	switch state {
	case "review-rejected":
		return kdeActionCardAction("show-guidance", "Show guidance", "compatibility-center-guidance", true)
	case "review-deferred":
		return kdeActionCardAction("resume-review", "Resume review", "compatibility-center-review", true)
	default:
		return kdeActionCardAction("review-required-gates", "Review required gates", "compatibility-center-gates", true)
	}
}

func kdeActionCardSecondaryActions(state string) []KDEActionCardAction {
	actions := []KDEActionCardAction{
		kdeActionCardAction("open-settings", "Open settings", "compatibility-settings", true),
		kdeActionCardAction("open-compatibility-center", "Open Compatibility Center", "compatibility-center", true),
	}
	if state == "waiting-for-runtime-gates" {
		actions = append(actions, kdeActionCardAction("show-status-details", "Show status details", "compatibility-center-status", true))
	}
	return actions
}

func kdeActionCardDisabledActions() []KDEActionCardAction {
	return []KDEActionCardAction{
		kdeActionCardAction("start-application", "Start application", "runtime-execution", false),
		kdeActionCardAction("record-review-receipt", "Record review receipt", "runtime-receipt", false),
		kdeActionCardAction("grant-resource-access", "Grant resource access", "runtime-resource-grant", false),
	}
}

func kdeActionCardAction(id string, label string, target string, enabled bool) KDEActionCardAction {
	return KDEActionCardAction{
		ID:             id,
		Label:          label,
		Target:         target,
		Enabled:        enabled,
		NavigationOnly: enabled,
		MutatesRuntime: false,
		StartsProgram:  false,
	}
}

func kdeActionCardDetailRows(status KDEActionStatusPreview) []KDEActionCardDetail {
	rows := []KDEActionCardDetail{
		{Label: "Application", Value: status.ApplicationName},
		{Label: "KDE surface", Value: status.Action.KDEComponent},
		{Label: "Current state", Value: status.UserVisibleState.ActionStateLabel},
		{Label: "Next step", Value: kdeActionCardNextStep(status.StatusState)},
	}
	if link := kdeAIAnalysisLinkForEntryPoint(status.Action.EntryPointID); link != nil {
		rows = append(rows, KDEActionCardDetail{Label: "AI analysis disclosure", Value: link.Disclosure})
	}
	return rows
}

func kdeActionCardNextStep(state string) string {
	switch state {
	case "review-rejected":
		return "Keep execution blocked and show safe guidance"
	case "review-deferred":
		return "Wait for the user to resume review"
	default:
		return "Wait for Runtime gates before execution"
	}
}

func kdeActionCardFooter(state string) string {
	switch state {
	case "review-rejected":
		return "No compatibility action will run from this card."
	case "review-deferred":
		return "This card keeps the review visible without changing system state."
	default:
		return "This card is a read-only preview; execution remains blocked."
	}
}
