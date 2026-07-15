package appidentity

import "errors"

type KDECenterPagePreview struct {
	SchemaVersion              string                    `json:"schema_version"`
	RequestType                string                    `json:"request_type"`
	PageType                   string                    `json:"page_type"`
	Source                     string                    `json:"source"`
	Desktop                    string                    `json:"desktop"`
	RuntimeMethod              string                    `json:"runtime_method"`
	ReadMethod                 string                    `json:"read_method"`
	ApplicationID              string                    `json:"application_id"`
	ApplicationName            string                    `json:"application_name"`
	Icon                       string                    `json:"icon"`
	DesktopFile                string                    `json:"desktop_file"`
	LauncherCommand            []string                  `json:"launcher_command"`
	Header                     KDECenterPageHeader       `json:"header"`
	ApplicationSummary         KDECenterPageApplication  `json:"application_summary"`
	ActionDeck                 KDECenterPageActionDeck   `json:"action_deck"`
	SettingsSnapshot           KDECenterPageSettings     `json:"settings_snapshot"`
	Navigation                 []KDECenterPageNavigation `json:"navigation"`
	NavigationCount            int                       `json:"navigation_count"`
	PrimaryNavigationTarget    string                    `json:"primary_navigation_target"`
	RuntimeOwned               bool                      `json:"runtime_owned"`
	GoRuntimeBacked            bool                      `json:"go_runtime_backed"`
	KDEPolicyOwner             bool                      `json:"kde_policy_owner"`
	OfficialDesktopOnly        bool                      `json:"official_desktop_only"`
	UserVisible                bool                      `json:"user_visible"`
	SafeForAIDiagnostics       bool                      `json:"safe_for_ai_diagnostics"`
	UserDecisionCaptured       bool                      `json:"user_decision_captured"`
	UserDecisionAllowsLaunch   bool                      `json:"user_decision_allows_launch"`
	PagePreviewCreated         bool                      `json:"page_preview_created"`
	PagePersisted              bool                      `json:"page_persisted"`
	DeckPersisted              bool                      `json:"deck_persisted"`
	CardsPersisted             bool                      `json:"cards_persisted"`
	CardActionsEnabled         bool                      `json:"card_actions_enabled"`
	SettingsPersisted          bool                      `json:"settings_persisted"`
	SettingsPersistenceEnabled bool                      `json:"settings_persistence_enabled"`
	NotificationsSent          bool                      `json:"notifications_sent"`
	ResourceGrantCreated       bool                      `json:"resource_grant_created"`
	RuntimeLaunchApproval      bool                      `json:"runtime_launch_approval"`
	LaunchAllowed              bool                      `json:"launch_allowed"`
	LaunchEnabled              bool                      `json:"launch_enabled"`
	ExecutionStarted           bool                      `json:"execution_started"`
	BackendProcessStarted      bool                      `json:"backend_process_started"`
	RequestObjectsCreated      bool                      `json:"request_objects_created"`
	PermissionGrantCreated     bool                      `json:"permission_grant_created"`
	HostRootModified           bool                      `json:"host_root_modified"`
	NetworkRequired            bool                      `json:"network_required"`
	BackendDetailsExposed      bool                      `json:"backend_details_exposed"`
	BlockedActions             []string                  `json:"blocked_actions"`
	UserFacingSettings         map[string]string         `json:"user_facing_settings"`
	DesktopSafeSummary         string                    `json:"desktop_safe_summary"`
}

