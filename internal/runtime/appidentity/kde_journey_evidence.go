package appidentity

import "errors"

type KDEJourneyEvidencePreview struct {
	SchemaVersion              string                    `json:"schema_version"`
	RequestType                string                    `json:"request_type"`
	JourneyType                string                    `json:"journey_type"`
	Source                     string                    `json:"source"`
	Desktop                    string                    `json:"desktop"`
	RuntimeMethod              string                    `json:"runtime_method"`
	ReadMethod                 string                    `json:"read_method"`
	ApplicationID              string                    `json:"application_id"`
	ApplicationName            string                    `json:"application_name"`
	Icon                       string                    `json:"icon"`
	DesktopFile                string                    `json:"desktop_file"`
	EntryPoints                []KDEJourneyEntryPoint    `json:"entry_points"`
	EntryPointIDs              []string                  `json:"entry_point_ids"`
	EntryPointCount            int                       `json:"entry_point_count"`
	CrossLinkedReadModels      []string                  `json:"cross_linked_read_models"`
	CrossLinkedReadModelCount  int                       `json:"cross_linked_read_model_count"`
	SharedReadinessStatus      string                    `json:"shared_readiness_status"`
	Ready                      bool                      `json:"ready"`
	ReadinessNodeCount         int                       `json:"readiness_node_count"`
	BlockedReadinessNodeCount  int                       `json:"blocked_readiness_node_count"`
	MissingEvidenceCount       int                       `json:"missing_evidence_count"`
	BlockedActionCount         int                       `json:"blocked_action_count"`
	SharedBlockedReasons       []string                  `json:"shared_blocked_reasons"`
	Agreement                  KDEJourneyAgreement       `json:"agreement"`
	NextReadOnlyChecks         []KDEJourneyReadOnlyCheck `json:"next_read_only_checks"`
	UserFacingSettings         map[string]string         `json:"user_facing_settings"`
	RuntimeOwned               bool                      `json:"runtime_owned"`
	GoRuntimeBacked            bool                      `json:"go_runtime_backed"`
	KDEPolicyOwner             bool                      `json:"kde_policy_owner"`
	OfficialDesktopOnly        bool                      `json:"official_desktop_only"`
	UserVisible                bool                      `json:"user_visible"`
	JourneyEvidenceCreated     bool                      `json:"journey_evidence_created"`
	JourneyEvidencePersisted   bool                      `json:"journey_evidence_persisted"`
	RuntimeWriteMethodsEnabled bool                      `json:"runtime_write_methods_enabled"`
	RequestObjectsCreated      bool                      `json:"request_objects_created"`
	PermissionGrantCreated     bool                      `json:"permission_grant_created"`
	SettingsPersisted          bool                      `json:"settings_persisted"`
	LaunchAllowed              bool                      `json:"launch_allowed"`
	LaunchEnabled              bool                      `json:"launch_enabled"`
	ExecutionStarted           bool                      `json:"execution_started"`
	BackendProcessStarted      bool                      `json:"backend_process_started"`
	KWinRuleApplied            bool                      `json:"kwin_rule_applied"`
	TrayBridgeActivated        bool                      `json:"tray_bridge_activated"`
	NotificationSent           bool                      `json:"notification_sent"`
	HostRootModified           bool                      `json:"host_root_modified"`
	StateRootPathExposed       bool                      `json:"state_root_path_exposed"`
	RawExecutableExposed       bool                      `json:"raw_executable_exposed"`
	RawCommandExposed          bool                      `json:"raw_command_exposed"`
	FileContentRead            bool                      `json:"file_content_read"`
	BackendDetailsExposed      bool                      `json:"backend_details_exposed"`
	DesktopSafeSummary         string                    `json:"desktop_safe_summary"`
}

type KDEJourneyEvidenceOptions struct {
	RuntimeRoot string
}

