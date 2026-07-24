package owner

import (
	"encoding/json"
	"errors"
	"strings"

	"xnix.local/xnix/internal/runtime/appidentity"
)

const (
	DesktopTriggerRequestPreflightSchemaVersion = "xnix.runtime.desktop_trigger_request_preflight.v1"
	DesktopTriggerRequestPreflightRequestType   = "desktop-trigger-request-preflight-preview"
)

type DesktopTriggerRequestPreflightRequest struct {
	StateRoot              string
	DesktopEntryContent    string
	EvidenceID             string
	EvidenceRelativePath   string
	ExpectedEvidenceSHA256 string
	FullCheckpointPromoted bool
	RouteID                string
	MethodID               string
	ActionID               string
	CallerRole             string
	RequestFreshness       string
	ReplayMarker           string
	KDEStateRoot           string
	KDECacheRoot           string
	KDELauncherPath        string
	KDETimeoutValue        string
	KDERawExecutablePath   string
	KDEBackendCommand      string
	KDEReceiptID           string
	KDESessionID           string
	KDEDispatchID          string
}

type DesktopTriggerRequestPreflightPreview struct {
	SchemaVersion                     string   `json:"schema_version"`
	RequestType                       string   `json:"request_type"`
	Source                            string   `json:"source"`
	RuntimeMethod                     string   `json:"runtime_method"`
	ReadMethod                        string   `json:"read_method"`
	PreflightState                    string   `json:"preflight_state"`
	MaterializationState              string   `json:"materialization_state"`
	DryRunReviewState                 string   `json:"dry_run_review_state"`
	OwnerTriggerState                 string   `json:"owner_trigger_state"`
	FullCheckpointState               string   `json:"full_checkpoint_state"`
	RuntimeStatusEvidenceState        string   `json:"runtime_status_evidence_state"`
	EvidenceHandleKind                string   `json:"evidence_handle_kind"`
	EvidenceHandle                    string   `json:"evidence_handle,omitempty"`
	EvidenceID                        string   `json:"evidence_id,omitempty"`
	EvidenceRelativePath              string   `json:"evidence_relative_path,omitempty"`
	EvidenceSHA256                    string   `json:"evidence_sha256,omitempty"`
	ExpectedEvidenceSHA256            string   `json:"expected_evidence_sha256,omitempty"`
	EvidenceDigestVerified            bool     `json:"evidence_digest_verified"`
	ExpectedDigestMatched             bool     `json:"expected_digest_matched"`
	DesktopCallableRoute              string   `json:"desktop_callable_route"`
	DesktopCallableRuntimeMethod      string   `json:"desktop_callable_runtime_method"`
	DesktopCallableExecutionType      string   `json:"desktop_callable_execution_type"`
	DesktopDBusMethod                 string   `json:"desktop_dbus_method"`
	OwnerServiceBoundary              string   `json:"owner_service_boundary"`
	OwnerServiceMethod                string   `json:"owner_service_method"`
	OwnerServiceCallType              string   `json:"owner_service_call_type"`
	OwnerServiceCallShapeVerified     bool     `json:"owner_service_call_shape_verified"`
	OwnerServiceCallReady             bool     `json:"owner_service_call_ready"`
	OperatorRequestReady              bool     `json:"operator_request_ready"`
	FormalPromotionRequired           bool     `json:"formal_promotion_required"`
	FormalPromotionObserved           bool     `json:"formal_promotion_observed"`
	FormalReleaseReady                bool     `json:"formal_release_ready"`
	HumanAuthorizationRequired        bool     `json:"human_authorization_required"`
	RuntimeOwnerServiceSuppliesInputs bool     `json:"runtime_owner_service_supplies_inputs"`
	DesktopEvidenceHandleForwarded    bool     `json:"desktop_evidence_handle_forwarded"`
	RuntimeOwned                      bool     `json:"runtime_owned"`
	RuntimeOwnedDispatch              bool     `json:"runtime_owned_dispatch"`
	GoRuntimeBacked                   bool     `json:"go_runtime_backed"`
	KDEPresentationOnly               bool     `json:"kde_presentation_only"`
	KDEForwardsOnlyEvidenceHandle     bool     `json:"kde_forwards_only_evidence_handle"`
	KDEReceivesMaterializedOwnerArgs  bool     `json:"kde_receives_materialized_owner_args"`
	DesktopReceiptFieldsReconstructed bool     `json:"desktop_receipt_fields_reconstructed"`
	DesktopKDEStateRootAccess         bool     `json:"desktop_kde_state_root_access"`
	StateRootPathExposed              bool     `json:"state_root_path_exposed"`
	CacheRootPathExposed              bool     `json:"cache_root_path_exposed"`
	LauncherPathExposed               bool     `json:"launcher_path_exposed"`
	RawExecutablePathExposed          bool     `json:"raw_executable_path_exposed"`
	RawCommandExposed                 bool     `json:"raw_command_exposed"`
	RawLauncherOutputExposed          bool     `json:"raw_launcher_output_exposed"`
	BackendDetailsExposed             bool     `json:"backend_details_exposed"`
	RequestObjectWritten              bool     `json:"request_object_written"`
	PermissionGrantCreated            bool     `json:"permission_grant_created"`
	ServiceCallDispatchEnabled        bool     `json:"service_call_dispatch_enabled"`
	ServiceCallDispatched             bool     `json:"service_call_dispatched"`
	DBusCalled                        bool     `json:"dbus_called"`
	DBusOwnershipEnabled              bool     `json:"dbus_ownership_enabled"`
	DesktopLaunchEnabled              bool     `json:"desktop_launch_enabled"`
	BackendLaunchEnabled              bool     `json:"backend_launch_enabled"`
	ExecutionStarted                  bool     `json:"execution_started"`
	BackendProcessStarted             bool     `json:"backend_process_started"`
	RuntimeStateWritten               bool     `json:"runtime_state_written"`
	KDEConfigurationWritten           bool     `json:"kde_configuration_written"`
	SmokeExecutedByPreview            bool     `json:"smoke_executed_by_preview"`
	NetworkRequired                   bool     `json:"network_required"`
	HostRootModified                  bool     `json:"host_root_modified"`
	DockerSocketMounted               bool     `json:"docker_socket_mounted"`
	BroadHostMountRequired            bool     `json:"broad_host_mount_required"`
	PrivilegedContainerRequired       bool     `json:"privileged_container_required"`
	HostNetworkRequired               bool     `json:"host_network_required"`
	ProductionAuthorizationAccepted   bool     `json:"production_authorization_accepted"`
	BlockedReason                     string   `json:"blocked_reason,omitempty"`
	BlockedActions                    []string `json:"blocked_actions"`
	NextRequirements                  []string `json:"next_requirements"`
	NextHumanAuthorizedAction         string   `json:"next_human_authorized_action"`
	NextHumanAuthorizedCommand        []string `json:"next_human_authorized_command"`
	DesktopSafeSummary                string   `json:"desktop_safe_summary"`
}

