package appidentity

import "errors"

type KDEActionCardDeckPreview struct {
	SchemaVersion            string                  `json:"schema_version"`
	RequestType              string                  `json:"request_type"`
	DeckType                 string                  `json:"deck_type"`
	Source                   string                  `json:"source"`
	Desktop                  string                  `json:"desktop"`
	RuntimeMethod            string                  `json:"runtime_method"`
	ReadMethod               string                  `json:"read_method"`
	ApplicationID            string                  `json:"application_id"`
	ApplicationName          string                  `json:"application_name"`
	Icon                     string                  `json:"icon"`
	DesktopFile              string                  `json:"desktop_file"`
	LauncherCommand          []string                `json:"launcher_command"`
	Queue                    KDEActionCardDeckQueue  `json:"queue"`
	Cards                    []KDEActionCardDeckItem `json:"cards"`
	CardIDs                  []string                `json:"card_ids"`
	PrimaryCardID            string                  `json:"primary_card_id"`
	CardCount                int                     `json:"card_count"`
	WaitingCardCount         int                     `json:"waiting_card_count"`
	DeferredCardCount        int                     `json:"deferred_card_count"`
	RejectedCardCount        int                     `json:"rejected_card_count"`
	NavigationActionCount    int                     `json:"navigation_action_count"`
	DisabledActionCount      int                     `json:"disabled_action_count"`
	FileCount                int                     `json:"file_count"`
	FileURIs                 []string                `json:"file_uris"`
	RuntimeOwned             bool                    `json:"runtime_owned"`
	GoRuntimeBacked          bool                    `json:"go_runtime_backed"`
	KDEPolicyOwner           bool                    `json:"kde_policy_owner"`
	OfficialDesktopOnly      bool                    `json:"official_desktop_only"`
	CompatibilityCenterDeck  bool                    `json:"compatibility_center_deck"`
	SafeForAIDiagnostics     bool                    `json:"safe_for_ai_diagnostics"`
	UserDecisionCaptured     bool                    `json:"user_decision_captured"`
	UserDecisionAllowsLaunch bool                    `json:"user_decision_allows_launch"`
	ActionQueueCreated       bool                    `json:"action_queue_created"`
	ActionQueuePersisted     bool                    `json:"action_queue_persisted"`
	DeckPreviewCreated       bool                    `json:"deck_preview_created"`
	DeckPersisted            bool                    `json:"deck_persisted"`
	CardsPersisted           bool                    `json:"cards_persisted"`
	CardActionsEnabled       bool                    `json:"card_actions_enabled"`
	CardActionsPersisted     bool                    `json:"card_actions_persisted"`
	StatusPersisted          bool                    `json:"status_persisted"`
	DecisionRecorded         bool                    `json:"decision_recorded"`
	ReviewReceiptRecorded    bool                    `json:"review_receipt_recorded"`
	QueueStateChanged        bool                    `json:"queue_state_changed"`
	SettingsPersisted        bool                    `json:"settings_persisted"`
	NotificationsSent        bool                    `json:"notifications_sent"`
	ResourceGrantCreated     bool                    `json:"resource_grant_created"`
	RuntimeLaunchApproval    bool                    `json:"runtime_launch_approval"`
	LaunchAllowed            bool                    `json:"launch_allowed"`
	LaunchEnabled            bool                    `json:"launch_enabled"`
	ExecutionStarted         bool                    `json:"execution_started"`
	BackendProcessStarted    bool                    `json:"backend_process_started"`
	RequestObjectsCreated    bool                    `json:"request_objects_created"`
	PermissionGrantCreated   bool                    `json:"permission_grant_created"`
	HostRootModified         bool                    `json:"host_root_modified"`
	NetworkRequired          bool                    `json:"network_required"`
	BackendDetailsExposed    bool                    `json:"backend_details_exposed"`
	BlockedActions           []string                `json:"blocked_actions"`
	UserFacingSettings       map[string]string       `json:"user_facing_settings"`
	DesktopSafeSummary       string                  `json:"desktop_safe_summary"`
}

type KDEActionCardDeckQueue struct {
	RequestType             string   `json:"request_type"`
	QueueType               string   `json:"queue_type"`
	Source                  string   `json:"source"`
	ActionIDs               []string `json:"action_ids"`
	ActionCount             int      `json:"action_count"`
	PendingActionCount      int      `json:"pending_action_count"`
	UserReviewRequiredCount int      `json:"user_review_required_count"`
	PortalActionCount       int      `json:"portal_action_count"`
	RuntimeGateActionCount  int      `json:"runtime_gate_action_count"`
	ActionQueueCreated      bool     `json:"action_queue_created"`
	ActionQueuePersisted    bool     `json:"action_queue_persisted"`
	RuntimeLaunchApproval   bool     `json:"runtime_launch_approval"`
	LaunchAllowed           bool     `json:"launch_allowed"`
	ExecutionStarted        bool     `json:"execution_started"`
	BackendDetailsExposed   bool     `json:"backend_details_exposed"`
}

