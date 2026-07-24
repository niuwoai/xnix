package appidentity

import (
	"errors"
	"path/filepath"
	"strings"
)

const (
	ManagedLauncherAcceptanceReportSchemaVersion = "xnix.runtime.managed_launcher_acceptance_report.v1"
	ManagedLauncherAcceptanceReportRequestType   = "managed-launcher-acceptance-report-preview"
)

type ManagedLauncherAcceptanceReportRequest struct {
	StateRoot              string
	EvidenceID             string
	EvidenceRelativePath   string
	ExpectedEvidenceSHA256 string
	FullCheckpointPromoted bool
}

type ManagedLauncherAcceptanceReportSection struct {
	ID       string `json:"id"`
	State    string `json:"state"`
	Evidence string `json:"evidence"`
	Summary  string `json:"summary"`
}

type ManagedLauncherAcceptanceReportPreview struct {
	SchemaVersion                      string                                   `json:"schema_version"`
	RequestType                        string                                   `json:"request_type"`
	Source                             string                                   `json:"source"`
	RuntimeMethod                      string                                   `json:"runtime_method"`
	ReadMethod                         string                                   `json:"read_method"`
	ApplicationID                      string                                   `json:"application_id,omitempty"`
	ApplicationName                    string                                   `json:"application_name,omitempty"`
	ApplicationVersion                 string                                   `json:"application_version,omitempty"`
	AcceptanceState                    string                                   `json:"acceptance_state"`
	KnownAppIdentityState              string                                   `json:"known_app_identity_state"`
	ArtifactDigestState                string                                   `json:"artifact_digest_state"`
	LaunchAuthorizationState           string                                   `json:"launch_authorization_state"`
	SessionGatedReviewState            string                                   `json:"session_gated_review_state"`
	ControlledExecutionSessionState    string                                   `json:"controlled_execution_session_state"`
	ManagedLauncherRequestState        string                                   `json:"managed_launcher_request_state"`
	GuestSmokeBoundaryState            string                                   `json:"guest_smoke_boundary_state"`
	RuntimeStatusEvidenceState         string                                   `json:"runtime_status_evidence_state"`
	KDESafeProjectionState             string                                   `json:"kde_safe_projection_state"`
	FullCheckpointState                string                                   `json:"full_checkpoint_state"`
	DesktopTriggerReadinessState       string                                   `json:"desktop_trigger_readiness_state"`
	EvidenceID                         string                                   `json:"evidence_id,omitempty"`
	EvidenceRelativePath               string                                   `json:"evidence_relative_path,omitempty"`
	EvidenceSHA256                     string                                   `json:"evidence_sha256,omitempty"`
	ExpectedEvidenceSHA256             string                                   `json:"expected_evidence_sha256,omitempty"`
	EvidenceDigestVerified             bool                                     `json:"evidence_digest_verified"`
	ExpectedDigestMatched              bool                                     `json:"expected_digest_matched"`
	CompatibilityCenterProjectionReady bool                                     `json:"compatibility_center_projection_ready"`
	KDECenterProjectionReady           bool                                     `json:"kde_center_projection_ready"`
	ManagedLauncherArgvReady           bool                                     `json:"managed_launcher_argv_ready"`
	KnownAppAcceptanceReady            bool                                     `json:"known_app_acceptance_ready"`
	KnownAppSmokeStatus                string                                   `json:"known_app_smoke_status,omitempty"`
	KnownAppArtifactVerified           bool                                     `json:"known_app_artifact_verified"`
	KnownAppMarkerObserved             bool                                     `json:"known_app_marker_observed"`
	GuestBoundary                      string                                   `json:"guest_boundary,omitempty"`
	LaunchAuthorizationReceiptID       string                                   `json:"launch_authorization_receipt_id,omitempty"`
	SessionGatedReviewReceiptID        string                                   `json:"session_gated_review_receipt_id,omitempty"`
	ControlledExecutionSessionID       string                                   `json:"controlled_execution_session_id,omitempty"`
	RuntimeOwned                       bool                                     `json:"runtime_owned"`
	RuntimeOwnedDispatch               bool                                     `json:"runtime_owned_dispatch"`
	GoRuntimeBacked                    bool                                     `json:"go_runtime_backed"`
	KDEPresentationOnly                bool                                     `json:"kde_presentation_only"`
	KDEForwardsOnlyEvidenceHandle      bool                                     `json:"kde_forwards_only_evidence_handle"`
	StateRootPathExposed               bool                                     `json:"state_root_path_exposed"`
	EvidencePathExposed                bool                                     `json:"evidence_path_exposed"`
	ManagedLauncherPathExposed         bool                                     `json:"managed_launcher_path_exposed"`
	RawArtifactPathExposed             bool                                     `json:"raw_artifact_path_exposed"`
	RawCommandExposed                  bool                                     `json:"raw_command_exposed"`
	RawLauncherOutputExposed           bool                                     `json:"raw_launcher_output_exposed"`
	BackendDetailsExposed              bool                                     `json:"backend_details_exposed"`
	DesktopLaunchEnabled               bool                                     `json:"desktop_launch_enabled"`
	BackendLaunchEnabled               bool                                     `json:"backend_launch_enabled"`
	ExecutionStarted                   bool                                     `json:"execution_started"`
	BackendProcessStarted              bool                                     `json:"backend_process_started"`
	RuntimeStateWritten                bool                                     `json:"runtime_state_written"`
	KDEConfigurationWritten            bool                                     `json:"kde_configuration_written"`
	DBusCalled                         bool                                     `json:"dbus_called"`
	SmokeExecutedByPreview             bool                                     `json:"smoke_executed_by_preview"`
	NetworkRequired                    bool                                     `json:"network_required"`
	HostRootModified                   bool                                     `json:"host_root_modified"`
	DockerSocketMounted                bool                                     `json:"docker_socket_mounted"`
	BroadHostMountRequired             bool                                     `json:"broad_host_mount_required"`
	PrivilegedContainerRequired        bool                                     `json:"privileged_container_required"`
	HostNetworkRequired                bool                                     `json:"host_network_required"`
	FormalReleaseReady                 bool                                     `json:"formal_release_ready"`
	NextHumanAuthorizedSmoke           string                                   `json:"next_human_authorized_smoke"`
	NextHumanAuthorizedSmokeCommand    []string                                 `json:"next_human_authorized_smoke_command"`
	Sections                           []ManagedLauncherAcceptanceReportSection `json:"sections"`
	DesktopSafeSummary                 string                                   `json:"desktop_safe_summary"`
}

