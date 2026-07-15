package appidentity

type RuntimeLiveOwnerGatePreview struct {
	Version                     string                         `json:"version"`
	SchemaVersion               string                         `json:"schema_version"`
	RequestType                 string                         `json:"request_type"`
	GateType                    string                         `json:"gate_type"`
	Source                      string                         `json:"source"`
	RuntimeMethod               string                         `json:"runtime_method"`
	ReadMethod                  string                         `json:"read_method"`
	BusName                     string                         `json:"bus_name"`
	ObjectPath                  string                         `json:"object_path"`
	Interface                   string                         `json:"interface"`
	ServiceBinding              RuntimeLiveOwnerBindingSummary `json:"service_binding"`
	RequiredGates               []RuntimeLiveOwnerRequiredGate `json:"required_gates"`
	GateIDs                     []string                       `json:"gate_ids"`
	Counts                      RuntimeLiveOwnerGateCounts     `json:"counts"`
	RuntimeOwned                bool                           `json:"runtime_owned"`
	GoRuntimeBacked             bool                           `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                           `json:"kde_policy_owner"`
	ActivationBindingReady      bool                           `json:"activation_binding_ready"`
	LiveDBusOwnerReady          bool                           `json:"live_dbus_owner_ready"`
	ProductionOwnerEnabled      bool                           `json:"production_owner_enabled"`
	OwnerTransitionReady        bool                           `json:"owner_transition_ready"`
	SmokeAdapterAvailable       bool                           `json:"smoke_adapter_available"`
	SmokeAdapterIsProduction    bool                           `json:"smoke_adapter_is_production_owner"`
	KDEMayClaimRuntimeOwnership bool                           `json:"kde_may_claim_runtime_ownership"`
	NetworkRequired             bool                           `json:"network_required"`
	HostRootModified            bool                           `json:"host_root_modified"`
	PrivilegedContainerRequired bool                           `json:"privileged_container_required"`
	BackendDetailsExposed       bool                           `json:"backend_details_exposed"`
	BlockedReasons              []string                       `json:"blocked_reasons"`
	BlockedActions              []string                       `json:"blocked_actions"`
	DesktopSafeSummary          string                         `json:"desktop_safe_summary"`
}

type RuntimeLiveOwnerBindingSummary struct {
	RequestType            string                      `json:"request_type"`
	ProductionStatus       string                      `json:"production_status"`
	ActivationBindingReady bool                        `json:"activation_binding_ready"`
	LiveDBusOwnerReady     bool                        `json:"live_dbus_owner_ready"`
	SmokeAdapterAvailable  bool                        `json:"smoke_adapter_available"`
	Counts                 RuntimeServiceBindingCounts `json:"counts"`
}

type RuntimeLiveOwnerRequiredGate struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type RuntimeLiveOwnerGateCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewRuntimeLiveOwnerGatePreview(root string) (RuntimeLiveOwnerGatePreview, error) {
	serviceBinding, err := NewRuntimeServiceBindingPreview(root)
	if err != nil {
		return RuntimeLiveOwnerGatePreview{}, err
	}

	requiredGates := runtimeLiveOwnerRequiredGates(serviceBinding.ActivationBindingReady)
	counts := countRuntimeLiveOwnerGates(requiredGates)

	preview := RuntimeLiveOwnerGatePreview{
		Version:       serviceBinding.Version,
		SchemaVersion: "xnix.runtime.live_owner_gate.v1",
		RequestType:   "runtime-live-owner-gate-preview",
		GateType:      "runtime-live-owner-gate",
		Source:        "runtime-service-binding-preview+owner-transition-gates",
		RuntimeMethod: "GetRuntimeLiveOwnerGate",
		ReadMethod:    "GetRuntimeLiveOwnerGatePreview",
		BusName:       serviceBinding.BusName,
		ObjectPath:    serviceBinding.ObjectPath,
		Interface:     serviceBinding.Interface,
		ServiceBinding: RuntimeLiveOwnerBindingSummary{
			RequestType:            serviceBinding.RequestType,
			ProductionStatus:       serviceBinding.ProductionStatus,
			ActivationBindingReady: serviceBinding.ActivationBindingReady,
			LiveDBusOwnerReady:     serviceBinding.LiveDBusOwnerReady,
			SmokeAdapterAvailable:  serviceBinding.SmokeAdapterAvailable,
			Counts:                 serviceBinding.Counts,
		},
		RequiredGates:               requiredGates,
		GateIDs:                     runtimeLiveOwnerGateIDs(requiredGates),
		Counts:                      counts,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		ActivationBindingReady:      serviceBinding.ActivationBindingReady,
		LiveDBusOwnerReady:          false,
		ProductionOwnerEnabled:      false,
		OwnerTransitionReady:        false,
		SmokeAdapterAvailable:       serviceBinding.SmokeAdapterAvailable,
		SmokeAdapterIsProduction:    false,
		KDEMayClaimRuntimeOwnership: false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		BlockedReasons:              runtimeLiveOwnerBlockedReasons(),
		BlockedActions:              runtimeLiveOwnerBlockedActions(),
		DesktopSafeSummary:          runtimeLiveOwnerSummary(serviceBinding.ActivationBindingReady),
	}
	if err := validateNoBackendTerms(preview, "Runtime live owner gate preview"); err != nil {
		return RuntimeLiveOwnerGatePreview{}, err
	}
	return preview, nil
}

func runtimeLiveOwnerRequiredGates(activationBindingReady bool) []RuntimeLiveOwnerRequiredGate {
	activationStatus := "blocked"
	if activationBindingReady {
		activationStatus = "pass"
	}
	return []RuntimeLiveOwnerRequiredGate{
		runtimeLiveOwnerGate("activation-binding", activationStatus, "D-Bus activation files, systemd unit, libexec wrapper, and contract must stay aligned."),
		runtimeLiveOwnerGate("long-running-runtime-owner", "pending", "The Runtime needs a packaged long-running process that owns the stable bus name."),
		runtimeLiveOwnerGate("bus-name-acquisition", "pending", "Production smoke must prove the packaged Runtime owns org.xnix.Compatibility1."),
		runtimeLiveOwnerGate("read-only-method-parity", "pending", "The live owner must answer the same read-only planning methods as the smoke adapter."),
		runtimeLiveOwnerGate("production-recipe-trust", "pending", "Production ownership must be gated by signed recipe validation instead of development registry trust."),
	}
}

func runtimeLiveOwnerGate(id string, status string, summary string) RuntimeLiveOwnerRequiredGate {
	return RuntimeLiveOwnerRequiredGate{
		ID:      id,
		Status:  status,
		Summary: summary,
	}
}

func runtimeLiveOwnerGateIDs(gates []RuntimeLiveOwnerRequiredGate) []string {
	ids := make([]string, 0, len(gates))
	for _, gate := range gates {
		ids = append(ids, gate.ID)
	}
	return ids
}

func countRuntimeLiveOwnerGates(gates []RuntimeLiveOwnerRequiredGate) RuntimeLiveOwnerGateCounts {
	counts := RuntimeLiveOwnerGateCounts{Total: len(gates)}
	for _, gate := range gates {
		switch gate.Status {
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

func runtimeLiveOwnerBlockedReasons() []string {
	return []string{
		"Do not treat the D-Bus smoke adapter as the production Runtime owner.",
		"Do not let KDE own Runtime policy or bus-name readiness decisions.",
		"Do not enable launch, install, repair, restore, or settings persistence from this gate.",
		"Do not mutate the host root while evaluating live-owner readiness.",
		"Do not expose compatibility backend implementation details in live-owner readiness.",
	}
}

func runtimeLiveOwnerBlockedActions() []string {
	return []string{
		"start production Runtime owner from preview",
		"claim production D-Bus name from preview",
		"mark smoke adapter as production owner",
		"let KDE claim Runtime ownership",
		"enable launch, install, repair, restore, or settings persistence",
		"mutate host root during live-owner gate evaluation",
	}
}

func runtimeLiveOwnerSummary(activationBindingReady bool) string {
	if !activationBindingReady {
		return "Runtime activation files are blocked; production D-Bus ownership cannot be enabled."
	}
	return "Runtime activation files are aligned, but production D-Bus ownership remains gated."
}
