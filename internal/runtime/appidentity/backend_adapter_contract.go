package appidentity

type BackendAdapterContractPreview struct {
	Version                     string                   `json:"version"`
	SchemaVersion               string                   `json:"schema_version"`
	RequestType                 string                   `json:"request_type"`
	ContractType                string                   `json:"contract_type"`
	Source                      string                   `json:"source"`
	RuntimeMethod               string                   `json:"runtime_method"`
	ReadMethod                  string                   `json:"read_method"`
	Adapters                    []BackendAdapterContract `json:"adapters"`
	AdapterIDs                  []string                 `json:"adapter_ids"`
	KDEFacingProfiles           []BackendAdapterProfile  `json:"kde_facing_profiles"`
	KDEFacingProfileIDs         []string                 `json:"kde_facing_profile_ids"`
	Counts                      BackendAdapterCounts     `json:"counts"`
	RuntimeOwned                bool                     `json:"runtime_owned"`
	GoRuntimeBacked             bool                     `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                     `json:"kde_policy_owner"`
	NoopImplementation          bool                     `json:"noop_implementation"`
	AdapterInvocationEnabled    bool                     `json:"adapter_invocation_enabled"`
	BackendInstallEnabled       bool                     `json:"backend_install_enabled"`
	BackendDownloadEnabled      bool                     `json:"backend_download_enabled"`
	BackendLaunchEnabled        bool                     `json:"backend_launch_enabled"`
	BackendProcessStarted       bool                     `json:"backend_process_started"`
	VMProcessStarted            bool                     `json:"vm_process_started"`
	CommandMaterialized         bool                     `json:"command_materialized"`
	ExecutablePathResolved      bool                     `json:"executable_path_resolved"`
	RawCommandExposed           bool                     `json:"raw_command_exposed"`
	ProfilePathExposed          bool                     `json:"profile_path_exposed"`
	StateRootPathExposed        bool                     `json:"state_root_path_exposed"`
	BackendDetailsExposedToKDE  bool                     `json:"backend_details_exposed_to_kde"`
	NetworkRequired             bool                     `json:"network_required"`
	HostRootModified            bool                     `json:"host_root_modified"`
	PrivilegedContainerRequired bool                     `json:"privileged_container_required"`
	SecretsExposed              bool                     `json:"secrets_exposed"`
	RequiredRuntimeGates        []string                 `json:"required_runtime_gates"`
	BlockedActions              []string                 `json:"blocked_actions"`
	DesktopSafeSummary          string                   `json:"desktop_safe_summary"`
}

type BackendAdapterContract struct {
	ID                         string   `json:"id"`
	Kind                       string   `json:"kind"`
	RuntimeRole                string   `json:"runtime_role"`
	UserFacingProfileID        string   `json:"user_facing_profile_id"`
	ContractStatus             string   `json:"contract_status"`
	Implementation             string   `json:"implementation"`
	AllowedOperations          []string `json:"allowed_operations"`
	RequiredInputs             []string `json:"required_inputs"`
	RequiredGates              []string `json:"required_gates"`
	AdapterInvocationEnabled   bool     `json:"adapter_invocation_enabled"`
	InstallEnabled             bool     `json:"install_enabled"`
	DownloadEnabled            bool     `json:"download_enabled"`
	LaunchEnabled              bool     `json:"launch_enabled"`
	ProcessStarted             bool     `json:"process_started"`
	VMProcessStarted           bool     `json:"vm_process_started"`
	CommandMaterialized        bool     `json:"command_materialized"`
	ExecutablePathResolved     bool     `json:"executable_path_resolved"`
	RawCommandExposed          bool     `json:"raw_command_exposed"`
	ProfilePathExposed         bool     `json:"profile_path_exposed"`
	StateRootPathExposed       bool     `json:"state_root_path_exposed"`
	BackendDetailsExposedToKDE bool     `json:"backend_details_exposed_to_kde"`
	NetworkRequired            bool     `json:"network_required"`
	HostRootModified           bool     `json:"host_root_modified"`
}

type BackendAdapterProfile struct {
	ID                    string `json:"id"`
	Label                 string `json:"label"`
	Summary               string `json:"summary"`
	ContractStatus        string `json:"contract_status"`
	RuntimeOwned          bool   `json:"runtime_owned"`
	KDEPolicyOwner        bool   `json:"kde_policy_owner"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
	LaunchEnabled         bool   `json:"launch_enabled"`
}

type BackendAdapterCounts struct {
	TotalAdapters       int `json:"total_adapters"`
	NoopAdapters        int `json:"noop_adapters"`
	KDEFacingProfiles   int `json:"kde_facing_profiles"`
	EnabledInvocations  int `json:"enabled_invocations"`
	EnabledLaunches     int `json:"enabled_launches"`
	EnabledInstallers   int `json:"enabled_installers"`
	MaterializedCommand int `json:"materialized_command"`
}

