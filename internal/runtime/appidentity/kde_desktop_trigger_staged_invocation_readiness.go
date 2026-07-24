package appidentity

import (
	"errors"
	"path/filepath"
	"strings"
)

const (
	DesktopTriggerStagedInvocationReadinessSchemaVersion = "xnix.runtime.desktop_trigger_staged_invocation_readiness.v1"
	DesktopTriggerStagedInvocationReadinessRequestType   = "desktop-trigger-staged-invocation-readiness-preview"
)

type DesktopTriggerStagedInvocationReadinessRequest struct {
	StateRoot              string
	EvidenceID             string
	EvidenceRelativePath   string
	ExpectedEvidenceSHA256 string
	FullCheckpointPromoted bool
}

type DesktopTriggerStagedInvocationReadinessSection struct {
	ID       string `json:"id"`
	State    string `json:"state"`
	Evidence string `json:"evidence"`
	Summary  string `json:"summary"`
}

type DesktopTriggerStagedInvocationReadinessPreview struct {
	SchemaVersion                     string                                           `json:"schema_version"`
	RequestType                       string                                           `json:"request_type"`
	Source                            string                                           `json:"source"`
	RuntimeMethod                     string                                           `json:"runtime_method"`
	ReadMethod                        string                                           `json:"read_method"`
	ApplicationID                     string                                           `json:"application_id,omitempty"`
	ApplicationName                   string                                           `json:"application_name,omitempty"`
	ApplicationVersion                string                                           `json:"application_version,omitempty"`
	ReadinessState                    string                                           `json:"readiness_state"`
	ReleaseGateState                  string                                           `json:"release_gate_state"`
	KDEActionState                    string                                           `json:"kde_action_state"`
	PublicDBusRouteState              string                                           `json:"public_dbus_route_state"`
	OwnerServiceTriggerState          string                                           `json:"owner_service_trigger_state"`
	RuntimeStatusEvidenceState        string                                           `json:"runtime_status_evidence_state"`
	ManagedLauncherRequestState       string                                           `json:"managed_launcher_request_state"`
	KnownAppArtifactState             string                                           `json:"known_app_artifact_state"`
	GuestSmokeBoundaryState           string                                           `json:"guest_smoke_boundary_state"`
	EvidenceID                        string                                           `json:"evidence_id,omitempty"`
	EvidenceRelativePath              string                                           `json:"evidence_relative_path,omitempty"`
	EvidenceSHA256                    string                                           `json:"evidence_sha256,omitempty"`
	ExpectedEvidenceSHA256            string                                           `json:"expected_evidence_sha256,omitempty"`
	EvidenceDigestVerified            bool                                             `json:"evidence_digest_verified"`
	ExpectedDigestMatched             bool                                             `json:"expected_digest_matched"`
	KDEActionID                       string                                           `json:"kde_action_id,omitempty"`
	KDEActionLabel                    string                                           `json:"kde_action_label,omitempty"`
	PublicDBusMethod                  string                                           `json:"public_dbus_method,omitempty"`
	DesktopCallableRoute              string                                           `json:"desktop_callable_route,omitempty"`
	DesktopCallableRuntimeMethod      string                                           `json:"desktop_callable_runtime_method,omitempty"`
	DesktopCallableExecutionType      string                                           `json:"desktop_callable_execution_type,omitempty"`
	ManagedLauncherArgvReady          bool                                             `json:"managed_launcher_argv_ready"`
	KnownAppSmokeStatus               string                                           `json:"known_app_smoke_status,omitempty"`
	KnownAppArtifactVerified          bool                                             `json:"known_app_artifact_verified"`
	KnownAppMarkerObserved            bool                                             `json:"known_app_marker_observed"`
	GuestBoundary                     string                                           `json:"guest_boundary,omitempty"`
	RuntimeOwned                      bool                                             `json:"runtime_owned"`
	RuntimeOwnedDispatch              bool                                             `json:"runtime_owned_dispatch"`
	GoRuntimeBacked                   bool                                             `json:"go_runtime_backed"`
	KDEPresentationOnly               bool                                             `json:"kde_presentation_only"`
	KDEForwardsOnlyEvidenceHandle     bool                                             `json:"kde_forwards_only_evidence_handle"`
	OwnerServiceArgsExposedToKDE      bool                                             `json:"owner_service_args_exposed_to_kde"`
	DesktopReceiptFieldsReconstructed bool                                             `json:"desktop_receipt_fields_reconstructed"`
	DesktopKDEStateRootAccess         bool                                             `json:"desktop_kde_state_root_access"`
	StateRootPathExposed              bool                                             `json:"state_root_path_exposed"`
	EvidencePathExposed               bool                                             `json:"evidence_path_exposed"`
	ManagedLauncherPathExposed        bool                                             `json:"managed_launcher_path_exposed"`
	RawArtifactPathExposed            bool                                             `json:"raw_artifact_path_exposed"`
	RawCommandExposed                 bool                                             `json:"raw_command_exposed"`
	RawLauncherOutputExposed          bool                                             `json:"raw_launcher_output_exposed"`
	BackendDetailsExposed             bool                                             `json:"backend_details_exposed"`
	DesktopLaunchEnabled              bool                                             `json:"desktop_launch_enabled"`
	BackendLaunchEnabled              bool                                             `json:"backend_launch_enabled"`
	ExecutionStarted                  bool                                             `json:"execution_started"`
	BackendProcessStarted             bool                                             `json:"backend_process_started"`
	RuntimeStateWritten               bool                                             `json:"runtime_state_written"`
	KDEConfigurationWritten           bool                                             `json:"kde_configuration_written"`
	DBusCalled                        bool                                             `json:"dbus_called"`
	SmokeExecutedByPreview            bool                                             `json:"smoke_executed_by_preview"`
	NetworkFetchRequired              bool                                             `json:"network_fetch_required"`
	HostRootModified                  bool                                             `json:"host_root_modified"`
	DockerSocketMounted               bool                                             `json:"docker_socket_mounted"`
	BroadHostMountRequired            bool                                             `json:"broad_host_mount_required"`
	PrivilegedContainerRequired       bool                                             `json:"privileged_container_required"`
	HostNetworkRequired               bool                                             `json:"host_network_required"`
	FormalReleaseReady                bool                                             `json:"formal_release_ready"`
	NextHumanAuthorizedSmoke          string                                           `json:"next_human_authorized_smoke"`
	NextHumanAuthorizedSmokeCommand   []string                                         `json:"next_human_authorized_smoke_command"`
	Sections                          []DesktopTriggerStagedInvocationReadinessSection `json:"sections"`
	DesktopSafeSummary                string                                           `json:"desktop_safe_summary"`
}

