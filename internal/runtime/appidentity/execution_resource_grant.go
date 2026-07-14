package appidentity

import "errors"

type ExecutionResourceGrantPreview struct {
	SchemaVersion             string                         `json:"schema_version"`
	RequestType               string                         `json:"request_type"`
	GrantType                 string                         `json:"grant_type"`
	RequestState              string                         `json:"request_state"`
	Source                    string                         `json:"source"`
	Desktop                   string                         `json:"desktop"`
	RuntimeMethod             string                         `json:"runtime_method"`
	ReadMethod                string                         `json:"read_method"`
	ApplicationID             string                         `json:"application_id"`
	ApplicationName           string                         `json:"application_name"`
	Icon                      string                         `json:"icon"`
	DesktopFile               string                         `json:"desktop_file"`
	LauncherCommand           []string                       `json:"launcher_command"`
	ExecutionPreflight        ExecutionResourcePreflight     `json:"execution_preflight"`
	ResourceGrants            []ExecutionResourceGrant       `json:"resource_grants"`
	BridgeSummary             ExecutionResourceBridgeSummary `json:"bridge_summary"`
	GrantCount                int                            `json:"grant_count"`
	AllowCount                int                            `json:"allow_count"`
	AskCount                  int                            `json:"ask_count"`
	DenyCount                 int                            `json:"deny_count"`
	PortalRequiredCount       int                            `json:"portal_required_count"`
	PendingReviewCount        int                            `json:"pending_review_count"`
	FileCount                 int                            `json:"file_count"`
	FileURIs                  []string                       `json:"file_uris"`
	RuntimeOwned              bool                           `json:"runtime_owned"`
	GoRuntimeBacked           bool                           `json:"go_runtime_backed"`
	KDEPolicyOwner            bool                           `json:"kde_policy_owner"`
	CompatibilityCenterCard   bool                           `json:"compatibility_center_card"`
	SafeForAIDiagnostics      bool                           `json:"safe_for_ai_diagnostics"`
	DesktopEntryLaunchVisible bool                           `json:"desktop_entry_launch_visible"`
	LaunchIntentCaptured      bool                           `json:"launch_intent_captured"`
	UserDecisionCaptured      bool                           `json:"user_decision_captured"`
	UserDecisionAllowsLaunch  bool                           `json:"user_decision_allows_launch"`
	PreflightPassed           bool                           `json:"preflight_passed"`
	PortalReviewRequired      bool                           `json:"portal_review_required"`
	SnapshotRequired          bool                           `json:"snapshot_required"`
	GrantPlanCreated          bool                           `json:"grant_plan_created"`
	GrantObjectsCreated       bool                           `json:"grant_objects_created"`
	RequestObjectsCreated     bool                           `json:"request_objects_created"`
	PermissionGranted         bool                           `json:"permission_granted"`
	ResourceBridgesEnabled    bool                           `json:"resource_bridges_enabled"`
	SettingsPersisted         bool                           `json:"settings_persisted"`
	RuntimeLaunchApproval     bool                           `json:"runtime_launch_approval"`
	LaunchAllowed             bool                           `json:"launch_allowed"`
	LaunchEnabled             bool                           `json:"launch_enabled"`
	ExecutionStarted          bool                           `json:"execution_started"`
	HostPermissionChanged     bool                           `json:"host_permission_changed"`
	HostRootModified          bool                           `json:"host_root_modified"`
	NetworkRequired           bool                           `json:"network_required"`
	BackendDetailsExposed     bool                           `json:"backend_details_exposed"`
	BlockedActions            []string                       `json:"blocked_actions"`
	UserFacingSettings        map[string]string              `json:"user_facing_settings"`
	DesktopSafeSummary        string                         `json:"desktop_safe_summary"`
}

type ExecutionResourcePreflight struct {
	SchemaVersion            string `json:"schema_version"`
	RequestType              string `json:"request_type"`
	PreflightType            string `json:"preflight_type"`
	RequestState             string `json:"request_state"`
	UserDecisionAllowsLaunch bool   `json:"user_decision_allows_launch"`
	PreflightPassed          bool   `json:"preflight_passed"`
	PortalRequired           bool   `json:"portal_required"`
	SnapshotRequired         bool   `json:"snapshot_required"`
	WriteGateOpen            bool   `json:"write_gate_open"`
	RuntimeLaunchApproval    bool   `json:"runtime_launch_approval"`
	PermissionGranted        bool   `json:"permission_granted"`
}

