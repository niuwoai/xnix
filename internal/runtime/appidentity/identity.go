package appidentity

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"path"
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

type SettingsChangePreview struct {
	SchemaVersion              string                       `json:"schema_version"`
	RequestType                string                       `json:"request_type"`
	PlanType                   string                       `json:"plan_type"`
	Source                     string                       `json:"source"`
	Desktop                    string                       `json:"desktop"`
	RuntimeMethod              string                       `json:"runtime_method"`
	ApplicationID              string                       `json:"application_id"`
	DisplayName                string                       `json:"display_name"`
	Icon                       string                       `json:"icon"`
	DesktopFile                string                       `json:"desktop_file"`
	SectionID                  string                       `json:"section_id"`
	FieldID                    string                       `json:"field_id"`
	RequestedValue             string                       `json:"requested_value"`
	ChangeState                string                       `json:"change_state"`
	ApplyEnabled               bool                         `json:"apply_enabled"`
	SettingsPersisted          bool                         `json:"settings_persisted"`
	SettingsPersistenceEnabled bool                         `json:"settings_persistence_enabled"`
	HostRootModified           bool                         `json:"host_root_modified"`
	BackendDetailsExposed      bool                         `json:"backend_details_exposed"`
	UserConfirmationRequired   bool                         `json:"user_confirmation_required"`
	SnapshotRecommended        bool                         `json:"snapshot_recommended"`
	PortalPolicyReviewRequired bool                         `json:"portal_policy_review_required"`
	RuntimeRestartRequired     bool                         `json:"runtime_restart_required"`
	AffectedPolicy             SettingsChangeAffectedPolicy `json:"affected_policy"`
	Steps                      []SettingsChangeStep         `json:"steps"`
	BlockedActions             []string                     `json:"blocked_actions"`
	RuntimeOwned               bool                         `json:"runtime_owned"`
	KDEPolicyOwner             bool                         `json:"kde_policy_owner"`
	UserVisible                bool                         `json:"user_visible"`
	UserFacingSettings         map[string]string            `json:"user_facing_settings"`
	DesktopSafeSummary         string                       `json:"desktop_safe_summary"`
}

type SettingsChangeAffectedPolicy struct {
	Section string   `json:"section"`
	Field   string   `json:"field"`
	Value   string   `json:"value"`
	Options []string `json:"options"`
}

type SettingsChangeStep struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
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

type PermissionReviewPreview struct {
	SchemaVersion              string                  `json:"schema_version"`
	RequestType                string                  `json:"request_type"`
	PlanType                   string                  `json:"plan_type"`
	Source                     string                  `json:"source"`
	Desktop                    string                  `json:"desktop"`
	RuntimeMethod              string                  `json:"runtime_method"`
	ApplicationID              string                  `json:"application_id"`
	DisplayName                string                  `json:"display_name"`
	Icon                       string                  `json:"icon"`
	DesktopFile                string                  `json:"desktop_file"`
	ReviewState                string                  `json:"review_state"`
	Permissions                []PermissionReviewEntry `json:"permissions"`
	PermissionCount            int                     `json:"permission_count"`
	AllowCount                 int                     `json:"allow_count"`
	AskCount                   int                     `json:"ask_count"`
	DenyCount                  int                     `json:"deny_count"`
	RequiredRuntimeGates       []string                `json:"required_runtime_gates"`
	RuntimeOwned               bool                    `json:"runtime_owned"`
	KDEPolicyOwner             bool                    `json:"kde_policy_owner"`
	UserVisible                bool                    `json:"user_visible"`
	UserReviewRequired         bool                    `json:"user_review_required"`
	PortalReviewRequired       bool                    `json:"portal_review_required"`
	PermissionChangesApplied   bool                    `json:"permission_changes_applied"`
	RequestObjectsCreated      bool                    `json:"request_objects_created"`
	PermissionsGranted         bool                    `json:"permissions_granted"`
	SettingsPersisted          bool                    `json:"settings_persisted"`
	SettingsPersistenceEnabled bool                    `json:"settings_persistence_enabled"`
	HostPermissionChanged      bool                    `json:"host_permission_changed"`
	HostRootModified           bool                    `json:"host_root_modified"`
	BackendDetailsExposed      bool                    `json:"backend_details_exposed"`
	UserFacingSettings         map[string]string       `json:"user_facing_settings"`
	Summary                    PermissionReviewSummary `json:"summary"`
}

