package appidentity

import (
	"errors"
	"strings"
)

type KDEIntegrationStatusPreview struct {
	SchemaVersion         string                      `json:"schema_version"`
	RequestType           string                      `json:"request_type"`
	StatusType            string                      `json:"status_type"`
	Source                string                      `json:"source"`
	RuntimeMethod         string                      `json:"runtime_method"`
	ReadMethod            string                      `json:"read_method"`
	Desktop               string                      `json:"desktop"`
	EntryPoints           []KDEIntegrationEntryPoint  `json:"entry_points"`
	EntryPointIDs         []string                    `json:"entry_point_ids"`
	EntryPointCount       int                         `json:"entry_point_count"`
	InitialCount          int                         `json:"initial_count"`
	PlannedCount          int                         `json:"planned_count"`
	CompleteCount         int                         `json:"complete_count"`
	Checks                []KDEIntegrationStatusCheck `json:"checks"`
	CheckIDs              []string                    `json:"check_ids"`
	RuntimeOwned          bool                        `json:"runtime_owned"`
	GoRuntimeBacked       bool                        `json:"go_runtime_backed"`
	KDEPolicyOwner        bool                        `json:"kde_policy_owner"`
	OfficialDesktopOnly   bool                        `json:"official_desktop_only"`
	StableDesktopContract bool                        `json:"stable_desktop_contract"`
	HostRootModified      bool                        `json:"host_root_modified"`
	BackendDetailsExposed bool                        `json:"backend_details_exposed"`
	BlockedActions        []string                    `json:"blocked_actions"`
	DesktopSafeSummary    string                      `json:"desktop_safe_summary"`
}

type KDEIntegrationEntryPoint struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	State             string `json:"state"`
	RuntimeMethod     string `json:"runtime_method"`
	AdapterRole       string `json:"adapter_role"`
	RuntimeBacked     bool   `json:"runtime_backed"`
	GoRuntimeBacked   bool   `json:"go_runtime_backed"`
	DBusReadAvailable bool   `json:"dbus_read_available"`
	KDEPolicyOwner    bool   `json:"kde_policy_owner"`
	Summary           string `json:"summary"`
	NextStep          string `json:"next_step"`
}

type KDEIntegrationStatusCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type KDEShellIntegrationPlanPreview struct {
	SchemaVersion               string                         `json:"schema_version"`
	RequestType                 string                         `json:"request_type"`
	PlanType                    string                         `json:"plan_type"`
	Source                      string                         `json:"source"`
	RuntimeMethod               string                         `json:"runtime_method"`
	ReadMethod                  string                         `json:"read_method"`
	DesktopShell                string                         `json:"desktop_shell"`
	Components                  []KDEShellIntegrationComponent `json:"components"`
	ComponentIDs                []string                       `json:"component_ids"`
	ComponentCount              int                            `json:"component_count"`
	InitialComponentCount       int                            `json:"initial_component_count"`
	PlannedComponentCount       int                            `json:"planned_component_count"`
	RuntimeOwned                bool                           `json:"runtime_owned"`
	GoRuntimeBacked             bool                           `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                           `json:"kde_policy_owner"`
	OfficialDesktopOnly         bool                           `json:"official_desktop_only"`
	FallbackDesktopsSupported   bool                           `json:"fallback_desktops_supported"`
	PlasmaForkRequired          bool                           `json:"plasma_fork_required"`
	PlasmaSourceModified        bool                           `json:"plasma_source_modified"`
	ShellConfigurationWritten   bool                           `json:"shell_configuration_written"`
	ComponentActivationEnabled  bool                           `json:"component_activation_enabled"`
	BackendLaunchEnabled        bool                           `json:"backend_launch_enabled"`
	HostRootModified            bool                           `json:"host_root_modified"`
	PrivilegedContainerRequired bool                           `json:"privileged_container_required"`
	BackendDetailsExposed       bool                           `json:"backend_details_exposed"`
	BlockedActions              []string                       `json:"blocked_actions"`
	DesktopSafeSummary          string                         `json:"desktop_safe_summary"`
}

type KDEShellIntegrationComponent struct {
	ID                    string `json:"id"`
	Label                 string `json:"label"`
	State                 string `json:"state"`
	RuntimeMethod         string `json:"runtime_method"`
	RuntimeOwned          bool   `json:"runtime_owned"`
	GoRuntimeBacked       bool   `json:"go_runtime_backed"`
	KDEPolicyOwner        bool   `json:"kde_policy_owner"`
	ShellWritesEnabled    bool   `json:"shell_writes_enabled"`
	BackendLaunchEnabled  bool   `json:"backend_launch_enabled"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
	Summary               string `json:"summary"`
}

