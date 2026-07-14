package appidentity

import (
	"errors"
)

type ExecutionReadinessPreview struct {
	SchemaVersion             string                   `json:"schema_version"`
	RequestType               string                   `json:"request_type"`
	ReadinessType             string                   `json:"readiness_type"`
	Source                    string                   `json:"source"`
	Desktop                   string                   `json:"desktop"`
	RuntimeMethod             string                   `json:"runtime_method"`
	Application               ReadinessApplication     `json:"application"`
	CompatibilityProfile      ReadinessProfile         `json:"compatibility_profile"`
	ExecutionState            string                   `json:"execution_state"`
	OverallStatus             string                   `json:"overall_status"`
	RecommendedAction         string                   `json:"recommended_action"`
	RuntimeOwned              bool                     `json:"runtime_owned"`
	GoRuntimeBacked           bool                     `json:"go_runtime_backed"`
	KDEPolicyOwner            bool                     `json:"kde_policy_owner"`
	CompatibilityCenterCard   bool                     `json:"compatibility_center_card"`
	SafeForAIDiagnostics      bool                     `json:"safe_for_ai_diagnostics"`
	DesktopEntryLaunchVisible bool                     `json:"desktop_entry_launch_visible"`
	LaunchAllowed             bool                     `json:"launch_allowed"`
	LaunchEnabled             bool                     `json:"launch_enabled"`
	ExecutionRequestCreated   bool                     `json:"execution_request_created"`
	BackendBindingReady       bool                     `json:"backend_binding_ready"`
	PortalPolicyRequired      bool                     `json:"portal_policy_required"`
	SnapshotRequired          bool                     `json:"snapshot_required"`
	UserActionRequired        bool                     `json:"user_action_required"`
	Gates                     []ExecutionReadinessGate `json:"gates"`
	GateCount                 int                      `json:"gate_count"`
	RequiredGateCount         int                      `json:"required_gate_count"`
	PendingGateCount          int                      `json:"pending_gate_count"`
	BlockedGateCount          int                      `json:"blocked_gate_count"`
	BlockedActions            []string                 `json:"blocked_actions"`
	HostRootModified          bool                     `json:"host_root_modified"`
	NetworkRequired           bool                     `json:"network_required"`
	BackendDetailsExposed     bool                     `json:"backend_details_exposed"`
	UserFacingSettings        map[string]string        `json:"user_facing_settings"`
	DesktopSafeSummary        string                   `json:"desktop_safe_summary"`
}

type ReadinessApplication struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Icon            string   `json:"icon"`
	RuntimeMode     string   `json:"runtime_mode"`
	DesktopFile     string   `json:"desktop_file"`
	LauncherCommand []string `json:"launcher_command"`
}

type ReadinessProfile struct {
	ID                    string   `json:"id"`
	Label                 string   `json:"label"`
	Kind                  string   `json:"kind"`
	Ready                 bool     `json:"ready"`
	LaunchEnabled         bool     `json:"launch_enabled"`
	RequiredPreflight     []string `json:"required_preflight"`
	BackendDetailsExposed bool     `json:"backend_details_exposed"`
}

type ExecutionReadinessGate struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

