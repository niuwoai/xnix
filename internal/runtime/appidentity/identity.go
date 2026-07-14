package appidentity

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var (
	idPattern        = regexp.MustCompile(`^[a-z][a-z0-9-]*(?:\.[a-z0-9-]+)+$`)
	extensionPattern = regexp.MustCompile(`^\.[A-Za-z0-9]{1,16}$`)
)

type Recipe struct {
	ID                  string   `json:"id"`
	Name                string   `json:"name"`
	Icon                string   `json:"icon"`
	Mode                string   `json:"mode"`
	SupportedExtensions []string `json:"supported_extensions"`
}

type Plan struct {
	SchemaVersion               string            `json:"schema_version"`
	ApplicationID               string            `json:"application_id"`
	DisplayName                 string            `json:"display_name"`
	Icon                        string            `json:"icon"`
	DesktopFile                 string            `json:"desktop_file"`
	StartupWMClass              string            `json:"startup_wm_class"`
	Categories                  []string          `json:"categories"`
	MIMETypes                   []string          `json:"mime_types"`
	LauncherAction              string            `json:"launcher_action"`
	LaunchCommand               []string          `json:"launch_command"`
	UserVisible                 bool              `json:"user_visible"`
	StandardDesktopEntry        bool              `json:"standard_desktop_entry"`
	AcceptsFileURIs             bool              `json:"accepts_file_uris"`
	RuntimeOwned                bool              `json:"runtime_owned"`
	BackendTerminologyHidden    bool              `json:"backend_terminology_hidden"`
	DesktopFileWriteEnabled     bool              `json:"desktop_file_write_enabled"`
	BackendLaunchEnabled        bool              `json:"backend_launch_enabled"`
	BackendDetailsExposed       bool              `json:"backend_details_exposed"`
	RawWindowsExecutableExposed bool              `json:"raw_windows_executable_exposed"`
	CompatibilityStorageExposed bool              `json:"compatibility_storage_exposed"`
	HostRootMutationEnabled     bool              `json:"host_root_mutation_enabled"`
	StableIdentityDigest        string            `json:"stable_identity_digest"`
	RecipeSource                string            `json:"recipe_source"`
	RegistryName                string            `json:"registry_name,omitempty"`
	RecipeDigestVerified        bool              `json:"recipe_digest_verified"`
	RecipeSignatureStatus       string            `json:"recipe_signature_status,omitempty"`
	KDEEntryPoints              []string          `json:"kde_entry_points"`
	UserFacingSettings          map[string]string `json:"user_facing_settings"`
	Summary                     string            `json:"summary"`
}

type WindowIdentityPreview struct {
	SchemaVersion         string            `json:"schema_version"`
	ApplicationID         string            `json:"application_id"`
	DisplayName           string            `json:"display_name"`
	Desktop               string            `json:"desktop"`
	DesktopFile           string            `json:"desktop_file"`
	LauncherURL           string            `json:"launcher_url"`
	WindowKind            string            `json:"window_kind"`
	ClassGroup            string            `json:"class_group"`
	ResourceName          string            `json:"resource_name"`
	TitleHint             string            `json:"title_hint"`
	TaskManager           TaskManagerHints  `json:"task_manager"`
	KWin                  KWinIdentityHints `json:"kwin"`
	Restore               RestoreHints      `json:"restore"`
	RuntimeOwned          bool              `json:"runtime_owned"`
	KDEPolicyOwner        bool              `json:"kde_policy_owner"`
	HostRootModified      bool              `json:"host_root_modified"`
	BackendDetailsExposed bool              `json:"backend_details_exposed"`
	UserFacingSettings    map[string]string `json:"user_facing_settings"`
	Summary               string            `json:"summary"`
}

type TaskManagerHints struct {
	GroupingKey          string `json:"grouping_key"`
	PinningAllowed       bool   `json:"pinning_allowed"`
	RestoreAllowed       bool   `json:"restore_allowed"`
	SkipTaskbar          bool   `json:"skip_taskbar"`
	ShowInSwitcher       bool   `json:"show_in_switcher"`
	PreferExistingWindow bool   `json:"prefer_existing_window"`
}