type KDECenterPageHeader struct {
	Title                 string `json:"title"`
	Subtitle              string `json:"subtitle"`
	Badge                 string `json:"badge"`
	BadgeTone             string `json:"badge_tone"`
	PrimaryActionLabel    string `json:"primary_action_label"`
	PrimaryActionTarget   string `json:"primary_action_target"`
	PrimaryActionEnabled  bool   `json:"primary_action_enabled"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

type KDECenterPageApplication struct {
	ApplicationID              string   `json:"application_id"`
	DisplayName                string   `json:"display_name"`
	CompatibilityState         string   `json:"compatibility_state"`
	CompatibilityLabel         string   `json:"compatibility_label"`
	DiagnosticsState           string   `json:"diagnostics_state"`
	RuntimeMode                string   `json:"runtime_mode"`
	SupportedExtensions        []string `json:"supported_extensions"`
	KnownIssueCount            int      `json:"known_issue_count"`
	RepairRecordState          string   `json:"repair_record_state"`
	RepairRecordCount          int      `json:"repair_record_count"`
	ActionExecutionEnabled     bool     `json:"action_execution_enabled"`
	RepairExecutionEnabled     bool     `json:"repair_execution_enabled"`
	BackendLaunchEnabled       bool     `json:"backend_launch_enabled"`
	SettingsPersistenceEnabled bool     `json:"settings_persistence_enabled"`
	HostRootModified           bool     `json:"host_root_modified"`
	BackendDetailsExposed      bool     `json:"backend_details_exposed"`
	Summary                    string   `json:"summary"`
}

type KDECenterPageActionDeck struct {
	RequestType           string   `json:"request_type"`
	DeckType              string   `json:"deck_type"`
	CardCount             int      `json:"card_count"`
	WaitingCardCount      int      `json:"waiting_card_count"`
	DeferredCardCount     int      `json:"deferred_card_count"`
	RejectedCardCount     int      `json:"rejected_card_count"`
	NavigationActionCount int      `json:"navigation_action_count"`
	DisabledActionCount   int      `json:"disabled_action_count"`
	PrimaryCardID         string   `json:"primary_card_id"`
	CardIDs               []string `json:"card_ids"`
	ActionQueueCreated    bool     `json:"action_queue_created"`
	ActionQueuePersisted  bool     `json:"action_queue_persisted"`
	DeckPreviewCreated    bool     `json:"deck_preview_created"`
	DeckPersisted         bool     `json:"deck_persisted"`
	CardsPersisted        bool     `json:"cards_persisted"`
	CardActionsEnabled    bool     `json:"card_actions_enabled"`
	RuntimeLaunchApproval bool     `json:"runtime_launch_approval"`
	LaunchAllowed         bool     `json:"launch_allowed"`
	ExecutionStarted      bool     `json:"execution_started"`
	BackendDetailsExposed bool     `json:"backend_details_exposed"`
}

type KDECenterPageSettings struct {
	RequestType                string            `json:"request_type"`
	SettingsState              string            `json:"settings_state"`
	SectionCount               int               `json:"section_count"`
	Sections                   []SettingsSection `json:"sections"`
	UserFacingSettings         map[string]string `json:"user_facing_settings"`
	SettingsPersisted          bool              `json:"settings_persisted"`
	SettingsPersistenceEnabled bool              `json:"settings_persistence_enabled"`
	HostRootModified           bool              `json:"host_root_modified"`
	BackendDetailsExposed      bool              `json:"backend_details_exposed"`
}

type KDECenterPageNavigation struct {
	ID             string `json:"id"`
	Label          string `json:"label"`
	Target         string `json:"target"`
	Enabled        bool   `json:"enabled"`
	NavigationOnly bool   `json:"navigation_only"`
	MutatesRuntime bool   `json:"mutates_runtime"`
	StartsProgram  bool   `json:"starts_program"`
}

type KDECenterPageSectionsPreview struct {
	SchemaVersion              string                 `json:"schema_version"`
	RequestType                string                 `json:"request_type"`
	PageType                   string                 `json:"page_type"`
	Source                     string                 `json:"source"`
	Desktop                    string                 `json:"desktop"`
	RuntimeMethod              string                 `json:"runtime_method"`
	ReadMethod                 string                 `json:"read_method"`
	ApplicationID              string                 `json:"application_id"`
	ApplicationName            string                 `json:"application_name"`
	SectionCount               int                    `json:"section_count"`
	ReadOnlySectionCount       int                    `json:"read_only_section_count"`
	NavigationOnlySectionCount int                    `json:"navigation_only_section_count"`
	ExecutableSectionCount     int                    `json:"executable_section_count"`
	PrimarySectionID           string                 `json:"primary_section_id"`
	Sections                   []KDECenterPageSection `json:"sections"`
	RuntimeOwned               bool                   `json:"runtime_owned"`
	GoRuntimeBacked            bool                   `json:"go_runtime_backed"`
	KDEPolicyOwner             bool                   `json:"kde_policy_owner"`
	OfficialDesktopOnly        bool                   `json:"official_desktop_only"`
	UserVisible                bool                   `json:"user_visible"`
	SafeForAIDiagnostics       bool                   `json:"safe_for_ai_diagnostics"`
	UserDecisionCaptured       bool                   `json:"user_decision_captured"`
	UserDecisionAllowsLaunch   bool                   `json:"user_decision_allows_launch"`
	SectionsPreviewCreated     bool                   `json:"sections_preview_created"`
	SectionsPersisted          bool                   `json:"sections_persisted"`
	SectionActionsEnabled      bool                   `json:"section_actions_enabled"`
	SettingsPersisted          bool                   `json:"settings_persisted"`
	SettingsPersistenceEnabled bool                   `json:"settings_persistence_enabled"`
	NotificationsSent          bool                   `json:"notifications_sent"`
	ResourceGrantCreated       bool                   `json:"resource_grant_created"`
	RuntimeLaunchApproval      bool                   `json:"runtime_launch_approval"`
	LaunchEnabled              bool                   `json:"launch_enabled"`
	ExecutionStarted           bool                   `json:"execution_started"`
	RequestObjectsCreated      bool                   `json:"request_objects_created"`
	PermissionGrantCreated     bool                   `json:"permission_grant_created"`
	HostRootModified           bool                   `json:"host_root_modified"`
	NetworkRequired            bool                   `json:"network_required"`
	BackendDetailsExposed      bool                   `json:"backend_details_exposed"`
	BlockedActions             []string               `json:"blocked_actions"`
	DesktopSafeSummary         string                 `json:"desktop_safe_summary"`
}

type KDECenterPageSection struct {
	ID                    string `json:"id"`
	Label                 string `json:"label"`
	Target                string `json:"target"`
	RuntimeMethod         string `json:"runtime_method"`
	ReadModel             string `json:"read_model"`
	State                 string `json:"state"`
	NavigationOnly        bool   `json:"navigation_only"`
	ReadOnly              bool   `json:"read_only"`
	MutatesRuntime        bool   `json:"mutates_runtime"`
	StartsProgram         bool   `json:"starts_program"`
	SettingsPersisted     bool   `json:"settings_persisted"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
	Summary               string `json:"summary"`
}