func PreviewManagedLauncherAcceptanceReport(request ManagedLauncherAcceptanceReportRequest) (ManagedLauncherAcceptanceReportPreview, error) {
	preview := baseManagedLauncherAcceptanceReport(request)
	evidence, err := PreviewKnownAppKDERuntimeStatusLaunchEvidence(KnownAppKDERuntimeStatusLaunchEvidencePreviewRequest{
		StateRoot:            request.StateRoot,
		EvidenceID:           request.EvidenceID,
		EvidenceRelativePath: request.EvidenceRelativePath,
	})
	if err != nil {
		state := classifyDesktopTriggerReadinessError(err)
		preview.AcceptanceState = state
		preview.RuntimeStatusEvidenceState = state
		preview.DesktopTriggerReadinessState = state
		preview.DesktopSafeSummary = "Managed launcher acceptance is blocked because Runtime-status launch evidence is unavailable or invalid."
		preview.Sections = append(preview.Sections, managedLauncherAcceptanceSection("runtime-status-evidence-handoff", state, "Runtime-status launch evidence handoff", preview.DesktopSafeSummary))
		return validateManagedLauncherAcceptanceReportPreview(preview)
	}

	preview.ApplicationID = evidence.AppID
	preview.ApplicationName = evidence.DisplayName
	preview.ApplicationVersion = evidence.AppVersion
	preview.EvidenceID = evidence.EvidenceID
	preview.EvidenceRelativePath = evidence.EvidenceRelativePath
	preview.EvidenceSHA256 = evidence.EvidenceSHA256
	preview.EvidenceDigestVerified = evidence.EvidenceDigestVerified
	preview.CompatibilityCenterProjectionReady = evidence.CompatibilityCenterProjectionReady
	preview.KDECenterProjectionReady = evidence.KDECenterProjectionReady
	preview.KnownAppSmokeStatus = evidence.KnownAppSmokeEvidence.SmokeStatus
	preview.KnownAppArtifactVerified = evidence.KnownAppSmokeEvidence.ChecksumVerified
	preview.KnownAppMarkerObserved = evidence.KnownAppSmokeEvidence.MarkerObserved
	preview.GuestBoundary = "managed-known-app-guest-smoke"
	preview.LaunchAuthorizationReceiptID = evidence.LaunchAuthorizationReceiptID
	preview.SessionGatedReviewReceiptID = evidence.SessionGatedReviewReceiptID
	preview.ControlledExecutionSessionID = evidence.ControlledExecutionSessionID
	preview.RuntimeOwned = evidence.RuntimeOwned
	preview.RuntimeOwnedDispatch = evidence.RuntimeOwnedDispatch
	preview.GoRuntimeBacked = evidence.GoRuntimeBacked
	preview.KDEPresentationOnly = !evidence.KDEPolicyOwner
	preview.RuntimeStatusEvidenceState = "ready"
	if strings.TrimSpace(request.ExpectedEvidenceSHA256) != "" && request.ExpectedEvidenceSHA256 != evidence.EvidenceSHA256 {
		preview.AcceptanceState = "stale-evidence"
		preview.RuntimeStatusEvidenceState = "stale-evidence"
		preview.ExpectedDigestMatched = false
		preview.DesktopTriggerReadinessState = "stale-evidence"
		preview.DesktopSafeSummary = evidence.DisplayName + " managed launcher acceptance is blocked because the caller's expected evidence digest is stale."
		preview.Sections = append(preview.Sections, managedLauncherAcceptanceSection("runtime-status-evidence-handoff", "stale-evidence", "Runtime-status launch evidence digest", "The evidence handoff is valid, but it does not match the caller's expected digest."))
		return validateManagedLauncherAcceptanceReportPreview(preview)
	}
	preview.ExpectedDigestMatched = true

	trigger, triggerErr := PreviewKnownAppKDERuntimeStatusLaunchActionTrigger(KnownAppKDERuntimeStatusLaunchActionTriggerRequest{
		StateRoot:            request.StateRoot,
		EvidenceID:           request.EvidenceID,
		EvidenceRelativePath: request.EvidenceRelativePath,
	})
	readiness, readinessErr := PreviewDesktopTriggerStagedInvocationReadiness(DesktopTriggerStagedInvocationReadinessRequest{
		StateRoot:              request.StateRoot,
		EvidenceID:             request.EvidenceID,
		EvidenceRelativePath:   request.EvidenceRelativePath,
		ExpectedEvidenceSHA256: request.ExpectedEvidenceSHA256,
		FullCheckpointPromoted: request.FullCheckpointPromoted,
	})

	smoke := evidence.KnownAppSmokeEvidence
	preview.KnownAppIdentityState = managedLauncherIdentityState(evidence)
	preview.ArtifactDigestState = managedLauncherArtifactState(smoke)
	preview.LaunchAuthorizationState = managedLauncherLaunchAuthorizationState(smoke, evidence)
	preview.SessionGatedReviewState = managedLauncherSessionGatedReviewState(smoke, evidence)
	preview.ControlledExecutionSessionState = managedLauncherControlledSessionState(smoke, evidence)
	preview.ManagedLauncherRequestState = stateFromError(triggerErr)
	preview.GuestSmokeBoundaryState = managedLauncherGuestSmokeState(smoke)
	preview.KDESafeProjectionState = managedLauncherProjectionState(evidence)
	if triggerErr == nil {
		preview.ManagedLauncherArgvReady = trigger.ManagedLauncherArgvReady
	}
	if readinessErr == nil {
		preview.DesktopTriggerReadinessState = readiness.ReadinessState
		preview.KDEForwardsOnlyEvidenceHandle = readiness.KDEForwardsOnlyEvidenceHandle
	} else {
		preview.DesktopTriggerReadinessState = stateFromError(readinessErr)
	}
	if readinessErr == nil && readiness.ReadinessState == "ready" {
		preview.KDEForwardsOnlyEvidenceHandle = readiness.KDEForwardsOnlyEvidenceHandle
	}
	preview.Sections = managedLauncherAcceptanceSections(preview)
	preview.AcceptanceState = managedLauncherOverallAcceptance(preview)
	preview.FormalReleaseReady = preview.AcceptanceState == "accepted"
	preview.KnownAppAcceptanceReady = preview.AcceptanceState == "accepted"
	preview.DesktopSafeSummary = managedLauncherAcceptanceSummary(preview)
	return validateManagedLauncherAcceptanceReportPreview(preview)
}

