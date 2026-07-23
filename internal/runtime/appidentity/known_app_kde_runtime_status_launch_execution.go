package appidentity

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"xnix.local/xnix/internal/runtime/winapp"
)

const (
	KnownAppKDERuntimeStatusLaunchExecutionSchemaVersion = "xnix.runtime.known_app_kde_runtime_status_launch_execution.v1"
	KnownAppKDERuntimeStatusLaunchExecutionRequestType   = "known-app-kde-runtime-status-launch-execution"
	knownAppRuntimeOwnedStateRootPlaceholder             = "<runtime-owned-state-root>"
)

type KnownAppKDERuntimeStatusLaunchExecutionRequest struct {
	AppID                        string
	StateRoot                    string
	CacheRoot                    string
	LaunchAuthorizationReceiptID string
	SessionGatedReviewReceiptID  string
	ControlledExecutionSessionID string
	CenterCardState              string
	PrimaryActionID              string
	PostReviewDispatchState      string
}

type KnownAppKDERuntimeStatusLaunchExecutionPlan struct {
	SchemaVersion                   string   `json:"schema_version"`
	RequestType                     string   `json:"request_type"`
	Source                          string   `json:"source"`
	Desktop                         string   `json:"desktop"`
	RuntimeMethod                   string   `json:"runtime_method"`
	ExecutionMethod                 string   `json:"execution_method"`
	RequestPreviewType              string   `json:"request_preview_type"`
	AppID                           string   `json:"app_id"`
	DisplayName                     string   `json:"display_name"`
	AppVersion                      string   `json:"app_version"`
	ActionID                        string   `json:"action_id"`
	ActionKind                      string   `json:"action_kind"`
	CenterCardState                 string   `json:"center_card_state"`
	PostReviewDispatchState         string   `json:"post_review_dispatch_state"`
	LaunchAuthorizationReceiptID    string   `json:"launch_authorization_receipt_id"`
	SessionGatedReviewReceiptID     string   `json:"session_gated_review_receipt_id"`
	ControlledExecutionSessionID    string   `json:"controlled_execution_session_id"`
	LaunchGateState                 string   `json:"launch_gate_state"`
	LaunchReceiptRevalidated        bool     `json:"launch_receipt_revalidated"`
	GuestBoundaryRevalidated        bool     `json:"guest_boundary_revalidated"`
	ReviewGateState                 string   `json:"review_gate_state"`
	ReviewReceiptRevalidated        bool     `json:"review_receipt_revalidated"`
	ControlledSessionRevalidated    bool     `json:"controlled_session_revalidated"`
	ControlledSessionDigestVerified bool     `json:"controlled_session_digest_verified"`
	ManagedLauncher                 string   `json:"managed_launcher"`
	ManagedLauncherArgv             []string `json:"managed_launcher_argv"`
	RuntimeManagedLauncherArgv      []string `json:"runtime_managed_launcher_argv"`
	RuntimeManagedLauncherArgvReady bool     `json:"runtime_managed_launcher_argv_ready"`
	RequiredOpaqueIDCount           int      `json:"required_opaque_id_count"`
	CollectedOpaqueIDCount          int      `json:"collected_opaque_id_count"`
	StateRootRequired               bool     `json:"state_root_required"`
	StateRootAccepted               bool     `json:"state_root_accepted"`
	StateRootInjectedByRuntime      bool     `json:"state_root_injected_by_runtime"`
	StateRootSuppliedByRuntime      bool     `json:"state_root_supplied_by_runtime"`
	KDEStateRootAccess              bool     `json:"kde_state_root_access"`
	ManagedLauncherInvocationReady  bool     `json:"managed_launcher_invocation_ready"`
	ExistingManagedLauncherPathUsed bool     `json:"existing_managed_launcher_path_used"`
	DelegatesArtifactGateToLauncher bool     `json:"delegates_artifact_gate_to_launcher"`
	DispatchReadyObserved           bool     `json:"dispatch_ready_observed"`
	DispatchPreparationRequired     bool     `json:"dispatch_preparation_required"`
	RuntimeOwnedRequest             bool     `json:"runtime_owned_request"`
	RuntimeOwnedLaunch              bool     `json:"runtime_owned_launch"`
	RuntimeOwnedDispatch            bool     `json:"runtime_owned_dispatch"`
	KDEPresentationOnly             bool     `json:"kde_presentation_only"`
	KDEActionForwarded              bool     `json:"kde_action_forwarded"`
	DirectLaunchEnabled             bool     `json:"direct_launch_enabled"`
	DesktopLaunchEnabled            bool     `json:"desktop_launch_enabled"`
	BackendLaunchEnabled            bool     `json:"backend_launch_enabled"`
	ExecutionStarted                bool     `json:"execution_started"`
	BackendProcessStarted           bool     `json:"backend_process_started"`
	RequestObjectsCreated           bool     `json:"request_objects_created"`
	PermissionGrantCreated          bool     `json:"permission_grant_created"`
	StateRootPathExposed            bool     `json:"state_root_path_exposed"`
	ReceiptPathExposed              bool     `json:"receipt_path_exposed"`
	SessionPathExposed              bool     `json:"session_path_exposed"`
	ManagedLauncherPathExposed      bool     `json:"managed_launcher_path_exposed"`
	RawArtifactPathExposed          bool     `json:"raw_artifact_path_exposed"`
	RawCommandExposed               bool     `json:"raw_command_exposed"`
	RawLauncherOutputExposed        bool     `json:"raw_launcher_output_exposed"`
	BackendDetailsExposed           bool     `json:"backend_details_exposed"`
	HostRootModified                bool     `json:"host_root_modified"`
	PrivilegedContainerRequired     bool     `json:"privileged_container_required"`
	DockerSocketMounted             bool     `json:"docker_socket_mounted"`
	BroadHostMountRequired          bool     `json:"broad_host_mount_required"`
	DesktopSafeSummary              string   `json:"desktop_safe_summary"`
}

