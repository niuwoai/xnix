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
	BackendSelectionSnapshot   KDECenterPageBackend      `json:"backend_selection_snapshot"`
	ActivationStatusSnapshot   KDECenterPageActivation   `json:"activation_status_snapshot"`
	ExecutionReadinessSnapshot KDECenterPageExecution    `json:"execution_readiness_snapshot"`
	LaunchIntentSnapshot       KDECenterPageLaunchIntent `json:"launch_intent_snapshot"`
	WindowIdentitySnapshot     KDECenterPageWindow       `json:"window_identity_snapshot"`
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

type KDECenterPageBackend struct {
	RequestType                 string `json:"request_type"`
	PlanType                    string `json:"plan_type"`
	RuntimeMethod               string `json:"runtime_method"`
	SelectedStrategy            string `json:"selected_strategy"`
	RecommendedProfileID        string `json:"recommended_profile_id"`
	CandidateCount              int    `json:"candidate_count"`
	ReadyCandidateCount         int    `json:"ready_candidate_count"`
	BlockedCandidateCount       int    `json:"blocked_candidate_count"`
	SelectionCommitted          bool   `json:"selection_committed"`
	SelectionChangeEnabled      bool   `json:"selection_change_enabled"`
	BackendLaunchEnabled        bool   `json:"backend_launch_enabled"`
	CapabilityActivationEnabled bool   `json:"capability_activation_enabled"`
	EnvironmentCreated          bool   `json:"environment_created"`
	HostRootModified            bool   `json:"host_root_modified"`
	BackendDetailsExposed       bool   `json:"backend_details_exposed"`
	Summary                     string `json:"summary"`
}

