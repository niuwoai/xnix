package appidentity

import "errors"

type RunPlanPreview struct {
	SchemaVersion           string             `json:"schema_version"`
	RequestType             string             `json:"request_type"`
	PlanType                string             `json:"plan_type"`
	Source                  string             `json:"source"`
	Desktop                 string             `json:"desktop"`
	RuntimeMethod           string             `json:"runtime_method"`
	ReadMethod              string             `json:"read_method"`
	Application             RunPlanApplication `json:"application"`
	Execution               RunPlanExecution   `json:"execution"`
	Preflight               RunPlanPreflight   `json:"preflight"`
	RuntimeOwned            bool               `json:"runtime_owned"`
	GoRuntimeBacked         bool               `json:"go_runtime_backed"`
	KDEPolicyOwner          bool               `json:"kde_policy_owner"`
	UserVisible             bool               `json:"user_visible"`
	LaunchEnabled           bool               `json:"launch_enabled"`
	ExecutionRequestCreated bool               `json:"execution_request_created"`
	ExecutionStarted        bool               `json:"execution_started"`
	BackendBindingReady     bool               `json:"backend_binding_ready"`
	RequestObjectCreated    bool               `json:"request_object_created"`
	PermissionGranted       bool               `json:"permission_granted"`
	HostRootModified        bool               `json:"host_root_modified"`
	NetworkRequired         bool               `json:"network_required"`
	BackendDetailsExposed   bool               `json:"backend_details_exposed"`
	RawCommandExposed       bool               `json:"raw_command_exposed"`
	UserFacingSettings      map[string]string  `json:"user_facing_settings"`
	BlockedActions          []string           `json:"blocked_actions"`
	DesktopSafeSummary      string             `json:"desktop_safe_summary"`
}

type RunPlanApplication struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	DesktopFile string `json:"desktop_file"`
	RuntimeMode string `json:"runtime_mode"`
}

type RunPlanExecution struct {
	Strategy                string               `json:"strategy"`
	SelectionSource         string               `json:"selection_source"`
	SelectedEngine          RunPlanEngineSummary `json:"selected_engine"`
	SelectedProfile         RunPlanProfile       `json:"selected_profile"`
	BackendBinding          RunPlanBinding       `json:"backend_binding"`
	BackendDetailsExposed   bool                 `json:"backend_details_exposed"`
	LaunchEnabled           bool                 `json:"launch_enabled"`
	ExecutionRequestCreated bool                 `json:"execution_request_created"`
	ExecutionStarted        bool                 `json:"execution_started"`
}