type PermissionReviewEntry struct {
	ID                    string `json:"id"`
	Label                 string `json:"label"`
	Operation             string `json:"operation"`
	Decision              string `json:"decision"`
	PortalInterface       string `json:"portal_interface"`
	PortalRequired        bool   `json:"portal_required"`
	UserMediationRequired bool   `json:"user_mediation_required"`
	CurrentValue          string `json:"current_value"`
	RequestedValue        string `json:"requested_value"`
	ChangePending         bool   `json:"change_pending"`
	RequestObjectCreated  bool   `json:"request_object_created"`
	PermissionGranted     bool   `json:"permission_granted"`
	DirectAccessAllowed   bool   `json:"direct_access_allowed"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

type PermissionReviewSummary struct {
	Headline string `json:"headline"`
	Detail   string `json:"detail"`
}

type DesktopResourceBridgePreview struct {
	SchemaVersion           string                          `json:"schema_version"`
	RequestType             string                          `json:"request_type"`
	PlanType                string                          `json:"plan_type"`
	Source                  string                          `json:"source"`
	Desktop                 string                          `json:"desktop"`
	RuntimeMethod           string                          `json:"runtime_method"`
	ApplicationID           string                          `json:"application_id"`
	DisplayName             string                          `json:"display_name"`
	Icon                    string                          `json:"icon"`
	DesktopFile             string                          `json:"desktop_file"`
	BridgeState             string                          `json:"bridge_state"`
	Resources               []DesktopResourceBridgeResource `json:"resources"`
	ResourceCount           int                             `json:"resource_count"`
	RequiredRuntimeGates    []string                        `json:"required_runtime_gates"`
	RuntimeOwned            bool                            `json:"runtime_owned"`
	KDEPolicyOwner          bool                            `json:"kde_policy_owner"`
	UserVisible             bool                            `json:"user_visible"`
	PortalMediated          bool                            `json:"portal_mediated"`
	FileBridgePlanned       bool                            `json:"file_bridge_planned"`
	URIBridgePlanned        bool                            `json:"uri_bridge_planned"`
	PrintBridgePlanned      bool                            `json:"print_bridge_planned"`
	ClipboardBridgePlanned  bool                            `json:"clipboard_bridge_planned"`
	ScreenshotBridgePlanned bool                            `json:"screenshot_bridge_planned"`
	BridgesEnabled          bool                            `json:"bridges_enabled"`
	RequestsCreated         bool                            `json:"requests_created"`
	BackendProcessStarted   bool                            `json:"backend_process_started"`
	DirectHostFileAccess    bool                            `json:"direct_host_file_access"`
	DirectClipboardAccess   bool                            `json:"direct_clipboard_access"`
	DirectPrintAccess       bool                            `json:"direct_print_access"`
	HostRootModified        bool                            `json:"host_root_modified"`
	BackendDetailsExposed   bool                            `json:"backend_details_exposed"`
	UserFacingSettings      map[string]string               `json:"user_facing_settings"`
	Summary                 DesktopResourceBridgeSummary    `json:"summary"`
}

type DesktopResourceBridgeResource struct {
	ID                         string `json:"id"`
	Name                       string `json:"name"`
	Operation                  string `json:"operation"`
	PortalInterface            string `json:"portal_interface"`
	RuntimeMethod              string `json:"runtime_method"`
	State                      string `json:"state"`
	PortalRequired             bool   `json:"portal_required"`
	UserApprovalRequired       bool   `json:"user_approval_required"`
	BridgeEnabled              bool   `json:"bridge_enabled"`
	RequestCreated             bool   `json:"request_created"`
	DirectBackendAccessAllowed bool   `json:"direct_backend_access_allowed"`
	BackendDetailsExposed      bool   `json:"backend_details_exposed"`
	Summary                    string `json:"summary"`
}

type DesktopResourceBridgeSummary struct {
	Headline string `json:"headline"`
	Detail   string `json:"detail"`
}

type PortalRequestPreview struct {
	SchemaVersion      string                  `json:"schema_version"`
	RequestType        string                  `json:"request_type"`
	Source             string                  `json:"source"`
	Desktop            string                  `json:"desktop"`
	RuntimeMethod      string                  `json:"runtime_method"`
	ApplicationID      string                  `json:"application_id"`
	DisplayName        string                  `json:"display_name"`
	Icon               string                  `json:"icon"`
	DesktopFile        string                  `json:"desktop_file"`
	Operation          string                  `json:"operation"`
	Reason             string                  `json:"reason"`
	Decision           string                  `json:"decision"`
	RequestAllowed     bool                    `json:"request_allowed"`
	RuntimeOwned       bool                    `json:"runtime_owned"`
	KDEPolicyOwner     bool                    `json:"kde_policy_owner"`
	UserVisible        bool                    `json:"user_visible"`
	Portal             PortalRequestEndpoint   `json:"portal"`
	Request            PortalRequestObject     `json:"request"`
	Completion         PortalRequestCompletion `json:"completion"`
	Denied             *PortalRequestDenied    `json:"denied"`
	Safety             PortalRequestSafety     `json:"safety"`
	UserFacingSettings map[string]string       `json:"user_facing_settings"`
	Summary            PortalRequestSummary    `json:"summary"`
}

type PortalRequestEndpoint struct {
	Destination string `json:"destination"`
	Interface   string `json:"interface"`
	Method      string `json:"method"`
	ObjectPath  string `json:"object_path"`
	DBusAPI     string `json:"dbus_api"`
}

type PortalRequestObject struct {
	ObjectPathRequired      bool     `json:"object_path_required"`
	RequestObjectCreated    bool     `json:"request_object_created"`
	HandleToken             string   `json:"handle_token"`
	UserMediationRequired   bool     `json:"user_mediation_required"`
	Resources               []string `json:"resources"`
	RuntimePolicyOwner      bool     `json:"runtime_policy_owner"`
	DesktopShellPolicyOwner bool     `json:"desktop_shell_policy_owner"`
}

type PortalRequestCompletion struct {
	Signal        string `json:"signal"`
	ResponseField string `json:"response_field"`
	SuccessCode   int    `json:"success_code"`
	CancelledCode int    `json:"cancelled_code"`
	DeniedCode    int    `json:"denied_code"`
	ResultOwner   string `json:"result_owner"`
}

type PortalRequestDenied struct {
	Reason            string `json:"reason"`
	NextAction        string `json:"next_action"`
	NotificationEvent string `json:"notification_event"`
}

type PortalRequestSafety struct {
	DirectAccessAllowed   bool `json:"direct_access_allowed"`
	PortalRequired        bool `json:"portal_required"`
	PermissionGranted     bool `json:"permission_granted"`
	HostPermissionChanged bool `json:"host_permission_changed"`
	HostRootModified      bool `json:"host_root_modified"`
	BackendDetailsExposed bool `json:"backend_details_exposed"`
}

type PortalRequestSummary struct {
	Headline string `json:"headline"`
	Detail   string `json:"detail"`
}

type portalOperationRule struct {
	Operation       string
	PortalInterface string
	PortalMethod    string
	Decision        string
	Resources       []string
	Summary         string
}

type settingsChangeRule struct {
	Section              string
	Field                string
	Options              []string
	UserConfirmation     bool
	SnapshotRecommended  bool
	PortalPolicyRequired bool
}

type KRunnerQueryPreview struct {
	SchemaVersion         string         `json:"schema_version"`
	QueryType             string         `json:"query_type"`
	EntryPoint            string         `json:"entry_point"`
	Desktop               string         `json:"desktop"`
	Query                 string         `json:"query"`
	Source                KRunnerSource  `json:"source"`
	RuntimeOwned          bool           `json:"runtime_owned"`
	KDEPolicyOwner        bool           `json:"kde_policy_owner"`
	Matches               []KRunnerMatch `json:"matches"`
	Summary               KRunnerSummary `json:"summary"`
	HostRootModified      bool           `json:"host_root_modified"`
	BackendDetailsExposed bool           `json:"backend_details_exposed"`
	DesktopSafeSummary    string         `json:"desktop_safe_summary"`
}

type KRunnerSource struct {
	Kind                  string `json:"kind"`
	RegistryName          string `json:"registry_name,omitempty"`
	RecipeDigestVerified  bool   `json:"recipe_digest_verified,omitempty"`
	RecipeSignatureStatus string `json:"recipe_signature_status,omitempty"`
}

type KRunnerMatch struct {
	RunnerID              string        `json:"runner_id"`
	ApplicationID         string        `json:"application_id"`
	Name                  string        `json:"name"`
	Icon                  string        `json:"icon"`
	Relevance             float64       `json:"relevance"`
	RelevancePercent      int           `json:"relevance_percent"`
	Subtitle              string        `json:"subtitle"`
	ModeLabel             string        `json:"mode_label"`
	SupportedExtensions   []string      `json:"supported_extensions"`
	RuntimeOwnedLaunch    bool          `json:"runtime_owned_launch"`
	BackendDetailsExposed bool          `json:"backend_details_exposed"`
	Action                KRunnerAction `json:"action"`
}

type KRunnerAction struct {
	Type           string   `json:"type"`
	DesktopEntryID string   `json:"desktop_entry_id"`
	Argv           []string `json:"argv"`
}

type KRunnerSummary struct {
	MatchCount            int    `json:"match_count"`
	OfficialDesktop       string `json:"official_desktop"`
	RuntimeOwnedLaunch    bool   `json:"runtime_owned_launch"`
	QueryExecutionEnabled bool   `json:"query_execution_enabled"`
	BackendLaunchEnabled  bool   `json:"backend_launch_enabled"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

type CompatibilityCenterPreview struct {
	SchemaVersion              string                         `json:"schema_version"`
	SummaryType                string                         `json:"summary_type"`
	Desktop                    string                         `json:"desktop"`
	Title                      string                         `json:"title"`
	Source                     KRunnerSource                  `json:"source"`
	RuntimeOwned               bool                           `json:"runtime_owned"`
	KDEPolicyOwner             bool                           `json:"kde_policy_owner"`
	ApplicationCount           int                            `json:"application_count"`
	KnownIssueCount            int                            `json:"known_issue_count"`
	RepairRecordCount          int                            `json:"repair_record_count"`
	PendingReviewCount         int                            `json:"pending_review_count"`
	Applications               []CompatibilityCenterApp       `json:"applications"`
	ActionExecutionEnabled     bool                           `json:"action_execution_enabled"`
	RepairExecutionEnabled     bool                           `json:"repair_execution_enabled"`
	BackendLaunchEnabled       bool                           `json:"backend_launch_enabled"`
	SettingsPersistenceEnabled bool                           `json:"settings_persistence_enabled"`
	HostRootModified           bool                           `json:"host_root_modified"`
	BackendDetailsExposed      bool                           `json:"backend_details_exposed"`
	Summary                    CompatibilityCenterSummaryText `json:"summary"`
}

type CompatibilityCenterApp struct {
	ApplicationID              string   `json:"application_id"`
	DisplayName                string   `json:"display_name"`
	Icon                       string   `json:"icon"`
	DesktopFile                string   `json:"desktop_file"`
	CompatibilityState         string   `json:"compatibility_state"`
	CompatibilityLabel         string   `json:"compatibility_label"`
	DiagnosticsState           string   `json:"diagnostics_state"`
	RuntimeMode                string   `json:"runtime_mode"`
	SupportedExtensions        []string `json:"supported_extensions"`
	KnownIssueCount            int      `json:"known_issue_count"`
	RepairRecordState          string   `json:"repair_record_state"`
	RepairRecordCount          int      `json:"repair_record_count"`
	LastRepairEvent            string   `json:"last_repair_event"`
	Actions                    []string `json:"actions"`
	UserVisible                bool     `json:"user_visible"`
	RuntimeOwned               bool     `json:"runtime_owned"`
	KDEPolicyOwner             bool     `json:"kde_policy_owner"`
	ActionExecutionEnabled     bool     `json:"action_execution_enabled"`
	RepairExecutionEnabled     bool     `json:"repair_execution_enabled"`
	BackendLaunchEnabled       bool     `json:"backend_launch_enabled"`
	SettingsPersistenceEnabled bool     `json:"settings_persistence_enabled"`
	HostRootModified           bool     `json:"host_root_modified"`
	BackendDetailsExposed      bool     `json:"backend_details_exposed"`
	Summary                    string   `json:"summary"`
}

type CompatibilityCenterSummaryText struct {
	Headline string `json:"headline"`
	Detail   string `json:"detail"`
}

type FileOpenPreview struct {
	SchemaVersion         string            `json:"schema_version"`
	RequestType           string            `json:"request_type"`
	Source                string            `json:"source"`
	Desktop               string            `json:"desktop"`
	ApplicationID         string            `json:"application_id"`
	DisplayName           string            `json:"display_name"`
	DesktopFile           string            `json:"desktop_file"`
	RuntimeMethod         string            `json:"runtime_method"`
	PortalRequired        bool              `json:"portal_required"`
	PortalInterface       string            `json:"portal_interface"`
	PortalMethod          string            `json:"portal_method"`
	FileCount             int               `json:"file_count"`
	FileURIs              []string          `json:"file_uris"`
	SelectedExtension     string            `json:"selected_extension"`
	SelectionMode         string            `json:"selection_mode"`
	Action                FileOpenAction    `json:"action"`
	RuntimeOwned          bool              `json:"runtime_owned"`
	KDEPolicyOwner        bool              `json:"kde_policy_owner"`
	UserVisible           bool              `json:"user_visible"`
	RequestObjectCreated  bool              `json:"request_object_created"`
	PermissionGranted     bool              `json:"permission_granted"`
	BackendLaunchEnabled  bool              `json:"backend_launch_enabled"`
	DirectHostFileAccess  bool              `json:"direct_host_file_access"`
	HostRootModified      bool              `json:"host_root_modified"`
	BackendDetailsExposed bool              `json:"backend_details_exposed"`
	SupportedExtensions   []string          `json:"supported_extensions"`
	UserFacingSettings    map[string]string `json:"user_facing_settings"`
	Summary               FileOpenSummary   `json:"summary"`
}

type FileOpenAction struct {
	Type string   `json:"type"`
	Argv []string `json:"argv"`
}

type FileOpenSummary struct {
	Headline string `json:"headline"`
	Detail   string `json:"detail"`
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

func (plan Plan) SettingsChangePreview(sectionID string, fieldID string, requestedValue string) (SettingsChangePreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return SettingsChangePreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile, sectionID, fieldID, requestedValue} {
		if !singleLine(value) {
			return SettingsChangePreview{}, errors.New("settings change preview requires single-line fields")
		}
	}
	rule, ok := settingsChangeRuleFor(sectionID, fieldID)
	if !ok {
		return SettingsChangePreview{}, fmt.Errorf("unknown settings field: %s/%s", sectionID, fieldID)
	}
	if !containsString(rule.Options, requestedValue) {
		return SettingsChangePreview{}, fmt.Errorf("unsupported settings value: %s", requestedValue)
	}

	preview := SettingsChangePreview{
		SchemaVersion:              "xnix.runtime.settings_change.v1",
		RequestType:                "settings-change-preview",
		PlanType:                   "settings-change-plan",
		Source:                     "unified-settings",
		Desktop:                    "KDE Plasma",
		RuntimeMethod:              "GetCompatibilitySettingsChangePlan",
		ApplicationID:              plan.ApplicationID,
		DisplayName:                plan.DisplayName,
		Icon:                       plan.Icon,
		DesktopFile:                plan.DesktopFile,
		SectionID:                  sectionID,
		FieldID:                    fieldID,
		RequestedValue:             requestedValue,
		ChangeState:                "planned",
		ApplyEnabled:               false,
		SettingsPersisted:          false,
		SettingsPersistenceEnabled: false,
		HostRootModified:           false,
		BackendDetailsExposed:      false,
		UserConfirmationRequired:   rule.UserConfirmation,
		SnapshotRecommended:        rule.SnapshotRecommended,
		PortalPolicyReviewRequired: rule.PortalPolicyRequired,
		RuntimeRestartRequired:     false,
		AffectedPolicy: SettingsChangeAffectedPolicy{
			Section: sectionID,
			Field:   fieldID,
			Value:   requestedValue,
			Options: append([]string(nil), rule.Options...),
		},
		Steps: []SettingsChangeStep{
			settingsChangeStep("validate-setting", "pass", "The requested settings value is valid for the Runtime settings schema."),
			settingsChangeStep("review-user-confirmation", requiredStatus(rule.UserConfirmation), "KDE must present the change for user review before persistence."),
			settingsChangeStep("review-portal-policy", requiredStatus(rule.PortalPolicyRequired), "Runtime Portal policy must be reviewed before desktop resource access changes."),
			settingsChangeStep("prepare-restore-point", recommendedStatus(rule.SnapshotRecommended), "Runtime should prepare a restore point before risky compatibility settings changes."),
			settingsChangeStep("persist-runtime-setting", "pending", "Runtime persistence is not enabled in this version."),
		},
		BlockedActions: []string{
			"persist compatibility settings before Runtime confirmation",
			"grant desktop resources without Portal policy review",
			"modify host root while planning settings changes",
			"expose backend implementation settings to KDE",
		},
		RuntimeOwned:       true,
		KDEPolicyOwner:     false,
		UserVisible:        true,
		UserFacingSettings: plan.UserFacingSettings,
		DesktopSafeSummary: "Compatibility settings change is planned and waiting for Runtime persistence support.",
	}
	if err := validateNoBackendTerms(preview, "settings change preview"); err != nil {
		return SettingsChangePreview{}, err
	}
	return preview, nil
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

func settingsChangeRuleFor(sectionID string, fieldID string) (settingsChangeRule, bool) {
	for _, rule := range []settingsChangeRule{
		{
			Section:             "run-mode",
			Field:               "mode",
			Options:             []string{"automatic", "performance", "compatibility"},
			UserConfirmation:    true,
			SnapshotRecommended: true,
		},
		{
			Section:             "run-mode",
			Field:               "preference",
			Options:             []string{"performance", "compatibility"},
			UserConfirmation:    true,
			SnapshotRecommended: true,
		},
		{
			Section:              "resource-access",
			Field:                "documents",
			Options:              []string{"allow", "ask", "deny"},
			UserConfirmation:     true,
			PortalPolicyRequired: true,
		},
		{
			Section:              "resource-access",
			Field:                "downloads",
			Options:              []string{"allow", "ask", "deny"},
			UserConfirmation:     true,
			PortalPolicyRequired: true,
		},
		{
			Section:              "devices",
			Field:                "camera",
			Options:              []string{"allow", "ask", "deny"},
			UserConfirmation:     true,
			PortalPolicyRequired: true,
		},
		{
			Section:              "network",
			Field:                "network",
			Options:              []string{"allow", "ask", "deny"},
			UserConfirmation:     true,
			PortalPolicyRequired: true,
		},
		{
			Section:             "snapshots",
			Field:               "snapshots",
			Options:             []string{"enabled", "disabled"},
			UserConfirmation:    true,
			SnapshotRecommended: true,
		},
	} {
		if rule.Section == sectionID && rule.Field == fieldID {
			return rule, true
		}
	}
	return settingsChangeRule{}, false
}

func settingsChangeStep(id string, status string, summary string) SettingsChangeStep {
	return SettingsChangeStep{
		ID:      id,
		Status:  status,
		Summary: summary,
	}
}

func requiredStatus(required bool) string {
	if required {
		return "required"
	}
	return "pass"
}

func recommendedStatus(recommended bool) string {
	if recommended {
		return "recommended"
	}
	return "pass"
}

func (plan Plan) PermissionReviewPreview() (PermissionReviewPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return PermissionReviewPreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return PermissionReviewPreview{}, errors.New("permission review preview requires single-line identity fields")
		}
	}

	permissions := defaultPermissionReviewEntries()
	allowCount := 0
	askCount := 0
	denyCount := 0
	for _, permission := range permissions {
		switch permission.Decision {
		case "allow":
			allowCount++
		case "ask":
			askCount++
		case "deny":
			denyCount++
		}
	}

	preview := PermissionReviewPreview{
		SchemaVersion:              "xnix.runtime.permission_review.v1",
		RequestType:                "permission-review-preview",
		PlanType:                   "compatibility-permission-review-plan",
		Source:                     "unified-settings",
		Desktop:                    "KDE Plasma",
		RuntimeMethod:              "GetCompatibilityPermissionReviewPlan",
		ApplicationID:              plan.ApplicationID,
		DisplayName:                plan.DisplayName,
		Icon:                       plan.Icon,
		DesktopFile:                plan.DesktopFile,
		ReviewState:                "planned",
		Permissions:                permissions,
		PermissionCount:            len(permissions),
		AllowCount:                 allowCount,
		AskCount:                   askCount,
		DenyCount:                  denyCount,
		RequiredRuntimeGates:       []string{"user-review", "portal-policy-review", "runtime-write-gate", "settings-persistence", "audit-log"},
		RuntimeOwned:               true,
		KDEPolicyOwner:             false,
		UserVisible:                true,
		UserReviewRequired:         true,
		PortalReviewRequired:       true,
		PermissionChangesApplied:   false,
		RequestObjectsCreated:      false,
		PermissionsGranted:         false,
		SettingsPersisted:          false,
		SettingsPersistenceEnabled: false,
		HostPermissionChanged:      false,
		HostRootModified:           false,
		BackendDetailsExposed:      false,
		UserFacingSettings:         plan.UserFacingSettings,
		Summary: PermissionReviewSummary{
			Headline: "KDE can review compatibility permissions before access changes.",
			Detail:   "The Runtime groups file, device, network, clipboard, print, and screenshot permissions while keeping changes and grants disabled.",
		},
	}
	if err := validateNoBackendTerms(preview, "permission review preview"); err != nil {
		return PermissionReviewPreview{}, err
	}
	return preview, nil
}