type KnownAppKDERuntimeStatusLaunchDelegatedEvidenceRequest struct {
	AppID                                  string
	DisplayName                            string
	AppVersion                             string
	RequestType                            string
	Status                                 string
	SkipReason                             string
	FailureReason                          string
	GuestBoundary                          string
	RuntimeOwnedDispatch                   bool
	ArtifactVerified                       bool
	MarkerObserved                         bool
	SessionGatedControlledDispatchConsumed bool
	SessionGatedControlledDispatchState    string
	SessionGatedReviewReceiptID            string
	LaunchAuthorizationReceiptID           string
	ControlledExecutionSessionConsumed     bool
	ControlledExecutionSessionID           string
	ControlledSessionDigestVerified        bool
	ControlledSessionRelativePath          string
	RuntimeOwnerConsumableSession          bool
	KDEReadModelConsumableSession          bool
	ControlledSessionLiveStateObserved     bool
	ControlledSessionRegistered            bool
	ControlledSessionWindowObserved        bool
	ControlledSessionHostRootModified      bool
	ControlledSessionBackendProcessStart   bool
	HostRootModified                       bool
	DockerSocketMounted                    bool
	BroadHostMountRequired                 bool
}

type KnownAppKDERuntimeStatusLaunchDelegatedEvidence struct {
	ProjectionType                         string `json:"projection_type"`
	RuntimeMethod                          string `json:"runtime_method"`
	RequestType                            string `json:"request_type"`
	Status                                 string `json:"status"`
	SkipReason                             string `json:"skip_reason,omitempty"`
	FailureReason                          string `json:"failure_reason,omitempty"`
	AppID                                  string `json:"app_id"`
	DisplayName                            string `json:"display_name"`
	AppVersion                             string `json:"app_version"`
	GuestBoundary                          string `json:"guest_boundary"`
	RuntimeOwnedDispatch                   bool   `json:"runtime_owned_dispatch"`
	ArtifactVerified                       bool   `json:"artifact_verified"`
	MarkerObserved                         bool   `json:"marker_observed"`
	SessionGatedControlledDispatchConsumed bool   `json:"session_gated_controlled_dispatch_consumed"`
	SessionGatedControlledDispatchState    string `json:"session_gated_controlled_dispatch_state"`
	SessionGatedReviewReceiptID            string `json:"session_gated_review_receipt_id"`
	LaunchAuthorizationReceiptID           string `json:"launch_authorization_receipt_id"`
	ControlledExecutionSessionConsumed     bool   `json:"controlled_execution_session_consumed"`
	ControlledExecutionSessionID           string `json:"controlled_execution_session_id"`
	ControlledSessionDigestVerified        bool   `json:"controlled_session_digest_verified"`
	ControlledSessionRelativePath          string `json:"controlled_session_relative_path"`
	RuntimeOwnerConsumableSession          bool   `json:"runtime_owner_consumable_session"`
	KDEReadModelConsumableSession          bool   `json:"kde_read_model_consumable_session"`
	ControlledSessionLiveStateObserved     bool   `json:"controlled_session_live_state_observed"`
	ControlledSessionRegistered            bool   `json:"controlled_session_registered"`
	ControlledSessionWindowObserved        bool   `json:"controlled_session_window_observed"`
	ControlledSessionHostRootModified      bool   `json:"controlled_session_host_root_modified"`
	ControlledSessionBackendProcessStart   bool   `json:"controlled_session_backend_process_start"`
	HostRootModified                       bool   `json:"host_root_modified"`
	DockerSocketMounted                    bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired                 bool   `json:"broad_host_mount_required"`
	StateRootPathExposed                   bool   `json:"state_root_path_exposed"`
	ManagedLauncherPathExposed             bool   `json:"managed_launcher_path_exposed"`
	RawLauncherOutputExposed               bool   `json:"raw_launcher_output_exposed"`
	BackendDetailsExposed                  bool   `json:"backend_details_exposed"`
	CompatibilityCenterProjectionReady     bool   `json:"compatibility_center_projection_ready"`
	KDECenterProjectionReady               bool   `json:"kde_center_projection_ready"`
	DesktopSafeSummary                     string `json:"desktop_safe_summary"`
}

