package appidentity

import "strings"

type AIDiagnosticInputPreview struct {
	SchemaVersion         string                       `json:"schema_version"`
	RequestType           string                       `json:"request_type"`
	InputType             string                       `json:"input_type"`
	Source                string                       `json:"source"`
	Desktop               string                       `json:"desktop"`
	RuntimeMethod         string                       `json:"runtime_method"`
	ReadMethod            string                       `json:"read_method"`
	Application           AIDiagnosticApplication      `json:"application"`
	Issue                 string                       `json:"issue"`
	TestType              string                       `json:"test_type"`
	RuntimeOwned          bool                         `json:"runtime_owned"`
	GoRuntimeBacked       bool                         `json:"go_runtime_backed"`
	KDEPolicyOwner        bool                         `json:"kde_policy_owner"`
	AIProviderCalled      bool                         `json:"ai_provider_called"`
	AIProviderCallEnabled bool                         `json:"ai_provider_call_enabled"`
	NetworkRequired       bool                         `json:"network_required"`
	SafeForAIDiagnostics  bool                         `json:"safe_for_ai_diagnostics"`
	ContextSections       []AIDiagnosticContextSection `json:"context_sections"`
	DiagnosticSignals     []AIDiagnosticSignal         `json:"diagnostic_signals"`
	PrivacyBoundaries     AIDiagnosticPrivacy          `json:"privacy_boundaries"`
	AllowedAITasks        []string                     `json:"allowed_ai_tasks"`
	BlockedAITasks        []string                     `json:"blocked_ai_tasks"`
	FileContentRead       bool                         `json:"file_content_read"`
	FilePathsExposed      bool                         `json:"file_paths_exposed"`
	RequestObjectCreated  bool                         `json:"request_object_created"`
	PermissionGranted     bool                         `json:"permission_granted"`
	HostRootModified      bool                         `json:"host_root_modified"`
	BackendDetailsExposed bool                         `json:"backend_details_exposed"`
	DesktopSafeSummary    string                       `json:"desktop_safe_summary"`
}

type AIDiagnosticApplication struct {
	ID                  string   `json:"id"`
	Name                string   `json:"name"`
	Icon                string   `json:"icon"`
	Mode                string   `json:"mode"`
	SupportedExtensions []string `json:"supported_extensions"`
	MIMETypes           []string `json:"mime_types"`
}

type AIDiagnosticContextSection struct {
	ID      string         `json:"id"`
	Summary string         `json:"summary"`
	Facts   map[string]any `json:"facts"`
}

type AIDiagnosticSignal struct {
	ID                    string `json:"id"`
	Severity              string `json:"severity"`
	Source                string `json:"source"`
	Summary               string `json:"summary"`
	RecommendedNextAction string `json:"recommended_next_action"`
}

type AIDiagnosticPrivacy struct {
	UserDocumentsIncluded                 bool `json:"user_documents_included"`
	HostPathsIncluded                     bool `json:"host_paths_included"`
	RawBackendLogsIncluded                bool `json:"raw_backend_logs_included"`
	SecretsIncluded                       bool `json:"secrets_included"`
	NetworkCallsAllowed                   bool `json:"network_calls_allowed"`
	RequiresUserApprovalForSensitiveTasks bool `json:"requires_user_approval_for_sensitive_actions"`
}

type AIDiagnosticRecommendationPreview struct {
	SchemaVersion            string                       `json:"schema_version"`
	RequestType              string                       `json:"request_type"`
	RecommendationType       string                       `json:"recommendation_type"`
	InputType                string                       `json:"input_type"`
	Source                   string                       `json:"source"`
	Desktop                  string                       `json:"desktop"`
	RuntimeMethod            string                       `json:"runtime_method"`
	ReadMethod               string                       `json:"read_method"`
	Application              AIDiagnosticApplication      `json:"application"`
	RuntimeOwned             bool                         `json:"runtime_owned"`
	GoRuntimeBacked          bool                         `json:"go_runtime_backed"`
	KDEPolicyOwner           bool                         `json:"kde_policy_owner"`
	AIProviderCalled         bool                         `json:"ai_provider_called"`
	AIProviderCallEnabled    bool                         `json:"ai_provider_call_enabled"`
	NetworkRequired          bool                         `json:"network_required"`
	SafeForAIDiagnostics     bool                         `json:"safe_for_ai_diagnostics"`
	AutoExecutionAllowed     bool                         `json:"auto_execution_allowed"`
	Recommendations          []AIDiagnosticRecommendation `json:"recommendations"`
	ApprovalRequiredActions  []AIApprovalRequiredAction   `json:"approval_required_actions"`
	BlockedActions           []string                     `json:"blocked_actions"`
	PrivacyBoundaries        AIDiagnosticPrivacy          `json:"privacy_boundaries"`
	RepairExecutionRequested bool                         `json:"repair_execution_requested"`
	RepairExecuted           bool                         `json:"repair_executed"`
	HostRootModified         bool                         `json:"host_root_modified"`
	BackendDetailsExposed    bool                         `json:"backend_details_exposed"`
	DesktopSafeSummary       string                       `json:"desktop_safe_summary"`
}