func PreviewDesktopTriggerRequestPreflight(request DesktopTriggerRequestPreflightRequest) (DesktopTriggerRequestPreflightPreview, error) {
	preview := baseDesktopTriggerRequestPreflight(request)
	handleKind, handle := desktopTriggerRequestPreflightEvidenceHandle(request)
	preview.EvidenceHandleKind = handleKind
	preview.EvidenceHandle = handle

	materialization, err := PreviewDesktopTriggerServiceCallMaterialization(desktopTriggerRequestPreflightMaterializationRequest(request))
	if err != nil {
		state := classifyDesktopTriggerDryRunError(err)
		preview.PreflightState = desktopTriggerRequestPreflightStateFromMaterializationState(state)
		preview.MaterializationState = state
		preview.BlockedReason = "desktop-trigger request preflight is blocked because service call materialization failed"
		preview.DesktopSafeSummary = desktopTriggerRequestPreflightSummary(preview)
		return validateDesktopTriggerRequestPreflightPreview(preview)
	}

	preview.MaterializationState = materialization.MaterializationState
	preview.DryRunReviewState = materialization.DryRunReviewState
	preview.OwnerTriggerState = materialization.OwnerTriggerState
	preview.FullCheckpointState = materialization.FullCheckpointState
	preview.RuntimeStatusEvidenceState = materialization.RuntimeStatusEvidenceState
	preview.EvidenceID = materialization.EvidenceID
	preview.EvidenceRelativePath = materialization.EvidenceRelativePath
	preview.EvidenceSHA256 = materialization.EvidenceSHA256
	preview.EvidenceDigestVerified = materialization.EvidenceDigestVerified
	preview.ExpectedDigestMatched = materialization.ExpectedDigestMatched
	preview.DesktopCallableRoute = materialization.DesktopCallableRoute
	preview.DesktopCallableRuntimeMethod = materialization.DesktopCallableRuntimeMethod
	preview.DesktopCallableExecutionType = materialization.DesktopCallableExecutionType
	preview.DesktopDBusMethod = materialization.DesktopDBusMethod
	preview.OwnerServiceBoundary = materialization.OwnerServiceBoundary
	preview.OwnerServiceMethod = materialization.OwnerServiceMethod
	preview.OwnerServiceCallType = materialization.OwnerServiceCallType
	preview.OwnerServiceCallReady = materialization.OwnerServiceCallReady
	preview.FormalPromotionObserved = materialization.FullCheckpointPromotionClaimed
	preview.RuntimeOwnerServiceSuppliesInputs = materialization.RuntimeOwnerServiceSuppliesInputs
	preview.DesktopEvidenceHandleForwarded = materialization.DesktopEvidenceHandleForwarded
	preview.RuntimeOwnedDispatch = materialization.RuntimeOwnedDispatch
	preview.PreflightState = desktopTriggerRequestPreflightState(materialization)
	preview.OperatorRequestReady = preview.PreflightState == "ready-for-operator-request"
	preview.OwnerServiceCallShapeVerified = preview.OperatorRequestReady && sameDesktopTriggerServiceCallMaterializationArgs(materialization.OwnerServiceCallArgs, []string{"ShowRuntimeControlledLaunch", "evidence-relative-path", materialization.EvidenceRelativePath})
	preview.BlockedReason = desktopTriggerRequestPreflightBlockedReason(preview)
	preview.DesktopSafeSummary = desktopTriggerRequestPreflightSummary(preview)
	return validateDesktopTriggerRequestPreflightPreview(preview)
}

