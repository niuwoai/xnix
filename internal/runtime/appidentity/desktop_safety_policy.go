package appidentity

type DesktopSafetyPolicyPreview struct {
	SchemaVersion             string   `json:"schema_version"`
	RequestType               string   `json:"request_type"`
	PolicyType                string   `json:"policy_type"`
	Desktop                   string   `json:"desktop"`
	RuntimeMethod             string   `json:"runtime_method"`
	EntryPoints               []string `json:"entrypoints"`
	EntryPointCount           int      `json:"entrypoint_count"`
	SettingsFieldIDs          []string `json:"settings_field_ids"`
	ForbiddenUserTerms        []string `json:"forbidden_user_terms"`
	SafetyFalseKeys           []string `json:"safety_false_keys"`
	RuntimeOwned              bool     `json:"runtime_owned"`
	GoRuntimeBacked           bool     `json:"go_runtime_backed"`
	KDEPolicyOwner            bool     `json:"kde_policy_owner"`
	UserVisible               bool     `json:"user_visible"`
	BackendTerminologyHidden  bool     `json:"backend_terminology_hidden"`
	WriteMethodsEnabled       bool     `json:"write_methods_enabled"`
	BackendLaunchEnabled      bool     `json:"backend_launch_enabled"`
	ExecutionEnabled          bool     `json:"execution_enabled"`
	RealPortalTransport       bool     `json:"real_portal_transport_enabled"`
	AIProviderCallEnabled     bool     `json:"ai_provider_call_enabled"`
	HostRootModified          bool     `json:"host_root_modified"`
	NetworkRequired           bool     `json:"network_required"`
	PrivilegedContainerNeeded bool     `json:"privileged_container_required"`
	DesktopSafeSummary        string   `json:"desktop_safe_summary"`
}

var kdeFirstEntryPoints = []string{
	"launcher",
	"task-manager",
	"file-manager",
	"system-tray",
	"notifications",
	"compatibility-center",
	"settings",
}

var userFacingSettingsFieldIDs = []string{
	"mode",
	"preference",
	"documents",
	"downloads",
	"camera",
	"network",
	"snapshots",
}

var forbiddenDesktopUserTerms = []string{
	"prefix",
	"bottle",
	"wine",
	"proton",
	".exe",
	"Program Files",
	"qemu-system",
	"/Users/",
	"/private/",
	"/var/",
	".wine",
	"docker.sock",
}

var desktopSafetyFalseKeys = []string{
	"backend_details_exposed",
	"backend_launch_enabled",
	"execution_started",
	"host_permission_changed",
	"host_root_modified",
	"network_required",
	"permission_granted",
	"privileged_container_required",
	"raw_command_exposed",
	"raw_executable_exposed",
	"raw_windows_executable_exposed",
	"real_portal_transport_enabled",
	"request_object_created",
	"settings_persisted",
	"task_manager_entry_active",
}

func NewDesktopSafetyPolicyPreview() DesktopSafetyPolicyPreview {
	entryPoints := append([]string(nil), kdeFirstEntryPoints...)
	settings := append([]string(nil), userFacingSettingsFieldIDs...)
	forbidden := append([]string(nil), forbiddenDesktopUserTerms...)
	falseKeys := append([]string(nil), desktopSafetyFalseKeys...)

	return DesktopSafetyPolicyPreview{
		SchemaVersion:             "xnix.runtime.desktop_safety_policy.v1",
		RequestType:               "desktop-safety-policy-preview",
		PolicyType:                "kde-first-user-facing-safety-policy",
		Desktop:                   "KDE Plasma",
		RuntimeMethod:             "GetKDEIntegrationStatus",
		EntryPoints:               entryPoints,
		EntryPointCount:           len(entryPoints),
		SettingsFieldIDs:          settings,
		ForbiddenUserTerms:        forbidden,
		SafetyFalseKeys:           falseKeys,
		RuntimeOwned:              true,
		GoRuntimeBacked:           true,
		KDEPolicyOwner:            false,
		UserVisible:               true,
		BackendTerminologyHidden:  true,
		WriteMethodsEnabled:       false,
		BackendLaunchEnabled:      false,
		ExecutionEnabled:          false,
		RealPortalTransport:       false,
		AIProviderCallEnabled:     false,
		HostRootModified:          false,
		NetworkRequired:           false,
		PrivilegedContainerNeeded: false,
		DesktopSafeSummary:        "KDE-facing Runtime output must use user concepts and hide compatibility backend terminology, raw executable paths, host paths, and disabled execution details.",
	}
}