func baseManagedLauncherAcceptanceReport(request ManagedLauncherAcceptanceReportRequest) ManagedLauncherAcceptanceReportPreview {
	fullCheckpointState := "needs-full-checkpoint"
	if request.FullCheckpointPromoted {
		fullCheckpointState = "ready"
	}
	return ManagedLauncherAcceptanceReportPreview{
		SchemaVersion:                   ManagedLauncherAcceptanceReportSchemaVersion,
		RequestType:                     ManagedLauncherAcceptanceReportRequestType,
		Source:                          KnownAppKDERuntimeStatusLaunchEvidencePreviewRequestType + "+managed-launcher-acceptance",
		RuntimeMethod:                   "PreviewManagedLauncherAcceptanceReport",
		ReadMethod:                      "GetManagedLauncherAcceptanceReport",
		AcceptanceState:                 "blocked",
		KnownAppIdentityState:           "blocked",
		ArtifactDigestState:             "blocked",
		LaunchAuthorizationState:        "blocked",
		SessionGatedReviewState:         "blocked",
		ControlledExecutionSessionState: "blocked",
		ManagedLauncherRequestState:     "blocked",
		GuestSmokeBoundaryState:         "blocked",
		RuntimeStatusEvidenceState:      "blocked",
		KDESafeProjectionState:          "blocked",
		FullCheckpointState:             fullCheckpointState,
		DesktopTriggerReadinessState:    "blocked",
		ExpectedEvidenceSHA256:          strings.TrimSpace(request.ExpectedEvidenceSHA256),
		RuntimeOwned:                    true,
		RuntimeOwnedDispatch:            true,
		GoRuntimeBacked:                 true,
		KDEPresentationOnly:             true,
		KDEForwardsOnlyEvidenceHandle:   true,
		StateRootPathExposed:            false,
		EvidencePathExposed:             false,
		ManagedLauncherPathExposed:      false,
		RawArtifactPathExposed:          false,
		RawCommandExposed:               false,
		RawLauncherOutputExposed:        false,
		BackendDetailsExposed:           false,
		DesktopLaunchEnabled:            false,
		BackendLaunchEnabled:            false,
		ExecutionStarted:                false,
		BackendProcessStarted:           false,
		RuntimeStateWritten:             false,
		KDEConfigurationWritten:         false,
		DBusCalled:                      false,
		SmokeExecutedByPreview:          false,
		NetworkRequired:                 false,
		HostRootModified:                false,
		DockerSocketMounted:             false,
		BroadHostMountRequired:          false,
		PrivilegedContainerRequired:     false,
		HostNetworkRequired:             false,
		FormalReleaseReady:              false,
		NextHumanAuthorizedSmoke:        "desktop-triggered-staged-launch-smoke",
		NextHumanAuthorizedSmokeCommand: []string{"ruby", "scripts/staged_launcher_dispatch_smoke.rb"},
	}
}

