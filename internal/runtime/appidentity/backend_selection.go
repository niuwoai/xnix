package appidentity

import (
	"errors"
)

type BackendSelectionPreview struct {
	SchemaVersion               string                    `json:"schema_version"`
	RequestType                 string                    `json:"request_type"`
	PlanType                    string                    `json:"plan_type"`
	Source                      string                    `json:"source"`
	Desktop                     string                    `json:"desktop"`
	RuntimeMethod               string                    `json:"runtime_method"`
	ApplicationID               string                    `json:"application_id"`
	DisplayName                 string                    `json:"display_name"`
	Icon                        string                    `json:"icon"`
	DesktopFile                 string                    `json:"desktop_file"`
	SelectedStrategy            string                    `json:"selected_strategy"`
	RecommendedProfileID        string                    `json:"recommended_profile_id"`
	CandidateProfiles           []BackendSelectionProfile `json:"candidate_profiles"`
	CandidateCount              int                       `json:"candidate_count"`
	ReadyCandidateCount         int                       `json:"ready_candidate_count"`
	BlockedCandidateCount       int                       `json:"blocked_candidate_count"`
	RequiredReviews             []string                  `json:"required_reviews"`
	BlockedActions              []string                  `json:"blocked_actions"`
	RuntimeOwned                bool                      `json:"runtime_owned"`
	GoRuntimeBacked             bool                      `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                      `json:"kde_policy_owner"`
	UserVisible                 bool                      `json:"user_visible"`
	SelectionCommitted          bool                      `json:"selection_committed"`
	SelectionChangeEnabled      bool                      `json:"selection_change_enabled"`
	BackendLaunchEnabled        bool                      `json:"backend_launch_enabled"`
	CapabilityActivationEnabled bool                      `json:"capability_activation_enabled"`
	EnvironmentCreated          bool                      `json:"environment_created"`
	RequestObjectCreated        bool                      `json:"request_object_created"`
	StateRootCreated            bool                      `json:"state_root_created"`
	SnapshotCreated             bool                      `json:"snapshot_created"`
	HostRootModified            bool                      `json:"host_root_modified"`
	PrivilegedContainerRequired bool                      `json:"privileged_container_required"`
	BackendDetailsExposed       bool                      `json:"backend_details_exposed"`
	UserFacingSettings          map[string]string         `json:"user_facing_settings"`
	DesktopSafeSummary          string                    `json:"desktop_safe_summary"`
}

type BackendSelectionProfile struct {
	ID                     string   `json:"id"`
	Label                  string   `json:"label"`
	Kind                   string   `json:"kind"`
	SelectionState         string   `json:"selection_state"`
	Recommended            bool     `json:"recommended"`
	Ready                  bool     `json:"ready"`
	Blocked                bool     `json:"blocked"`
	ReadyCapabilityCount   int      `json:"ready_capability_count"`
	PendingCapabilityCount int      `json:"pending_capability_count"`
	RequiredPreflight      []string `json:"required_preflight"`
	SelectionCommitted     bool     `json:"selection_committed"`
	EnvironmentCreated     bool     `json:"environment_created"`
	BackendProcessStarted  bool     `json:"backend_process_started"`
	BackendDetailsExposed  bool     `json:"backend_details_exposed"`
	Summary                string   `json:"summary"`
}

func (plan Plan) BackendSelectionPreview() (BackendSelectionPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return BackendSelectionPreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return BackendSelectionPreview{}, errors.New("backend selection preview requires single-line identity fields")
		}
	}

	recommendedProfileID := plan.recommendedCompatibilityProfileID()
	profiles := backendSelectionProfiles(recommendedProfileID)
	readyCandidateCount := 0
	blockedCandidateCount := 0
	for _, profile := range profiles {
		if profile.Ready {
			readyCandidateCount++
		}
		if profile.Blocked {
			blockedCandidateCount++
		}
	}

	preview := BackendSelectionPreview{
		SchemaVersion:               "xnix.runtime.backend_selection.v1",
		RequestType:                 "backend-selection-preview",
		PlanType:                    "compatibility-backend-selection-plan",
		Source:                      "compatibility-center",
		Desktop:                     "KDE Plasma",
		RuntimeMethod:               "GetBackendSelectionPlan",
		ApplicationID:               plan.ApplicationID,
		DisplayName:                 plan.DisplayName,
		Icon:                        plan.Icon,
		DesktopFile:                 plan.DesktopFile,
		SelectedStrategy:            recommendedProfileID + "-strategy",
		RecommendedProfileID:        recommendedProfileID,
		CandidateProfiles:           profiles,
		CandidateCount:              len(profiles),
		ReadyCandidateCount:         readyCandidateCount,
		BlockedCandidateCount:       blockedCandidateCount,
		RequiredReviews:             []string{"backend-capability-review", "backend-binding-review", "application-state-root-review", "portal-policy-review", "snapshot-baseline-review"},
		BlockedActions:              []string{"commit compatibility profile from KDE", "start compatibility profile after selection", "create profile environment before Runtime preflight", "expose backend implementation details to the desktop shell"},
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		UserVisible:                 true,
		SelectionCommitted:          false,
		SelectionChangeEnabled:      false,
		BackendLaunchEnabled:        false,
		CapabilityActivationEnabled: false,
		EnvironmentCreated:          false,
		RequestObjectCreated:        false,
		StateRootCreated:            false,
		SnapshotCreated:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		UserFacingSettings:          plan.UserFacingSettings,
		DesktopSafeSummary:          "Runtime can explain the recommended compatibility profile, but no profile selection is committed yet.",
	}
	if err := validateNoBackendTerms(preview, "backend selection preview"); err != nil {
		return BackendSelectionPreview{}, err
	}
	return preview, nil
}

func (plan Plan) recommendedCompatibilityProfileID() string {
	if plan.RecipeMode == "vm" || plan.UserFacingSettings["run_mode"] == "isolated-execution" {
		return "isolated-compatibility"
	}
	return "local-compatibility"
}

func backendSelectionProfiles(recommendedProfileID string) []BackendSelectionProfile {
	return []BackendSelectionProfile{
		backendSelectionProfile(
			"local-compatibility",
			"Local compatibility",
			"local",
			recommendedProfileID == "local-compatibility",
			2,
			4,
		),
		backendSelectionProfile(
			"isolated-compatibility",
			"Isolated compatibility",
			"isolated",
			recommendedProfileID == "isolated-compatibility",
			1,
			5,
		),
	}
}

func backendSelectionProfile(id string, label string, kind string, recommended bool, readyCount int, pendingCount int) BackendSelectionProfile {
	selectionState := "available-after-review"
	summaryState := "available"
	if recommended {
		selectionState = "recommended"
		summaryState = "recommended"
	}

	return BackendSelectionProfile{
		ID:                     id,
		Label:                  label,
		Kind:                   kind,
		SelectionState:         selectionState,
		Recommended:            recommended,
		Ready:                  false,
		Blocked:                true,
		ReadyCapabilityCount:   readyCount,
		PendingCapabilityCount: pendingCount,
		RequiredPreflight:      []string{"backend-capability-review", "backend-binding-review", "state-root-review", "portal-policy-review", "snapshot-baseline-review"},
		SelectionCommitted:     false,
		EnvironmentCreated:     false,
		BackendProcessStarted:  false,
		BackendDetailsExposed:  false,
		Summary:                label + " is " + summaryState + " only after Runtime preflight passes.",
	}
}