func PrepareKnownAppKDERuntimeStatusLaunchExecution(request KnownAppKDERuntimeStatusLaunchExecutionRequest) (KnownAppKDERuntimeStatusLaunchExecutionPlan, error) {
	stateRoot := strings.TrimSpace(request.StateRoot)
	if err := validateKnownAppRuntimeOwnerStateRoot(stateRoot); err != nil {
		return KnownAppKDERuntimeStatusLaunchExecutionPlan{}, err
	}
	requestPreview, err := PreviewKnownAppKDERuntimeStatusLaunchRequest(KnownAppKDERuntimeStatusLaunchRequest{
		AppID:                        request.AppID,
		LaunchAuthorizationReceiptID: request.LaunchAuthorizationReceiptID,
		SessionGatedReviewReceiptID:  request.SessionGatedReviewReceiptID,
		ControlledExecutionSessionID: request.ControlledExecutionSessionID,
		CenterCardState:              request.CenterCardState,
		PrimaryActionID:              request.PrimaryActionID,
		PostReviewDispatchState:      request.PostReviewDispatchState,
	})
	if err != nil {
		return KnownAppKDERuntimeStatusLaunchExecutionPlan{}, err
	}
	cacheRoot := strings.TrimSpace(request.CacheRoot)
	if cacheRoot == "" {
		cacheRoot = winapp.DefaultKnownAppCacheRoot
	}
	launchGate, err := PreviewKnownAppLaunchGate(KnownAppLaunchGateRequest{
		AppID:         requestPreview.AppID,
		StateRoot:     stateRoot,
		ReceiptID:     requestPreview.LaunchAuthorizationReceiptID,
		CacheRoot:     cacheRoot,
		GuestBoundary: winapp.KnownDispatchGuestBoundary,
	})
	if err != nil {
		return KnownAppKDERuntimeStatusLaunchExecutionPlan{}, err
	}
	if !launchGate.ReceiptAccepted || !launchGate.GuestBoundaryAccepted {
		return KnownAppKDERuntimeStatusLaunchExecutionPlan{}, errors.New("known app KDE Runtime-status launch execution requires an accepted launch receipt and managed guest boundary")
	}
	reviewGate, err := PreviewKnownAppSessionGatedLaunchReviewGate(KnownAppSessionGatedLaunchReviewGateRequest{
		AppID:     requestPreview.AppID,
		StateRoot: stateRoot,
		SessionID: requestPreview.ControlledExecutionSessionID,
		ReceiptID: requestPreview.SessionGatedReviewReceiptID,
	})
	if err != nil {
		return KnownAppKDERuntimeStatusLaunchExecutionPlan{}, err
	}
	if !reviewGate.ReviewReceiptAccepted || !reviewGate.DispatchStateAdvanceReady {
		return KnownAppKDERuntimeStatusLaunchExecutionPlan{}, errors.New("known app KDE Runtime-status launch execution requires an accepted session-gated review receipt")
	}
	sessionConsume, err := PreviewKnownAppControlledExecutionSessionConsumption(KnownAppControlledExecutionSessionConsumeRequest{
		AppID:     requestPreview.AppID,
		StateRoot: stateRoot,
		SessionID: requestPreview.ControlledExecutionSessionID,
	})
	if err != nil {
		return KnownAppKDERuntimeStatusLaunchExecutionPlan{}, err
	}
	if !sessionConsume.RecordConsumed || !sessionConsume.SessionDigestVerified {
		return KnownAppKDERuntimeStatusLaunchExecutionPlan{}, errors.New("known app KDE Runtime-status launch execution requires digest-verified controlled session evidence")
	}
	runtimeArgv := knownAppRuntimeStatusLaunchRuntimeArgv(requestPreview)
	plan := KnownAppKDERuntimeStatusLaunchExecutionPlan{
		SchemaVersion:                   KnownAppKDERuntimeStatusLaunchExecutionSchemaVersion,
		RequestType:                     KnownAppKDERuntimeStatusLaunchExecutionRequestType,
		Source:                          KnownAppKDERuntimeStatusLaunchRequestType + "+runtime-owner-execution-entrypoint",
		Desktop:                         "KDE Plasma",
		RuntimeMethod:                   "PrepareKnownAppKDERuntimeStatusLaunchExecution",
		ExecutionMethod:                 "RunKnownAppKDERuntimeStatusLaunchExecution",
		RequestPreviewType:              requestPreview.RequestType,
		AppID:                           requestPreview.AppID,
		DisplayName:                     requestPreview.DisplayName,
		AppVersion:                      requestPreview.AppVersion,
		ActionID:                        requestPreview.ActionID,
		ActionKind:                      requestPreview.ActionKind,
		CenterCardState:                 requestPreview.CenterCardState,
		PostReviewDispatchState:         requestPreview.PostReviewDispatchState,
		LaunchAuthorizationReceiptID:    requestPreview.LaunchAuthorizationReceiptID,
		SessionGatedReviewReceiptID:     requestPreview.SessionGatedReviewReceiptID,
		ControlledExecutionSessionID:    requestPreview.ControlledExecutionSessionID,
		LaunchGateState:                 launchGate.LaunchGateState,
		LaunchReceiptRevalidated:        true,
		GuestBoundaryRevalidated:        true,
		ReviewGateState:                 reviewGate.ReviewGateState,
		ReviewReceiptRevalidated:        true,
		ControlledSessionRevalidated:    true,
		ControlledSessionDigestVerified: true,
		ManagedLauncher:                 requestPreview.ManagedLauncher,
		ManagedLauncherArgv:             requestPreview.ManagedLauncherArgv,
		RuntimeManagedLauncherArgv:      runtimeArgv,
		RuntimeManagedLauncherArgvReady: true,
		RequiredOpaqueIDCount:           requestPreview.RequiredOpaqueIDCount,
		CollectedOpaqueIDCount:          requestPreview.CollectedOpaqueIDCount,
		StateRootRequired:               true,
		StateRootAccepted:               true,
		StateRootInjectedByRuntime:      true,
		StateRootSuppliedByRuntime:      true,
		KDEStateRootAccess:              false,
		ManagedLauncherInvocationReady:  true,
		ExistingManagedLauncherPathUsed: true,
		DelegatesArtifactGateToLauncher: true,
		DispatchReadyObserved:           launchGate.DispatchReady,
		DispatchPreparationRequired:     !launchGate.DispatchReady,
		RuntimeOwnedRequest:             true,
		RuntimeOwnedLaunch:              true,
		RuntimeOwnedDispatch:            true,
		KDEPresentationOnly:             true,
		KDEActionForwarded:              true,
		DirectLaunchEnabled:             false,
		DesktopLaunchEnabled:            false,
		BackendLaunchEnabled:            false,
		ExecutionStarted:                false,
		BackendProcessStarted:           false,
		RequestObjectsCreated:           false,
		PermissionGrantCreated:          false,
		StateRootPathExposed:            false,
		ReceiptPathExposed:              false,
		SessionPathExposed:              false,
		ManagedLauncherPathExposed:      false,
		RawArtifactPathExposed:          false,
		RawCommandExposed:               false,
		RawLauncherOutputExposed:        false,
		BackendDetailsExposed:           false,
		HostRootModified:                false,
		PrivilegedContainerRequired:     false,
		DockerSocketMounted:             false,
		BroadHostMountRequired:          false,
		DesktopSafeSummary:              requestPreview.DisplayName + " passed Runtime-owner receipt, session, post-review, and managed-boundary checks; the existing managed launcher is ready to receive a Runtime-injected state root.",
	}
	return validateKnownAppKDERuntimeStatusLaunchExecutionPlan(plan)
}