type KDEJourneyEntryPoint struct {
	ID                     string   `json:"id"`
	Label                  string   `json:"label"`
	Surface                string   `json:"surface"`
	RuntimeMethod          string   `json:"runtime_method"`
	ReadModel              string   `json:"read_model"`
	ApplicationID          string   `json:"application_id"`
	ApplicationName        string   `json:"application_name"`
	SharedReadinessStatus  string   `json:"shared_readiness_status"`
	Ready                  bool     `json:"ready"`
	EvidenceIDs            []string `json:"evidence_ids"`
	BlockedReasons         []string `json:"blocked_reasons"`
	NextSafeReadOnlyCheck  string   `json:"next_safe_read_only_check"`
	UserFacingSummary      string   `json:"user_facing_summary"`
	RuntimeOwned           bool     `json:"runtime_owned"`
	GoRuntimeBacked        bool     `json:"go_runtime_backed"`
	KDEPolicyOwner         bool     `json:"kde_policy_owner"`
	LaunchAllowed          bool     `json:"launch_allowed"`
	ExecutionStarted       bool     `json:"execution_started"`
	RequestObjectsCreated  bool     `json:"request_objects_created"`
	PermissionGrantCreated bool     `json:"permission_grant_created"`
	SettingsPersisted      bool     `json:"settings_persisted"`
	KWinRuleApplied        bool     `json:"kwin_rule_applied"`
	TrayBridgeActivated    bool     `json:"tray_bridge_activated"`
	NotificationSent       bool     `json:"notification_sent"`
	HostRootModified       bool     `json:"host_root_modified"`
	BackendDetailsExposed  bool     `json:"backend_details_exposed"`
	StateRootPathExposed   bool     `json:"state_root_path_exposed"`
	RawExecutableExposed   bool     `json:"raw_executable_exposed"`
	RawCommandExposed      bool     `json:"raw_command_exposed"`
}

type KDEJourneyAgreement struct {
	AppIDConsistent           bool     `json:"app_id_consistent"`
	DisplayNameConsistent     bool     `json:"display_name_consistent"`
	DesktopFileConsistent     bool     `json:"desktop_file_consistent"`
	ReadinessStateConsistent  bool     `json:"readiness_state_consistent"`
	DisabledActionStateShared bool     `json:"disabled_action_state_shared"`
	UnsafeSideEffectsDisabled bool     `json:"unsafe_side_effects_disabled"`
	CheckedReadModels         []string `json:"checked_read_models"`
}

type KDEJourneyReadOnlyCheck struct {
	ID             string `json:"id"`
	ReadModel      string `json:"read_model"`
	Reason         string `json:"reason"`
	MutatesRuntime bool   `json:"mutates_runtime"`
	StartsProgram  bool   `json:"starts_program"`
}

func (plan Plan) KDEJourneyEvidencePreview(decision string, fileURIs []string) (KDEJourneyEvidencePreview, error) {
	return plan.KDEJourneyEvidencePreviewWithOptions(decision, fileURIs, KDEJourneyEvidenceOptions{})
}