func PreviewDesktopTriggerStagedInvocationReadiness(request DesktopTriggerStagedInvocationReadinessRequest) (DesktopTriggerStagedInvocationReadinessPreview, error) {
	preview := baseDesktopTriggerStagedInvocationReadiness(request)
	evidence, err := PreviewKnownAppKDERuntimeStatusLaunchEvidence(KnownAppKDERuntimeStatusLaunchEvidencePreviewRequest{
		StateRoot:            request.StateRoot,
		EvidenceID:           request.EvidenceID,
		EvidenceRelativePath: request.EvidenceRelativePath,
	})
	if err != nil {
		preview.ReadinessState = classifyDesktopTriggerReadinessError(err)
		preview.RuntimeStatusEvidenceState = preview.ReadinessState
		preview.DesktopSafeSummary = "Desktop-trigger staged invocation readiness is blocked because Runtime-status launch evidence is unavailable or invalid."
		preview.Sections = append(preview.Sections, desktopTriggerReadinessSection("runtime-status-evidence-handoff", preview.RuntimeStatusEvidenceState, "Runtime-status launch evidence handoff", preview.DesktopSafeSummary))
		return validateDesktopTriggerStagedInvocationReadinessPreview(preview)
	}
	preview.ApplicationID = evidence.AppID
	preview.ApplicationName = evidence.DisplayName
	preview.ApplicationVersion = evidence.AppVersion
	preview.EvidenceID = evidence.EvidenceID
	preview.EvidenceRelativePath = evidence.EvidenceRelativePath
	preview.EvidenceSHA256 = evidence.EvidenceSHA256
	preview.EvidenceDigestVerified = evidence.EvidenceDigestVerified
	preview.RuntimeOwned = evidence.RuntimeOwned
	preview.RuntimeOwnedDispatch = evidence.RuntimeOwnedDispatch
	preview.GoRuntimeBacked = evidence.GoRuntimeBacked
	preview.KDEPresentationOnly = !evidence.KDEPolicyOwner
	preview.KnownAppSmokeStatus = evidence.KnownAppSmokeEvidence.SmokeStatus
	preview.KnownAppArtifactVerified = evidence.KnownAppSmokeEvidence.ChecksumVerified
	preview.KnownAppMarkerObserved = evidence.KnownAppSmokeEvidence.MarkerObserved
	preview.GuestBoundary = "managed-known-app-guest-smoke"
	preview.RuntimeStatusEvidenceState = "ready"
	if strings.TrimSpace(request.ExpectedEvidenceSHA256) != "" && request.ExpectedEvidenceSHA256 != evidence.EvidenceSHA256 {
		preview.ReadinessState = "stale-evidence"
		preview.RuntimeStatusEvidenceState = "stale-evidence"
		preview.ExpectedDigestMatched = false
		preview.DesktopSafeSummary = evidence.DisplayName + " readiness is blocked because the caller's expected evidence digest is stale."
		preview.Sections = append(preview.Sections, desktopTriggerReadinessSection("runtime-status-evidence-handoff", "stale-evidence", "Runtime-status launch evidence digest", "The evidence handoff is valid, but it does not match the caller's expected digest."))
		return validateDesktopTriggerStagedInvocationReadinessPreview(preview)
	}
	preview.ExpectedDigestMatched = true

	_, ownerErr := PreviewKnownAppRuntimeStatusLaunchOwnerTrigger(KnownAppRuntimeStatusLaunchOwnerTriggerRequest{
		StateRoot:            request.StateRoot,
		EvidenceID:           request.EvidenceID,
		EvidenceRelativePath: request.EvidenceRelativePath,
	})
	action, actionErr := PreviewKDEControlledLaunchAction(KDEControlledLaunchActionRequest{
		StateRoot:            request.StateRoot,
		EvidenceID:           request.EvidenceID,
		EvidenceRelativePath: request.EvidenceRelativePath,
	})
	_, smokeErr := PreviewKDEControlledLaunchSessionBusSmokePlan(KDEControlledLaunchSessionBusSmokePlanRequest{
		StateRoot:            request.StateRoot,
		EvidenceID:           request.EvidenceID,
		EvidenceRelativePath: request.EvidenceRelativePath,
	})
	actionTrigger, triggerErr := PreviewKnownAppKDERuntimeStatusLaunchActionTrigger(KnownAppKDERuntimeStatusLaunchActionTriggerRequest{
		StateRoot:            request.StateRoot,
		EvidenceID:           request.EvidenceID,
		EvidenceRelativePath: request.EvidenceRelativePath,
	})

	preview.KDEActionState = stateFromError(actionErr)
	preview.PublicDBusRouteState = stateFromError(actionErr)
	preview.OwnerServiceTriggerState = stateFromError(ownerErr)
	preview.ManagedLauncherRequestState = stateFromError(triggerErr)
	preview.GuestSmokeBoundaryState = "ready"
	preview.KnownAppArtifactState = "ready"
	if !evidence.KnownAppSmokeEvidence.ChecksumVerified {
		preview.KnownAppArtifactState = "missing-evidence"
	}
	if !evidence.KnownAppSmokeEvidence.MarkerObserved || evidence.KnownAppSmokeEvidence.SmokeStatus != "passed" {
		preview.GuestSmokeBoundaryState = "blocked"
	}
	if actionErr == nil {
		preview.KDEActionID = action.KDEActionID
		preview.KDEActionLabel = action.KDEActionLabel
		preview.PublicDBusMethod = action.PublicDBusMethod
		preview.DesktopCallableRoute = action.DesktopCallableRoute
		preview.DesktopCallableRuntimeMethod = action.DesktopCallableRuntimeMethod
		preview.DesktopCallableExecutionType = action.DesktopCallableExecutionType
		preview.KDEForwardsOnlyEvidenceHandle = action.KDEForwardsOnlyEvidenceHandle
	}
	if ownerErr == nil {
		preview.OwnerServiceArgsExposedToKDE = false
	}
	if triggerErr == nil {
		preview.ManagedLauncherArgvReady = actionTrigger.ManagedLauncherArgvReady
	}
	if smokeErr != nil {
		preview.GuestSmokeBoundaryState = stateFromError(smokeErr)
	}
	preview.Sections = desktopTriggerReadinessSections(evidence, actionErr, smokeErr, ownerErr, triggerErr, preview.KnownAppArtifactState, preview.GuestSmokeBoundaryState, preview.ReleaseGateState)
	preview.ReadinessState = desktopTriggerOverallReadiness(preview)
	preview.FormalReleaseReady = preview.ReadinessState == "ready"
	preview.DesktopSafeSummary = desktopTriggerReadinessSummary(preview)
	return validateDesktopTriggerStagedInvocationReadinessPreview(preview)
}