type KWinIdentityHints struct {
	ScriptRole               string `json:"script_role"`
	ResourceName             string `json:"resource_name"`
	ClassGroup               string `json:"class_group"`
	DesktopFile              string `json:"desktop_file"`
	TaskManagerGroupingKey   string `json:"task_manager_grouping_key"`
	LauncherURL              string `json:"launcher_url"`
	Placement                string `json:"placement"`
	WindowManagerPolicyOnly  bool   `json:"window_manager_policy_only"`
	RuntimeOwnsBackendPolicy bool   `json:"runtime_owns_backend_policy"`
}

type RestoreHints struct {
	RestoreKey           string `json:"restore_key"`
	PinningAllowed       bool   `json:"pinning_allowed"`
	RestoreAllowed       bool   `json:"restore_allowed"`
	PreferExistingWindow bool   `json:"prefer_existing_window"`
}

type TrayStatusPreview struct {
	SchemaVersion                string               `json:"schema_version"`
	StatusType                   string               `json:"status_type"`
	ApplicationID                string               `json:"application_id"`
	DisplayName                  string               `json:"display_name"`
	Desktop                      string               `json:"desktop"`
	Icon                         string               `json:"icon"`
	DesktopFile                  string               `json:"desktop_file"`
	RuntimeActivity              TrayRuntimeActivity  `json:"runtime_activity"`
	CompatibilityStatus          TrayCompatibility    `json:"compatibility_status"`
	TrayBridge                   TrayBridgeStatus     `json:"tray_bridge"`
	ApplicationEntry             TrayApplicationEntry `json:"application_entry"`
	Actions                      []string             `json:"actions"`
	RuntimeOwned                 bool                 `json:"runtime_owned"`
	KDEPolicyOwner               bool                 `json:"kde_policy_owner"`
	UserVisible                  bool                 `json:"user_visible"`
	LiveBackendBridgeEnabled     bool                 `json:"live_backend_bridge_enabled"`
	BridgeConfigurationPersisted bool                 `json:"bridge_configuration_persisted"`
	HostRootModified             bool                 `json:"host_root_modified"`
	BackendDetailsExposed        bool                 `json:"backend_details_exposed"`
	Summary                      string               `json:"summary"`
}

type TrayRuntimeActivity struct {
	ActiveApplicationCount     int    `json:"active_application_count"`
	AttentionRequiredCount     int    `json:"attention_required_count"`
	RegisteredApplicationCount int    `json:"registered_application_count"`
	Summary                    string `json:"summary"`
}

type TrayCompatibility struct {
	State string `json:"state"`
	Label string `json:"label"`
}

type TrayBridgeStatus struct {
	BridgedTrayApplicationCount int    `json:"bridged_tray_application_count"`
	State                       string `json:"state"`
	Label                       string `json:"label"`
}

type TrayApplicationEntry struct {
	ApplicationID      string `json:"application_id"`
	DisplayName        string `json:"display_name"`
	Icon               string `json:"icon"`
	DesktopFile        string `json:"desktop_file"`
	CompatibilityState string `json:"compatibility_state"`
	AttentionRequired  bool   `json:"attention_required"`
	UserVisible        bool   `json:"user_visible"`
}