func defaultPermissionReviewEntries() []PermissionReviewEntry {
	return []PermissionReviewEntry{
		permissionReviewEntry("documents", "Documents", "file-open", "ask", "org.freedesktop.portal.FileChooser", true),
		permissionReviewEntry("downloads", "Downloads", "file-open", "ask", "org.freedesktop.portal.FileChooser", true),
		permissionReviewEntry("camera", "Camera", "camera", "deny", "org.freedesktop.portal.Camera", true),
		permissionReviewEntry("network", "Network", "network", "allow", "none", false),
		permissionReviewEntry("clipboard", "Clipboard", "clipboard", "ask", "org.freedesktop.portal.Clipboard", true),
		permissionReviewEntry("print", "Print", "print", "ask", "org.freedesktop.portal.Print", true),
		permissionReviewEntry("screenshot", "Screenshot", "screenshot", "ask", "org.freedesktop.portal.Screenshot", true),
	}
}

func permissionReviewEntry(id string, label string, operation string, decision string, portalInterface string, portalRequired bool) PermissionReviewEntry {
	return PermissionReviewEntry{
		ID:                    id,
		Label:                 label,
		Operation:             operation,
		Decision:              decision,
		PortalInterface:       portalInterface,
		PortalRequired:        portalRequired,
		UserMediationRequired: true,
		CurrentValue:          decision,
		RequestedValue:        decision,
		ChangePending:         false,
		RequestObjectCreated:  false,
		PermissionGranted:     false,
		DirectAccessAllowed:   false,
		BackendDetailsExposed: false,
	}
}