func baseDesktopTriggerStagedInvocationReadiness(request DesktopTriggerStagedInvocationReadinessRequest) DesktopTriggerStagedInvocationReadinessPreview {
	releaseGateState := "needs-full-checkpoint"
	if request.FullCheckpointPromoted {
		releaseGateState = "ready"
	}
	return DesktopTriggerStagedInvocationReadinessPreview{
		SchemaVersion:                     DesktopTriggerStagedInvocationReadinessSchemaVersion,
		RequestType:                       DesktopTriggerStagedInvocationReadinessRequestType,
		Source:                            KDEControlledLaunchSessionBusSmokePlanRequestType + "+desktop-trigger-readiness",
		RuntimeMethod:                     "PreviewDesktopTriggerStagedInvocationReadiness",
		ReadMethod:                        "GetDesktopTriggerStagedInvocationReadiness",
		ReadinessState:                    "blocked",
		ReleaseGateState:                  releaseGateState,
		KDEActionState:                    "blocked",
		PublicDBusRouteState:              "blocked",
		OwnerServiceTriggerState:          "blocked",
		RuntimeStatusEvidenceState:        "blocked",
		ManagedLauncherRequestState:       "blocked",
		KnownAppArtifactState:             "blocked",
		GuestSmokeBoundaryState:           "blocked",
		ExpectedEvidenceSHA256:            strings.TrimSpace(request.ExpectedEvidenceSHA256),
		KDEPresentationOnly:               true,
		OwnerServiceArgsExposedToKDE:      false,
		DesktopReceiptFieldsReconstructed: false,
		DesktopKDEStateRootAccess:         false,
		StateRootPathExposed:              false,
		EvidencePathExposed:               false,
		ManagedLauncherPathExposed:        false,
		RawArtifactPathExposed:            false,
		RawCommandExposed:                 false,
		RawLauncherOutputExposed:          false,
		BackendDetailsExposed:             false,
		DesktopLaunchEnabled:              false,
		BackendLaunchEnabled:              false,
		ExecutionStarted:                  false,
		BackendProcessStarted:             false,
		RuntimeStateWritten:               false,
		KDEConfigurationWritten:           false,
		DBusCalled:                        false,
		SmokeExecutedByPreview:            false,
		NetworkFetchRequired:              false,
		HostRootModified:                  false,
		DockerSocketMounted:               false,
		BroadHostMountRequired:            false,
		PrivilegedContainerRequired:       false,
		HostNetworkRequired:               false,
		FormalReleaseReady:                false,
		NextHumanAuthorizedSmoke:          "desktop-triggered-staged-launch-smoke",
		NextHumanAuthorizedSmokeCommand:   []string{"ruby", "scripts/staged_launcher_dispatch_smoke.rb"},
	}
}

