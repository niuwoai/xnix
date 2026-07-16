package appidentity

type WindowsCompatibilityWorkstreamsPreview struct {
	SchemaVersion                     string                               `json:"schema_version"`
	RequestType                       string                               `json:"request_type"`
	PlanType                          string                               `json:"plan_type"`
	Source                            string                               `json:"source"`
	ReadMethod                        string                               `json:"read_method"`
	ProductTarget                     string                               `json:"product_target"`
	OfficialDesktop                   string                               `json:"official_desktop"`
	FutureDesktopOrder                []string                             `json:"future_desktop_order"`
	ArchitectureLayers                []string                             `json:"architecture_layers"`
	EntryPoints                       []WindowsCompatibilityEntryPoint     `json:"entry_points"`
	EntryPointIDs                     []string                             `json:"entry_point_ids"`
	EntryPointCount                   int                                  `json:"entry_point_count"`
	Workstreams                       []WindowsCompatibilityWorkstream     `json:"workstreams"`
	WorkstreamIDs                     []string                             `json:"workstream_ids"`
	WorkstreamCount                   int                                  `json:"workstream_count"`
	FirstWaveWorkstreams              []WindowsCompatibilityFirstWaveEntry `json:"first_wave_workstreams"`
	FirstWaveIDs                      []string                             `json:"first_wave_ids"`
	FirstWaveCount                    int                                  `json:"first_wave_count"`
	RuntimeOwned                      bool                                 `json:"runtime_owned"`
	GoRuntimeBacked                   bool                                 `json:"go_runtime_backed"`
	CCoreAllowed                      bool                                 `json:"c_core_allowed"`
	RubyCoreLogicAllowed              bool                                 `json:"ruby_core_logic_allowed"`
	RubyTestHarness                   bool                                 `json:"ruby_test_harness"`
	KDEPolicyOwner                    bool                                 `json:"kde_policy_owner"`
	KDEPresentationOnly               bool                                 `json:"kde_presentation_only"`
	DeepDesktopForkRequired           bool                                 `json:"deep_desktop_fork_required"`
	GNOMEFirstReleaseSupported        bool                                 `json:"gnome_first_release_supported"`
	XFCEFirstReleaseSupported         bool                                 `json:"xfce_first_release_supported"`
	ProductionDBusOwnershipEnabled    bool                                 `json:"production_dbus_ownership_enabled"`
	RuntimeWriteMethodsEnabled        bool                                 `json:"runtime_write_methods_enabled"`
	BackendLaunchEnabled              bool                                 `json:"backend_launch_enabled"`
	PortalTransportCallsEnabled       bool                                 `json:"portal_transport_calls_enabled"`
	NetworkFetchEnabled               bool                                 `json:"network_fetch_enabled"`
	HostPackageManagerEnabled         bool                                 `json:"host_package_manager_enabled"`
	PrivilegedContainerRequired       bool                                 `json:"privileged_container_required"`
	HostNetworkingRequired            bool                                 `json:"host_networking_required"`
	DockerSocketMounted               bool                                 `json:"docker_socket_mounted"`
	BroadHostMountRequired            bool                                 `json:"broad_host_mount_required"`
	HostRootModified                  bool                                 `json:"host_root_modified"`
	RawExecutablePathExposed          bool                                 `json:"raw_executable_path_exposed"`
	CompatibilityStoragePathExposed   bool                                 `json:"compatibility_storage_path_exposed"`
	BackendCommandExposed             bool                                 `json:"backend_command_exposed"`
	BackendDetailsExposed             bool                                 `json:"backend_details_exposed"`
	ProtectedImplementationPackageDoc string                               `json:"protected_implementation_package_doc"`
	BlockedActions                    []string                             `json:"blocked_actions"`
	NextRecommendedDispatch           []string                             `json:"next_recommended_dispatch"`
	DesktopSafeSummary                string                               `json:"desktop_safe_summary"`
}

type WindowsCompatibilityEntryPoint struct {
	ID                     string `json:"id"`
	Label                  string `json:"label"`
	KDEComponent           string `json:"kde_component"`
	RuntimeSource          string `json:"runtime_source"`
	RequiresRuntimeGate    bool   `json:"requires_runtime_gate"`
	RequiresPortalReview   bool   `json:"requires_portal_review"`
	KDEOwnsPolicy          bool   `json:"kde_owns_policy"`
	HostRootModified       bool   `json:"host_root_modified"`
	BackendLaunchEnabled   bool   `json:"backend_launch_enabled"`
	BackendDetailsExposed  bool   `json:"backend_details_exposed"`
	UserFacingOnly         bool   `json:"user_facing_only"`
	NormalApplicationShape bool   `json:"normal_application_shape"`
}