func baseDesktopTriggerRequestPreflight(request DesktopTriggerRequestPreflightRequest) DesktopTriggerRequestPreflightPreview {
	fullCheckpointState := "blocked-missing-promotion"
	if request.FullCheckpointPromoted {
		fullCheckpointState = "ready"
	}
	return DesktopTriggerRequestPreflightPreview{
		SchemaVersion:                     DesktopTriggerRequestPreflightSchemaVersion,
		RequestType:                       DesktopTriggerRequestPreflightRequestType,
		Source:                            DesktopTriggerServiceCallMaterializationRequestType + "+post-release-desktop-request-preflight",
		RuntimeMethod:                     "PreviewDesktopTriggerRequestPreflight",
		ReadMethod:                        "GetDesktopTriggerRequestPreflight",
		PreflightState:                    "blocked",
		MaterializationState:              "blocked",
		DryRunReviewState:                 "blocked",
		OwnerTriggerState:                 "blocked",
		FullCheckpointState:               fullCheckpointState,
		RuntimeStatusEvidenceState:        "blocked",
		ExpectedEvidenceSHA256:            strings.TrimSpace(request.ExpectedEvidenceSHA256),
		DesktopCallableRoute:              "kde-dbus-runtime-status-action",
		DesktopCallableRuntimeMethod:      "ShowRuntimeControlledLaunch",
		DesktopCallableExecutionType:      appidentity.KnownAppKDERuntimeStatusLaunchExecutionRequestType,
		DesktopDBusMethod:                 "org.xnix.Compatibility1.ShowRuntimeControlledLaunch",
		OwnerServiceBoundary:              "go-runtime-owner-in-process-service",
		OwnerServiceMethod:                "ShowRuntimeControlledLaunch",
		OwnerServiceCallType:              "desktop-action-dispatch",
		FormalPromotionRequired:           true,
		FormalPromotionObserved:           request.FullCheckpointPromoted,
		FormalReleaseReady:                false,
		HumanAuthorizationRequired:        true,
		RuntimeOwnerServiceSuppliesInputs: true,
		DesktopEvidenceHandleForwarded:    true,
		RuntimeOwned:                      true,
		RuntimeOwnedDispatch:              true,
		GoRuntimeBacked:                   true,
		KDEPresentationOnly:               true,
		KDEForwardsOnlyEvidenceHandle:     true,
		KDEReceivesMaterializedOwnerArgs:  false,
		StateRootPathExposed:              false,
		CacheRootPathExposed:              false,
		LauncherPathExposed:               false,
		RawExecutablePathExposed:          false,
		RawCommandExposed:                 false,
		RawLauncherOutputExposed:          false,
		BackendDetailsExposed:             false,
		RequestObjectWritten:              false,
		PermissionGrantCreated:            false,
		ServiceCallDispatchEnabled:        false,
		ServiceCallDispatched:             false,
		DBusCalled:                        false,
		DBusOwnershipEnabled:              false,
		DesktopLaunchEnabled:              false,
		BackendLaunchEnabled:              false,
		ExecutionStarted:                  false,
		BackendProcessStarted:             false,
		RuntimeStateWritten:               false,
		KDEConfigurationWritten:           false,
		SmokeExecutedByPreview:            false,
		NetworkRequired:                   false,
		HostRootModified:                  false,
		DockerSocketMounted:               false,
		BroadHostMountRequired:            false,
		PrivilegedContainerRequired:       false,
		HostNetworkRequired:               false,
		ProductionAuthorizationAccepted:   false,
		BlockedActions:                    desktopTriggerRequestPreflightBlockedActions(),
		NextRequirements:                  desktopTriggerRequestPreflightNextRequirements(),
		NextHumanAuthorizedAction:         "operator-dispatches-desktop-triggered-staged-launch",
		NextHumanAuthorizedCommand:        []string{"ruby", "scripts/staged_launcher_dispatch_smoke.rb"},
		DesktopSafeSummary:                "Desktop-trigger request preflight is blocked before service dispatch.",
	}
}

