package appidentity

type ApplicationsPreview struct {
	SchemaVersion               string                      `json:"schema_version"`
	RequestType                 string                      `json:"request_type"`
	CatalogType                 string                      `json:"catalog_type"`
	Source                      string                      `json:"source"`
	Desktop                     string                      `json:"desktop"`
	RuntimeMethod               string                      `json:"runtime_method"`
	ReadMethod                  string                      `json:"read_method"`
	RegistryName                string                      `json:"registry_name,omitempty"`
	RecipeDigestVerified        bool                        `json:"recipe_digest_verified"`
	RecipeSignatureStatus       string                      `json:"recipe_signature_status,omitempty"`
	ApplicationCount            int                         `json:"application_count"`
	Applications                []RuntimeApplicationPreview `json:"applications"`
	ApplicationIDs              []string                    `json:"application_ids"`
	EntryPointIDs               []string                    `json:"entry_point_ids"`
	RuntimeOwned                bool                        `json:"runtime_owned"`
	GoRuntimeBacked             bool                        `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                        `json:"kde_policy_owner"`
	UserVisible                 bool                        `json:"user_visible"`
	StandardDesktopEntries      bool                        `json:"standard_desktop_entries"`
	RecipeRegistryVerified      bool                        `json:"recipe_registry_verified"`
	BackendTerminologyHidden    bool                        `json:"backend_terminology_hidden"`
	LaunchEnabled               bool                        `json:"launch_enabled"`
	ExecutionStarted            bool                        `json:"execution_started"`
	DesktopFilesWritten         bool                        `json:"desktop_files_written"`
	HostRootModified            bool                        `json:"host_root_modified"`
	NetworkRequired             bool                        `json:"network_required"`
	BackendDetailsExposed       bool                        `json:"backend_details_exposed"`
	RawWindowsExecutableExposed bool                        `json:"raw_windows_executable_exposed"`
	BlockedActions              []string                    `json:"blocked_actions"`
	DesktopSafeSummary          string                      `json:"desktop_safe_summary"`
}

type ApplicationPreview struct {
	SchemaVersion               string                    `json:"schema_version"`
	RequestType                 string                    `json:"request_type"`
	CatalogType                 string                    `json:"catalog_type"`
	Source                      string                    `json:"source"`
	Desktop                     string                    `json:"desktop"`
	RuntimeMethod               string                    `json:"runtime_method"`
	ReadMethod                  string                    `json:"read_method"`
	Application                 RuntimeApplicationPreview `json:"application"`
	RuntimeOwned                bool                      `json:"runtime_owned"`
	GoRuntimeBacked             bool                      `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                      `json:"kde_policy_owner"`
	UserVisible                 bool                      `json:"user_visible"`
	StandardDesktopEntry        bool                      `json:"standard_desktop_entry"`
	RecipeRegistryVerified      bool                      `json:"recipe_registry_verified"`
	BackendTerminologyHidden    bool                      `json:"backend_terminology_hidden"`
	LaunchEnabled               bool                      `json:"launch_enabled"`
	ExecutionStarted            bool                      `json:"execution_started"`
	DesktopFileWritten          bool                      `json:"desktop_file_written"`
	HostRootModified            bool                      `json:"host_root_modified"`
	NetworkRequired             bool                      `json:"network_required"`
	BackendDetailsExposed       bool                      `json:"backend_details_exposed"`
	RawWindowsExecutableExposed bool                      `json:"raw_windows_executable_exposed"`
	BlockedActions              []string                  `json:"blocked_actions"`
	DesktopSafeSummary          string                    `json:"desktop_safe_summary"`
}

type RuntimeApplicationPreview struct {
	ID                          string            `json:"id"`
	Name                        string            `json:"name"`
	Icon                        string            `json:"icon"`
	RuntimeMode                 string            `json:"runtime_mode"`
	DesktopFile                 string            `json:"desktop_file"`
	StartupWMClass              string            `json:"startup_wm_class"`
	Categories                  []string          `json:"categories"`
	MIMETypes                   []string          `json:"mime_types"`
	SupportedExtensions         []string          `json:"supported_extensions"`
	LauncherAction              string            `json:"launcher_action"`
	LaunchCommand               []string          `json:"launch_command"`
	KDEEntryPoints              []string          `json:"kde_entry_points"`
	UserFacingSettings          map[string]string `json:"user_facing_settings"`
	RegistryName                string            `json:"registry_name,omitempty"`
	RecipeDigestVerified        bool              `json:"recipe_digest_verified"`
	RecipeSignatureStatus       string            `json:"recipe_signature_status,omitempty"`
	UserVisible                 bool              `json:"user_visible"`
	StandardDesktopEntry        bool              `json:"standard_desktop_entry"`
	AcceptsFileURIs             bool              `json:"accepts_file_uris"`
	RuntimeOwned                bool              `json:"runtime_owned"`
	BackendTerminologyHidden    bool              `json:"backend_terminology_hidden"`
	BackendLaunchEnabled        bool              `json:"backend_launch_enabled"`
	BackendDetailsExposed       bool              `json:"backend_details_exposed"`
	RawWindowsExecutableExposed bool              `json:"raw_windows_executable_exposed"`
	HostRootMutationEnabled     bool              `json:"host_root_mutation_enabled"`
	StableIdentityDigest        string            `json:"stable_identity_digest"`
}