type WindowsCompatibilityWorkstream struct {
	ID                    string `json:"id"`
	Name                  string `json:"name"`
	PrimaryLayer          string `json:"primary_layer"`
	MainlinePackage       string `json:"mainline_package"`
	FirstWave             bool   `json:"first_wave"`
	CanStartNow           bool   `json:"can_start_now"`
	ImplementationOwner   string `json:"implementation_owner"`
	MainEvidence          string `json:"main_evidence"`
	UnsafeBehaviorEnabled bool   `json:"unsafe_behavior_enabled"`
	BackendLaunchEnabled  bool   `json:"backend_launch_enabled"`
	HostRootModified      bool   `json:"host_root_modified"`
}

type WindowsCompatibilityFirstWaveEntry struct {
	Workstream              string `json:"workstream"`
	MainlinePackage         string `json:"mainline_package"`
	SuggestedBranch         string `json:"suggested_branch"`
	Reason                  string `json:"reason"`
	MinimalMergeableOutcome string `json:"minimal_mergeable_outcome"`
	HostRootModified        bool   `json:"host_root_modified"`
	NetworkRequired         bool   `json:"network_required"`
	BackendLaunchEnabled    bool   `json:"backend_launch_enabled"`
}

func NewWindowsCompatibilityWorkstreamsPreview() (WindowsCompatibilityWorkstreamsPreview, error) {
	entryPoints := windowsCompatibilityEntryPoints()
	workstreams := windowsCompatibilityWorkstreams()
	firstWave := windowsCompatibilityFirstWave(workstreams)
	preview := WindowsCompatibilityWorkstreamsPreview{
		SchemaVersion:                     "xnix.runtime.windows_compatibility_workstreams.v1",
		RequestType:                       "windows-compatibility-workstreams-preview",
		PlanType:                          "kde-first-windows-compatibility-workstreams",
		Source:                            "go-runtime-product-workstream-model",
		ReadMethod:                        "GetWindowsCompatibilityWorkstreamsPreview",
		ProductTarget:                     "best Linux desktop for existing Windows applications",
		OfficialDesktop:                   "KDE Plasma",
		FutureDesktopOrder:                []string{"KDE Plasma", "GNOME", "XFCE"},
		ArchitectureLayers:                []string{"KDE Plasma desktop shell", "KDE integration layer", "Xnix AI Compatibility Runtime", "compatibility backend providers", "Linux system"},
		EntryPoints:                       entryPoints,
		EntryPointIDs:                     windowsCompatibilityEntryPointIDs(entryPoints),
		EntryPointCount:                   len(entryPoints),
		Workstreams:                       workstreams,
		WorkstreamIDs:                     windowsCompatibilityWorkstreamIDs(workstreams),
		WorkstreamCount:                   len(workstreams),
		FirstWaveWorkstreams:              firstWave,
		FirstWaveIDs:                      windowsCompatibilityFirstWaveIDs(firstWave),
		FirstWaveCount:                    len(firstWave),
		RuntimeOwned:                      true,
		GoRuntimeBacked:                   true,
		CCoreAllowed:                      true,
		RubyCoreLogicAllowed:              false,
		RubyTestHarness:                   true,
		KDEPolicyOwner:                    false,
		KDEPresentationOnly:               true,
		DeepDesktopForkRequired:           false,
		GNOMEFirstReleaseSupported:        false,
		XFCEFirstReleaseSupported:         false,
		ProductionDBusOwnershipEnabled:    false,
		RuntimeWriteMethodsEnabled:        false,
		BackendLaunchEnabled:              false,
		PortalTransportCallsEnabled:       false,
		NetworkFetchEnabled:               false,
		HostPackageManagerEnabled:         false,
		PrivilegedContainerRequired:       false,
		HostNetworkingRequired:            false,
		DockerSocketMounted:               false,
		BroadHostMountRequired:            false,
		HostRootModified:                  false,
		RawExecutablePathExposed:          false,
		CompatibilityStoragePathExposed:   false,
		BackendCommandExposed:             false,
		BackendDetailsExposed:             false,
		ProtectedImplementationPackageDoc: "docs/claude-code-implementation-packages.md",
		BlockedActions: []string{
			"fork KDE desktop shell for compatibility policy",
			"let KDE own compatibility backend selection",
			"enable production D-Bus ownership from workstream preview",
			"enable Runtime write methods from workstream preview",
			"start compatibility backend providers from workstream preview",
			"call real Portal transports from workstream preview",
			"fetch network artifacts from workstream preview",
			"mutate host root from workstream preview",
			"expose raw executable paths to desktop shell",
			"modify docs/claude-code-implementation-packages.md from this workstream",
		},
		NextRecommendedDispatch: []string{"CW1", "CW2", "CW3", "CW10"},
		DesktopSafeSummary:      "KDE Plasma is the first official shell while the Go Runtime owns compatibility policy; the first implementation wave is CW1, CW2, CW3, and CW10.",
	}
	if err := validateNoBackendTerms(preview, "Windows compatibility workstreams preview"); err != nil {
		return WindowsCompatibilityWorkstreamsPreview{}, err
	}
	return preview, nil
}