func desktopTriggerRequestPreflightMaterializationRequest(request DesktopTriggerRequestPreflightRequest) DesktopTriggerServiceCallMaterializationRequest {
	return DesktopTriggerServiceCallMaterializationRequest{
		StateRoot:              request.StateRoot,
		DesktopEntryContent:    request.DesktopEntryContent,
		EvidenceID:             request.EvidenceID,
		EvidenceRelativePath:   request.EvidenceRelativePath,
		ExpectedEvidenceSHA256: request.ExpectedEvidenceSHA256,
		FullCheckpointPromoted: request.FullCheckpointPromoted,
		HumanAuthorizedSmoke:   false,
		RouteID:                request.RouteID,
		MethodID:               request.MethodID,
		ActionID:               request.ActionID,
		CallerRole:             request.CallerRole,
		RequestFreshness:       request.RequestFreshness,
		ReplayMarker:           request.ReplayMarker,
		KDEStateRoot:           request.KDEStateRoot,
		KDECacheRoot:           request.KDECacheRoot,
		KDELauncherPath:        request.KDELauncherPath,
		KDETimeoutValue:        request.KDETimeoutValue,
		KDERawExecutablePath:   request.KDERawExecutablePath,
		KDEBackendCommand:      request.KDEBackendCommand,
		KDEReceiptID:           request.KDEReceiptID,
		KDESessionID:           request.KDESessionID,
		KDEDispatchID:          request.KDEDispatchID,
	}
}

func desktopTriggerRequestPreflightEvidenceHandle(request DesktopTriggerRequestPreflightRequest) (string, string) {
	if strings.TrimSpace(request.EvidenceID) != "" {
		return "evidence-id", strings.TrimSpace(request.EvidenceID)
	}
	return "evidence-relative-path", strings.TrimSpace(request.EvidenceRelativePath)
}

func desktopTriggerRequestPreflightState(materialization DesktopTriggerServiceCallMaterializationPreview) string {
	if materialization.FullCheckpointState != "ready" {
		return "blocked-missing-promotion"
	}
	return desktopTriggerRequestPreflightStateFromMaterializationState(materialization.MaterializationState)
}

func desktopTriggerRequestPreflightStateFromMaterializationState(state string) string {
	switch state {
	case "ready-for-human-authorized-service-call":
		return "ready-for-operator-request"
	case "blocked-missing-full-checkpoint", "needs-full-checkpoint":
		return "blocked-missing-promotion"
	case "blocked-unsafe-envelope", "unsafe":
		return "blocked-unsafe-envelope"
	case "blocked-unsafe-action", "unsupported":
		return "blocked-unsafe-action"
	case "stale-evidence":
		return "blocked-stale-evidence"
	case "missing-evidence":
		return "missing-evidence"
	case "malformed":
		return "malformed"
	default:
		return "blocked"
	}
}

