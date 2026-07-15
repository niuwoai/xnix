package appidentity

type RuntimeOwnerSmokePlanPreview struct {
	Version                     string                       `json:"version"`
	SchemaVersion               string                       `json:"schema_version"`
	RequestType                 string                       `json:"request_type"`
	PlanType                    string                       `json:"plan_type"`
	Source                      string                       `json:"source"`
	RuntimeMethod               string                       `json:"runtime_method"`
	ReadMethod                  string                       `json:"read_method"`
	BusName                     string                       `json:"bus_name"`
	ObjectPath                  string                       `json:"object_path"`
	Interface                   string                       `json:"interface"`
	LiveOwnerGate               RuntimeOwnerSmokeGateSummary `json:"live_owner_gate"`
	ActivationBindingReady      bool                         `json:"activation_binding_ready"`
	LiveDBusOwnerReady          bool                         `json:"live_dbus_owner_ready"`
	ProductionOwnerEnabled      bool                         `json:"production_owner_enabled"`
	OwnerTransitionReady        bool                         `json:"owner_transition_ready"`
	SmokeState                  string                       `json:"smoke_state"`
	SmokeEnvironment            string                       `json:"smoke_environment"`
	Steps                       []RuntimeOwnerSmokeStep      `json:"steps"`
	StepIDs                     []string                     `json:"step_ids"`
	Counts                      RuntimeOwnerSmokeCounts      `json:"counts"`
	PendingStepCount            int                          `json:"pending_step_count"`
	RuntimeOwned                bool                         `json:"runtime_owned"`
	GoRuntimeBacked             bool                         `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                         `json:"kde_policy_owner"`
	NetworkRequired             bool                         `json:"network_required"`
	HostRootModified            bool                         `json:"host_root_modified"`
	PrivilegedContainerRequired bool                         `json:"privileged_container_required"`
	SystemServiceStarted        bool                         `json:"system_service_started"`
	ProductionBusClaimed        bool                         `json:"production_bus_claimed"`
	BackendDetailsExposed       bool                         `json:"backend_details_exposed"`
	BlockedActions              []string                     `json:"blocked_actions"`
	DesktopSafeSummary          string                       `json:"desktop_safe_summary"`
}

type RuntimeOwnerSmokeGateSummary struct {
	RequestType            string                     `json:"request_type"`
	GateType               string                     `json:"gate_type"`
	ActivationBindingReady bool                       `json:"activation_binding_ready"`
	LiveDBusOwnerReady     bool                       `json:"live_dbus_owner_ready"`
	ProductionOwnerEnabled bool                       `json:"production_owner_enabled"`
	OwnerTransitionReady   bool                       `json:"owner_transition_ready"`
	Counts                 RuntimeLiveOwnerGateCounts `json:"counts"`
}

type RuntimeOwnerSmokeStep struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type RuntimeOwnerSmokeCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewRuntimeOwnerSmokePlanPreview(root string) (RuntimeOwnerSmokePlanPreview, error) {
	liveOwnerGate, err := NewRuntimeLiveOwnerGatePreview(root)
	if err != nil {
		return RuntimeOwnerSmokePlanPreview{}, err
	}

	steps := runtimeOwnerSmokeSteps(liveOwnerGate.ActivationBindingReady)
	counts := countRuntimeOwnerSmokeSteps(steps)
	preview := RuntimeOwnerSmokePlanPreview{
		Version:       liveOwnerGate.Version,
		SchemaVersion: "xnix.runtime.owner_smoke_plan.v1",
		RequestType:   "runtime-owner-smoke-plan-preview",
		PlanType:      "runtime-owner-smoke-plan",
		Source:        "runtime-live-owner-gate-preview+runtime-service-binding-preview",
		RuntimeMethod: "GetRuntimeOwnerSmokePlan",
		ReadMethod:    "GetRuntimeOwnerSmokePlanPreview",
		BusName:       liveOwnerGate.BusName,
		ObjectPath:    liveOwnerGate.ObjectPath,
		Interface:     liveOwnerGate.Interface,
		LiveOwnerGate: RuntimeOwnerSmokeGateSummary{
			RequestType:            liveOwnerGate.RequestType,
			GateType:               liveOwnerGate.GateType,
			ActivationBindingReady: liveOwnerGate.ActivationBindingReady,
			LiveDBusOwnerReady:     liveOwnerGate.LiveDBusOwnerReady,
			ProductionOwnerEnabled: liveOwnerGate.ProductionOwnerEnabled,
			OwnerTransitionReady:   liveOwnerGate.OwnerTransitionReady,
			Counts:                 liveOwnerGate.Counts,
		},
		ActivationBindingReady:      liveOwnerGate.ActivationBindingReady,
		LiveDBusOwnerReady:          false,
		ProductionOwnerEnabled:      false,
		OwnerTransitionReady:        false,
		SmokeState:                  "planned",
		SmokeEnvironment:            "restricted-session",
		Steps:                       steps,
		StepIDs:                     runtimeOwnerSmokeStepIDs(steps),
		Counts:                      counts,
		PendingStepCount:            counts.Pending,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		SystemServiceStarted:        false,
		ProductionBusClaimed:        false,
		BackendDetailsExposed:       false,
		BlockedActions:              runtimeOwnerSmokeBlockedActions(),
		DesktopSafeSummary:          runtimeOwnerSmokeSummary(liveOwnerGate.ActivationBindingReady),
	}
	if err := validateNoBackendTerms(preview, "Runtime owner smoke plan preview"); err != nil {
		return RuntimeOwnerSmokePlanPreview{}, err
	}
	return preview, nil
}

func runtimeOwnerSmokeSteps(activationBindingReady bool) []RuntimeOwnerSmokeStep {
	activationStatus := "blocked"
	if activationBindingReady {
		activationStatus = "pass"
	}
	return []RuntimeOwnerSmokeStep{
		runtimeOwnerSmokeStep("validate-activation-files", activationStatus, "Verify D-Bus service activation, systemd hardening, libexec wrapper, and contract alignment."),
		runtimeOwnerSmokeStep("start-packaged-runtime-owner", "pending", "Start the packaged Runtime owner in an isolated session without modifying the host root."),
		runtimeOwnerSmokeStep("assert-stable-bus-name", "pending", "Prove the packaged Runtime owner owns org.xnix.Compatibility1 on the test bus."),
		runtimeOwnerSmokeStep("check-read-only-method-parity", "pending", "Call the read-only planning methods required by KDE and compare them with the contract."),
		runtimeOwnerSmokeStep("reject-write-methods", "pending", "Confirm launch, install, snapshot, restore, repair, and settings persistence remain gated."),
		runtimeOwnerSmokeStep("verify-non-production-smoke-adapter-boundary", "pending", "Confirm the smoke adapter is never accepted as a production Runtime owner."),
		runtimeOwnerSmokeStep("report-kde-safe-summary", "pending", "Return a Compatibility Center summary without backend details or host paths."),
	}
}

func runtimeOwnerSmokeStep(id string, status string, summary string) RuntimeOwnerSmokeStep {
	return RuntimeOwnerSmokeStep{
		ID:      id,
		Status:  status,
		Summary: summary,
	}
}

func runtimeOwnerSmokeStepIDs(steps []RuntimeOwnerSmokeStep) []string {
	ids := make([]string, 0, len(steps))
	for _, step := range steps {
		ids = append(ids, step.ID)
	}
	return ids
}

func countRuntimeOwnerSmokeSteps(steps []RuntimeOwnerSmokeStep) RuntimeOwnerSmokeCounts {
	counts := RuntimeOwnerSmokeCounts{Total: len(steps)}
	for _, step := range steps {
		switch step.Status {
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

func runtimeOwnerSmokeBlockedActions() []string {
	return []string{
		"Do not start a host system service from the smoke plan.",
		"Do not claim the production Runtime bus name from the smoke adapter.",
		"Do not enable backend launch, install, repair, restore, or settings persistence.",
		"Do not mutate the host root while planning owner smoke.",
		"Do not expose backend implementation details or host paths to KDE.",
	}
}

func runtimeOwnerSmokeSummary(activationBindingReady bool) string {
	if !activationBindingReady {
		return "Runtime activation files must be repaired before owner smoke can run."
	}
	return "Runtime owner smoke is planned; production bus ownership remains disabled until all gates pass."
}
