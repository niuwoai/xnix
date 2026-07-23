package appidentity

import (
	"errors"
	"path/filepath"
	"strings"
)

const (
	KnownAppKDERuntimeStatusLaunchActionTriggerSchemaVersion = "xnix.runtime.known_app_kde_runtime_status_launch_action_trigger.v1"
	KnownAppKDERuntimeStatusLaunchActionTriggerRequestType   = "known-app-kde-runtime-status-launch-action-trigger-preview"
)

type KnownAppKDERuntimeStatusLaunchActionTriggerRequest struct {
	StateRoot            string
	EvidenceID           string
	EvidenceRelativePath string
}

type KnownAppKDERuntimeStatusLaunchActionTriggerPreview struct {
	SchemaVersion               string                                `json:"schema_version"`
	RequestType                 string                                `json:"request_type"`
	Source                      string                                `json:"source"`
	Desktop                     string                                `json:"desktop"`
	RuntimeMethod               string                                `json:"runtime_method"`
	ReadMethod                  string                                `json:"read_method"`
	AppID                       string                                `json:"app_id"`
	DisplayName                 string                                `json:"display_name"`
	AppVersion                  string                                `json:"app_version"`
	ActionID                    string                                `json:"action_id"`
	ActionKind                  string                                `json:"action_kind"`
	TriggerState                string                                `json:"trigger_state"`
	EvidenceID                  string                                `json:"evidence_id"`
	EvidenceReadState           string                                `json:"evidence_read_state"`
	EvidenceHandoffConsumed     bool                                  `json:"evidence_handoff_consumed"`
	EvidenceRelativePath        string                                `json:"evidence_relative_path"`
	EvidenceSHA256              string                                `json:"evidence_sha256"`
	EvidenceDigestVerified      bool                                  `json:"evidence_digest_verified"`
	LaunchRequestType           string                                `json:"launch_request_type"`
	LaunchRequestRuntimeMethod  string                                `json:"launch_request_runtime_method"`
	LaunchRequestReadMethod     string                                `json:"launch_request_read_method"`
	LaunchRequestCreated        bool                                  `json:"launch_request_created"`
	LaunchRequest               KnownAppKDERuntimeStatusLaunchPreview `json:"launch_request"`
	ManagedLauncherArgv         []string                              `json:"managed_launcher_argv"`
	ManagedLauncherArgvReady    bool                                  `json:"managed_launcher_argv_ready"`
	RequiredOpaqueIDCount       int                                   `json:"required_opaque_id_count"`
	CollectedOpaqueIDCount      int                                   `json:"collected_opaque_id_count"`
	RuntimeOwnedTrigger         bool                                  `json:"runtime_owned_trigger"`
	RuntimeOwnedRequest         bool                                  `json:"runtime_owned_request"`
	RuntimeOwnedLaunch          bool                                  `json:"runtime_owned_launch"`
	RuntimeOwnedDispatch        bool                                  `json:"runtime_owned_dispatch"`
	KDEPresentationOnly         bool                                  `json:"kde_presentation_only"`
	KDEActionForwarded          bool                                  `json:"kde_action_forwarded"`
	StateRootRequired           bool                                  `json:"state_root_required"`
	StateRootSuppliedByRuntime  bool                                  `json:"state_root_supplied_by_runtime"`
	KDEStateRootAccess          bool                                  `json:"kde_state_root_access"`
	DirectLaunchEnabled         bool                                  `json:"direct_launch_enabled"`
	DesktopLaunchEnabled        bool                                  `json:"desktop_launch_enabled"`
	BackendLaunchEnabled        bool                                  `json:"backend_launch_enabled"`
	ExecutionStarted            bool                                  `json:"execution_started"`
	BackendProcessStarted       bool                                  `json:"backend_process_started"`
	RequestObjectsCreated       bool                                  `json:"request_objects_created"`
	PermissionGrantCreated      bool                                  `json:"permission_grant_created"`
	StateRootPathExposed        bool                                  `json:"state_root_path_exposed"`
	EvidencePathExposed         bool                                  `json:"evidence_path_exposed"`
	ReceiptPathExposed          bool                                  `json:"receipt_path_exposed"`
	SessionPathExposed          bool                                  `json:"session_path_exposed"`
	RawArtifactPathExposed      bool                                  `json:"raw_artifact_path_exposed"`
	RawCommandExposed           bool                                  `json:"raw_command_exposed"`
	RawLauncherOutputExposed    bool                                  `json:"raw_launcher_output_exposed"`
	BackendDetailsExposed       bool                                  `json:"backend_details_exposed"`
	HostRootModified            bool                                  `json:"host_root_modified"`
	PrivilegedContainerRequired bool                                  `json:"privileged_container_required"`
	DockerSocketMounted         bool                                  `json:"docker_socket_mounted"`
	BroadHostMountRequired      bool                                  `json:"broad_host_mount_required"`
	DesktopSafeSummary          string                                `json:"desktop_safe_summary"`
}

