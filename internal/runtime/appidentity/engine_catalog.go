package appidentity

type EngineCatalogPreview struct {
	SchemaVersion            string                       `json:"schema_version"`
	RequestType              string                       `json:"request_type"`
	CatalogType              string                       `json:"catalog_type"`
	Source                   string                       `json:"source"`
	Desktop                  string                       `json:"desktop"`
	RuntimeMethod            string                       `json:"runtime_method"`
	ReadMethod               string                       `json:"read_method"`
	Engines                  []CompatibilityEnginePreview `json:"engines"`
	EngineIDs                []string                     `json:"engine_ids"`
	DefaultEngineID          string                       `json:"default_engine_id"`
	EngineCount              int                          `json:"engine_count"`
	RuntimeOwned             bool                         `json:"runtime_owned"`
	GoRuntimeBacked          bool                         `json:"go_runtime_backed"`
	KDEPolicyOwner           bool                         `json:"kde_policy_owner"`
	UserVisible              bool                         `json:"user_visible"`
	BackendTerminologyHidden bool                         `json:"backend_terminology_hidden"`
	SelectionPersisted       bool                         `json:"selection_persisted"`
	BackendInstalled         bool                         `json:"backend_installed"`
	BackendLaunchEnabled     bool                         `json:"backend_launch_enabled"`
	ExecutionStarted         bool                         `json:"execution_started"`
	HostRootModified         bool                         `json:"host_root_modified"`
	NetworkRequired          bool                         `json:"network_required"`
	BackendDetailsExposed    bool                         `json:"backend_details_exposed"`
	RawCommandExposed        bool                         `json:"raw_command_exposed"`
	BlockedActions           []string                     `json:"blocked_actions"`
	DesktopSafeSummary       string                       `json:"desktop_safe_summary"`
}

type CompatibilityEnginePreview struct {
	ID                    string `json:"id"`
	Label                 string `json:"label"`
	Strategy              string `json:"strategy"`
	Isolation             string `json:"isolation"`
	Availability          string `json:"availability"`
	Status                string `json:"status"`
	Summary               string `json:"summary"`
	Recommended           bool   `json:"recommended"`
	UserSelectable        bool   `json:"user_selectable"`
	RequiresInstall       bool   `json:"requires_install"`
	RequiresRuntimeReview bool   `json:"requires_runtime_review"`
	BackendInstalled      bool   `json:"backend_installed"`
	BackendLaunchEnabled  bool   `json:"backend_launch_enabled"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
	RawCommandExposed     bool   `json:"raw_command_exposed"`
}

func NewEngineCatalogPreview() (EngineCatalogPreview, error) {
	engines := []CompatibilityEnginePreview{
		compatibilityEnginePreview(
			"automatic",
			"Automatic",
			"runtime-selected",
			"policy-managed",
			"Runtime will choose a compatibility strategy after all launch gates and user reviews are complete.",
			true,
		),
		compatibilityEnginePreview(
			"local-compatible",
			"Local compatibility",
			"local-compatible",
			"shared-desktop-session",
			"Local compatibility keeps the application close to the desktop session while Runtime policy keeps execution disabled in preview.",
			false,
		),
		compatibilityEnginePreview(
			"isolated-compatible",
			"Isolated compatibility",
			"isolated-compatible",
			"separate-runtime-environment",
			"Isolated compatibility reserves a separate Runtime environment while install and launch remain blocked in preview.",
			false,
		),
	}
	preview := EngineCatalogPreview{
		SchemaVersion:            "xnix.runtime.engine_catalog.v1",
		RequestType:              "engine-catalog-preview",
		CatalogType:              "compatibility-engine-catalog",
		Source:                   "go-runtime-engine-catalog",
		Desktop:                  "KDE Plasma",
		RuntimeMethod:            "GetEngineCatalog",
		ReadMethod:               "GetEngineCatalogPreview",
		Engines:                  engines,
		EngineIDs:                engineCatalogIDs(engines),
		DefaultEngineID:          "automatic",
		EngineCount:              len(engines),
		RuntimeOwned:             true,
		GoRuntimeBacked:          true,
		KDEPolicyOwner:           false,
		UserVisible:              true,
		BackendTerminologyHidden: true,
		SelectionPersisted:       false,
		BackendInstalled:         false,
		BackendLaunchEnabled:     false,
		ExecutionStarted:         false,
		HostRootModified:         false,
		NetworkRequired:          false,
		BackendDetailsExposed:    false,
		RawCommandExposed:        false,
		BlockedActions: []string{
			"persist compatibility engine selection from preview",
			"install compatibility engine from preview",
			"launch compatibility engine from preview",
			"expose raw engine commands to KDE",
			"mutate host root during engine catalog preview",
		},
		DesktopSafeSummary: "Engine catalog preview is Go-owned and exposes user-facing compatibility choices while installation, launch, persistence, and backend details remain disabled.",
	}
	if err := validateNoBackendTerms(preview, "Engine catalog preview"); err != nil {
		return EngineCatalogPreview{}, err
	}
	return preview, nil
}

func compatibilityEnginePreview(id string, label string, strategy string, isolation string, summary string, recommended bool) CompatibilityEnginePreview {
	return CompatibilityEnginePreview{
		ID:                    id,
		Label:                 label,
		Strategy:              strategy,
		Isolation:             isolation,
		Availability:          "planned",
		Status:                "ready-for-selection",
		Summary:               summary,
		Recommended:           recommended,
		UserSelectable:        true,
		RequiresInstall:       true,
		RequiresRuntimeReview: true,
		BackendInstalled:      false,
		BackendLaunchEnabled:  false,
		BackendDetailsExposed: false,
		RawCommandExposed:     false,
	}
}

func engineCatalogIDs(engines []CompatibilityEnginePreview) []string {
	ids := make([]string, 0, len(engines))
	for _, engine := range engines {
		ids = append(ids, engine.ID)
	}
	return ids
}
