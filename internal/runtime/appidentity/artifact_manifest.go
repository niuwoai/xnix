package appidentity

type ArtifactManifestPreview struct {
	SchemaVersion               string                   `json:"schema_version"`
	RequestType                 string                   `json:"request_type"`
	ManifestType                string                   `json:"manifest_type"`
	Source                      string                   `json:"source"`
	Desktop                     string                   `json:"desktop"`
	RuntimeMethod               string                   `json:"runtime_method"`
	ReadMethod                  string                   `json:"read_method"`
	Application                 PackageSourceApplication `json:"application"`
	SelectedStrategy            string                   `json:"selected_strategy"`
	ManifestState               string                   `json:"manifest_state"`
	ArtifactGroups              []ArtifactManifestGroup  `json:"artifact_groups"`
	ArtifactGroupIDs            []string                 `json:"artifact_group_ids"`
	RequiredPreflight           []ArtifactManifestCheck  `json:"required_preflight"`
	RequiredPreflightIDs        []string                 `json:"required_preflight_ids"`
	RuntimeOwned                bool                     `json:"runtime_owned"`
	GoRuntimeBacked             bool                     `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                     `json:"kde_policy_owner"`
	UserVisible                 bool                     `json:"user_visible"`
	ManifestReady               bool                     `json:"manifest_ready"`
	SignatureVerified           bool                     `json:"signature_verified"`
	AcquisitionPreflightReady   bool                     `json:"acquisition_preflight_ready"`
	DownloadEnabled             bool                     `json:"download_enabled"`
	InstallEnabled              bool                     `json:"install_enabled"`
	NetworkRequestCreated       bool                     `json:"network_request_created"`
	ArtifactsDownloaded         bool                     `json:"artifacts_downloaded"`
	HostRootModified            bool                     `json:"host_root_modified"`
	PrivilegedContainerRequired bool                     `json:"privileged_container_required"`
	DesktopShellCommandExposed  bool                     `json:"desktop_shell_command_exposed"`
	BackendDetailsExposed       bool                     `json:"backend_details_exposed"`
	UserFacingSettings          map[string]string        `json:"user_facing_settings"`
	BlockedActions              []string                 `json:"blocked_actions"`
	DesktopSafeSummary          string                   `json:"desktop_safe_summary"`
}

type ArtifactManifestGroup struct {
	ID             string `json:"id"`
	Kind           string `json:"kind"`
	Required       bool   `json:"required"`
	Resolved       bool   `json:"resolved"`
	Downloaded     bool   `json:"downloaded"`
	CacheNamespace string `json:"cache_namespace"`
	Summary        string `json:"summary"`
}

type ArtifactManifestCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

func (plan Plan) ArtifactManifestPreview() (ArtifactManifestPreview, error) {
	acquisition, err := plan.AcquisitionPreflightPreview()
	if err != nil {
		return ArtifactManifestPreview{}, err
	}

	groups := artifactManifestGroups(plan.ApplicationID)
	preflight := artifactManifestPreflight()

	preview := ArtifactManifestPreview{
		SchemaVersion:               "xnix.runtime.artifact_manifest.v1",
		RequestType:                 "artifact-manifest-preview",
		ManifestType:                "compatibility-artifact-manifest",
		Source:                      "acquisition-preflight-preview+go-runtime-artifact-manifest",
		Desktop:                     "KDE Plasma",
		RuntimeMethod:               "GetCompatibilityArtifactManifest",
		ReadMethod:                  "GetCompatibilityArtifactManifestPreview",
		Application:                 acquisition.Application,
		SelectedStrategy:            acquisition.SelectedStrategy,
		ManifestState:               "planned",
		ArtifactGroups:              groups,
		ArtifactGroupIDs:            artifactManifestGroupIDs(groups),
		RequiredPreflight:           preflight,
		RequiredPreflightIDs:        artifactManifestPreflightIDs(preflight),
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		UserVisible:                 true,
		ManifestReady:               false,
		SignatureVerified:           false,
		AcquisitionPreflightReady:   acquisition.AcquisitionReady,
		DownloadEnabled:             false,
		InstallEnabled:              false,
		NetworkRequestCreated:       false,
		ArtifactsDownloaded:         false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		DesktopShellCommandExposed:  false,
		BackendDetailsExposed:       false,
		UserFacingSettings:          plan.UserFacingSettings,
		BlockedActions: []string{
			"download artifacts before signed manifest verification",
			"activate artifacts before digest verification",
			"expose artifact cache paths to KDE",
			"mutate host root during artifact manifest planning",
		},
		DesktopSafeSummary: "Compatibility artifact manifest is planned and waiting for acquisition preflight; no download, activation, or host write is enabled.",
	}
	if err := validateNoBackendTerms(preview, "artifact manifest preview"); err != nil {
		return ArtifactManifestPreview{}, err
	}
	return preview, nil
}

func artifactManifestGroups(applicationID string) []ArtifactManifestGroup {
	return []ArtifactManifestGroup{
		{
			ID:             "runtime-launch-metadata",
			Kind:           "metadata",
			Required:       true,
			Resolved:       false,
			Downloaded:     false,
			CacheNamespace: artifactCacheNamespace(applicationID, "launch-metadata"),
			Summary:        "Runtime launch metadata must be resolved before launch binding.",
		},
		{
			ID:             "local-execution-artifacts",
			Kind:           "execution-artifacts",
			Required:       true,
			Resolved:       false,
			Downloaded:     false,
			CacheNamespace: artifactCacheNamespace(applicationID, "local-execution"),
			Summary:        "Local execution artifacts remain planned until signed manifest verification passes.",
		},
		{
			ID:             "isolated-environment-artifacts",
			Kind:           "environment-artifacts",
			Required:       false,
			Resolved:       false,
			Downloaded:     false,
			CacheNamespace: artifactCacheNamespace(applicationID, "isolated-environment"),
			Summary:        "Isolated environment artifacts remain optional until Runtime policy selects them.",
		},
	}
}

func artifactManifestPreflight() []ArtifactManifestCheck {
	return []ArtifactManifestCheck{
		{ID: "acquisition-preflight-ready", Status: "pending", Summary: "Runtime acquisition preflight must be ready before artifact manifest resolution."},
		{ID: "manifest-signature-verification", Status: "required", Summary: "Runtime must verify the signed artifact manifest before artifact use."},
		{ID: "artifact-digest-verification", Status: "required", Summary: "Runtime must verify artifact digests before cache activation."},
		{ID: "cache-namespace-allocation", Status: "pending", Summary: "Runtime must allocate cache namespaces before artifact acquisition."},
		{ID: "rollback-reference", Status: "required", Summary: "Runtime must record rollback references before artifacts can affect state."},
	}
}

func artifactCacheNamespace(applicationID string, suffix string) string {
	sanitized := make([]rune, 0, len(applicationID))
	for _, r := range applicationID {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9', r == '.', r == '-':
			sanitized = append(sanitized, r)
		default:
			sanitized = append(sanitized, '-')
		}
	}
	return string(sanitized) + "." + suffix
}

func artifactManifestGroupIDs(groups []ArtifactManifestGroup) []string {
	ids := make([]string, 0, len(groups))
	for _, group := range groups {
		ids = append(ids, group.ID)
	}
	return ids
}

func artifactManifestPreflightIDs(checks []ArtifactManifestCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}