func managedLauncherIdentityState(evidence KnownAppKDERuntimeStatusLaunchEvidencePreview) string {
	if evidence.AppID == "" || evidence.DisplayName == "" || evidence.AppVersion == "" {
		return "missing-evidence"
	}
	return "ready"
}

func managedLauncherArtifactState(smoke KnownAppSmokeEvidenceSummary) string {
	if !smoke.ChecksumVerified {
		return "missing-evidence"
	}
	return "ready"
}

func managedLauncherLaunchAuthorizationState(smoke KnownAppSmokeEvidenceSummary, evidence KnownAppKDERuntimeStatusLaunchEvidencePreview) string {
	if smoke.LaunchAuthorizationReceiptState != "recorded" || smoke.LaunchAuthorizationReceiptID == "" || evidence.LaunchAuthorizationReceiptID == "" {
		return "blocked"
	}
	return "ready"
}

func managedLauncherSessionGatedReviewState(smoke KnownAppSmokeEvidenceSummary, evidence KnownAppKDERuntimeStatusLaunchEvidencePreview) string {
	if !smoke.PostReviewDispatchConsumed || smoke.SessionGatedReviewReceiptID == "" || evidence.SessionGatedReviewReceiptID == "" || smoke.PostReviewDispatchState != "created-after-session-gated-review" {
		return "blocked"
	}
	return "ready"
}