type KDEActionCardDeckItem struct {
	CardID                string                   `json:"card_id"`
	ActionID              string                   `json:"action_id"`
	EntryPointID          string                   `json:"entry_point_id"`
	EntryPointLabel       string                   `json:"entry_point_label"`
	KDEComponent          string                   `json:"kde_component"`
	CardState             string                   `json:"card_state"`
	RequiredRuntimeGate   string                   `json:"required_runtime_gate"`
	UserReviewRequired    bool                     `json:"user_review_required"`
	RequiresPortal        bool                     `json:"requires_portal"`
	RequiresRuntimeGate   bool                     `json:"requires_runtime_gate"`
	Card                  KDEActionCardVisualState `json:"card"`
	PrimaryAction         KDEActionCardAction      `json:"primary_action"`
	SecondaryActions      []KDEActionCardAction    `json:"secondary_actions"`
	DisabledActions       []KDEActionCardAction    `json:"disabled_actions"`
	DetailRows            []KDEActionCardDetail    `json:"detail_rows"`
	NavigationActionCount int                      `json:"navigation_action_count"`
	DisabledActionCount   int                      `json:"disabled_action_count"`
	CardPreviewCreated    bool                     `json:"card_preview_created"`
	CardPersisted         bool                     `json:"card_persisted"`
	ExecutionEnabled      bool                     `json:"execution_enabled"`
	RequestObjectsCreated bool                     `json:"request_objects_created"`
	ResourceGrantCreated  bool                     `json:"resource_grant_created"`
	NotificationsSent     bool                     `json:"notifications_sent"`
	MutatesRuntime        bool                     `json:"mutates_runtime"`
	StartsProgram         bool                     `json:"starts_program"`
	HostRootModified      bool                     `json:"host_root_modified"`
	BackendDetailsExposed bool                     `json:"backend_details_exposed"`
}

func (plan Plan) KDEActionCardDeckPreview(decision string, fileURIs []string) (KDEActionCardDeckPreview, error) {
	if !singleLine(decision) {
		return KDEActionCardDeckPreview{}, errors.New("KDE action card deck preview requires a single-line decision")
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return KDEActionCardDeckPreview{}, errors.New("KDE action card deck preview requires single-line identity fields")
		}
	}

	queue, err := plan.KDEActionQueuePreview(decision, fileURIs)
	if err != nil {
		return KDEActionCardDeckPreview{}, err
	}

	cards := make([]KDEActionCardDeckItem, 0, len(queue.Actions))
	for _, action := range queue.Actions {
		card, err := plan.KDEActionCardPreview(action.ID, decision, fileURIs)
		if err != nil {
			return KDEActionCardDeckPreview{}, err
		}
		cards = append(cards, kdeActionCardDeckItem(card))
	}
	cardIDs, waitingCount, deferredCount, rejectedCount, navigationCount, disabledCount := summarizeKDEActionCardDeck(cards)

	preview := KDEActionCardDeckPreview{
		SchemaVersion:   "xnix.runtime.kde_action_card_deck.v1",
		RequestType:     "kde-action-card-deck-preview",
		DeckType:        "compatibility-center-kde-action-card-deck",
		Source:          "kde-action-card-preview",
		Desktop:         "KDE Plasma",
		RuntimeMethod:   "GetKDEActionCardDeck",
		ReadMethod:      "GetKDEActionCardDeckPreview",
		ApplicationID:   queue.ApplicationID,
		ApplicationName: queue.ApplicationName,
		Icon:            queue.Icon,
		DesktopFile:     queue.DesktopFile,
		LauncherCommand: queue.LauncherCommand,
		Queue: KDEActionCardDeckQueue{
			RequestType:             queue.RequestType,
			QueueType:               queue.QueueType,
			Source:                  queue.Source,
			ActionIDs:               queue.ActionIDs,
			ActionCount:             queue.ActionCount,
			PendingActionCount:      queue.PendingActionCount,
			UserReviewRequiredCount: queue.UserReviewRequiredCount,
			PortalActionCount:       queue.PortalActionCount,
			RuntimeGateActionCount:  queue.RuntimeGateActionCount,
			ActionQueueCreated:      queue.ActionQueueCreated,
			ActionQueuePersisted:    queue.ActionQueuePersisted,
			RuntimeLaunchApproval:   queue.RuntimeLaunchApproval,
			LaunchAllowed:           queue.LaunchAllowed,
			ExecutionStarted:        queue.ExecutionStarted,
			BackendDetailsExposed:   queue.BackendDetailsExposed,
		},
		Cards:                    cards,
		CardIDs:                  cardIDs,
		PrimaryCardID:            primaryKDEActionCardID(cardIDs),
		CardCount:                len(cards),
		WaitingCardCount:         waitingCount,
		DeferredCardCount:        deferredCount,
		RejectedCardCount:        rejectedCount,
		NavigationActionCount:    navigationCount,
		DisabledActionCount:      disabledCount,
		FileCount:                queue.FileCount,
		FileURIs:                 queue.FileURIs,
		RuntimeOwned:             true,
		GoRuntimeBacked:          true,
		KDEPolicyOwner:           false,
		OfficialDesktopOnly:      true,
		CompatibilityCenterDeck:  true,
		SafeForAIDiagnostics:     true,
		UserDecisionCaptured:     queue.UserDecisionCaptured,
		UserDecisionAllowsLaunch: queue.UserDecisionAllowsLaunch,
		ActionQueueCreated:       queue.ActionQueueCreated,
		ActionQueuePersisted:     false,
		DeckPreviewCreated:       true,
		DeckPersisted:            false,
		CardsPersisted:           false,
		CardActionsEnabled:       false,
		CardActionsPersisted:     false,
		StatusPersisted:          false,
		DecisionRecorded:         false,
		ReviewReceiptRecorded:    false,
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
		BlockedActions:           []string{"persist KDE action card deck from preview state", "enable card actions from deck preview", "record KDE action receipts from deck preview", "persist KDE action queue from deck preview", "create Runtime request objects from deck preview", "grant desktop resources from deck preview", "persist compatibility settings from deck preview", "send desktop notifications from deck preview", "start compatibility profile from deck preview", "mutate host root during KDE action card deck preview", "expose raw backend command to desktop shell"},
		UserFacingSettings:       queue.UserFacingSettings,
		DesktopSafeSummary:       "Compatibility Center can render a deck of KDE action cards from Runtime previews, but the deck is read-only and cannot approve, grant, persist, notify, or start execution.",
	}
	if err := validateNoBackendTerms(preview, "KDE action card deck preview"); err != nil {
		return KDEActionCardDeckPreview{}, err
	}
	return preview, nil
}