func PreviewKnownAppKDERuntimeStatusLaunchActionTrigger(request KnownAppKDERuntimeStatusLaunchActionTriggerRequest) (KnownAppKDERuntimeStatusLaunchActionTriggerPreview, error) {
	evidencePreview, err := PreviewKnownAppKDERuntimeStatusLaunchEvidence(KnownAppKDERuntimeStatusLaunchEvidencePreviewRequest{
		StateRoot:            request.StateRoot,
		EvidenceID:           request.EvidenceID,
		EvidenceRelativePath: request.EvidenceRelativePath,
	})
	if err != nil {
		return KnownAppKDERuntimeStatusLaunchActionTriggerPreview{}, err
	}
	evidence := evidencePreview.KnownAppSmokeEvidence
	if evidence.CenterCardState != "validated-post-review-dispatch" ||
		evidence.PrimaryActionID != KnownAppKDERuntimeStatusLaunchAction ||
		evidence.PrimaryActionKind != "runtime-status" ||
		!evidence.PostReviewDispatchConsumed {
		return KnownAppKDERuntimeStatusLaunchActionTriggerPreview{}, errors.New("known app KDE Runtime-status launch action trigger requires consumed Runtime-status handoff evidence")
	}
	launchRequest, err := PreviewKnownAppKDERuntimeStatusLaunchRequest(KnownAppKDERuntimeStatusLaunchRequest{
		AppID:                        evidence.AppID,
		LaunchAuthorizationReceiptID: evidence.LaunchAuthorizationReceiptID,
		SessionGatedReviewReceiptID:  evidence.SessionGatedReviewReceiptID,
		ControlledExecutionSessionID: evidence.ControlledExecutionSessionID,
		CenterCardState:              evidence.CenterCardState,
		PrimaryActionID:              evidence.PrimaryActionID,
		PostReviewDispatchState:      evidence.PostReviewDispatchState,
	})
	if err != nil {
		return KnownAppKDERuntimeStatusLaunchActionTriggerPreview{}, err
	}
	preview := KnownAppKDERuntimeStatusLaunchActionTriggerPreview{
		SchemaVersion:               KnownAppKDERuntimeStatusLaunchActionTriggerSchemaVersion,
		RequestType:                 KnownAppKDERuntimeStatusLaunchActionTriggerRequestType,
		Source:                      evidencePreview.RequestType + "+show-runtime-controlled-launch",
		Desktop:                     "KDE Plasma",
		RuntimeMethod:               "PreviewKnownAppKDERuntimeStatusLaunchActionTrigger",
		ReadMethod:                  "GetKnownAppKDERuntimeStatusLaunchActionTrigger",
		AppID:                       launchRequest.AppID,
		DisplayName:                 launchRequest.DisplayName,
		AppVersion:                  launchRequest.AppVersion,
		ActionID:                    KnownAppKDERuntimeStatusLaunchAction,
		ActionKind:                  "runtime-status",
		TriggerState:                "runtime-launch-request-assembled",
		EvidenceID:                  evidencePreview.EvidenceID,
		EvidenceReadState:           evidencePreview.EvidenceReadState,
		EvidenceHandoffConsumed:     evidencePreview.EvidenceHandoffConsumed,
		EvidenceRelativePath:        evidencePreview.EvidenceRelativePath,
		EvidenceSHA256:              evidencePreview.EvidenceSHA256,
		EvidenceDigestVerified:      evidencePreview.EvidenceDigestVerified,
		LaunchRequestType:           launchRequest.RequestType,
		LaunchRequestRuntimeMethod:  launchRequest.RuntimeMethod,
		LaunchRequestReadMethod:     launchRequest.ReadMethod,
		LaunchRequestCreated:        launchRequest.LaunchRequestCreated,
		LaunchRequest:               launchRequest,
		ManagedLauncherArgv:         launchRequest.ManagedLauncherArgv,
		ManagedLauncherArgvReady:    launchRequest.ManagedLauncherArgvReady,
		RequiredOpaqueIDCount:       launchRequest.RequiredOpaqueIDCount,
		CollectedOpaqueIDCount:      launchRequest.CollectedOpaqueIDCount,
		RuntimeOwnedTrigger:         true,
		RuntimeOwnedRequest:         launchRequest.RuntimeOwnedRequest,
		RuntimeOwnedLaunch:          launchRequest.RuntimeOwnedLaunch,
		RuntimeOwnedDispatch:        launchRequest.RuntimeOwnedDispatch,
		KDEPresentationOnly:         launchRequest.KDEPresentationOnly,
		KDEActionForwarded:          launchRequest.KDEActionForwarded,
		StateRootRequired:           launchRequest.StateRootRequired,
		StateRootSuppliedByRuntime:  launchRequest.StateRootSuppliedByRuntime,
		KDEStateRootAccess:          false,
		DirectLaunchEnabled:         false,
		DesktopLaunchEnabled:        false,
		BackendLaunchEnabled:        false,
		ExecutionStarted:            false,
		BackendProcessStarted:       false,
		RequestObjectsCreated:       false,
		PermissionGrantCreated:      false,
		StateRootPathExposed:        false,
		EvidencePathExposed:         false,
		ReceiptPathExposed:          false,
		SessionPathExposed:          false,
		RawArtifactPathExposed:      false,
		RawCommandExposed:           false,
		RawLauncherOutputExposed:    false,
		BackendDetailsExposed:       false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		DockerSocketMounted:         false,
		BroadHostMountRequired:      false,
		DesktopSafeSummary:          launchRequest.DisplayName + " Runtime-status action trigger assembled a Runtime-owned launch request from the consumed handoff.",
	}
	return validateKnownAppKDERuntimeStatusLaunchActionTriggerPreview(preview)
}