func managedLauncherControlledSessionState(smoke KnownAppSmokeEvidenceSummary, evidence KnownAppKDERuntimeStatusLaunchEvidencePreview) string {
	if !smoke.LauncherSessionGateConsumed || !smoke.LauncherSessionDigestVerified || !smoke.LauncherSessionRuntimeOwnerConsumable || !smoke.LauncherSessionKDEReadModelConsumable || smoke.ControlledExecutionSessionID == "" || evidence.ControlledExecutionSessionID == "" {
		return "blocked"
	}
	return "ready"
}

func managedLauncherGuestSmokeState(smoke KnownAppSmokeEvidenceSummary) string {
	if !smoke.MarkerObserved || smoke.SmokeStatus != "passed" {
		return "blocked"
	}
	return "ready"
}

func managedLauncherProjectionState(evidence KnownAppKDERuntimeStatusLaunchEvidencePreview) string {
	if !evidence.CompatibilityCenterProjectionReady || !evidence.KDECenterProjectionReady {
		return "blocked"
	}
	return "ready"
}

func managedLauncherAcceptanceSections(preview ManagedLauncherAcceptanceReportPreview) []ManagedLauncherAcceptanceReportSection {
	return []ManagedLauncherAcceptanceReportSection{
		managedLauncherAcceptanceSection("known-app-identity", preview.KnownAppIdentityState, "Known application identity", "The report can identify the managed known Windows application without exposing local paths."),
		managedLauncherAcceptanceSection("artifact-digest", preview.ArtifactDigestState, "Known application artifact digest", "The known application evidence reports whether the managed artifact checksum was verified."),
		managedLauncherAcceptanceSection("launch-authorization", preview.LaunchAuthorizationState, "Runtime launch authorization receipt", "The Runtime launch authorization receipt is present as an opaque id."),
		managedLauncherAcceptanceSection("session-gated-review", preview.SessionGatedReviewState, "Session-gated review receipt", "The accepted review gate is present as an opaque id before controlled dispatch."),
		managedLauncherAcceptanceSection("controlled-execution-session", preview.ControlledExecutionSessionState, "Controlled execution session evidence", "The controlled execution session is digest-verified and consumable by Runtime and KDE read models."),
		managedLauncherAcceptanceSection("managed-launcher-request", preview.ManagedLauncherRequestState, "Runtime-managed launcher request", "The managed launcher argv can be assembled inside Runtime from the evidence handoff."),
		managedLauncherAcceptanceSection("guest-smoke-boundary", preview.GuestSmokeBoundaryState, "Managed known-app guest smoke", "The existing guest smoke boundary reports whether the known Windows app lane passed."),
		managedLauncherAcceptanceSection("runtime-status-evidence-handoff", preview.RuntimeStatusEvidenceState, "Runtime-status launch evidence handoff", "The persisted handoff is consumed and digest verified."),
		managedLauncherAcceptanceSection("kde-safe-projection", preview.KDESafeProjectionState, "Compatibility Center projection", "The Compatibility Center and KDE Center projections are ready without exposing Runtime internals."),
		managedLauncherAcceptanceSection("desktop-trigger-readiness", preview.DesktopTriggerReadinessState, "Desktop-trigger staged invocation readiness", "The desktop-trigger readiness packet agrees with this acceptance report."),
		managedLauncherAcceptanceSection("full-checkpoint", preview.FullCheckpointState, "Full checkpoint dependency", "Formal acceptance still depends on the human-authorized full checkpoint gate."),
	}
}

