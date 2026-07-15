package appidentity

import "errors"

type BackendBindingPreview struct {
	SchemaVersion               string                    `json:"schema_version"`
	RequestType                 string                    `json:"request_type"`
	BindingType                 string                    `json:"binding_type"`
	Source                      string                    `json:"source"`
	Desktop                     string                    `json:"desktop"`
	RuntimeMethod               string                    `json:"runtime_method"`
	ApplicationID               string                    `json:"application_id"`
	DisplayName                 string                    `json:"display_name"`
	Icon                        string                    `json:"icon"`
	DesktopFile                 string                    `json:"desktop_file"`
	SelectedStrategy            string                    `json:"selected_strategy"`
	RecommendedProfileID        string                    `json:"recommended_profile_id"`
	ProfileKind                 string                    `json:"profile_kind"`
	BindingState                string                    `json:"binding_state"`
	BindingKey                  string                    `json:"binding_key"`
	Selection                   BackendBindingSelection   `json:"selection"`
	Environment                 BackendBindingEnvironment `json:"environment"`
	PreflightGates              []BackendBindingGate      `json:"preflight_gates"`
	PreflightGateIDs            []string                  `json:"preflight_gate_ids"`
	PreflightGateCount          int                       `json:"preflight_gate_count"`
	PassedGateCount             int                       `json:"passed_gate_count"`
	PendingGateCount            int                       `json:"pending_gate_count"`
	BlockedGateCount            int                       `json:"blocked_gate_count"`
	RequiredReviews             []string                  `json:"required_reviews"`
	RuntimeOwned                bool                      `json:"runtime_owned"`
	GoRuntimeBacked             bool                      `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                      `json:"kde_policy_owner"`
	UserVisible                 bool                      `json:"user_visible"`
	ManagedBindingReady         bool                      `json:"managed_binding_ready"`
	BindingCommitted            bool                      `json:"binding_committed"`
	BindingPersisted            bool                      `json:"binding_persisted"`
	LaunchEnabled               bool                      `json:"launch_enabled"`
	ExecutionRequestCreated     bool                      `json:"execution_request_created"`
	EnvironmentCreated          bool                      `json:"environment_created"`
	BackendProcessStarted       bool                      `json:"backend_process_started"`
	StateRootCreated            bool                      `json:"state_root_created"`
	PortalRequestCreated        bool                      `json:"portal_request_created"`
	SnapshotCreated             bool                      `json:"snapshot_created"`
	HostRootModified            bool                      `json:"host_root_modified"`
	NetworkRequired             bool                      `json:"network_required"`
	PrivilegedContainerRequired bool                      `json:"privileged_container_required"`
	BackendDetailsExposed       bool                      `json:"backend_details_exposed"`
	CompatibilityStorageExposed bool                      `json:"compatibility_storage_exposed"`
	RawBackendCommandExposed    bool                      `json:"raw_backend_command_exposed"`
	BlockedActions              []string                  `json:"blocked_actions"`
	UserFacingSettings          map[string]string         `json:"user_facing_settings"`
	DesktopSafeSummary          string                    `json:"desktop_safe_summary"`
}

type BackendBindingSelection struct {
	RequestType           string `json:"request_type"`
	RecommendedProfileID  string `json:"recommended_profile_id"`
	CandidateCount        int    `json:"candidate_count"`
	SelectionCommitted    bool   `json:"selection_committed"`
	BackendLaunchEnabled  bool   `json:"backend_launch_enabled"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

type BackendBindingEnvironment struct {
	RequestType           string `json:"request_type"`
	EnvironmentState      string `json:"environment_state"`
	ProfileCount          int    `json:"profile_count"`
	BridgeCapabilityCount int    `json:"bridge_capability_count"`
	EnvironmentCreated    bool   `json:"environment_created"`
	BackendProcessStarted bool   `json:"backend_process_started"`
	BackendBindingReady   bool   `json:"backend_binding_ready"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

type BackendBindingGate struct {
	ID       string `json:"id"`
	Status   string `json:"status"`
	Required bool   `json:"required"`
	Summary  string `json:"summary"`
}

func (plan Plan) BackendBindingPreview() (BackendBindingPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return BackendBindingPreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return BackendBindingPreview{}, errors.New("backend binding preview requires single-line identity fields")
		}
	}

	selection, err := plan.BackendSelectionPreview()
	if err != nil {
		return BackendBindingPreview{}, err
	}
	environment, err := plan.BackendEnvironmentPreview()
	if err != nil {
		return BackendBindingPreview{}, err
	}
	gates := backendBindingGates()
	passedGateCount, pendingGateCount, blockedGateCount := countBackendBindingGates(gates)
	profileID := selection.RecommendedProfileID

	preview := BackendBindingPreview{
		SchemaVersion:        "xnix.runtime.backend_binding.v1",
		RequestType:          "backend-binding-preview",
		BindingType:          "compatibility-backend-binding",
		Source:               "backend-selection-preview+backend-environment-preview+execution-readiness-preview",
		Desktop:              "KDE Plasma",
		RuntimeMethod:        "GetBackendBinding",
		ApplicationID:        plan.ApplicationID,
		DisplayName:          plan.DisplayName,
		Icon:                 plan.Icon,
		DesktopFile:          plan.DesktopFile,
		SelectedStrategy:     selection.SelectedStrategy,
		RecommendedProfileID: profileID,
		ProfileKind:          readinessProfileKind(profileID),
		BindingState:         "planned-blocked",
		BindingKey:           plan.ApplicationID + ":" + profileID,
		Selection: BackendBindingSelection{
			RequestType:           selection.RequestType,
			RecommendedProfileID:  selection.RecommendedProfileID,
			CandidateCount:        selection.CandidateCount,
			SelectionCommitted:    selection.SelectionCommitted,
			BackendLaunchEnabled:  selection.BackendLaunchEnabled,
			BackendDetailsExposed: selection.BackendDetailsExposed,
		},
		Environment: BackendBindingEnvironment{
			RequestType:           environment.RequestType,
			EnvironmentState:      environment.EnvironmentState,
			ProfileCount:          environment.ProfileCount,
			BridgeCapabilityCount: environment.BridgeCapabilityCount,
			EnvironmentCreated:    environment.EnvironmentCreated,
			BackendProcessStarted: environment.BackendProcessStarted,
			BackendBindingReady:   environment.BackendBindingReady,
			BackendDetailsExposed: environment.BackendDetailsExposed,
		},
		PreflightGates:              gates,
		PreflightGateIDs:            backendBindingGateIDs(gates),
		PreflightGateCount:          len(gates),
		PassedGateCount:             passedGateCount,
		PendingGateCount:            pendingGateCount,
		BlockedGateCount:            blockedGateCount,
		RequiredReviews:             []string{"package-source-review", "application-state-root-review", "portal-policy-review", "snapshot-baseline-review", "runtime-launch-write-gate"},
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		UserVisible:                 true,
		ManagedBindingReady:         false,
		BindingCommitted:            false,
		BindingPersisted:            false,
		LaunchEnabled:               false,
		ExecutionRequestCreated:     false,
		EnvironmentCreated:          false,
		BackendProcessStarted:       false,
		StateRootCreated:            false,
		PortalRequestCreated:        false,
		SnapshotCreated:             false,
		HostRootModified:            false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		CompatibilityStorageExposed: false,
		RawBackendCommandExposed:    false,
		BlockedActions: []string{
			"commit compatibility profile binding from preview",
			"persist backend binding before Runtime preflight",
			"create execution request from backend binding preview",
			"create compatibility environment from backend binding preview",
			"start compatibility backend from backend binding preview",
			"create state root from backend binding preview",
			"create Portal request from backend binding preview",
			"create snapshot from backend binding preview",
			"expose backend command to desktop shell",
			"mutate host root during backend binding planning",
		},
		UserFacingSettings: plan.UserFacingSettings,
		DesktopSafeSummary: "Runtime can explain the selected compatibility profile binding, but the binding is not committed and cannot launch until package source, state, Portal, snapshot, and write gates pass.",
	}
	if err := validateNoBackendTerms(preview, "backend binding preview"); err != nil {
		return BackendBindingPreview{}, err
	}
	return preview, nil
}

func backendBindingGates() []BackendBindingGate {
	return []BackendBindingGate{
		backendBindingGate("recipe-validation", "pass", true, "Recipe metadata is loaded and can produce a desktop-safe binding plan."),
		backendBindingGate("package-source-review", "pending", true, "Runtime package source policy must approve the selected profile."),
		backendBindingGate("application-state-root", "pending", true, "Runtime state ownership must be allocated before binding."),
		backendBindingGate("portal-policy-review", "pending", true, "Portal-mediated desktop resource access must be reviewed before binding."),
		backendBindingGate("snapshot-baseline", "pending", true, "A restore point baseline is required before binding."),
		backendBindingGate("runtime-launch-write-gate", "blocked", true, "Runtime Launch remains disabled until production compatibility ownership is ready."),
	}
}

func backendBindingGate(id string, status string, required bool, summary string) BackendBindingGate {
	return BackendBindingGate{
		ID:       id,
		Status:   status,
		Required: required,
		Summary:  summary,
	}
}

func countBackendBindingGates(gates []BackendBindingGate) (int, int, int) {
	passed := 0
	pending := 0
	blocked := 0
	for _, gate := range gates {
		switch gate.Status {
		case "pass":
			passed++
		case "pending":
			pending++
		case "blocked":
			blocked++
		}
	}
	return passed, pending, blocked
}

func backendBindingGateIDs(gates []BackendBindingGate) []string {
	ids := make([]string, 0, len(gates))
	for _, gate := range gates {
		ids = append(ids, gate.ID)
	}
	return ids
}