type AIDiagnosticRecommendation struct {
	ID                   string `json:"id"`
	Priority             string `json:"priority"`
	SourceSignal         string `json:"source_signal"`
	Title                string `json:"title"`
	Summary              string `json:"summary"`
	UserVisible          bool   `json:"user_visible"`
	AutoExecute          bool   `json:"auto_execute"`
	RequiresUserApproval bool   `json:"requires_user_approval"`
	NextRuntimeAction    string `json:"next_runtime_action"`
}

type AIApprovalRequiredAction struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	Reason          string `json:"reason"`
	ApprovalSurface string `json:"approval_surface"`
}

type AIRepairApprovalGatePreview struct {
	SchemaVersion            string                     `json:"schema_version"`
	RequestType              string                     `json:"request_type"`
	GateType                 string                     `json:"gate_type"`
	RecommendationType       string                     `json:"recommendation_type"`
	Source                   string                     `json:"source"`
	Desktop                  string                     `json:"desktop"`
	Application              AIDiagnosticApplication    `json:"application"`
	RuntimeMethod            string                     `json:"runtime_method"`
	ReadMethod               string                     `json:"read_method"`
	RuntimeOwned             bool                       `json:"runtime_owned"`
	GoRuntimeBacked          bool                       `json:"go_runtime_backed"`
	KDEPolicyOwner           bool                       `json:"kde_policy_owner"`
	AIProviderCalled         bool                       `json:"ai_provider_called"`
	AIProviderCallEnabled    bool                       `json:"ai_provider_call_enabled"`
	NetworkRequired          bool                       `json:"network_required"`
	SafeForAIDiagnostics     bool                       `json:"safe_for_ai_diagnostics"`
	RepairExecutionRequested bool                       `json:"repair_execution_requested"`
	RepairExecuted           bool                       `json:"repair_executed"`
	AutoExecutionAllowed     bool                       `json:"auto_execution_allowed"`
	GateDecision             string                     `json:"gate_decision"`
	ApprovalSurface          string                     `json:"approval_surface"`
	RequiredGates            []AIRepairApprovalGateStep `json:"required_gates"`
	ApprovalRequiredActions  []AIApprovalRequiredAction `json:"approval_required_actions"`
	BlockedActions           []string                   `json:"blocked_actions"`
	HostRootModified         bool                       `json:"host_root_modified"`
	BackendDetailsExposed    bool                       `json:"backend_details_exposed"`
	DesktopSafeSummary       string                     `json:"desktop_safe_summary"`
}

type AIRepairApprovalGateStep struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