type KDEApplicationSurfacePlanPreview struct {
	SchemaVersion                 string                        `json:"schema_version"`
	RequestType                   string                        `json:"request_type"`
	PlanType                      string                        `json:"plan_type"`
	Source                        string                        `json:"source"`
	RuntimeMethod                 string                        `json:"runtime_method"`
	ReadMethod                    string                        `json:"read_method"`
	Application                   KDEApplicationSurfaceIdentity `json:"application"`
	SurfaceState                  string                        `json:"surface_state"`
	DesktopShell                  string                        `json:"desktop_shell"`
	EntryPoints                   []KDEApplicationSurfaceEntry  `json:"entry_points"`
	EntryPointIDs                 []string                      `json:"entry_point_ids"`
	EntryPointCount               int                           `json:"entry_point_count"`
	RequiredRuntimeGates          []string                      `json:"required_runtime_gates"`
	ActivationReceiptRoot         bool                          `json:"activation_receipt_root"`
	ActivationReceiptBacked       bool                          `json:"activation_receipt_backed"`
	ActivationReceiptPath         string                        `json:"activation_receipt_path,omitempty"`
	RuntimeOwned                  bool                          `json:"runtime_owned"`
	GoRuntimeBacked               bool                          `json:"go_runtime_backed"`
	KDEPolicyOwner                bool                          `json:"kde_policy_owner"`
	OfficialDesktopOnly           bool                          `json:"official_desktop_only"`
	NormalLinuxApplicationSurface bool                          `json:"normal_linux_application_surface"`
	StandardLauncherVisible       bool                          `json:"standard_launcher_visible"`
	TaskManagerIdentityReady      bool                          `json:"task_manager_identity_ready"`
	FileAssociationsPlanned       bool                          `json:"file_associations_planned"`
	DolphinActionPlanned          bool                          `json:"dolphin_action_planned"`
	KRunnerQueryPlanned           bool                          `json:"krunner_query_planned"`
	TrayStatusPlanned             bool                          `json:"tray_status_planned"`
	NotificationRoutePlanned      bool                          `json:"notification_route_planned"`
	SettingsSurfacePlanned        bool                          `json:"settings_surface_planned"`
	PortalReviewRequired          bool                          `json:"portal_review_required"`
	ExecutionReady                bool                          `json:"execution_ready"`
	LaunchEnabled                 bool                          `json:"launch_enabled"`
	BackendProcessStarted         bool                          `json:"backend_process_started"`
	DesktopFilesWritten           bool                          `json:"desktop_files_written"`
	MIMEAppsWritten               bool                          `json:"mimeapps_written"`
	HostRootModified              bool                          `json:"host_root_modified"`
	BackendCommandExposed         bool                          `json:"backend_command_exposed"`
	RawWindowsExecutableExposed   bool                          `json:"raw_windows_executable_exposed"`
	BackendDetailsExposed         bool                          `json:"backend_details_exposed"`
	DesktopSafeSummary            string                        `json:"desktop_safe_summary"`
}

type KDEApplicationSurfaceOptions struct {
	ActivationRoot string
}

type KDEApplicationSurfaceIdentity struct {
	ID                  string   `json:"id"`
	Name                string   `json:"name"`
	Icon                string   `json:"icon"`
	RequestedMode       string   `json:"requested_mode"`
	DesktopFile         string   `json:"desktop_file"`
	SupportedExtensions []string `json:"supported_extensions"`
}