func managedLauncherAcceptanceSection(id string, state string, evidence string, summary string) ManagedLauncherAcceptanceReportSection {
	return ManagedLauncherAcceptanceReportSection{ID: id, State: state, Evidence: evidence, Summary: summary}
}

func managedLauncherOverallAcceptance(preview ManagedLauncherAcceptanceReportPreview) string {
	for _, state := range []string{
		preview.KnownAppIdentityState,
		preview.ArtifactDigestState,
		preview.LaunchAuthorizationState,
		preview.SessionGatedReviewState,
		preview.ControlledExecutionSessionState,
		preview.ManagedLauncherRequestState,
		preview.GuestSmokeBoundaryState,
		preview.RuntimeStatusEvidenceState,
		preview.KDESafeProjectionState,
		preview.DesktopTriggerReadinessState,
	} {
		if state != "ready" && state != "needs-full-checkpoint" {
			return state
		}
	}
	if preview.FullCheckpointState != "ready" || preview.DesktopTriggerReadinessState == "needs-full-checkpoint" {
		return "needs-full-checkpoint"
	}
	return "accepted"
}

func managedLauncherAcceptanceSummary(preview ManagedLauncherAcceptanceReportPreview) string {
	if preview.AcceptanceState == "accepted" {
		return preview.ApplicationName + " managed launcher evidence is accepted for the next human-authorized desktop-triggered staged smoke."
	}
	if preview.AcceptanceState == "needs-full-checkpoint" {
		return preview.ApplicationName + " managed launcher evidence is ready, but formal full checkpoint promotion is still required."
	}
	if preview.ArtifactDigestState != "ready" {
		return preview.ApplicationName + " managed launcher acceptance is blocked until artifact digest evidence is verified."
	}
	if preview.GuestSmokeBoundaryState != "ready" {
		return preview.ApplicationName + " managed launcher acceptance is blocked until the managed guest smoke boundary passes."
	}
	return preview.ApplicationName + " managed launcher acceptance is blocked by incomplete Runtime evidence."
}