func (plan Plan) AIDiagnosticInputPreview(issue string, testType string) (AIDiagnosticInputPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return AIDiagnosticInputPreview{}, err
	}
	runPlan, err := plan.RunPlanPreview()
	if err != nil {
		return AIDiagnosticInputPreview{}, err
	}
	testResult, err := plan.TestResultPreview(testType)
	if err != nil {
		return AIDiagnosticInputPreview{}, err
	}
	repairPlan, err := plan.RepairPlanPreview(issue)
	if err != nil {
		return AIDiagnosticInputPreview{}, err
	}

	preview := AIDiagnosticInputPreview{
		SchemaVersion:         "xnix.runtime.ai_diagnostic_input.v1",
		RequestType:           "ai-diagnostic-input-preview",
		InputType:             "ai-diagnostic-input",
		Source:                "registry+go-runtime-ai-diagnostics",
		Desktop:               "KDE Plasma",
		RuntimeMethod:         "GetAIDiagnosticInput",
		ReadMethod:            "GetAIDiagnosticInputPreview",
		Application:           plan.aiDiagnosticApplication(),
		Issue:                 issue,
		TestType:              testType,
		RuntimeOwned:          true,
		GoRuntimeBacked:       true,
		KDEPolicyOwner:        false,
		AIProviderCalled:      false,
		AIProviderCallEnabled: false,
		NetworkRequired:       false,
		SafeForAIDiagnostics:  true,
		ContextSections:       plan.aiDiagnosticContextSections(runPlan, testResult, repairPlan),
		DiagnosticSignals:     aiDiagnosticSignals(testResult),
		PrivacyBoundaries:     aiDiagnosticPrivacy(),
		AllowedAITasks: []string{
			"summarize compatibility status",
			"explain pending Runtime work",
			"suggest safe next diagnostic steps",
			"prepare user-facing Compatibility Center text",
		},
		BlockedAITasks: []string{
			"start a compatibility backend",
			"read user documents",
			"change desktop permissions",
			"execute repair actions without Runtime approval",
		},
		FileContentRead:       false,
		FilePathsExposed:      false,
		RequestObjectCreated:  false,
		PermissionGranted:     false,
		HostRootModified:      false,
		BackendDetailsExposed: false,
		DesktopSafeSummary:    "AI diagnostics can explain current compatibility status using Runtime-safe metadata only.",
	}
	if err := validateNoBackendTerms(preview, "AI diagnostic input preview"); err != nil {
		return AIDiagnosticInputPreview{}, err
	}
	return preview, nil
}

func (plan Plan) AIDiagnosticRecommendationPreview(issue string, testType string) (AIDiagnosticRecommendationPreview, error) {
	input, err := plan.AIDiagnosticInputPreview(issue, testType)
	if err != nil {
		return AIDiagnosticRecommendationPreview{}, err
	}

	recommendations := aiDiagnosticRecommendations()
	preview := AIDiagnosticRecommendationPreview{
		SchemaVersion:            "xnix.runtime.ai_diagnostic_recommendation.v1",
		RequestType:              "ai-diagnostic-recommendation-preview",
		RecommendationType:       "ai-diagnostic-recommendation",
		InputType:                input.InputType,
		Source:                   "ai-diagnostic-input-preview+go-runtime-ai-diagnostics",
		Desktop:                  "KDE Plasma",
		RuntimeMethod:            "GetAIDiagnosticRecommendation",
		ReadMethod:               "GetAIDiagnosticRecommendationPreview",
		Application:              input.Application,
		RuntimeOwned:             true,
		GoRuntimeBacked:          true,
		KDEPolicyOwner:           false,
		AIProviderCalled:         false,
		AIProviderCallEnabled:    false,
		NetworkRequired:          false,
		SafeForAIDiagnostics:     true,
		AutoExecutionAllowed:     false,
		Recommendations:          recommendations,
		ApprovalRequiredActions:  aiApprovalRequiredActions(recommendations),
		BlockedActions:           input.BlockedAITasks,
		PrivacyBoundaries:        input.PrivacyBoundaries,
		RepairExecutionRequested: false,
		RepairExecuted:           false,
		HostRootModified:         false,
		BackendDetailsExposed:    false,
		DesktopSafeSummary:       "AI diagnostic recommendations are ready for review without executing compatibility changes.",
	}
	if err := validateNoBackendTerms(preview, "AI diagnostic recommendation preview"); err != nil {
		return AIDiagnosticRecommendationPreview{}, err
	}
	return preview, nil
}

