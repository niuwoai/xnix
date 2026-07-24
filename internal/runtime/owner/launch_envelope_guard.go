package owner

import (
	"errors"
	"path/filepath"
	"strings"

	"xnix.local/xnix/internal/runtime/appidentity"
)

const (
	LaunchEnvelopeGuardSchemaVersion = "xnix.runtime.owner_service_launch_envelope_guard.v1"
	LaunchEnvelopeGuardRequestType   = "owner-service-launch-envelope-guard-preview"
)

const (
	launchEnvelopeExpectedRoute      = "kde-dbus-runtime-status-action"
	launchEnvelopeExpectedMethod     = "ShowRuntimeControlledLaunch"
	launchEnvelopeExpectedActionID   = "xnix.runtime-status.controlled-launch"
	launchEnvelopeExpectedCallerRole = "kde-presentation-shell"
)

type LaunchEnvelopeGuardRequest struct {
	StateRoot              string
	RouteID                string
	MethodID               string
	EvidenceHandleKind     string
	EvidenceHandle         string
	ExpectedEvidenceSHA256 string
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

type LaunchEnvelopeGuardPreview struct {
	SchemaVersion                     string   `json:"schema_version"`
	RequestType                       string   `json:"request_type"`
	Source                            string   `json:"source"`
	RuntimeMethod                     string   `json:"runtime_method"`
	ReadMethod                        string   `json:"read_method"`
	GuardState                        string   `json:"guard_state"`
	RouteID                           string   `json:"route_id"`
	MethodID                          string   `json:"method_id"`
	EvidenceHandleKind                string   `json:"evidence_handle_kind"`
	EvidenceHandle                    string   `json:"evidence_handle,omitempty"`
	ExpectedEvidenceSHA256            string   `json:"expected_evidence_sha256,omitempty"`
	EvidenceSHA256                    string   `json:"evidence_sha256,omitempty"`
	EvidenceDigestVerified            bool     `json:"evidence_digest_verified"`
	EvidenceDigestMatched             bool     `json:"evidence_digest_matched"`
	ExpectedActionID                  string   `json:"expected_action_id"`
	CallerRole                        string   `json:"caller_role"`
	RequestFreshness                  string   `json:"request_freshness"`
	ReplayMarkerState                 string   `json:"replay_marker_state"`
	RouteMatched                      bool     `json:"route_matched"`
	MethodMatched                     bool     `json:"method_matched"`
	ActionMatched                     bool     `json:"action_matched"`
	CallerRoleAccepted                bool     `json:"caller_role_accepted"`
	EvidenceHandleShapeAccepted       bool     `json:"evidence_handle_shape_accepted"`
	OwnerOnlyArgumentsPresent         bool     `json:"owner_only_arguments_present"`
	OwnerOnlyArgumentCount            int      `json:"owner_only_argument_count"`
	KDECanOnlyForwardEvidenceHandle   bool     `json:"kde_can_only_forward_evidence_handle"`
	KDEStateRootRejected              bool     `json:"kde_state_root_rejected"`
	KDECacheRootRejected              bool     `json:"kde_cache_root_rejected"`
	KDELauncherPathRejected           bool     `json:"kde_launcher_path_rejected"`
	KDETimeoutValueRejected           bool     `json:"kde_timeout_value_rejected"`
	KDERawExecutablePathRejected      bool     `json:"kde_raw_executable_path_rejected"`
	KDEBackendCommandRejected         bool     `json:"kde_backend_command_rejected"`
	KDEReceiptFieldsRejected          bool     `json:"kde_receipt_fields_rejected"`
	KDESessionFieldsRejected          bool     `json:"kde_session_fields_rejected"`
	KDEDispatchFieldsRejected         bool     `json:"kde_dispatch_fields_rejected"`
	RuntimeOwned                      bool     `json:"runtime_owned"`
	GoRuntimeBacked                   bool     `json:"go_runtime_backed"`
	KDEPolicyOwner                    bool     `json:"kde_policy_owner"`
	DesktopReceiptFieldsReconstructed bool     `json:"desktop_receipt_fields_reconstructed"`
	DesktopKDEStateRootAccess         bool     `json:"desktop_kde_state_root_access"`
	StateRootPathExposed              bool     `json:"state_root_path_exposed"`
	CacheRootPathExposed              bool     `json:"cache_root_path_exposed"`
	LauncherPathExposed               bool     `json:"launcher_path_exposed"`
	RawExecutablePathExposed          bool     `json:"raw_executable_path_exposed"`
	RawCommandExposed                 bool     `json:"raw_command_exposed"`
	BackendDetailsExposed             bool     `json:"backend_details_exposed"`
	ReceiptWriteEnabled               bool     `json:"receipt_write_enabled"`
	ReceiptAccepted                   bool     `json:"receipt_accepted"`
	ProductionAuthorizationAccepted   bool     `json:"production_authorization_accepted"`
	ServiceCallDispatched             bool     `json:"service_call_dispatched"`
	DBusOwnershipEnabled              bool     `json:"dbus_ownership_enabled"`
	DesktopLaunchEnabled              bool     `json:"desktop_launch_enabled"`
	BackendLaunchEnabled              bool     `json:"backend_launch_enabled"`
	ExecutionStarted                  bool     `json:"execution_started"`
	BackendProcessStarted             bool     `json:"backend_process_started"`
	RuntimeStateWritten               bool     `json:"runtime_state_written"`
	NetworkRequired                   bool     `json:"network_required"`
	HostRootModified                  bool     `json:"host_root_modified"`
	PrivilegedContainerRequired       bool     `json:"privileged_container_required"`
	BlockedReason                     string   `json:"blocked_reason,omitempty"`
	AcceptedForReview                 bool     `json:"accepted_for_review"`
	BlockedActions                    []string `json:"blocked_actions"`
	NextRequirements                  []string `json:"next_requirements"`
	DesktopSafeSummary                string   `json:"desktop_safe_summary"`
}

func PreviewLaunchEnvelopeGuard(request LaunchEnvelopeGuardRequest) (LaunchEnvelopeGuardPreview, error) {
	preview := baseLaunchEnvelopeGuardPreview(request)
	ownerOnlyArgumentCount := launchEnvelopeOwnerOnlyArgumentCount(request)
	preview.OwnerOnlyArgumentCount = ownerOnlyArgumentCount
	preview.OwnerOnlyArgumentsPresent = ownerOnlyArgumentCount > 0
	preview.KDEStateRootRejected = strings.TrimSpace(request.KDEStateRoot) != ""
	preview.KDECacheRootRejected = strings.TrimSpace(request.KDECacheRoot) != ""
	preview.KDELauncherPathRejected = strings.TrimSpace(request.KDELauncherPath) != ""
	preview.KDETimeoutValueRejected = strings.TrimSpace(request.KDETimeoutValue) != ""
	preview.KDERawExecutablePathRejected = strings.TrimSpace(request.KDERawExecutablePath) != ""
	preview.KDEBackendCommandRejected = strings.TrimSpace(request.KDEBackendCommand) != ""
	preview.KDEReceiptFieldsRejected = strings.TrimSpace(request.KDEReceiptID) != ""
	preview.KDESessionFieldsRejected = strings.TrimSpace(request.KDESessionID) != ""
	preview.KDEDispatchFieldsRejected = strings.TrimSpace(request.KDEDispatchID) != ""

	switch {
	case !preview.RouteMatched || !preview.MethodMatched || !preview.ActionMatched:
		preview.GuardState = "blocked-mismatched-route"
		preview.BlockedReason = "desktop launch envelope route, method, or action does not match the Runtime-controlled launch route"
	case !preview.CallerRoleAccepted:
		preview.GuardState = "unsupported"
		preview.BlockedReason = "desktop launch envelope caller role is not supported"
	case !preview.EvidenceHandleShapeAccepted:
		preview.GuardState = "malformed"
		preview.BlockedReason = "desktop launch envelope must contain exactly one safe evidence handle"
	case preview.OwnerOnlyArgumentsPresent:
		preview.GuardState = "blocked-owner-args"
		preview.BlockedReason = "desktop launch envelope contains owner-only arguments"
	case preview.ReplayMarkerState != "fresh":
		preview.GuardState = "blocked-replay"
		preview.BlockedReason = "desktop launch envelope replay marker is not fresh"
	case preview.RequestFreshness != "fresh":
		preview.GuardState = "blocked-stale"
		preview.BlockedReason = "desktop launch envelope freshness state is stale"
	default:
		preview = bindLaunchEnvelopeEvidence(preview, request)
	}
	preview.KDECanOnlyForwardEvidenceHandle = preview.EvidenceHandleShapeAccepted && !preview.OwnerOnlyArgumentsPresent
	preview.AcceptedForReview = preview.GuardState == "accepted-for-review"
	preview.DesktopSafeSummary = launchEnvelopeGuardSummary(preview)
	return validateLaunchEnvelopeGuardPreview(preview)
}

func baseLaunchEnvelopeGuardPreview(request LaunchEnvelopeGuardRequest) LaunchEnvelopeGuardPreview {
	routeID := strings.TrimSpace(request.RouteID)
	methodID := strings.TrimSpace(request.MethodID)
	evidenceHandleKind := strings.TrimSpace(request.EvidenceHandleKind)
	actionID := strings.TrimSpace(request.ActionID)
	callerRole := strings.TrimSpace(request.CallerRole)
	requestFreshness := strings.TrimSpace(request.RequestFreshness)
	replayMarkerState := strings.TrimSpace(request.ReplayMarker)
	if requestFreshness == "" {
		requestFreshness = "fresh"
	}
	if replayMarkerState == "" {
		replayMarkerState = "fresh"
	}
	return LaunchEnvelopeGuardPreview{
		SchemaVersion:                     LaunchEnvelopeGuardSchemaVersion,
		RequestType:                       LaunchEnvelopeGuardRequestType,
		Source:                            "kde-controlled-launch-action+runtime-owner-service-envelope",
		RuntimeMethod:                     "PreviewLaunchEnvelopeGuard",
		ReadMethod:                        "GetLaunchEnvelopeGuardPreview",
		GuardState:                        "blocked-missing-evidence",
		RouteID:                           routeID,
		MethodID:                          methodID,
		EvidenceHandleKind:                evidenceHandleKind,
		EvidenceHandle:                    strings.TrimSpace(request.EvidenceHandle),
		ExpectedEvidenceSHA256:            strings.TrimSpace(request.ExpectedEvidenceSHA256),
		ExpectedActionID:                  actionID,
		CallerRole:                        callerRole,
		RequestFreshness:                  requestFreshness,
		ReplayMarkerState:                 replayMarkerState,
		RouteMatched:                      routeID == launchEnvelopeExpectedRoute,
		MethodMatched:                     methodID == launchEnvelopeExpectedMethod,
		ActionMatched:                     actionID == launchEnvelopeExpectedActionID,
		CallerRoleAccepted:                callerRole == launchEnvelopeExpectedCallerRole,
		EvidenceHandleShapeAccepted:       launchEnvelopeEvidenceHandleShapeAccepted(evidenceHandleKind, strings.TrimSpace(request.EvidenceHandle)),
		RuntimeOwned:                      true,
		GoRuntimeBacked:                   true,
		KDEPolicyOwner:                    false,
		DesktopReceiptFieldsReconstructed: false,
		DesktopKDEStateRootAccess:         false,
		StateRootPathExposed:              false,
		CacheRootPathExposed:              false,
		LauncherPathExposed:               false,
		RawExecutablePathExposed:          false,
		RawCommandExposed:                 false,
		BackendDetailsExposed:             false,
		ReceiptWriteEnabled:               false,
		ReceiptAccepted:                   false,
		ProductionAuthorizationAccepted:   false,
		ServiceCallDispatched:             false,
		DBusOwnershipEnabled:              false,
		DesktopLaunchEnabled:              false,
		BackendLaunchEnabled:              false,
		ExecutionStarted:                  false,
		BackendProcessStarted:             false,
		RuntimeStateWritten:               false,
		NetworkRequired:                   false,
		HostRootModified:                  false,
		PrivilegedContainerRequired:       false,
		BlockedActions:                    launchEnvelopeBlockedActions(),
		NextRequirements:                  launchEnvelopeNextRequirements(),
	}
}

func bindLaunchEnvelopeEvidence(preview LaunchEnvelopeGuardPreview, request LaunchEnvelopeGuardRequest) LaunchEnvelopeGuardPreview {
	if strings.TrimSpace(request.StateRoot) == "" {
		preview.GuardState = "blocked-missing-evidence"
		preview.BlockedReason = "Runtime owner state root is required internally to bind envelope evidence"
		return preview
	}
	evidenceRequest := appidentity.KnownAppKDERuntimeStatusLaunchEvidencePreviewRequest{StateRoot: request.StateRoot}
	switch preview.EvidenceHandleKind {
	case "evidence-id":
		evidenceRequest.EvidenceID = preview.EvidenceHandle
	case "evidence-relative-path":
		evidenceRequest.EvidenceRelativePath = preview.EvidenceHandle
	default:
		preview.GuardState = "malformed"
		preview.BlockedReason = "unsupported evidence handle kind"
		return preview
	}
	evidence, err := appidentity.PreviewKnownAppKDERuntimeStatusLaunchEvidence(evidenceRequest)
	if err != nil {
		preview.GuardState = classifyLaunchEnvelopeEvidenceError(err)
		preview.BlockedReason = "Runtime-status launch evidence could not be bound to the envelope"
		return preview
	}
	preview.EvidenceSHA256 = evidence.EvidenceSHA256
	preview.EvidenceDigestVerified = evidence.EvidenceDigestVerified
	preview.EvidenceDigestMatched = preview.ExpectedEvidenceSHA256 == "" || preview.ExpectedEvidenceSHA256 == evidence.EvidenceSHA256
	if !preview.EvidenceDigestMatched {
		preview.GuardState = "blocked-stale"
		preview.BlockedReason = "desktop launch envelope evidence digest is stale"
		return preview
	}
	preview.GuardState = "accepted-for-review"
	return preview
}

func launchEnvelopeEvidenceHandleShapeAccepted(kind string, handle string) bool {
	if strings.TrimSpace(handle) == "" {
		return false
	}
	switch kind {
	case "evidence-id":
		return !strings.ContainsAny(handle, "/\\ \t\r\n")
	case "evidence-relative-path":
		clean := filepath.Clean(filepath.ToSlash(handle))
		return !filepath.IsAbs(handle) && !strings.Contains(clean, "..") && strings.HasPrefix(filepath.ToSlash(handle), "runtime/kde-runtime-status-launch-evidence/") && strings.HasSuffix(handle, ".json")
	default:
		return false
	}
}

func launchEnvelopeOwnerOnlyArgumentCount(request LaunchEnvelopeGuardRequest) int {
	count := 0
	for _, value := range []string{
		request.KDEStateRoot,
		request.KDECacheRoot,
		request.KDELauncherPath,
		request.KDETimeoutValue,
		request.KDERawExecutablePath,
		request.KDEBackendCommand,
		request.KDEReceiptID,
		request.KDESessionID,
		request.KDEDispatchID,
	} {
		if strings.TrimSpace(value) != "" {
			count++
		}
	}
	return count
}

func classifyLaunchEnvelopeEvidenceError(err error) string {
	if err == nil {
		return "accepted-for-review"
	}
	message := strings.ToLower(err.Error())
	switch {
	case strings.Contains(message, "no such file") || strings.Contains(message, "requires an evidence id") || strings.Contains(message, "requires a runtime-status launch evidence relative path"):
		return "blocked-missing-evidence"
	case strings.Contains(message, "invalid character") || strings.Contains(message, "unexpected end of json") || strings.Contains(message, "cannot unmarshal") || strings.Contains(message, "invalid schema") || strings.Contains(message, "invalid projection"):
		return "malformed"
	case strings.Contains(message, "unsafe") || strings.Contains(message, "exposes forbidden"):
		return "malformed"
	default:
		return "blocked-missing-evidence"
	}
}

func launchEnvelopeGuardSummary(preview LaunchEnvelopeGuardPreview) string {
	if preview.GuardState == "accepted-for-review" {
		return "Runtime owner service launch envelope is accepted for review with evidence-only desktop forwarding."
	}
	if preview.GuardState == "blocked-owner-args" {
		return "Runtime owner service launch envelope is blocked because desktop input included owner-only arguments."
	}
	if preview.GuardState == "blocked-stale" {
		return "Runtime owner service launch envelope is blocked because its freshness or evidence digest is stale."
	}
	if preview.GuardState == "blocked-replay" {
		return "Runtime owner service launch envelope is blocked because the replay marker is not fresh."
	}
	return "Runtime owner service launch envelope is blocked before any service call dispatch."
}

func launchEnvelopeBlockedActions() []string {
	return []string{
		"dispatch Runtime owner service call from guard preview",
		"accept production authorization from guard preview",
		"write launch receipts from guard preview",
		"let KDE supply owner-only launch inputs",
		"enable production D-Bus ownership from guard preview",
		"start desktop or compatibility execution from guard preview",
		"mutate host root from guard preview",
	}
}

func launchEnvelopeNextRequirements() []string {
	return []string{
		"Keep this guard in front of future desktop-triggered owner service dispatch.",
		"Bind accepted envelopes to fresh opaque evidence handles before real desktop-triggered smoke.",
		"Keep owner-only state root, cache root, launcher, timeout, receipt, session, and dispatch inputs inside Runtime.",
	}
}

func validateLaunchEnvelopeGuardPreview(preview LaunchEnvelopeGuardPreview) (LaunchEnvelopeGuardPreview, error) {
	validStates := map[string]bool{
		"accepted-for-review":      true,
		"blocked-missing-evidence": true,
		"blocked-mismatched-route": true,
		"blocked-stale":            true,
		"blocked-replay":           true,
		"blocked-owner-args":       true,
		"malformed":                true,
		"unsupported":              true,
	}
	switch {
	case preview.SchemaVersion != LaunchEnvelopeGuardSchemaVersion || preview.RequestType != LaunchEnvelopeGuardRequestType:
		return LaunchEnvelopeGuardPreview{}, errors.New("Runtime owner launch envelope guard has invalid schema")
	case !validStates[preview.GuardState]:
		return LaunchEnvelopeGuardPreview{}, errors.New("Runtime owner launch envelope guard has invalid state")
	case preview.AcceptedForReview != (preview.GuardState == "accepted-for-review"):
		return LaunchEnvelopeGuardPreview{}, errors.New("Runtime owner launch envelope guard acceptance flag is inconsistent")
	case preview.AcceptedForReview && (!preview.RouteMatched || !preview.MethodMatched || !preview.ActionMatched || !preview.CallerRoleAccepted || !preview.EvidenceHandleShapeAccepted || !preview.EvidenceDigestVerified || !preview.EvidenceDigestMatched):
		return LaunchEnvelopeGuardPreview{}, errors.New("Runtime owner launch envelope guard accepted an incomplete envelope")
	case preview.AcceptedForReview && preview.OwnerOnlyArgumentsPresent:
		return LaunchEnvelopeGuardPreview{}, errors.New("Runtime owner launch envelope guard accepted owner-only arguments")
	case preview.GuardState != "malformed" && preview.EvidenceHandle != "" && preview.EvidenceHandleKind == "evidence-relative-path" && !launchEnvelopeEvidenceHandleShapeAccepted(preview.EvidenceHandleKind, preview.EvidenceHandle):
		return LaunchEnvelopeGuardPreview{}, errors.New("Runtime owner launch envelope guard requires safe relative evidence")
	case preview.OwnerOnlyArgumentsPresent && preview.OwnerOnlyArgumentCount == 0:
		return LaunchEnvelopeGuardPreview{}, errors.New("Runtime owner launch envelope guard owner-only count is inconsistent")
	case preview.KDEPolicyOwner || preview.DesktopReceiptFieldsReconstructed || preview.DesktopKDEStateRootAccess || preview.StateRootPathExposed || preview.CacheRootPathExposed || preview.LauncherPathExposed || preview.RawExecutablePathExposed || preview.RawCommandExposed || preview.BackendDetailsExposed:
		return LaunchEnvelopeGuardPreview{}, errors.New("Runtime owner launch envelope guard must not expose Runtime internals")
	case preview.ReceiptWriteEnabled || preview.ReceiptAccepted || preview.ProductionAuthorizationAccepted || preview.ServiceCallDispatched || preview.DBusOwnershipEnabled || preview.DesktopLaunchEnabled || preview.BackendLaunchEnabled || preview.ExecutionStarted || preview.BackendProcessStarted || preview.RuntimeStateWritten:
		return LaunchEnvelopeGuardPreview{}, errors.New("Runtime owner launch envelope guard must not dispatch, authorize, write, or launch")
	case preview.NetworkRequired || preview.HostRootModified || preview.PrivilegedContainerRequired:
		return LaunchEnvelopeGuardPreview{}, errors.New("Runtime owner launch envelope guard must keep host boundaries closed")
	}
	for _, value := range []string{preview.SchemaVersion, preview.RequestType, preview.Source, preview.RuntimeMethod, preview.ReadMethod, preview.GuardState, preview.RouteID, preview.MethodID, preview.EvidenceHandleKind, preview.EvidenceHandle, preview.ExpectedEvidenceSHA256, preview.EvidenceSHA256, preview.ExpectedActionID, preview.CallerRole, preview.RequestFreshness, preview.ReplayMarkerState, preview.BlockedReason, preview.DesktopSafeSummary} {
		if value != "" && strings.ContainsAny(value, "\r\n") {
			return LaunchEnvelopeGuardPreview{}, errors.New("Runtime owner launch envelope guard requires single-line fields")
		}
	}
	for _, value := range append(append([]string{}, preview.BlockedActions...), preview.NextRequirements...) {
		if strings.TrimSpace(value) == "" || strings.ContainsAny(value, "\r\n") {
			return LaunchEnvelopeGuardPreview{}, errors.New("Runtime owner launch envelope guard requires single-line list values")
		}
	}
	if err := validateNoBackendTerms(preview, "Runtime owner launch envelope guard"); err != nil {
		return LaunchEnvelopeGuardPreview{}, err
	}
	return preview, nil
}