func NewBackendAdapterContractPreview(root string) (BackendAdapterContractPreview, error) {
	if root == "" {
		root = "."
	}
	version, err := readRuntimeServiceBindingVersion(root)
	if err != nil {
		return BackendAdapterContractPreview{}, err
	}

	requiredGates := []string{
		"recipe-trust",
		"artifact-receipt",
		"state-root",
		"backend-binding",
		"portal-policy-review",
		"snapshot-baseline",
		"diagnostics-review",
		"runtime-write-gate",
		"restricted-launch-preflight",
		"test-only-materialization",
	}
	adapters := []BackendAdapterContract{
		backendAdapterContract("wine", "local-windows-api-adapter", "local-compatibility-adapter", "local-compatibility", requiredGates),
		backendAdapterContract("proton", "gaming-compatibility-adapter", "gaming-compatibility-adapter", "game-compatibility", requiredGates),
		backendAdapterContract("windows-vm", "isolated-runtime-adapter", "isolated-compatibility-adapter", "isolated-compatibility", append(requiredGates, "isolated-image-review")),
	}
	profiles := []BackendAdapterProfile{
		backendAdapterProfile("local-compatibility", "Local compatibility", "A reviewed local compatibility profile is planned, but adapter invocation is disabled."),
		backendAdapterProfile("game-compatibility", "Game compatibility", "A reviewed game compatibility profile is planned, but adapter invocation is disabled."),
		backendAdapterProfile("isolated-compatibility", "Isolated compatibility", "A reviewed isolated compatibility profile is planned, but adapter invocation is disabled."),
	}

	preview := BackendAdapterContractPreview{
		Version:              version,
		SchemaVersion:        "xnix.runtime.backend_adapter_contract.v1",
		RequestType:          "backend-adapter-contract-preview",
		ContractType:         "compatibility-backend-adapter-noop-contract",
		Source:               "go-runtime-backend-manager+adapter-noop-boundary",
		RuntimeMethod:        "GetBackendAdapterContract",
		ReadMethod:           "GetBackendAdapterContractPreview",
		Adapters:             adapters,
		AdapterIDs:           backendAdapterIDs(adapters),
		KDEFacingProfiles:    profiles,
		KDEFacingProfileIDs:  backendAdapterProfileIDs(profiles),
		Counts:               backendAdapterCounts(adapters, profiles),
		RuntimeOwned:         true,
		GoRuntimeBacked:      true,
		KDEPolicyOwner:       false,
		NoopImplementation:   true,
		RequiredRuntimeGates: requiredGates,
		BlockedActions: []string{
			"invoke compatibility backend adapter",
			"install compatibility backend through adapter contract",
			"download compatibility backend through adapter contract",
			"materialize backend command from adapter contract",
			"resolve executable path from adapter contract",
			"start local backend process from adapter contract",
			"start isolated runtime process from adapter contract",
			"expose backend details, paths, commands, or storage to KDE",
			"mutate host root during adapter contract preview",
		},
		DesktopSafeSummary: "Runtime defines reviewed compatibility adapter profiles as no-op contracts; KDE sees only safe profiles while invocation, install, download, launch, command materialization, paths, and host mutation remain disabled.",
	}
	return preview, nil
}

func backendAdapterContract(id, kind, role, profileID string, requiredGates []string) BackendAdapterContract {
	return BackendAdapterContract{
		ID:                         id,
		Kind:                       kind,
		RuntimeRole:                role,
		UserFacingProfileID:        profileID,
		ContractStatus:             "noop-contract",
		Implementation:             "noop",
		AllowedOperations:          []string{"inspect-contract", "report-disabled-gates"},
		RequiredInputs:             []string{"verified-recipe", "artifact-receipt", "state-root-record", "portal-review", "snapshot-baseline", "diagnostic-summary"},
		RequiredGates:              append([]string{}, requiredGates...),
		AdapterInvocationEnabled:   false,
		InstallEnabled:             false,
		DownloadEnabled:            false,
		LaunchEnabled:              false,
		ProcessStarted:             false,
		VMProcessStarted:           false,
		CommandMaterialized:        false,
		ExecutablePathResolved:     false,
		RawCommandExposed:          false,
		ProfilePathExposed:         false,
		StateRootPathExposed:       false,
		BackendDetailsExposedToKDE: false,
		NetworkRequired:            false,
		HostRootModified:           false,
	}
}

func backendAdapterProfile(id, label, summary string) BackendAdapterProfile {
	return BackendAdapterProfile{
		ID:                    id,
		Label:                 label,
		Summary:               summary,
		ContractStatus:        "noop-contract",
		RuntimeOwned:          true,
		KDEPolicyOwner:        false,
		BackendDetailsExposed: false,
		LaunchEnabled:         false,
	}
}

func backendAdapterIDs(adapters []BackendAdapterContract) []string {
	ids := make([]string, 0, len(adapters))
	for _, adapter := range adapters {
		ids = append(ids, adapter.ID)
	}
	return ids
}

func backendAdapterProfileIDs(profiles []BackendAdapterProfile) []string {
	ids := make([]string, 0, len(profiles))
	for _, profile := range profiles {
		ids = append(ids, profile.ID)
	}
	return ids
}

func backendAdapterCounts(adapters []BackendAdapterContract, profiles []BackendAdapterProfile) BackendAdapterCounts {
	counts := BackendAdapterCounts{
		TotalAdapters:     len(adapters),
		KDEFacingProfiles: len(profiles),
	}
	for _, adapter := range adapters {
		if adapter.Implementation == "noop" {
			counts.NoopAdapters++
		}
		if adapter.AdapterInvocationEnabled {
			counts.EnabledInvocations++
		}
		if adapter.LaunchEnabled {
			counts.EnabledLaunches++
		}
		if adapter.InstallEnabled {
			counts.EnabledInstallers++
		}
		if adapter.CommandMaterialized {
			counts.MaterializedCommand++
		}
	}
	return counts
}