func (plan Plan) DesktopResourceBridgePreview() (DesktopResourceBridgePreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return DesktopResourceBridgePreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return DesktopResourceBridgePreview{}, errors.New("desktop resource bridge preview requires single-line identity fields")
		}
	}

	resources := defaultDesktopResourceBridgeResources()
	preview := DesktopResourceBridgePreview{
		SchemaVersion:           "xnix.runtime.desktop_resource_bridge.v1",
		RequestType:             "desktop-resource-bridge-preview",
		PlanType:                "desktop-resource-bridge-plan",
		Source:                  "runtime-resource-boundary",
		Desktop:                 "KDE Plasma",
		RuntimeMethod:           "GetDesktopResourceBridgePlan",
		ApplicationID:           plan.ApplicationID,
		DisplayName:             plan.DisplayName,
		Icon:                    plan.Icon,
		DesktopFile:             plan.DesktopFile,
		BridgeState:             "planned",
		Resources:               resources,
		ResourceCount:           len(resources),
		RequiredRuntimeGates:    []string{"portal-policy-review", "portal-request-plan", "snapshot-baseline", "backend-environment-plan", "runtime-write-gate"},
		RuntimeOwned:            true,
		KDEPolicyOwner:          false,
		UserVisible:             true,
		PortalMediated:          true,
		FileBridgePlanned:       true,
		URIBridgePlanned:        true,
		PrintBridgePlanned:      true,
		ClipboardBridgePlanned:  true,
		ScreenshotBridgePlanned: true,
		BridgesEnabled:          false,
		RequestsCreated:         false,
		BackendProcessStarted:   false,
		DirectHostFileAccess:    false,
		DirectClipboardAccess:   false,
		DirectPrintAccess:       false,
		HostRootModified:        false,
		BackendDetailsExposed:   false,
		UserFacingSettings:      plan.UserFacingSettings,
		Summary: DesktopResourceBridgeSummary{
			Headline: "KDE can show desktop resource bridges before any access is enabled.",
			Detail:   "The Runtime plans file, URI, print, clipboard, and screenshot bridges through XDG Desktop Portal review while keeping bridge activation disabled.",
		},
	}
	if err := validateNoBackendTerms(preview, "desktop resource bridge preview"); err != nil {
		return DesktopResourceBridgePreview{}, err
	}
	return preview, nil
}