func NewKDECenterPagePreview(recipe Recipe, provenance Provenance, decision string, fileURIs []string) (KDECenterPagePreview, error) {
	if !singleLine(decision) {
		return KDECenterPagePreview{}, errors.New("KDE center page preview requires a single-line decision")
	}

	center, err := NewCompatibilityCenterPreview([]Recipe{recipe}, provenance)
	if err != nil {
		return KDECenterPagePreview{}, err
	}
	if len(center.Applications) != 1 {
		return KDECenterPagePreview{}, errors.New("KDE center page preview requires exactly one application summary")
	}

	plan, err := NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return KDECenterPagePreview{}, err
	}
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return KDECenterPagePreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return KDECenterPagePreview{}, errors.New("KDE center page preview requires single-line identity fields")
		}
	}

	deck, err := plan.KDEActionCardDeckPreview(decision, fileURIs)
	if err != nil {
		return KDECenterPagePreview{}, err
	}
	settings, err := plan.SettingsPreview()
	if err != nil {
		return KDECenterPagePreview{}, err
	}

	application := center.Applications[0]
	navigation := kdeCenterPageNavigation()
	preview := KDECenterPagePreview{
		SchemaVersion:   "xnix.runtime.kde_center_page.v1",
		RequestType:     "kde-center-page-preview",
		PageType:        "compatibility-center-application-page",
		Source:          "compatibility-center-preview+kde-action-card-deck-preview+settings-preview",
		Desktop:         "KDE Plasma",
		RuntimeMethod:   "GetKDECenterPage",
		ReadMethod:      "GetKDECenterPagePreview",
		ApplicationID:   plan.ApplicationID,
		ApplicationName: plan.DisplayName,
		Icon:            plan.Icon,
		DesktopFile:     plan.DesktopFile,
		LauncherCommand: deck.LauncherCommand,
		Header: KDECenterPageHeader{
			Title:                 plan.DisplayName,
			Subtitle:              "Compatibility Center application details",
			Badge:                 kdeCenterPageBadge(deck),
			BadgeTone:             kdeCenterPageBadgeTone(deck),
			PrimaryActionLabel:    "Review required gates",
			PrimaryActionTarget:   "compatibility-center-gates",
			PrimaryActionEnabled:  true,
			BackendDetailsExposed: false,
		},
		ApplicationSummary: KDECenterPageApplication{
			ApplicationID:              application.ApplicationID,
			DisplayName:                application.DisplayName,
			CompatibilityState:         application.CompatibilityState,
			CompatibilityLabel:         application.CompatibilityLabel,
			DiagnosticsState:           application.DiagnosticsState,
			RuntimeMode:                application.RuntimeMode,
			SupportedExtensions:        application.SupportedExtensions,
			KnownIssueCount:            application.KnownIssueCount,
			RepairRecordState:          application.RepairRecordState,
			RepairRecordCount:          application.RepairRecordCount,
			ActionExecutionEnabled:     application.ActionExecutionEnabled,
			RepairExecutionEnabled:     application.RepairExecutionEnabled,
			BackendLaunchEnabled:       application.BackendLaunchEnabled,
			SettingsPersistenceEnabled: application.SettingsPersistenceEnabled,
			HostRootModified:           application.HostRootModified,
			BackendDetailsExposed:      application.BackendDetailsExposed,
			Summary:                    application.Summary,
		},
		ActionDeck: KDECenterPageActionDeck{
			RequestType:           deck.RequestType,
			DeckType:              deck.DeckType,
			CardCount:             deck.CardCount,
			WaitingCardCount:      deck.WaitingCardCount,
			DeferredCardCount:     deck.DeferredCardCount,
			RejectedCardCount:     deck.RejectedCardCount,
			NavigationActionCount: deck.NavigationActionCount,
			DisabledActionCount:   deck.DisabledActionCount,
			PrimaryCardID:         deck.PrimaryCardID,
			CardIDs:               deck.CardIDs,
			ActionQueueCreated:    deck.ActionQueueCreated,
			ActionQueuePersisted:  deck.ActionQueuePersisted,
			DeckPreviewCreated:    deck.DeckPreviewCreated,
			DeckPersisted:         deck.DeckPersisted,
			CardsPersisted:        deck.CardsPersisted,
			CardActionsEnabled:    deck.CardActionsEnabled,
			RuntimeLaunchApproval: deck.RuntimeLaunchApproval,
			LaunchAllowed:         deck.LaunchAllowed,
			ExecutionStarted:      deck.ExecutionStarted,
			BackendDetailsExposed: deck.BackendDetailsExposed,
		},
		SettingsSnapshot: KDECenterPageSettings{
			RequestType:                settings.RequestType,
			SettingsState:              settings.SettingsState,
			SectionCount:               settings.SectionCount,
			Sections:                   settings.Sections,
			UserFacingSettings:         settings.UserFacingSettings,
			SettingsPersisted:          settings.SettingsPersisted,
			SettingsPersistenceEnabled: settings.SettingsPersistenceEnabled,
			HostRootModified:           settings.HostRootModified,
			BackendDetailsExposed:      settings.BackendDetailsExposed,
		},
		Navigation:                 navigation,
		NavigationCount:            len(navigation),
		PrimaryNavigationTarget:    "compatibility-center-gates",
		RuntimeOwned:               true,
		GoRuntimeBacked:            true,
		KDEPolicyOwner:             false,
		OfficialDesktopOnly:        true,
		UserVisible:                true,
		SafeForAIDiagnostics:       true,
		UserDecisionCaptured:       deck.UserDecisionCaptured,
		UserDecisionAllowsLaunch:   deck.UserDecisionAllowsLaunch,
		PagePreviewCreated:         true,
		PagePersisted:              false,
		DeckPersisted:              false,
		CardsPersisted:             false,
		CardActionsEnabled:         false,
		SettingsPersisted:          false,
		SettingsPersistenceEnabled: false,
		NotificationsSent:          false,
		ResourceGrantCreated:       false,
		RuntimeLaunchApproval:      false,
		LaunchAllowed:              false,
		LaunchEnabled:              false,
		ExecutionStarted:           false,
		BackendProcessStarted:      false,
		RequestObjectsCreated:      false,
		PermissionGrantCreated:     false,
		HostRootModified:           false,
		NetworkRequired:            false,
		BackendDetailsExposed:      false,
		BlockedActions:             []string{"persist KDE center page from preview state", "enable center page action buttons from preview state", "persist KDE action card deck from center page preview", "persist compatibility settings from center page preview", "record review receipts from center page preview", "create Runtime request objects from center page preview", "grant desktop resources from center page preview", "send desktop notifications from center page preview", "start compatibility profile from center page preview", "mutate host root during KDE center page preview", "expose raw backend command to desktop shell"},
		UserFacingSettings:         settings.UserFacingSettings,
		DesktopSafeSummary:         "KDE can render an application page from Runtime-owned summary, action deck, and settings previews, but the page remains read-only and cannot approve, persist, grant, notify, or start execution.",
	}
	if err := validateNoBackendTerms(preview, "KDE center page preview"); err != nil {
		return KDECenterPagePreview{}, err
	}
	return preview, nil
}

