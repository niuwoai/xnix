package appidentity

import "errors"

type BackendEnvironmentPreview struct {
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
	SelectedStrategy            string                      `json:"selected_strategy"`
	RecommendedProfileID        string                      `json:"recommended_profile_id"`
	EnvironmentState            string                      `json:"environment_state"`
	Profiles                    []BackendEnvironmentProfile `json:"profiles"`
	ProfileIDs                  []string                    `json:"profile_ids"`
	ProfileCount                int                         `json:"profile_count"`
	ReadyProfileCount           int                         `json:"ready_profile_count"`
	BlockedProfileCount         int                         `json:"blocked_profile_count"`
	RequiredReviews             []string                    `json:"required_reviews"`
	BridgeCapabilities          []BackendEnvironmentBridge  `json:"bridge_capabilities"`
	BridgeCapabilityIDs         []string                    `json:"bridge_capability_ids"`
	BridgeCapabilityCount       int                         `json:"bridge_capability_count"`
	RuntimeOwned                bool                        `json:"runtime_owned"`
	GoRuntimeBacked             bool                        `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                        `json:"kde_policy_owner"`
	UserVisible                 bool                        `json:"user_visible"`
	LocalEnvironmentReady       bool                        `json:"local_environment_ready"`
	IsolatedEnvironmentReady    bool                        `json:"isolated_environment_ready"`
	EnvironmentCreated          bool                        `json:"environment_created"`
	BackendProcessStarted       bool                        `json:"backend_process_started"`
	BackendBindingReady         bool                        `json:"backend_binding_ready"`
	LaunchEnabled               bool                        `json:"launch_enabled"`
	RequestObjectCreated        bool                        `json:"request_object_created"`
	HostStorageExposed          bool                        `json:"host_storage_exposed"`
	ClipboardBridgeEnabled      bool                        `json:"clipboard_bridge_enabled"`
	PrintBridgeEnabled          bool                        `json:"print_bridge_enabled"`
	FileBridgeEnabled           bool                        `json:"file_bridge_enabled"`
	PortalReviewRequired        bool                        `json:"portal_review_required"`
	SnapshotRequired            bool                        `json:"snapshot_required"`
	HostRootModified            bool                        `json:"host_root_modified"`
	NetworkRequired             bool                        `json:"network_required"`
	PrivilegedContainerRequired bool                        `json:"privileged_container_required"`
	BackendDetailsExposed       bool                        `json:"backend_details_exposed"`
	CompatibilityStorageExposed bool                        `json:"compatibility_storage_exposed"`
	RawBackendCommandExposed    bool                        `json:"raw_backend_command_exposed"`
	BlockedActions              []string                    `json:"blocked_actions"`
	UserFacingSettings          map[string]string           `json:"user_facing_settings"`
	DesktopSafeSummary          string                      `json:"desktop_safe_summary"`
}

type BackendEnvironmentProfile struct {
	ID                       string   `json:"id"`
	Label                    string   `json:"label"`
	Kind                     string   `json:"kind"`
	Status                   string   `json:"status"`
	Recommended              bool     `json:"recommended"`
	Ready                    bool     `json:"ready"`
	Blocked                  bool     `json:"blocked"`
	RequiredReviews          []string `json:"required_reviews"`
	EnvironmentCreated       bool     `json:"environment_created"`
	BackendProcessStarted    bool     `json:"backend_process_started"`
	StoragePathExposed       bool     `json:"storage_path_exposed"`
	BackendDetailsExposed    bool     `json:"backend_details_exposed"`
	PrivilegedContainerUsed  bool     `json:"privileged_container_used"`
	CompatibilityProfileOnly bool     `json:"compatibility_profile_only"`
	Summary                  string   `json:"summary"`
}

type BackendEnvironmentBridge struct {
	ID                         string `json:"id"`
	Label                      string `json:"label"`
	RuntimeMethod              string `json:"runtime_method"`
	State                      string `json:"state"`
	PortalRequired             bool   `json:"portal_required"`
	UserReviewRequired         bool   `json:"user_review_required"`
	BridgeEnabled              bool   `json:"bridge_enabled"`
	RequestObjectCreated       bool   `json:"request_object_created"`
	DirectBackendAccessAllowed bool   `json:"direct_backend_access_allowed"`
	BackendDetailsExposed      bool   `json:"backend_details_exposed"`
}

func (plan Plan) BackendEnvironmentPreview() (BackendEnvironmentPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return BackendEnvironmentPreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return BackendEnvironmentPreview{}, errors.New("backend environment preview requires single-line identity fields")
		}
	}

	recommendedProfileID := plan.recommendedCompatibilityProfileID()
	profiles := backendEnvironmentProfiles(recommendedProfileID)
	bridges := backendEnvironmentBridges()
	readyProfileCount := 0
	blockedProfileCount := 0
	for _, profile := range profiles {
		if profile.Ready {
			readyProfileCount++
		}
		if profile.Blocked {
			blockedProfileCount++
		}
	}

	preview := BackendEnvironmentPreview{
		SchemaVersion:               "xnix.runtime.backend_environment.v1",
		RequestType:                 "backend-environment-preview",
		PlanType:                    "compatibility-backend-environment-plan",
		Source:                      "backend-selection-preview+desktop-resource-bridge-preview",
		Desktop:                     "KDE Plasma",
		RuntimeMethod:               "GetBackendEnvironmentPlan",
		ApplicationID:               plan.ApplicationID,
		DisplayName:                 plan.DisplayName,
		Icon:                        plan.Icon,
		DesktopFile:                 plan.DesktopFile,
		SelectedStrategy:            recommendedProfileID + "-strategy",
		RecommendedProfileID:        recommendedProfileID,
		EnvironmentState:            "planned-blocked",
		Profiles:                    profiles,
		ProfileIDs:                  backendEnvironmentProfileIDs(profiles),
		ProfileCount:                len(profiles),
		ReadyProfileCount:           readyProfileCount,
		BlockedProfileCount:         blockedProfileCount,
		RequiredReviews:             []string{"package-source-review", "application-state-root-review", "portal-policy-review", "snapshot-baseline-review"},
		BridgeCapabilities:          bridges,
		BridgeCapabilityIDs:         backendEnvironmentBridgeIDs(bridges),
		BridgeCapabilityCount:       len(bridges),
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		UserVisible:                 true,
		LocalEnvironmentReady:       false,
		IsolatedEnvironmentReady:    false,
		EnvironmentCreated:          false,
		BackendProcessStarted:       false,
		BackendBindingReady:         false,
		LaunchEnabled:               false,
		RequestObjectCreated:        false,
		HostStorageExposed:          false,
		ClipboardBridgeEnabled:      false,
		PrintBridgeEnabled:          false,
		FileBridgeEnabled:           false,
		PortalReviewRequired:        true,
		SnapshotRequired:            true,
		HostRootModified:            false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		CompatibilityStorageExposed: false,
		RawBackendCommandExposed:    false,
		BlockedActions: []string{
			"create local compatibility environment from KDE",
			"create isolated compatibility environment from KDE",
			"bind compatibility profile before Runtime preflight",
			"start compatibility backend from environment preview",
			"enable clipboard bridge before Portal review",
			"enable print bridge before Portal review",
			"expose backend command to desktop shell",
			"mutate host root during backend environment planning",
		},
		UserFacingSettings: plan.UserFacingSettings,
		DesktopSafeSummary: "Runtime can describe the compatibility environment plan and required desktop bridges, but environment creation and backend startup remain blocked until Runtime-owned preflight passes.",
	}
	if err := validateNoBackendTerms(preview, "backend environment preview"); err != nil {
		return BackendEnvironmentPreview{}, err
	}
	return preview, nil
}

func backendEnvironmentProfiles(recommendedProfileID string) []BackendEnvironmentProfile {
	return []BackendEnvironmentProfile{
		backendEnvironmentProfile("local-compatibility-environment", "Local compatibility environment", "local", recommendedProfileID == "local-compatibility"),
		backendEnvironmentProfile("isolated-compatibility-environment", "Isolated compatibility environment", "isolated", recommendedProfileID == "isolated-compatibility"),
	}
}

func backendEnvironmentProfile(id string, label string, kind string, recommended bool) BackendEnvironmentProfile {
	status := "blocked"
	summaryState := "available after Runtime preflight"
	if recommended {
		summaryState = "recommended after Runtime preflight"
	}

	return BackendEnvironmentProfile{
		ID:                       id,
		Label:                    label,
		Kind:                     kind,
		Status:                   status,
		Recommended:              recommended,
		Ready:                    false,
		Blocked:                  true,
		RequiredReviews:          []string{"package-source-review", "state-root-review", "portal-policy-review", "snapshot-baseline-review"},
		EnvironmentCreated:       false,
		BackendProcessStarted:    false,
		StoragePathExposed:       false,
		BackendDetailsExposed:    false,
		PrivilegedContainerUsed:  false,
		CompatibilityProfileOnly: true,
		Summary:                  label + " is " + summaryState + ".",
	}
}

func backendEnvironmentBridges() []BackendEnvironmentBridge {
	return []BackendEnvironmentBridge{
		backendEnvironmentBridge("file-open", "File access"),
		backendEnvironmentBridge("uri-open", "URI open"),
		backendEnvironmentBridge("clipboard", "Clipboard"),
		backendEnvironmentBridge("print", "Print"),
		backendEnvironmentBridge("screenshot", "Screenshot"),
	}
}

func backendEnvironmentBridge(id string, label string) BackendEnvironmentBridge {
	return BackendEnvironmentBridge{
		ID:                         id,
		Label:                      label,
		RuntimeMethod:              "GetPortalRequestPlan",
		State:                      "planned",
		PortalRequired:             true,
		UserReviewRequired:         true,
		BridgeEnabled:              false,
		RequestObjectCreated:       false,
		DirectBackendAccessAllowed: false,
		BackendDetailsExposed:      false,
	}
}

func backendEnvironmentProfileIDs(profiles []BackendEnvironmentProfile) []string {
	ids := make([]string, 0, len(profiles))
	for _, profile := range profiles {
		ids = append(ids, profile.ID)
	}
	return ids
}

func backendEnvironmentBridgeIDs(bridges []BackendEnvironmentBridge) []string {
	ids := make([]string, 0, len(bridges))
	for _, bridge := range bridges {
		ids = append(ids, bridge.ID)
	}
	return ids
}