func defaultDesktopResourceBridgeResources() []DesktopResourceBridgeResource {
	return []DesktopResourceBridgeResource{
		desktopResourceBridgeResource(
			"file-open",
			"File Open",
			"file-open",
			"org.freedesktop.portal.FileChooser",
			"File opens require XDG Desktop Portal review before selected files can be handed to the Runtime.",
		),
		desktopResourceBridgeResource(
			"uri-open",
			"URI Open",
			"uri-open",
			"org.freedesktop.portal.OpenURI",
			"URI opens are routed through a Runtime-owned Portal request plan.",
		),
		desktopResourceBridgeResource(
			"print",
			"Print",
			"print",
			"org.freedesktop.portal.Print",
			"Printing stays behind user-approved desktop Portal review.",
		),
		desktopResourceBridgeResource(
			"clipboard",
			"Clipboard",
			"clipboard",
			"org.freedesktop.portal.Clipboard",
			"Clipboard bridging stays disabled until Runtime records a reviewed Portal path.",
		),
		desktopResourceBridgeResource(
			"screenshot",
			"Screenshot",
			"screenshot",
			"org.freedesktop.portal.Screenshot",
			"Screenshot access stays Portal-mediated and disabled until user review exists.",
		),
	}
}

func desktopResourceBridgeResource(id string, name string, operation string, portalInterface string, summary string) DesktopResourceBridgeResource {
	return DesktopResourceBridgeResource{
		ID:                         id,
		Name:                       name,
		Operation:                  operation,
		PortalInterface:            portalInterface,
		RuntimeMethod:              "GetPortalRequestPlan",
		State:                      "planned",
		PortalRequired:             true,
		UserApprovalRequired:       true,
		BridgeEnabled:              false,
		RequestCreated:             false,
		DirectBackendAccessAllowed: false,
		BackendDetailsExposed:      false,
		Summary:                    summary,
	}
}