func ProjectKnownAppKDERuntimeStatusLaunchDelegatedEvidence(request KnownAppKDERuntimeStatusLaunchDelegatedEvidenceRequest) (KnownAppKDERuntimeStatusLaunchDelegatedEvidence, error) {
	projection := KnownAppKDERuntimeStatusLaunchDelegatedEvidence{
		ProjectionType:                         "known-app-kde-runtime-status-launch-delegated-evidence",
		RuntimeMethod:                          "ProjectKnownAppKDERuntimeStatusLaunchDelegatedEvidence",
		RequestType:                            strings.TrimSpace(request.RequestType),
		Status:                                 strings.TrimSpace(request.Status),
		SkipReason:                             strings.TrimSpace(request.SkipReason),
		FailureReason:                          strings.TrimSpace(request.FailureReason),
		AppID:                                  strings.TrimSpace(request.AppID),
		DisplayName:                            strings.TrimSpace(request.DisplayName),
		AppVersion:                             strings.TrimSpace(request.AppVersion),
		GuestBoundary:                          strings.TrimSpace(request.GuestBoundary),
		RuntimeOwnedDispatch:                   request.RuntimeOwnedDispatch,
		ArtifactVerified:                       request.ArtifactVerified,
		MarkerObserved:                         request.MarkerObserved,
		SessionGatedControlledDispatchConsumed: request.SessionGatedControlledDispatchConsumed,
		SessionGatedControlledDispatchState:    strings.TrimSpace(request.SessionGatedControlledDispatchState),
		SessionGatedReviewReceiptID:            strings.TrimSpace(request.SessionGatedReviewReceiptID),
		LaunchAuthorizationReceiptID:           strings.TrimSpace(request.LaunchAuthorizationReceiptID),
		ControlledExecutionSessionConsumed:     request.ControlledExecutionSessionConsumed,
		ControlledExecutionSessionID:           strings.TrimSpace(request.ControlledExecutionSessionID),
		ControlledSessionDigestVerified:        request.ControlledSessionDigestVerified,
		ControlledSessionRelativePath:          filepath.ToSlash(strings.TrimSpace(request.ControlledSessionRelativePath)),
		RuntimeOwnerConsumableSession:          request.RuntimeOwnerConsumableSession,
		KDEReadModelConsumableSession:          request.KDEReadModelConsumableSession,
		ControlledSessionLiveStateObserved:     request.ControlledSessionLiveStateObserved,
		ControlledSessionRegistered:            request.ControlledSessionRegistered,
		ControlledSessionWindowObserved:        request.ControlledSessionWindowObserved,
		ControlledSessionHostRootModified:      request.ControlledSessionHostRootModified,
		ControlledSessionBackendProcessStart:   request.ControlledSessionBackendProcessStart,
		HostRootModified:                       request.HostRootModified,
		DockerSocketMounted:                    request.DockerSocketMounted,
		BroadHostMountRequired:                 request.BroadHostMountRequired,
		StateRootPathExposed:                   false,
		ManagedLauncherPathExposed:             false,
		RawLauncherOutputExposed:               false,
		BackendDetailsExposed:                  false,
		CompatibilityCenterProjectionReady:     true,
		KDECenterProjectionReady:               true,
		DesktopSafeSummary:                     strings.TrimSpace(request.DisplayName) + " delegated Runtime launcher evidence is projected for Compatibility Center and KDE Center consumption without exposing launcher output or state-root paths.",
	}
	return validateKnownAppKDERuntimeStatusLaunchDelegatedEvidence(projection)
}