func desktopTriggerReadinessSections(evidence KnownAppKDERuntimeStatusLaunchEvidencePreview, actionErr error, smokeErr error, ownerErr error, triggerErr error, artifactState string, guestState string, releaseGateState string) []DesktopTriggerStagedInvocationReadinessSection {
	return []DesktopTriggerStagedInvocationReadinessSection{
		desktopTriggerReadinessSection("kde-action-metadata", stateFromError(actionErr), "KDE controlled-launch action preview", "KDE action metadata is presentation-only and forwards only the evidence handle."),
		desktopTriggerReadinessSection("public-dbus-route", stateFromError(actionErr), "Public desktop route metadata", "The desktop route is represented by a public Runtime D-Bus method without enabling production ownership."),
		desktopTriggerReadinessSection("owner-service-trigger", stateFromError(ownerErr), "Runtime owner trigger preview", "The Runtime owner service trigger is internally available without exposing owner-only arguments to KDE."),
		desktopTriggerReadinessSection("runtime-status-evidence-handoff", "ready", "Runtime-status launch evidence handoff", "The Runtime-status launch evidence handoff is consumed and digest verified."),
		desktopTriggerReadinessSection("managed-launcher-request", stateFromError(triggerErr), "Runtime-managed launcher request", "The managed launcher request can be assembled inside Runtime without KDE reconstructing receipts or sessions."),
		desktopTriggerReadinessSection("known-app-artifact-verification", artifactState, "Known application artifact evidence", "The known application smoke evidence reports whether the managed artifact was verified."),
		desktopTriggerReadinessSection("guest-smoke-boundary", guestState, evidence.KnownAppSmokeEvidence.EvidenceSource, "The existing staged smoke evidence reports whether the managed guest boundary completed."),
		desktopTriggerReadinessSection("release-gate-dependency", releaseGateState, "Full checkpoint gate", "The next desktop-triggered staged smoke must wait for formal full checkpoint promotion."),
	}
}