type NotificationPreview struct {
	SchemaVersion              string   `json:"schema_version"`
	RequestType                string   `json:"request_type"`
	Source                     string   `json:"source"`
	Desktop                    string   `json:"desktop"`
	ApplicationID              string   `json:"application_id"`
	DisplayName                string   `json:"display_name"`
	DesktopFile                string   `json:"desktop_file"`
	EventType                  string   `json:"event_type"`
	NotificationID             string   `json:"notification_id"`
	Urgency                    string   `json:"urgency"`
	Category                   string   `json:"category"`
	Title                      string   `json:"title"`
	Body                       string   `json:"body"`
	Actions                    []string `json:"actions"`
	RequiresUserReview         bool     `json:"requires_user_review"`
	RuntimeOwned               bool     `json:"runtime_owned"`
	KDEPolicyOwner             bool     `json:"kde_policy_owner"`
	UserVisible                bool     `json:"user_visible"`
	ActionExecutionEnabled     bool     `json:"action_execution_enabled"`
	RepairExecutionEnabled     bool     `json:"repair_execution_enabled"`
	SettingsPersistenceEnabled bool     `json:"settings_persistence_enabled"`
	HostRootModified           bool     `json:"host_root_modified"`
	BackendDetailsExposed      bool     `json:"backend_details_exposed"`
	Summary                    string   `json:"summary"`
}

type SettingsPreview struct {
	SchemaVersion              string            `json:"schema_version"`
	RequestType                string            `json:"request_type"`
	Desktop                    string            `json:"desktop"`
	ApplicationID              string            `json:"application_id"`
	DisplayName                string            `json:"display_name"`
	Icon                       string            `json:"icon"`
	DesktopFile                string            `json:"desktop_file"`
	RuntimeOwned               bool              `json:"runtime_owned"`
	KDEPolicyOwner             bool              `json:"kde_policy_owner"`
	UserVisible                bool              `json:"user_visible"`
	SettingsState              string            `json:"settings_state"`
	SettingsPersisted          bool              `json:"settings_persisted"`
	SettingsPersistenceEnabled bool              `json:"settings_persistence_enabled"`
	HostRootModified           bool              `json:"host_root_modified"`
	BackendDetailsExposed      bool              `json:"backend_details_exposed"`
	SectionCount               int               `json:"section_count"`
	Sections                   []SettingsSection `json:"sections"`
	UserFacingSettings         map[string]string `json:"user_facing_settings"`
	Summary                    string            `json:"summary"`
}