func KnownAppKDERuntimeStatusLaunchExecutionArgv(plan KnownAppKDERuntimeStatusLaunchExecutionPlan, stateRoot string, cacheRoot string, extraArgs []string) ([]string, error) {
	if _, err := validateKnownAppKDERuntimeStatusLaunchExecutionPlan(plan); err != nil {
		return nil, err
	}
	stateRoot = strings.TrimSpace(stateRoot)
	if err := validateKnownAppRuntimeOwnerStateRoot(stateRoot); err != nil {
		return nil, err
	}
	cacheRoot = strings.TrimSpace(cacheRoot)
	if cacheRoot == "" {
		cacheRoot = winapp.DefaultKnownAppCacheRoot
	}
	argv := []string{
		"--app", plan.AppID,
		"--cache-root", cacheRoot,
		"--guest-boundary", winapp.KnownDispatchGuestBoundary,
		"--state-root", stateRoot,
		"--receipt-id", plan.LaunchAuthorizationReceiptID,
		"--review-receipt-id", plan.SessionGatedReviewReceiptID,
		"--session-id", plan.ControlledExecutionSessionID,
	}
	for _, value := range extraArgs {
		if !singleLine(value) {
			return nil, errors.New("known app KDE Runtime-status launch execution requires single-line launcher arguments")
		}
		argv = append(argv, value)
	}
	return argv, nil
}