func (plan Plan) KDEJourneyEvidencePreviewWithOptions(decision string, fileURIs []string, options KDEJourneyEvidenceOptions) (KDEJourneyEvidencePreview, error) {
	if !singleLine(decision) {
		return KDEJourneyEvidencePreview{}, errors.New("KDE journey evidence preview requires a single-line decision")
	}
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return KDEJourneyEvidencePreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return KDEJourneyEvidencePreview{}, errors.New("KDE journey evidence preview requires single-line identity fields")
		}
	}

	if options.RuntimeRoot == "" {
		options.RuntimeRoot = "."
	}

	readiness, err := plan.ApplicationReadinessPreview(ApplicationReadinessOptions{RuntimeRoot: options.RuntimeRoot})
	if err != nil {
		return KDEJourneyEvidencePreview{}, err
	}
	graph, err := plan.KDEActionDependencyGraphPreview(decision, fileURIs)
	if err != nil {
		return KDEJourneyEvidencePreview{}, err
	}
	taskManager, err := plan.TaskManagerIdentityPlanPreview()
	if err != nil {
		return KDEJourneyEvidencePreview{}, err
	}
	kwin, err := plan.KWinWindowRulePlanPreview()
	if err != nil {
		return KDEJourneyEvidencePreview{}, err
	}
	tray, err := plan.TrayStatusPreview()
	if err != nil {
		return KDEJourneyEvidencePreview{}, err
	}
	notification, err := plan.NotificationPreview("approval-required")
	if err != nil {
		return KDEJourneyEvidencePreview{}, err
	}
	settings, err := plan.SettingsPreview()
	if err != nil {
		return KDEJourneyEvidencePreview{}, err
	}

	crossLinks := []string{
		"application-readiness-preview",
		"kde-action-dependency-graph-preview",
		"desktop-entry-preview",
		"task-manager-identity-preview",
		"kwin-window-rule-preview",
		"file-open-preview",
		"tray-status-preview",
		"notification-preview",
		"kde-center-page-preview",
		"settings-preview",
	}
	sharedBlocked := kdeJourneySharedBlockedReasons(readiness, graph)
	entryPoints := kdeJourneyEntryPoints(plan, readiness, graph, sharedBlocked)
	checkedModels := []string{
		readiness.RequestType,
		graph.RequestType,
		taskManager.RequestType,
		kwin.RequestType,
		tray.StatusType,
		notification.RequestType,
		settings.RequestType,
	}
	entryIDs := make([]string, 0, len(entryPoints))
	for _, entry := range entryPoints {
		entryIDs = append(entryIDs, entry.ID)
	}

	appIDConsistent := readiness.Application.ID == plan.ApplicationID &&
		graph.ApplicationID == plan.ApplicationID &&
		taskManager.ApplicationID == plan.ApplicationID &&
		kwin.ApplicationID == plan.ApplicationID &&
		tray.ApplicationID == plan.ApplicationID &&
		notification.ApplicationID == plan.ApplicationID &&
		settings.ApplicationID == plan.ApplicationID
	displayNameConsistent := readiness.Application.Name == plan.DisplayName &&
		graph.ApplicationName == plan.DisplayName &&
		taskManager.DisplayName == plan.DisplayName &&
		kwin.DisplayName == plan.DisplayName &&
		tray.DisplayName == plan.DisplayName &&
		notification.DisplayName == plan.DisplayName &&
		settings.DisplayName == plan.DisplayName
	desktopFileConsistent := readiness.Application.DesktopFile == plan.DesktopFile &&
		graph.DesktopFile == plan.DesktopFile &&
		taskManager.DesktopFile == plan.DesktopFile &&
		kwin.DesktopFile == plan.DesktopFile &&
		tray.DesktopFile == plan.DesktopFile &&
		notification.DesktopFile == plan.DesktopFile &&
		settings.DesktopFile == plan.DesktopFile

	preview := KDEJourneyEvidencePreview{
		SchemaVersion:             "xnix.runtime.kde_journey_evidence.v1",
		RequestType:               "kde-journey-evidence-preview",
		JourneyType:               "kde-seven-entrypoint-runtime-evidence",
		Source:                    "application-readiness-preview+kde-action-dependency-graph-preview+task-manager-identity-preview+kwin-window-rule-preview+tray-status-preview+notification-preview+settings-preview",
		Desktop:                   "KDE Plasma",
		RuntimeMethod:             "GetKDEJourneyEvidence",
		ReadMethod:                "GetKDEJourneyEvidencePreview",
		ApplicationID:             plan.ApplicationID,
		ApplicationName:           plan.DisplayName,
		Icon:                      plan.Icon,
		DesktopFile:               plan.DesktopFile,
		EntryPoints:               entryPoints,
		EntryPointIDs:             entryIDs,
		EntryPointCount:           len(entryPoints),
		CrossLinkedReadModels:     crossLinks,
		CrossLinkedReadModelCount: len(crossLinks),
		SharedReadinessStatus:     readiness.OverallStatus,
		Ready:                     readiness.Ready,
		ReadinessNodeCount:        readiness.NodeCount,
		BlockedReadinessNodeCount: readiness.BlockedNodeCount,
		MissingEvidenceCount:      graph.MissingEvidenceCount,
		BlockedActionCount:        graph.BlockedActionCount,
		SharedBlockedReasons:      sharedBlocked,
		Agreement: KDEJourneyAgreement{
			AppIDConsistent:           appIDConsistent,
			DisplayNameConsistent:     displayNameConsistent,
			DesktopFileConsistent:     desktopFileConsistent,
			ReadinessStateConsistent:  true,
			DisabledActionStateShared: true,
			UnsafeSideEffectsDisabled: true,
			CheckedReadModels:         checkedModels,
		},
		NextReadOnlyChecks: []KDEJourneyReadOnlyCheck{
			{ID: "refresh-readiness", ReadModel: "application-readiness-preview", Reason: "refresh shared install, lifecycle, Portal, snapshot, execution, and write-gate evidence", MutatesRuntime: false, StartsProgram: false},
			{ID: "refresh-action-dependencies", ReadModel: "kde-action-dependency-graph-preview", Reason: "refresh missing evidence and blocked action counts for every KDE entry point", MutatesRuntime: false, StartsProgram: false},
			{ID: "refresh-center", ReadModel: "kde-center-page-preview", Reason: "refresh the user-visible AI Compatibility Center page from the same Runtime evidence", MutatesRuntime: false, StartsProgram: false},
		},
		UserFacingSettings:         plan.UserFacingSettings,
		RuntimeOwned:               true,
		GoRuntimeBacked:            true,
		KDEPolicyOwner:             false,
		OfficialDesktopOnly:        true,
		UserVisible:                true,
		JourneyEvidenceCreated:     true,
		JourneyEvidencePersisted:   false,
		RuntimeWriteMethodsEnabled: false,
		RequestObjectsCreated:      false,
		PermissionGrantCreated:     false,
		SettingsPersisted:          false,
		LaunchAllowed:              false,
		LaunchEnabled:              false,
		ExecutionStarted:           false,
		BackendProcessStarted:      false,
		KWinRuleApplied:            false,
		TrayBridgeActivated:        false,
		NotificationSent:           false,
		HostRootModified:           false,
		StateRootPathExposed:       false,
		RawExecutableExposed:       false,
		RawCommandExposed:          false,
		FileContentRead:            false,
		BackendDetailsExposed:      false,
		DesktopSafeSummary:         "KDE can show one Runtime-owned journey state across launcher, task manager, file manager, tray, notifications, AI Compatibility Center, and unified settings while every unsafe action remains disabled.",
	}
	if err := validateNoBackendTerms(preview, "KDE journey evidence preview"); err != nil {
		return KDEJourneyEvidencePreview{}, err
	}
	return preview, nil
}