type KDECenterPageActivation struct {
	RequestType           string `json:"request_type"`
	StatusType            string `json:"status_type"`
	RuntimeMethod         string `json:"runtime_method"`
	Renderer              string `json:"renderer"`
	ActivationState       string `json:"activation_state"`
	TransactionState      string `json:"transaction_state"`
	PreflightDecision     string `json:"preflight_decision"`
	StatusSignalCount     int    `json:"status_signal_count"`
	BlockedReasonCount    int    `json:"blocked_reason_count"`
	NextSafeActionCount   int    `json:"next_safe_action_count"`
	ActivationReady       bool   `json:"activation_ready"`
	ActivationCommitted   bool   `json:"activation_committed"`
	CommitEnabled         bool   `json:"commit_enabled"`
	LaunchEnabled         bool   `json:"launch_enabled"`
	HostRootModified      bool   `json:"host_root_modified"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
	Summary               string `json:"summary"`
}

type KDECenterPageExecution struct {
	RequestType               string `json:"request_type"`
	ReadinessType             string `json:"readiness_type"`
	RuntimeMethod             string `json:"runtime_method"`
	ExecutionState            string `json:"execution_state"`
	OverallStatus             string `json:"overall_status"`
	RecommendedAction         string `json:"recommended_action"`
	GateCount                 int    `json:"gate_count"`
	RequiredGateCount         int    `json:"required_gate_count"`
	PendingGateCount          int    `json:"pending_gate_count"`
	BlockedGateCount          int    `json:"blocked_gate_count"`
	DesktopEntryLaunchVisible bool   `json:"desktop_entry_launch_visible"`
	LaunchAllowed             bool   `json:"launch_allowed"`
	LaunchEnabled             bool   `json:"launch_enabled"`
	ExecutionRequestCreated   bool   `json:"execution_request_created"`
	BackendBindingReady       bool   `json:"backend_binding_ready"`
	PortalPolicyRequired      bool   `json:"portal_policy_required"`
	SnapshotRequired          bool   `json:"snapshot_required"`
	UserActionRequired        bool   `json:"user_action_required"`
	HostRootModified          bool   `json:"host_root_modified"`
	NetworkRequired           bool   `json:"network_required"`
	BackendDetailsExposed     bool   `json:"backend_details_exposed"`
	Summary                   string `json:"summary"`
}

type KDECenterPageLaunchIntent struct {
	RequestType             string `json:"request_type"`
	IntentType              string `json:"intent_type"`
	Source                  string `json:"source"`
	RuntimeMethod           string `json:"runtime_method"`
	ReadMethod              string `json:"read_method"`
	FileCount               int    `json:"file_count"`
	PortalRequired          bool   `json:"portal_required"`
	SnapshotRequired        bool   `json:"snapshot_required"`
	StandardDesktopEntry    bool   `json:"standard_desktop_entry"`
	LaunchUsesRuntime       bool   `json:"launch_uses_runtime"`
	LaunchAllowed           bool   `json:"launch_allowed"`
	LaunchEnabled           bool   `json:"launch_enabled"`
	ExecutionRequestCreated bool   `json:"execution_request_created"`
	ExecutionStarted        bool   `json:"execution_started"`
	BackendBindingReady     bool   `json:"backend_binding_ready"`
	RequestObjectCreated    bool   `json:"request_object_created"`
	PermissionGranted       bool   `json:"permission_granted"`
	HostRootModified        bool   `json:"host_root_modified"`
	NetworkRequired         bool   `json:"network_required"`
	BackendDetailsExposed   bool   `json:"backend_details_exposed"`
	Summary                 string `json:"summary"`
}

type KDECenterPageWindow struct {
	SchemaVersion             string `json:"schema_version"`
	DesktopFile               string `json:"desktop_file"`
	LauncherURL               string `json:"launcher_url"`
	WindowKind                string `json:"window_kind"`
	ClassGroup                string `json:"class_group"`
	ResourceName              string `json:"resource_name"`
	TitleHint                 string `json:"title_hint"`
	TaskManagerGroupingKey    string `json:"task_manager_grouping_key"`
	TaskManagerPinningAllowed bool   `json:"task_manager_pinning_allowed"`
	TaskManagerRestoreAllowed bool   `json:"task_manager_restore_allowed"`
	TaskManagerSkipTaskbar    bool   `json:"task_manager_skip_taskbar"`
	TaskManagerShowInSwitcher bool   `json:"task_manager_show_in_switcher"`
	PreferExistingWindow      bool   `json:"prefer_existing_window"`
	KWinScriptRole            string `json:"kwin_script_role"`
	KWinPlacement             string `json:"kwin_placement"`
	WindowManagerPolicyOnly   bool   `json:"window_manager_policy_only"`
	RuntimeOwnsBackendPolicy  bool   `json:"runtime_owns_backend_policy"`
	TaskManagerEntryActive    bool   `json:"task_manager_entry_active"`
	KWinRuleApplied           bool   `json:"kwin_rule_applied"`
	HostRootModified          bool   `json:"host_root_modified"`
	BackendDetailsExposed     bool   `json:"backend_details_exposed"`
	Summary                   string `json:"summary"`
}

type KDECenterPageActionDeck struct {
	RequestType           string             `json:"request_type"`
	DeckType              string             `json:"deck_type"`
	CardCount             int                `json:"card_count"`
	WaitingCardCount      int                `json:"waiting_card_count"`
	DeferredCardCount     int                `json:"deferred_card_count"`
	RejectedCardCount     int                `json:"rejected_card_count"`
	AIAnalysis            *KDEAIAnalysisLink `json:"ai_analysis,omitempty"`
	AIAnalysisCardCount   int                `json:"ai_analysis_card_count"`
	NavigationActionCount int                `json:"navigation_action_count"`
	DisabledActionCount   int                `json:"disabled_action_count"`
	PrimaryCardID         string             `json:"primary_card_id"`
	CardIDs               []string           `json:"card_ids"`
	ActionQueueCreated    bool               `json:"action_queue_created"`
	ActionQueuePersisted  bool               `json:"action_queue_persisted"`
	DeckPreviewCreated    bool               `json:"deck_preview_created"`
	DeckPersisted         bool               `json:"deck_persisted"`
	CardsPersisted        bool               `json:"cards_persisted"`
	CardActionsEnabled    bool               `json:"card_actions_enabled"`
	RuntimeLaunchApproval bool               `json:"runtime_launch_approval"`
	LaunchAllowed         bool               `json:"launch_allowed"`
	ExecutionStarted      bool               `json:"execution_started"`
	BackendDetailsExposed bool               `json:"backend_details_exposed"`
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
	AIAnalysis                 *KDEAIAnalysisLink     `json:"ai_analysis,omitempty"`
	AIAnalysisSectionCount     int                    `json:"ai_analysis_section_count"`
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
	ID                    string             `json:"id"`
	Label                 string             `json:"label"`
	Target                string             `json:"target"`
	RuntimeMethod         string             `json:"runtime_method"`
	ReadModel             string             `json:"read_model"`
	State                 string             `json:"state"`
	AIAnalysis            *KDEAIAnalysisLink `json:"ai_analysis,omitempty"`
	NavigationOnly        bool               `json:"navigation_only"`
	ReadOnly              bool               `json:"read_only"`
	MutatesRuntime        bool               `json:"mutates_runtime"`
	StartsProgram         bool               `json:"starts_program"`
	SettingsPersisted     bool               `json:"settings_persisted"`
	BackendDetailsExposed bool               `json:"backend_details_exposed"`
	Summary               string             `json:"summary"`
}

type KDECenterPageSectionDetailPreview struct {
	SchemaVersion              string                        `json:"schema_version"`
	RequestType                string                        `json:"request_type"`
	PageType                   string                        `json:"page_type"`
	Source                     string                        `json:"source"`
	Desktop                    string                        `json:"desktop"`
	RuntimeMethod              string                        `json:"runtime_method"`
	ReadMethod                 string                        `json:"read_method"`
	ApplicationID              string                        `json:"application_id"`
	ApplicationName            string                        `json:"application_name"`
	SectionID                  string                        `json:"section_id"`
	SectionLabel               string                        `json:"section_label"`
	SectionTarget              string                        `json:"section_target"`
	SectionState               string                        `json:"section_state"`
	SectionRuntimeMethod       string                        `json:"section_runtime_method"`
	SectionReadModel           string                        `json:"section_read_model"`
	SectionSummary             string                        `json:"section_summary"`
	AIAnalysis                 *KDEAIAnalysisLink            `json:"ai_analysis,omitempty"`
	AIAnalysisInput            *KDECenterPageAIAnalysisInput `json:"ai_analysis_input,omitempty"`
	AvailableSectionIDs        []string                      `json:"available_section_ids"`
	ReadOnlyNavigation         bool                          `json:"read_only_navigation"`
	DetailPreviewCreated       bool                          `json:"detail_preview_created"`
	DetailPersisted            bool                          `json:"detail_persisted"`
	SectionActionsEnabled      bool                          `json:"section_actions_enabled"`
	SettingsPersisted          bool                          `json:"settings_persisted"`
	SettingsPersistenceEnabled bool                          `json:"settings_persistence_enabled"`
	NotificationsSent          bool                          `json:"notifications_sent"`
	ResourceGrantCreated       bool                          `json:"resource_grant_created"`
	RuntimeLaunchApproval      bool                          `json:"runtime_launch_approval"`
	LaunchEnabled              bool                          `json:"launch_enabled"`
	ExecutionStarted           bool                          `json:"execution_started"`
	RequestObjectsCreated      bool                          `json:"request_objects_created"`
	PermissionGrantCreated     bool                          `json:"permission_grant_created"`
	HostRootModified           bool                          `json:"host_root_modified"`
	NetworkRequired            bool                          `json:"network_required"`
	BackendDetailsExposed      bool                          `json:"backend_details_exposed"`
	RuntimeOwned               bool                          `json:"runtime_owned"`
	GoRuntimeBacked            bool                          `json:"go_runtime_backed"`
	KDEPolicyOwner             bool                          `json:"kde_policy_owner"`
	OfficialDesktopOnly        bool                          `json:"official_desktop_only"`
	UserVisible                bool                          `json:"user_visible"`
	SafeForAIDiagnostics       bool                          `json:"safe_for_ai_diagnostics"`
	UserDecisionCaptured       bool                          `json:"user_decision_captured"`
	UserDecisionAllowsLaunch   bool                          `json:"user_decision_allows_launch"`
	BlockedActions             []string                      `json:"blocked_actions"`
	DesktopSafeSummary         string                        `json:"desktop_safe_summary"`
}

type KDECenterPageAIAnalysisInput struct {
	RequestType            string `json:"request_type"`
	Source                 string `json:"source"`
	RuntimeMethod          string `json:"runtime_method"`
	AnalysisTask           string `json:"analysis_task"`
	AnalysisSurface        string `json:"analysis_surface"`
	SelectionMode          string `json:"selection_mode"`
	FileCount              int    `json:"file_count"`
	SelectedExtension      string `json:"selected_extension"`
	SelectedFileDisclosure string `json:"selected_file_disclosure"`
	UserReviewRequired     bool   `json:"user_review_required"`
	SafeForAIDiagnostics   bool   `json:"safe_for_ai_diagnostics"`
	AIProviderCallEnabled  bool   `json:"ai_provider_call_enabled"`
	NetworkRequired        bool   `json:"network_required"`
	FileContentRead        bool   `json:"file_content_read"`
	FilePathsExposed       bool   `json:"file_paths_exposed"`
	RequestObjectCreated   bool   `json:"request_object_created"`
	PermissionGranted      bool   `json:"permission_granted"`
	BackendLaunchEnabled   bool   `json:"backend_launch_enabled"`
	HostRootModified       bool   `json:"host_root_modified"`
	BackendDetailsExposed  bool   `json:"backend_details_exposed"`
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
	backendSelection, err := plan.BackendSelectionPreview()
	if err != nil {
		return KDECenterPagePreview{}, err
	}
	activationStatus, err := plan.DesktopActivationStatusPreview("development")
	if err != nil {
		return KDECenterPagePreview{}, err
	}
	executionReadiness, err := plan.ExecutionReadinessPreview()
	if err != nil {
		return KDECenterPagePreview{}, err
	}
	launchIntent, err := plan.LaunchIntentPreview(fileURIs)
	if err != nil {
		return KDECenterPagePreview{}, err
	}
	windowIdentity, err := plan.WindowIdentityPreview()
	if err != nil {
		return KDECenterPagePreview{}, err
	}

	application := center.Applications[0]
	navigation := kdeCenterPageNavigation()
	preview := KDECenterPagePreview{
		SchemaVersion:   "xnix.runtime.kde_center_page.v1",
		RequestType:     "kde-center-page-preview",
		PageType:        "compatibility-center-application-page",
		Source:          "compatibility-center-preview+backend-selection-preview+desktop-activation-status-preview+execution-readiness-preview+launch-intent-preview+window-identity-preview+kde-action-card-deck-preview+settings-preview",
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
		BackendSelectionSnapshot: KDECenterPageBackend{
			RequestType:                 backendSelection.RequestType,
			PlanType:                    backendSelection.PlanType,
			RuntimeMethod:               backendSelection.RuntimeMethod,
			SelectedStrategy:            backendSelection.SelectedStrategy,
			RecommendedProfileID:        backendSelection.RecommendedProfileID,
			CandidateCount:              backendSelection.CandidateCount,
			ReadyCandidateCount:         backendSelection.ReadyCandidateCount,
			BlockedCandidateCount:       backendSelection.BlockedCandidateCount,
			SelectionCommitted:          backendSelection.SelectionCommitted,
			SelectionChangeEnabled:      backendSelection.SelectionChangeEnabled,
			BackendLaunchEnabled:        backendSelection.BackendLaunchEnabled,
			CapabilityActivationEnabled: backendSelection.CapabilityActivationEnabled,
			EnvironmentCreated:          backendSelection.EnvironmentCreated,
			HostRootModified:            backendSelection.HostRootModified,
			BackendDetailsExposed:       backendSelection.BackendDetailsExposed,
			Summary:                     backendSelection.DesktopSafeSummary,
		},
		ActivationStatusSnapshot: KDECenterPageActivation{
			RequestType:           activationStatus.RequestType,
			StatusType:            activationStatus.StatusType,
			RuntimeMethod:         activationStatus.PlannedRuntimeMethod,
			Renderer:              activationStatus.CurrentRenderer,
			ActivationState:       activationStatus.ActivationState,
			TransactionState:      activationStatus.Transaction.TransactionState,
			PreflightDecision:     activationStatus.PreflightDecision,
			StatusSignalCount:     len(activationStatus.StatusSignals),
			BlockedReasonCount:    len(activationStatus.BlockedReasons),
			NextSafeActionCount:   len(activationStatus.NextSafeActions),
			ActivationReady:       activationStatus.ActivationReady,
			ActivationCommitted:   activationStatus.ActivationCommitted,
			CommitEnabled:         activationStatus.CommitGate.CommitEnabled,
			LaunchEnabled:         activationStatus.LaunchEnabled,
			HostRootModified:      activationStatus.HostRootModified,
			BackendDetailsExposed: activationStatus.BackendDetailsExposed,
			Summary:               activationStatus.DesktopSafeSummary,
		},
		ExecutionReadinessSnapshot: KDECenterPageExecution{
			RequestType:               executionReadiness.RequestType,
			ReadinessType:             executionReadiness.ReadinessType,
			RuntimeMethod:             executionReadiness.RuntimeMethod,
			ExecutionState:            executionReadiness.ExecutionState,
			OverallStatus:             executionReadiness.OverallStatus,
			RecommendedAction:         executionReadiness.RecommendedAction,
			GateCount:                 executionReadiness.GateCount,
			RequiredGateCount:         executionReadiness.RequiredGateCount,
			PendingGateCount:          executionReadiness.PendingGateCount,
			BlockedGateCount:          executionReadiness.BlockedGateCount,
			DesktopEntryLaunchVisible: executionReadiness.DesktopEntryLaunchVisible,
			LaunchAllowed:             executionReadiness.LaunchAllowed,
			LaunchEnabled:             executionReadiness.LaunchEnabled,
			ExecutionRequestCreated:   executionReadiness.ExecutionRequestCreated,
			BackendBindingReady:       executionReadiness.BackendBindingReady,
			PortalPolicyRequired:      executionReadiness.PortalPolicyRequired,
			SnapshotRequired:          executionReadiness.SnapshotRequired,
			UserActionRequired:        executionReadiness.UserActionRequired,
			HostRootModified:          executionReadiness.HostRootModified,
			NetworkRequired:           executionReadiness.NetworkRequired,
			BackendDetailsExposed:     executionReadiness.BackendDetailsExposed,
			Summary:                   executionReadiness.DesktopSafeSummary,
		},
		LaunchIntentSnapshot: KDECenterPageLaunchIntent{
			RequestType:             launchIntent.RequestType,
			IntentType:              launchIntent.IntentType,
			Source:                  launchIntent.Source,
			RuntimeMethod:           launchIntent.RuntimeMethod,
			ReadMethod:              launchIntent.ReadMethod,
			FileCount:               launchIntent.FileCount,
			PortalRequired:          launchIntent.PortalRequired,
			SnapshotRequired:        launchIntent.SnapshotRequired,
			StandardDesktopEntry:    launchIntent.StandardDesktopEntry,
			LaunchUsesRuntime:       launchIntent.LaunchUsesRuntime,
			LaunchAllowed:           launchIntent.LaunchAllowed,
			LaunchEnabled:           launchIntent.LaunchEnabled,
			ExecutionRequestCreated: launchIntent.ExecutionRequestCreated,
			ExecutionStarted:        launchIntent.ExecutionStarted,
			BackendBindingReady:     launchIntent.BackendBindingReady,
			RequestObjectCreated:    launchIntent.RequestObjectCreated,
			PermissionGranted:       launchIntent.PermissionGranted,
			HostRootModified:        launchIntent.HostRootModified,
			NetworkRequired:         launchIntent.NetworkRequired,
			BackendDetailsExposed:   launchIntent.BackendDetailsExposed,
			Summary:                 launchIntent.DesktopSafeSummary,
		},
		WindowIdentitySnapshot: KDECenterPageWindow{
			SchemaVersion:             windowIdentity.SchemaVersion,
			DesktopFile:               windowIdentity.DesktopFile,
			LauncherURL:               windowIdentity.LauncherURL,
			WindowKind:                windowIdentity.WindowKind,
			ClassGroup:                windowIdentity.ClassGroup,
			ResourceName:              windowIdentity.ResourceName,
			TitleHint:                 windowIdentity.TitleHint,
			TaskManagerGroupingKey:    windowIdentity.TaskManager.GroupingKey,
			TaskManagerPinningAllowed: windowIdentity.TaskManager.PinningAllowed,
			TaskManagerRestoreAllowed: windowIdentity.TaskManager.RestoreAllowed,
			TaskManagerSkipTaskbar:    windowIdentity.TaskManager.SkipTaskbar,
			TaskManagerShowInSwitcher: windowIdentity.TaskManager.ShowInSwitcher,
			PreferExistingWindow:      windowIdentity.TaskManager.PreferExistingWindow,
			KWinScriptRole:            windowIdentity.KWin.ScriptRole,
			KWinPlacement:             windowIdentity.KWin.Placement,
			WindowManagerPolicyOnly:   windowIdentity.KWin.WindowManagerPolicyOnly,
			RuntimeOwnsBackendPolicy:  windowIdentity.KWin.RuntimeOwnsBackendPolicy,
			TaskManagerEntryActive:    false,
			KWinRuleApplied:           false,
			HostRootModified:          windowIdentity.HostRootModified,
			BackendDetailsExposed:     windowIdentity.BackendDetailsExposed,
			Summary:                   windowIdentity.Summary,
		},
		ActionDeck: KDECenterPageActionDeck{
			RequestType:           deck.RequestType,
			DeckType:              deck.DeckType,
			CardCount:             deck.CardCount,
			WaitingCardCount:      deck.WaitingCardCount,
			DeferredCardCount:     deck.DeferredCardCount,
			RejectedCardCount:     deck.RejectedCardCount,
			AIAnalysis:            deck.AIAnalysis,
			AIAnalysisCardCount:   deck.AIAnalysisCardCount,
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
		DesktopSafeSummary:         "KDE can render an application page from Runtime-owned summary, execution readiness, launch intent, window identity, action deck, and settings previews, but the page remains read-only and cannot approve, persist, grant, notify, or start execution.",
	}
	if err := validateNoBackendTerms(preview, "KDE center page preview"); err != nil {
		return KDECenterPagePreview{}, err
	}
	return preview, nil
}

func kdeCenterPageAIAnalysisInput(recipe Recipe, provenance Provenance, sectionID string, fileURIs []string) (*KDECenterPageAIAnalysisInput, error) {
	if sectionID != "diagnostics" || len(fileURIs) == 0 {
		return nil, nil
	}
	preview, err := NewDolphinAIAnalysisPreview([]Recipe{recipe}, provenance, fileURIs, recipe.ID)
	if err != nil {
		return nil, err
	}
	input := KDECenterPageAIAnalysisInput{
		RequestType:            preview.RequestType,
		Source:                 preview.Source,
		RuntimeMethod:          preview.RuntimeMethod,
		AnalysisTask:           preview.AnalysisTask,
		AnalysisSurface:        preview.AnalysisSurface,
		SelectionMode:          preview.SelectionMode,
		FileCount:              preview.FileCount,
		SelectedExtension:      preview.SelectedExtension,
		SelectedFileDisclosure: preview.SelectedFileDisclosure,
		UserReviewRequired:     preview.UserReviewRequired,
		SafeForAIDiagnostics:   preview.SafeForAIDiagnostics,
		AIProviderCallEnabled:  preview.AIProviderCallEnabled,
		NetworkRequired:        preview.NetworkRequired,
		FileContentRead:        preview.FileContentRead,
		FilePathsExposed:       preview.FilePathsExposed,
		RequestObjectCreated:   preview.RequestObjectCreated,
		PermissionGranted:      preview.PermissionGranted,
		BackendLaunchEnabled:   preview.BackendLaunchEnabled,
		HostRootModified:       preview.HostRootModified,
		BackendDetailsExposed:  preview.BackendDetailsExposed,
	}
	return &input, nil
}

func NewKDECenterPageSectionsPreview(recipe Recipe, provenance Provenance, decision string, fileURIs []string) (KDECenterPageSectionsPreview, error) {
	page, err := NewKDECenterPagePreview(recipe, provenance, decision, fileURIs)
	if err != nil {
		return KDECenterPageSectionsPreview{}, err
	}

	sections := kdeCenterPageSections()
	aiAnalysis := firstKDECenterPageSectionAIAnalysis(sections)
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
		AIAnalysis:                 aiAnalysis,
		AIAnalysisSectionCount:     countKDECenterPageSectionAIAnalysis(sections),
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

func NewKDECenterPageSectionDetailPreview(recipe Recipe, provenance Provenance, sectionID string, decision string, fileURIs []string) (KDECenterPageSectionDetailPreview, error) {
	if !singleLine(sectionID) {
		return KDECenterPageSectionDetailPreview{}, errors.New("KDE center page section detail preview requires a single-line section id")
	}

	sections, err := NewKDECenterPageSectionsPreview(recipe, provenance, decision, fileURIs)
	if err != nil {
		return KDECenterPageSectionDetailPreview{}, err
	}

	section, ok := findKDECenterPageSection(sections.Sections, sectionID)
	if !ok {
		return KDECenterPageSectionDetailPreview{}, errors.New("KDE center page section detail preview requires a known section id")
	}
	aiAnalysisInput, err := kdeCenterPageAIAnalysisInput(recipe, provenance, section.ID, fileURIs)
	if err != nil {
		return KDECenterPageSectionDetailPreview{}, err
	}

	preview := KDECenterPageSectionDetailPreview{
		SchemaVersion:              "xnix.runtime.kde_center_page_section_detail.v1",
		RequestType:                "kde-center-page-section-detail-preview",
		PageType:                   sections.PageType,
		Source:                     "kde-center-page-sections-preview",
		Desktop:                    "KDE Plasma",
		RuntimeMethod:              "GetKDECenterPageSectionDetail",
		ReadMethod:                 "GetKDECenterPageSectionDetailPreview",
		ApplicationID:              sections.ApplicationID,
		ApplicationName:            sections.ApplicationName,
		SectionID:                  section.ID,
		SectionLabel:               section.Label,
		SectionTarget:              section.Target,
		SectionState:               section.State,
		SectionRuntimeMethod:       section.RuntimeMethod,
		SectionReadModel:           section.ReadModel,
		SectionSummary:             section.Summary,
		AIAnalysis:                 section.AIAnalysis,
		AIAnalysisInput:            aiAnalysisInput,
		AvailableSectionIDs:        kdeCenterPageSectionIDs(sections.Sections),
		ReadOnlyNavigation:         true,
		DetailPreviewCreated:       true,
		DetailPersisted:            false,
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
		RuntimeOwned:               true,
		GoRuntimeBacked:            true,
		KDEPolicyOwner:             false,
		OfficialDesktopOnly:        true,
		UserVisible:                true,
		SafeForAIDiagnostics:       true,
		UserDecisionCaptured:       sections.UserDecisionCaptured,
		UserDecisionAllowsLaunch:   sections.UserDecisionAllowsLaunch,
		BlockedActions:             []string{"persist KDE center page section detail from preview state", "enable section action buttons from detail preview", "persist compatibility settings from section detail", "record review receipts from section detail", "create Runtime request objects from section detail", "grant desktop resources from section detail", "send desktop notifications from section detail", "start compatibility profile from section detail", "mutate host root during KDE center page section detail preview", "expose raw backend command to desktop shell"},
		DesktopSafeSummary:         "KDE can open a selected Compatibility Center section and route it to a Runtime read model, but the detail remains read-only and cannot persist, grant, notify, or start execution.",
	}
	if err := validateNoBackendTerms(preview, "KDE center page section detail preview"); err != nil {
		return KDECenterPageSectionDetailPreview{}, err
	}
	return preview, nil
}

func kdeCenterPageNavigation() []KDECenterPageNavigation {
	return []KDECenterPageNavigation{
		kdeCenterPageNavigationItem("overview", "Overview", "compatibility-center-overview"),
		kdeCenterPageNavigationItem("backend", "Backend", "compatibility-backend-selection"),
		kdeCenterPageNavigationItem("activation", "Activation", "compatibility-activation-status"),
		kdeCenterPageNavigationItem("execution", "Execution", "compatibility-execution-readiness"),
		kdeCenterPageNavigationItem("launch", "Launch", "compatibility-launch-intent"),
		kdeCenterPageNavigationItem("window", "Window", "compatibility-window-identity"),
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
		kdeCenterPageSection("backend", "Backend", "compatibility-backend-selection", "GetBackendSelectionPlan", "backend-selection-preview", "selection-pending", "Backend reads Runtime-owned recommended compatibility profile while selection commit, environment creation, and launch remain closed."),
		kdeCenterPageSection("activation", "Activation", "compatibility-activation-status", "GetDesktopActivationStatus", "desktop-activation-status-preview", "ready-for-runtime-commit", "Activation reads Runtime-owned KDE desktop activation status while commit, launch, and host mutation gates remain closed."),
		kdeCenterPageSection("execution", "Execution", "compatibility-execution-readiness", "GetExecutionReadiness", "execution-readiness-preview", "blocked", "Execution reads Runtime-owned launch readiness while request creation, launch, and backend process gates remain closed."),
		kdeCenterPageSection("launch", "Launch", "compatibility-launch-intent", "GetLaunchIntent", "launch-intent-preview", "blocked", "Launch reads Runtime-owned desktop-launch intent while Launch request creation, permission grants, execution, and backend process gates remain closed."),
		kdeCenterPageSection("window", "Window", "compatibility-window-identity", "GetTaskManagerIdentityPlan", "window-identity-preview", "planned", "Window reads Runtime-owned task-manager and KWin identity hints while task-manager activation, KWin rule application, execution, and backend policy stay closed."),
		kdeCenterPageSection("actions", "Actions", "compatibility-center-gates", "GetCompatibilityActionQueue", "compatibility-center-action-queue", "waiting-for-runtime-gates", "Actions read queued review cards while all execution and mutation gates remain closed."),
		kdeCenterPageSection("settings", "Settings", "compatibility-settings", "GetCompatibilitySettings", "settings-model", "planned", "Settings read user-facing Runtime policy without persisting changes from KDE."),
		kdeCenterPageSection("diagnostics", "Diagnostics", "compatibility-diagnostics", "GetAIDiagnosticInput", "ai-diagnostic-input", "planned", "Diagnostics read AI-safe Runtime status and Dolphin file analysis metadata without exposing backend implementation details."),
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
		AIAnalysis:            kdeCenterPageSectionAIAnalysis(id),
		NavigationOnly:        true,
		ReadOnly:              true,
		MutatesRuntime:        false,
		StartsProgram:         false,
		SettingsPersisted:     false,
		BackendDetailsExposed: false,
		Summary:               summary,
	}
}

func kdeCenterPageSectionAIAnalysis(sectionID string) *KDEAIAnalysisLink {
	if sectionID != "diagnostics" {
		return nil
	}
	link := kdeAIAnalysisLinkForEntryPoint("file-manager")
	return link
}

func firstKDECenterPageSectionAIAnalysis(sections []KDECenterPageSection) *KDEAIAnalysisLink {
	for _, section := range sections {
		if section.AIAnalysis != nil {
			return section.AIAnalysis
		}
	}
	return nil
}

func countKDECenterPageSectionAIAnalysis(sections []KDECenterPageSection) int {
	count := 0
	for _, section := range sections {
		if section.AIAnalysis != nil {
			count++
		}
	}
	return count
}

func findKDECenterPageSection(sections []KDECenterPageSection, sectionID string) (KDECenterPageSection, bool) {
	for _, section := range sections {
		if section.ID == sectionID {
			return section, true
		}
	}
	return KDECenterPageSection{}, false
}

func kdeCenterPageSectionIDs(sections []KDECenterPageSection) []string {
	ids := make([]string, 0, len(sections))
	for _, section := range sections {
		ids = append(ids, section.ID)
	}
	return ids
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