func NewKDECenterPageSectionsPreview(recipe Recipe, provenance Provenance, decision string, fileURIs []string) (KDECenterPageSectionsPreview, error) {
	page, err := NewKDECenterPagePreview(recipe, provenance, decision, fileURIs)
	if err != nil {
		return KDECenterPageSectionsPreview{}, err
	}

	sections := kdeCenterPageSections()
	preview := KDECenterPageSectionsPreview{
		SchemaVersion:              "xnix.runtime.kde_center_page_sections.v1",
		RequestType:                "kde-center-page-sections-preview",
		PageType:                   page.PageType,
		Source:                     "kde-center-page-preview",
		Desktop:                    "KDE Plasma",
		RuntimeMethod:              "GetKDECenterPageSections",
		ReadMethod:                 "GetKDECenterPageSectionsPreview",
		ApplicationID:              page.ApplicationID,
		ApplicationName:            page.ApplicationName,
		SectionCount:               len(sections),
		ReadOnlySectionCount:       len(sections),
		NavigationOnlySectionCount: len(sections),
		ExecutableSectionCount:     0,
		PrimarySectionID:           "overview",
		Sections:                   sections,
		RuntimeOwned:               true,
		GoRuntimeBacked:            true,
		KDEPolicyOwner:             false,
		OfficialDesktopOnly:        true,
		UserVisible:                true,
		SafeForAIDiagnostics:       true,
		UserDecisionCaptured:       page.UserDecisionCaptured,
		UserDecisionAllowsLaunch:   page.UserDecisionAllowsLaunch,
		SectionsPreviewCreated:     true,
		SectionsPersisted:          false,
		SectionActionsEnabled:      false,
		SettingsPersisted:          false,
		SettingsPersistenceEnabled: false,
		NotificationsSent:          false,
		ResourceGrantCreated:       false,
		RuntimeLaunchApproval:      false,
		LaunchEnabled:              false,
		ExecutionStarted:           false,
		RequestObjectsCreated:      false,
		PermissionGrantCreated:     false,
		HostRootModified:           false,
		NetworkRequired:            false,
		BackendDetailsExposed:      false,
		BlockedActions:             []string{"persist KDE center page sections from preview state", "enable section action buttons from preview state", "persist compatibility settings from section navigation", "record review receipts from page sections", "create Runtime request objects from page sections", "grant desktop resources from page sections", "send desktop notifications from page sections", "start compatibility profile from page sections", "mutate host root during KDE center page section preview", "expose raw backend command to desktop shell"},
		DesktopSafeSummary:         "KDE can navigate Compatibility Center page sections backed by Runtime read models, but sections remain read-only and cannot persist, grant, notify, or start execution.",
	}
	if err := validateNoBackendTerms(preview, "KDE center page sections preview"); err != nil {
		return KDECenterPageSectionsPreview{}, err
	}
	return preview, nil
}

