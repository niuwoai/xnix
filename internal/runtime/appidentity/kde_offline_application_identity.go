package appidentity

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
)

type KDEOfflineApplicationIdentityPreview struct {
	SchemaVersion                     string                         `json:"schema_version"`
	RequestType                       string                         `json:"request_type"`
	IdentityType                      string                         `json:"identity_type"`
	Source                            string                         `json:"source"`
	RuntimeMethod                     string                         `json:"runtime_method"`
	ReadMethod                        string                         `json:"read_method"`
	Desktop                           string                         `json:"desktop"`
	ApplicationID                     string                         `json:"application_id"`
	DisplayName                       string                         `json:"display_name"`
	Icon                              string                         `json:"icon"`
	DesktopFile                       string                         `json:"desktop_file"`
	LauncherURL                       string                         `json:"launcher_url"`
	StableIdentityDigest              string                         `json:"stable_identity_digest"`
	RegistryName                      string                         `json:"registry_name"`
	RecipeDigestVerified              bool                           `json:"recipe_digest_verified"`
	RecipeSignatureStatus             string                         `json:"recipe_signature_status"`
	SurfaceIDs                        []string                       `json:"surface_ids"`
	SurfaceCount                      int                            `json:"surface_count"`
	DesktopEntry                      KDEOfflineDesktopEntryIdentity `json:"desktop_entry"`
	MIMEAssociations                  KDEOfflineMIMEIdentity         `json:"mime_associations"`
	KRunner                           KDEOfflineKRunnerIdentity      `json:"krunner"`
	TaskManager                       KDEOfflineTaskManagerIdentity  `json:"task_manager"`
	KWin                              KDEOfflineKWinIdentity         `json:"kwin"`
	Tray                              KDEOfflineTrayIdentity         `json:"tray"`
	Notification                      KDEOfflineNotificationIdentity `json:"notification"`
	Settings                          KDEOfflineSettingsIdentity     `json:"settings"`
	CompatibilityCenter               KDEOfflineCenterIdentity       `json:"compatibility_center"`
	CrossSurfaceIdentityConsistent    bool                           `json:"cross_surface_identity_consistent"`
	RuntimeOwned                      bool                           `json:"runtime_owned"`
	GoRuntimeBacked                   bool                           `json:"go_runtime_backed"`
	KDEPolicyOwner                    bool                           `json:"kde_policy_owner"`
	ReviewOnly                        bool                           `json:"review_only"`
	Offline                           bool                           `json:"offline"`
	DesktopFilesWritten               bool                           `json:"desktop_files_written"`
	MIMEDefaultsWritten               bool                           `json:"mime_defaults_written"`
	KRunnerIndexPersisted             bool                           `json:"krunner_index_persisted"`
	TaskManagerEntryActive            bool                           `json:"task_manager_entry_active"`
	KWinRuleApplied                   bool                           `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled             bool                           `json:"live_tray_bridge_enabled"`
	TrayBridgePersisted               bool                           `json:"tray_bridge_persisted"`
	NotificationSent                  bool                           `json:"notification_sent"`
	NotificationDeliveryEnabled       bool                           `json:"notification_delivery_enabled"`
	NotificationActionsEnabled        bool                           `json:"notification_actions_enabled"`
	SettingsPersisted                 bool                           `json:"settings_persisted"`
	SettingsPersistenceEnabled        bool                           `json:"settings_persistence_enabled"`
	CompatibilityCenterPersisted      bool                           `json:"compatibility_center_persisted"`
	CompatibilityCenterActionsEnabled bool                           `json:"compatibility_center_actions_enabled"`
	LaunchEnabled                     bool                           `json:"launch_enabled"`
	ExecutionStarted                  bool                           `json:"execution_started"`
	BackendProcessStarted             bool                           `json:"backend_process_started"`
	NetworkRequired                   bool                           `json:"network_required"`
	HostRootModified                  bool                           `json:"host_root_modified"`
	RawCommandExposed                 bool                           `json:"raw_command_exposed"`
	BackendDetailsExposed             bool                           `json:"backend_details_exposed"`
	DesktopSafeSummary                string                         `json:"desktop_safe_summary"`
}

type KDEOfflineDesktopEntryIdentity struct {
	ContentSHA256       string `json:"content_sha256"`
	ApplicationID       string `json:"application_id"`
	DisplayName         string `json:"display_name"`
	Icon                string `json:"icon"`
	DesktopFile         string `json:"desktop_file"`
	StartupWMClass      string `json:"startup_wm_class"`
	MIMETypeCount       int    `json:"mime_type_count"`
	StandardEntry       bool   `json:"standard_entry"`
	IdentityFieldsMatch bool   `json:"identity_fields_match"`
	WriteEnabled        bool   `json:"write_enabled"`
}

type KDEOfflineMIMEIdentity struct {
	ContentSHA256       string   `json:"content_sha256"`
	MIMETypes           []string `json:"mime_types"`
	DesktopFile         string   `json:"desktop_file"`
	DefaultCount        int      `json:"default_count"`
	AssociationCount    int      `json:"association_count"`
	IdentityFieldsMatch bool     `json:"identity_fields_match"`
	WriteEnabled        bool     `json:"write_enabled"`
}

type KDEOfflineKRunnerIdentity struct {
	RunnerID              string `json:"runner_id"`
	ApplicationID         string `json:"application_id"`
	DisplayName           string `json:"display_name"`
	Icon                  string `json:"icon"`
	DesktopEntryID        string `json:"desktop_entry_id"`
	MatchCount            int    `json:"match_count"`
	RelevancePercent      int    `json:"relevance_percent"`
	RuntimeOwnedLaunch    bool   `json:"runtime_owned_launch"`
	IdentityFieldsMatch   bool   `json:"identity_fields_match"`
	QueryExecutionEnabled bool   `json:"query_execution_enabled"`
}

type KDEOfflineTaskManagerIdentity struct {
	ApplicationID       string `json:"application_id"`
	DisplayName         string `json:"display_name"`
	DesktopFile         string `json:"desktop_file"`
	LauncherURL         string `json:"launcher_url"`
	GroupingKey         string `json:"grouping_key"`
	IdentityFieldsMatch bool   `json:"identity_fields_match"`
	EntryActive         bool   `json:"entry_active"`
}

type KDEOfflineKWinIdentity struct {
	ApplicationID          string `json:"application_id"`
	DisplayName            string `json:"display_name"`
	DesktopFile            string `json:"desktop_file"`
	LauncherURL            string `json:"launcher_url"`
	ResourceName           string `json:"resource_name"`
	TaskManagerGroupingKey string `json:"task_manager_grouping_key"`
	IdentityFieldsMatch    bool   `json:"identity_fields_match"`
	RuleApplied            bool   `json:"rule_applied"`
}

type KDEOfflineTrayIdentity struct {
	ApplicationID       string   `json:"application_id"`
	DisplayName         string   `json:"display_name"`
	Icon                string   `json:"icon"`
	DesktopFile         string   `json:"desktop_file"`
	CompatibilityState  string   `json:"compatibility_state"`
	RegisteredAppCount  int      `json:"registered_app_count"`
	Actions             []string `json:"actions"`
	IdentityFieldsMatch bool     `json:"identity_fields_match"`
	LiveBridgeEnabled   bool     `json:"live_bridge_enabled"`
	BridgePersisted     bool     `json:"bridge_persisted"`
}

type KDEOfflineNotificationIdentity struct {
	ApplicationID           string   `json:"application_id"`
	DisplayName             string   `json:"display_name"`
	DesktopFile             string   `json:"desktop_file"`
	EventType               string   `json:"event_type"`
	NotificationID          string   `json:"notification_id"`
	NotificationIDNamespace string   `json:"notification_id_namespace"`
	Category                string   `json:"category"`
	Actions                 []string `json:"actions"`
	IdentityFieldsMatch     bool     `json:"identity_fields_match"`
	DeliveryEnabled         bool     `json:"delivery_enabled"`
	ActionExecutionEnabled  bool     `json:"action_execution_enabled"`
}

type KDEOfflineSettingsIdentity struct {
	ApplicationID       string `json:"application_id"`
	DisplayName         string `json:"display_name"`
	Icon                string `json:"icon"`
	DesktopFile         string `json:"desktop_file"`
	SettingsState       string `json:"settings_state"`
	SectionCount        int    `json:"section_count"`
	IdentityFieldsMatch bool   `json:"identity_fields_match"`
	Persisted           bool   `json:"persisted"`
	PersistenceEnabled  bool   `json:"persistence_enabled"`
}

type KDEOfflineCenterIdentity struct {
	ApplicationID       string `json:"application_id"`
	DisplayName         string `json:"display_name"`
	Icon                string `json:"icon"`
	DesktopFile         string `json:"desktop_file"`
	PageType            string `json:"page_type"`
	PageTitle           string `json:"page_title"`
	NavigationCount     int    `json:"navigation_count"`
	IdentityFieldsMatch bool   `json:"identity_fields_match"`
	PageCreated         bool   `json:"page_created"`
	Persisted           bool   `json:"persisted"`
	ActionsEnabled      bool   `json:"actions_enabled"`
}

func NewKDEOfflineApplicationIdentityPreview(recipe Recipe, provenance Provenance) (KDEOfflineApplicationIdentityPreview, error) {
	plan, err := NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return KDEOfflineApplicationIdentityPreview{}, err
	}
	if provenance.Source != "registry" || !provenance.DigestVerified {
		return KDEOfflineApplicationIdentityPreview{}, errors.New("offline KDE application identity requires a digest-verified registry recipe")
	}

	desktopEntry, err := plan.RenderDesktopEntry()
	if err != nil {
		return KDEOfflineApplicationIdentityPreview{}, err
	}
	mimeApps, err := plan.RenderMIMEApps()
	if err != nil {
		return KDEOfflineApplicationIdentityPreview{}, err
	}
	krunner, err := NewKRunnerQueryPreview([]Recipe{recipe}, provenance, recipe.Name)
	if err != nil {
		return KDEOfflineApplicationIdentityPreview{}, err
	}
	if len(krunner.Matches) != 1 {
		return KDEOfflineApplicationIdentityPreview{}, errors.New("offline KDE application identity requires exactly one KRunner match")
	}
	taskManager, err := plan.TaskManagerIdentityPlanPreview()
	if err != nil {
		return KDEOfflineApplicationIdentityPreview{}, err
	}
	kwin, err := plan.KWinWindowRulePlanPreview()
	if err != nil {
		return KDEOfflineApplicationIdentityPreview{}, err
	}
	tray, err := plan.TrayStatusPreview()
	if err != nil {
		return KDEOfflineApplicationIdentityPreview{}, err
	}
	notification, err := plan.NotificationPreview("approval-required")
	if err != nil {
		return KDEOfflineApplicationIdentityPreview{}, err
	}
	settings, err := plan.SettingsPreview()
	if err != nil {
		return KDEOfflineApplicationIdentityPreview{}, err
	}
	center, err := NewKDECenterPagePreview(recipe, provenance, "reviewed", nil)
	if err != nil {
		return KDEOfflineApplicationIdentityPreview{}, err
	}

	match := krunner.Matches[0]
	desktopEntryMatches := desktopEntryIdentityMatches(plan, desktopEntry)
	mimeMatches := mimeIdentityMatches(plan, mimeApps)
	krunnerMatches := match.ApplicationID == plan.ApplicationID && match.Name == plan.DisplayName &&
		match.Icon == plan.Icon && match.Action.DesktopEntryID == plan.DesktopFile
	taskManagerMatches := taskManager.ApplicationID == plan.ApplicationID && taskManager.DisplayName == plan.DisplayName &&
		taskManager.DesktopFile == plan.DesktopFile && taskManager.LauncherURL == "applications:"+plan.DesktopFile &&
		taskManager.GroupingKey == plan.ApplicationID
	kwinMatches := kwin.ApplicationID == plan.ApplicationID && kwin.DisplayName == plan.DisplayName &&
		kwin.DesktopFile == plan.DesktopFile && kwin.LauncherURL == "applications:"+plan.DesktopFile &&
		kwin.Match.ResourceName == plan.ApplicationID && kwin.Set.TaskManagerGroupingKey == plan.ApplicationID
	trayMatches := tray.ApplicationID == plan.ApplicationID && tray.DisplayName == plan.DisplayName &&
		tray.Icon == plan.Icon && tray.DesktopFile == plan.DesktopFile &&
		tray.ApplicationEntry.ApplicationID == plan.ApplicationID && tray.ApplicationEntry.DesktopFile == plan.DesktopFile
	notificationNamespace := plan.ApplicationID + "."
	notificationMatches := notification.ApplicationID == plan.ApplicationID && notification.DisplayName == plan.DisplayName &&
		notification.DesktopFile == plan.DesktopFile && strings.HasPrefix(notification.NotificationID, notificationNamespace)
	settingsMatches := settings.ApplicationID == plan.ApplicationID && settings.DisplayName == plan.DisplayName &&
		settings.Icon == plan.Icon && settings.DesktopFile == plan.DesktopFile
	centerMatches := center.ApplicationID == plan.ApplicationID && center.ApplicationName == plan.DisplayName &&
		center.Icon == plan.Icon && center.DesktopFile == plan.DesktopFile && center.Header.Title == plan.DisplayName

	preview := KDEOfflineApplicationIdentityPreview{
		SchemaVersion:         "xnix.runtime.kde_offline_application_identity.v1",
		RequestType:           "kde-offline-application-identity-preview",
		IdentityType:          "digest-verified-offline-kde-application",
		Source:                "registry+desktop-entry+mimeapps+krunner+task-manager+kwin+tray+notification+settings+compatibility-center",
		RuntimeMethod:         "GetKDEOfflineApplicationIdentity",
		ReadMethod:            "GetKDEOfflineApplicationIdentityPreview",
		Desktop:               "KDE Plasma",
		ApplicationID:         plan.ApplicationID,
		DisplayName:           plan.DisplayName,
		Icon:                  plan.Icon,
		DesktopFile:           plan.DesktopFile,
		LauncherURL:           "applications:" + plan.DesktopFile,
		StableIdentityDigest:  plan.StableIdentityDigest,
		RegistryName:          provenance.RegistryName,
		RecipeDigestVerified:  provenance.DigestVerified,
		RecipeSignatureStatus: provenance.SignatureStatus,
		SurfaceIDs:            []string{"desktop-entry", "mime-associations", "krunner", "task-manager", "kwin", "system-tray", "notification-center", "unified-settings", "compatibility-center"},
		SurfaceCount:          9,
		DesktopEntry: KDEOfflineDesktopEntryIdentity{
			ContentSHA256:       contentSHA256(desktopEntry),
			ApplicationID:       plan.ApplicationID,
			DisplayName:         plan.DisplayName,
			Icon:                plan.Icon,
			DesktopFile:         plan.DesktopFile,
			StartupWMClass:      plan.StartupWMClass,
			MIMETypeCount:       len(plan.MIMETypes),
			StandardEntry:       plan.StandardDesktopEntry,
			IdentityFieldsMatch: desktopEntryMatches,
			WriteEnabled:        false,
		},
		MIMEAssociations: KDEOfflineMIMEIdentity{
			ContentSHA256:       contentSHA256(mimeApps),
			MIMETypes:           append([]string(nil), plan.MIMETypes...),
			DesktopFile:         plan.DesktopFile,
			DefaultCount:        len(plan.MIMETypes),
			AssociationCount:    len(plan.MIMETypes),
			IdentityFieldsMatch: mimeMatches,
			WriteEnabled:        false,
		},
		KRunner: KDEOfflineKRunnerIdentity{
			RunnerID:              match.RunnerID,
			ApplicationID:         match.ApplicationID,
			DisplayName:           match.Name,
			Icon:                  match.Icon,
			DesktopEntryID:        match.Action.DesktopEntryID,
			MatchCount:            len(krunner.Matches),
			RelevancePercent:      match.RelevancePercent,
			RuntimeOwnedLaunch:    match.RuntimeOwnedLaunch,
			IdentityFieldsMatch:   krunnerMatches,
			QueryExecutionEnabled: krunner.Summary.QueryExecutionEnabled,
		},
		TaskManager: KDEOfflineTaskManagerIdentity{
			ApplicationID:       taskManager.ApplicationID,
			DisplayName:         taskManager.DisplayName,
			DesktopFile:         taskManager.DesktopFile,
			LauncherURL:         taskManager.LauncherURL,
			GroupingKey:         taskManager.GroupingKey,
			IdentityFieldsMatch: taskManagerMatches,
			EntryActive:         taskManager.TaskManagerEntryActive,
		},
		KWin: KDEOfflineKWinIdentity{
			ApplicationID:          kwin.ApplicationID,
			DisplayName:            kwin.DisplayName,
			DesktopFile:            kwin.DesktopFile,
			LauncherURL:            kwin.LauncherURL,
			ResourceName:           kwin.Match.ResourceName,
			TaskManagerGroupingKey: kwin.Set.TaskManagerGroupingKey,
			IdentityFieldsMatch:    kwinMatches,
			RuleApplied:            kwin.KWinRuleApplied,
		},
		Tray: KDEOfflineTrayIdentity{
			ApplicationID:       tray.ApplicationID,
			DisplayName:         tray.DisplayName,
			Icon:                tray.Icon,
			DesktopFile:         tray.DesktopFile,
			CompatibilityState:  tray.CompatibilityStatus.State,
			RegisteredAppCount:  tray.RuntimeActivity.RegisteredApplicationCount,
			Actions:             append([]string(nil), tray.Actions...),
			IdentityFieldsMatch: trayMatches,
			LiveBridgeEnabled:   tray.LiveBackendBridgeEnabled,
			BridgePersisted:     tray.BridgeConfigurationPersisted,
		},
		Notification: KDEOfflineNotificationIdentity{
			ApplicationID:           notification.ApplicationID,
			DisplayName:             notification.DisplayName,
			DesktopFile:             notification.DesktopFile,
			EventType:               notification.EventType,
			NotificationID:          notification.NotificationID,
			NotificationIDNamespace: notificationNamespace,
			Category:                notification.Category,
			Actions:                 append([]string(nil), notification.Actions...),
			IdentityFieldsMatch:     notificationMatches,
			DeliveryEnabled:         false,
			ActionExecutionEnabled:  notification.ActionExecutionEnabled,
		},
		Settings: KDEOfflineSettingsIdentity{
			ApplicationID:       settings.ApplicationID,
			DisplayName:         settings.DisplayName,
			Icon:                settings.Icon,
			DesktopFile:         settings.DesktopFile,
			SettingsState:       settings.SettingsState,
			SectionCount:        settings.SectionCount,
			IdentityFieldsMatch: settingsMatches,
			Persisted:           settings.SettingsPersisted,
			PersistenceEnabled:  settings.SettingsPersistenceEnabled,
		},
		CompatibilityCenter: KDEOfflineCenterIdentity{
			ApplicationID:       center.ApplicationID,
			DisplayName:         center.ApplicationName,
			Icon:                center.Icon,
			DesktopFile:         center.DesktopFile,
			PageType:            center.PageType,
			PageTitle:           center.Header.Title,
			NavigationCount:     center.NavigationCount,
			IdentityFieldsMatch: centerMatches,
			PageCreated:         center.PagePreviewCreated,
			Persisted:           center.PagePersisted,
			ActionsEnabled:      center.CardActionsEnabled,
		},
		CrossSurfaceIdentityConsistent: desktopEntryMatches && mimeMatches && krunnerMatches && taskManagerMatches && kwinMatches && trayMatches && notificationMatches && settingsMatches && centerMatches,
		RuntimeOwned:                   true,
		GoRuntimeBacked:                true,
		KDEPolicyOwner:                 false,
		ReviewOnly:                     true,
		Offline:                        true,
		DesktopSafeSummary:             "A digest-verified registry application has one consistent offline identity across KDE launch, file association, search, task manager, window policy, tray, notifications, settings, and Compatibility Center while persistence, delivery, and execution remain disabled.",
	}
	if !preview.CrossSurfaceIdentityConsistent {
		return KDEOfflineApplicationIdentityPreview{}, errors.New("offline KDE application identity fields disagree across surfaces")
	}
	if err := validateNoBackendTerms(preview, "offline KDE application identity preview"); err != nil {
		return KDEOfflineApplicationIdentityPreview{}, err
	}
	return preview, nil
}

func desktopEntryIdentityMatches(plan Plan, rendered string) bool {
	required := []string{
		"Name=" + plan.DisplayName,
		"Icon=" + plan.Icon,
		"StartupWMClass=" + plan.StartupWMClass,
		"X-Xnix-ApplicationId=" + plan.ApplicationID,
	}
	for _, line := range required {
		if !containsExactLine(rendered, line) {
			return false
		}
	}
	return true
}

func mimeIdentityMatches(plan Plan, rendered string) bool {
	for _, mimeType := range plan.MIMETypes {
		if !containsExactLine(rendered, mimeType+"="+plan.DesktopFile) ||
			!containsExactLine(rendered, mimeType+"="+plan.DesktopFile+";") {
			return false
		}
	}
	return len(plan.MIMETypes) > 0
}

func containsExactLine(rendered string, expected string) bool {
	for _, line := range strings.Split(rendered, "\n") {
		if line == expected {
			return true
		}
	}
	return false
}

func contentSHA256(value string) string {
	sum := sha256.Sum256([]byte(value))
	return hex.EncodeToString(sum[:])
}