func desktopTriggerRequestPreflightBlockedReason(preview DesktopTriggerRequestPreflightPreview) string {
	switch preview.PreflightState {
	case "ready-for-operator-request":
		return ""
	case "blocked-missing-promotion":
		return "desktop-trigger request preflight is blocked until formal full checkpoint promotion is observed"
	case "blocked-unsafe-envelope":
		return "desktop-trigger request preflight is blocked by the Runtime owner launch envelope guard"
	case "blocked-unsafe-action":
		return "desktop-trigger request preflight is blocked by the KDE action surface audit"
	case "blocked-stale-evidence":
		return "desktop-trigger request preflight is blocked because the Runtime evidence digest is stale"
	case "missing-evidence":
		return "desktop-trigger request preflight is blocked because required Runtime evidence is missing"
	case "malformed":
		return "desktop-trigger request preflight is blocked because the Runtime evidence or desktop action is malformed"
	default:
		return "desktop-trigger request preflight is blocked by incomplete Runtime-controlled launch materialization"
	}
}

func desktopTriggerRequestPreflightSummary(preview DesktopTriggerRequestPreflightPreview) string {
	if preview.PreflightState == "ready-for-operator-request" {
		return "Desktop-trigger request preflight is ready for a human-authorized staged Runtime owner service request without exposing owner arguments to KDE."
	}
	if preview.PreflightState == "blocked-missing-promotion" {
		return "Desktop-trigger request preflight is prepared, but formal full checkpoint promotion is still required before a real desktop-triggered staged request."
	}
	return "Desktop-trigger request preflight is blocked before request-object writes, permission grants, D-Bus calls, owner service dispatch, desktop launch, backend launch, Runtime writes, or host mutation."
}

func desktopTriggerRequestPreflightBlockedActions() []string {
	return []string{
		"write desktop-trigger Runtime request objects from preflight",
		"create permission grants from preflight",
		"dispatch Runtime owner service calls from preflight",
		"call D-Bus from preflight",
		"write KDE configuration from preflight",
		"start desktop or compatibility execution from preflight",
		"mutate host root from preflight",
	}
}

func desktopTriggerRequestPreflightNextRequirements() []string {
	return []string{
		"Observe formal v0.2.640 promotion evidence before accepting a real desktop-triggered staged request.",
		"Keep KDE restricted to opaque Runtime evidence handles.",
		"Let the Runtime owner service inject state roots, cache roots, launcher paths, and guest execution options.",
	}
}

