package appidentity

import "errors"

type DesktopActivationManifestPreview struct {
	SchemaVersion               string                          `json:"schema_version"`
	RequestType                 string                          `json:"request_type"`
	ManifestType                string                          `json:"manifest_type"`
	Source                      string                          `json:"source"`
	Desktop                     string                          `json:"desktop"`
	RuntimeMethod               string                          `json:"runtime_method"`
	ReadMethod                  string                          `json:"read_method"`
	ApplicationID               string                          `json:"application_id"`
	DisplayName                 string                          `json:"display_name"`
	Icon                        string                          `json:"icon"`
	DesktopFile                 string                          `json:"desktop_file"`
	LauncherURL                 string                          `json:"launcher_url"`
	Bundle                      DesktopActivationManifestBundle `json:"bundle"`
	EntryPoints                 []KDEEntryPointPreview          `json:"entry_points"`
	EntryPointIDs               []string                        `json:"entry_point_ids"`
	EntryPointCount             int                             `json:"entry_point_count"`
	VisibleEntryPointCount      int                             `json:"visible_entry_point_count"`
	ActivationMaterials         []DesktopActivationMaterial     `json:"activation_materials"`
	ActivationMaterialIDs       []string                        `json:"activation_material_ids"`
	ActivationMaterialCount     int                             `json:"activation_material_count"`
	ContractSectionIDs          []string                        `json:"contract_section_ids"`
	ContractSectionCount        int                             `json:"contract_section_count"`
	RuntimeOwned                bool                            `json:"runtime_owned"`
	GoRuntimeBacked             bool                            `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                            `json:"kde_policy_owner"`
	UserVisible                 bool                            `json:"user_visible"`
	OfficialDesktopOnly         bool                            `json:"official_desktop_only"`
	StableDesktopContract       bool                            `json:"stable_desktop_contract"`
	NormalApplicationSurface    bool                            `json:"normal_application_surface"`
	StandardDesktopEntry        bool                            `json:"standard_desktop_entry"`
	SevenEntryPointContract     bool                            `json:"seven_entry_point_contract"`
	PortalMediatedFileAccess    bool                            `json:"portal_mediated_file_access"`
	DesktopFilesWritten         bool                            `json:"desktop_files_written"`
	MIMEAppsWritten             bool                            `json:"mimeapps_written"`
	ManifestWritten             bool                            `json:"manifest_written"`
	SettingsPersisted           bool                            `json:"settings_persisted"`
	NotificationsSent           bool                            `json:"notifications_sent"`
	TaskManagerEntryActive      bool                            `json:"task_manager_entry_active"`
	KWinRuleApplied             bool                            `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled       bool                            `json:"live_tray_bridge_enabled"`
	LaunchEnabled               bool                            `json:"launch_enabled"`
	BackendLaunchEnabled        bool                            `json:"backend_launch_enabled"`
	ExecutionStarted            bool                            `json:"execution_started"`
	HostRootModified            bool                            `json:"host_root_modified"`
	NetworkRequired             bool                            `json:"network_required"`
	PrivilegedContainerRequired bool                            `json:"privileged_container_required"`
	BackendDetailsExposed       bool                            `json:"backend_details_exposed"`
	RawWindowsExecutableExposed bool                            `json:"raw_windows_executable_exposed"`
	CompatibilityStorageExposed bool                            `json:"compatibility_storage_exposed"`
	BlockedActions              []string                        `json:"blocked_actions"`
	DesktopSafeSummary          string                          `json:"desktop_safe_summary"`
}

type DesktopActivationManifestBundle struct {
	RequestType              string   `json:"request_type"`
	MaterialCount            int      `json:"material_count"`
	MaterialIDs              []string `json:"material_ids"`
	StandardDesktopEntry     bool     `json:"standard_desktop_entry"`
	FileAssociationReady     bool     `json:"file_association_ready"`
	TaskManagerIdentityReady bool     `json:"task_manager_identity_ready"`
	KWinIdentityReady        bool     `json:"kwin_identity_ready"`
	TrayStatusReady          bool     `json:"tray_status_ready"`
	NotificationReady        bool     `json:"notification_ready"`
	SettingsReady            bool     `json:"settings_ready"`
	CompatibilityCenterReady bool     `json:"compatibility_center_ready"`
	DesktopFilesWritten      bool     `json:"desktop_files_written"`
	HostRootModified         bool     `json:"host_root_modified"`
	BackendDetailsExposed    bool     `json:"backend_details_exposed"`
}

