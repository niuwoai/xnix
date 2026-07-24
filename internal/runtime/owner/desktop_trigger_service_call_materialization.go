package owner

import (
	"encoding/json"
	"errors"
	"strings"

	"xnix.local/xnix/internal/runtime/appidentity"
)

const (
	DesktopTriggerServiceCallMaterializationSchemaVersion = "xnix.runtime.desktop_trigger_service_call_materialization.v1"
	DesktopTriggerServiceCallMaterializationRequestType   = "desktop-trigger-service-call-materialization-preview"
)

type DesktopTriggerServiceCallMaterializationRequest struct {
	StateRoot              string
	DesktopEntryContent    string
	EvidenceID             string
	EvidenceRelativePath   string
	ExpectedEvidenceSHA256 string
	FullCheckpointPromoted bool
	HumanAuthorizedSmoke   bool
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

type DesktopTriggerServiceCallMaterializationPreview struct {
	SchemaVersion                     string   `json:"schema_version"`
	RequestType                       string   `json:"request_type"`
	Source                            string   `json:"source"`
	RuntimeMethod                     string   `json:"runtime_method"`
	ReadMethod                        string   `json:"read_method"`
	MaterializationState              string   `json:"materialization_state"`
	DryRunReviewState                 string   `json:"dry_run_review_state"`
	OwnerTriggerState                 string   `json:"owner_trigger_state"`
	FullCheckpointState               string   `json:"full_checkpoint_state"`
	RuntimeStatusEvidenceState        string   `json:"runtime_status_evidence_state"`
	EvidenceID                        string   `json:"evidence_id,omitempty"`
	EvidenceRelativePath              string   `json:"evidence_relative_path,omitempty"`
	EvidenceSHA256                    string   `json:"evidence_sha256,omitempty"`
	ExpectedEvidenceSHA256            string   `json:"expected_evidence_sha256,omitempty"`
	EvidenceDigestVerified            bool     `json:"evidence_digest_verified"`
	ExpectedDigestMatched             bool     `json:"expected_digest_matched"`
	HumanAuthorizedSmoke              bool     `json:"human_authorized_smoke"`
	FullCheckpointPromotionClaimed    bool     `json:"full_checkpoint_promotion_claimed"`
	FormalReleaseReady                bool     `json:"formal_release_ready"`
	DesktopCallableRoute              string   `json:"desktop_callable_route"`
	DesktopCallableRuntimeMethod      string   `json:"desktop_callable_runtime_method"`
	DesktopCallableExecutionType      string   `json:"desktop_callable_execution_type"`
	DesktopDBusMethod                 string   `json:"desktop_dbus_method"`
	OwnerServiceBoundary              string   `json:"owner_service_boundary"`
	OwnerServiceMethod                string   `json:"owner_service_method"`
	OwnerServiceCallType              string   `json:"owner_service_call_type"`
	OwnerServiceCallArgs              []string `json:"owner_service_call_args,omitempty"`
	OwnerServiceCLIArgs               []string `json:"owner_service_cli_args,omitempty"`
	OwnerServiceCallReady             bool     `json:"owner_service_call_ready"`
	HumanAuthorizationRequired        bool     `json:"human_authorization_required"`
	MaterializedForHumanSmoke         bool     `json:"materialized_for_human_smoke"`
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
	NextHumanAuthorizedSmoke          string   `json:"next_human_authorized_smoke"`
	NextHumanAuthorizedSmokeCommand   []string `json:"next_human_authorized_smoke_command"`
	DesktopSafeSummary                string   `json:"desktop_safe_summary"`
}

func PreviewDesktopTriggerServiceCallMaterialization(request DesktopTriggerServiceCallMaterializationRequest) (DesktopTriggerServiceCallMaterializationPreview, error) {
	preview := baseDesktopTriggerServiceCallMaterialization(request)
	review, reviewErr := PreviewDesktopTriggerDryRunRequestReview(desktopTriggerServiceCallReviewRequest(request))
	if reviewErr != nil {
		state := classifyDesktopTriggerDryRunError(reviewErr)
		preview.MaterializationState = state
		preview.DryRunReviewState = state
		preview.OwnerTriggerState = "blocked"
		preview.BlockedReason = "desktop-trigger service call materialization is blocked because dry-run review failed"
		return validateDesktopTriggerServiceCallMaterializationPreview(preview)
	}

	preview.DryRunReviewState = review.ReviewState
	preview.FullCheckpointState = review.FullCheckpointState
	preview.RuntimeStatusEvidenceState = review.RuntimeStatusEvidenceState
	preview.EvidenceID = review.EvidenceID
	preview.EvidenceRelativePath = review.EvidenceRelativePath
	preview.EvidenceSHA256 = review.EvidenceSHA256
	preview.EvidenceDigestVerified = review.EvidenceDigestVerified
	preview.ExpectedDigestMatched = review.ExpectedDigestMatched
	if !desktopTriggerServiceCallReviewAllowsMaterialization(review, request) {
		preview.MaterializationState = review.ReviewState
		preview.OwnerTriggerState = "blocked"
		preview.BlockedReason = "desktop-trigger service call materialization requires an accepted dry-run review"
		preview.DesktopSafeSummary = desktopTriggerServiceCallMaterializationSummary(preview)
		return validateDesktopTriggerServiceCallMaterializationPreview(preview)
	}

	trigger, triggerErr := appidentity.PreviewKnownAppRuntimeStatusLaunchOwnerTrigger(appidentity.KnownAppRuntimeStatusLaunchOwnerTriggerRequest{
		StateRoot:            request.StateRoot,
		EvidenceID:           request.EvidenceID,
		EvidenceRelativePath: request.EvidenceRelativePath,
	})
	if triggerErr != nil {
		state := classifyDesktopTriggerDryRunError(triggerErr)
		preview.MaterializationState = state
		preview.OwnerTriggerState = state
		preview.BlockedReason = "desktop-trigger service call materialization could not derive owner service call arguments"
		preview.DesktopSafeSummary = desktopTriggerServiceCallMaterializationSummary(preview)
		return validateDesktopTriggerServiceCallMaterializationPreview(preview)
	}

	preview.OwnerTriggerState = "ready"
	preview.EvidenceID = trigger.EvidenceID
	preview.EvidenceRelativePath = trigger.EvidenceRelativePath
	preview.EvidenceSHA256 = trigger.EvidenceSHA256
	preview.EvidenceDigestVerified = trigger.EvidenceDigestVerified
	preview.DesktopCallableRoute = trigger.DesktopCallableRoute
	preview.DesktopCallableRuntimeMethod = trigger.DesktopCallableRuntimeMethod
	preview.DesktopCallableExecutionType = trigger.DesktopCallableExecutionType
	preview.DesktopDBusMethod = trigger.DesktopDBusMethod
	preview.OwnerServiceBoundary = trigger.OwnerServiceBoundary
	preview.OwnerServiceMethod = trigger.OwnerServiceMethod
	preview.OwnerServiceCallType = trigger.OwnerServiceCallType
	preview.OwnerServiceCallArgs = append([]string{}, trigger.OwnerServiceCallArgs...)
	preview.OwnerServiceCLIArgs = append([]string{}, trigger.OwnerServiceCLIArgs...)
	preview.OwnerServiceCallReady = trigger.OwnerServiceCallReady
	preview.MaterializedForHumanSmoke = true
	preview.RuntimeOwnerServiceSuppliesInputs = trigger.RuntimeOwnerServiceSuppliesInputs
	preview.DesktopEvidenceHandleForwarded = trigger.DesktopEvidenceHandleForwarded
	preview.RuntimeOwnedDispatch = trigger.RuntimeOwnedDispatch
	preview.MaterializationState = "ready-for-human-authorized-service-call"
	preview.DesktopSafeSummary = desktopTriggerServiceCallMaterializationSummary(preview)
	return validateDesktopTriggerServiceCallMaterializationPreview(preview)
}

func baseDesktopTriggerServiceCallMaterialization(request DesktopTriggerServiceCallMaterializationRequest) DesktopTriggerServiceCallMaterializationPreview {
	fullCheckpointState := "blocked-missing-full-checkpoint"
	if request.FullCheckpointPromoted {
		fullCheckpointState = "ready"
	}
	return DesktopTriggerServiceCallMaterializationPreview{
		SchemaVersion:                     DesktopTriggerServiceCallMaterializationSchemaVersion,
		RequestType:                       DesktopTriggerServiceCallMaterializationRequestType,
		Source:                            DesktopTriggerDryRunRequestReviewRequestType + "+known-app-runtime-status-launch-owner-trigger-preview",
		RuntimeMethod:                     "PreviewDesktopTriggerServiceCallMaterialization",
		ReadMethod:                        "GetDesktopTriggerServiceCallMaterialization",
		MaterializationState:              "blocked",
		DryRunReviewState:                 "blocked",
		OwnerTriggerState:                 "blocked",
		FullCheckpointState:               fullCheckpointState,
		RuntimeStatusEvidenceState:        "blocked",
		ExpectedEvidenceSHA256:            strings.TrimSpace(request.ExpectedEvidenceSHA256),
		HumanAuthorizedSmoke:              request.HumanAuthorizedSmoke,
		FullCheckpointPromotionClaimed:    request.FullCheckpointPromoted,
		FormalReleaseReady:                false,
		DesktopCallableRoute:              "kde-dbus-runtime-status-action",
		DesktopCallableRuntimeMethod:      "ShowRuntimeControlledLaunch",
		DesktopCallableExecutionType:      appidentity.KnownAppKDERuntimeStatusLaunchExecutionRequestType,
		DesktopDBusMethod:                 "org.xnix.Compatibility1.ShowRuntimeControlledLaunch",
		OwnerServiceBoundary:              "go-runtime-owner-in-process-service",
		OwnerServiceMethod:                "ShowRuntimeControlledLaunch",
		OwnerServiceCallType:              "desktop-action-dispatch",
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
		BlockedActions:                    desktopTriggerServiceCallMaterializationBlockedActions(),
		NextRequirements:                  desktopTriggerServiceCallMaterializationNextRequirements(),
		NextHumanAuthorizedSmoke:          "desktop-triggered-staged-launch-smoke",
		NextHumanAuthorizedSmokeCommand:   []string{"ruby", "scripts/staged_launcher_dispatch_smoke.rb"},
		DesktopSafeSummary:                "Desktop-trigger service call materialization is blocked before owner service dispatch.",
	}
}

func desktopTriggerServiceCallReviewRequest(request DesktopTriggerServiceCallMaterializationRequest) DesktopTriggerDryRunRequestReviewRequest {
	return DesktopTriggerDryRunRequestReviewRequest{
		StateRoot:              request.StateRoot,
		DesktopEntryContent:    request.DesktopEntryContent,
		EvidenceID:             request.EvidenceID,
		EvidenceRelativePath:   request.EvidenceRelativePath,
		ExpectedEvidenceSHA256: request.ExpectedEvidenceSHA256,
		FullCheckpointPromoted: request.FullCheckpointPromoted,
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

func desktopTriggerServiceCallReviewAllowsMaterialization(review DesktopTriggerDryRunRequestReviewPreview, request DesktopTriggerServiceCallMaterializationRequest) bool {
	if review.ReviewState == "accepted-review" {
		return true
	}
	if !request.HumanAuthorizedSmoke || review.ReviewState != "blocked-missing-full-checkpoint" {
		return false
	}
	return review.EnvelopeGuardState == "accepted-for-review" &&
		review.ActionSurfaceState == "safe" &&
		review.ManagedLauncherAcceptanceState == "needs-full-checkpoint" &&
		review.FullCheckpointState == "needs-full-checkpoint" &&
		review.RuntimeStatusEvidenceState == "ready" &&
		review.EvidenceDigestVerified &&
		review.ExpectedDigestMatched
}

func desktopTriggerServiceCallMaterializationSummary(preview DesktopTriggerServiceCallMaterializationPreview) string {
	if preview.MaterializationState == "ready-for-human-authorized-service-call" {
		if preview.HumanAuthorizedSmoke && !preview.FullCheckpointPromotionClaimed {
			return "Desktop-trigger owner service call arguments are materialized for a human-authorized smoke candidate without claiming formal full-checkpoint promotion."
		}
		return "Desktop-trigger owner service call arguments are materialized for the next human-authorized staged launch smoke without dispatching the service call."
	}
	if preview.MaterializationState == "blocked-missing-full-checkpoint" {
		return "Desktop-trigger owner service call materialization is blocked until the full checkpoint gate is promoted."
	}
	return "Desktop-trigger owner service call materialization is blocked before D-Bus, owner service dispatch, desktop launch, backend launch, Runtime writes, or host mutation."
}

func desktopTriggerServiceCallMaterializationBlockedActions() []string {
	return []string{
		"dispatch Runtime owner service call from materialization preview",
		"call D-Bus from materialization preview",
		"write Runtime request objects from materialization preview",
		"create permission grants from materialization preview",
		"write KDE configuration from materialization preview",
		"start desktop or compatibility execution from materialization preview",
		"mutate host root from materialization preview",
	}
}

func desktopTriggerServiceCallMaterializationNextRequirements() []string {
	return []string{
		"Run the formal full checkpoint before promoting v0.2.640.",
		"Use the materialized owner service call only inside a human-authorized desktop-triggered staged launch smoke.",
		"Keep Runtime owner environment values outside KDE and outside this preview output.",
	}
}

func validateDesktopTriggerServiceCallMaterializationPreview(preview DesktopTriggerServiceCallMaterializationPreview) (DesktopTriggerServiceCallMaterializationPreview, error) {
	validStates := map[string]bool{
		"ready-for-human-authorized-service-call": true,
		"accepted-review":                         true,
		"blocked":                                 true,
		"blocked-missing-full-checkpoint":         true,
		"blocked-unsafe-envelope":                 true,
		"blocked-unsafe-action":                   true,
		"blocked-acceptance":                      true,
		"missing-evidence":                        true,
		"stale-evidence":                          true,
		"malformed":                               true,
		"unsafe":                                  true,
		"ready":                                   true,
		"needs-full-checkpoint":                   true,
	}
	for _, state := range []string{preview.MaterializationState, preview.DryRunReviewState, preview.OwnerTriggerState, preview.FullCheckpointState, preview.RuntimeStatusEvidenceState} {
		if !validStates[state] {
			return DesktopTriggerServiceCallMaterializationPreview{}, errors.New("desktop-trigger service call materialization has invalid state")
		}
	}
	reviewAllowsReady := preview.DryRunReviewState == "accepted-review" ||
		(preview.HumanAuthorizedSmoke && preview.DryRunReviewState == "blocked-missing-full-checkpoint")
	checkpointAllowsReady := preview.FullCheckpointState == "ready" ||
		(preview.HumanAuthorizedSmoke && !preview.FullCheckpointPromotionClaimed && preview.FullCheckpointState == "needs-full-checkpoint")
	switch {
	case preview.SchemaVersion != DesktopTriggerServiceCallMaterializationSchemaVersion || preview.RequestType != DesktopTriggerServiceCallMaterializationRequestType:
		return DesktopTriggerServiceCallMaterializationPreview{}, errors.New("desktop-trigger service call materialization has invalid schema")
	case preview.MaterializationState == "ready-for-human-authorized-service-call" && (!preview.OwnerServiceCallReady || !preview.MaterializedForHumanSmoke || !reviewAllowsReady || preview.OwnerTriggerState != "ready" || !checkpointAllowsReady):
		return DesktopTriggerServiceCallMaterializationPreview{}, errors.New("desktop-trigger service call materialization accepted incomplete dependencies")
	case preview.MaterializationState != "ready-for-human-authorized-service-call" && (preview.OwnerServiceCallReady || preview.MaterializedForHumanSmoke || len(preview.OwnerServiceCallArgs) > 0 || len(preview.OwnerServiceCLIArgs) > 0):
		return DesktopTriggerServiceCallMaterializationPreview{}, errors.New("desktop-trigger service call materialization must not emit service call args while blocked")
	case preview.FullCheckpointPromotionClaimed && preview.FullCheckpointState != "ready":
		return DesktopTriggerServiceCallMaterializationPreview{}, errors.New("desktop-trigger service call materialization claimed checkpoint promotion without a ready checkpoint")
	case preview.FormalReleaseReady:
		return DesktopTriggerServiceCallMaterializationPreview{}, errors.New("desktop-trigger service call materialization must not claim formal release readiness")
	case preview.MaterializationState == "ready-for-human-authorized-service-call" && (!sameDesktopTriggerServiceCallMaterializationArgs(preview.OwnerServiceCallArgs, []string{"ShowRuntimeControlledLaunch", "evidence-relative-path", preview.EvidenceRelativePath}) || !sameDesktopTriggerServiceCallMaterializationArgs(preview.OwnerServiceCLIArgs, []string{"--service-call", "ShowRuntimeControlledLaunch", "evidence-relative-path", preview.EvidenceRelativePath})):
		return DesktopTriggerServiceCallMaterializationPreview{}, errors.New("desktop-trigger service call materialization requires evidence-only service call args")
	case !preview.HumanAuthorizationRequired || !preview.RuntimeOwnerServiceSuppliesInputs || !preview.DesktopEvidenceHandleForwarded:
		return DesktopTriggerServiceCallMaterializationPreview{}, errors.New("desktop-trigger service call materialization has invalid authorization or trigger flags")
	case !preview.RuntimeOwned || !preview.RuntimeOwnedDispatch || !preview.GoRuntimeBacked || !preview.KDEPresentationOnly || !preview.KDEForwardsOnlyEvidenceHandle || preview.KDEReceivesMaterializedOwnerArgs:
		return DesktopTriggerServiceCallMaterializationPreview{}, errors.New("desktop-trigger service call materialization has invalid ownership flags")
	case preview.DesktopReceiptFieldsReconstructed || preview.DesktopKDEStateRootAccess || preview.StateRootPathExposed || preview.CacheRootPathExposed || preview.LauncherPathExposed || preview.RawExecutablePathExposed || preview.RawCommandExposed || preview.RawLauncherOutputExposed || preview.BackendDetailsExposed:
		return DesktopTriggerServiceCallMaterializationPreview{}, errors.New("desktop-trigger service call materialization must not expose Runtime internals")
	case preview.RequestObjectWritten || preview.PermissionGrantCreated || preview.ServiceCallDispatchEnabled || preview.ServiceCallDispatched || preview.DBusCalled || preview.DBusOwnershipEnabled || preview.DesktopLaunchEnabled || preview.BackendLaunchEnabled || preview.ExecutionStarted || preview.BackendProcessStarted || preview.RuntimeStateWritten || preview.KDEConfigurationWritten || preview.SmokeExecutedByPreview:
		return DesktopTriggerServiceCallMaterializationPreview{}, errors.New("desktop-trigger service call materialization must not dispatch, write, or launch")
	case preview.NetworkRequired || preview.HostRootModified || preview.DockerSocketMounted || preview.BroadHostMountRequired || preview.PrivilegedContainerRequired || preview.HostNetworkRequired || preview.ProductionAuthorizationAccepted:
		return DesktopTriggerServiceCallMaterializationPreview{}, errors.New("desktop-trigger service call materialization must keep host and container boundaries closed")
	}
	for _, value := range []string{
		preview.SchemaVersion, preview.RequestType, preview.Source, preview.RuntimeMethod, preview.ReadMethod, preview.MaterializationState, preview.DryRunReviewState, preview.OwnerTriggerState, preview.FullCheckpointState, preview.RuntimeStatusEvidenceState, preview.EvidenceID, preview.EvidenceRelativePath, preview.EvidenceSHA256, preview.ExpectedEvidenceSHA256, preview.DesktopCallableRoute, preview.DesktopCallableRuntimeMethod, preview.DesktopCallableExecutionType, preview.DesktopDBusMethod, preview.OwnerServiceBoundary, preview.OwnerServiceMethod, preview.OwnerServiceCallType, preview.BlockedReason, preview.NextHumanAuthorizedSmoke, preview.DesktopSafeSummary,
	} {
		if value != "" && strings.ContainsAny(value, "\r\n") {
			return DesktopTriggerServiceCallMaterializationPreview{}, errors.New("desktop-trigger service call materialization requires single-line fields")
		}
	}
	for _, value := range append(append(append(append([]string{}, preview.OwnerServiceCallArgs...), preview.OwnerServiceCLIArgs...), preview.BlockedActions...), append(preview.NextRequirements, preview.NextHumanAuthorizedSmokeCommand...)...) {
		if strings.TrimSpace(value) == "" || strings.ContainsAny(value, "\r\n") {
			return DesktopTriggerServiceCallMaterializationPreview{}, errors.New("desktop-trigger service call materialization requires single-line list values")
		}
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		return DesktopTriggerServiceCallMaterializationPreview{}, errors.New("desktop-trigger service call materialization must be serializable")
	}
	lower := strings.ToLower(string(encoded))
	for _, forbidden := range []string{
		"xnix_runtime_owner_",
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
			return DesktopTriggerServiceCallMaterializationPreview{}, errors.New("desktop-trigger service call materialization exposes forbidden terms")
		}
	}
	return preview, nil
}

func sameDesktopTriggerServiceCallMaterializationArgs(left []string, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}
