package appidentity

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

type LaunchIntentPreview struct {
	SchemaVersion             string            `json:"schema_version"`
	RequestType               string            `json:"request_type"`
	IntentType                string            `json:"intent_type"`
	Source                    string            `json:"source"`
	Desktop                   string            `json:"desktop"`
	RuntimeMethod             string            `json:"runtime_method"`
	ReadMethod                string            `json:"read_method"`
	ApplicationID             string            `json:"application_id"`
	ApplicationName           string            `json:"application_name"`
	Icon                      string            `json:"icon"`
	DesktopFile               string            `json:"desktop_file"`
	LauncherCommand           []string          `json:"launcher_command"`
	ExecutionState            string            `json:"execution_state"`
	OverallStatus             string            `json:"overall_status"`
	WriteGateDecision         string            `json:"write_gate_decision"`
	DenialErrorName           string            `json:"denial_error_name"`
	PortalRequired            bool              `json:"portal_required"`
	SnapshotRequired          bool              `json:"snapshot_required"`
	CompatibilityProfile      ReadinessProfile  `json:"compatibility_profile"`
	RunPlan                   LaunchRunPlan     `json:"run_plan"`
	FileCount                 int               `json:"file_count"`
	FileURIs                  []string          `json:"file_uris"`
	RuntimeOwned              bool              `json:"runtime_owned"`
	GoRuntimeBacked           bool              `json:"go_runtime_backed"`
	KDEPolicyOwner            bool              `json:"kde_policy_owner"`
	StandardDesktopEntry      bool              `json:"standard_desktop_entry"`
	LaunchUsesRuntime         bool              `json:"launch_uses_runtime"`
	DesktopEntryLaunchVisible bool              `json:"desktop_entry_launch_visible"`
	LaunchAllowed             bool              `json:"launch_allowed"`
	LaunchEnabled             bool              `json:"launch_enabled"`
	ExecutionRequestCreated   bool              `json:"execution_request_created"`
	ExecutionStarted          bool              `json:"execution_started"`
	BackendBindingReady       bool              `json:"backend_binding_ready"`
	RequestObjectCreated      bool              `json:"request_object_created"`
	PermissionGranted         bool              `json:"permission_granted"`
	HostRootModified          bool              `json:"host_root_modified"`
	NetworkRequired           bool              `json:"network_required"`
	BackendDetailsExposed     bool              `json:"backend_details_exposed"`
	BlockedActions            []string          `json:"blocked_actions"`
	UserFacingSettings        map[string]string `json:"user_facing_settings"`
	DesktopSafeSummary        string            `json:"desktop_safe_summary"`
}

type LaunchRunPlan struct {
	PlanType                  string `json:"plan_type"`
	Strategy                  string `json:"strategy"`
	BackendDetailsExposed     bool   `json:"backend_details_exposed"`
	BackendReady              bool   `json:"backend_ready"`
	PortalPolicyRequired      bool   `json:"portal_policy_required"`
	SnapshotBeforeRiskyChange bool   `json:"snapshot_before_risky_change"`
	RuntimeWriteGateRequired  bool   `json:"runtime_write_gate_required"`
	ExecutionRequestCreated   bool   `json:"execution_request_created"`
}

func (plan Plan) LaunchIntentPreview(fileURIs []string) (LaunchIntentPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return LaunchIntentPreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return LaunchIntentPreview{}, errors.New("launch intent preview requires single-line identity fields")
		}
	}
	normalizedURIs, err := normalizeLaunchIntentFileURIs(fileURIs)
	if err != nil {
		return LaunchIntentPreview{}, err
	}
	readiness, err := plan.ExecutionReadinessPreview()
	if err != nil {
		return LaunchIntentPreview{}, err
	}

	preview := LaunchIntentPreview{
		SchemaVersion:        "xnix.runtime.launch_intent.v1",
		RequestType:          "launch-intent-preview",
		IntentType:           "runtime-launch-intent",
		Source:               "desktop-launcher",
		Desktop:              "KDE Plasma",
		RuntimeMethod:        "Launch",
		ReadMethod:           "GetLaunchIntent",
		ApplicationID:        plan.ApplicationID,
		ApplicationName:      plan.DisplayName,
		Icon:                 plan.Icon,
		DesktopFile:          plan.DesktopFile,
		LauncherCommand:      plan.LaunchCommand,
		ExecutionState:       readiness.ExecutionState,
		OverallStatus:        readiness.OverallStatus,
		WriteGateDecision:    "blocked-until-production-backend",
		DenialErrorName:      "org.xnix.Compatibility1.Error.WriteMethodDisabled",
		PortalRequired:       len(normalizedURIs) > 0,
		SnapshotRequired:     readiness.SnapshotRequired,
		CompatibilityProfile: readiness.CompatibilityProfile,
		RunPlan: LaunchRunPlan{
			PlanType:                  "compatibility-run",
			Strategy:                  plan.runtimeModeID() + "-managed",
			BackendDetailsExposed:     false,
			BackendReady:              false,
			PortalPolicyRequired:      true,
			SnapshotBeforeRiskyChange: true,
			RuntimeWriteGateRequired:  true,
			ExecutionRequestCreated:   false,
		},
		FileCount:                 len(normalizedURIs),
		FileURIs:                  normalizedURIs,
		RuntimeOwned:              true,
		GoRuntimeBacked:           true,
		KDEPolicyOwner:            false,
		StandardDesktopEntry:      true,
		LaunchUsesRuntime:         true,
		DesktopEntryLaunchVisible: true,
		LaunchAllowed:             false,
		LaunchEnabled:             false,
		ExecutionRequestCreated:   false,
		ExecutionStarted:          false,
		BackendBindingReady:       false,
		RequestObjectCreated:      false,
		PermissionGranted:         false,
		HostRootModified:          false,
		NetworkRequired:           false,
		BackendDetailsExposed:     false,
		BlockedActions:            []string{"create launch request object before Runtime gates pass", "start compatibility profile from KDE", "expose raw backend command to desktop shell", "grant desktop resources without Portal review", "mutate host root during launch intent planning"},
		UserFacingSettings:        plan.UserFacingSettings,
		DesktopSafeSummary:        "KDE launch intent is captured without executing; Launch remains gated by the Runtime.",
	}
	if err := validateNoBackendTerms(preview, "launch intent preview"); err != nil {
		return LaunchIntentPreview{}, err
	}
	return preview, nil
}

func normalizeLaunchIntentFileURIs(fileURIs []string) ([]string, error) {
	normalized := make([]string, 0, len(fileURIs))
	for _, fileURI := range fileURIs {
		parsed, err := url.Parse(fileURI)
		if err != nil {
			return nil, fmt.Errorf("invalid file URI: %s", fileURI)
		}
		if parsed.Scheme != "file" {
			return nil, errors.New("only file URIs are accepted")
		}
		if parsed.Path == "" || !strings.HasPrefix(parsed.Path, "/") {
			return nil, errors.New("file URI must include an absolute path")
		}
		normalized = append(normalized, parsed.String())
	}
	return normalized, nil
}