func validateKnownAppKDERuntimeStatusLaunchDelegatedEvidence(projection KnownAppKDERuntimeStatusLaunchDelegatedEvidence) (KnownAppKDERuntimeStatusLaunchDelegatedEvidence, error) {
	switch {
	case projection.ProjectionType != "known-app-kde-runtime-status-launch-delegated-evidence":
		return KnownAppKDERuntimeStatusLaunchDelegatedEvidence{}, errors.New("known app KDE Runtime-status delegated evidence has invalid projection type")
	case projection.RequestType != winapp.KnownDispatchSmokeRequestType:
		return KnownAppKDERuntimeStatusLaunchDelegatedEvidence{}, errors.New("known app KDE Runtime-status delegated evidence requires dispatch smoke evidence")
	case projection.AppID == "" || projection.DisplayName == "" || projection.AppVersion == "":
		return KnownAppKDERuntimeStatusLaunchDelegatedEvidence{}, errors.New("known app KDE Runtime-status delegated evidence requires known app identity")
	case projection.GuestBoundary != winapp.KnownDispatchGuestBoundary:
		return KnownAppKDERuntimeStatusLaunchDelegatedEvidence{}, errors.New("known app KDE Runtime-status delegated evidence requires managed guest boundary")
	case projection.ControlledSessionRelativePath != "" && (filepath.IsAbs(projection.ControlledSessionRelativePath) || strings.Contains(filepath.Clean(projection.ControlledSessionRelativePath), "..")):
		return KnownAppKDERuntimeStatusLaunchDelegatedEvidence{}, errors.New("known app KDE Runtime-status delegated evidence requires relative session evidence")
	case projection.StateRootPathExposed || projection.ManagedLauncherPathExposed || projection.RawLauncherOutputExposed || projection.BackendDetailsExposed:
		return KnownAppKDERuntimeStatusLaunchDelegatedEvidence{}, errors.New("known app KDE Runtime-status delegated evidence must not expose paths, launcher output, or backend details")
	case projection.HostRootModified || projection.ControlledSessionHostRootModified || projection.DockerSocketMounted || projection.BroadHostMountRequired:
		return KnownAppKDERuntimeStatusLaunchDelegatedEvidence{}, errors.New("known app KDE Runtime-status delegated evidence must keep host and container boundaries closed")
	case !projection.CompatibilityCenterProjectionReady || !projection.KDECenterProjectionReady:
		return KnownAppKDERuntimeStatusLaunchDelegatedEvidence{}, errors.New("known app KDE Runtime-status delegated evidence must be ready for Center projections")
	}
	for _, value := range []string{projection.ProjectionType, projection.RuntimeMethod, projection.RequestType, projection.Status, projection.SkipReason, projection.FailureReason, projection.AppID, projection.DisplayName, projection.AppVersion, projection.GuestBoundary, projection.SessionGatedControlledDispatchState, projection.SessionGatedReviewReceiptID, projection.LaunchAuthorizationReceiptID, projection.ControlledExecutionSessionID, projection.ControlledSessionRelativePath} {
		if value != "" && !singleLine(value) {
			return KnownAppKDERuntimeStatusLaunchDelegatedEvidence{}, errors.New("known app KDE Runtime-status delegated evidence requires single-line fields")
		}
	}
	if err := validateNoBackendTerms(projection, "known app KDE Runtime-status delegated evidence"); err != nil {
		return KnownAppKDERuntimeStatusLaunchDelegatedEvidence{}, err
	}
	return projection, nil
}

