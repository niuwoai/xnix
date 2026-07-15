package appidentity

type DiagnosticsPreview struct {
	SchemaVersion              string                   `json:"schema_version"`
	RequestType                string                   `json:"request_type"`
	DiagnosticsType            string                   `json:"diagnostics_type"`
	Source                     string                   `json:"source"`
	Desktop                    string                   `json:"desktop"`
	RuntimeMethod              string                   `json:"runtime_method"`
	ReadMethod                 string                   `json:"read_method"`
	Application                DiagnosticsApplication   `json:"application"`
	Status                     string                   `json:"status"`
	RuntimeMode                string                   `json:"runtime_mode"`
	Checks                     []DiagnosticsCheck       `json:"checks"`
	CheckIDs                   []string                 `json:"check_ids"`
	Counts                     DiagnosticsCounts        `json:"counts"`
	ExecutionReadiness         DiagnosticsExecution     `json:"execution_readiness"`
	LaunchIntent               DiagnosticsLaunchIntent  `json:"launch_intent"`
	ActionQueue                DiagnosticsActionQueue   `json:"action_queue"`
	CompatibilityCenterSummary DiagnosticsCenterSummary `json:"compatibility_center_summary"`
	KDECenterSections          DiagnosticsKDESections   `json:"kde_center_sections"`
	AI                         DiagnosticsAISummary     `json:"ai"`
	RuntimeOwned               bool                     `json:"runtime_owned"`
	GoRuntimeBacked            bool                     `json:"go_runtime_backed"`
	KDEPolicyOwner             bool                     `json:"kde_policy_owner"`
	UserVisible                bool                     `json:"user_visible"`
	CompatibilityCenterCard    bool                     `json:"compatibility_center_card"`
	SafeForAIDiagnostics       bool                     `json:"safe_for_ai_diagnostics"`
	AIProviderCallEnabled      bool                     `json:"ai_provider_call_enabled"`
	FileContentRead            bool                     `json:"file_content_read"`
	FilePathsExposed           bool                     `json:"file_paths_exposed"`
	RequestObjectCreated       bool                     `json:"request_object_created"`
	PermissionGranted          bool                     `json:"permission_granted"`
	LaunchEnabled              bool                     `json:"launch_enabled"`
	ExecutionStarted           bool                     `json:"execution_started"`
	RepairExecutionEnabled     bool                     `json:"repair_execution_enabled"`
	SettingsPersisted          bool                     `json:"settings_persisted"`
	HostRootModified           bool                     `json:"host_root_modified"`
	NetworkRequired            bool                     `json:"network_required"`
	BackendDetailsExposed      bool                     `json:"backend_details_exposed"`
	BlockedActions             []string                 `json:"blocked_actions"`
	DesktopSafeSummary         string                   `json:"desktop_safe_summary"`
}

type DiagnosticsApplication struct {
	ID                  string   `json:"id"`
	Name                string   `json:"name"`
	Icon                string   `json:"icon"`
	DesktopFile         string   `json:"desktop_file"`
	SupportedExtensions []string `json:"supported_extensions"`
	RegistryName        string   `json:"registry_name,omitempty"`
	DigestVerified      bool     `json:"digest_verified"`
	SignatureStatus     string   `json:"signature_status,omitempty"`
}

type DiagnosticsCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type DiagnosticsCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