func windowsCompatibilityEntryPoints() []WindowsCompatibilityEntryPoint {
	return []WindowsCompatibilityEntryPoint{
		windowsCompatibilityEntryPoint("start-menu", "Start menu", "Plasma application launcher", "desktop-entry-preview", false),
		windowsCompatibilityEntryPoint("task-manager", "Task manager", "Plasma task manager", "task-manager-identity-preview", false),
		windowsCompatibilityEntryPoint("file-manager", "File manager", "Dolphin", "file-open-preview", true),
		windowsCompatibilityEntryPoint("system-tray", "System tray", "Plasma system tray", "tray-status-preview", false),
		windowsCompatibilityEntryPoint("notification-center", "Notification center", "Plasma notification center", "notification-preview", false),
		windowsCompatibilityEntryPoint("ai-compatibility-center", "AI Compatibility Center", "Plasma widget", "compatibility-center-preview", false),
		windowsCompatibilityEntryPoint("unified-settings", "Unified settings", "KDE system settings", "settings-preview", false),
	}
}

func windowsCompatibilityEntryPoint(id string, label string, component string, source string, portalReview bool) WindowsCompatibilityEntryPoint {
	return WindowsCompatibilityEntryPoint{
		ID:                     id,
		Label:                  label,
		KDEComponent:           component,
		RuntimeSource:          source,
		RequiresRuntimeGate:    true,
		RequiresPortalReview:   portalReview,
		KDEOwnsPolicy:          false,
		HostRootModified:       false,
		BackendLaunchEnabled:   false,
		BackendDetailsExposed:  false,
		UserFacingOnly:         true,
		NormalApplicationShape: true,
	}
}

func windowsCompatibilityWorkstreams() []WindowsCompatibilityWorkstream {
	return []WindowsCompatibilityWorkstream{
		windowsCompatibilityWorkstream("CW1", "Runtime owner read boundary", "Go Runtime", "M1", true, "Go Runtime", "private smoke owner, route parity, disabled writes"),
		windowsCompatibilityWorkstream("CW2", "Recipe and artifact trust pipeline", "Go Runtime", "M2", true, "Go Runtime", "digest-verified recipes and artifact staging receipts"),
		windowsCompatibilityWorkstream("CW3", "Runtime state root and backend lifecycle", "Go Runtime", "M3", true, "Go Runtime", "durable lifecycle states without backend launch"),
		windowsCompatibilityWorkstream("CW4", "KDE seven entry points", "KDE integration plus Runtime reads", "M5", false, "Go Runtime plus KDE integration", "KDE entry points consume Runtime read models"),
		windowsCompatibilityWorkstream("CW5", "Portal permission broker", "Go Runtime plus Portal model", "M4", false, "Go Runtime", "fake-mode request records and review states"),
		windowsCompatibilityWorkstream("CW6", "Snapshot and rollback store", "Go Runtime", "M4", false, "Go Runtime", "content-addressed restore points and rollback receipts"),
		windowsCompatibilityWorkstream("CW7", "AI diagnostics and repair boundary", "Go Runtime", "M7", false, "Go Runtime", "redacted diagnostic inputs and review-only recommendations"),
		windowsCompatibilityWorkstream("CW8", "Execution transaction ledger", "Go Runtime", "M6", false, "Go Runtime", "reviewed transaction records that keep launch blocked"),
		windowsCompatibilityWorkstream("CW9", "KDE materialization writer", "Go Runtime plus KDE files", "M5", false, "Go Runtime plus KDE integration", "desktop, MIME, service-menu, icon, and activation receipts under target roots"),
		windowsCompatibilityWorkstream("CW10", "Evidence and drift harness", "Ruby tooling plus Go fixtures", "M8", true, "Ruby tooling plus Go fixtures", "reports fail on orphan contracts and preview-only regressions"),
		windowsCompatibilityWorkstream("CW11", "Product image and QEMU acceptance", "Build and image tooling", "M9", false, "Build tooling", "KDE-first image smoke with loopback-only access and persisted logs"),
	}
}