func (plan Plan) PortalRequestPreview(operation string, reason string) (PortalRequestPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return PortalRequestPreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile, operation, reason} {
		if !singleLineOrBlank(value) {
			return PortalRequestPreview{}, errors.New("portal request preview requires single-line fields")
		}
	}
	rule, ok := portalOperationRuleFor(operation)
	if !ok {
		return PortalRequestPreview{}, fmt.Errorf("unsupported Portal operation: %s", operation)
	}
	if strings.TrimSpace(reason) == "" {
		reason = "Compatibility application requested " + operation + " access."
	}

	requestAllowed := rule.Decision != "deny"
	var denied *PortalRequestDenied
	if !requestAllowed {
		denied = &PortalRequestDenied{
			Reason:            rule.Summary,
			NextAction:        "open-compatibility-settings",
			NotificationEvent: "approval-required",
		}
	}

	preview := PortalRequestPreview{
		SchemaVersion:  "xnix.runtime.portal_request.v1",
		RequestType:    "portal-request-preview",
		Source:         "runtime-portal-request-plan",
		Desktop:        "KDE Plasma",
		RuntimeMethod:  "GetPortalRequestPlan",
		ApplicationID:  plan.ApplicationID,
		DisplayName:    plan.DisplayName,
		Icon:           plan.Icon,
		DesktopFile:    plan.DesktopFile,
		Operation:      rule.Operation,
		Reason:         reason,
		Decision:       rule.Decision,
		RequestAllowed: requestAllowed,
		RuntimeOwned:   true,
		KDEPolicyOwner: false,
		UserVisible:    true,
		Portal: PortalRequestEndpoint{
			Destination: "org.freedesktop.portal.Desktop",
			Interface:   rule.PortalInterface,
			Method:      rule.PortalMethod,
			ObjectPath:  "/org/freedesktop/portal/desktop",
			DBusAPI:     "XDG Desktop Portal",
		},
		Request: PortalRequestObject{
			ObjectPathRequired:      true,
			RequestObjectCreated:    false,
			HandleToken:             portalRequestHandleToken(plan.ApplicationID, rule.Operation),
			UserMediationRequired:   true,
			Resources:               append([]string(nil), rule.Resources...),
			RuntimePolicyOwner:      true,
			DesktopShellPolicyOwner: false,
		},
		Completion: PortalRequestCompletion{
			Signal:        "Response",
			ResponseField: "response",
			SuccessCode:   0,
			CancelledCode: 1,
			DeniedCode:    2,
			ResultOwner:   "Runtime",
		},
		Denied: denied,
		Safety: PortalRequestSafety{
			DirectAccessAllowed:   false,
			PortalRequired:        true,
			PermissionGranted:     false,
			HostPermissionChanged: false,
			HostRootModified:      false,
			BackendDetailsExposed: false,
		},
		UserFacingSettings: plan.UserFacingSettings,
		Summary: PortalRequestSummary{
			Headline: "KDE can show the Portal request before any desktop resource is granted.",
			Detail:   "The Runtime owns the Portal request plan, waits for the Portal response, and keeps direct access disabled.",
		},
	}
	if err := validateNoBackendTerms(preview, "Portal request preview"); err != nil {
		return PortalRequestPreview{}, err
	}
	return preview, nil
}

func portalOperationRuleFor(operation string) (portalOperationRule, bool) {
	for _, rule := range []portalOperationRule{
		{
			Operation:       "file-open",
			PortalInterface: "org.freedesktop.portal.FileChooser",
			PortalMethod:    "OpenFile",
			Decision:        "ask",
			Resources:       []string{"documents", "downloads", "selected-files"},
			Summary:         "File access requires a user-approved desktop Portal request.",
		},
		{
			Operation:       "uri-open",
			PortalInterface: "org.freedesktop.portal.OpenURI",
			PortalMethod:    "OpenURI",
			Decision:        "ask",
			Resources:       []string{"external-uri"},
			Summary:         "URI handling requires a user-approved desktop Portal request.",
		},
		{
			Operation:       "print",
			PortalInterface: "org.freedesktop.portal.Print",
			PortalMethod:    "Print",
			Decision:        "ask",
			Resources:       []string{"printer"},
			Summary:         "Printing requires a user-approved desktop Portal request.",
		},
		{
			Operation:       "screenshot",
			PortalInterface: "org.freedesktop.portal.Screenshot",
			PortalMethod:    "Screenshot",
			Decision:        "ask",
			Resources:       []string{"screen"},
			Summary:         "Screenshots require a user-approved desktop Portal request.",
		},
		{
			Operation:       "clipboard",
			PortalInterface: "org.freedesktop.portal.Clipboard",
			PortalMethod:    "RequestClipboard",
			Decision:        "ask",
			Resources:       []string{"clipboard"},
			Summary:         "Clipboard access requires a user-approved desktop Portal request.",
		},
		{
			Operation:       "camera",
			PortalInterface: "org.freedesktop.portal.Camera",
			PortalMethod:    "AccessCamera",
			Decision:        "deny",
			Resources:       []string{"camera"},
			Summary:         "Camera access is denied until the user changes the application policy.",
		},
		{
			Operation:       "remote-desktop",
			PortalInterface: "org.freedesktop.portal.RemoteDesktop",
			PortalMethod:    "CreateSession",
			Decision:        "deny",
			Resources:       []string{"screen", "input-devices"},
			Summary:         "Remote desktop access is denied until the user changes the application policy.",
		},
	} {
		if rule.Operation == operation {
			return rule, true
		}
	}
	return portalOperationRule{}, false
}

func portalRequestHandleToken(applicationID string, operation string) string {
	replacer := strings.NewReplacer(".", "_", "-", "_")
	return "xnix_" + replacer.Replace(applicationID) + "_" + replacer.Replace(operation)
}

func NewKRunnerQueryPreview(recipes []Recipe, provenance Provenance, query string) (KRunnerQueryPreview, error) {
	if !singleLineOrBlank(query) {
		return KRunnerQueryPreview{}, errors.New("KRunner query preview requires a single-line query")
	}
	if provenance.Source == "" {
		provenance.Source = "registry"
	}

	normalizedQuery := normalizeQuery(query)
	matches := make([]KRunnerMatch, 0, len(recipes))
	for _, recipe := range recipes {
		plan, err := NewPlanWithProvenance(recipe, provenance)
		if err != nil {
			return KRunnerQueryPreview{}, err
		}
		if err := plan.ValidateSafeForDesktop(); err != nil {
			return KRunnerQueryPreview{}, err
		}
		relevancePercent := krunnerRelevancePercent(recipe, normalizedQuery)
		if relevancePercent == 0 {
			continue
		}
		matches = append(matches, krunnerMatch(plan, recipe, relevancePercent))
	}
	sort.Slice(matches, func(left int, right int) bool {
		if matches[left].RelevancePercent != matches[right].RelevancePercent {
			return matches[left].RelevancePercent > matches[right].RelevancePercent
		}
		return matches[left].Name < matches[right].Name
	})

	preview := KRunnerQueryPreview{
		SchemaVersion: "xnix.runtime.krunner_query.v1",
		QueryType:     "krunner-query-plan",
		EntryPoint:    "krunner",
		Desktop:       "KDE Plasma",
		Query:         query,
		Source: KRunnerSource{
			Kind:                  "runtime-go-registry",
			RegistryName:          provenance.RegistryName,
			RecipeDigestVerified:  provenance.DigestVerified,
			RecipeSignatureStatus: provenance.SignatureStatus,
		},
		RuntimeOwned:     true,
		KDEPolicyOwner:   false,
		Matches:          matches,
		HostRootModified: false,
		Summary: KRunnerSummary{
			MatchCount:            len(matches),
			OfficialDesktop:       "KDE Plasma",
			RuntimeOwnedLaunch:    true,
			QueryExecutionEnabled: false,
			BackendLaunchEnabled:  false,
			BackendDetailsExposed: false,
		},
		BackendDetailsExposed: false,
		DesktopSafeSummary:    "KRunner query planning is Runtime-owned and returns safe launcher actions only.",
	}
	if err := validateNoBackendTerms(preview, "KRunner query preview"); err != nil {
		return KRunnerQueryPreview{}, err
	}
	return preview, nil
}

