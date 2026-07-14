package appidentity

import "errors"

type ExecutionTransactionPreview struct {
	SchemaVersion               string                             `json:"schema_version"`
	RequestType                 string                             `json:"request_type"`
	TransactionType             string                             `json:"transaction_type"`
	RequestState                string                             `json:"request_state"`
	Source                      string                             `json:"source"`
	Desktop                     string                             `json:"desktop"`
	RuntimeMethod               string                             `json:"runtime_method"`
	ReadMethod                  string                             `json:"read_method"`
	ApplicationID               string                             `json:"application_id"`
	ApplicationName             string                             `json:"application_name"`
	Icon                        string                             `json:"icon"`
	DesktopFile                 string                             `json:"desktop_file"`
	LauncherCommand             []string                           `json:"launcher_command"`
	ResourceGrant               ExecutionTransactionResourceGrant  `json:"resource_grant"`
	Readiness                   ExecutionTransactionReadiness      `json:"readiness"`
	BackendBinding              ExecutionTransactionBackendBinding `json:"backend_binding"`
	SnapshotBaseline            ExecutionTransactionSnapshot       `json:"snapshot_baseline"`
	WriteGate                   ExecutionTransactionWriteGate      `json:"write_gate"`
	TransactionSteps            []ExecutionTransactionStep         `json:"transaction_steps"`
	StepCount                   int                                `json:"step_count"`
	PassedStepCount             int                                `json:"passed_step_count"`
	RequiredStepCount           int                                `json:"required_step_count"`
	PendingStepCount            int                                `json:"pending_step_count"`
	BlockedStepCount            int                                `json:"blocked_step_count"`
	FileCount                   int                                `json:"file_count"`
	FileURIs                    []string                           `json:"file_uris"`
	RuntimeOwned                bool                               `json:"runtime_owned"`
	GoRuntimeBacked             bool                               `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                               `json:"kde_policy_owner"`
	CompatibilityCenterCard     bool                               `json:"compatibility_center_card"`
	SafeForAIDiagnostics        bool                               `json:"safe_for_ai_diagnostics"`
	DesktopEntryLaunchVisible   bool                               `json:"desktop_entry_launch_visible"`
	LaunchIntentCaptured        bool                               `json:"launch_intent_captured"`
	UserDecisionCaptured        bool                               `json:"user_decision_captured"`
	UserDecisionAllowsLaunch    bool                               `json:"user_decision_allows_launch"`
	TransactionPlanCreated      bool                               `json:"transaction_plan_created"`
	TransactionCommitted        bool                               `json:"transaction_committed"`
	ReadinessPassed             bool                               `json:"readiness_passed"`
	ResourceGrantsCommitted     bool                               `json:"resource_grants_committed"`
	PortalApprovalRecorded      bool                               `json:"portal_approval_recorded"`
	SnapshotBaselineCreated     bool                               `json:"snapshot_baseline_created"`
	BackendBindingCommitted     bool                               `json:"backend_binding_committed"`
	RuntimeLaunchApproval       bool                               `json:"runtime_launch_approval"`
	LaunchAllowed               bool                               `json:"launch_allowed"`
	LaunchEnabled               bool                               `json:"launch_enabled"`
	ExecutionStarted            bool                               `json:"execution_started"`
	RequestObjectsCreated       bool                               `json:"request_objects_created"`
	HostPermissionChanged       bool                               `json:"host_permission_changed"`
	HostRootModified            bool                               `json:"host_root_modified"`
	NetworkRequired             bool                               `json:"network_required"`
	PrivilegedContainerRequired bool                               `json:"privileged_container_required"`
	BackendDetailsExposed       bool                               `json:"backend_details_exposed"`
	BlockedActions              []string                           `json:"blocked_actions"`
	UserFacingSettings          map[string]string                  `json:"user_facing_settings"`
	DesktopSafeSummary          string                             `json:"desktop_safe_summary"`
}

type ExecutionTransactionResourceGrant struct {
	RequestType            string `json:"request_type"`
	GrantType              string `json:"grant_type"`
	GrantCount             int    `json:"grant_count"`
	PortalRequiredCount    int    `json:"portal_required_count"`
	PendingReviewCount     int    `json:"pending_review_count"`
	GrantPlanCreated       bool   `json:"grant_plan_created"`
	GrantObjectsCreated    bool   `json:"grant_objects_created"`
	PermissionGranted      bool   `json:"permission_granted"`
	ResourceBridgesEnabled bool   `json:"resource_bridges_enabled"`
	PortalReviewRequired   bool   `json:"portal_review_required"`
}

type ExecutionTransactionReadiness struct {
	RequestType          string `json:"request_type"`
	ReadinessType        string `json:"readiness_type"`
	ExecutionState       string `json:"execution_state"`
	OverallStatus        string `json:"overall_status"`
	GateCount            int    `json:"gate_count"`
	RequiredGateCount    int    `json:"required_gate_count"`
	PendingGateCount     int    `json:"pending_gate_count"`
	BlockedGateCount     int    `json:"blocked_gate_count"`
	LaunchAllowed        bool   `json:"launch_allowed"`
	LaunchEnabled        bool   `json:"launch_enabled"`
	BackendBindingReady  bool   `json:"backend_binding_ready"`
	SnapshotRequired     bool   `json:"snapshot_required"`
	PortalPolicyRequired bool   `json:"portal_policy_required"`
}

type ExecutionTransactionBackendBinding struct {
	RecommendedProfileID  string `json:"recommended_profile_id"`
	SelectedStrategy      string `json:"selected_strategy"`
	CandidateCount        int    `json:"candidate_count"`
	ReadyCandidateCount   int    `json:"ready_candidate_count"`
	BlockedCandidateCount int    `json:"blocked_candidate_count"`
	BindingRequired       bool   `json:"binding_required"`
	BindingCommitted      bool   `json:"binding_committed"`
	EnvironmentCreated    bool   `json:"environment_created"`
	BackendLaunchEnabled  bool   `json:"backend_launch_enabled"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

type ExecutionTransactionSnapshot struct {
	Required              bool   `json:"required"`
	State                 string `json:"state"`
	Scope                 string `json:"scope"`
	BaselineCreated       bool   `json:"baseline_created"`
	RestorePointAvailable bool   `json:"restore_point_available"`
	UserDocumentsIncluded bool   `json:"user_documents_included"`
	HostSystemIncluded    bool   `json:"host_system_included"`
}

type ExecutionTransactionWriteGate struct {
	MethodName           string `json:"method_name"`
	GateDecision         string `json:"gate_decision"`
	DenialErrorName      string `json:"denial_error_name"`
	WriteMethodEnabled   bool   `json:"write_method_enabled"`
	DispatchEnabled      bool   `json:"dispatch_enabled"`
	RequestObjectCreated bool   `json:"request_object_created"`
}

type ExecutionTransactionStep struct {
	ID                  string `json:"id"`
	Status              string `json:"status"`
	Required            bool   `json:"required"`
	RuntimeGateRequired bool   `json:"runtime_gate_required"`
	BlocksCommit        bool   `json:"blocks_commit"`
	Summary             string `json:"summary"`
}

func (plan Plan) ExecutionTransactionPreview(decision string, fileURIs []string) (ExecutionTransactionPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return ExecutionTransactionPreview{}, err
	}
	if !singleLine(decision) {
		return ExecutionTransactionPreview{}, errors.New("execution transaction preview requires a single-line decision")
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return ExecutionTransactionPreview{}, errors.New("execution transaction preview requires single-line identity fields")
		}
	}

	resourceGrant, err := plan.ExecutionResourceGrantPreview(decision, fileURIs)
	if err != nil {
		return ExecutionTransactionPreview{}, err
	}
	readiness, err := plan.ExecutionReadinessPreview()
	if err != nil {
		return ExecutionTransactionPreview{}, err
	}
	backendSelection, err := plan.BackendSelectionPreview()
	if err != nil {
		return ExecutionTransactionPreview{}, err
	}

	steps := executionTransactionSteps(decision)
	passedStepCount, requiredStepCount, pendingStepCount, blockedStepCount := countExecutionTransactionSteps(steps)
	preview := ExecutionTransactionPreview{
		SchemaVersion:   "xnix.runtime.launch_transaction.v1",
		RequestType:     "execution-transaction-preview",
		TransactionType: "compatibility-launch-transaction",
		RequestState:    resourceGrant.RequestState,
		Source:          "execution-resource-grant-preview",
		Desktop:         "KDE Plasma",
		RuntimeMethod:   "Launch",
		ReadMethod:      "GetExecutionTransactionPreview",
		ApplicationID:   plan.ApplicationID,
		ApplicationName: plan.DisplayName,
		Icon:            plan.Icon,
		DesktopFile:     plan.DesktopFile,
		LauncherCommand: plan.LaunchCommand,
		ResourceGrant: ExecutionTransactionResourceGrant{
			RequestType:            resourceGrant.RequestType,
			GrantType:              resourceGrant.GrantType,
			GrantCount:             resourceGrant.GrantCount,
			PortalRequiredCount:    resourceGrant.PortalRequiredCount,
			PendingReviewCount:     resourceGrant.PendingReviewCount,
			GrantPlanCreated:       resourceGrant.GrantPlanCreated,
			GrantObjectsCreated:    resourceGrant.GrantObjectsCreated,
			PermissionGranted:      resourceGrant.PermissionGranted,
			ResourceBridgesEnabled: resourceGrant.ResourceBridgesEnabled,
			PortalReviewRequired:   resourceGrant.PortalReviewRequired,
		},
		Readiness: ExecutionTransactionReadiness{
			RequestType:          readiness.RequestType,
			ReadinessType:        readiness.ReadinessType,
			ExecutionState:       readiness.ExecutionState,
			OverallStatus:        readiness.OverallStatus,
			GateCount:            readiness.GateCount,
			RequiredGateCount:    readiness.RequiredGateCount,
			PendingGateCount:     readiness.PendingGateCount,
			BlockedGateCount:     readiness.BlockedGateCount,
			LaunchAllowed:        readiness.LaunchAllowed,
			LaunchEnabled:        readiness.LaunchEnabled,
			BackendBindingReady:  readiness.BackendBindingReady,
			SnapshotRequired:     readiness.SnapshotRequired,
			PortalPolicyRequired: readiness.PortalPolicyRequired,
		},
		BackendBinding: ExecutionTransactionBackendBinding{
			RecommendedProfileID:  backendSelection.RecommendedProfileID,
			SelectedStrategy:      backendSelection.SelectedStrategy,
			CandidateCount:        backendSelection.CandidateCount,
			ReadyCandidateCount:   backendSelection.ReadyCandidateCount,
			BlockedCandidateCount: backendSelection.BlockedCandidateCount,
			BindingRequired:       true,
			BindingCommitted:      false,
			EnvironmentCreated:    false,
			BackendLaunchEnabled:  false,
			BackendDetailsExposed: false,
		},
		SnapshotBaseline: ExecutionTransactionSnapshot{
			Required:              true,
			State:                 "required",
			Scope:                 "application-state-and-runtime-metadata",
			BaselineCreated:       false,
			RestorePointAvailable: false,
			UserDocumentsIncluded: false,
			HostSystemIncluded:    false,
		},
		WriteGate: ExecutionTransactionWriteGate{
			MethodName:           "Launch",
			GateDecision:         "blocked-until-production-backend",
			DenialErrorName:      "org.xnix.Compatibility1.Error.WriteMethodDisabled",
			WriteMethodEnabled:   false,
			DispatchEnabled:      false,
			RequestObjectCreated: false,
		},
		TransactionSteps:            steps,
		StepCount:                   len(steps),
		PassedStepCount:             passedStepCount,
		RequiredStepCount:           requiredStepCount,
		PendingStepCount:            pendingStepCount,
		BlockedStepCount:            blockedStepCount,
		FileCount:                   resourceGrant.FileCount,
		FileURIs:                    resourceGrant.FileURIs,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		CompatibilityCenterCard:     true,
		SafeForAIDiagnostics:        true,
		DesktopEntryLaunchVisible:   true,
		LaunchIntentCaptured:        resourceGrant.LaunchIntentCaptured,
		UserDecisionCaptured:        resourceGrant.UserDecisionCaptured,
		UserDecisionAllowsLaunch:    resourceGrant.UserDecisionAllowsLaunch,
		TransactionPlanCreated:      true,
		TransactionCommitted:        false,
		ReadinessPassed:             false,
		ResourceGrantsCommitted:     false,
		PortalApprovalRecorded:      false,
		SnapshotBaselineCreated:     false,
		BackendBindingCommitted:     false,
		RuntimeLaunchApproval:       false,
		LaunchAllowed:               false,
		LaunchEnabled:               false,
		ExecutionStarted:            false,
		RequestObjectsCreated:       false,
		HostPermissionChanged:       false,
		HostRootModified:            false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		BlockedActions:              []string{"commit launch transaction from preview", "create execution request objects from transaction preview", "record Portal approval from transaction preview", "create restore point from transaction preview", "commit compatibility profile binding from transaction preview", "start compatibility profile from transaction preview", "mutate host root during launch transaction planning", "expose raw backend command to desktop shell"},
		UserFacingSettings:          plan.UserFacingSettings,
		DesktopSafeSummary:          "KDE can show the launch transaction plan, but the Runtime does not commit the transaction or start the application.",
	}
	if err := validateNoBackendTerms(preview, "execution transaction preview"); err != nil {
		return ExecutionTransactionPreview{}, err
	}
	return preview, nil
}