func NewApplicationsPreview(recipes []Recipe, provenance Provenance) (ApplicationsPreview, error) {
	applications := make([]RuntimeApplicationPreview, 0, len(recipes))
	for _, recipe := range recipes {
		application, err := runtimeApplicationPreview(recipe, provenance)
		if err != nil {
			return ApplicationsPreview{}, err
		}
		applications = append(applications, application)
	}

	preview := ApplicationsPreview{
		SchemaVersion:               "xnix.runtime.applications.v1",
		RequestType:                 "applications-preview",
		CatalogType:                 "runtime-application-catalog",
		Source:                      "registry+go-runtime-desktop-identity",
		Desktop:                     "KDE Plasma",
		RuntimeMethod:               "ListApplications",
		ReadMethod:                  "ListApplicationsPreview",
		RegistryName:                provenance.RegistryName,
		RecipeDigestVerified:        provenance.DigestVerified,
		RecipeSignatureStatus:       provenance.SignatureStatus,
		ApplicationCount:            len(applications),
		Applications:                applications,
		ApplicationIDs:              runtimeApplicationIDs(applications),
		EntryPointIDs:               runtimeApplicationEntryPoints(),
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		UserVisible:                 true,
		StandardDesktopEntries:      true,
		RecipeRegistryVerified:      provenance.DigestVerified,
		BackendTerminologyHidden:    true,
		LaunchEnabled:               false,
		ExecutionStarted:            false,
		DesktopFilesWritten:         false,
		HostRootModified:            false,
		NetworkRequired:             false,
		BackendDetailsExposed:       false,
		RawWindowsExecutableExposed: false,
		BlockedActions: []string{
			"launch application from catalog preview",
			"write desktop files from catalog preview",
			"start compatibility execution from catalog preview",
			"mutate host root during catalog preview",
			"expose backend implementation details from catalog preview",
		},
		DesktopSafeSummary: "Application catalog preview is Go-owned and presents registered compatibility applications as normal Linux applications while launch and host mutation remain disabled.",
	}
	if err := validateNoBackendTerms(preview, "Applications preview"); err != nil {
		return ApplicationsPreview{}, err
	}
	return preview, nil
}

func NewApplicationPreview(recipe Recipe, provenance Provenance) (ApplicationPreview, error) {
	application, err := runtimeApplicationPreview(recipe, provenance)
	if err != nil {
		return ApplicationPreview{}, err
	}

	preview := ApplicationPreview{
		SchemaVersion:               "xnix.runtime.application.v1",
		RequestType:                 "application-preview",
		CatalogType:                 "runtime-application",
		Source:                      "registry+go-runtime-desktop-identity",
		Desktop:                     "KDE Plasma",
		RuntimeMethod:               "GetApplication",
		ReadMethod:                  "GetApplicationPreview",
		Application:                 application,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		UserVisible:                 true,
		StandardDesktopEntry:        true,
		RecipeRegistryVerified:      provenance.DigestVerified,
		BackendTerminologyHidden:    true,
		LaunchEnabled:               false,
		ExecutionStarted:            false,
		DesktopFileWritten:          false,
		HostRootModified:            false,
		NetworkRequired:             false,
		BackendDetailsExposed:       false,
		RawWindowsExecutableExposed: false,
		BlockedActions: []string{
			"launch application from application preview",
			"write desktop file from application preview",
			"start compatibility execution from application preview",
			"mutate host root during application preview",
			"expose backend implementation details from application preview",
		},
		DesktopSafeSummary: "Application preview is Go-owned and presents one registered compatibility application as a normal Linux application while launch and host mutation remain disabled.",
	}
	if err := validateNoBackendTerms(preview, "Application preview"); err != nil {
		return ApplicationPreview{}, err
	}
	return preview, nil
}

func runtimeApplicationPreview(recipe Recipe, provenance Provenance) (RuntimeApplicationPreview, error) {
	plan, err := NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return RuntimeApplicationPreview{}, err
	}
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return RuntimeApplicationPreview{}, err
	}

	return RuntimeApplicationPreview{
		ID:                          plan.ApplicationID,
		Name:                        plan.DisplayName,
		Icon:                        plan.Icon,
		RuntimeMode:                 modeLabel(recipe.Mode),
		DesktopFile:                 plan.DesktopFile,
		StartupWMClass:              plan.StartupWMClass,
		Categories:                  plan.Categories,
		MIMETypes:                   plan.MIMETypes,
		SupportedExtensions:         normalizedExtensions(recipe.SupportedExtensions),
		LauncherAction:              plan.LauncherAction,
		LaunchCommand:               plan.LaunchCommand,
		KDEEntryPoints:              plan.KDEEntryPoints,
		UserFacingSettings:          plan.UserFacingSettings,
		RegistryName:                plan.RegistryName,
		RecipeDigestVerified:        plan.RecipeDigestVerified,
		RecipeSignatureStatus:       plan.RecipeSignatureStatus,
		UserVisible:                 plan.UserVisible,
		StandardDesktopEntry:        plan.StandardDesktopEntry,
		AcceptsFileURIs:             plan.AcceptsFileURIs,
		RuntimeOwned:                plan.RuntimeOwned,
		BackendTerminologyHidden:    plan.BackendTerminologyHidden,
		BackendLaunchEnabled:        plan.BackendLaunchEnabled,
		BackendDetailsExposed:       plan.BackendDetailsExposed,
		RawWindowsExecutableExposed: plan.RawWindowsExecutableExposed,
		HostRootMutationEnabled:     plan.HostRootMutationEnabled,
		StableIdentityDigest:        plan.StableIdentityDigest,
	}, nil
}

func runtimeApplicationIDs(applications []RuntimeApplicationPreview) []string {
	ids := make([]string, 0, len(applications))
	for _, application := range applications {
		ids = append(ids, application.ID)
	}
	return ids
}

func runtimeApplicationEntryPoints() []string {
	return []string{
		"start-menu",
		"task-manager",
		"file-manager",
		"system-tray",
		"notification-center",
		"compatibility-center",
		"unified-settings",
	}
}