func NewCompatibilityCenterPreview(recipes []Recipe, provenance Provenance) (CompatibilityCenterPreview, error) {
	if provenance.Source == "" {
		provenance.Source = "registry"
	}

	applications := make([]CompatibilityCenterApp, 0, len(recipes))
	for _, recipe := range recipes {
		plan, err := NewPlanWithProvenance(recipe, provenance)
		if err != nil {
			return CompatibilityCenterPreview{}, err
		}
		if err := plan.ValidateSafeForDesktop(); err != nil {
			return CompatibilityCenterPreview{}, err
		}
		applications = append(applications, compatibilityCenterApp(plan, recipe))
	}
	sort.Slice(applications, func(left int, right int) bool {
		return applications[left].DisplayName < applications[right].DisplayName
	})

	preview := CompatibilityCenterPreview{
		SchemaVersion:    "xnix.runtime.compatibility_center.v1",
		SummaryType:      "compatibility-center-preview",
		Desktop:          "KDE Plasma",
		Title:            "Xnix Compatibility Center",
		RuntimeOwned:     true,
		KDEPolicyOwner:   false,
		ApplicationCount: len(applications),
		Applications:     applications,
		Source: KRunnerSource{
			Kind:                  "runtime-go-registry",
			RegistryName:          provenance.RegistryName,
			RecipeDigestVerified:  provenance.DigestVerified,
			RecipeSignatureStatus: provenance.SignatureStatus,
		},
		ActionExecutionEnabled:     false,
		RepairExecutionEnabled:     false,
		BackendLaunchEnabled:       false,
		SettingsPersistenceEnabled: false,
		HostRootModified:           false,
		BackendDetailsExposed:      false,
		Summary: CompatibilityCenterSummaryText{
			Headline: "Compatibility applications are ready for KDE review.",
			Detail:   "The Runtime exposes application status, user-facing mode, diagnostics state, and repair records without enabling execution.",
		},
	}
	if err := validateNoBackendTerms(preview, "Compatibility Center preview"); err != nil {
		return CompatibilityCenterPreview{}, err
	}
	return preview, nil
}

func compatibilityCenterApp(plan Plan, recipe Recipe) CompatibilityCenterApp {
	return CompatibilityCenterApp{
		ApplicationID:              plan.ApplicationID,
		DisplayName:                plan.DisplayName,
		Icon:                       plan.Icon,
		DesktopFile:                plan.DesktopFile,
		CompatibilityState:         "registered",
		CompatibilityLabel:         "Registered",
		DiagnosticsState:           "not-run",
		RuntimeMode:                modeLabel(recipe.Mode),
		SupportedExtensions:        normalizedExtensions(recipe.SupportedExtensions),
		KnownIssueCount:            0,
		RepairRecordState:          "none",
		RepairRecordCount:          0,
		LastRepairEvent:            "none",
		Actions:                    []string{"open-settings", "show-diagnostics", "review-application"},
		UserVisible:                true,
		RuntimeOwned:               true,
		KDEPolicyOwner:             false,
		ActionExecutionEnabled:     false,
		RepairExecutionEnabled:     false,
		BackendLaunchEnabled:       false,
		SettingsPersistenceEnabled: false,
		HostRootModified:           false,
		BackendDetailsExposed:      false,
		Summary:                    "KDE can display this compatibility application, but execution and repair actions remain gated in the Runtime.",
	}
}

func NewFileOpenPreview(recipes []Recipe, provenance Provenance, fileURIs []string, applicationID string) (FileOpenPreview, error) {
	if len(fileURIs) == 0 {
		return FileOpenPreview{}, errors.New("file-open preview requires at least one file URI")
	}
	if applicationID != "" && !idPattern.MatchString(applicationID) {
		return FileOpenPreview{}, errors.New("application id must be a reverse-DNS identifier")
	}

	normalizedURIs, selectedExtension, err := normalizeFileURIs(fileURIs)
	if err != nil {
		return FileOpenPreview{}, err
	}
	recipe, selectionMode, err := selectFileOpenRecipe(recipes, normalizedURIs, selectedExtension, applicationID)
	if err != nil {
		return FileOpenPreview{}, err
	}
	plan, err := NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return FileOpenPreview{}, err
	}
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return FileOpenPreview{}, err
	}

	preview := FileOpenPreview{
		SchemaVersion:         "xnix.runtime.file_open.v1",
		RequestType:           "file-open-preview",
		Source:                "dolphin-service-menu",
		Desktop:               "KDE Plasma",
		ApplicationID:         plan.ApplicationID,
		DisplayName:           plan.DisplayName,
		DesktopFile:           plan.DesktopFile,
		RuntimeMethod:         "Launch",
		PortalRequired:        true,
		PortalInterface:       "org.freedesktop.portal.FileChooser",
		PortalMethod:          "OpenFile",
		FileCount:             len(normalizedURIs),
		FileURIs:              normalizedURIs,
		SelectedExtension:     selectedExtension,
		SelectionMode:         selectionMode,
		RuntimeOwned:          true,
		KDEPolicyOwner:        false,
		UserVisible:           true,
		RequestObjectCreated:  false,
		PermissionGranted:     false,
		BackendLaunchEnabled:  false,
		DirectHostFileAccess:  false,
		HostRootModified:      false,
		BackendDetailsExposed: false,
		SupportedExtensions:   normalizedExtensions(recipe.SupportedExtensions),
		UserFacingSettings:    plan.UserFacingSettings,
		Action: FileOpenAction{
			Type: "runtime-file-open",
			Argv: []string{"xnix-compat-open", "--app", plan.ApplicationID, "%U"},
		},
		Summary: FileOpenSummary{
			Headline: "Dolphin can hand selected files to the Runtime for review.",
			Detail:   "The request remains Portal-mediated and does not start a backend or grant file access in this preview.",
		},
	}
	if err := validateNoBackendTerms(preview, "file-open preview"); err != nil {
		return FileOpenPreview{}, err
	}
	return preview, nil
}