type ExecutionResourceGrant struct {
	ID                    string `json:"id"`
	Label                 string `json:"label"`
	Operation             string `json:"operation"`
	Decision              string `json:"decision"`
	GrantState            string `json:"grant_state"`
	PortalInterface       string `json:"portal_interface"`
	PortalRequired        bool   `json:"portal_required"`
	UserMediationRequired bool   `json:"user_mediation_required"`
	RequestObjectCreated  bool   `json:"request_object_created"`
	PermissionGranted     bool   `json:"permission_granted"`
	DirectAccessAllowed   bool   `json:"direct_access_allowed"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
	Summary               string `json:"summary"`
}

type ExecutionResourceBridgeSummary struct {
	RequestType            string `json:"request_type"`
	PlanType               string `json:"plan_type"`
	BridgeState            string `json:"bridge_state"`
	ResourceCount          int    `json:"resource_count"`
	PortalMediated         bool   `json:"portal_mediated"`
	ResourceBridgesEnabled bool   `json:"resource_bridges_enabled"`
	RequestObjectsCreated  bool   `json:"request_objects_created"`
	DirectHostFileAccess   bool   `json:"direct_host_file_access"`
}

func (plan Plan) ExecutionResourceGrantPreview(decision string, fileURIs []string) (ExecutionResourceGrantPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return ExecutionResourceGrantPreview{}, err
	}
	if !singleLine(decision) {
		return ExecutionResourceGrantPreview{}, errors.New("execution resource grant preview requires a single-line decision")
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return ExecutionResourceGrantPreview{}, errors.New("execution resource grant preview requires single-line identity fields")
		}
	}

	preflight, err := plan.ExecutionPreflightPreview(decision, fileURIs)
	if err != nil {
		return ExecutionResourceGrantPreview{}, err
	}
	permissionReview, err := plan.PermissionReviewPreview()
	if err != nil {
		return ExecutionResourceGrantPreview{}, err
	}
	resourceBridge, err := plan.DesktopResourceBridgePreview()
	if err != nil {
		return ExecutionResourceGrantPreview{}, err
	}

	grants, portalRequiredCount, pendingReviewCount := executionResourceGrants(permissionReview.Permissions)
	preview := ExecutionResourceGrantPreview{
		SchemaVersion:   "xnix.runtime.resource_grant.v1",
		RequestType:     "execution-resource-grant-preview",
		GrantType:       "compatibility-launch-resource-grant",
		RequestState:    preflight.RequestState,
		Source:          "execution-preflight-preview",
		Desktop:         "KDE Plasma",
		RuntimeMethod:   "Launch",
		ReadMethod:      "GetExecutionResourceGrantPreview",
		ApplicationID:   plan.ApplicationID,
		ApplicationName: plan.DisplayName,
		Icon:            plan.Icon,
		DesktopFile:     plan.DesktopFile,
		LauncherCommand: plan.LaunchCommand,
		ExecutionPreflight: ExecutionResourcePreflight{
			SchemaVersion:            preflight.SchemaVersion,
			RequestType:              preflight.RequestType,
			PreflightType:            preflight.PreflightType,
			RequestState:             preflight.RequestState,
			UserDecisionAllowsLaunch: preflight.UserDecisionAllowsLaunch,
			PreflightPassed:          preflight.PreflightPassed,
			PortalRequired:           preflight.PortalRequired,
			SnapshotRequired:         preflight.SnapshotRequired,
			WriteGateOpen:            preflight.WriteGateOpen,
			RuntimeLaunchApproval:    preflight.RuntimeLaunchApproval,
			PermissionGranted:        preflight.PermissionGranted,
		},
		ResourceGrants: grants,
		BridgeSummary: ExecutionResourceBridgeSummary{
			RequestType:            resourceBridge.RequestType,
			PlanType:               resourceBridge.PlanType,
			BridgeState:            resourceBridge.BridgeState,
			ResourceCount:          resourceBridge.ResourceCount,
			PortalMediated:         resourceBridge.PortalMediated,
			ResourceBridgesEnabled: resourceBridge.BridgesEnabled,
			RequestObjectsCreated:  resourceBridge.RequestsCreated,
			DirectHostFileAccess:   resourceBridge.DirectHostFileAccess,
		},
		GrantCount:                len(grants),
		AllowCount:                permissionReview.AllowCount,
		AskCount:                  permissionReview.AskCount,
		DenyCount:                 permissionReview.DenyCount,
		PortalRequiredCount:       portalRequiredCount,
		PendingReviewCount:        pendingReviewCount,
		FileCount:                 preflight.FileCount,
		FileURIs:                  preflight.FileURIs,
		RuntimeOwned:              true,
		GoRuntimeBacked:           true,
		KDEPolicyOwner:            false,
		CompatibilityCenterCard:   true,
		SafeForAIDiagnostics:      true,
		DesktopEntryLaunchVisible: true,
		LaunchIntentCaptured:      preflight.LaunchIntentCaptured,
		UserDecisionCaptured:      preflight.UserDecisionCaptured,
		UserDecisionAllowsLaunch:  preflight.UserDecisionAllowsLaunch,
		PreflightPassed:           false,
		PortalReviewRequired:      true,
		SnapshotRequired:          preflight.SnapshotRequired,
		GrantPlanCreated:          true,
		GrantObjectsCreated:       false,
		RequestObjectsCreated:     false,
		PermissionGranted:         false,
		ResourceBridgesEnabled:    false,
		SettingsPersisted:         false,
		RuntimeLaunchApproval:     false,
		LaunchAllowed:             false,
		LaunchEnabled:             false,
		ExecutionStarted:          false,
		HostPermissionChanged:     false,
		HostRootModified:          false,
		NetworkRequired:           false,
		BackendDetailsExposed:     false,
		BlockedActions:            []string{"create resource grant objects from preview", "enable desktop resource bridges before Portal review", "persist resource permissions before Runtime approval", "start compatibility profile from resource grant preview", "treat allowed network policy as launch approval", "mutate host root during resource grant planning", "expose raw backend command to desktop shell"},
		UserFacingSettings:        plan.UserFacingSettings,
		DesktopSafeSummary:        "KDE can show launch resource access before execution, but the Runtime does not grant permissions or start the application.",
	}
	if err := validateNoBackendTerms(preview, "execution resource grant preview"); err != nil {
		return ExecutionResourceGrantPreview{}, err
	}
	return preview, nil
}

func executionResourceGrants(permissions []PermissionReviewEntry) ([]ExecutionResourceGrant, int, int) {
	grants := make([]ExecutionResourceGrant, 0, len(permissions))
	portalRequiredCount := 0
	pendingReviewCount := 0
	for _, permission := range permissions {
		if permission.PortalRequired {
			portalRequiredCount++
		}
		if permission.Decision == "ask" {
			pendingReviewCount++
		}
		grants = append(grants, ExecutionResourceGrant{
			ID:                    permission.ID,
			Label:                 permission.Label,
			Operation:             permission.Operation,
			Decision:              permission.Decision,
			GrantState:            executionResourceGrantState(permission),
			PortalInterface:       permission.PortalInterface,
			PortalRequired:        permission.PortalRequired,
			UserMediationRequired: permission.UserMediationRequired,
			RequestObjectCreated:  false,
			PermissionGranted:     false,
			DirectAccessAllowed:   false,
			BackendDetailsExposed: false,
			Summary:               executionResourceGrantSummary(permission),
		})
	}
	return grants, portalRequiredCount, pendingReviewCount
}

func executionResourceGrantState(permission PermissionReviewEntry) string {
	switch permission.Decision {
	case "allow":
		return "planned-allow"
	case "deny":
		return "planned-deny"
	default:
		return "requires-review"
	}
}

func executionResourceGrantSummary(permission PermissionReviewEntry) string {
	switch permission.Decision {
	case "allow":
		return permission.Label + " access is planned but not granted until Runtime launch gates pass."
	case "deny":
		return permission.Label + " access remains denied for this launch."
	default:
		return permission.Label + " access requires user review before launch can proceed."
	}
}