func validateDesktopTriggerRequestPreflightPreview(preview DesktopTriggerRequestPreflightPreview) (DesktopTriggerRequestPreflightPreview, error) {
	validStates := map[string]bool{
		"ready-for-operator-request":              true,
		"blocked-missing-promotion":               true,
		"blocked-unsafe-envelope":                 true,
		"blocked-unsafe-action":                   true,
		"blocked-stale-evidence":                  true,
		"missing-evidence":                        true,
		"malformed":                               true,
		"blocked":                                 true,
		"ready-for-human-authorized-service-call": true,
		"accepted-review":                         true,
		"ready":                                   true,
		"blocked-missing-full-checkpoint":         true,
		"blocked-acceptance":                      true,
		"stale-evidence":                          true,
		"unsafe":                                  true,
		"needs-full-checkpoint":                   true,
	}
	for _, state := range []string{preview.PreflightState, preview.MaterializationState, preview.DryRunReviewState, preview.OwnerTriggerState, preview.FullCheckpointState, preview.RuntimeStatusEvidenceState} {
		if !validStates[state] {
			return DesktopTriggerRequestPreflightPreview{}, errors.New("desktop-trigger request preflight has invalid state")
		}
	}
	switch {
	case preview.SchemaVersion != DesktopTriggerRequestPreflightSchemaVersion || preview.RequestType != DesktopTriggerRequestPreflightRequestType:
		return DesktopTriggerRequestPreflightPreview{}, errors.New("desktop-trigger request preflight has invalid schema")
	case preview.OperatorRequestReady && (!preview.OwnerServiceCallShapeVerified || !preview.OwnerServiceCallReady || preview.FullCheckpointState != "ready" || !preview.FormalPromotionObserved):
		return DesktopTriggerRequestPreflightPreview{}, errors.New("desktop-trigger request preflight accepted incomplete dependencies")
	case !preview.OperatorRequestReady && (preview.OwnerServiceCallShapeVerified || preview.PreflightState == "ready-for-operator-request"):
		return DesktopTriggerRequestPreflightPreview{}, errors.New("desktop-trigger request preflight has inconsistent ready state")
	case !preview.FormalPromotionRequired || preview.FormalReleaseReady:
		return DesktopTriggerRequestPreflightPreview{}, errors.New("desktop-trigger request preflight must require formal promotion without claiming release readiness")
	case !preview.HumanAuthorizationRequired || !preview.RuntimeOwnerServiceSuppliesInputs || !preview.DesktopEvidenceHandleForwarded:
		return DesktopTriggerRequestPreflightPreview{}, errors.New("desktop-trigger request preflight has invalid authorization or trigger flags")
	case !preview.RuntimeOwned || !preview.RuntimeOwnedDispatch || !preview.GoRuntimeBacked || !preview.KDEPresentationOnly || !preview.KDEForwardsOnlyEvidenceHandle || preview.KDEReceivesMaterializedOwnerArgs:
		return DesktopTriggerRequestPreflightPreview{}, errors.New("desktop-trigger request preflight has invalid ownership flags")
	case preview.DesktopReceiptFieldsReconstructed || preview.DesktopKDEStateRootAccess || preview.StateRootPathExposed || preview.CacheRootPathExposed || preview.LauncherPathExposed || preview.RawExecutablePathExposed || preview.RawCommandExposed || preview.RawLauncherOutputExposed || preview.BackendDetailsExposed:
		return DesktopTriggerRequestPreflightPreview{}, errors.New("desktop-trigger request preflight must not expose Runtime internals")
	case preview.RequestObjectWritten || preview.PermissionGrantCreated || preview.ServiceCallDispatchEnabled || preview.ServiceCallDispatched || preview.DBusCalled || preview.DBusOwnershipEnabled || preview.DesktopLaunchEnabled || preview.BackendLaunchEnabled || preview.ExecutionStarted || preview.BackendProcessStarted || preview.RuntimeStateWritten || preview.KDEConfigurationWritten || preview.SmokeExecutedByPreview:
		return DesktopTriggerRequestPreflightPreview{}, errors.New("desktop-trigger request preflight must not dispatch, write, or launch")
	case preview.NetworkRequired || preview.HostRootModified || preview.DockerSocketMounted || preview.BroadHostMountRequired || preview.PrivilegedContainerRequired || preview.HostNetworkRequired || preview.ProductionAuthorizationAccepted:
		return DesktopTriggerRequestPreflightPreview{}, errors.New("desktop-trigger request preflight must keep host and container boundaries closed")
	}
	for _, value := range []string{
		preview.SchemaVersion, preview.RequestType, preview.Source, preview.RuntimeMethod, preview.ReadMethod, preview.PreflightState, preview.MaterializationState, preview.DryRunReviewState, preview.OwnerTriggerState, preview.FullCheckpointState, preview.RuntimeStatusEvidenceState, preview.EvidenceHandleKind, preview.EvidenceHandle, preview.EvidenceID, preview.EvidenceRelativePath, preview.EvidenceSHA256, preview.ExpectedEvidenceSHA256, preview.DesktopCallableRoute, preview.DesktopCallableRuntimeMethod, preview.DesktopCallableExecutionType, preview.DesktopDBusMethod, preview.OwnerServiceBoundary, preview.OwnerServiceMethod, preview.OwnerServiceCallType, preview.BlockedReason, preview.NextHumanAuthorizedAction, preview.DesktopSafeSummary,
	} {
		if value != "" && strings.ContainsAny(value, "\r\n") {
			return DesktopTriggerRequestPreflightPreview{}, errors.New("desktop-trigger request preflight requires single-line fields")
		}
	}
	for _, value := range append(append(append([]string{}, preview.BlockedActions...), preview.NextRequirements...), preview.NextHumanAuthorizedCommand...) {
		if strings.TrimSpace(value) == "" || strings.ContainsAny(value, "\r\n") {
			return DesktopTriggerRequestPreflightPreview{}, errors.New("desktop-trigger request preflight requires single-line list values")
		}
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		return DesktopTriggerRequestPreflightPreview{}, errors.New("desktop-trigger request preflight must be serializable")
	}
	lower := strings.ToLower(string(encoded))
	for _, forbidden := range []string{
		"owner_service_call_args",
		"owner_service_cli_args",
		"xnix_runtime_owner_",
		" --service-call ",
		" --state-root ",
		" --cache-root ",
		" --launcher ",
		" --receipt-id ",
		" --session-id ",
		" --dispatch-id ",
		".exe",
		"wine ",
		"proton",
		"qemu-system",
		"program files",
		".wine",
		"secret",
		"token",
		"password",
	} {
		if strings.Contains(lower, forbidden) {
			return DesktopTriggerRequestPreflightPreview{}, errors.New("desktop-trigger request preflight exposes forbidden terms")
		}
	}
	return preview, nil
}