func (plan Plan) DesktopActivationManifestPreview() (DesktopActivationManifestPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return DesktopActivationManifestPreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return DesktopActivationManifestPreview{}, errors.New("desktop activation manifest requires single-line identity fields")
		}
	}

	bundle, err := plan.DesktopActivationBundlePreview()
	if err != nil {
		return DesktopActivationManifestPreview{}, err
	}
	entrypoints, err := plan.KDEEntryPointsPreview("deferred", nil)
	if err != nil {
		return DesktopActivationManifestPreview{}, err
	}
	contractSections := desktopActivationManifestContractSections()
	preview := DesktopActivationManifestPreview{
		SchemaVersion:               "xnix.runtime.desktop_activation_manifest.v1",
		RequestType:                 "desktop-activation-manifest-preview",
		ManifestType:                "kde-desktop-activation-manifest",
		Source:                      "desktop-activation-bundle-preview+kde-entrypoints-preview",
		Desktop:                     "KDE Plasma",
		RuntimeMethod:               "GetDesktopActivationManifest",
		ReadMethod:                  "GetDesktopActivationManifestPreview",
		ApplicationID:               plan.ApplicationID,
		DisplayName:                 plan.DisplayName,
		Icon:                        plan.Icon,
		DesktopFile:                 plan.DesktopFile,
		LauncherURL:                 "applications:" + plan.DesktopFile,
		Bundle:                      desktopActivationManifestBundle(bundle),
		EntryPoints:                 entrypoints.EntryPoints,
		EntryPointIDs:               entrypoints.EntryPointIDs,
		EntryPointCount:             entrypoints.EntryPointCount,
		VisibleEntryPointCount:      entrypoints.VisibleEntryPointCount,
		ActivationMaterials:         bundle.Materials,
		ActivationMaterialIDs:       bundle.MaterialIDs,
		ActivationMaterialCount:     bundle.MaterialCount,
		ContractSectionIDs:          contractSections,
		ContractSectionCount:        len(contractSections),
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		UserVisible:                 true,
		OfficialDesktopOnly:         true,
		StableDesktopContract:       true,
		NormalApplicationSurface:    true,
		StandardDesktopEntry:        true,
		SevenEntryPointContract:     entrypoints.EntryPointCount == 7,
		PortalMediatedFileAccess:    entrypoints.PortalEntryPointCount > 0,
		DesktopFilesWritten:         false,
		MIMEAppsWritten:             false,
		ManifestWritten:             false,
		SettingsPersisted:           false,
		NotificationsSent:           false,
		TaskManagerEntryActive:      false,
		KWinRuleApplied:             false,
		LiveTrayBridgeEnabled:       false,
		LaunchEnabled:               false,
		BackendLaunchEnabled:        false,
		ExecutionStarted:            false,
		HostRootModified:            false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		RawWindowsExecutableExposed: false,
		CompatibilityStorageExposed: false,
		BlockedActions: []string{
			"write desktop activation files from manifest preview",
			"write MIME defaults from manifest preview",
			"write activation manifest from manifest preview",
			"persist compatibility settings from manifest preview",
			"send desktop notifications from manifest preview",
			"activate task manager entry before a live session",
			"apply KWin rules before Runtime launch approval",
			"enable live tray bridge from manifest preview",
			"start compatibility backend from manifest preview",
			"mutate host root during desktop activation manifest preview",
			"expose raw backend command to desktop shell",
		},
		DesktopSafeSummary: "KDE can read the seven-entry desktop activation manifest for this compatibility application, but Runtime preview does not write desktop files, start execution, or mutate the host root.",
	}
	if err := validateNoBackendTerms(preview, "desktop activation manifest preview"); err != nil {
		return DesktopActivationManifestPreview{}, err
	}
	return preview, nil
}

func desktopActivationManifestBundle(bundle DesktopActivationBundlePreview) DesktopActivationManifestBundle {
	return DesktopActivationManifestBundle{
		RequestType:              bundle.RequestType,
		MaterialCount:            bundle.MaterialCount,
		MaterialIDs:              bundle.MaterialIDs,
		StandardDesktopEntry:     bundle.StandardDesktopEntry,
		FileAssociationReady:     bundle.FileAssociationReady,
		TaskManagerIdentityReady: bundle.TaskManagerIdentityReady,
		KWinIdentityReady:        bundle.KWinIdentityReady,
		TrayStatusReady:          bundle.TrayStatusReady,
		NotificationReady:        bundle.NotificationReady,
		SettingsReady:            bundle.SettingsReady,
		CompatibilityCenterReady: bundle.CompatibilityCenterReady,
		DesktopFilesWritten:      bundle.DesktopFilesWritten,
		HostRootModified:         bundle.HostRootModified,
		BackendDetailsExposed:    bundle.BackendDetailsExposed,
	}
}

func desktopActivationManifestContractSections() []string {
	return []string{
		"launcher",
		"task-manager",
		"file-manager",
		"system-tray",
		"notifications",
		"compatibility-center",
		"unified-settings",
	}
}