func (plan Plan) ExecutionReadinessPreview() (ExecutionReadinessPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return ExecutionReadinessPreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return ExecutionReadinessPreview{}, errors.New("execution readiness preview requires single-line identity fields")
		}
	}

	profileID := plan.recommendedCompatibilityProfileID()
	gates := []ExecutionReadinessGate{
		executionReadinessGate("recipe-validation", "pass", "Recipe metadata is loaded and can produce a desktop-safe execution plan."),
		executionReadinessGate("portal-policy-review", "required", "Portal-mediated resource access must be reviewed before execution."),
		executionReadinessGate("snapshot-baseline", "required", "A Runtime-managed restore point must exist before compatibility execution."),
		executionReadinessGate("backend-binding", "pending", "A managed compatibility profile binding is required before launch."),
		executionReadinessGate("runtime-launch-write-gate", "blocked", "Runtime Launch remains disabled until production compatibility ownership is ready."),
	}
	requiredGateCount, pendingGateCount, blockedGateCount := countExecutionReadinessGates(gates)

	preview := ExecutionReadinessPreview{
		SchemaVersion: "xnix.runtime.launch_readiness.v1",
		RequestType:   "execution-readiness-preview",
		ReadinessType: "compatibility-execution-readiness",
		Source:        "compatibility-center",
		Desktop:       "KDE Plasma",
		RuntimeMethod: "GetExecutionReadiness",
		Application: ReadinessApplication{
			ID:              plan.ApplicationID,
			Name:            plan.DisplayName,
			Icon:            plan.Icon,
			RuntimeMode:     plan.runtimeModeID(),
			DesktopFile:     plan.DesktopFile,
			LauncherCommand: plan.LaunchCommand,
		},
		CompatibilityProfile: ReadinessProfile{
			ID:                    profileID,
			Label:                 readinessProfileLabel(profileID),
			Kind:                  readinessProfileKind(profileID),
			Ready:                 false,
			LaunchEnabled:         false,
			RequiredPreflight:     []string{"portal-policy-review", "snapshot-baseline", "backend-binding", "runtime-launch-write-gate"},
			BackendDetailsExposed: false,
		},
		ExecutionState:            "blocked",
		OverallStatus:             "not-ready",
		RecommendedAction:         "Complete Runtime-owned compatibility, Portal, snapshot, and Launch gates before execution.",
		RuntimeOwned:              true,
		GoRuntimeBacked:           true,
		KDEPolicyOwner:            false,
		CompatibilityCenterCard:   true,
		SafeForAIDiagnostics:      true,
		DesktopEntryLaunchVisible: true,
		LaunchAllowed:             false,
		LaunchEnabled:             false,
		ExecutionRequestCreated:   false,
		BackendBindingReady:       false,
		PortalPolicyRequired:      true,
		SnapshotRequired:          true,
		UserActionRequired:        true,
		Gates:                     gates,
		GateCount:                 len(gates),
		RequiredGateCount:         requiredGateCount,
		PendingGateCount:          pendingGateCount,
		BlockedGateCount:          blockedGateCount,
		BlockedActions:            []string{"create execution request", "launch compatibility profile", "expose raw backend command to desktop shell", "grant desktop resources without Portal review", "mutate host root during execution readiness planning"},
		HostRootModified:          false,
		NetworkRequired:           false,
		BackendDetailsExposed:     false,
		UserFacingSettings:        plan.UserFacingSettings,
		DesktopSafeSummary:        "Compatibility execution is not ready; KDE may show the desktop entry but must route launch intent through Runtime gates.",
	}
	if err := validateNoBackendTerms(preview, "execution readiness preview"); err != nil {
		return ExecutionReadinessPreview{}, err
	}
	return preview, nil
}

func executionReadinessGate(id string, status string, summary string) ExecutionReadinessGate {
	return ExecutionReadinessGate{
		ID:      id,
		Status:  status,
		Summary: summary,
	}
}

func countExecutionReadinessGates(gates []ExecutionReadinessGate) (int, int, int) {
	requiredGateCount := 0
	pendingGateCount := 0
	blockedGateCount := 0
	for _, gate := range gates {
		switch gate.Status {
		case "required":
			requiredGateCount++
		case "pending":
			pendingGateCount++
		case "blocked":
			blockedGateCount++
		}
	}
	return requiredGateCount, pendingGateCount, blockedGateCount
}

func readinessProfileLabel(profileID string) string {
	if profileID == "isolated-compatibility" {
		return "Isolated compatibility"
	}
	return "Local compatibility"
}

func readinessProfileKind(profileID string) string {
	if profileID == "isolated-compatibility" {
		return "isolated"
	}
	return "local"
}
