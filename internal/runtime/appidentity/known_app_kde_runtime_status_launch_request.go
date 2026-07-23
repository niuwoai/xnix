package appidentity

import (
	"errors"
	"fmt"
	"strings"

	"xnix.local/xnix/internal/runtime/winapp"
)

const (
	KnownAppKDERuntimeStatusLaunchRequestSchemaVersion = "xnix.runtime.known_app_kde_runtime_status_launch_request.v1"
	KnownAppKDERuntimeStatusLaunchRequestType          = "known-app-kde-runtime-status-launch-request-preview"
	KnownAppKDERuntimeStatusLaunchAction               = "show-runtime-controlled-launch"
)

type KnownAppKDERuntimeStatusLaunchRequest struct {
	AppID                        string
	LaunchAuthorizationReceiptID string
	SessionGatedReviewReceiptID  string
	ControlledExecutionSessionID string
	CenterCardState              string
	PrimaryActionID              string
	PostReviewDispatchState      string
}

type KnownAppKDERuntimeStatusLaunchPreview struct {
	SchemaVersion                string   `json:"schema_version"`
	RequestType                  string   `json:"request_type"`
	Source                       string   `json:"source"`
	Desktop                      string   `json:"desktop"`
	RuntimeMethod                string   `json:"runtime_method"`
	ReadMethod                   string   `json:"read_method"`
	AppID                        string   `json:"app_id"`
	DisplayName                  string   `json:"display_name"`
	AppVersion                   string   `json:"app_version"`
	ActionID                     string   `json:"action_id"`
	ActionKind                   string   `json:"action_kind"`
	CenterCardState              string   `json:"center_card_state"`
	PostReviewDispatchState      string   `json:"post_review_dispatch_state"`
	LaunchAuthorizationReceiptID string   `json:"launch_authorization_receipt_id"`
	SessionGatedReviewReceiptID  string   `json:"session_gated_review_receipt_id"`
	ControlledExecutionSessionID string   `json:"controlled_execution_session_id"`
	ManagedLauncher              string   `json:"managed_launcher"`
	ManagedLauncherArgv          []string `json:"managed_launcher_argv"`
	ManagedLauncherArgvReady     bool     `json:"managed_launcher_argv_ready"`
	RequiredOpaqueIDCount        int      `json:"required_opaque_id_count"`
	CollectedOpaqueIDCount       int      `json:"collected_opaque_id_count"`
	LaunchRequestCreated         bool     `json:"launch_request_created"`
	RuntimeOwnedRequest          bool     `json:"runtime_owned_request"`
	RuntimeOwnedLaunch           bool     `json:"runtime_owned_launch"`
	RuntimeOwnedDispatch         bool     `json:"runtime_owned_dispatch"`
	KDEPresentationOnly          bool     `json:"kde_presentation_only"`
	KDEActionForwarded           bool     `json:"kde_action_forwarded"`
	StateRootRequired            bool     `json:"state_root_required"`
	StateRootSuppliedByRuntime   bool     `json:"state_root_supplied_by_runtime"`
	KDEStateRootAccess           bool     `json:"kde_state_root_access"`
	DirectLaunchEnabled          bool     `json:"direct_launch_enabled"`
	DesktopLaunchEnabled         bool     `json:"desktop_launch_enabled"`
	BackendLaunchEnabled         bool     `json:"backend_launch_enabled"`
	ExecutionStarted             bool     `json:"execution_started"`
	BackendProcessStarted        bool     `json:"backend_process_started"`
	RequestObjectsCreated        bool     `json:"request_objects_created"`
	PermissionGrantCreated       bool     `json:"permission_grant_created"`
	StateRootPathExposed         bool     `json:"state_root_path_exposed"`
	ReceiptPathExposed           bool     `json:"receipt_path_exposed"`
	SessionPathExposed           bool     `json:"session_path_exposed"`
	RawArtifactPathExposed       bool     `json:"raw_artifact_path_exposed"`
	RawCommandExposed            bool     `json:"raw_command_exposed"`
	BackendDetailsExposed        bool     `json:"backend_details_exposed"`
	HostRootModified             bool     `json:"host_root_modified"`
	PrivilegedContainerRequired  bool     `json:"privileged_container_required"`
	DockerSocketMounted          bool     `json:"docker_socket_mounted"`
	BroadHostMountRequired       bool     `json:"broad_host_mount_required"`
	DesktopSafeSummary           string   `json:"desktop_safe_summary"`
}