func (plan Plan) AIRepairApprovalGatePreview(issue string, testType string) (AIRepairApprovalGatePreview, error) {
	recommendation, err := plan.AIDiagnosticRecommendationPreview(issue, testType)
	if err != nil {
		return AIRepairApprovalGatePreview{}, err
	}

	gateDecision := "review-only"
	if len(recommendation.ApprovalRequiredActions) > 0 {
		gateDecision = "blocked-until-approval"
	}
	preview := AIRepairApprovalGatePreview{
		SchemaVersion:            "xnix.runtime.ai_repair_approval_gate.v1",
		RequestType:              "ai-repair-approval-gate-preview",
		GateType:                 "ai-repair-approval-gate",
		RecommendationType:       recommendation.RecommendationType,
		Source:                   "ai-diagnostic-recommendation-preview+go-runtime-ai-diagnostics",
		Desktop:                  "KDE Plasma",
		Application:              recommendation.Application,
		RuntimeMethod:            "GetAIRepairApprovalGate",
		ReadMethod:               "GetAIRepairApprovalGatePreview",
		RuntimeOwned:             true,
		GoRuntimeBacked:          true,
		KDEPolicyOwner:           false,
		AIProviderCalled:         false,
		AIProviderCallEnabled:    false,
		NetworkRequired:          false,
		SafeForAIDiagnostics:     true,
		RepairExecutionRequested: false,
		RepairExecuted:           false,
		AutoExecutionAllowed:     false,
		GateDecision:             gateDecision,
		ApprovalSurface:          "Compatibility Center",
		RequiredGates:            aiRepairApprovalGates(),
		ApprovalRequiredActions:  recommendation.ApprovalRequiredActions,
		BlockedActions:           aiRepairBlockedActions(recommendation.BlockedActions),
		HostRootModified:         false,
		BackendDetailsExposed:    false,
		DesktopSafeSummary:       "AI repair recommendations are blocked from execution until Runtime approval gates pass.",
	}
	if err := validateNoBackendTerms(preview, "AI repair approval gate preview"); err != nil {
		return AIRepairApprovalGatePreview{}, err
	}
	return preview, nil
}

func (plan Plan) aiDiagnosticApplication() AIDiagnosticApplication {
	return AIDiagnosticApplication{
		ID:                  plan.ApplicationID,
		Name:                plan.DisplayName,
		Icon:                plan.Icon,
		Mode:                plan.RecipeMode,
		SupportedExtensions: plan.supportedExtensionsFromMIMETypes(),
		MIMETypes:           append([]string(nil), plan.MIMETypes...),
	}
}

func (plan Plan) supportedExtensionsFromMIMETypes() []string {
	extensions := make([]string, 0, len(plan.MIMETypes))
	for _, mimeType := range plan.MIMETypes {
		if strings.HasPrefix(mimeType, "application/x-xnix-") {
			extensions = append(extensions, "."+strings.TrimPrefix(mimeType, "application/x-xnix-"))
		}
	}
	return normalizedExtensions(extensions)
}

func (plan Plan) aiDiagnosticContextSections(runPlan RunPlanPreview, testResult TestResultPreview, repairPlan RepairPlanPreview) []AIDiagnosticContextSection {
	return []AIDiagnosticContextSection{
		{
			ID:      "recipe",
			Summary: "Application recipe metadata is available.",
			Facts: map[string]any{
				"application_id":         plan.ApplicationID,
				"display_name":           plan.DisplayName,
				"requested_mode":         plan.RecipeMode,
				"supported_extensions":   plan.supportedExtensionsFromMIMETypes(),
				"recipe_digest_verified": plan.RecipeDigestVerified,
			},
		},
		{
			ID:      "run-plan",
			Summary: runPlan.DesktopSafeSummary,
			Facts: map[string]any{
				"strategy":                runPlan.Execution.Strategy,
				"backend_binding_ready":   runPlan.Execution.BackendBinding.Ready,
				"launch_enabled":          runPlan.Execution.BackendBinding.LaunchEnabled,
				"backend_details_exposed": runPlan.Execution.BackendDetailsExposed,
			},
		},
		{
			ID:      "test-result",
			Summary: testResult.DesktopSafeSummary,
			Facts: map[string]any{
				"test_type":       testResult.TestType,
				"execution_state": testResult.ExecutionState,
				"overall_status":  testResult.OverallStatus,
				"counts":          testResult.Counts,
			},
		},
		{
			ID:      "repair-plan",
			Summary: repairPlan.DesktopSafeSummary,
			Facts: map[string]any{
				"issue":                  repairPlan.Issue,
				"severity":               repairPlan.Severity,
				"user_approval_required": repairPlan.UserApprovalRequired,
				"snapshot_required":      repairPlan.SnapshotRequired,
				"rollback_available":     repairPlan.RollbackAvailable,
			},
		},
	}
}

