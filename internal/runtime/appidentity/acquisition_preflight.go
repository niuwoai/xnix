package appidentity

type AcquisitionPreflightPreview struct {
	SchemaVersion               string                      `json:"schema_version"`
	RequestType                 string                      `json:"request_type"`
	PreflightType               string                      `json:"preflight_type"`
	Source                      string                      `json:"source"`
	Desktop                     string                      `json:"desktop"`
	RuntimeMethod               string                      `json:"runtime_method"`
	ReadMethod                  string                      `json:"read_method"`
	Application                 PackageSourceApplication    `json:"application"`
	SelectedStrategy            string                      `json:"selected_strategy"`
	PreflightState              string                      `json:"preflight_state"`
	Checks                      []AcquisitionPreflightCheck `json:"checks"`
	CheckIDs                    []string                    `json:"check_ids"`
	RuntimeOwned                bool                        `json:"runtime_owned"`
	GoRuntimeBacked             bool                        `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                        `json:"kde_policy_owner"`
	UserVisible                 bool                        `json:"user_visible"`
	AcquisitionReady            bool                        `json:"acquisition_ready"`
	PackageSourceReady          bool                        `json:"package_source_ready"`
	DownloadEnabled             bool                        `json:"download_enabled"`
	InstallEnabled              bool                        `json:"install_enabled"`
	NetworkRequiredForPlanning  bool                        `json:"network_required_for_planning"`
	NetworkRequestCreated       bool                        `json:"network_request_created"`
	ArtifactsDownloaded         bool                        `json:"artifacts_downloaded"`
	HostRootModified            bool                        `json:"host_root_modified"`
	PrivilegedContainerRequired bool                        `json:"privileged_container_required"`
	DesktopShellCommandExposed  bool                        `json:"desktop_shell_command_exposed"`
	BackendDetailsExposed       bool                        `json:"backend_details_exposed"`
	UserFacingSettings          map[string]string           `json:"user_facing_settings"`
	BlockedActions              []string                    `json:"blocked_actions"`
	DesktopSafeSummary          string                      `json:"desktop_safe_summary"`
}

type AcquisitionPreflightCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

func (plan Plan) AcquisitionPreflightPreview() (AcquisitionPreflightPreview, error) {
	packageSource, err := plan.PackageSourcePreview()
	if err != nil {
		return AcquisitionPreflightPreview{}, err
	}

	checks := acquisitionPreflightChecks()
	preview := AcquisitionPreflightPreview{
		SchemaVersion:               "xnix.runtime.acquisition_preflight.v1",
		RequestType:                 "acquisition-preflight-preview",
		PreflightType:               "compatibility-acquisition-preflight",
		Source:                      "package-source-preview+go-runtime-acquisition-preflight",
		Desktop:                     "KDE Plasma",
		RuntimeMethod:               "GetCompatibilityAcquisitionPreflight",
		ReadMethod:                  "GetCompatibilityAcquisitionPreflightPreview",
		Application:                 packageSource.Application,
		SelectedStrategy:            packageSource.SelectedStrategy,
		PreflightState:              "planned",
		Checks:                      checks,
		CheckIDs:                    acquisitionPreflightCheckIDs(checks),
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		UserVisible:                 true,
		AcquisitionReady:            false,
		PackageSourceReady:          packageSource.PackageSourceReady,
		DownloadEnabled:             false,
		InstallEnabled:              false,
		NetworkRequiredForPlanning:  false,
		NetworkRequestCreated:       false,
		ArtifactsDownloaded:         false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		DesktopShellCommandExposed:  false,
		BackendDetailsExposed:       false,
		UserFacingSettings:          plan.UserFacingSettings,
		BlockedActions: []string{
			"download compatibility artifacts before acquisition preflight",
			"install compatibility artifacts before signed manifest verification",
			"invoke network access from KDE",
			"mutate host root during acquisition preflight",
		},
		DesktopSafeSummary: "Compatibility acquisition preflight is planned and waiting for Runtime source readiness; no download, install, or host write is enabled.",
	}
	if err := validateNoBackendTerms(preview, "acquisition preflight preview"); err != nil {
		return AcquisitionPreflightPreview{}, err
	}
	return preview, nil
}

func acquisitionPreflightChecks() []AcquisitionPreflightCheck {
	return []AcquisitionPreflightCheck{
		{ID: "package-source-ready", Status: "pending", Summary: "Runtime package source selection must be ready before acquisition."},
		{ID: "signed-artifact-manifest", Status: "required", Summary: "Runtime must verify a signed artifact manifest before acquisition."},
		{ID: "runtime-cache-space", Status: "pending", Summary: "Runtime cache capacity must be checked before artifact acquisition."},
		{ID: "network-policy-review", Status: "required", Summary: "Runtime must approve network policy before any acquisition request."},
		{ID: "rollback-marker", Status: "required", Summary: "Runtime must define rollback markers before acquisition can change state."},
	}
}

func acquisitionPreflightCheckIDs(checks []AcquisitionPreflightCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}
