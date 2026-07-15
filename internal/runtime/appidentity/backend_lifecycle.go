package appidentity

import "errors"

type BackendLifecyclePreview struct {
	SchemaVersion           string                      `json:"schema_version"`
	RequestType             string                      `json:"request_type"`
	LifecycleType           string                      `json:"lifecycle_type"`
	Source                  string                      `json:"source"`
	Desktop                 string                      `json:"desktop"`
	RuntimeMethod           string                      `json:"runtime_method"`
	ReadMethod              string                      `json:"read_method"`
	Application             BackendLifecycleApplication `json:"application"`
	SelectedStrategy        string                      `json:"selected_strategy"`
	LifecycleState          string                      `json:"lifecycle_state"`
	OverallStatus           string                      `json:"overall_status"`
	Stages                  []BackendLifecycleStage     `json:"stages"`
	StageIDs                []string                    `json:"stage_ids"`
	RuntimeOwned            bool                        `json:"runtime_owned"`
	GoRuntimeBacked         bool                        `json:"go_runtime_backed"`
	KDEPolicyOwner          bool                        `json:"kde_policy_owner"`
	BackendBindingReady     bool                        `json:"backend_binding_ready"`
	LaunchEnabled           bool                        `json:"launch_enabled"`
	ExecutionRequestCreated bool                        `json:"execution_request_created"`
	BackendProcessStarted   bool                        `json:"backend_process_started"`
	LocalBackendStarted     bool                        `json:"local_backend_started"`
	IsolatedBackendStarted  bool                        `json:"isolated_backend_started"`
	StateRootReady          bool                        `json:"state_root_ready"`
	PortalReviewRequired    bool                        `json:"portal_review_required"`
	SnapshotRequired        bool                        `json:"snapshot_required"`
	HostRootModified        bool                        `json:"host_root_modified"`
	NetworkRequired         bool                        `json:"network_required"`
	BackendDetailsExposed   bool                        `json:"backend_details_exposed"`
	UserFacingSettings      map[string]string           `json:"user_facing_settings"`
	BlockedActions          []string                    `json:"blocked_actions"`
	DesktopSafeSummary      string                      `json:"desktop_safe_summary"`
}

type BackendLifecycleApplication struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Icon          string `json:"icon"`
	DesktopFile   string `json:"desktop_file"`
	RequestedMode string `json:"requested_mode"`
}

type BackendLifecycleStage struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

func (plan Plan) BackendLifecyclePreview() (BackendLifecyclePreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return BackendLifecyclePreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return BackendLifecyclePreview{}, errors.New("backend lifecycle preview requires single-line identity fields")
		}
	}

	stages := backendLifecycleStages()
	preview := BackendLifecyclePreview{
		SchemaVersion: "xnix.runtime.backend_lifecycle.v1",
		RequestType:   "backend-lifecycle-preview",
		LifecycleType: "compatibility-backend-lifecycle",
		Source:        "run-plan-preview+go-runtime-backend-lifecycle",
		Desktop:       "KDE Plasma",
		RuntimeMethod: "GetBackendLifecycle",
		ReadMethod:    "GetBackendLifecyclePreview",
		Application: BackendLifecycleApplication{
			ID:            plan.ApplicationID,
			Name:          plan.DisplayName,
			Icon:          plan.Icon,
			DesktopFile:   plan.DesktopFile,
			RequestedMode: plan.runtimeModeID(),
		},
		SelectedStrategy:        plan.runPlanStrategy(),
		LifecycleState:          "blocked",
		OverallStatus:           "not-ready",
		Stages:                  stages,
		StageIDs:                backendLifecycleStageIDs(stages),
		RuntimeOwned:            true,
		GoRuntimeBacked:         true,
		KDEPolicyOwner:          false,
		BackendBindingReady:     false,
		LaunchEnabled:           false,
		ExecutionRequestCreated: false,
		BackendProcessStarted:   false,
		LocalBackendStarted:     false,
		IsolatedBackendStarted:  false,
		StateRootReady:          false,
		PortalReviewRequired:    true,
		SnapshotRequired:        true,
		HostRootModified:        false,
		NetworkRequired:         false,
		BackendDetailsExposed:   false,
		UserFacingSettings:      plan.UserFacingSettings,
		BlockedActions: []string{
			"start local compatibility backend from KDE",
			"start isolated compatibility backend from KDE",
			"create backend process before Runtime lifecycle gates pass",
			"expose backend command to desktop shell",
			"mutate host root during lifecycle planning",
		},
		DesktopSafeSummary: "Runtime backend lifecycle is modeled but blocked; KDE may display state and must not start compatibility backends.",
	}
	if err := validateNoBackendTerms(preview, "backend lifecycle preview"); err != nil {
		return BackendLifecyclePreview{}, err
	}
	return preview, nil
}

func backendLifecycleStages() []BackendLifecycleStage {
	return []BackendLifecycleStage{
		{ID: "recipe-loaded", Status: "pass", Summary: "Recipe metadata is loaded and mapped to a Runtime-owned backend lifecycle."},
		{ID: "state-root-ready", Status: "pending", Summary: "A Runtime-owned application state root must exist before backend lifecycle activation."},
		{ID: "backend-binding-ready", Status: "pending", Summary: "A managed backend binding must be ready before any lifecycle start request."},
		{ID: "portal-and-snapshot-review", Status: "required", Summary: "Portal access and restore-point policy must be reviewed before backend lifecycle activation."},
		{ID: "runtime-launch-write-gate", Status: "blocked", Summary: "Runtime Launch write dispatch remains disabled until production backend ownership is ready."},
	}
}

func backendLifecycleStageIDs(stages []BackendLifecycleStage) []string {
	ids := make([]string, 0, len(stages))
	for _, stage := range stages {
		ids = append(ids, stage.ID)
	}
	return ids
}
