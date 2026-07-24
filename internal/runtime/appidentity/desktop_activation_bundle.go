package appidentity

import "errors"

type DesktopActivationBundlePreview struct {
	SchemaVersion               string                      `json:"schema_version"`
	RequestType                 string                      `json:"request_type"`
	PlanType                    string                      `json:"plan_type"`
	Source                      string                      `json:"source"`
	Desktop                     string                      `json:"desktop"`
	RuntimeMethod               string                      `json:"runtime_method"`
	ApplicationID               string                      `json:"application_id"`
	DisplayName                 string                      `json:"display_name"`
	Icon                        string                      `json:"icon"`
	DesktopFile                 string                      `json:"desktop_file"`
	LauncherURL                 string                      `json:"launcher_url"`
	LaunchCommand               []string                    `json:"launch_command"`
	MIMETypes                   []string                    `json:"mime_types"`
	DesktopEntryPreview         string                      `json:"desktop_entry_preview"`
	MIMEAppsPreview             string                      `json:"mimeapps_preview"`
	DesktopIcon                 DesktopIconPreview          `json:"desktop_icon"`
	WindowIdentity              WindowIdentityPreview       `json:"window_identity"`
	TrayStatus                  TrayStatusPreview           `json:"tray_status"`
	Notification                NotificationPreview         `json:"notification"`
	Settings                    SettingsPreview             `json:"settings"`
	Materials                   []DesktopActivationMaterial `json:"materials"`
	MaterialIDs                 []string                    `json:"material_ids"`
	MaterialCount               int                         `json:"material_count"`
	RuntimeOwned                bool                        `json:"runtime_owned"`
	GoRuntimeBacked             bool                        `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                        `json:"kde_policy_owner"`
	OfficialDesktopOnly         bool                        `json:"official_desktop_only"`
	NormalApplicationSurface    bool                        `json:"normal_application_surface"`
	StandardDesktopEntry        bool                        `json:"standard_desktop_entry"`
	FileAssociationReady        bool                        `json:"file_association_ready"`
	TaskManagerIdentityReady    bool                        `json:"task_manager_identity_ready"`
	KWinIdentityReady           bool                        `json:"kwin_identity_ready"`
	TrayStatusReady             bool                        `json:"tray_status_ready"`
	NotificationReady           bool                        `json:"notification_ready"`
	SettingsReady               bool                        `json:"settings_ready"`
	CompatibilityCenterReady    bool                        `json:"compatibility_center_ready"`
	DesktopFilesWritten         bool                        `json:"desktop_files_written"`
	MIMEAppsWritten             bool                        `json:"mimeapps_written"`
	SettingsPersisted           bool                        `json:"settings_persisted"`
	NotificationsSent           bool                        `json:"notifications_sent"`
	TaskManagerEntryActive      bool                        `json:"task_manager_entry_active"`
	KWinRuleApplied             bool                        `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled       bool                        `json:"live_tray_bridge_enabled"`
	LaunchEnabled               bool                        `json:"launch_enabled"`
	BackendLaunchEnabled        bool                        `json:"backend_launch_enabled"`
	ExecutionStarted            bool                        `json:"execution_started"`
	HostRootModified            bool                        `json:"host_root_modified"`
	BackendDetailsExposed       bool                        `json:"backend_details_exposed"`
	RawWindowsExecutableExposed bool                        `json:"raw_windows_executable_exposed"`
	CompatibilityStorageExposed bool                        `json:"compatibility_storage_exposed"`
	BlockedActions              []string                    `json:"blocked_actions"`
	DesktopSafeSummary          string                      `json:"desktop_safe_summary"`
}