func PreviewKnownAppKDERuntimeStatusLaunchRequest(request KnownAppKDERuntimeStatusLaunchRequest) (KnownAppKDERuntimeStatusLaunchPreview, error) {
	app, err := winapp.LookupKnownPortableApp(request.AppID)
	if err != nil {
		return KnownAppKDERuntimeStatusLaunchPreview{}, err
	}
	launchReceiptID := strings.TrimSpace(request.LaunchAuthorizationReceiptID)
	reviewReceiptID := strings.TrimSpace(request.SessionGatedReviewReceiptID)
	sessionID := strings.TrimSpace(request.ControlledExecutionSessionID)
	centerCardState := strings.TrimSpace(request.CenterCardState)
	primaryActionID := strings.TrimSpace(request.PrimaryActionID)
	postReviewState := strings.TrimSpace(request.PostReviewDispatchState)
	if centerCardState != "validated-post-review-dispatch" {
		return KnownAppKDERuntimeStatusLaunchPreview{}, errors.New("known app KDE Runtime-status launch request requires validated post-review dispatch state")
	}
	if primaryActionID != KnownAppKDERuntimeStatusLaunchAction {
		return KnownAppKDERuntimeStatusLaunchPreview{}, fmt.Errorf("known app KDE Runtime-status launch request requires action %s", KnownAppKDERuntimeStatusLaunchAction)
	}
	if postReviewState != "created-after-session-gated-review" {
		return KnownAppKDERuntimeStatusLaunchPreview{}, errors.New("known app KDE Runtime-status launch request requires post-review dispatch state")
	}
	if !safeOpaqueKnownAppID(launchReceiptID) || !safeOpaqueKnownAppID(reviewReceiptID) || !safeOpaqueKnownAppID(sessionID) {
		return KnownAppKDERuntimeStatusLaunchPreview{}, errors.New("known app KDE Runtime-status launch request requires safe opaque ids")
	}
	if launchReceiptID != KnownAppLaunchAuthorizationReceiptID(app.ID, app.Version) {
		return KnownAppKDERuntimeStatusLaunchPreview{}, errors.New("known app KDE Runtime-status launch request received mismatched launch authorization receipt id")
	}
	if sessionID != KnownAppControlledExecutionSessionID(app.ID, app.Version) {
		return KnownAppKDERuntimeStatusLaunchPreview{}, errors.New("known app KDE Runtime-status launch request received mismatched controlled execution session id")
	}
	if reviewReceiptID != KnownAppSessionGatedLaunchReviewReceiptID(app.ID, app.Version, sessionID) {
		return KnownAppKDERuntimeStatusLaunchPreview{}, errors.New("known app KDE Runtime-status launch request received mismatched session-gated review receipt id")
	}
	argv := []string{
		"xnix-compat-launch",
		"--app", app.ID,
		"--guest-boundary", winapp.KnownDispatchGuestBoundary,
		"--receipt-id", launchReceiptID,
		"--review-receipt-id", reviewReceiptID,
		"--session-id", sessionID,
	}
	preview := KnownAppKDERuntimeStatusLaunchPreview{
		SchemaVersion:                KnownAppKDERuntimeStatusLaunchRequestSchemaVersion,
		RequestType:                  KnownAppKDERuntimeStatusLaunchRequestType,
		Source:                       "kde-center-page-preview+compatibility-center-preview",
		Desktop:                      "KDE Plasma",
		RuntimeMethod:                "PreviewKnownAppKDERuntimeStatusLaunchRequest",
		ReadMethod:                   "GetKnownAppKDERuntimeStatusLaunchRequest",
		AppID:                        app.ID,
		DisplayName:                  app.DisplayName,
		AppVersion:                   app.Version,
		ActionID:                     KnownAppKDERuntimeStatusLaunchAction,
		ActionKind:                   "runtime-status",
		CenterCardState:              centerCardState,
		PostReviewDispatchState:      postReviewState,
		LaunchAuthorizationReceiptID: launchReceiptID,
		SessionGatedReviewReceiptID:  reviewReceiptID,
		ControlledExecutionSessionID: sessionID,
		ManagedLauncher:              strings.Join(argv, " "),
		ManagedLauncherArgv:          argv,
		ManagedLauncherArgvReady:     true,
		RequiredOpaqueIDCount:        3,
		CollectedOpaqueIDCount:       3,
		LaunchRequestCreated:         true,
		RuntimeOwnedRequest:          true,
		RuntimeOwnedLaunch:           true,
		RuntimeOwnedDispatch:         true,
		KDEPresentationOnly:          true,
		KDEActionForwarded:           true,
		StateRootRequired:            true,
		StateRootSuppliedByRuntime:   true,
		KDEStateRootAccess:           false,
		DirectLaunchEnabled:          false,
		DesktopLaunchEnabled:         false,
		BackendLaunchEnabled:         false,
		ExecutionStarted:             false,
		BackendProcessStarted:        false,
		RequestObjectsCreated:        false,
		PermissionGrantCreated:       false,
		StateRootPathExposed:         false,
		ReceiptPathExposed:           false,
		SessionPathExposed:           false,
		RawArtifactPathExposed:       false,
		RawCommandExposed:            false,
		BackendDetailsExposed:        false,
		HostRootModified:             false,
		PrivilegedContainerRequired:  false,
		DockerSocketMounted:          false,
		BroadHostMountRequired:       false,
		DesktopSafeSummary:           app.DisplayName + " has a Runtime-owned managed launcher request assembled from KDE-safe opaque ids.",
	}
	return validateKnownAppKDERuntimeStatusLaunchPreview(preview)
}