type DiagnosticsExecution struct {
	RequestType           string `json:"request_type"`
	OverallStatus         string `json:"overall_status"`
	ExecutionState        string `json:"execution_state"`
	GateCount             int    `json:"gate_count"`
	RequiredGateCount     int    `json:"required_gate_count"`
	PendingGateCount      int    `json:"pending_gate_count"`
	BlockedGateCount      int    `json:"blocked_gate_count"`
	LaunchEnabled         bool   `json:"launch_enabled"`
	BackendBindingReady   bool   `json:"backend_binding_ready"`
	SafeForAIDiagnostics  bool   `json:"safe_for_ai_diagnostics"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

type DiagnosticsLaunchIntent struct {
	RequestType           string `json:"request_type"`
	OverallStatus         string `json:"overall_status"`
	WriteGateDecision     string `json:"write_gate_decision"`
	PortalRequired        bool   `json:"portal_required"`
	SnapshotRequired      bool   `json:"snapshot_required"`
	LaunchEnabled         bool   `json:"launch_enabled"`
	ExecutionStarted      bool   `json:"execution_started"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

type DiagnosticsActionQueue struct {
	RequestType             string `json:"request_type"`
	ActionCount             int    `json:"action_count"`
	PendingActionCount      int    `json:"pending_action_count"`
	UserReviewRequiredCount int    `json:"user_review_required_count"`
	RuntimeGateActionCount  int    `json:"runtime_gate_action_count"`
	ActionQueueCreated      bool   `json:"action_queue_created"`
	ActionQueuePersisted    bool   `json:"action_queue_persisted"`
	ExecutionStarted        bool   `json:"execution_started"`
	BackendDetailsExposed   bool   `json:"backend_details_exposed"`
}

type DiagnosticsCenterSummary struct {
	RequestType                string `json:"request_type"`
	CompatibilityState         string `json:"compatibility_state"`
	DiagnosticsState           string `json:"diagnostics_state"`
	KnownIssueCount            int    `json:"known_issue_count"`
	RepairRecordState          string `json:"repair_record_state"`
	ActionExecutionEnabled     bool   `json:"action_execution_enabled"`
	RepairExecutionEnabled     bool   `json:"repair_execution_enabled"`
	BackendLaunchEnabled       bool   `json:"backend_launch_enabled"`
	SettingsPersistenceEnabled bool   `json:"settings_persistence_enabled"`
	BackendDetailsExposed      bool   `json:"backend_details_exposed"`
}

type DiagnosticsKDESections struct {
	RequestType           string   `json:"request_type"`
	SectionCount          int      `json:"section_count"`
	ReadOnlySectionCount  int      `json:"read_only_section_count"`
	SectionIDs            []string `json:"section_ids"`
	RuntimeMethods        []string `json:"runtime_methods"`
	SectionActionsEnabled bool     `json:"section_actions_enabled"`
	RequestObjectsCreated bool     `json:"request_objects_created"`
	ExecutionStarted      bool     `json:"execution_started"`
	BackendDetailsExposed bool     `json:"backend_details_exposed"`
}

type DiagnosticsAISummary struct {
	RuntimeMethod         string `json:"runtime_method"`
	AnalysisTask          string `json:"analysis_task"`
	Disclosure            string `json:"disclosure"`
	SafeForAIDiagnostics  bool   `json:"safe_for_ai_diagnostics"`
	AIProviderCallEnabled bool   `json:"ai_provider_call_enabled"`
	NetworkRequired       bool   `json:"network_required"`
	FileContentRead       bool   `json:"file_content_read"`
	FilePathsExposed      bool   `json:"file_paths_exposed"`
	RequestObjectCreated  bool   `json:"request_object_created"`
	PermissionGranted     bool   `json:"permission_granted"`
}

func NewDiagnosticsPreview(recipe Recipe, provenance Provenance) (DiagnosticsPreview, error) {
	plan, err := NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return DiagnosticsPreview{}, err
	}
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return DiagnosticsPreview{}, err
	}

	readiness, err := plan.ExecutionReadinessPreview()
	if err != nil {
		return DiagnosticsPreview{}, err
	}
	launchIntent, err := plan.LaunchIntentPreview(nil)
	if err != nil {
		return DiagnosticsPreview{}, err
	}
	actionQueue, err := plan.KDEActionQueuePreview("approved", nil)
	if err != nil {
		return DiagnosticsPreview{}, err
	}
	center, err := NewCompatibilityCenterPreview([]Recipe{recipe}, provenance)
	if err != nil {
		return DiagnosticsPreview{}, err
	}
	centerApp := center.Applications[0]
	sections, err := NewKDECenterPageSectionsPreview(recipe, provenance, "approved", nil)
	if err != nil {
		return DiagnosticsPreview{}, err
	}

	checks := []DiagnosticsCheck{
		diagnosticsCheck("recipe-validation", "pass", "Recipe metadata is valid and can produce desktop-safe Runtime read models."),
		diagnosticsCheck("desktop-identity", "pass", "The application is represented as a normal KDE application entry."),
		diagnosticsCheck("execution-readiness", "blocked", "Compatibility execution remains blocked until Runtime-owned launch gates are complete."),
		diagnosticsCheck("launch-write-gate", "blocked", "Launch remains disabled until production Runtime ownership is proven."),
		diagnosticsCheck("action-queue", "pending", "KDE review actions are available but are not persisted or executable from diagnostics."),
		diagnosticsCheck("ai-diagnostics-safety", "pass", "Diagnostics use Runtime-safe metadata only and do not call an AI provider."),
	}
	counts := countDiagnosticsChecks(checks)

	preview := DiagnosticsPreview{
		SchemaVersion:   "xnix.runtime.diagnostics.v1",
		RequestType:     "diagnostics-preview",
		DiagnosticsType: "runtime-diagnostics",
		Source:          "registry+go-runtime-previews",
		Desktop:         "KDE Plasma",
		RuntimeMethod:   "GetDiagnostics",
		ReadMethod:      "GetDiagnosticsPreview",
		Application: DiagnosticsApplication{
			ID:                  plan.ApplicationID,
			Name:                plan.DisplayName,
			Icon:                plan.Icon,
			DesktopFile:         plan.DesktopFile,
			SupportedExtensions: normalizedExtensions(recipe.SupportedExtensions),
			RegistryName:        provenance.RegistryName,
			DigestVerified:      provenance.DigestVerified,
			SignatureStatus:     provenance.SignatureStatus,
		},
		Status:      "known",
		RuntimeMode: modeLabel(recipe.Mode),
		Checks:      checks,
		CheckIDs:    diagnosticsCheckIDs(checks),
		Counts:      counts,
		ExecutionReadiness: DiagnosticsExecution{
			RequestType:           readiness.RequestType,
			OverallStatus:         readiness.OverallStatus,
			ExecutionState:        readiness.ExecutionState,
			GateCount:             readiness.GateCount,
			RequiredGateCount:     readiness.RequiredGateCount,
			PendingGateCount:      readiness.PendingGateCount,
			BlockedGateCount:      readiness.BlockedGateCount,
			LaunchEnabled:         readiness.LaunchEnabled,
			BackendBindingReady:   readiness.BackendBindingReady,
			SafeForAIDiagnostics:  readiness.SafeForAIDiagnostics,
			BackendDetailsExposed: readiness.BackendDetailsExposed,
		},
		LaunchIntent: DiagnosticsLaunchIntent{
			RequestType:           launchIntent.RequestType,
			OverallStatus:         launchIntent.OverallStatus,
			WriteGateDecision:     launchIntent.WriteGateDecision,
			PortalRequired:        launchIntent.PortalRequired,
			SnapshotRequired:      launchIntent.SnapshotRequired,
			LaunchEnabled:         launchIntent.LaunchEnabled,
			ExecutionStarted:      launchIntent.ExecutionStarted,
			BackendDetailsExposed: launchIntent.BackendDetailsExposed,
		},
		ActionQueue: DiagnosticsActionQueue{
			RequestType:             actionQueue.RequestType,
			ActionCount:             actionQueue.ActionCount,
			PendingActionCount:      actionQueue.PendingActionCount,
			UserReviewRequiredCount: actionQueue.UserReviewRequiredCount,
			RuntimeGateActionCount:  actionQueue.RuntimeGateActionCount,
			ActionQueueCreated:      actionQueue.ActionQueueCreated,
			ActionQueuePersisted:    actionQueue.ActionQueuePersisted,
			ExecutionStarted:        actionQueue.ExecutionStarted,
			BackendDetailsExposed:   actionQueue.BackendDetailsExposed,
		},
		CompatibilityCenterSummary: DiagnosticsCenterSummary{
			RequestType:                center.SummaryType,
			CompatibilityState:         centerApp.CompatibilityState,
			DiagnosticsState:           centerApp.DiagnosticsState,
			KnownIssueCount:            centerApp.KnownIssueCount,
			RepairRecordState:          centerApp.RepairRecordState,
			ActionExecutionEnabled:     centerApp.ActionExecutionEnabled,
			RepairExecutionEnabled:     centerApp.RepairExecutionEnabled,
			BackendLaunchEnabled:       centerApp.BackendLaunchEnabled,
			SettingsPersistenceEnabled: centerApp.SettingsPersistenceEnabled,
			BackendDetailsExposed:      centerApp.BackendDetailsExposed,
		},
		KDECenterSections: DiagnosticsKDESections{
			RequestType:           sections.RequestType,
			SectionCount:          sections.SectionCount,
			ReadOnlySectionCount:  sections.ReadOnlySectionCount,
			SectionIDs:            diagnosticsKDESectionIDs(sections.Sections),
			RuntimeMethods:        diagnosticsKDERuntimeMethods(sections.Sections),
			SectionActionsEnabled: sections.SectionActionsEnabled,
			RequestObjectsCreated: sections.RequestObjectsCreated,
			ExecutionStarted:      sections.ExecutionStarted,
			BackendDetailsExposed: sections.BackendDetailsExposed,
		},
		AI: DiagnosticsAISummary{
			RuntimeMethod:         "GetAIDiagnosticInput",
			AnalysisTask:          "compatibility-status-review",
			Disclosure:            "runtime-metadata-only",
			SafeForAIDiagnostics:  true,
			AIProviderCallEnabled: false,
			NetworkRequired:       false,
			FileContentRead:       false,
			FilePathsExposed:      false,
			RequestObjectCreated:  false,
			PermissionGranted:     false,
		},
		RuntimeOwned:            true,
		GoRuntimeBacked:         true,
		KDEPolicyOwner:          false,
		UserVisible:             true,
		CompatibilityCenterCard: true,
		SafeForAIDiagnostics:    true,
		AIProviderCallEnabled:   false,
		FileContentRead:         false,
		FilePathsExposed:        false,
		RequestObjectCreated:    false,
		PermissionGranted:       false,
		LaunchEnabled:           false,
		ExecutionStarted:        false,
		RepairExecutionEnabled:  false,
		SettingsPersisted:       false,
		HostRootModified:        false,
		NetworkRequired:         false,
		BackendDetailsExposed:   false,
		BlockedActions: []string{
			"start compatibility execution from diagnostics preview",
			"create launch request from diagnostics preview",
			"persist action queue from diagnostics preview",
			"call AI provider from diagnostics preview",
			"read selected file contents from diagnostics preview",
			"mutate host root during diagnostics preview",
		},
		DesktopSafeSummary: "Diagnostics preview is Go-owned and summarizes Runtime-safe application status while execution, repair, AI calls, and host mutation remain disabled.",
	}
	if err := validateNoBackendTerms(preview, "Diagnostics preview"); err != nil {
		return DiagnosticsPreview{}, err
	}
	return preview, nil
}

func diagnosticsCheck(id string, status string, summary string) DiagnosticsCheck {
	return DiagnosticsCheck{ID: id, Status: status, Summary: summary}
}

func countDiagnosticsChecks(checks []DiagnosticsCheck) DiagnosticsCounts {
	counts := DiagnosticsCounts{Total: len(checks)}
	for _, check := range checks {
		switch check.Status {
		case "pass":
			counts.Passed++
		case "pending":
			counts.Pending++
		case "blocked":
			counts.Blocked++
		}
	}
	return counts
}

func diagnosticsCheckIDs(checks []DiagnosticsCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func diagnosticsKDESectionIDs(sections []KDECenterPageSection) []string {
	ids := make([]string, 0, len(sections))
	for _, section := range sections {
		ids = append(ids, section.ID)
	}
	return ids
}

func diagnosticsKDERuntimeMethods(sections []KDECenterPageSection) []string {
	methods := make([]string, 0, len(sections))
	for _, section := range sections {
		methods = append(methods, section.RuntimeMethod)
	}
	return methods
}
