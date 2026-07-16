package appidentity

import (
	"sort"

	"xnix.local/xnix/internal/runtime/safety"
)

// CompatibilityBackendFallbackOptions carries the external, already-resolved
// evidence the command layer supplies (diagnostics risk counts). The read model
// itself joins the Runtime backend evidence.
type CompatibilityBackendFallbackOptions struct {
	DiagnosticsFailingRuns int
	DiagnosticsBlockedRuns int
}

type CompatibilityBackendFallbackPreview struct {
	SchemaVersion               string                                  `json:"schema_version"`
	RequestType                 string                                  `json:"request_type"`
	PreviewType                 string                                  `json:"preview_type"`
	Source                      string                                  `json:"source"`
	Desktop                     string                                  `json:"desktop"`
	RuntimeMethod               string                                  `json:"runtime_method"`
	ReadMethod                  string                                  `json:"read_method"`
	ApplicationID               string                                  `json:"application_id"`
	DisplayName                 string                                  `json:"display_name"`
	OverallState                string                                  `json:"overall_state"`
	PrimaryProfileID            string                                  `json:"primary_profile_id"`
	FallbackProfileID           string                                  `json:"fallback_profile_id"`
	SelectedProfileID           string                                  `json:"selected_profile_id"`
	FallbackAvailable           bool                                    `json:"fallback_available"`
	IsolationRequired           bool                                    `json:"isolation_required"`
	Candidates                  []CompatibilityBackendFallbackCandidate `json:"candidates"`
	CandidateIDs                []string                                `json:"candidate_ids"`
	CandidateCount              int                                     `json:"candidate_count"`
	ReasonCodes                 []string                                `json:"reason_codes"`
	RemediationHints            []string                                `json:"remediation_hints"`
	NextSafeReadOnlyChecks      []string                                `json:"next_safe_read_only_checks"`
	RuntimeOwned                bool                                    `json:"runtime_owned"`
	GoRuntimeBacked             bool                                    `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                                    `json:"kde_policy_owner"`
	UserVisible                 bool                                    `json:"user_visible"`
	ReviewOnly                  bool                                    `json:"review_only"`
	SelectionPersisted          bool                                    `json:"selection_persisted"`
	EngineInstallEnabled        bool                                    `json:"engine_install_enabled"`
	BackendLaunchEnabled        bool                                    `json:"backend_launch_enabled"`
	VMStartEnabled              bool                                    `json:"vm_start_enabled"`
	BackendDetailsExposed       bool                                    `json:"backend_details_exposed"`
	NetworkRequired             bool                                    `json:"network_required"`
	PrivilegedContainerRequired bool                                    `json:"privileged_container_required"`
	StateRootPathExposed        bool                                    `json:"state_root_path_exposed"`
	HostRootModified            bool                                    `json:"host_root_modified"`
	BlockedActions              []string                                `json:"blocked_actions"`
	DesktopSafeSummary          string                                  `json:"desktop_safe_summary"`
}

type CompatibilityBackendFallbackCandidate struct {
	ID                     string   `json:"id"`
	Label                  string   `json:"label"`
	Role                   string   `json:"role"`
	FallbackState          string   `json:"fallback_state"`
	ReasonCodes            []string `json:"reason_codes"`
	ReadyCapabilityCount   int      `json:"ready_capability_count"`
	PendingCapabilityCount int      `json:"pending_capability_count"`
	BlockedCapabilityCount int      `json:"blocked_capability_count"`
	LifecycleStatus        string   `json:"lifecycle_status"`
	PortalReviewRequired   bool     `json:"portal_review_required"`
	SnapshotRequired       bool     `json:"snapshot_required"`
	UserSafeSummary        string   `json:"user_safe_summary"`
	SelectionEnabled       bool     `json:"selection_enabled"`
	BackendLaunchEnabled   bool     `json:"backend_launch_enabled"`
	BackendDetailsExposed  bool     `json:"backend_details_exposed"`
}

// CompatibilityBackendFallbackPreview explains how Automatic mode would choose
// between Local compatibility and Isolated compatibility when evidence is
// missing, blocked, or risky. It joins the recipe-derived recommendation, the
// backend capability matrix, backend lifecycle status, and diagnostics risk
// without persisting any selection, installing engines, or starting a backend.
func (plan Plan) CompatibilityBackendFallbackPreview(options CompatibilityBackendFallbackOptions) (CompatibilityBackendFallbackPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return CompatibilityBackendFallbackPreview{}, err
	}

	primaryID := plan.recommendedCompatibilityProfileID()
	isolationRequired := primaryID == "isolated-compatibility"
	fallbackID := "isolated-compatibility"
	if primaryID == "isolated-compatibility" {
		fallbackID = "local-compatibility"
	}

	lifecycle, err := plan.BackendLifecyclePreview()
	if err != nil {
		return CompatibilityBackendFallbackPreview{}, err
	}
	matrix, err := NewBackendCapabilityMatrixPreview()
	if err != nil {
		return CompatibilityBackendFallbackPreview{}, err
	}
	capabilityByID := map[string]BackendCapabilityProfile{}
	for _, profile := range matrix.Profiles {
		capabilityByID[profile.ID] = profile
	}

	diagnosticsRisk := options.DiagnosticsFailingRuns > 0 || options.DiagnosticsBlockedRuns > 0
	lifecycleBlocked := backendFallbackLifecycleBlocked(lifecycle.OverallStatus)

	var candidates []CompatibilityBackendFallbackCandidate
	// Automatic orchestrator candidate is appended after the managed profiles are
	// evaluated so it can report which one would be selected.
	managedOrder := []string{"local-compatibility", "isolated-compatibility"}
	selectedProfileID := ""
	fallbackAvailable := false
	for _, id := range managedOrder {
		candidate := plan.backendFallbackManagedCandidate(id, primaryID, isolationRequired, lifecycle, lifecycleBlocked, diagnosticsRisk, capabilityByID[id])
		switch candidate.FallbackState {
		case "selected-primary":
			selectedProfileID = id
		case "available-fallback":
			fallbackAvailable = true
			if selectedProfileID == "" {
				selectedProfileID = id
			}
		}
		candidates = append(candidates, candidate)
	}

	automatic := backendFallbackAutomaticCandidate(primaryID, selectedProfileID, fallbackAvailable, isolationRequired)
	candidates = append([]CompatibilityBackendFallbackCandidate{automatic}, candidates...)

	overallState := backendFallbackOverallState(selectedProfileID, fallbackAvailable, isolationRequired)
	reasonCodes := backendFallbackAggregateReasons(candidates)

	preview := CompatibilityBackendFallbackPreview{
		SchemaVersion:     "xnix.runtime.compatibility_backend_fallback.v1",
		RequestType:       "compatibility-backend-fallback-preview",
		PreviewType:       "automatic-mode-backend-fallback",
		Source:            "backend-manager+backend-capability-matrix+backend-lifecycle+diagnostics",
		Desktop:           "KDE Plasma",
		RuntimeMethod:     "GetCompatibilityBackendFallback",
		ReadMethod:        "GetCompatibilityBackendFallbackPreview",
		ApplicationID:     plan.ApplicationID,
		DisplayName:       plan.DisplayName,
		OverallState:      overallState,
		PrimaryProfileID:  primaryID,
		FallbackProfileID: fallbackID,
		SelectedProfileID: selectedProfileID,
		FallbackAvailable: fallbackAvailable,
		IsolationRequired: isolationRequired,
		Candidates:        candidates,
		CandidateIDs:      backendFallbackCandidateIDs(candidates),
		CandidateCount:    len(candidates),
		ReasonCodes:       reasonCodes,
		RemediationHints:  backendFallbackRemediation(candidates, overallState),
		NextSafeReadOnlyChecks: []string{
			"review backend capability matrix readiness",
			"review backend lifecycle gates and repair hints",
			"review diagnostic run history for regressions",
		},
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		UserVisible:                 true,
		ReviewOnly:                  true,
		SelectionPersisted:          false,
		EngineInstallEnabled:        false,
		BackendLaunchEnabled:        false,
		VMStartEnabled:              false,
		BackendDetailsExposed:       false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		StateRootPathExposed:        false,
		HostRootModified:            false,
		BlockedActions: []string{
			"commit a compatibility profile selection from the fallback preview",
			"install a compatibility engine from the fallback preview",
			"start a compatibility backend from the fallback preview",
			"start an isolated compatibility environment from the fallback preview",
			"mutate host root during the fallback preview",
		},
		DesktopSafeSummary: backendFallbackSummary(overallState, primaryID),
	}
	if err := validateNoBackendTerms(preview, "compatibility backend fallback preview"); err != nil {
		return CompatibilityBackendFallbackPreview{}, err
	}
	if err := safety.ValidatePayload("compatibility backend fallback preview", preview); err != nil {
		return CompatibilityBackendFallbackPreview{}, err
	}
	return preview, nil
}

func (plan Plan) backendFallbackManagedCandidate(id, primaryID string, isolationRequired bool, lifecycle BackendLifecyclePreview, lifecycleBlocked, diagnosticsRisk bool, capability BackendCapabilityProfile) CompatibilityBackendFallbackCandidate {
	candidate := CompatibilityBackendFallbackCandidate{
		ID:                     id,
		Label:                  backendFallbackProfileLabel(id),
		ReadyCapabilityCount:   capability.ReadyCount,
		PendingCapabilityCount: capability.PendingCount,
		BlockedCapabilityCount: capability.BlockedCount,
		LifecycleStatus:        lifecycle.OverallStatus,
		PortalReviewRequired:   lifecycle.PortalReviewRequired,
		SnapshotRequired:       lifecycle.SnapshotRequired,
		SelectionEnabled:       false,
		BackendLaunchEnabled:   false,
		BackendDetailsExposed:  false,
	}
	isPrimary := id == primaryID
	if isPrimary {
		candidate.Role = "primary"
	} else {
		candidate.Role = "fallback"
	}

	var reasons []string
	// Isolation-only recipes forbid the local profile as a fallback.
	if isolationRequired && id == "local-compatibility" {
		candidate.FallbackState = "blocked-by-policy"
		reasons = append(reasons, "recipe-requests-isolation")
		candidate.ReasonCodes = backendFallbackUnique(reasons)
		candidate.UserSafeSummary = "Local compatibility is not offered because this application requests an isolated environment."
		return candidate
	}

	evidenceReady := capability.ProfileReady && !lifecycleBlocked && !diagnosticsRisk
	if evidenceReady {
		if isPrimary {
			candidate.FallbackState = "selected-primary"
		} else {
			candidate.FallbackState = "available-fallback"
		}
	} else {
		candidate.FallbackState = "blocked-missing-evidence"
	}

	if isPrimary {
		if isolationRequired {
			reasons = append(reasons, "recipe-requests-isolation")
		} else {
			reasons = append(reasons, "local-selected-by-default")
		}
	}
	if !capability.ProfileReady {
		reasons = append(reasons, "capability-not-ready")
	}
	if lifecycleBlocked {
		reasons = append(reasons, "lifecycle-"+lifecycle.OverallStatus)
	}
	if diagnosticsRisk {
		reasons = append(reasons, "diagnostics-regressions-present")
	}
	if lifecycle.PortalReviewRequired {
		reasons = append(reasons, "portal-review-required")
	}
	if lifecycle.SnapshotRequired {
		reasons = append(reasons, "snapshot-required")
	}
	candidate.ReasonCodes = backendFallbackUnique(reasons)
	candidate.UserSafeSummary = backendFallbackCandidateSummary(candidate.FallbackState, candidate.Label)
	return candidate
}

func backendFallbackAutomaticCandidate(primaryID, selectedProfileID string, fallbackAvailable, isolationRequired bool) CompatibilityBackendFallbackCandidate {
	candidate := CompatibilityBackendFallbackCandidate{
		ID:                    "automatic",
		Label:                 "Automatic",
		Role:                  "orchestrator",
		SelectionEnabled:      false,
		BackendLaunchEnabled:  false,
		BackendDetailsExposed: false,
	}
	var reasons []string
	switch {
	case selectedProfileID == primaryID && selectedProfileID != "":
		candidate.FallbackState = "selects-primary"
		reasons = append(reasons, "primary-evidence-ready")
	case fallbackAvailable && selectedProfileID != "":
		candidate.FallbackState = "falls-back"
		reasons = append(reasons, "primary-evidence-missing", "fallback-evidence-ready")
	case isolationRequired:
		candidate.FallbackState = "blocked"
		reasons = append(reasons, "recipe-requests-isolation", "no-selectable-profile")
	default:
		candidate.FallbackState = "blocked"
		reasons = append(reasons, "no-selectable-profile")
	}
	candidate.ReasonCodes = backendFallbackUnique(reasons)
	candidate.UserSafeSummary = backendFallbackAutomaticSummary(candidate.FallbackState)
	return candidate
}

func backendFallbackLifecycleBlocked(status string) bool {
	switch status {
	case "blocked", "repair-required", "not-ready":
		return true
	default:
		return false
	}
}

func backendFallbackOverallState(selectedProfileID string, fallbackAvailable, isolationRequired bool) string {
	switch {
	case selectedProfileID != "" && !fallbackAvailable:
		return "primary-selectable"
	case selectedProfileID != "" && fallbackAvailable:
		return "fallback-required"
	case isolationRequired:
		return "blocked-isolation-required"
	default:
		return "blocked-missing-evidence"
	}
}

func backendFallbackProfileLabel(id string) string {
	switch id {
	case "local-compatibility":
		return "Local compatibility"
	case "isolated-compatibility":
		return "Isolated compatibility"
	case "automatic":
		return "Automatic"
	default:
		return id
	}
}

func backendFallbackCandidateSummary(state, label string) string {
	switch state {
	case "selected-primary":
		return label + " would be selected as the primary compatibility mode after review gates pass."
	case "available-fallback":
		return label + " is available as a fallback compatibility mode after review gates pass."
	case "blocked-by-policy":
		return label + " is not offered for this application."
	default:
		return label + " cannot be selected yet because required Runtime evidence is missing."
	}
}

func backendFallbackAutomaticSummary(state string) string {
	switch state {
	case "selects-primary":
		return "Automatic mode would select the primary compatibility profile once review gates pass."
	case "falls-back":
		return "Automatic mode would fall back to the alternate compatibility profile because the primary lacks evidence."
	default:
		return "Automatic mode cannot select a compatibility profile yet because required Runtime evidence is missing."
	}
}

func backendFallbackSummary(overallState, primaryID string) string {
	switch overallState {
	case "primary-selectable":
		return "Automatic mode has enough evidence to prefer the primary compatibility profile; nothing has been selected or started."
	case "fallback-required":
		return "Automatic mode would fall back to the alternate compatibility profile; nothing has been selected or started."
	case "blocked-isolation-required":
		return "This application requests an isolated compatibility profile that is not ready yet; nothing has been selected or started."
	default:
		return "Neither compatibility profile has enough evidence yet, so Automatic mode stays blocked; nothing has been selected or started."
	}
}

func backendFallbackAggregateReasons(candidates []CompatibilityBackendFallbackCandidate) []string {
	var reasons []string
	for _, candidate := range candidates {
		reasons = append(reasons, candidate.ReasonCodes...)
	}
	return backendFallbackUnique(reasons)
}

func backendFallbackRemediation(candidates []CompatibilityBackendFallbackCandidate, overallState string) []string {
	hints := []string{}
	switch overallState {
	case "primary-selectable", "fallback-required":
		hints = append(hints, "Review the recommended compatibility profile in the Compatibility Center before enabling it.")
	default:
		hints = append(hints, "Complete the Runtime review gates before any compatibility profile can be selected.")
	}
	for _, candidate := range candidates {
		if candidate.Role == "orchestrator" {
			continue
		}
		if candidate.UserSafeSummary != "" {
			hints = append(hints, candidate.UserSafeSummary)
		}
	}
	return backendFallbackUnique(hints)
}

func backendFallbackCandidateIDs(candidates []CompatibilityBackendFallbackCandidate) []string {
	ids := make([]string, 0, len(candidates))
	for _, candidate := range candidates {
		ids = append(ids, candidate.ID)
	}
	return ids
}

func backendFallbackUnique(values []string) []string {
	seen := map[string]bool{}
	result := make([]string, 0, len(values))
	for _, value := range values {
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}