func desktopTriggerReadinessSection(id string, state string, evidence string, summary string) DesktopTriggerStagedInvocationReadinessSection {
	return DesktopTriggerStagedInvocationReadinessSection{ID: id, State: state, Evidence: evidence, Summary: summary}
}

func desktopTriggerOverallReadiness(preview DesktopTriggerStagedInvocationReadinessPreview) string {
	for _, state := range []string{
		preview.KDEActionState,
		preview.PublicDBusRouteState,
		preview.OwnerServiceTriggerState,
		preview.RuntimeStatusEvidenceState,
		preview.ManagedLauncherRequestState,
		preview.KnownAppArtifactState,
		preview.GuestSmokeBoundaryState,
	} {
		if state != "ready" {
			return state
		}
	}
	if preview.ReleaseGateState != "ready" {
		return preview.ReleaseGateState
	}
	return "ready"
}

func desktopTriggerReadinessSummary(preview DesktopTriggerStagedInvocationReadinessPreview) string {
	if preview.ReadinessState == "ready" {
		return preview.ApplicationName + " is ready for a human-authorized desktop-triggered staged launch smoke."
	}
	if preview.ReadinessState == "needs-full-checkpoint" {
		return preview.ApplicationName + " desktop-triggered staged launch readiness is prepared, but formal full checkpoint promotion is still required."
	}
	if preview.KnownAppArtifactState != "ready" {
		return preview.ApplicationName + " desktop-triggered staged launch readiness is blocked until known application artifact evidence is verified."
	}
	if preview.GuestSmokeBoundaryState != "ready" {
		return preview.ApplicationName + " desktop-triggered staged launch readiness is blocked until existing managed guest smoke evidence passes."
	}
	return preview.ApplicationName + " desktop-triggered staged launch readiness is blocked by incomplete Runtime evidence."
}

func stateFromError(err error) string {
	if err == nil {
		return "ready"
	}
	return classifyDesktopTriggerReadinessError(err)
}

func classifyDesktopTriggerReadinessError(err error) string {
	if err == nil {
		return "ready"
	}
	message := strings.ToLower(err.Error())
	switch {
	case strings.Contains(message, "no such file") || strings.Contains(message, "requires an evidence id") || strings.Contains(message, "requires a runtime-status launch evidence relative path"):
		return "missing-evidence"
	case strings.Contains(message, "invalid character") || strings.Contains(message, "unexpected end of json") || strings.Contains(message, "cannot unmarshal") || strings.Contains(message, "invalid schema") || strings.Contains(message, "invalid projection"):
		return "malformed"
	case strings.Contains(message, "unsafe") || strings.Contains(message, "exposes forbidden"):
		return "unsafe"
	case strings.Contains(message, "unsupported"):
		return "unsupported"
	default:
		return "blocked"
	}
}