func executionTransactionSteps(decision string) []ExecutionTransactionStep {
	userDecisionStatus := "pass"
	userDecisionSummary := "User launch review intent can be evaluated by Runtime gates."
	if decision == "rejected" {
		userDecisionStatus = "blocked"
		userDecisionSummary = "User rejected this launch; the transaction must remain blocked."
	} else if decision == "deferred" {
		userDecisionStatus = "pending"
		userDecisionSummary = "User deferred this launch; the transaction must wait for later review."
	}

	return []ExecutionTransactionStep{
		executionTransactionStep("identity-validation", "pass", true, false, false, "Application identity and desktop entry data are available."),
		executionTransactionStep("user-decision", userDecisionStatus, true, true, userDecisionStatus != "pass", userDecisionSummary),
		executionTransactionStep("portal-resource-review", "required", true, true, true, "Desktop resource access requires user-mediated Portal review."),
		executionTransactionStep("resource-grant-objects", "pending", true, true, true, "Resource grants remain pending until Portal review is recorded."),
		executionTransactionStep("snapshot-baseline", "required", true, true, true, "Runtime restore baseline must exist before launch commit."),
		executionTransactionStep("backend-binding", "pending", true, true, true, "A managed compatibility profile binding is required before launch commit."),
		executionTransactionStep("runtime-launch-write-gate", "blocked", true, true, true, "Runtime Launch dispatch remains disabled until production compatibility ownership is ready."),
	}
}

func executionTransactionStep(id string, status string, required bool, runtimeGateRequired bool, blocksCommit bool, summary string) ExecutionTransactionStep {
	return ExecutionTransactionStep{
		ID:                  id,
		Status:              status,
		Required:            required,
		RuntimeGateRequired: runtimeGateRequired,
		BlocksCommit:        blocksCommit,
		Summary:             summary,
	}
}

func countExecutionTransactionSteps(steps []ExecutionTransactionStep) (int, int, int, int) {
	passedStepCount := 0
	requiredStepCount := 0
	pendingStepCount := 0
	blockedStepCount := 0
	for _, step := range steps {
		switch step.Status {
		case "pass":
			passedStepCount++
		case "required":
			requiredStepCount++
		case "pending":
			pendingStepCount++
		case "blocked":
			blockedStepCount++
		}
	}
	return passedStepCount, requiredStepCount, pendingStepCount, blockedStepCount
}