func validateKnownAppKDERuntimeStatusLaunchPreview(preview KnownAppKDERuntimeStatusLaunchPreview) (KnownAppKDERuntimeStatusLaunchPreview, error) {
	switch {
	case preview.SchemaVersion != KnownAppKDERuntimeStatusLaunchRequestSchemaVersion || preview.RequestType != KnownAppKDERuntimeStatusLaunchRequestType:
		return KnownAppKDERuntimeStatusLaunchPreview{}, errors.New("known app KDE Runtime-status launch request has invalid schema")
	case preview.ActionID != KnownAppKDERuntimeStatusLaunchAction || preview.ActionKind != "runtime-status":
		return KnownAppKDERuntimeStatusLaunchPreview{}, errors.New("known app KDE Runtime-status launch request has invalid action")
	case preview.CenterCardState != "validated-post-review-dispatch" || preview.PostReviewDispatchState != "created-after-session-gated-review":
		return KnownAppKDERuntimeStatusLaunchPreview{}, errors.New("known app KDE Runtime-status launch request requires accepted post-review state")
	case !preview.ManagedLauncherArgvReady || len(preview.ManagedLauncherArgv) != 11:
		return KnownAppKDERuntimeStatusLaunchPreview{}, errors.New("known app KDE Runtime-status launch request requires managed launcher argv")
	case preview.RequiredOpaqueIDCount != 3 || preview.CollectedOpaqueIDCount != 3:
		return KnownAppKDERuntimeStatusLaunchPreview{}, errors.New("known app KDE Runtime-status launch request must collect required opaque ids")
	case !preview.StateRootRequired || !preview.StateRootSuppliedByRuntime || preview.KDEStateRootAccess:
		return KnownAppKDERuntimeStatusLaunchPreview{}, errors.New("known app KDE Runtime-status launch request must keep state root Runtime-owned")
	case !preview.RuntimeOwnedRequest || !preview.RuntimeOwnedLaunch || !preview.RuntimeOwnedDispatch || !preview.KDEPresentationOnly:
		return KnownAppKDERuntimeStatusLaunchPreview{}, errors.New("known app KDE Runtime-status launch request must keep Runtime ownership")
	case preview.DirectLaunchEnabled || preview.DesktopLaunchEnabled || preview.BackendLaunchEnabled || preview.ExecutionStarted || preview.BackendProcessStarted:
		return KnownAppKDERuntimeStatusLaunchPreview{}, errors.New("known app KDE Runtime-status launch request must not launch or start backend processes")
	case preview.RequestObjectsCreated || preview.PermissionGrantCreated:
		return KnownAppKDERuntimeStatusLaunchPreview{}, errors.New("known app KDE Runtime-status launch request must not write request objects or grants")
	case preview.StateRootPathExposed || preview.ReceiptPathExposed || preview.SessionPathExposed || preview.RawArtifactPathExposed || preview.RawCommandExposed || preview.BackendDetailsExposed:
		return KnownAppKDERuntimeStatusLaunchPreview{}, errors.New("known app KDE Runtime-status launch request must not expose paths, raw commands, or backend details")
	case preview.HostRootModified || preview.PrivilegedContainerRequired || preview.DockerSocketMounted || preview.BroadHostMountRequired:
		return KnownAppKDERuntimeStatusLaunchPreview{}, errors.New("known app KDE Runtime-status launch request must keep host and container boundaries closed")
	}
	for _, value := range []string{preview.AppID, preview.DisplayName, preview.AppVersion, preview.ActionID, preview.CenterCardState, preview.PostReviewDispatchState, preview.LaunchAuthorizationReceiptID, preview.SessionGatedReviewReceiptID, preview.ControlledExecutionSessionID} {
		if !singleLine(value) {
			return KnownAppKDERuntimeStatusLaunchPreview{}, errors.New("known app KDE Runtime-status launch request requires single-line fields")
		}
	}
	if err := validateNoBackendTerms(preview, "known app KDE Runtime-status launch request"); err != nil {
		return KnownAppKDERuntimeStatusLaunchPreview{}, err
	}
	return preview, nil
}

func safeOpaqueKnownAppID(value string) bool {
	return singleLine(value) && !strings.ContainsAny(value, `/\`) && !strings.Contains(value, "..")
}