func kdeJourneySharedBlockedReasons(readiness ApplicationReadinessPreview, graph KDEActionDependencyGraphPreview) []string {
	return []string{
		"application readiness status is " + readiness.OverallStatus,
		"Runtime write gate is disabled",
		"KDE action dependency graph is missing evidence",
		"Compatibility Center review receipts are not recorded",
		"launch and execution remain disabled",
	}
}

func kdeJourneyEntryPoints(plan Plan, readiness ApplicationReadinessPreview, graph KDEActionDependencyGraphPreview, sharedBlocked []string) []KDEJourneyEntryPoint {
	sharedEvidence := []string{
		"application-readiness",
		"kde-action-dependency-graph",
		"runtime-write-gate",
	}
	return []KDEJourneyEntryPoint{
		kdeJourneyEntryPoint(plan, readiness, graph, "launcher", "Application launcher", "Plasma application launcher", "GetDesktopEntryPlan", "desktop-entry-preview", append(sharedEvidence, "desktop-entry-plan"), sharedBlocked, "desktop-entry-preview", "KDE can list this application beside native Linux applications while launch stays gated by Runtime evidence."),
		kdeJourneyEntryPoint(plan, readiness, graph, "task-manager", "Task manager", "Plasma task manager", "GetTaskManagerIdentityPlan", "task-manager-identity-preview", append(sharedEvidence, "task-manager-identity-plan"), sharedBlocked, "task-manager-identity-preview", "KDE can group and pin the application as a normal window identity while execution remains disabled."),
		kdeJourneyEntryPoint(plan, readiness, graph, "file-manager", "File manager", "Dolphin file manager", "GetFileOpenPlan", "file-open-preview", append(sharedEvidence, "file-association-plan", "portal-file-access-receipt"), sharedBlocked, "file-open-preview", "Dolphin can explain which files may be reviewed for opening while Portal and execution evidence remain required."),
		kdeJourneyEntryPoint(plan, readiness, graph, "system-tray", "System tray", "Plasma system tray", "GetTrayStatus", "tray-status-preview", append(sharedEvidence, "tray-status"), sharedBlocked, "tray-status-preview", "KDE can show compatibility status in the tray while live tray bridging stays disabled."),
		kdeJourneyEntryPoint(plan, readiness, graph, "notification-center", "Notification center", "Plasma notifications", "GetNotificationPlan", "notification-preview", append(sharedEvidence, "notification-plan"), sharedBlocked, "notification-preview", "KDE can prepare review notifications while no desktop notification is sent by this preview."),
		kdeJourneyEntryPoint(plan, readiness, graph, "ai-compatibility-center", "AI Compatibility Center", "KDE Compatibility Center", "GetKDECenterPage", "kde-center-page-preview", append(sharedEvidence, "kde-center-page", "action-card-deck"), sharedBlocked, "kde-center-page-preview", "The Compatibility Center can show readiness, blocked actions, and next read-only checks without approving execution."),
		kdeJourneyEntryPoint(plan, readiness, graph, "unified-settings", "Unified settings", "KDE system settings", "GetCompatibilitySettings", "settings-preview", append(sharedEvidence, "settings-preview"), sharedBlocked, "settings-preview", "KDE can show user-facing compatibility settings while persistence and permission grants remain disabled."),
	}
}