func kdeActionCardDeckItem(card KDEActionCardPreview) KDEActionCardDeckItem {
	return KDEActionCardDeckItem{
		CardID:                kdeActionCardDeckID(card.ApplicationID, card.Action.ID),
		ActionID:              card.Action.ID,
		EntryPointID:          card.Action.EntryPointID,
		EntryPointLabel:       card.Action.EntryPointLabel,
		KDEComponent:          card.Action.KDEComponent,
		CardState:             card.CardState,
		RequiredRuntimeGate:   card.RequiredRuntimeGate,
		UserReviewRequired:    card.Action.UserReviewRequired,
		RequiresPortal:        card.Action.RequiresPortal,
		RequiresRuntimeGate:   card.Action.RequiresRuntimeGate,
		Card:                  card.Card,
		PrimaryAction:         card.Card.PrimaryAction,
		SecondaryActions:      card.Card.SecondaryActions,
		DisabledActions:       card.Card.DisabledActions,
		DetailRows:            card.Card.DetailRows,
		NavigationActionCount: kdeActionCardNavigationActionCount(card.Card),
		DisabledActionCount:   len(card.Card.DisabledActions),
		CardPreviewCreated:    card.CardPreviewCreated,
		CardPersisted:         card.CardPersisted,
		ExecutionEnabled:      false,
		RequestObjectsCreated: false,
		ResourceGrantCreated:  false,
		NotificationsSent:     false,
		MutatesRuntime:        false,
		StartsProgram:         false,
		HostRootModified:      false,
		BackendDetailsExposed: false,
	}
}

func kdeActionCardDeckID(applicationID string, actionID string) string {
	return applicationID + ":" + actionID + ":card"
}

func kdeActionCardNavigationActionCount(card KDEActionCardVisualState) int {
	count := 0
	if card.PrimaryAction.NavigationOnly {
		count++
	}
	for _, action := range card.SecondaryActions {
		if action.NavigationOnly {
			count++
		}
	}
	return count
}

func summarizeKDEActionCardDeck(cards []KDEActionCardDeckItem) ([]string, int, int, int, int, int) {
	cardIDs := make([]string, 0, len(cards))
	waitingCount := 0
	deferredCount := 0
	rejectedCount := 0
	navigationCount := 0
	disabledCount := 0
	for _, card := range cards {
		cardIDs = append(cardIDs, card.CardID)
		switch card.CardState {
		case "action-deferred":
			deferredCount++
		case "action-rejected":
			rejectedCount++
		default:
			waitingCount++
		}
		navigationCount += card.NavigationActionCount
		disabledCount += card.DisabledActionCount
	}
	return cardIDs, waitingCount, deferredCount, rejectedCount, navigationCount, disabledCount
}

func primaryKDEActionCardID(cardIDs []string) string {
	if len(cardIDs) == 0 {
		return ""
	}
	return cardIDs[0]
}