func validateKnownAppKDERuntimeStatusLaunchExecutionPlan(plan KnownAppKDERuntimeStatusLaunchExecutionPlan) (KnownAppKDERuntimeStatusLaunchExecutionPlan, error) {
	switch {
	case plan.SchemaVersion != KnownAppKDERuntimeStatusLaunchExecutionSchemaVersion || plan.RequestType != KnownAppKDERuntimeStatusLaunchExecutionRequestType:
		return KnownAppKDERuntimeStatusLaunchExecutionPlan{}, errors.New("known app KDE Runtime-status launch execution has invalid schema")
	case plan.RequestPreviewType != KnownAppKDERuntimeStatusLaunchRequestType:
		return KnownAppKDERuntimeStatusLaunchExecutionPlan{}, errors.New("known app KDE Runtime-status launch execution requires the KDE request preview")
	case plan.ActionID != KnownAppKDERuntimeStatusLaunchAction || plan.ActionKind != "runtime-status":
		return KnownAppKDERuntimeStatusLaunchExecutionPlan{}, errors.New("known app KDE Runtime-status launch execution has invalid action")
	case plan.CenterCardState != "validated-post-review-dispatch" || plan.PostReviewDispatchState != "created-after-session-gated-review":
		return KnownAppKDERuntimeStatusLaunchExecutionPlan{}, errors.New("known app KDE Runtime-status launch execution requires validated post-review state")
	case !plan.LaunchReceiptRevalidated || !plan.GuestBoundaryRevalidated || !plan.ReviewReceiptRevalidated || !plan.ControlledSessionRevalidated || !plan.ControlledSessionDigestVerified:
		return KnownAppKDERuntimeStatusLaunchExecutionPlan{}, errors.New("known app KDE Runtime-status launch execution requires revalidated receipts and session evidence")
	case !plan.RuntimeManagedLauncherArgvReady || len(plan.RuntimeManagedLauncherArgv) != 13:
		return KnownAppKDERuntimeStatusLaunchExecutionPlan{}, errors.New("known app KDE Runtime-status launch execution requires Runtime-managed launcher argv")
	case plan.RequiredOpaqueIDCount != 3 || plan.CollectedOpaqueIDCount != 3:
		return KnownAppKDERuntimeStatusLaunchExecutionPlan{}, errors.New("known app KDE Runtime-status launch execution must collect required opaque ids")
	case !plan.StateRootRequired || !plan.StateRootAccepted || !plan.StateRootInjectedByRuntime || !plan.StateRootSuppliedByRuntime || plan.KDEStateRootAccess:
		return KnownAppKDERuntimeStatusLaunchExecutionPlan{}, errors.New("known app KDE Runtime-status launch execution must keep state root Runtime-owned")
	case !plan.ManagedLauncherInvocationReady || !plan.ExistingManagedLauncherPathUsed || !plan.DelegatesArtifactGateToLauncher:
		return KnownAppKDERuntimeStatusLaunchExecutionPlan{}, errors.New("known app KDE Runtime-status launch execution must delegate to the existing managed launcher")
	case !plan.RuntimeOwnedRequest || !plan.RuntimeOwnedLaunch || !plan.RuntimeOwnedDispatch || !plan.KDEPresentationOnly:
		return KnownAppKDERuntimeStatusLaunchExecutionPlan{}, errors.New("known app KDE Runtime-status launch execution must keep Runtime ownership")
	case plan.DirectLaunchEnabled || plan.DesktopLaunchEnabled || plan.BackendLaunchEnabled || plan.ExecutionStarted || plan.BackendProcessStarted:
		return KnownAppKDERuntimeStatusLaunchExecutionPlan{}, errors.New("known app KDE Runtime-status launch execution plan must not start execution before the launcher is invoked")
	case plan.RequestObjectsCreated || plan.PermissionGrantCreated:
		return KnownAppKDERuntimeStatusLaunchExecutionPlan{}, errors.New("known app KDE Runtime-status launch execution must not create request objects or grants")
	case plan.StateRootPathExposed || plan.ReceiptPathExposed || plan.SessionPathExposed || plan.ManagedLauncherPathExposed || plan.RawArtifactPathExposed || plan.RawCommandExposed || plan.RawLauncherOutputExposed || plan.BackendDetailsExposed:
		return KnownAppKDERuntimeStatusLaunchExecutionPlan{}, errors.New("known app KDE Runtime-status launch execution must not expose paths, raw commands, launcher output, or backend details")
	case plan.HostRootModified || plan.PrivilegedContainerRequired || plan.DockerSocketMounted || plan.BroadHostMountRequired:
		return KnownAppKDERuntimeStatusLaunchExecutionPlan{}, errors.New("known app KDE Runtime-status launch execution must keep host and container boundaries closed")
	}
	for _, value := range []string{plan.AppID, plan.DisplayName, plan.AppVersion, plan.ActionID, plan.CenterCardState, plan.PostReviewDispatchState, plan.LaunchAuthorizationReceiptID, plan.SessionGatedReviewReceiptID, plan.ControlledExecutionSessionID, plan.LaunchGateState, plan.ReviewGateState} {
		if !singleLine(value) {
			return KnownAppKDERuntimeStatusLaunchExecutionPlan{}, errors.New("known app KDE Runtime-status launch execution requires single-line fields")
		}
	}
	if !containsString(plan.RuntimeManagedLauncherArgv, knownAppRuntimeOwnedStateRootPlaceholder) {
		return KnownAppKDERuntimeStatusLaunchExecutionPlan{}, errors.New("known app KDE Runtime-status launch execution must redact Runtime state root")
	}
	if err := validateNoBackendTerms(plan, "known app KDE Runtime-status launch execution"); err != nil {
		return KnownAppKDERuntimeStatusLaunchExecutionPlan{}, err
	}
	return plan, nil
}