func normalizeFileURIs(fileURIs []string) ([]string, string, error) {
	normalized := make([]string, 0, len(fileURIs))
	selectedExtension := ""
	for _, fileURI := range fileURIs {
		parsed, err := url.Parse(fileURI)
		if err != nil {
			return nil, "", fmt.Errorf("invalid file URI: %s", fileURI)
		}
		if parsed.Scheme != "file" {
			return nil, "", errors.New("only file URIs are accepted")
		}
		if parsed.Path == "" || !strings.HasPrefix(parsed.Path, "/") {
			return nil, "", errors.New("file URI must include an absolute path")
		}
		extension := strings.ToLower(path.Ext(parsed.Path))
		if selectedExtension == "" {
			selectedExtension = extension
		}
		normalized = append(normalized, parsed.String())
	}
	if selectedExtension == "" {
		return nil, "", errors.New("selected file must have an extension")
	}
	return normalized, selectedExtension, nil
}

func selectFileOpenRecipe(recipes []Recipe, fileURIs []string, selectedExtension string, applicationID string) (Recipe, string, error) {
	if applicationID != "" {
		for _, recipe := range recipes {
			if recipe.ID == applicationID {
				return recipe, "explicit-application", nil
			}
		}
		return Recipe{}, "", fmt.Errorf("unknown application: %s", applicationID)
	}
	for _, recipe := range recipes {
		if containsString(normalizedExtensions(recipe.SupportedExtensions), selectedExtension) {
			return recipe, "extension-match", nil
		}
	}
	return Recipe{}, "", fmt.Errorf("no compatible application is registered for %s", selectedExtension)
}

func krunnerMatch(plan Plan, recipe Recipe, relevancePercent int) KRunnerMatch {
	return KRunnerMatch{
		RunnerID:              "xnix.compatibility." + plan.ApplicationID,
		ApplicationID:         plan.ApplicationID,
		Name:                  plan.DisplayName,
		Icon:                  plan.Icon,
		Relevance:             float64(relevancePercent) / 100.0,
		RelevancePercent:      relevancePercent,
		Subtitle:              "Open as a normal Linux application",
		ModeLabel:             modeLabel(recipe.Mode),
		SupportedExtensions:   normalizedExtensions(recipe.SupportedExtensions),
		RuntimeOwnedLaunch:    true,
		BackendDetailsExposed: false,
		Action: KRunnerAction{
			Type:           "runtime-launch",
			DesktopEntryID: plan.DesktopFile,
			Argv:           []string{"xnix-compat-launch", "--app", plan.ApplicationID},
		},
	}
}

func krunnerRelevancePercent(recipe Recipe, normalizedQuery string) int {
	if normalizedQuery == "" {
		return 0
	}
	name := normalizeQuery(recipe.Name)
	applicationID := normalizeQuery(recipe.ID)
	extensions := make([]string, 0, len(recipe.SupportedExtensions))
	for _, extension := range recipe.SupportedExtensions {
		extensions = append(extensions, normalizeQuery(strings.TrimPrefix(extension, ".")))
	}

	switch {
	case name == normalizedQuery || applicationID == normalizedQuery:
		return 100
	case strings.HasPrefix(name, normalizedQuery):
		return 95
	case strings.Contains(name, normalizedQuery):
		return 90
	case containsString(extensions, strings.TrimPrefix(normalizedQuery, ".")):
		return 85
	case strings.Contains(applicationID, normalizedQuery):
		return 75
	case naturalLaunchQuery(normalizedQuery, name):
		return 65
	case extensionLaunchQuery(normalizedQuery, extensions):
		return 55
	default:
		return 0
	}
}

func naturalLaunchQuery(normalizedQuery string, normalizedName string) bool {
	for _, verb := range []string{"open", "launch", "start", "run"} {
		if normalizedQuery == verb+" "+normalizedName || strings.HasSuffix(normalizedQuery, " "+normalizedName) {
			return true
		}
	}
	return false
}

func extensionLaunchQuery(normalizedQuery string, extensions []string) bool {
	tokens := strings.Fields(normalizedQuery)
	if !intersects(tokens, []string{"open", "launch", "start", "run", "file"}) {
		return false
	}
	for _, extension := range extensions {
		if containsString(tokens, extension) {
			return true
		}
	}
	return false
}

func modeLabel(mode string) string {
	switch mode {
	case "automatic":
		return "Automatic"
	case "wine":
		return "Managed compatibility"
	case "vm":
		return "Isolated environment"
	default:
		return "Automatic"
	}
}

func normalizedExtensions(extensions []string) []string {
	result := make([]string, 0, len(extensions))
	seen := make(map[string]bool, len(extensions))
	for _, extension := range extensions {
		normalized := "." + strings.ToLower(strings.TrimPrefix(extension, "."))
		if !seen[normalized] {
			result = append(result, normalized)
			seen[normalized] = true
		}
	}
	sort.Strings(result)
	return result
}

func singleLine(value string) bool {
	return value != "" && !strings.ContainsAny(value, "\r\n")
}

func singleLineOrBlank(value string) bool {
	return !strings.ContainsAny(value, "\r\n")
}

func normalizeQuery(value string) string {
	return strings.Join(strings.Fields(strings.ToLower(value)), " ")
}

func containsString(values []string, candidate string) bool {
	for _, value := range values {
		if value == candidate {
			return true
		}
	}
	return false
}

func intersects(left []string, right []string) bool {
	for _, value := range left {
		if containsString(right, value) {
			return true
		}
	}
	return false
}

func validateNoBackendTerms(value any, label string) error {
	encoded, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("encode %s JSON: %w", label, err)
	}
	text := strings.ToLower(string(encoded))
	forbidden := []string{"prefix", ".exe", "wine ", "wine/", "proton", "qemu-system", "program files", ".wine"}
	for _, term := range forbidden {
		if strings.Contains(text, term) {
			return fmt.Errorf("%s exposes forbidden backend term: %s", label, term)
		}
	}
	return nil
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
