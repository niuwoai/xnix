package owner

import (
	"encoding/json"
	"errors"
	"strings"

	"xnix.local/xnix/internal/runtime/appidentity"
)

const (
	DesktopTriggerDryRunRequestReviewSchemaVersion = "xnix.runtime.desktop_trigger_dry_run_request_review.v1"
	DesktopTriggerDryRunRequestReviewRequestType   = "desktop-trigger-dry-run-request-review-preview"
)

type DesktopTriggerDryRunRequestReviewRequest struct {
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

type DesktopTriggerDryRunRequestReviewPreview struct {
	SchemaVersion                   string   `json:"schema_version"`
	RequestType                     string   `json:"request_type"`
	Source                          string   `json:"source"`
	RuntimeMethod                   string   `json:"runtime_method"`
	ReadMethod                      string   `json:"read_method"`
	ReviewState                     string   `json:"review_state"`
	EnvelopeGuardState              string   `json:"envelope_guard_state"`
	ActionSurfaceState              string   `json:"action_surface_state"`
	ManagedLauncherAcceptanceState  string   `json:"managed_launcher_acceptance_state"`
	FullCheckpointState             string   `json:"full_checkpoint_state"`
	RuntimeStatusEvidenceState      string   `json:"runtime_status_evidence_state"`
	EvidenceHandleKind              string   `json:"evidence_handle_kind"`
	EvidenceHandle                  string   `json:"evidence_handle,omitempty"`
	EvidenceID                      string   `json:"evidence_id,omitempty"`
	EvidenceRelativePath            string   `json:"evidence_relative_path,omitempty"`
	EvidenceSHA256                  string   `json:"evidence_sha256,omitempty"`
	ExpectedEvidenceSHA256          string   `json:"expected_evidence_sha256,omitempty"`
	EvidenceDigestVerified          bool     `json:"evidence_digest_verified"`
	ExpectedDigestMatched           bool     `json:"expected_digest_matched"`
	RouteID                         string   `json:"route_id"`
	MethodID                        string   `json:"method_id"`
	ActionID                        string   `json:"action_id"`
	CallerRole                      string   `json:"caller_role"`
	KDEActionID                     string   `json:"kde_action_id,omitempty"`
	PublicDBusMethod                string   `json:"public_dbus_method,omitempty"`
	ForwardedArgumentKind           string   `json:"forwarded_argument_kind,omitempty"`
	EnvelopeAcceptedForReview       bool     `json:"envelope_accepted_for_review"`
	ActionSurfaceSafeForHumanSmoke  bool     `json:"action_surface_safe_for_human_smoke"`
	ManagedLauncherAccepted         bool     `json:"managed_launcher_accepted"`
	DryRunReviewOnly                bool     `json:"dry_run_review_only"`
	RequestObjectWritten            bool     `json:"request_object_written"`
	PermissionGrantCreated          bool     `json:"permission_grant_created"`
	RuntimeOwned                    bool     `json:"runtime_owned"`
	GoRuntimeBacked                 bool     `json:"go_runtime_backed"`
	KDEPresentationOnly             bool     `json:"kde_presentation_only"`
	KDEForwardsOnlyEvidenceHandle   bool     `json:"kde_forwards_only_evidence_handle"`
	StateRootPathExposed            bool     `json:"state_root_path_exposed"`
	CacheRootPathExposed            bool     `json:"cache_root_path_exposed"`
	LauncherPathExposed             bool     `json:"launcher_path_exposed"`
	RawExecutablePathExposed        bool     `json:"raw_executable_path_exposed"`
	RawCommandExposed               bool     `json:"raw_command_exposed"`
	RawLauncherOutputExposed        bool     `json:"raw_launcher_output_exposed"`
	BackendDetailsExposed           bool     `json:"backend_details_exposed"`
	ServiceCallDispatchEnabled      bool     `json:"service_call_dispatch_enabled"`
	ServiceCallDispatched           bool     `json:"service_call_dispatched"`
	DBusCalled                      bool     `json:"dbus_called"`
	DBusOwnershipEnabled            bool     `json:"dbus_ownership_enabled"`
	DesktopLaunchEnabled            bool     `json:"desktop_launch_enabled"`
	BackendLaunchEnabled            bool     `json:"backend_launch_enabled"`
	ExecutionStarted                bool     `json:"execution_started"`
	BackendProcessStarted           bool     `json:"backend_process_started"`
	RuntimeStateWritten             bool     `json:"runtime_state_written"`
	KDEConfigurationWritten         bool     `json:"kde_configuration_written"`
	SmokeExecutedByPreview          bool     `json:"smoke_executed_by_preview"`
	NetworkRequired                 bool     `json:"network_required"`
	HostRootModified                bool     `json:"host_root_modified"`
	DockerSocketMounted             bool     `json:"docker_socket_mounted"`
	BroadHostMountRequired          bool     `json:"broad_host_mount_required"`
	PrivilegedContainerRequired     bool     `json:"privileged_container_required"`
	HostNetworkRequired             bool     `json:"host_network_required"`
	ProductionAuthorizationAccepted bool     `json:"production_authorization_accepted"`
	BlockedReason                   string   `json:"blocked_reason,omitempty"`
	BlockedActions                  []string `json:"blocked_actions"`
	NextRequirements                []string `json:"next_requirements"`
	NextHumanAuthorizedSmoke        string   `json:"next_human_authorized_smoke"`
	NextHumanAuthorizedSmokeCommand []string `json:"next_human_authorized_smoke_command"`
	DesktopSafeSummary              string   `json:"desktop_safe_summary"`
}

func PreviewDesktopTriggerDryRunRequestReview(request DesktopTriggerDryRunRequestReviewRequest) (DesktopTriggerDryRunRequestReviewPreview, error) {
	preview := baseDesktopTriggerDryRunRequestReview(request)
	handleKind, handle := desktopTriggerDryRunEvidenceHandle(request)
	preview.EvidenceHandleKind = handleKind
	preview.EvidenceHandle = handle

	envelope, envelopeErr := PreviewLaunchEnvelopeGuard(desktopTriggerDryRunEnvelopeRequest(request, handleKind, handle))
	if envelopeErr == nil {
		preview.EnvelopeGuardState = envelope.GuardState
		preview.EnvelopeAcceptedForReview = envelope.AcceptedForReview
		preview.EvidenceSHA256 = envelope.EvidenceSHA256
		preview.EvidenceDigestVerified = envelope.EvidenceDigestVerified
		preview.ExpectedDigestMatched = envelope.EvidenceDigestMatched
	} else {
		preview.EnvelopeGuardState = classifyDesktopTriggerDryRunError(envelopeErr)
	}

	action, actionErr := appidentity.PreviewKDEControlledLaunchActionSurfaceAudit(appidentity.KDEControlledLaunchActionSurfaceAuditRequest{
		DesktopEntryContent: request.DesktopEntryContent,
		EvidenceHandle:      handle,
	})
	if actionErr == nil {
		preview.ActionSurfaceState = action.AuditState
		preview.ActionSurfaceSafeForHumanSmoke = action.SurfaceSafeForHumanSmoke
		preview.KDEActionID = action.KDEActionID
		preview.PublicDBusMethod = action.PublicDBusMethod
		preview.ForwardedArgumentKind = action.ForwardedArgumentKind
		preview.KDEForwardsOnlyEvidenceHandle = action.KDEForwardsOnlyEvidenceHandle
	} else {
		preview.ActionSurfaceState = classifyDesktopTriggerDryRunError(actionErr)
	}

	acceptance, acceptanceErr := appidentity.PreviewManagedLauncherAcceptanceReport(appidentity.ManagedLauncherAcceptanceReportRequest{
		StateRoot:              request.StateRoot,
		EvidenceID:             request.EvidenceID,
		EvidenceRelativePath:   request.EvidenceRelativePath,
		ExpectedEvidenceSHA256: request.ExpectedEvidenceSHA256,
		FullCheckpointPromoted: request.FullCheckpointPromoted,
	})
	if acceptanceErr == nil {
		preview.ManagedLauncherAcceptanceState = acceptance.AcceptanceState
		preview.ManagedLauncherAccepted = acceptance.AcceptanceState == "accepted"
		preview.RuntimeStatusEvidenceState = acceptance.RuntimeStatusEvidenceState
		preview.FullCheckpointState = acceptance.FullCheckpointState
		preview.EvidenceID = acceptance.EvidenceID
		preview.EvidenceRelativePath = acceptance.EvidenceRelativePath
		if preview.EvidenceSHA256 == "" {
			preview.EvidenceSHA256 = acceptance.EvidenceSHA256
		}
		if !preview.EvidenceDigestVerified {
			preview.EvidenceDigestVerified = acceptance.EvidenceDigestVerified
		}
		preview.ExpectedDigestMatched = preview.ExpectedDigestMatched || acceptance.ExpectedDigestMatched
	} else {
		state := classifyDesktopTriggerDryRunError(acceptanceErr)
		preview.ManagedLauncherAcceptanceState = state
		preview.RuntimeStatusEvidenceState = state
	}

	preview.ReviewState = desktopTriggerDryRunOverallState(preview)
	preview.BlockedReason = desktopTriggerDryRunBlockedReason(preview)
	preview.DesktopSafeSummary = desktopTriggerDryRunSummary(preview)
	return validateDesktopTriggerDryRunRequestReviewPreview(preview)
}

func baseDesktopTriggerDryRunRequestReview(request DesktopTriggerDryRunRequestReviewRequest) DesktopTriggerDryRunRequestReviewPreview {
	fullCheckpointState := "blocked-missing-full-checkpoint"
	if request.FullCheckpointPromoted {
		fullCheckpointState = "ready"
	}
	routeID := strings.TrimSpace(request.RouteID)
	if routeID == "" {
		routeID = launchEnvelopeExpectedRoute
	}
	methodID := strings.TrimSpace(request.MethodID)
	if methodID == "" {
		methodID = launchEnvelopeExpectedMethod
	}
	actionID := strings.TrimSpace(request.ActionID)
	if actionID == "" {
		actionID = launchEnvelopeExpectedActionID
	}
	callerRole := strings.TrimSpace(request.CallerRole)
	if callerRole == "" {
		callerRole = launchEnvelopeExpectedCallerRole
	}
	return DesktopTriggerDryRunRequestReviewPreview{
		SchemaVersion:                   DesktopTriggerDryRunRequestReviewSchemaVersion,
		RequestType:                     DesktopTriggerDryRunRequestReviewRequestType,
		Source:                          "owner-service-launch-envelope-guard-preview+kde-controlled-launch-action-surface-audit-preview+managed-launcher-acceptance-report-preview",
		RuntimeMethod:                   "PreviewDesktopTriggerDryRunRequestReview",
		ReadMethod:                      "GetDesktopTriggerDryRunRequestReview",
		ReviewState:                     "blocked",
		EnvelopeGuardState:              "blocked",
		ActionSurfaceState:              "blocked",
		ManagedLauncherAcceptanceState:  "blocked",
		FullCheckpointState:             fullCheckpointState,
		RuntimeStatusEvidenceState:      "blocked",
		ExpectedEvidenceSHA256:          strings.TrimSpace(request.ExpectedEvidenceSHA256),
		RouteID:                         routeID,
		MethodID:                        methodID,
		ActionID:                        actionID,
		CallerRole:                      callerRole,
		DryRunReviewOnly:                true,
		RequestObjectWritten:            false,
		PermissionGrantCreated:          false,
		RuntimeOwned:                    true,
		GoRuntimeBacked:                 true,
		KDEPresentationOnly:             true,
		KDEForwardsOnlyEvidenceHandle:   true,
		StateRootPathExposed:            false,
		CacheRootPathExposed:            false,
		LauncherPathExposed:             false,
		RawExecutablePathExposed:        false,
		RawCommandExposed:               false,
		RawLauncherOutputExposed:        false,
		BackendDetailsExposed:           false,
		ServiceCallDispatchEnabled:      false,
		ServiceCallDispatched:           false,
		DBusCalled:                      false,
		DBusOwnershipEnabled:            false,
		DesktopLaunchEnabled:            false,
		BackendLaunchEnabled:            false,
		ExecutionStarted:                false,
		BackendProcessStarted:           false,
		RuntimeStateWritten:             false,
		KDEConfigurationWritten:         false,
		SmokeExecutedByPreview:          false,
		NetworkRequired:                 false,
		HostRootModified:                false,
		DockerSocketMounted:             false,
		BroadHostMountRequired:          false,
		PrivilegedContainerRequired:     false,
		HostNetworkRequired:             false,
		ProductionAuthorizationAccepted: false,
		BlockedActions:                  desktopTriggerDryRunBlockedActions(),
		NextRequirements:                desktopTriggerDryRunNextRequirements(),
		NextHumanAuthorizedSmoke:        "desktop-triggered-staged-launch-smoke",
		NextHumanAuthorizedSmokeCommand: []string{"ruby", "scripts/staged_launcher_dispatch_smoke.rb"},
		DesktopSafeSummary:              "Desktop-trigger dry-run request review is blocked before service dispatch.",
	}
}

func desktopTriggerDryRunEvidenceHandle(request DesktopTriggerDryRunRequestReviewRequest) (string, string) {
	if strings.TrimSpace(request.EvidenceID) != "" {
		return "evidence-id", strings.TrimSpace(request.EvidenceID)
	}
	return "evidence-relative-path", strings.TrimSpace(request.EvidenceRelativePath)
}

func desktopTriggerDryRunEnvelopeRequest(request DesktopTriggerDryRunRequestReviewRequest, handleKind string, handle string) LaunchEnvelopeGuardRequest {
	routeID := strings.TrimSpace(request.RouteID)
	if routeID == "" {
		routeID = launchEnvelopeExpectedRoute
	}
	methodID := strings.TrimSpace(request.MethodID)
	if methodID == "" {
		methodID = launchEnvelopeExpectedMethod
	}
	actionID := strings.TrimSpace(request.ActionID)
	if actionID == "" {
		actionID = launchEnvelopeExpectedActionID
	}
	callerRole := strings.TrimSpace(request.CallerRole)
	if callerRole == "" {
		callerRole = launchEnvelopeExpectedCallerRole
	}
	return LaunchEnvelopeGuardRequest{
		StateRoot:              request.StateRoot,
		RouteID:                routeID,
		MethodID:               methodID,
		EvidenceHandleKind:     handleKind,
		EvidenceHandle:         handle,
		ExpectedEvidenceSHA256: request.ExpectedEvidenceSHA256,
		ActionID:               actionID,
		CallerRole:             callerRole,
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

func desktopTriggerDryRunOverallState(preview DesktopTriggerDryRunRequestReviewPreview) string {
	if preview.FullCheckpointState != "ready" {
		return "blocked-missing-full-checkpoint"
	}
	if preview.RuntimeStatusEvidenceState == "missing-evidence" || preview.EnvelopeGuardState == "blocked-missing-evidence" || preview.ActionSurfaceState == "missing-evidence-handle" || preview.ManagedLauncherAcceptanceState == "missing-evidence" {
		return "missing-evidence"
	}
	if preview.RuntimeStatusEvidenceState == "malformed" || preview.EnvelopeGuardState == "malformed" || preview.ActionSurfaceState == "malformed" || preview.ManagedLauncherAcceptanceState == "malformed" {
		return "malformed"
	}
	if preview.RuntimeStatusEvidenceState == "stale-evidence" || preview.EnvelopeGuardState == "blocked-stale" || preview.ManagedLauncherAcceptanceState == "stale-evidence" {
		return "stale-evidence"
	}
	if preview.EnvelopeGuardState != "accepted-for-review" {
		return "blocked-unsafe-envelope"
	}
	if preview.ActionSurfaceState != "safe" {
		return "blocked-unsafe-action"
	}
	if preview.ManagedLauncherAcceptanceState != "accepted" {
		return "blocked-acceptance"
	}
	return "accepted-review"
}

func desktopTriggerDryRunBlockedReason(preview DesktopTriggerDryRunRequestReviewPreview) string {
	switch preview.ReviewState {
	case "accepted-review":
		return ""
	case "blocked-missing-full-checkpoint":
		return "desktop-trigger dry-run request review is blocked until the full checkpoint is promoted"
	case "blocked-unsafe-envelope":
		return "desktop-trigger dry-run request review is blocked by the Runtime owner launch envelope guard"
	case "blocked-unsafe-action":
		return "desktop-trigger dry-run request review is blocked by the KDE controlled-launch action surface audit"
	case "missing-evidence":
		return "desktop-trigger dry-run request review is blocked because required Runtime evidence is missing"
	case "stale-evidence":
		return "desktop-trigger dry-run request review is blocked because the evidence digest is stale"
	case "malformed":
		return "desktop-trigger dry-run request review is blocked because an evidence handle or action surface is malformed"
	default:
		return "desktop-trigger dry-run request review is blocked by incomplete managed-launcher acceptance"
	}
}

func desktopTriggerDryRunSummary(preview DesktopTriggerDryRunRequestReviewPreview) string {
	if preview.ReviewState == "accepted-review" {
		return "Desktop-trigger dry-run request is accepted for the next human-authorized staged launch smoke without dispatching a service call."
	}
	if preview.ReviewState == "blocked-missing-full-checkpoint" {
		return "Desktop-trigger dry-run request is blocked until the full checkpoint gate is promoted."
	}
	return "Desktop-trigger dry-run request is blocked before D-Bus, desktop launch, backend launch, Runtime writes, or host mutation."
}

func desktopTriggerDryRunBlockedActions() []string {
	return []string{
		"dispatch Runtime owner service call from dry-run review",
		"call D-Bus from dry-run review",
		"write Runtime request objects from dry-run review",
		"create permission grants from dry-run review",
		"write KDE configuration from dry-run review",
		"start desktop or compatibility execution from dry-run review",
		"mutate host root from dry-run review",
	}
}

func desktopTriggerDryRunNextRequirements() []string {
	return []string{
		"Promote the full checkpoint only after human-authorized full smoke passes.",
		"Keep the Runtime owner launch envelope guard in front of any desktop-triggered service call.",
		"Use the accepted dry-run review as evidence before any human-authorized desktop-triggered staged launch smoke.",
	}
}

func classifyDesktopTriggerDryRunError(err error) string {
	if err == nil {
		return "ready"
	}
	message := strings.ToLower(err.Error())
	switch {
	case strings.Contains(message, "no such file") || strings.Contains(message, "requires an evidence") || strings.Contains(message, "missing"):
		return "missing-evidence"
	case strings.Contains(message, "invalid character") || strings.Contains(message, "unexpected end of json") || strings.Contains(message, "cannot unmarshal") || strings.Contains(message, "malformed") || strings.Contains(message, "invalid"):
		return "malformed"
	case strings.Contains(message, "stale"):
		return "stale-evidence"
	case strings.Contains(message, "unsafe") || strings.Contains(message, "forbidden") || strings.Contains(message, "exposes"):
		return "unsafe"
	default:
		return "blocked"
	}
}

func validateDesktopTriggerDryRunRequestReviewPreview(preview DesktopTriggerDryRunRequestReviewPreview) (DesktopTriggerDryRunRequestReviewPreview, error) {
	validStates := map[string]bool{
		"accepted-review":                 true,
		"blocked":                         true,
		"blocked-missing-full-checkpoint": true,
		"blocked-unsafe-envelope":         true,
		"blocked-unsafe-action":           true,
		"blocked-acceptance":              true,
		"missing-evidence":                true,
		"stale-evidence":                  true,
		"malformed":                       true,
		"unsafe":                          true,
		"ready":                           true,
		"accepted":                        true,
		"needs-full-checkpoint":           true,
		"accepted-for-review":             true,
		"safe":                            true,
		"blocked-missing-evidence":        true,
		"blocked-mismatched-route":        true,
		"blocked-stale":                   true,
		"blocked-replay":                  true,
		"blocked-owner-args":              true,
		"unsupported":                     true,
		"unsafe-owner-args":               true,
		"unsafe-backend-terms":            true,
		"missing-evidence-handle":         true,
	}
	for _, state := range []string{preview.ReviewState, preview.EnvelopeGuardState, preview.ActionSurfaceState, preview.ManagedLauncherAcceptanceState, preview.FullCheckpointState, preview.RuntimeStatusEvidenceState} {
		if !validStates[state] {
			return DesktopTriggerDryRunRequestReviewPreview{}, errors.New("desktop-trigger dry-run request review has invalid state")
		}
	}
	switch {
	case preview.SchemaVersion != DesktopTriggerDryRunRequestReviewSchemaVersion || preview.RequestType != DesktopTriggerDryRunRequestReviewRequestType:
		return DesktopTriggerDryRunRequestReviewPreview{}, errors.New("desktop-trigger dry-run request review has invalid schema")
	case preview.ReviewState == "accepted-review" && (!preview.EnvelopeAcceptedForReview || !preview.ActionSurfaceSafeForHumanSmoke || !preview.ManagedLauncherAccepted || preview.FullCheckpointState != "ready"):
		return DesktopTriggerDryRunRequestReviewPreview{}, errors.New("desktop-trigger dry-run request review accepted incomplete evidence")
	case preview.ReviewState != "accepted-review" && preview.ManagedLauncherAccepted:
		if preview.FullCheckpointState == "ready" && preview.EnvelopeAcceptedForReview && preview.ActionSurfaceSafeForHumanSmoke {
			return DesktopTriggerDryRunRequestReviewPreview{}, errors.New("desktop-trigger dry-run request review blocked despite accepted dependencies")
		}
	case !preview.DryRunReviewOnly || preview.RequestObjectWritten || preview.PermissionGrantCreated || preview.ProductionAuthorizationAccepted:
		return DesktopTriggerDryRunRequestReviewPreview{}, errors.New("desktop-trigger dry-run request review must stay review-only")
	case !preview.RuntimeOwned || !preview.GoRuntimeBacked || !preview.KDEPresentationOnly || !preview.KDEForwardsOnlyEvidenceHandle:
		return DesktopTriggerDryRunRequestReviewPreview{}, errors.New("desktop-trigger dry-run request review has invalid ownership flags")
	case preview.StateRootPathExposed || preview.CacheRootPathExposed || preview.LauncherPathExposed || preview.RawExecutablePathExposed || preview.RawCommandExposed || preview.RawLauncherOutputExposed || preview.BackendDetailsExposed:
		return DesktopTriggerDryRunRequestReviewPreview{}, errors.New("desktop-trigger dry-run request review must not expose Runtime internals")
	case preview.ServiceCallDispatchEnabled || preview.ServiceCallDispatched || preview.DBusCalled || preview.DBusOwnershipEnabled || preview.DesktopLaunchEnabled || preview.BackendLaunchEnabled || preview.ExecutionStarted || preview.BackendProcessStarted || preview.RuntimeStateWritten || preview.KDEConfigurationWritten || preview.SmokeExecutedByPreview:
		return DesktopTriggerDryRunRequestReviewPreview{}, errors.New("desktop-trigger dry-run request review must not dispatch, write, or launch")
	case preview.NetworkRequired || preview.HostRootModified || preview.DockerSocketMounted || preview.BroadHostMountRequired || preview.PrivilegedContainerRequired || preview.HostNetworkRequired:
		return DesktopTriggerDryRunRequestReviewPreview{}, errors.New("desktop-trigger dry-run request review must keep host and container boundaries closed")
	}
	for _, value := range []string{
		preview.SchemaVersion, preview.RequestType, preview.Source, preview.RuntimeMethod, preview.ReadMethod, preview.ReviewState, preview.EnvelopeGuardState, preview.ActionSurfaceState, preview.ManagedLauncherAcceptanceState, preview.FullCheckpointState, preview.RuntimeStatusEvidenceState, preview.EvidenceHandleKind, preview.EvidenceHandle, preview.EvidenceID, preview.EvidenceRelativePath, preview.EvidenceSHA256, preview.ExpectedEvidenceSHA256, preview.RouteID, preview.MethodID, preview.ActionID, preview.CallerRole, preview.KDEActionID, preview.PublicDBusMethod, preview.ForwardedArgumentKind, preview.BlockedReason, preview.NextHumanAuthorizedSmoke, preview.DesktopSafeSummary,
	} {
		if value != "" && strings.ContainsAny(value, "\r\n") {
			return DesktopTriggerDryRunRequestReviewPreview{}, errors.New("desktop-trigger dry-run request review requires single-line fields")
		}
	}
	for _, value := range append(append(append([]string{}, preview.BlockedActions...), preview.NextRequirements...), preview.NextHumanAuthorizedSmokeCommand...) {
		if strings.TrimSpace(value) == "" || strings.ContainsAny(value, "\r\n") {
			return DesktopTriggerDryRunRequestReviewPreview{}, errors.New("desktop-trigger dry-run request review requires single-line list values")
		}
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		return DesktopTriggerDryRunRequestReviewPreview{}, errors.New("desktop-trigger dry-run request review must be serializable")
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
			return DesktopTriggerDryRunRequestReviewPreview{}, errors.New("desktop-trigger dry-run request review exposes forbidden terms")
		}
	}
	return preview, nil
}
