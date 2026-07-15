package appidentity

import "errors"

type PackageSourcePreview struct {
	SchemaVersion               string                   `json:"schema_version"`
	RequestType                 string                   `json:"request_type"`
	SourceType                  string                   `json:"source_type"`
	Source                      string                   `json:"source"`
	Desktop                     string                   `json:"desktop"`
	RuntimeMethod               string                   `json:"runtime_method"`
	ReadMethod                  string                   `json:"read_method"`
	Application                 PackageSourceApplication `json:"application"`
	SelectedStrategy            string                   `json:"selected_strategy"`
	SourceSelectionState        string                   `json:"source_selection_state"`
	SourcePolicy                PackageSourcePolicy      `json:"source_policy"`
	SourceChannels              []PackageSourceChannel   `json:"source_channels"`
	SourceChannelIDs            []string                 `json:"source_channel_ids"`
	RequiredPreflight           []PackageSourcePreflight `json:"required_preflight"`
	RequiredPreflightIDs        []string                 `json:"required_preflight_ids"`
	RuntimeOwned                bool                     `json:"runtime_owned"`
	GoRuntimeBacked             bool                     `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                     `json:"kde_policy_owner"`
	UserVisible                 bool                     `json:"user_visible"`
	PackageSourceReady          bool                     `json:"package_source_ready"`
	InstallEnabled              bool                     `json:"install_enabled"`
	NetworkRequiredForPlanning  bool                     `json:"network_required_for_planning"`
	HostRootModified            bool                     `json:"host_root_modified"`
	PrivilegedContainerRequired bool                     `json:"privileged_container_required"`
	DesktopShellCommandExposed  bool                     `json:"desktop_shell_command_exposed"`
	BackendDetailsExposed       bool                     `json:"backend_details_exposed"`
	UserFacingSettings          map[string]string        `json:"user_facing_settings"`
	BlockedActions              []string                 `json:"blocked_actions"`
	DesktopSafeSummary          string                   `json:"desktop_safe_summary"`
}

type PackageSourceApplication struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Icon          string `json:"icon"`
	DesktopFile   string `json:"desktop_file"`
	RequestedMode string `json:"requested_mode"`
}

type PackageSourcePolicy struct {
	SignedSourceRequired      bool `json:"signed_source_required"`
	RuntimeCacheRequired      bool `json:"runtime_cache_required"`
	DirectDesktopInstallAllow bool `json:"direct_desktop_install_allowed"`
	UserVisibleBackendNames   bool `json:"user_visible_backend_names"`
	HostPackageManagerInvoked bool `json:"host_package_manager_invoked"`
}

type PackageSourceChannel struct {
	ID                  string   `json:"id"`
	Kind                string   `json:"kind"`
	RuntimeOwned        bool     `json:"runtime_owned"`
	SelectionState      string   `json:"selection_state"`
	SupportedStrategies []string `json:"supported_strategies"`
	Summary             string   `json:"summary"`
}

type PackageSourcePreflight struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

func (plan Plan) PackageSourcePreview() (PackageSourcePreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return PackageSourcePreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return PackageSourcePreview{}, errors.New("package source preview requires single-line identity fields")
		}
	}

	channels := packageSourceChannels()
	preflight := packageSourcePreflight()

	preview := PackageSourcePreview{
		SchemaVersion: "xnix.runtime.package_source.v1",
		RequestType:   "package-source-preview",
		SourceType:    "compatibility-package-source",
		Source:        "registry+go-runtime-package-source",
		Desktop:       "KDE Plasma",
		RuntimeMethod: "GetCompatibilityPackageSource",
		ReadMethod:    "GetCompatibilityPackageSourcePreview",
		Application: PackageSourceApplication{
			ID:            plan.ApplicationID,
			Name:          plan.DisplayName,
			Icon:          plan.Icon,
			DesktopFile:   plan.DesktopFile,
			RequestedMode: plan.runtimeModeID(),
		},
		SelectedStrategy:     plan.runPlanStrategy(),
		SourceSelectionState: "planned",
		SourcePolicy: PackageSourcePolicy{
			SignedSourceRequired:      true,
			RuntimeCacheRequired:      true,
			DirectDesktopInstallAllow: false,
			UserVisibleBackendNames:   false,
			HostPackageManagerInvoked: false,
		},
		SourceChannels:              channels,
		SourceChannelIDs:            packageSourceChannelIDs(channels),
		RequiredPreflight:           preflight,
		RequiredPreflightIDs:        packageSourcePreflightIDs(preflight),
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		UserVisible:                 true,
		PackageSourceReady:          false,
		InstallEnabled:              false,
		NetworkRequiredForPlanning:  false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		DesktopShellCommandExposed:  false,
		BackendDetailsExposed:       false,
		UserFacingSettings:          plan.UserFacingSettings,
		BlockedActions: []string{
			"install compatibility packages without Runtime source selection",
			"expose package source commands to KDE",
			"use unsigned package sources",
			"mutate the host root during package-source planning",
		},
		DesktopSafeSummary: "Compatibility package source selection is planned and Runtime-owned; no install, network, or host write is enabled.",
	}
	if err := validateNoBackendTerms(preview, "package source preview"); err != nil {
		return PackageSourcePreview{}, err
	}
	return preview, nil
}

func packageSourceChannels() []PackageSourceChannel {
	return []PackageSourceChannel{
		{
			ID:                  "os-managed-compatibility-packages",
			Kind:                "distribution-packages",
			RuntimeOwned:        true,
			SelectionState:      "planned",
			SupportedStrategies: []string{"local-compatible-managed"},
			Summary:             "Distribution-provided compatibility packages can be selected only through Runtime policy.",
		},
		{
			ID:                  "runtime-managed-toolcache",
			Kind:                "runtime-cache",
			RuntimeOwned:        true,
			SelectionState:      "planned",
			SupportedStrategies: []string{"automatic-managed", "local-compatible-managed", "isolated-compatible-managed"},
			Summary:             "Runtime-managed tool cache keeps package selection outside desktop shell code.",
		},
		{
			ID:                  "isolated-environment-template-catalog",
			Kind:                "template-catalog",
			RuntimeOwned:        true,
			SelectionState:      "planned",
			SupportedStrategies: []string{"isolated-compatible-managed"},
			Summary:             "Isolated environment templates remain Runtime-owned and are not launched during planning.",
		},
	}
}

func packageSourcePreflight() []PackageSourcePreflight {
	return []PackageSourcePreflight{
		{ID: "signed-source-verification", Status: "required", Summary: "Runtime must verify a signed source before package installation is enabled."},
		{ID: "source-policy-review", Status: "pending", Summary: "Runtime policy must select the package source before launch binding."},
		{ID: "runtime-cache-quota", Status: "pending", Summary: "Runtime cache quota must be checked before package acquisition."},
		{ID: "offline-fallback", Status: "pending", Summary: "Runtime must define the offline behavior before package acquisition."},
	}
}

func packageSourceChannelIDs(channels []PackageSourceChannel) []string {
	ids := make([]string, 0, len(channels))
	for _, channel := range channels {
		ids = append(ids, channel.ID)
	}
	return ids
}

func packageSourcePreflightIDs(preflight []PackageSourcePreflight) []string {
	ids := make([]string, 0, len(preflight))
	for _, check := range preflight {
		ids = append(ids, check.ID)
	}
	return ids
}