type SettingsSection struct {
	ID          string          `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Fields      []SettingsField `json:"fields"`
}

type SettingsField struct {
	ID      string   `json:"id"`
	Label   string   `json:"label"`
	Value   string   `json:"value"`
	Options []string `json:"options"`
}

func NewPlan(recipe Recipe) (Plan, error) {
	return NewPlanWithProvenance(recipe, Provenance{Source: "direct-file"})
}

func NewPlanWithProvenance(recipe Recipe, provenance Provenance) (Plan, error) {
	if err := recipe.Validate(); err != nil {
		return Plan{}, err
	}
	if provenance.Source == "" {
		provenance.Source = "direct-file"
	}

	mimeTypes := recipe.MIMETypes()
	identityDigest := digestIdentity(recipe, mimeTypes)

	return Plan{
		SchemaVersion:            "xnix.runtime.desktop_identity.v1",
		ApplicationID:            recipe.ID,
		DisplayName:              recipe.Name,
		Icon:                     recipe.Icon,
		DesktopFile:              fmt.Sprintf("xnix-%s.desktop", recipe.ID),
		StartupWMClass:           fmt.Sprintf("xnix-%s", recipe.ID),
		Categories:               []string{"Utility"},
		MIMETypes:                mimeTypes,
		LauncherAction:           "runtime-launch",
		LaunchCommand:            []string{"xnix-compat-launch", "--app", recipe.ID, "%U"},
		UserVisible:              true,
		StandardDesktopEntry:     true,
		AcceptsFileURIs:          len(mimeTypes) > 0,
		RuntimeOwned:             true,
		BackendTerminologyHidden: true,
		StableIdentityDigest:     identityDigest,
		RecipeSource:             provenance.Source,
		RegistryName:             provenance.RegistryName,
		RecipeDigestVerified:     provenance.DigestVerified,
		RecipeSignatureStatus:    provenance.SignatureStatus,
		KDEEntryPoints: []string{
			"start-menu",
			"task-manager",
			"file-manager",
			"system-tray",
			"notification-center",
			"compatibility-center",
			"unified-settings",
		},
		UserFacingSettings: map[string]string{
			"run_mode":        "automatic",
			"resource_access": "review-required",
			"snapshot":        "enabled",
		},
		Summary: "desktop identity plan presents a compatibility application as a normal Linux application while keeping backend details hidden behind the Runtime.",
	}, nil
}

func ParseRecipe(data []byte) (Recipe, error) {
	var recipe Recipe
	if err := json.Unmarshal(data, &recipe); err != nil {
		return Recipe{}, fmt.Errorf("parse recipe JSON: %w", err)
	}
	return recipe, nil
}

func (recipe Recipe) Validate() error {
	if !idPattern.MatchString(recipe.ID) {
		return errors.New("recipe id must be a reverse-DNS identifier")
	}
	if !singleLine(recipe.Name) {
		return errors.New("recipe name must be a non-empty single-line string")
	}
	if !singleLine(recipe.Icon) {
		return errors.New("recipe icon must be a non-empty single-line string")
	}
	switch recipe.Mode {
	case "automatic", "wine", "vm":
	default:
		return errors.New("recipe mode must be automatic, wine, or vm")
	}
	for _, extension := range recipe.SupportedExtensions {
		if !extensionPattern.MatchString(extension) {
			return fmt.Errorf("invalid supported extension: %s", extension)
		}
	}
	return nil
}

func (recipe Recipe) MIMETypes() []string {
	mimeTypes := make([]string, 0, len(recipe.SupportedExtensions))
	seen := make(map[string]bool, len(recipe.SupportedExtensions))
	for _, extension := range recipe.SupportedExtensions {
		normalized := strings.ToLower(strings.TrimPrefix(extension, "."))
		mimeType := "application/x-xnix-" + normalized
		if !seen[mimeType] {
			mimeTypes = append(mimeTypes, mimeType)
			seen[mimeType] = true
		}
	}
	sort.Strings(mimeTypes)
	return mimeTypes
}

func (plan Plan) ValidateSafeForDesktop() error {
	encoded, err := json.Marshal(plan)
	if err != nil {
		return fmt.Errorf("encode plan JSON: %w", err)
	}
	text := strings.ToLower(string(encoded))
	forbidden := []string{"prefix", ".exe", "wine ", "wine/", "proton", "qemu-system", "program files"}
	for _, term := range forbidden {
		if strings.Contains(text, term) {
			return fmt.Errorf("desktop identity plan exposes forbidden backend term: %s", term)
		}
	}
	if !plan.RuntimeOwned || !plan.BackendTerminologyHidden {
		return errors.New("desktop identity plan must be Runtime-owned and hide backend terminology")
	}
	if plan.DesktopFileWriteEnabled || plan.BackendLaunchEnabled || plan.BackendDetailsExposed ||
		plan.RawWindowsExecutableExposed || plan.CompatibilityStorageExposed || plan.HostRootMutationEnabled {
		return errors.New("desktop identity plan must keep unsafe gates disabled")
	}
	return nil
}

func (plan Plan) RenderDesktopEntry() (string, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return "", err
	}
	if !plan.StandardDesktopEntry || !plan.UserVisible {
		return "", errors.New("desktop entry requires a standard user-visible plan")
	}
	if len(plan.LaunchCommand) != 4 {
		return "", errors.New("desktop entry requires a complete managed launcher command")
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile, plan.StartupWMClass} {
		if !singleLine(value) {
			return "", errors.New("desktop entry fields must be non-empty single-line strings")
		}
	}

	lines := []string{
		"[Desktop Entry]",
		"Type=Application",
		"Version=1.0",
		"Name=" + plan.DisplayName,
		"Comment=Run with Xnix Compatibility Runtime",
		"Exec=" + strings.Join(plan.LaunchCommand, " "),
		"Icon=" + plan.Icon,
		"Categories=" + strings.Join(plan.Categories, ";") + ";",
		"StartupNotify=true",
		"StartupWMClass=" + plan.StartupWMClass,
		"X-Xnix-ApplicationId=" + plan.ApplicationID,
		"X-Xnix-RuntimeOwned=true",
	}
	if len(plan.MIMETypes) > 0 {
		lines = append(lines, "MimeType="+strings.Join(plan.MIMETypes, ";")+";")
	}
	return strings.Join(lines, "\n") + "\n", nil
}

func (plan Plan) RenderMIMEApps() (string, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return "", err
	}
	if !singleLine(plan.DesktopFile) {
		return "", errors.New("MIME association preview requires a desktop file")
	}
	if len(plan.MIMETypes) == 0 {
		return "", errors.New("MIME association preview requires at least one MIME type")
	}
	for _, mimeType := range plan.MIMETypes {
		if !singleLine(mimeType) {
			return "", errors.New("MIME association preview requires single-line MIME types")
		}
	}

	lines := []string{"[Default Applications]"}
	for _, mimeType := range plan.MIMETypes {
		lines = append(lines, mimeType+"="+plan.DesktopFile)
	}
	lines = append(lines, "", "[Added Associations]")
	for _, mimeType := range plan.MIMETypes {
		lines = append(lines, mimeType+"="+plan.DesktopFile+";")
	}
	return strings.Join(lines, "\n") + "\n", nil
}

func (plan Plan) WindowIdentityPreview() (WindowIdentityPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return WindowIdentityPreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.DesktopFile} {
		if !singleLine(value) {
			return WindowIdentityPreview{}, errors.New("window identity preview requires single-line identity fields")
		}
	}

	launcherURL := "applications:" + plan.DesktopFile
	taskManager := TaskManagerHints{
		GroupingKey:          plan.ApplicationID,
		PinningAllowed:       true,
		RestoreAllowed:       true,
		SkipTaskbar:          false,
		ShowInSwitcher:       true,
		PreferExistingWindow: true,
	}

	return WindowIdentityPreview{
		SchemaVersion: "xnix.runtime.window_identity.v1",
		ApplicationID: plan.ApplicationID,
		DisplayName:   plan.DisplayName,
		Desktop:       "KDE Plasma",
		DesktopFile:   plan.DesktopFile,
		LauncherURL:   launcherURL,
		WindowKind:    "compatibility-application",
		ClassGroup:    "xnix-compatibility",
		ResourceName:  plan.ApplicationID,
		TitleHint:     plan.DisplayName,
		TaskManager:   taskManager,
		KWin: KWinIdentityHints{
			ScriptRole:               "identity-and-layout",
			ResourceName:             plan.ApplicationID,
			ClassGroup:               "xnix-compatibility",
			DesktopFile:              plan.DesktopFile,
			TaskManagerGroupingKey:   plan.ApplicationID,
			LauncherURL:              launcherURL,
			Placement:                "normal-window",
			WindowManagerPolicyOnly:  true,
			RuntimeOwnsBackendPolicy: true,
		},
		Restore: RestoreHints{
			RestoreKey:           plan.ApplicationID,
			PinningAllowed:       taskManager.PinningAllowed,
			RestoreAllowed:       taskManager.RestoreAllowed,
			PreferExistingWindow: taskManager.PreferExistingWindow,
		},
		RuntimeOwned:          true,
		KDEPolicyOwner:        false,
		HostRootModified:      false,
		BackendDetailsExposed: false,
		UserFacingSettings:    plan.UserFacingSettings,
		Summary:               "window identity preview lets KDE group, pin, switch, and restore compatibility windows as normal Linux application windows while backend policy stays in the Runtime.",
	}, nil
}

func (plan Plan) TrayStatusPreview() (TrayStatusPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return TrayStatusPreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return TrayStatusPreview{}, errors.New("tray status preview requires single-line identity fields")
		}
	}

	return TrayStatusPreview{
		SchemaVersion: "xnix.runtime.tray_status.v1",
		StatusType:    "tray-status-preview",
		ApplicationID: plan.ApplicationID,
		DisplayName:   plan.DisplayName,
		Desktop:       "KDE Plasma",
		Icon:          plan.Icon,
		DesktopFile:   plan.DesktopFile,
		RuntimeActivity: TrayRuntimeActivity{
			ActiveApplicationCount:     0,
			AttentionRequiredCount:     0,
			RegisteredApplicationCount: 1,
			Summary:                    plan.DisplayName + " is registered for KDE tray visibility",
		},
		CompatibilityStatus: TrayCompatibility{
			State: "ready",
			Label: "Ready",
		},
		TrayBridge: TrayBridgeStatus{
			BridgedTrayApplicationCount: 0,
			State:                       "planned",
			Label:                       "Tray bridge is planned",
		},
		ApplicationEntry: TrayApplicationEntry{
			ApplicationID:      plan.ApplicationID,
			DisplayName:        plan.DisplayName,
			Icon:               plan.Icon,
			DesktopFile:        plan.DesktopFile,
			CompatibilityState: "ready",
			AttentionRequired:  false,
			UserVisible:        true,
		},
		Actions:                      []string{"open-compatibility-center", "open-settings"},
		RuntimeOwned:                 true,
		KDEPolicyOwner:               false,
		UserVisible:                  true,
		LiveBackendBridgeEnabled:     false,
		BridgeConfigurationPersisted: false,
		HostRootModified:             false,
		BackendDetailsExposed:        false,
		Summary:                      "tray status preview lets KDE show compatibility application status while live tray bridging and backend policy stay gated in the Runtime.",
	}, nil
}

func (plan Plan) NotificationPreview(eventType string) (NotificationPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return NotificationPreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.DesktopFile, eventType} {
		if !singleLine(value) {
			return NotificationPreview{}, errors.New("notification preview requires single-line identity fields")
		}
	}

	preview := NotificationPreview{
		SchemaVersion:              "xnix.runtime.notification.v1",
		RequestType:                "desktop-notification-preview",
		Source:                     "runtime-event",
		Desktop:                    "KDE Plasma",
		ApplicationID:              plan.ApplicationID,
		DisplayName:                plan.DisplayName,
		DesktopFile:                plan.DesktopFile,
		EventType:                  eventType,
		NotificationID:             plan.ApplicationID + "." + eventType,
		RuntimeOwned:               true,
		KDEPolicyOwner:             false,
		UserVisible:                true,
		ActionExecutionEnabled:     false,
		RepairExecutionEnabled:     false,
		SettingsPersistenceEnabled: false,
		HostRootModified:           false,
		BackendDetailsExposed:      false,
	}

	switch eventType {
	case "install-failed":
		preview.Urgency = "critical"
		preview.Category = "compatibility.install"
		preview.Title = plan.DisplayName + " installation needs attention"
		preview.Body = "The Runtime could not complete installation planning and needs Compatibility Center review."
		preview.Actions = []string{"open-compatibility-center", "show-diagnostics"}
		preview.RequiresUserReview = true
		preview.Summary = "notification preview surfaces an installation failure for KDE review without starting repair or install execution."
	case "repair-applied":
		preview.Urgency = "normal"
		preview.Category = "compatibility.repair"
		preview.Title = plan.DisplayName + " repair receipt is ready"
		preview.Body = "A Runtime repair receipt is available for review before any further action is enabled."
		preview.Actions = []string{"open-compatibility-center"}
		preview.RequiresUserReview = false
		preview.Summary = "notification preview reports repair receipts while repair execution remains disabled in this milestone."
	case "mode-changed":
		preview.Urgency = "low"
		preview.Category = "compatibility.settings"
		preview.Title = plan.DisplayName + " compatibility mode changed"
		preview.Body = "The Runtime recorded a compatibility mode change plan for KDE settings review."
		preview.Actions = []string{"open-settings"}
		preview.RequiresUserReview = false
		preview.Summary = "notification preview reports mode changes while settings persistence remains disabled."
	case "approval-required":
		preview.Urgency = "critical"
		preview.Category = "compatibility.approval"
		preview.Title = plan.DisplayName + " needs approval"
		preview.Body = "A compatibility action needs explicit review before the Runtime can proceed."
		preview.Actions = []string{"open-compatibility-center", "review-request"}
		preview.RequiresUserReview = true
		preview.Summary = "notification preview asks KDE to show a user-review notification while Runtime execution gates stay closed."
	default:
		return NotificationPreview{}, fmt.Errorf("unsupported notification event type: %s", eventType)
	}

	return preview, nil
}

func (plan Plan) SettingsPreview() (SettingsPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return SettingsPreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return SettingsPreview{}, errors.New("settings preview requires single-line identity fields")
		}
	}

	sections := []SettingsSection{
		settingsSection("run-mode", "Run mode", "Choose how Xnix balances speed and compatibility.", []SettingsField{
			settingsField("mode", "Run mode", "automatic", []string{"automatic", "performance", "compatibility"}),
			settingsField("preference", "Priority", "compatibility", []string{"performance", "compatibility"}),
		}),
		settingsSection("resource-access", "File access", "Control which user folders this application may request.", []SettingsField{
			settingsField("documents", "Documents", "ask", []string{"allow", "ask", "deny"}),
			settingsField("downloads", "Downloads", "ask", []string{"allow", "ask", "deny"}),
		}),
		settingsSection("devices", "Devices", "Control sensitive device access.", []SettingsField{
			settingsField("camera", "Camera", "deny", []string{"allow", "ask", "deny"}),
		}),
		settingsSection("network", "Network", "Control network access for compatibility actions.", []SettingsField{
			settingsField("network", "Network", "allow", []string{"allow", "ask", "deny"}),
		}),
		settingsSection("snapshots", "Snapshots", "Keep restore points before risky compatibility changes.", []SettingsField{
			settingsField("snapshots", "Environment snapshots", "enabled", []string{"enabled", "disabled"}),
		}),
	}

	return SettingsPreview{
		SchemaVersion:              "xnix.runtime.settings.v1",
		RequestType:                "settings-preview",
		Desktop:                    "KDE Plasma",
		ApplicationID:              plan.ApplicationID,
		DisplayName:                plan.DisplayName,
		Icon:                       plan.Icon,
		DesktopFile:                plan.DesktopFile,
		RuntimeOwned:               true,
		KDEPolicyOwner:             false,
		UserVisible:                true,
		SettingsState:              "planned",
		SettingsPersisted:          false,
		SettingsPersistenceEnabled: false,
		HostRootModified:           false,
		BackendDetailsExposed:      false,
		SectionCount:               len(sections),
		Sections:                   sections,
		UserFacingSettings:         plan.UserFacingSettings,
		Summary:                    "settings preview lets KDE show user-facing compatibility controls while persistence and implementation details stay gated in the Runtime.",
	}, nil
}

func settingsSection(id string, title string, description string, fields []SettingsField) SettingsSection {
	return SettingsSection{
		ID:          id,
		Title:       title,
		Description: description,
		Fields:      fields,
	}
}

func settingsField(id string, label string, value string, options []string) SettingsField {
	return SettingsField{
		ID:      id,
		Label:   label,
		Value:   value,
		Options: options,
	}
}

func singleLine(value string) bool {
	return value != "" && !strings.ContainsAny(value, "\r\n")
}

func digestIdentity(recipe Recipe, mimeTypes []string) string {
	parts := []string{
		recipe.ID,
		recipe.Name,
		recipe.Icon,
		strings.Join(mimeTypes, ";"),
	}
	sum := sha256.Sum256([]byte(strings.Join(parts, "\x00")))
	return hex.EncodeToString(sum[:])
}