type DesktopActivationMaterial struct {
	ID                    string `json:"id"`
	Label                 string `json:"label"`
	KDEComponent          string `json:"kde_component"`
	RuntimeSource         string `json:"runtime_source"`
	RuntimeMethod         string `json:"runtime_method"`
	State                 string `json:"state"`
	UserVisible           bool   `json:"user_visible"`
	RequiredForNormalApp  bool   `json:"required_for_normal_app"`
	WritesHost            bool   `json:"writes_host"`
	StartsBackend         bool   `json:"starts_backend"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

func (plan Plan) DesktopActivationBundlePreview() (DesktopActivationBundlePreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return DesktopActivationBundlePreview{}, err
	}
	if len(plan.LaunchCommand) != 4 {
		return DesktopActivationBundlePreview{}, errors.New("desktop activation bundle preview requires a complete managed launcher command")
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return DesktopActivationBundlePreview{}, errors.New("desktop activation bundle preview requires single-line identity fields")
		}
	}

	desktopEntry, err := plan.RenderDesktopEntry()
	if err != nil {
		return DesktopActivationBundlePreview{}, err
	}
	mimeapps := ""
	if len(plan.MIMETypes) > 0 {
		rendered, err := plan.RenderMIMEApps()
		if err != nil {
			return DesktopActivationBundlePreview{}, err
		}
		mimeapps = rendered
	}
	desktopIcon, err := plan.DesktopIconPreview()
	if err != nil {
		return DesktopActivationBundlePreview{}, err
	}
	windowIdentity, err := plan.WindowIdentityPreview()
	if err != nil {
		return DesktopActivationBundlePreview{}, err
	}
	trayStatus, err := plan.TrayStatusPreview()
	if err != nil {
		return DesktopActivationBundlePreview{}, err
	}
	notification, err := plan.NotificationPreview("approval-required")
	if err != nil {
		return DesktopActivationBundlePreview{}, err
	}
	settings, err := plan.SettingsPreview()
	if err != nil {
		return DesktopActivationBundlePreview{}, err
	}

	materials := desktopActivationMaterials()
	preview := DesktopActivationBundlePreview{
		SchemaVersion:               "xnix.runtime.desktop_activation_bundle.v1",
		RequestType:                 "desktop-activation-bundle-preview",
		PlanType:                    "normal-linux-application-activation",
		Source:                      "desktop-entry-preview+mimeapps-preview+window-identity-preview+tray-status-preview+notification-preview+settings-preview",
		Desktop:                     "KDE Plasma",
		RuntimeMethod:               "GetDesktopActivationBundlePreview",
		ApplicationID:               plan.ApplicationID,
		DisplayName:                 plan.DisplayName,
		Icon:                        plan.Icon,
		DesktopFile:                 plan.DesktopFile,
		LauncherURL:                 "applications:" + plan.DesktopFile,
		LaunchCommand:               plan.LaunchCommand,
		MIMETypes:                   plan.MIMETypes,
		DesktopEntryPreview:         desktopEntry,
		MIMEAppsPreview:             mimeapps,
		DesktopIcon:                 desktopIcon,
		WindowIdentity:              windowIdentity,
		TrayStatus:                  trayStatus,
		Notification:                notification,
		Settings:                    settings,
		Materials:                   materials,
		MaterialIDs:                 desktopActivationMaterialIDs(materials),
		MaterialCount:               len(materials),
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		OfficialDesktopOnly:         true,
		NormalApplicationSurface:    true,
		StandardDesktopEntry:        true,
		FileAssociationReady:        true,
		TaskManagerIdentityReady:    true,
		KWinIdentityReady:           true,
		TrayStatusReady:             true,
		NotificationReady:           true,
		SettingsReady:               true,
		CompatibilityCenterReady:    true,
		DesktopFilesWritten:         false,
		MIMEAppsWritten:             false,
		SettingsPersisted:           false,
		NotificationsSent:           false,
		TaskManagerEntryActive:      false,
		KWinRuleApplied:             false,
		LiveTrayBridgeEnabled:       false,
		LaunchEnabled:               false,
		BackendLaunchEnabled:        false,
		ExecutionStarted:            false,
		HostRootModified:            false,
		BackendDetailsExposed:       false,
		RawWindowsExecutableExposed: false,
		CompatibilityStorageExposed: false,
		BlockedActions: []string{
			"write desktop activation files from preview",
			"overwrite MIME defaults from preview",
			"activate task manager entry before a live session",
			"apply KWin rules before Runtime launch approval",
			"send desktop notifications from preview",
			"persist compatibility settings from preview",
			"enable live tray bridge from preview",
			"start compatibility backend from preview",
			"mutate host root during desktop activation bundle preview",
			"expose raw backend command to desktop shell",
		},
		DesktopSafeSummary: "desktop activation bundle preview gives KDE the launcher, file association, window identity, tray, notification, and settings materials needed to present a compatibility application as a normal Linux application without writing host files or exposing backend details.",
	}
	if err := validateNoBackendTerms(preview, "desktop activation bundle preview"); err != nil {
		return DesktopActivationBundlePreview{}, err
	}
	return preview, nil
}

func desktopActivationMaterials() []DesktopActivationMaterial {
	return []DesktopActivationMaterial{
		desktopActivationMaterial("launcher", "Start menu launcher", "Plasma application launcher", "desktop-entry-preview", "Launch"),
		desktopActivationMaterial("file-association", "File association", "Dolphin and MIME applications", "mimeapps-preview", "Launch"),
		desktopActivationMaterial("desktop-icon", "Desktop icon", "Plasma desktop", "desktop-icon-preview", "GetDesktopIconPlan"),
		desktopActivationMaterial("task-manager", "Task manager identity", "Plasma task manager", "window-identity-preview", "GetTaskManagerIdentityPlan"),
		desktopActivationMaterial("kwin-window-rule", "KWin window identity", "KWin", "window-identity-preview", "GetKWinWindowRulePlan"),
		desktopActivationMaterial("system-tray", "System tray status", "Plasma system tray", "tray-status-preview", "GetTrayStatus"),
		desktopActivationMaterial("notification-center", "Notification routing", "Plasma notification center", "notification-preview", "GetNotificationPlan"),
		desktopActivationMaterial("unified-settings", "Unified settings", "KDE system settings", "settings-preview", "GetCompatibilitySettings"),
		desktopActivationMaterial("compatibility-center", "Compatibility Center link", "Plasma widget", "compatibility-center-preview", "GetCompatibilityCenterSummary"),
	}
}

func desktopActivationMaterial(id string, label string, component string, source string, method string) DesktopActivationMaterial {
	return DesktopActivationMaterial{
		ID:                    id,
		Label:                 label,
		KDEComponent:          component,
		RuntimeSource:         source,
		RuntimeMethod:         method,
		State:                 "preview-ready",
		UserVisible:           true,
		RequiredForNormalApp:  true,
		WritesHost:            false,
		StartsBackend:         false,
		BackendDetailsExposed: false,
	}
}

func desktopActivationMaterialIDs(materials []DesktopActivationMaterial) []string {
	ids := make([]string, 0, len(materials))
	for _, material := range materials {
		ids = append(ids, material.ID)
	}
	return ids
}