func aiDiagnosticSignals(testResult TestResultPreview) []AIDiagnosticSignal {
	return []AIDiagnosticSignal{
		{
			ID:                    "pending-runtime-launch-binding",
			Severity:              "medium",
			Source:                "run-plan",
			Summary:               "Managed compatibility launch binding is not ready yet.",
			RecommendedNextAction: "Keep automated smoke execution pending until a launch backend is bound.",
		},
		{
			ID:                    "pending-test-work",
			Severity:              "low",
			Source:                "test-result",
			Summary:               "Compatibility test steps are pending.",
			RecommendedNextAction: "Show pending Runtime-controlled preflight work in the Compatibility Center.",
		},
		{
			ID:                    "snapshot-before-risky-change",
			Severity:              "low",
			Source:                "repair-plan",
			Summary:               "A Runtime restore point is required before risky compatibility changes.",
			RecommendedNextAction: "Prepare a restore point before repair execution.",
		},
	}
}

func aiDiagnosticPrivacy() AIDiagnosticPrivacy {
	return AIDiagnosticPrivacy{
		UserDocumentsIncluded:                 false,
		HostPathsIncluded:                     false,
		RawBackendLogsIncluded:                false,
		SecretsIncluded:                       false,
		NetworkCallsAllowed:                   false,
		RequiresUserApprovalForSensitiveTasks: true,
	}
}

func aiDiagnosticRecommendations() []AIDiagnosticRecommendation {
	return []AIDiagnosticRecommendation{
		{
			ID:                   "explain-pending-runtime-work",
			Priority:             "high",
			SourceSignal:         "pending-runtime-launch-binding",
			Title:                "Explain pending Runtime launch binding",
			Summary:              "Tell the user that automated compatibility execution is waiting for a managed Runtime launch backend.",
			UserVisible:          true,
			AutoExecute:          false,
			RequiresUserApproval: false,
			NextRuntimeAction:    "Keep smoke execution pending until backend binding is available.",
		},
		{
			ID:                   "prepare-safe-restore-point",
			Priority:             "medium",
			SourceSignal:         "snapshot-before-risky-change",
			Title:                "Prepare a restore point before repair",
			Summary:              "Ask the Runtime to create an application-scoped restore point before risky compatibility changes.",
			UserVisible:          true,
			AutoExecute:          false,
			RequiresUserApproval: true,
			NextRuntimeAction:    "Create a Runtime restore point only after approval.",
		},
		{
			ID:                   "surface-test-progress",
			Priority:             "medium",
			SourceSignal:         "pending-test-work",
			Title:                "Show compatibility test progress",
			Summary:              "Display pending preflight work in the Compatibility Center without exposing backend details.",
			UserVisible:          true,
			AutoExecute:          false,
			RequiresUserApproval: false,
			NextRuntimeAction:    "Update the Compatibility Center card with pending test status.",
		},
	}
}

func aiApprovalRequiredActions(recommendations []AIDiagnosticRecommendation) []AIApprovalRequiredAction {
	actions := make([]AIApprovalRequiredAction, 0)
	for _, recommendation := range recommendations {
		if !recommendation.RequiresUserApproval {
			continue
		}
		actions = append(actions, AIApprovalRequiredAction{
			ID:              recommendation.ID,
			Title:           recommendation.Title,
			Reason:          recommendation.Summary,
			ApprovalSurface: "Compatibility Center",
		})
	}
	return actions
}

func aiRepairApprovalGates() []AIRepairApprovalGateStep {
	return []AIRepairApprovalGateStep{
		{
			ID:      "compatibility-center-review",
			Status:  "required",
			Summary: "A user-visible Compatibility Center review is required before any repair action.",
		},
		{
			ID:      "runtime-approval-token",
			Status:  "missing",
			Summary: "No Runtime approval token is present for repair execution.",
		},
		{
			ID:      "restore-point-preflight",
			Status:  "required",
			Summary: "A Runtime restore point must be prepared before risky compatibility repair.",
		},
	}
}

func aiRepairBlockedActions(existing []string) []string {
	seen := map[string]bool{}
	actions := make([]string, 0, len(existing)+3)
	for _, action := range append(existing,
		"execute repair without approval",
		"create restore point without approval",
		"change compatibility mode without Runtime gate",
	) {
		if seen[action] {
			continue
		}
		seen[action] = true
		actions = append(actions, action)
	}
	return actions
}