func validateManagedLauncherAcceptanceReportPreview(preview ManagedLauncherAcceptanceReportPreview) (ManagedLauncherAcceptanceReportPreview, error) {
	if preview.SchemaVersion != ManagedLauncherAcceptanceReportSchemaVersion || preview.RequestType != ManagedLauncherAcceptanceReportRequestType {
		return ManagedLauncherAcceptanceReportPreview{}, errors.New("managed launcher acceptance report has invalid schema")
	}
	validStates := map[string]bool{
		"accepted": true, "ready": true, "blocked": true, "missing-evidence": true, "stale-evidence": true, "malformed": true, "unsafe": true, "unsupported": true, "needs-full-checkpoint": true,
	}
	for _, state := range []string{preview.AcceptanceState, preview.KnownAppIdentityState, preview.ArtifactDigestState, preview.LaunchAuthorizationState, preview.SessionGatedReviewState, preview.ControlledExecutionSessionState, preview.ManagedLauncherRequestState, preview.GuestSmokeBoundaryState, preview.RuntimeStatusEvidenceState, preview.KDESafeProjectionState, preview.FullCheckpointState, preview.DesktopTriggerReadinessState} {
		if !validStates[state] {
			return ManagedLauncherAcceptanceReportPreview{}, errors.New("managed launcher acceptance report has invalid state")
		}
	}
	if preview.AcceptanceState == "accepted" && (!preview.FormalReleaseReady || !preview.KnownAppAcceptanceReady) {
		return ManagedLauncherAcceptanceReportPreview{}, errors.New("managed launcher acceptance report cannot be accepted without release and known-app readiness")
	}
	if preview.AcceptanceState != "accepted" && (preview.FormalReleaseReady || preview.KnownAppAcceptanceReady) {
		return ManagedLauncherAcceptanceReportPreview{}, errors.New("managed launcher acceptance report cannot claim readiness while blocked")
	}
	if preview.EvidenceRelativePath != "" && (filepath.IsAbs(preview.EvidenceRelativePath) || strings.Contains(filepath.Clean(preview.EvidenceRelativePath), "..")) {
		return ManagedLauncherAcceptanceReportPreview{}, errors.New("managed launcher acceptance report requires safe relative evidence")
	}
	if preview.StateRootPathExposed || preview.EvidencePathExposed || preview.ManagedLauncherPathExposed || preview.RawArtifactPathExposed || preview.RawCommandExposed || preview.RawLauncherOutputExposed || preview.BackendDetailsExposed {
		return ManagedLauncherAcceptanceReportPreview{}, errors.New("managed launcher acceptance report must not expose Runtime internals")
	}
	if preview.DesktopLaunchEnabled || preview.BackendLaunchEnabled || preview.ExecutionStarted || preview.BackendProcessStarted || preview.RuntimeStateWritten || preview.KDEConfigurationWritten || preview.DBusCalled || preview.SmokeExecutedByPreview {
		return ManagedLauncherAcceptanceReportPreview{}, errors.New("managed launcher acceptance report preview must not execute or write")
	}
	if preview.NetworkRequired || preview.HostRootModified || preview.DockerSocketMounted || preview.BroadHostMountRequired || preview.PrivilegedContainerRequired || preview.HostNetworkRequired {
		return ManagedLauncherAcceptanceReportPreview{}, errors.New("managed launcher acceptance report must keep host and container boundaries closed")
	}
	for _, value := range []string{
		preview.SchemaVersion, preview.RequestType, preview.Source, preview.RuntimeMethod, preview.ReadMethod, preview.ApplicationID, preview.ApplicationName, preview.ApplicationVersion, preview.AcceptanceState, preview.KnownAppIdentityState, preview.ArtifactDigestState, preview.LaunchAuthorizationState, preview.SessionGatedReviewState, preview.ControlledExecutionSessionState, preview.ManagedLauncherRequestState, preview.GuestSmokeBoundaryState, preview.RuntimeStatusEvidenceState, preview.KDESafeProjectionState, preview.FullCheckpointState, preview.DesktopTriggerReadinessState, preview.EvidenceID, preview.EvidenceRelativePath, preview.EvidenceSHA256, preview.ExpectedEvidenceSHA256, preview.KnownAppSmokeStatus, preview.GuestBoundary, preview.LaunchAuthorizationReceiptID, preview.SessionGatedReviewReceiptID, preview.ControlledExecutionSessionID, preview.NextHumanAuthorizedSmoke, preview.DesktopSafeSummary,
	} {
		if value != "" && !singleLine(value) {
			return ManagedLauncherAcceptanceReportPreview{}, errors.New("managed launcher acceptance report requires single-line fields")
		}
	}
	for _, section := range preview.Sections {
		if section.ID == "" || !validStates[section.State] || section.Evidence == "" || section.Summary == "" {
			return ManagedLauncherAcceptanceReportPreview{}, errors.New("managed launcher acceptance report requires complete sections")
		}
		for _, value := range []string{section.ID, section.State, section.Evidence, section.Summary} {
			if !singleLine(value) {
				return ManagedLauncherAcceptanceReportPreview{}, errors.New("managed launcher acceptance report requires single-line section values")
			}
		}
	}
	for _, value := range preview.NextHumanAuthorizedSmokeCommand {
		if strings.TrimSpace(value) == "" || !singleLine(value) {
			return ManagedLauncherAcceptanceReportPreview{}, errors.New("managed launcher acceptance report requires single-line smoke command values")
		}
	}
	if err := validateNoBackendTerms(preview, "managed launcher acceptance report"); err != nil {
		return ManagedLauncherAcceptanceReportPreview{}, err
	}
	return preview, nil
}