func windowsCompatibilityWorkstream(id string, name string, layer string, mainline string, firstWave bool, owner string, evidence string) WindowsCompatibilityWorkstream {
	return WindowsCompatibilityWorkstream{
		ID:                    id,
		Name:                  name,
		PrimaryLayer:          layer,
		MainlinePackage:       mainline,
		FirstWave:             firstWave,
		CanStartNow:           id != "CW8" && id != "CW11",
		ImplementationOwner:   owner,
		MainEvidence:          evidence,
		UnsafeBehaviorEnabled: false,
		BackendLaunchEnabled:  false,
		HostRootModified:      false,
	}
}

func windowsCompatibilityFirstWave(workstreams []WindowsCompatibilityWorkstream) []WindowsCompatibilityFirstWaveEntry {
	return []WindowsCompatibilityFirstWaveEntry{
		windowsCompatibilityFirstWaveEntry(workstreams, "CW1", "codex/cw-runtime-owner-read-boundary", "Runtime read ownership must exist before desktop and execution surfaces depend on it.", "A constrained Go Runtime owner serves read-only Runtime methods through owner routes and keeps writes disabled."),
		windowsCompatibilityFirstWaveEntry(workstreams, "CW2", "codex/cw-recipe-artifact-trust", "Desktop presentation and execution planning must consume verified local inputs.", "Local recipes and fixture artifacts verify digests, stage under explicit roots, and fail closed for invalid inputs."),
		windowsCompatibilityFirstWaveEntry(workstreams, "CW3", "codex/cw-runtime-state-lifecycle", "Application status, diagnostics, and future launch gates need durable lifecycle evidence.", "Runtime state-root lifecycle records expose missing, planned, staged, ready, repair-required, and blocked states without backend launch."),
		windowsCompatibilityFirstWaveEntry(workstreams, "CW10", "codex/cw-evidence-drift-harness", "The KDE-first strategy needs gates that prevent contract-only growth.", "Reports fail on orphan Runtime contracts, KDE entry-point drift, preview-only regressions, and missing workstream ownership."),
	}
}

func windowsCompatibilityFirstWaveEntry(workstreams []WindowsCompatibilityWorkstream, id string, branch string, reason string, outcome string) WindowsCompatibilityFirstWaveEntry {
	mainline := ""
	for _, workstream := range workstreams {
		if workstream.ID == id {
			mainline = workstream.MainlinePackage
			break
		}
	}
	return WindowsCompatibilityFirstWaveEntry{
		Workstream:              id,
		MainlinePackage:         mainline,
		SuggestedBranch:         branch,
		Reason:                  reason,
		MinimalMergeableOutcome: outcome,
		HostRootModified:        false,
		NetworkRequired:         false,
		BackendLaunchEnabled:    false,
	}
}

func windowsCompatibilityEntryPointIDs(entryPoints []WindowsCompatibilityEntryPoint) []string {
	ids := make([]string, 0, len(entryPoints))
	for _, entryPoint := range entryPoints {
		ids = append(ids, entryPoint.ID)
	}
	return ids
}

func windowsCompatibilityWorkstreamIDs(workstreams []WindowsCompatibilityWorkstream) []string {
	ids := make([]string, 0, len(workstreams))
	for _, workstream := range workstreams {
		ids = append(ids, workstream.ID)
	}
	return ids
}

func windowsCompatibilityFirstWaveIDs(firstWave []WindowsCompatibilityFirstWaveEntry) []string {
	ids := make([]string, 0, len(firstWave))
	for _, entry := range firstWave {
		ids = append(ids, entry.Workstream)
	}
	return ids
}