func validateKnownAppKDERuntimeStatusLaunchActionTriggerPreview(preview KnownAppKDERuntimeStatusLaunchActionTriggerPreview) (KnownAppKDERuntimeStatusLaunchActionTriggerPreview, error) {
	switch {
	case preview.SchemaVersion != KnownAppKDERuntimeStatusLaunchActionTriggerSchemaVersion || preview.RequestType != KnownAppKDERuntimeStatusLaunchActionTriggerRequestType:
		return KnownAppKDERuntimeStatusLaunchActionTriggerPreview{}, errors.New("known app KDE Runtime-status launch action trigger has invalid schema")
	case preview.ActionID != KnownAppKDERuntimeStatusLaunchAction || preview.ActionKind != "runtime-status" || preview.TriggerState != "runtime-launch-request-assembled":
		return KnownAppKDERuntimeStatusLaunchActionTriggerPreview{}, errors.New("known app KDE Runtime-status launch action trigger has invalid action state")
	case preview.EvidenceID == "" || preview.EvidenceReadState != "consumed" || !preview.EvidenceHandoffConsumed || preview.EvidenceRelativePath == "":
		return KnownAppKDERuntimeStatusLaunchActionTriggerPreview{}, errors.New("known app KDE Runtime-status launch action trigger requires consumed handoff evidence")
	case filepath.IsAbs(preview.EvidenceRelativePath) || strings.Contains(filepath.Clean(preview.EvidenceRelativePath), "..") || preview.EvidenceSHA256 == "" || !preview.EvidenceDigestVerified:
		return KnownAppKDERuntimeStatusLaunchActionTriggerPreview{}, errors.New("known app KDE Runtime-status launch action trigger requires safe verified evidence")
	case preview.LaunchRequestType != KnownAppKDERuntimeStatusLaunchRequestType || preview.LaunchRequestRuntimeMethod != "PreviewKnownAppKDERuntimeStatusLaunchRequest" || preview.LaunchRequestReadMethod != "GetKnownAppKDERuntimeStatusLaunchRequest":
		return KnownAppKDERuntimeStatusLaunchActionTriggerPreview{}, errors.New("known app KDE Runtime-status launch action trigger requires Runtime launch request preview")
	case !preview.LaunchRequestCreated || !preview.ManagedLauncherArgvReady || len(preview.ManagedLauncherArgv) != 11:
		return KnownAppKDERuntimeStatusLaunchActionTriggerPreview{}, errors.New("known app KDE Runtime-status launch action trigger requires managed launcher argv")
	case preview.RequiredOpaqueIDCount != 3 || preview.CollectedOpaqueIDCount != 3:
		return KnownAppKDERuntimeStatusLaunchActionTriggerPreview{}, errors.New("known app KDE Runtime-status launch action trigger requires collected opaque ids")
	case !preview.RuntimeOwnedTrigger || !preview.RuntimeOwnedRequest || !preview.RuntimeOwnedLaunch || !preview.RuntimeOwnedDispatch:
		return KnownAppKDERuntimeStatusLaunchActionTriggerPreview{}, errors.New("known app KDE Runtime-status launch action trigger requires Runtime ownership")
	case !preview.KDEPresentationOnly || !preview.KDEActionForwarded || !preview.StateRootRequired || !preview.StateRootSuppliedByRuntime || preview.KDEStateRootAccess:
		return KnownAppKDERuntimeStatusLaunchActionTriggerPreview{}, errors.New("known app KDE Runtime-status launch action trigger must keep KDE as a presentation shell")
	case preview.DirectLaunchEnabled || preview.DesktopLaunchEnabled || preview.BackendLaunchEnabled || preview.ExecutionStarted || preview.BackendProcessStarted:
		return KnownAppKDERuntimeStatusLaunchActionTriggerPreview{}, errors.New("known app KDE Runtime-status launch action trigger must not launch or start backend processes")
	case preview.RequestObjectsCreated || preview.PermissionGrantCreated:
		return KnownAppKDERuntimeStatusLaunchActionTriggerPreview{}, errors.New("known app KDE Runtime-status launch action trigger must not write request objects or grants")
	case preview.StateRootPathExposed || preview.EvidencePathExposed || preview.ReceiptPathExposed || preview.SessionPathExposed || preview.RawArtifactPathExposed || preview.RawCommandExposed || preview.RawLauncherOutputExposed || preview.BackendDetailsExposed:
		return KnownAppKDERuntimeStatusLaunchActionTriggerPreview{}, errors.New("known app KDE Runtime-status launch action trigger must not expose paths, raw commands, launcher output, or backend details")
	case preview.HostRootModified || preview.PrivilegedContainerRequired || preview.DockerSocketMounted || preview.BroadHostMountRequired:
		return KnownAppKDERuntimeStatusLaunchActionTriggerPreview{}, errors.New("known app KDE Runtime-status launch action trigger must keep host and container boundaries closed")
	}
	for _, value := range []string{preview.SchemaVersion, preview.RequestType, preview.Source, preview.Desktop, preview.RuntimeMethod, preview.ReadMethod, preview.AppID, preview.DisplayName, preview.AppVersion, preview.ActionID, preview.ActionKind, preview.TriggerState, preview.EvidenceID, preview.EvidenceReadState, preview.EvidenceRelativePath, preview.EvidenceSHA256, preview.LaunchRequestType, preview.LaunchRequestRuntimeMethod, preview.LaunchRequestReadMethod} {
		if value != "" && !singleLine(value) {
			return KnownAppKDERuntimeStatusLaunchActionTriggerPreview{}, errors.New("known app KDE Runtime-status launch action trigger requires single-line fields")
		}
	}
	if err := validateNoBackendTerms(preview, "known app KDE Runtime-status launch action trigger"); err != nil {
		return KnownAppKDERuntimeStatusLaunchActionTriggerPreview{}, err
	}
	return preview, nil
}