func kdeJourneyEntryPoint(plan Plan, readiness ApplicationReadinessPreview, graph KDEActionDependencyGraphPreview, id string, label string, surface string, runtimeMethod string, readModel string, evidenceIDs []string, sharedBlocked []string, nextCheck string, summary string) KDEJourneyEntryPoint {
	return KDEJourneyEntryPoint{
		ID:                     id,
		Label:                  label,
		Surface:                surface,
		RuntimeMethod:          runtimeMethod,
		ReadModel:              readModel,
		ApplicationID:          plan.ApplicationID,
		ApplicationName:        plan.DisplayName,
		SharedReadinessStatus:  readiness.OverallStatus,
		Ready:                  readiness.Ready && graph.BlockedActionCount == 0,
		EvidenceIDs:            evidenceIDs,
		BlockedReasons:         sharedBlocked,
		NextSafeReadOnlyCheck:  nextCheck,
		UserFacingSummary:      summary,
		RuntimeOwned:           true,
		GoRuntimeBacked:        true,
		KDEPolicyOwner:         false,
		LaunchAllowed:          false,
		ExecutionStarted:       false,
		RequestObjectsCreated:  false,
		PermissionGrantCreated: false,
		SettingsPersisted:      false,
		KWinRuleApplied:        false,
		TrayBridgeActivated:    false,
		NotificationSent:       false,
		HostRootModified:       false,
		BackendDetailsExposed:  false,
		StateRootPathExposed:   false,
		RawExecutableExposed:   false,
		RawCommandExposed:      false,
	}
}
