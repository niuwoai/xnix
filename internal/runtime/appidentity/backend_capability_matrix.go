package appidentity

type BackendCapabilityMatrixPreview struct {
	SchemaVersion               string                     `json:"schema_version"`
	RequestType                 string                     `json:"request_type"`
	MatrixType                  string                     `json:"matrix_type"`
	Source                      string                     `json:"source"`
	Desktop                     string                     `json:"desktop"`
	RuntimeMethod               string                     `json:"runtime_method"`
	ReadMethod                  string                     `json:"read_method"`
	Profiles                    []BackendCapabilityProfile `json:"profiles"`
	ProfileIDs                  []string                   `json:"profile_ids"`
	ProfileCount                int                        `json:"profile_count"`
	CapabilityCount             int                        `json:"capability_count"`
	ReadyCapabilityCount        int                        `json:"ready_capability_count"`
	PendingCapabilityCount      int                        `json:"pending_capability_count"`
	BlockedCapabilityCount      int                        `json:"blocked_capability_count"`
	RuntimeOwned                bool                       `json:"runtime_owned"`
	GoRuntimeBacked             bool                       `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                       `json:"kde_policy_owner"`
	SelectionEnabled            bool                       `json:"selection_enabled"`
	BackendLaunchEnabled        bool                       `json:"backend_launch_enabled"`
	CapabilityActivationEnabled bool                       `json:"capability_activation_enabled"`
	RequestObjectsCreated       bool                       `json:"request_objects_created"`
	StateRootCreated            bool                       `json:"state_root_created"`
	SnapshotsCreated            bool                       `json:"snapshots_created"`
	HostRootModified            bool                       `json:"host_root_modified"`
	NetworkRequired             bool                       `json:"network_required"`
	PrivilegedContainerRequired bool                       `json:"privileged_container_required"`
	BackendDetailsExposed       bool                       `json:"backend_details_exposed"`
	DesktopSafeSummary          string                     `json:"desktop_safe_summary"`
}

type BackendCapabilityProfile struct {
	ID                          string              `json:"id"`
	Label                       string              `json:"label"`
	Kind                        string              `json:"kind"`
	SelectionState              string              `json:"selection_state"`
	Capabilities                []BackendCapability `json:"capabilities"`
	CapabilityCount             int                 `json:"capability_count"`
	ReadyCount                  int                 `json:"ready_count"`
	PendingCount                int                 `json:"pending_count"`
	BlockedCount                int                 `json:"blocked_count"`
	RuntimeOwned                bool                `json:"runtime_owned"`
	KDEPolicyOwner              bool                `json:"kde_policy_owner"`
	ProfileReady                bool                `json:"profile_ready"`
	SelectionEnabled            bool                `json:"selection_enabled"`
	BackendProcessStarted       bool                `json:"backend_process_started"`
	NetworkRequiredForPlanning  bool                `json:"network_required_for_planning"`
	HostRootModified            bool                `json:"host_root_modified"`
	PrivilegedContainerRequired bool                `json:"privileged_container_required"`
	BackendDetailsExposed       bool                `json:"backend_details_exposed"`
	Summary                     string              `json:"summary"`
}

type BackendCapability struct {
	ID                    string `json:"id"`
	Label                 string `json:"label"`
	Status                string `json:"status"`
	Summary               string `json:"summary"`
	UserVisible           bool   `json:"user_visible"`
	RequiresPortalReview  bool   `json:"requires_portal_review"`
	RequiresStateRoot     bool   `json:"requires_state_root"`
	RequiresSnapshot      bool   `json:"requires_snapshot"`
	ActivationEnabled     bool   `json:"activation_enabled"`
	RequestObjectCreated  bool   `json:"request_object_created"`
	BackendProcessStarted bool   `json:"backend_process_started"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

type backendCapabilitySpec struct {
	id      string
	label   string
	status  string
	summary string
}

var backendCapabilitySpecs = []backendCapabilitySpec{
	{"application-launch", "Application launch", "pending", "Runtime launch binding is still gated."},
	{"package-management", "Managed packages", "pending", "Runtime package sources are planned before backend availability."},
	{"file-bridge", "File bridge", "pending", "File access must pass Portal policy and state-root preflight."},
	{"clipboard-bridge", "Clipboard bridge", "pending", "Clipboard access must pass Portal policy review."},
	{"print-bridge", "Print bridge", "pending", "Print access must pass Portal policy review."},
	{"snapshot-restore", "Snapshot and restore", "pending", "Restore points must exist before risky compatibility changes."},
	{"diagnostics", "Diagnostics", "ready", "Runtime can describe diagnostics without starting a backend."},
}

type backendCapabilityProfileSpec struct {
	id    string
	label string
	kind  string
}

var backendCapabilityProfileSpecs = []backendCapabilityProfileSpec{
	{"local-compatibility", "Local compatibility profile", "local"},
	{"isolated-compatibility", "Isolated compatibility profile", "isolated"},
}

func NewBackendCapabilityMatrixPreview() (BackendCapabilityMatrixPreview, error) {
	profiles := backendCapabilityProfiles()
	readyTotal := 0
	pendingTotal := 0
	blockedTotal := 0
	profileIDs := make([]string, 0, len(profiles))
	for _, profile := range profiles {
		readyTotal += profile.ReadyCount
		pendingTotal += profile.PendingCount
		blockedTotal += profile.BlockedCount
		profileIDs = append(profileIDs, profile.ID)
	}

	preview := BackendCapabilityMatrixPreview{
		SchemaVersion:               "xnix.runtime.backend_capability_matrix.v1",
		RequestType:                 "backend-capability-matrix-preview",
		MatrixType:                  "compatibility-backend-capability-matrix",
		Source:                      "go-runtime-backend-capability-matrix",
		Desktop:                     "KDE Plasma",
		RuntimeMethod:               "GetBackendCapabilityMatrix",
		ReadMethod:                  "GetBackendCapabilityMatrixPreview",
		Profiles:                    profiles,
		ProfileIDs:                  profileIDs,
		ProfileCount:                len(profiles),
		CapabilityCount:             len(backendCapabilitySpecs),
		ReadyCapabilityCount:        readyTotal,
		PendingCapabilityCount:      pendingTotal,
		BlockedCapabilityCount:      blockedTotal,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		SelectionEnabled:            false,
		BackendLaunchEnabled:        false,
		CapabilityActivationEnabled: false,
		RequestObjectsCreated:       false,
		StateRootCreated:            false,
		SnapshotsCreated:            false,
		HostRootModified:            false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		DesktopSafeSummary:          "Runtime backend capability planning is visible to KDE, while backend selection and launch remain disabled.",
	}
	if err := validateNoBackendTerms(preview, "backend capability matrix preview"); err != nil {
		return BackendCapabilityMatrixPreview{}, err
	}
	return preview, nil
}

func backendCapabilityProfiles() []BackendCapabilityProfile {
	profiles := make([]BackendCapabilityProfile, 0, len(backendCapabilityProfileSpecs))
	for _, spec := range backendCapabilityProfileSpecs {
		capabilities := backendCapabilities()
		ready := 0
		pending := 0
		blocked := 0
		for _, capability := range capabilities {
			switch capability.Status {
			case "ready":
				ready++
			case "pending":
				pending++
			case "blocked":
				blocked++
			}
		}
		profiles = append(profiles, BackendCapabilityProfile{
			ID:                          spec.id,
			Label:                       spec.label,
			Kind:                        spec.kind,
			SelectionState:              "planned",
			Capabilities:                capabilities,
			CapabilityCount:             len(capabilities),
			ReadyCount:                  ready,
			PendingCount:                pending,
			BlockedCount:                blocked,
			RuntimeOwned:                true,
			KDEPolicyOwner:              false,
			ProfileReady:                false,
			SelectionEnabled:            false,
			BackendProcessStarted:       false,
			NetworkRequiredForPlanning:  false,
			HostRootModified:            false,
			PrivilegedContainerRequired: false,
			BackendDetailsExposed:       false,
			Summary:                     spec.label + " is planned but not selectable until Runtime preflight passes.",
		})
	}
	return profiles
}

func backendCapabilities() []BackendCapability {
	capabilities := make([]BackendCapability, 0, len(backendCapabilitySpecs))
	for _, spec := range backendCapabilitySpecs {
		capabilities = append(capabilities, BackendCapability{
			ID:                    spec.id,
			Label:                 spec.label,
			Status:                spec.status,
			Summary:               spec.summary,
			UserVisible:           true,
			RequiresPortalReview:  spec.id == "file-bridge" || spec.id == "clipboard-bridge" || spec.id == "print-bridge",
			RequiresStateRoot:     spec.id == "application-launch" || spec.id == "file-bridge" || spec.id == "snapshot-restore",
			RequiresSnapshot:      spec.id == "snapshot-restore",
			ActivationEnabled:     false,
			RequestObjectCreated:  false,
			BackendProcessStarted: false,
			BackendDetailsExposed: false,
		})
	}
	return capabilities
}