type KDEApplicationSurfaceEntry struct {
	ID                    string `json:"id"`
	Name                  string `json:"name"`
	RuntimeMethod         string `json:"runtime_method"`
	State                 string `json:"state"`
	RuntimeBacked         bool   `json:"runtime_backed"`
	GoRuntimeBacked       bool   `json:"go_runtime_backed"`
	KDEWritesPolicy       bool   `json:"kde_writes_policy"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

func NewKDEIntegrationStatusPreview() (KDEIntegrationStatusPreview, error) {
	entryPoints := kdeIntegrationEntryPoints()
	ids, initial, planned, complete := summarizeKDEIntegrationEntryPoints(entryPoints)
	checks := []KDEIntegrationStatusCheck{
		{ID: "official-desktop", Status: "pass", Summary: "KDE Plasma is the only official first-release desktop shell."},
		{ID: "runtime-policy-owner", Status: "pass", Summary: "KDE entry points read Runtime policy instead of owning compatibility decisions."},
		{ID: "desktop-safety", Status: "pass", Summary: "The status read model does not write shell configuration or expose implementation details."},
	}
	preview := KDEIntegrationStatusPreview{
		SchemaVersion:         "xnix.runtime.kde_integration_status.v1",
		RequestType:           "kde-integration-status-preview",
		StatusType:            "kde-integration-status",
		Source:                "go-runtime-kde-integration-status",
		RuntimeMethod:         "GetKDEIntegrationStatus",
		ReadMethod:            "GetKDEIntegrationStatusPreview",
		Desktop:               "KDE Plasma",
		EntryPoints:           entryPoints,
		EntryPointIDs:         ids,
		EntryPointCount:       len(entryPoints),
		InitialCount:          initial,
		PlannedCount:          planned,
		CompleteCount:         complete,
		Checks:                checks,
		CheckIDs:              kdeIntegrationCheckIDs(checks),
		RuntimeOwned:          true,
		GoRuntimeBacked:       true,
		KDEPolicyOwner:        false,
		OfficialDesktopOnly:   true,
		StableDesktopContract: true,
		HostRootModified:      false,
		BackendDetailsExposed: false,
		BlockedActions:        []string{"treat a non-KDE shell as first-release complete", "let KDE own Runtime compatibility policy", "write shell configuration from status preview", "start compatibility profile from KDE status preview"},
		DesktopSafeSummary:    "KDE Plasma is the only official first-release shell, and all seven entry points are Runtime-backed Go read models.",
	}
	if err := validateNoBackendTerms(preview, "KDE integration status preview"); err != nil {
		return KDEIntegrationStatusPreview{}, err
	}
	return preview, nil
}

func NewKDEShellIntegrationPlanPreview() (KDEShellIntegrationPlanPreview, error) {
	components := kdeShellIntegrationComponents()
	ids, initial, planned := summarizeKDEShellComponents(components)
	preview := KDEShellIntegrationPlanPreview{
		SchemaVersion:               "xnix.runtime.kde_shell_integration.v1",
		RequestType:                 "kde-shell-integration-preview",
		PlanType:                    "kde-shell-integration-plan",
		Source:                      "go-runtime-kde-shell-integration",
		RuntimeMethod:               "GetKDEShellIntegrationPlan",
		ReadMethod:                  "GetKDEShellIntegrationPlanPreview",
		DesktopShell:                "KDE Plasma",
		Components:                  components,
		ComponentIDs:                ids,
		ComponentCount:              len(components),
		InitialComponentCount:       initial,
		PlannedComponentCount:       planned,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		OfficialDesktopOnly:         true,
		FallbackDesktopsSupported:   false,
		PlasmaForkRequired:          false,
		PlasmaSourceModified:        false,
		ShellConfigurationWritten:   false,
		ComponentActivationEnabled:  false,
		BackendLaunchEnabled:        false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		BlockedActions:              []string{"fork Plasma from shell integration preview", "write panel or component configuration from preview", "activate live shell components from preview", "start compatibility profile from shell integration preview", "mutate host root from shell integration preview"},
		DesktopSafeSummary:          "KDE Plasma remains a replaceable presentation shell while the Runtime owns compatibility policy and state.",
	}
	if err := validateNoBackendTerms(preview, "KDE shell integration preview"); err != nil {
		return KDEShellIntegrationPlanPreview{}, err
	}
	return preview, nil
}

func (plan Plan) KDEApplicationSurfacePlanPreview() (KDEApplicationSurfacePlanPreview, error) {
	return plan.KDEApplicationSurfacePlanPreviewWithOptions(KDEApplicationSurfaceOptions{})
}

func (plan Plan) KDEApplicationSurfacePlanPreviewWithOptions(options KDEApplicationSurfaceOptions) (KDEApplicationSurfacePlanPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return KDEApplicationSurfacePlanPreview{}, err
	}
	if !singleLine(plan.ApplicationID) || !singleLine(plan.DisplayName) || !singleLine(plan.Icon) || !singleLine(plan.DesktopFile) {
		return KDEApplicationSurfacePlanPreview{}, errors.New("KDE application surface preview requires single-line identity fields")
	}
	entries := kdeApplicationSurfaceEntries()
	source := "go-runtime-kde-application-surface"
	activationReceiptRoot := strings.TrimSpace(options.ActivationRoot) != ""
	activationReceiptBacked := false
	activationReceiptPath := ""
	if activationReceiptRoot {
		evidence, err := plan.DesktopActivationReceiptEvidence(options.ActivationRoot)
		if err != nil {
			return KDEApplicationSurfacePlanPreview{}, err
		}
		source = "go-runtime-kde-application-surface+desktop-activation-receipt"
		activationReceiptBacked = evidence.SafeForKDE
		activationReceiptPath = evidence.ReceiptRelativePath
	}
	preview := KDEApplicationSurfacePlanPreview{
		SchemaVersion: "xnix.runtime.kde_application_surface.v1",
		RequestType:   "kde-application-surface-preview",
		PlanType:      "kde-application-surface-plan",
		Source:        source,
		RuntimeMethod: "GetKDEApplicationSurfacePlan",
		ReadMethod:    "GetKDEApplicationSurfacePlanPreview",
		Application: KDEApplicationSurfaceIdentity{
			ID:                  plan.ApplicationID,
			Name:                plan.DisplayName,
			Icon:                plan.Icon,
			RequestedMode:       plan.RecipeMode,
			DesktopFile:         plan.DesktopFile,
			SupportedExtensions: plan.MIMETypes,
		},
		SurfaceState:                  "planned",
		DesktopShell:                  "KDE Plasma",
		EntryPoints:                   entries,
		EntryPointIDs:                 kdeApplicationSurfaceEntryIDs(entries),
		EntryPointCount:               len(entries),
		RequiredRuntimeGates:          []string{"recipe-install-gate", "portal-policy-review", "snapshot-baseline", "backend-environment-plan", "backend-lifecycle-plan", "runtime-write-gate"},
		ActivationReceiptRoot:         activationReceiptRoot,
		ActivationReceiptBacked:       activationReceiptBacked,
		ActivationReceiptPath:         activationReceiptPath,
		RuntimeOwned:                  true,
		GoRuntimeBacked:               true,
		KDEPolicyOwner:                false,
		OfficialDesktopOnly:           true,
		NormalLinuxApplicationSurface: true,
		StandardLauncherVisible:       true,
		TaskManagerIdentityReady:      true,
		FileAssociationsPlanned:       true,
		DolphinActionPlanned:          true,
		KRunnerQueryPlanned:           true,
		TrayStatusPlanned:             true,
		NotificationRoutePlanned:      true,
		SettingsSurfacePlanned:        true,
		PortalReviewRequired:          true,
		ExecutionReady:                false,
		LaunchEnabled:                 false,
		BackendProcessStarted:         false,
		DesktopFilesWritten:           false,
		MIMEAppsWritten:               false,
		HostRootModified:              false,
		BackendCommandExposed:         false,
		RawWindowsExecutableExposed:   false,
		BackendDetailsExposed:         false,
		DesktopSafeSummary:            "KDE can present this compatibility application as a normal Linux application while Runtime gates still block execution.",
	}
	if err := validateNoBackendTerms(preview, "KDE application surface preview"); err != nil {
		return KDEApplicationSurfacePlanPreview{}, err
	}
	return preview, nil
}

func kdeIntegrationEntryPoints() []KDEIntegrationEntryPoint {
	return []KDEIntegrationEntryPoint{
		kdeIntegrationEntryPoint("launcher", "Start menu", "initial", "GetDesktopEntryPlan", "standard-desktop-entry", "Compatibility applications appear in the KDE launcher through generated desktop entries.", "Connect activation receipts to the production recipe installer."),
		kdeIntegrationEntryPoint("task-manager", "Task manager", "initial", "GetTaskManagerIdentityPlan", "window-identity-and-restore", "Compatibility windows expose grouping, pinning, switcher, and restore identity.", "Connect Runtime identity plans to the production KWin script and task manager bridge."),
		kdeIntegrationEntryPoint("file-manager", "File manager", "initial", "GetFileAssociationPlan", "dolphin-service-menu-and-mime", "Dolphin opens selected files through portal-mediated Runtime file-open planning.", "Connect file-open requests to production Runtime launch requests."),
		kdeIntegrationEntryPoint("system-tray", "System tray", "planned", "GetTrayStatus", "runtime-status-surface", "The tray can show Runtime activity, attention state, and bridge readiness.", "Connect tray status plans to a production Plasma tray surface."),
		kdeIntegrationEntryPoint("notifications", "Notifications", "initial", "GetNotificationPlan", "runtime-event-notification", "Runtime events map to KDE notification payloads for install, repair, mode, and approval states.", "Connect notification plans to the production KDE notification path."),
		kdeIntegrationEntryPoint("compatibility-center", "Compatibility Center", "initial", "GetCompatibilityCenterSummary", "plasma-read-model", "The Compatibility Center can show Runtime-owned application state and safe action cards.", "Render live Runtime applications and diagnostics in the Plasmoid."),
		kdeIntegrationEntryPoint("settings", "Unified settings", "initial", "GetCompatibilitySettings", "user-facing-policy-controls", "Settings expose user-facing Runtime policy without implementation terminology.", "Connect settings plans to a production KDE settings module and persisted Runtime policy."),
	}
}

func kdeIntegrationEntryPoint(id, name, state, runtimeMethod, adapterRole, summary, nextStep string) KDEIntegrationEntryPoint {
	return KDEIntegrationEntryPoint{
		ID:                id,
		Name:              name,
		State:             state,
		RuntimeMethod:     runtimeMethod,
		AdapterRole:       adapterRole,
		RuntimeBacked:     true,
		GoRuntimeBacked:   true,
		DBusReadAvailable: true,
		KDEPolicyOwner:    false,
		Summary:           summary,
		NextStep:          nextStep,
	}
}

func kdeShellIntegrationComponents() []KDEShellIntegrationComponent {
	return []KDEShellIntegrationComponent{
		kdeShellIntegrationComponent("start-menu", "Start menu", "initial", "GetDesktopEntryPlan", "Generated desktop entries make compatibility applications visible beside native Linux applications."),
		kdeShellIntegrationComponent("task-manager", "Task manager", "initial", "GetTaskManagerIdentityPlan", "Window identity plans let compatibility windows group and restore like normal applications."),
		kdeShellIntegrationComponent("file-manager", "Dolphin file manager", "initial", "GetFileAssociationPlan", "Dolphin actions and MIME associations route file opens through Runtime planning."),
		kdeShellIntegrationComponent("system-tray", "System tray", "planned", "GetTrayStatus", "Tray status is modeled while live tray bridging remains disabled."),
		kdeShellIntegrationComponent("notification-center", "Notification center", "initial", "GetNotificationPlan", "Runtime notification plans describe install, repair, approval, and mode-change events."),
		kdeShellIntegrationComponent("compatibility-center", "Compatibility Center", "initial", "GetCompatibilityCenterSummary", "Compatibility Center reads Runtime diagnostics, review tasks, and repair state."),
		kdeShellIntegrationComponent("unified-settings", "Unified settings", "initial", "GetCompatibilitySettings", "Settings expose user-facing Runtime policy without implementation terminology."),
		kdeShellIntegrationComponent("krunner-search", "KRunner search", "initial", "GetKRunnerQueryPlan", "KRunner query plans map natural-language queries to Runtime application identities."),
		kdeShellIntegrationComponent("kwin-window-management", "KWin window management", "initial", "GetKWinWindowRulePlan", "KWin rule plans provide identity and layout hints without KDE owning compatibility policy."),
	}
}

func kdeShellIntegrationComponent(id, label, state, runtimeMethod, summary string) KDEShellIntegrationComponent {
	return KDEShellIntegrationComponent{
		ID:                    id,
		Label:                 label,
		State:                 state,
		RuntimeMethod:         runtimeMethod,
		RuntimeOwned:          true,
		GoRuntimeBacked:       true,
		KDEPolicyOwner:        false,
		ShellWritesEnabled:    false,
		BackendLaunchEnabled:  false,
		BackendDetailsExposed: false,
		Summary:               summary,
	}
}

func kdeApplicationSurfaceEntries() []KDEApplicationSurfaceEntry {
	return []KDEApplicationSurfaceEntry{
		kdeApplicationSurfaceEntry("launcher", "Launcher", "GetDesktopEntryPlan"),
		kdeApplicationSurfaceEntry("task-manager", "Task Manager", "GetTaskManagerIdentityPlan"),
		kdeApplicationSurfaceEntry("file-manager", "File Manager", "GetFileAssociationPlan"),
		kdeApplicationSurfaceEntry("system-tray", "System Tray", "GetTrayStatus"),
		kdeApplicationSurfaceEntry("notifications", "Notifications", "GetNotificationPlan"),
		kdeApplicationSurfaceEntry("compatibility-center", "Compatibility Center", "GetCompatibilityCenterSummary"),
		kdeApplicationSurfaceEntry("settings", "Settings", "GetCompatibilitySettings"),
	}
}

func kdeApplicationSurfaceEntry(id, name, runtimeMethod string) KDEApplicationSurfaceEntry {
	return KDEApplicationSurfaceEntry{
		ID:                    id,
		Name:                  name,
		RuntimeMethod:         runtimeMethod,
		State:                 "planned",
		RuntimeBacked:         true,
		GoRuntimeBacked:       true,
		KDEWritesPolicy:       false,
		BackendDetailsExposed: false,
	}
}

func summarizeKDEIntegrationEntryPoints(entries []KDEIntegrationEntryPoint) ([]string, int, int, int) {
	ids := make([]string, 0, len(entries))
	initial := 0
	planned := 0
	complete := 0
	for _, entry := range entries {
		ids = append(ids, entry.ID)
		switch entry.State {
		case "initial":
			initial++
		case "planned":
			planned++
		case "complete":
			complete++
		}
	}
	return ids, initial, planned, complete
}

func summarizeKDEShellComponents(components []KDEShellIntegrationComponent) ([]string, int, int) {
	ids := make([]string, 0, len(components))
	initial := 0
	planned := 0
	for _, component := range components {
		ids = append(ids, component.ID)
		if component.State == "initial" {
			initial++
		}
		if component.State == "planned" {
			planned++
		}
	}
	return ids, initial, planned
}

func kdeApplicationSurfaceEntryIDs(entries []KDEApplicationSurfaceEntry) []string {
	ids := make([]string, 0, len(entries))
	for _, entry := range entries {
		ids = append(ids, entry.ID)
	}
	return ids
}

func kdeIntegrationCheckIDs(checks []KDEIntegrationStatusCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}