func knownAppRuntimeStatusLaunchRuntimeArgv(request KnownAppKDERuntimeStatusLaunchPreview) []string {
	return []string{
		"xnix-compat-launch",
		"--app", request.AppID,
		"--guest-boundary", winapp.KnownDispatchGuestBoundary,
		"--state-root", knownAppRuntimeOwnedStateRootPlaceholder,
		"--receipt-id", request.LaunchAuthorizationReceiptID,
		"--review-receipt-id", request.SessionGatedReviewReceiptID,
		"--session-id", request.ControlledExecutionSessionID,
	}
}

func validateKnownAppRuntimeOwnerStateRoot(root string) error {
	root = strings.TrimSpace(root)
	if root == "" {
		return errors.New("known app KDE Runtime-status launch execution requires --state-root")
	}
	abs, err := filepath.Abs(root)
	if err != nil {
		return fmt.Errorf("resolve Runtime owner state root: %w", err)
	}
	if abs == string(os.PathSeparator) {
		return errors.New("refusing to use filesystem root for Runtime owner state root")
	}
	info, err := os.Lstat(abs)
	if err != nil {
		return fmt.Errorf("Runtime owner state root must exist: %w", err)
	}
	if !info.IsDir() || info.Mode()&os.ModeSymlink != 0 {
		return errors.New("Runtime owner state root must be a real directory")
	}
	return nil
}