func validateDesktopTriggerStagedInvocationReadinessPreview(preview DesktopTriggerStagedInvocationReadinessPreview) (DesktopTriggerStagedInvocationReadinessPreview, error) {
	if preview.SchemaVersion != DesktopTriggerStagedInvocationReadinessSchemaVersion || preview.RequestType != DesktopTriggerStagedInvocationReadinessRequestType {
		return DesktopTriggerStagedInvocationReadinessPreview{}, errors.New("desktop-trigger staged invocation readiness has invalid schema")
	}
	validStates := map[string]bool{
		"ready": true, "blocked": true, "missing-evidence": true, "stale-evidence": true, "unsafe": true, "unsupported": true, "needs-full-checkpoint": true, "malformed": true,
	}
	for _, state := range []string{preview.ReadinessState, preview.ReleaseGateState, preview.KDEActionState, preview.PublicDBusRouteState, preview.OwnerServiceTriggerState, preview.RuntimeStatusEvidenceState, preview.ManagedLauncherRequestState, preview.KnownAppArtifactState, preview.GuestSmokeBoundaryState} {
		if !validStates[state] {
			return DesktopTriggerStagedInvocationReadinessPreview{}, errors.New("desktop-trigger staged invocation readiness has invalid state")
		}
	}
	if preview.ReadinessState == "ready" && !preview.FormalReleaseReady {
		return DesktopTriggerStagedInvocationReadinessPreview{}, errors.New("desktop-trigger staged invocation readiness cannot be ready without formal release readiness")
	}
	if preview.ReadinessState != "ready" && preview.FormalReleaseReady {
		return DesktopTriggerStagedInvocationReadinessPreview{}, errors.New("desktop-trigger staged invocation readiness cannot claim formal release readiness while blocked")
	}
	if preview.EvidenceRelativePath != "" && (filepath.IsAbs(preview.EvidenceRelativePath) || strings.Contains(filepath.Clean(preview.EvidenceRelativePath), "..")) {
		return DesktopTriggerStagedInvocationReadinessPreview{}, errors.New("desktop-trigger staged invocation readiness requires safe relative evidence")
	}
	if preview.OwnerServiceArgsExposedToKDE || preview.DesktopReceiptFieldsReconstructed || preview.DesktopKDEStateRootAccess || preview.StateRootPathExposed || preview.EvidencePathExposed || preview.ManagedLauncherPathExposed || preview.RawArtifactPathExposed || preview.RawCommandExposed || preview.RawLauncherOutputExposed || preview.BackendDetailsExposed {
		return DesktopTriggerStagedInvocationReadinessPreview{}, errors.New("desktop-trigger staged invocation readiness must not expose Runtime internals")
	}
	if preview.DesktopLaunchEnabled || preview.BackendLaunchEnabled || preview.ExecutionStarted || preview.BackendProcessStarted || preview.RuntimeStateWritten || preview.KDEConfigurationWritten || preview.DBusCalled || preview.SmokeExecutedByPreview {
		return DesktopTriggerStagedInvocationReadinessPreview{}, errors.New("desktop-trigger staged invocation readiness preview must not execute or write")
	}
	if preview.NetworkFetchRequired || preview.HostRootModified || preview.DockerSocketMounted || preview.BroadHostMountRequired || preview.PrivilegedContainerRequired || preview.HostNetworkRequired {
		return DesktopTriggerStagedInvocationReadinessPreview{}, errors.New("desktop-trigger staged invocation readiness must keep host and container boundaries closed")
	}
	for _, value := range []string{
		preview.SchemaVersion, preview.RequestType, preview.Source, preview.RuntimeMethod, preview.ReadMethod, preview.ApplicationID, preview.ApplicationName, preview.ApplicationVersion, preview.ReadinessState, preview.ReleaseGateState, preview.KDEActionState, preview.PublicDBusRouteState, preview.OwnerServiceTriggerState, preview.RuntimeStatusEvidenceState, preview.ManagedLauncherRequestState, preview.KnownAppArtifactState, preview.GuestSmokeBoundaryState, preview.EvidenceID, preview.EvidenceRelativePath, preview.EvidenceSHA256, preview.ExpectedEvidenceSHA256, preview.KDEActionID, preview.KDEActionLabel, preview.PublicDBusMethod, preview.DesktopCallableRoute, preview.DesktopCallableRuntimeMethod, preview.DesktopCallableExecutionType, preview.KnownAppSmokeStatus, preview.GuestBoundary, preview.NextHumanAuthorizedSmoke, preview.DesktopSafeSummary,
	} {
		if value != "" && !singleLine(value) {
			return DesktopTriggerStagedInvocationReadinessPreview{}, errors.New("desktop-trigger staged invocation readiness requires single-line fields")
		}
	}
	for _, section := range preview.Sections {
		if section.ID == "" || !validStates[section.State] || section.Evidence == "" || section.Summary == "" {
			return DesktopTriggerStagedInvocationReadinessPreview{}, errors.New("desktop-trigger staged invocation readiness requires complete sections")
		}
		for _, value := range []string{section.ID, section.State, section.Evidence, section.Summary} {
			if !singleLine(value) {
				return DesktopTriggerStagedInvocationReadinessPreview{}, errors.New("desktop-trigger staged invocation readiness requires single-line section values")
			}
		}
	}
	for _, value := range preview.NextHumanAuthorizedSmokeCommand {
		if strings.TrimSpace(value) == "" || !singleLine(value) {
			return DesktopTriggerStagedInvocationReadinessPreview{}, errors.New("desktop-trigger staged invocation readiness requires single-line smoke command values")
		}
	}
	if err := validateNoBackendTerms(preview, "desktop-trigger staged invocation readiness"); err != nil {
		return DesktopTriggerStagedInvocationReadinessPreview{}, err
	}
	return preview, nil
}
