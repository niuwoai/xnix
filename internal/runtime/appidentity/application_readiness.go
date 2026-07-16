package appidentity

import (
	"errors"

	"xnix.local/xnix/internal/runtime/artifact"
)

type ApplicationReadinessOptions struct {
	Environment     string
	RuntimeRoot     string
	StateRoot       string
	ArtifactReceipt *artifact.StageReceipt
	PortalOperation string
	SnapshotReason  string
	WriteMethod     string
}

type ApplicationReadinessPreview struct {
	Version                     string                        `json:"version"`
	SchemaVersion               string                        `json:"schema_version"`
	RequestType                 string                        `json:"request_type"`
	GraphType                   string                        `json:"graph_type"`
	Source                      string                        `json:"source"`
	Desktop                     string                        `json:"desktop"`
	RuntimeMethod               string                        `json:"runtime_method"`
	ReadMethod                  string                        `json:"read_method"`
	Application                 ApplicationReadinessIdentity  `json:"application"`
	Environment                 string                        `json:"environment"`
	OverallStatus               string                        `json:"overall_status"`
	Ready                       bool                          `json:"ready"`
	RecommendedAction           string                        `json:"recommended_action"`
	Nodes                       []ApplicationReadinessNode    `json:"nodes"`
	NodeIDs                     []string                      `json:"node_ids"`
	NodeCount                   int                           `json:"node_count"`
	ReadyNodeCount              int                           `json:"ready_node_count"`
	RequiredNodeCount           int                           `json:"required_node_count"`
	BlockedNodeCount            int                           `json:"blocked_node_count"`
	RecipeTrustDecision         string                        `json:"recipe_trust_decision"`
	InstallReadiness            CompatibilityInstallReadiness `json:"install_readiness"`
	BackendLifecycleState       string                        `json:"backend_lifecycle_state"`
	PortalDecision              string                        `json:"portal_decision"`
	SnapshotReason              string                        `json:"snapshot_reason"`
	ExecutionState              string                        `json:"execution_state"`
	WriteGateDecision           string                        `json:"write_gate_decision"`
	RuntimeOwned                bool                          `json:"runtime_owned"`
	GoRuntimeBacked             bool                          `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                          `json:"kde_policy_owner"`
	UserVisible                 bool                          `json:"user_visible"`
	CompatibilityCenterCard     bool                          `json:"compatibility_center_card"`
	LaunchAllowed               bool                          `json:"launch_allowed"`
	LaunchEnabled               bool                          `json:"launch_enabled"`
	WriteMethodsEnabled         bool                          `json:"write_methods_enabled"`
	ExecutionRequestCreated     bool                          `json:"execution_request_created"`
	ExecutionStarted            bool                          `json:"execution_started"`
	BackendLaunchEnabled        bool                          `json:"backend_launch_enabled"`
	BackendProcessStarted       bool                          `json:"backend_process_started"`
	RealPortalTransportEnabled  bool                          `json:"real_portal_transport_enabled"`
	RequestObjectCreated        bool                          `json:"request_object_created"`
	PermissionGranted           bool                          `json:"permission_granted"`
	SnapshotCreated             bool                          `json:"snapshot_created"`
	RestoreExecuted             bool                          `json:"restore_executed"`
	HostRootModified            bool                          `json:"host_root_modified"`
	NetworkRequired             bool                          `json:"network_required"`
	PrivilegedContainerRequired bool                          `json:"privileged_container_required"`
	StateRootPathExposed        bool                          `json:"state_root_path_exposed"`
	BackendDetailsExposed       bool                          `json:"backend_details_exposed"`
	RawCommandExposed           bool                          `json:"raw_command_exposed"`
	RawExecutableExposed        bool                          `json:"raw_executable_exposed"`
	UserFacingSettings          map[string]string             `json:"user_facing_settings"`
	BlockedActions              []string                      `json:"blocked_actions"`
	DesktopSafeSummary          string                        `json:"desktop_safe_summary"`
}

type ApplicationReadinessIdentity struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	DesktopFile string `json:"desktop_file"`
	RuntimeMode string `json:"runtime_mode"`
}

type ApplicationReadinessNode struct {
	ID              string   `json:"id"`
	Source          string   `json:"source"`
	Status          string   `json:"status"`
	Ready           bool     `json:"ready"`
	Required        bool     `json:"required"`
	BlockingReasons []string `json:"blocking_reasons"`
	Summary         string   `json:"summary"`
}

func (plan Plan) ApplicationReadinessPreview(options ApplicationReadinessOptions) (ApplicationReadinessPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return ApplicationReadinessPreview{}, err
	}
	if options.Environment == "" {
		options.Environment = "development"
	}
	if options.RuntimeRoot == "" {
		options.RuntimeRoot = "."
	}
	if options.PortalOperation == "" {
		options.PortalOperation = "file-open"
	}
	if options.SnapshotReason == "" {
		options.SnapshotReason = "before-repair"
	}
	if options.WriteMethod == "" {
		options.WriteMethod = "Launch"
	}
	if options.WriteMethod != "Launch" {
		return ApplicationReadinessPreview{}, errors.New("application readiness preview currently supports Launch write-gate evaluation only")
	}

	install, err := plan.CompatibilityInstallPlanPreviewWithArtifactReceipt(options.Environment, options.ArtifactReceipt)
	if err != nil {
		return ApplicationReadinessPreview{}, err
	}
	backend, err := plan.BackendLifecyclePreviewWithStateRoot(options.StateRoot)
	if err != nil {
		return ApplicationReadinessPreview{}, err
	}
	portalPolicy, err := NewPortalAccessPolicyPreview(plan.ApplicationID, options.PortalOperation)
	if err != nil {
		return ApplicationReadinessPreview{}, err
	}
	snapshot, err := NewSnapshotPlanPreview(plan.ApplicationID, options.SnapshotReason)
	if err != nil {
		return ApplicationReadinessPreview{}, err
	}
	execution, err := plan.ExecutionReadinessPreview()
	if err != nil {
		return ApplicationReadinessPreview{}, err
	}
	writeGate, err := NewRuntimeWriteGatePreview(options.RuntimeRoot, options.WriteMethod)
	if err != nil {
		return ApplicationReadinessPreview{}, err
	}

	nodes := applicationReadinessNodes(install, backend, portalPolicy, snapshot, execution, writeGate)
	readyCount, requiredCount, blockedCount := countApplicationReadinessNodes(nodes)
	preview := ApplicationReadinessPreview{
		Version:       writeGate.Version,
		SchemaVersion: "xnix.runtime.application_readiness.v1",
		RequestType:   "application-readiness-preview",
		GraphType:     "runtime-application-readiness-evidence-graph",
		Source:        "go-runtime-application-readiness-evidence-graph",
		Desktop:       "KDE Plasma",
		RuntimeMethod: "GetApplicationReadiness",
		ReadMethod:    "GetApplicationReadinessPreview",
		Application: ApplicationReadinessIdentity{
			ID:          plan.ApplicationID,
			Name:        plan.DisplayName,
			Icon:        plan.Icon,
			DesktopFile: plan.DesktopFile,
			RuntimeMode: plan.runtimeModeID(),
		},
		Environment:                 install.Environment,
		OverallStatus:               "not-ready",
		Ready:                       false,
		RecommendedAction:           "Complete Runtime-owned artifact, backend lifecycle, Portal, snapshot, execution, and Launch write-gate evidence before launch.",
		Nodes:                       nodes,
		NodeIDs:                     applicationReadinessNodeIDs(nodes),
		NodeCount:                   len(nodes),
		ReadyNodeCount:              readyCount,
		RequiredNodeCount:           requiredCount,
		BlockedNodeCount:            blockedCount,
		RecipeTrustDecision:         install.Readiness.RecipeInstallDecision,
		InstallReadiness:            install.Readiness,
		BackendLifecycleState:       backend.LifecycleState,
		PortalDecision:              portalPolicy.Decision,
		SnapshotReason:              snapshot.Reason,
		ExecutionState:              execution.ExecutionState,
		WriteGateDecision:           writeGate.GateDecision,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		UserVisible:                 true,
		CompatibilityCenterCard:     true,
		LaunchAllowed:               false,
		LaunchEnabled:               false,
		WriteMethodsEnabled:         false,
		ExecutionRequestCreated:     false,
		ExecutionStarted:            false,
		BackendLaunchEnabled:        false,
		BackendProcessStarted:       false,
		RealPortalTransportEnabled:  false,
		RequestObjectCreated:        false,
		PermissionGranted:           false,
		SnapshotCreated:             false,
		RestoreExecuted:             false,
		HostRootModified:            false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		StateRootPathExposed:        false,
		BackendDetailsExposed:       false,
		RawCommandExposed:           false,
		RawExecutableExposed:        false,
		UserFacingSettings:          plan.UserFacingSettings,
		BlockedActions: []string{
			"create execution request before readiness evidence passes",
			"launch application before Runtime Launch write gate passes",
			"start compatibility backend from readiness preview",
			"create real Portal request from readiness preview",
			"create snapshot from readiness preview",
			"expose state-root paths or backend commands to KDE",
			"mutate host root during readiness inspection",
		},
		DesktopSafeSummary: "Application readiness evidence is visible to KDE, but launch remains blocked until Runtime-owned artifact, backend, Portal, snapshot, execution, and write-gate checks pass.",
	}
	if err := validateNoBackendTerms(preview, "application readiness preview"); err != nil {
		return ApplicationReadinessPreview{}, err
	}
	return preview, nil
}

func applicationReadinessNodes(install CompatibilityInstallPlanPreview, backend BackendLifecyclePreview, portal PortalAccessPolicyPreview, snapshot SnapshotPlanPreview, execution ExecutionReadinessPreview, writeGate RuntimeWriteGatePreview) []ApplicationReadinessNode {
	return []ApplicationReadinessNode{
		applicationReadinessNode("recipe-trust", install.Source, install.Readiness.RecipeInstallAllowed, true, install.Readiness.RecipeTrustBlockingReasons, "Recipe trust and install gate evidence must allow this environment."),
		applicationReadinessNode("artifact-stage-receipt", install.Source, install.Readiness.ArtifactStageReceiptReady && install.Readiness.RequiredArtifactsStaged, true, install.Readiness.ArtifactStageBlockingReasons, "A verified artifact staging receipt must cover required artifacts."),
		applicationReadinessNode("backend-lifecycle", backend.Source, backend.BackendBindingReady && backend.StateRootReady, true, backendReadinessBlockingReasons(backend), "Backend lifecycle state must be ready without exposing backend implementation details."),
		applicationReadinessNode("portal-review", portal.Source, portal.PermissionGranted, true, []string{"Portal permission receipt is required before sensitive desktop access."}, "Sensitive desktop resources require user-mediated Portal evidence."),
		applicationReadinessNode("snapshot-baseline", snapshot.Source, snapshot.SnapshotCreated, true, []string{"Runtime restore-point receipt is required before risky compatibility changes."}, "Snapshot evidence must exist before execution can become ready."),
		applicationReadinessNode("execution-readiness", execution.Source, execution.LaunchAllowed, true, executionReadinessBlockingReasons(execution), "Execution readiness must pass all Runtime-owned launch gates."),
		applicationReadinessNode("runtime-write-gate", writeGate.Source, writeGate.WriteMethodEnabled && writeGate.DispatchEnabled, true, writeGateBlockingReasons(writeGate), "Launch write gate must remain closed until production Runtime gates pass."),
	}
}

func applicationReadinessNode(id string, source string, ready bool, required bool, blockingReasons []string, summary string) ApplicationReadinessNode {
	status := "pass"
	if !ready && required {
		status = "required"
	}
	if id == "runtime-write-gate" && !ready {
		status = "blocked"
	}
	if id == "execution-readiness" && !ready {
		status = "blocked"
	}
	reasons := append([]string(nil), blockingReasons...)
	if !ready && len(reasons) == 0 {
		reasons = []string{"Runtime evidence is not ready."}
	}
	if ready {
		reasons = []string{}
	}
	return ApplicationReadinessNode{
		ID:              id,
		Source:          source,
		Status:          status,
		Ready:           ready,
		Required:        required,
		BlockingReasons: reasons,
		Summary:         summary,
	}
}

func backendReadinessBlockingReasons(backend BackendLifecyclePreview) []string {
	reasons := []string{}
	if backend.BlockReason != "" {
		reasons = append(reasons, backend.BlockReason)
	}
	reasons = append(reasons, backend.PendingGates...)
	if len(reasons) == 0 && !backend.BackendBindingReady {
		reasons = append(reasons, "Backend lifecycle and binding evidence are not ready.")
	}
	return reasons
}

func executionReadinessBlockingReasons(execution ExecutionReadinessPreview) []string {
	reasons := []string{}
	for _, gate := range execution.Gates {
		if gate.Status != "pass" {
			reasons = append(reasons, gate.ID)
		}
	}
	return reasons
}

func writeGateBlockingReasons(writeGate RuntimeWriteGatePreview) []string {
	reasons := make([]string, 0, len(writeGate.RequiredGateIDs)+1)
	reasons = append(reasons, writeGate.RequiredGateIDs...)
	reasons = append(reasons, writeGate.DenialErrorName)
	return reasons
}

func applicationReadinessNodeIDs(nodes []ApplicationReadinessNode) []string {
	ids := make([]string, 0, len(nodes))
	for _, node := range nodes {
		ids = append(ids, node.ID)
	}
	return ids
}

func countApplicationReadinessNodes(nodes []ApplicationReadinessNode) (int, int, int) {
	ready := 0
	required := 0
	blocked := 0
	for _, node := range nodes {
		if node.Ready {
			ready++
		}
		if node.Required {
			required++
		}
		if node.Status == "blocked" {
			blocked++
		}
	}
	return ready, required, blocked
}