func kdeCenterPageNavigation() []KDECenterPageNavigation {
	return []KDECenterPageNavigation{
		kdeCenterPageNavigationItem("overview", "Overview", "compatibility-center-overview"),
		kdeCenterPageNavigationItem("actions", "Actions", "compatibility-center-gates"),
		kdeCenterPageNavigationItem("settings", "Settings", "compatibility-settings"),
		kdeCenterPageNavigationItem("diagnostics", "Diagnostics", "compatibility-diagnostics"),
	}
}

func kdeCenterPageNavigationItem(id string, label string, target string) KDECenterPageNavigation {
	return KDECenterPageNavigation{
		ID:             id,
		Label:          label,
		Target:         target,
		Enabled:        true,
		NavigationOnly: true,
		MutatesRuntime: false,
		StartsProgram:  false,
	}
}

func kdeCenterPageSections() []KDECenterPageSection {
	return []KDECenterPageSection{
		kdeCenterPageSection("overview", "Overview", "compatibility-center-overview", "GetCompatibilityCenterSummary", "compatibility-center-summary", "ready", "Overview reads Runtime-owned application state, known issue counts, and repair record state."),
		kdeCenterPageSection("actions", "Actions", "compatibility-center-gates", "GetCompatibilityActionQueue", "compatibility-center-action-queue", "waiting-for-runtime-gates", "Actions read queued review cards while all execution and mutation gates remain closed."),
		kdeCenterPageSection("settings", "Settings", "compatibility-settings", "GetCompatibilitySettings", "settings-model", "planned", "Settings read user-facing Runtime policy without persisting changes from KDE."),
		kdeCenterPageSection("diagnostics", "Diagnostics", "compatibility-diagnostics", "GetDiagnostics", "runtime-diagnostics", "planned", "Diagnostics read safe Runtime status for AI review without exposing backend implementation details."),
	}
}

func kdeCenterPageSection(id string, label string, target string, runtimeMethod string, readModel string, state string, summary string) KDECenterPageSection {
	return KDECenterPageSection{
		ID:                    id,
		Label:                 label,
		Target:                target,
		RuntimeMethod:         runtimeMethod,
		ReadModel:             readModel,
		State:                 state,
		NavigationOnly:        true,
		ReadOnly:              true,
		MutatesRuntime:        false,
		StartsProgram:         false,
		SettingsPersisted:     false,
		BackendDetailsExposed: false,
		Summary:               summary,
	}
}

func kdeCenterPageBadge(deck KDEActionCardDeckPreview) string {
	if deck.RejectedCardCount > 0 {
		return "Needs guidance"
	}
	if deck.DeferredCardCount > 0 {
		return "Deferred"
	}
	return "Review ready"
}

func kdeCenterPageBadgeTone(deck KDEActionCardDeckPreview) string {
	if deck.RejectedCardCount > 0 {
		return "critical"
	}
	if deck.DeferredCardCount > 0 {
		return "neutral"
	}
	return "warning"
}
