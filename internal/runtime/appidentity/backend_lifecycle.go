package appidentity

import (
	"errors"

	"xnix.local/xnix/internal/runtime/environment"
)

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
	StateRootBacked         bool                        `json:"state_root_backed"`
	StateRootPathExposed    bool                        `json:"state_root_path_exposed"`
	EnvironmentProfile      string                      `json:"environment_profile"`
	SatisfiedGates          []string                    `json:"satisfied_gates"`
	PendingGates            []string                    `json:"pending_gates"`
	RepairHints             []string                    `json:"repair_hints"`
	BlockReason             string                      `json:"block_reason,omitempty"`
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

type BackendLifecycleRecord struct {
	SchemaVersion               string                  `json:"schema_version"`
	RecordType                  string                  `json:"record_type"`
	Source                      string                  `json:"source"`
	Action                      string                  `json:"action"`
	RelativePath                string                  `json:"relative_path"`
	EnvironmentRecord           environment.Record      `json:"environment_record"`
	Preview                     BackendLifecyclePreview `json:"preview"`
	RuntimeOwned                bool                    `json:"runtime_owned"`
	GoRuntimeBacked             bool                    `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                    `json:"kde_policy_owner"`
	StateRootPathExposed        bool                    `json:"state_root_path_exposed"`
	LaunchEnabled               bool                    `json:"launch_enabled"`
	BackendProcessStarted       bool                    `json:"backend_process_started"`
	HostRootModified            bool                    `json:"host_root_modified"`
	BackendDetailsExposed       bool                    `json:"backend_details_exposed"`
	NetworkRequired             bool                    `json:"network_required"`
	PrivilegedContainerRequired bool                    `json:"privileged_container_required"`
	Summary                     string                  `json:"summary"`
}

func (plan Plan) BackendLifecyclePreview() (BackendLifecyclePreview, error) {
	return plan.backendLifecyclePreview(nil)
}

func (plan Plan) BackendLifecyclePreviewWithStateRoot(stateRoot string) (BackendLifecyclePreview, error) {
	if stateRoot == "" {
		return plan.BackendLifecyclePreview()
	}
	lifecycle, err := environment.New(stateRoot)
	if err != nil {
		return BackendLifecyclePreview{}, err
	}
	profile := plan.backendLifecycleProfile()
	record, err := lifecycle.Get(plan.ApplicationID, profile)
	if err != nil {
		return BackendLifecyclePreview{}, err
	}
	return plan.backendLifecyclePreview(&record)
}

func (plan Plan) RecordBackendLifecycleState(stateRoot string, action string, gate string, repairHint string, blockReason string) (BackendLifecycleRecord, error) {
	if stateRoot == "" {
		return BackendLifecycleRecord{}, errors.New("backend lifecycle record requires a state root")
	}
	if action == "" {
		return BackendLifecycleRecord{}, errors.New("backend lifecycle record requires an action")
	}
	if !singleLine(action) {
		return BackendLifecycleRecord{}, errors.New("backend lifecycle record action must be a single-line value")
	}
	for _, value := range []string{gate, repairHint, blockReason} {
		if !singleLineOrBlank(value) {
			return BackendLifecycleRecord{}, errors.New("backend lifecycle record arguments must be single-line values")
		}
	}
	lifecycle, err := environment.New(stateRoot)
	if err != nil {
		return BackendLifecycleRecord{}, err
	}
	profile := plan.backendLifecycleProfile()
	var record environment.Record
	switch action {
	case "inspect":
		record, err = lifecycle.Get(plan.ApplicationID, profile)
	case "plan":
		record, err = lifecycle.Plan(plan.ApplicationID, profile)
	case "stage":
		record, err = lifecycle.Stage(plan.ApplicationID, profile)
	case "satisfy-gate":
		if gate == "" {
			return BackendLifecycleRecord{}, errors.New("backend lifecycle satisfy-gate action requires --gate")
		}
		record, err = lifecycle.SatisfyGate(plan.ApplicationID, profile, gate)
	case "mark-ready":
		record, err = lifecycle.MarkReady(plan.ApplicationID, profile)
	case "flag-repair":
		if repairHint == "" {
			return BackendLifecycleRecord{}, errors.New("backend lifecycle flag-repair action requires --repair-hint")
		}
		record, err = lifecycle.FlagRepair(plan.ApplicationID, profile, []string{repairHint})
	case "block":
		if blockReason == "" {
			return BackendLifecycleRecord{}, errors.New("backend lifecycle block action requires --block-reason")
		}
		record, err = lifecycle.Block(plan.ApplicationID, profile, blockReason)
	case "retire":
		record, err = lifecycle.Retire(plan.ApplicationID, profile)
	default:
		return BackendLifecycleRecord{}, errors.New("unsupported backend lifecycle action")
	}
	if err != nil {
		return BackendLifecycleRecord{}, err
	}
	relativePath, err := environment.RecordRelativePath(plan.ApplicationID, profile)
	if err != nil {
		return BackendLifecycleRecord{}, err
	}
	preview, err := plan.backendLifecyclePreview(&record)
	if err != nil {
		return BackendLifecycleRecord{}, err
	}
	return BackendLifecycleRecord{
		SchemaVersion:               "xnix.runtime.backend_lifecycle_record.v1",
		RecordType:                  "backend-lifecycle-state-record",
		Source:                      "go-runtime-state-root-backend-lifecycle",
		Action:                      action,
		RelativePath:                relativePath,
		EnvironmentRecord:           record,
		Preview:                     preview,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		StateRootPathExposed:        false,
		LaunchEnabled:               false,
		BackendProcessStarted:       false,
		HostRootModified:            false,
		BackendDetailsExposed:       false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		Summary:                     "Runtime persisted backend lifecycle state under the configured state root while launch, backend process start, host mutation, and backend detail exposure remain disabled.",
	}, nil
}

func (plan Plan) backendLifecyclePreview(record *environment.Record) (BackendLifecyclePreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return BackendLifecyclePreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return BackendLifecyclePreview{}, errors.New("backend lifecycle preview requires single-line identity fields")
		}
	}

	profile := plan.backendLifecycleProfile()
	lifecycleState := "blocked"
	overallStatus := "not-ready"
	satisfiedGates := []string{}
	pendingGates := environment.RequiredGates()
	repairHints := []string{}
	blockReason := ""
	stateRootBacked := false
	stateRootReady := false
	backendBindingReady := false
	portalReviewRequired := true
	snapshotRequired := true
	if record != nil {
		stateRootBacked = true
		lifecycleState = string(record.State)
		overallStatus = backendLifecycleOverallStatus(*record)
		satisfiedGates = append([]string{}, record.SatisfiedGates...)
		pendingGates = record.PendingGates()
		repairHints = append([]string{}, record.RepairHints...)
		blockReason = record.BlockReason
		stateRootReady = record.State != environment.StateMissing
		backendBindingReady = backendLifecycleGateSatisfied(*record, "backend-binding")
		portalReviewRequired = !backendLifecycleGateSatisfied(*record, "portal-policy-review")
		snapshotRequired = !backendLifecycleGateSatisfied(*record, "snapshot-baseline")
	}

	stages := backendLifecycleStages(record)
	preview := BackendLifecyclePreview{
		SchemaVersion: "xnix.runtime.backend_lifecycle.v1",
		RequestType:   "backend-lifecycle-preview",
		LifecycleType: "compatibility-backend-lifecycle",
		Source:        backendLifecycleSource(stateRootBacked),
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
		LifecycleState:          lifecycleState,
		OverallStatus:           overallStatus,
		StateRootBacked:         stateRootBacked,
		StateRootPathExposed:    false,
		EnvironmentProfile:      string(profile),
		SatisfiedGates:          satisfiedGates,
		PendingGates:            pendingGates,
		RepairHints:             repairHints,
		BlockReason:             blockReason,
		Stages:                  stages,
		StageIDs:                backendLifecycleStageIDs(stages),
		RuntimeOwned:            true,
		GoRuntimeBacked:         true,
		KDEPolicyOwner:          false,
		BackendBindingReady:     backendBindingReady,
		LaunchEnabled:           false,
		ExecutionRequestCreated: false,
		BackendProcessStarted:   false,
		LocalBackendStarted:     false,
		IsolatedBackendStarted:  false,
		StateRootReady:          stateRootReady,
		PortalReviewRequired:    portalReviewRequired,
		SnapshotRequired:        snapshotRequired,
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
		DesktopSafeSummary: backendLifecycleSummary(lifecycleState, stateRootBacked),
	}
	if err := validateNoBackendTerms(preview, "backend lifecycle preview"); err != nil {
		return BackendLifecyclePreview{}, err
	}
	return preview, nil
}

func (plan Plan) backendLifecycleProfile() environment.Profile {
	if plan.runPlanStrategy() == "isolated-compatible-managed" {
		return environment.ProfileIsolated
	}
	return environment.ProfileLocal
}

func backendLifecycleSource(stateRootBacked bool) string {
	if stateRootBacked {
		return "run-plan-preview+go-runtime-backend-lifecycle+environment-state-root"
	}
	return "run-plan-preview+go-runtime-backend-lifecycle"
}

func backendLifecycleStages(record *environment.Record) []BackendLifecycleStage {
	stateRootStatus := "pending"
	backendBindingStatus := "pending"
	portalSnapshotStatus := "required"
	if record != nil {
		if record.State != environment.StateMissing {
			stateRootStatus = "pass"
		}
		if backendLifecycleGateSatisfied(*record, "backend-binding") {
			backendBindingStatus = "pass"
		}
		if backendLifecycleGateSatisfied(*record, "portal-policy-review") && backendLifecycleGateSatisfied(*record, "snapshot-baseline") {
			portalSnapshotStatus = "pass"
		}
	}
	return []BackendLifecycleStage{
		{ID: "recipe-loaded", Status: "pass", Summary: "Recipe metadata is loaded and mapped to a Runtime-owned backend lifecycle."},
		{ID: "state-root-ready", Status: stateRootStatus, Summary: "A Runtime-owned application state root must exist before backend lifecycle activation."},
		{ID: "backend-binding-ready", Status: backendBindingStatus, Summary: "A managed backend binding must be ready before any lifecycle start request."},
		{ID: "portal-and-snapshot-review", Status: portalSnapshotStatus, Summary: "Portal access and restore-point policy must be reviewed before backend lifecycle activation."},
		{ID: "runtime-launch-write-gate", Status: "blocked", Summary: "Runtime Launch write dispatch remains disabled until production backend ownership is ready."},
	}
}

func backendLifecycleGateSatisfied(record environment.Record, gate string) bool {
	for _, satisfied := range record.SatisfiedGates {
		if satisfied == gate {
			return true
		}
	}
	return false
}

func backendLifecycleOverallStatus(record environment.Record) string {
	switch record.State {
	case environment.StateReady:
		return "ready-with-launch-disabled"
	case environment.StateRepairRequired:
		return "repair-required"
	case environment.StateBlocked:
		return "blocked"
	default:
		return "not-ready"
	}
}

func backendLifecycleSummary(lifecycleState string, stateRootBacked bool) string {
	if stateRootBacked {
		return "Runtime backend lifecycle is backed by a state-root record; KDE may display state and must not start compatibility backends."
	}
	return "Runtime backend lifecycle is modeled but blocked; KDE may display state and must not start compatibility backends."
}

func backendLifecycleStageIDs(stages []BackendLifecycleStage) []string {
	ids := make([]string, 0, len(stages))
	for _, stage := range stages {
		ids = append(ids, stage.ID)
	}
	return ids
}