type RunPlanEngineSummary struct {
	ID                    string `json:"id"`
	Label                 string `json:"label"`
	Availability          string `json:"availability"`
	Ready                 bool   `json:"ready"`
	LaunchEnabled         bool   `json:"launch_enabled"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

type RunPlanProfile struct {
	ID                    string   `json:"id"`
	Label                 string   `json:"label"`
	Kind                  string   `json:"kind"`
	RequiredPreflight     []string `json:"required_preflight"`
	Ready                 bool     `json:"ready"`
	LaunchEnabled         bool     `json:"launch_enabled"`
	BackendDetailsExposed bool     `json:"backend_details_exposed"`
}

type RunPlanBinding struct {
	Ready                 bool   `json:"ready"`
	LaunchEnabled         bool   `json:"launch_enabled"`
	Reason                string `json:"reason"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

type RunPlanPreflight struct {
	PortalPolicyRequired      bool     `json:"portal_policy_required"`
	SnapshotBeforeRiskyChange bool     `json:"snapshot_before_risky_change"`
	DiagnosticsRequired       bool     `json:"diagnostics_required"`
	BackendBindingRequired    bool     `json:"backend_binding_required"`
	RuntimeWriteGateRequired  bool     `json:"runtime_write_gate_required"`
	RequiredGateIDs           []string `json:"required_gate_ids"`
}

func (plan Plan) RunPlanPreview() (RunPlanPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return RunPlanPreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return RunPlanPreview{}, errors.New("run plan preview requires single-line identity fields")
		}
	}

	profileID := plan.recommendedCompatibilityProfileID()
	engine := plan.runPlanEngine()
	profile := RunPlanProfile{
		ID:                    profileID,
		Label:                 readinessProfileLabel(profileID),
		Kind:                  readinessProfileKind(profileID),
		RequiredPreflight:     []string{"portal-policy-review", "snapshot-baseline", "backend-binding", "runtime-launch-write-gate"},
		Ready:                 false,
		LaunchEnabled:         false,
		BackendDetailsExposed: false,
	}
	preflight := RunPlanPreflight{
		PortalPolicyRequired:      true,
		SnapshotBeforeRiskyChange: true,
		DiagnosticsRequired:       true,
		BackendBindingRequired:    true,
		RuntimeWriteGateRequired:  true,
		RequiredGateIDs:           []string{"portal-policy-review", "snapshot-baseline", "diagnostics", "backend-binding", "runtime-launch-write-gate"},
	}

	preview := RunPlanPreview{
		SchemaVersion: "xnix.runtime.run_plan.v1",
		RequestType:   "run-plan-preview",
		PlanType:      "compatibility-run",
		Source:        "registry+go-runtime-run-plan",
		Desktop:       "KDE Plasma",
		RuntimeMethod: "GetRunPlan",
		ReadMethod:    "GetRunPlanPreview",
		Application: RunPlanApplication{
			ID:          plan.ApplicationID,
			Name:        plan.DisplayName,
			Icon:        plan.Icon,
			DesktopFile: plan.DesktopFile,
			RuntimeMode: plan.runtimeModeID(),
		},
		Execution: RunPlanExecution{
			Strategy:        plan.runPlanStrategy(),
			SelectionSource: "recipe",
			SelectedEngine:  engine,
			SelectedProfile: profile,
			BackendBinding: RunPlanBinding{
				Ready:                 false,
				LaunchEnabled:         false,
				Reason:                "Compatibility profile binding is pending.",
				BackendDetailsExposed: false,
			},
			BackendDetailsExposed:   false,
			LaunchEnabled:           false,
			ExecutionRequestCreated: false,
			ExecutionStarted:        false,
		},
		Preflight:               preflight,
		RuntimeOwned:            true,
		GoRuntimeBacked:         true,
		KDEPolicyOwner:          false,
		UserVisible:             true,
		LaunchEnabled:           false,
		ExecutionRequestCreated: false,
		ExecutionStarted:        false,
		BackendBindingReady:     false,
		RequestObjectCreated:    false,
		PermissionGranted:       false,
		HostRootModified:        false,
		NetworkRequired:         false,
		BackendDetailsExposed:   false,
		RawCommandExposed:       false,
		UserFacingSettings:      plan.UserFacingSettings,
		BlockedActions: []string{
			"create execution request before Runtime gates pass",
			"launch compatibility profile from run plan preview",
			"persist compatibility profile binding from run plan preview",
			"grant desktop resources without Portal review",
			"expose raw backend command to KDE",
			"mutate host root during run plan preview",
		},
		DesktopSafeSummary: "Runtime can explain the compatibility run plan, but launch, binding, permissions, and host writes remain disabled.",
	}
	if err := validateNoBackendTerms(preview, "run plan preview"); err != nil {
		return RunPlanPreview{}, err
	}
	return preview, nil
}

func (plan Plan) runPlanStrategy() string {
	switch plan.RecipeMode {
	case "vm":
		return "isolated-compatible-managed"
	case "wine":
		return "local-compatible-managed"
	default:
		return "automatic-managed"
	}
}

func (plan Plan) runPlanEngine() RunPlanEngineSummary {
	switch plan.RecipeMode {
	case "vm":
		return runPlanEngineSummary("isolated-compatible", "Isolated compatibility")
	case "wine":
		return runPlanEngineSummary("local-compatible", "Local compatibility")
	default:
		return runPlanEngineSummary("automatic", "Automatic")
	}
}

func runPlanEngineSummary(id string, label string) RunPlanEngineSummary {
	return RunPlanEngineSummary{
		ID:                    id,
		Label:                 label,
		Availability:          "planned",
		Ready:                 false,
		LaunchEnabled:         false,
		BackendDetailsExposed: false,
	}
}
