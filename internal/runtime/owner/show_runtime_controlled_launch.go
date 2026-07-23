package owner

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/winapp"
)

type ShowRuntimeControlledLaunchResult struct {
	appidentity.KnownAppKDERuntimeStatusLaunchExecutionPlan
	SchemaVersion                                   string                                                      `json:"owner_schema_version"`
	OwnerRequestType                                string                                                      `json:"owner_request_type"`
	OwnerRuntimeMethod                              string                                                      `json:"owner_runtime_method"`
	OwnerServiceBoundary                            string                                                      `json:"owner_service_boundary"`
	DesktopCallableActionID                         string                                                      `json:"desktop_callable_action_id"`
	DesktopCallableRoute                            string                                                      `json:"desktop_callable_route"`
	DesktopCallableRuntimeMethod                    string                                                      `json:"desktop_callable_runtime_method"`
	DesktopCallableExecutionType                    string                                                      `json:"desktop_callable_execution_type"`
	DesktopEvidenceHandleForwarded                  bool                                                        `json:"desktop_evidence_handle_forwarded"`
	DesktopReceiptFieldsReconstructed               bool                                                        `json:"desktop_receipt_fields_reconstructed"`
	DesktopKDEStateRootAccess                       bool                                                        `json:"desktop_kde_state_root_access"`
	DesktopRuntimeOwnerAdapterUsed                  bool                                                        `json:"desktop_runtime_owner_adapter_used"`
	DesktopStateRootSuppliedByRuntimeOwner          bool                                                        `json:"desktop_state_root_supplied_by_runtime_owner"`
	DesktopCacheRootSuppliedByRuntimeOwner          bool                                                        `json:"desktop_cache_root_supplied_by_runtime_owner"`
	DesktopLauncherSuppliedByRuntimeOwner           bool                                                        `json:"desktop_launcher_supplied_by_runtime_owner"`
	DesktopTimeoutSuppliedByRuntimeOwner            bool                                                        `json:"desktop_timeout_supplied_by_runtime_owner"`
	RuntimeOwnerServiceActionDispatch               bool                                                        `json:"runtime_owner_service_action_dispatch"`
	RuntimeOwnerServiceCallReady                    bool                                                        `json:"runtime_owner_service_call_ready"`
	RuntimeOwnerServiceSuppliesOwnerInputs          bool                                                        `json:"runtime_owner_service_supplies_owner_inputs"`
	KDEForwardsOnlyEvidenceHandle                   bool                                                        `json:"kde_forwards_only_evidence_handle"`
	WriteMethodsEnabled                             bool                                                        `json:"write_methods_enabled"`
	DispatchReady                                   bool                                                        `json:"dispatch_ready"`
	CompatibilityCenterKnownAppEvidence             appidentity.KnownAppKDERuntimeStatusLaunchDelegatedEvidence `json:"compatibility_center_known_app_evidence"`
	ManagedLauncherName                             string                                                      `json:"managed_launcher_name"`
	ManagedLauncherInvoked                          bool                                                        `json:"managed_launcher_invoked"`
	ExistingManagedLauncherInvoked                  bool                                                        `json:"existing_managed_launcher_invoked"`
	LauncherExitCode                                int                                                         `json:"launcher_exit_code"`
	LauncherOutputJSONObserved                      bool                                                        `json:"launcher_output_json_observed"`
	DelegatedRequestType                            string                                                      `json:"delegated_request_type"`
	DelegatedStatus                                 string                                                      `json:"delegated_status"`
	DelegatedSkipReason                             string                                                      `json:"delegated_skip_reason,omitempty"`
	DelegatedFailureReason                          string                                                      `json:"delegated_failure_reason,omitempty"`
	DelegatedGuestBoundary                          string                                                      `json:"delegated_guest_boundary"`
	DelegatedRuntimeOwnedDispatch                   bool                                                        `json:"delegated_runtime_owned_dispatch"`
	DelegatedArtifactVerified                       bool                                                        `json:"delegated_artifact_verified"`
	DelegatedMarkerObserved                         bool                                                        `json:"delegated_marker_observed"`
	DelegatedSmokePassed                            bool                                                        `json:"delegated_smoke_passed"`
	DelegatedExecutionStarted                       bool                                                        `json:"delegated_execution_started"`
	DelegatedBackendProcessStarted                  bool                                                        `json:"delegated_backend_process_started"`
	DelegatedSessionGatedControlledDispatchConsumed bool                                                        `json:"delegated_session_gated_controlled_dispatch_consumed"`
	DelegatedSessionGatedControlledDispatchState    string                                                      `json:"delegated_session_gated_controlled_dispatch_state"`
	DelegatedSessionGatedReviewReceiptID            string                                                      `json:"delegated_session_gated_review_receipt_id"`
	DelegatedLaunchAuthorizationReceiptID           string                                                      `json:"delegated_launch_authorization_receipt_id"`
	DelegatedControlledExecutionSessionConsumed     bool                                                        `json:"delegated_controlled_execution_session_consumed"`
	DelegatedControlledExecutionSessionID           string                                                      `json:"delegated_controlled_execution_session_id"`
	DelegatedHostRootModified                       bool                                                        `json:"delegated_host_root_modified"`
	DelegatedDockerSocketMounted                    bool                                                        `json:"delegated_docker_socket_mounted"`
	DelegatedBroadHostMountRequired                 bool                                                        `json:"delegated_broad_host_mount_required"`
	DelegatedRawCommandExposed                      bool                                                        `json:"delegated_raw_command_exposed"`
	DelegatedBackendDetailsExposed                  bool                                                        `json:"delegated_backend_details_exposed"`
}

type runtimeControlledLaunchOwnerConfig struct {
	StateRoot    string
	CacheRoot    string
	LauncherPath string
	OwnerTimeout string
	GuestTimeout string
	GuestKeyPath string
}

func (service Service) ShowRuntimeControlledLaunch(args []string) (ShowRuntimeControlledLaunchResult, error) {
	evidenceID, evidenceRelativePath, err := parseShowRuntimeControlledLaunchArgs(args)
	if err != nil {
		return ShowRuntimeControlledLaunchResult{}, err
	}
	config, err := runtimeControlledLaunchOwnerConfigFromEnv()
	if err != nil {
		return ShowRuntimeControlledLaunchResult{}, err
	}
	plan, err := appidentity.PrepareKnownAppKDERuntimeStatusLaunchExecutionFromActionTrigger(appidentity.KnownAppKDERuntimeStatusLaunchExecutionFromActionTriggerRequest{
		StateRoot:            config.StateRoot,
		CacheRoot:            config.CacheRoot,
		EvidenceID:           evidenceID,
		EvidenceRelativePath: evidenceRelativePath,
	})
	if err != nil {
		return ShowRuntimeControlledLaunchResult{}, err
	}
	extra, err := runtimeControlledLaunchExtraArgs(config)
	if err != nil {
		return ShowRuntimeControlledLaunchResult{}, err
	}
	launcherArgs, err := appidentity.KnownAppKDERuntimeStatusLaunchExecutionArgv(plan, config.StateRoot, config.CacheRoot, extra)
	if err != nil {
		return ShowRuntimeControlledLaunchResult{}, err
	}
	ownerTimeout, err := time.ParseDuration(config.OwnerTimeout)
	if err != nil {
		return ShowRuntimeControlledLaunchResult{}, fmt.Errorf("parse Runtime owner action timeout: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), ownerTimeout)
	defer cancel()
	output, err := exec.CommandContext(ctx, config.LauncherPath, launcherArgs...).Output()
	if ctx.Err() != nil {
		return ShowRuntimeControlledLaunchResult{}, errors.New("Runtime owner action managed launcher invocation timed out")
	}
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return ShowRuntimeControlledLaunchResult{}, fmt.Errorf("Runtime owner action managed launcher failed with exit code %d", exitErr.ExitCode())
		}
		return ShowRuntimeControlledLaunchResult{}, errors.New("Runtime owner action managed launcher invocation failed")
	}
	return showRuntimeControlledLaunchResultFromOutput(plan, config.LauncherPath, output)
}

func parseShowRuntimeControlledLaunchArgs(args []string) (string, string, error) {
	if len(args) == 2 {
		switch args[0] {
		case "evidence-id":
			if strings.TrimSpace(args[1]) == "" {
				return "", "", errors.New("ShowRuntimeControlledLaunch requires a non-empty evidence id")
			}
			return args[1], "", nil
		case "evidence-relative-path":
			if strings.TrimSpace(args[1]) == "" {
				return "", "", errors.New("ShowRuntimeControlledLaunch requires a non-empty evidence relative path")
			}
			return "", args[1], nil
		}
	}
	flags := flag.NewFlagSet("ShowRuntimeControlledLaunch", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	evidenceID := flags.String("evidence-id", "", "opaque Runtime-status launch evidence handoff id")
	evidenceRelativePath := flags.String("evidence-relative-path", "", "relative Runtime-status launch evidence handoff path")
	if err := flags.Parse(args); err != nil {
		return "", "", err
	}
	if flags.NArg() != 0 {
		return "", "", errors.New("ShowRuntimeControlledLaunch does not accept positional arguments")
	}
	if strings.TrimSpace(*evidenceID) == "" && strings.TrimSpace(*evidenceRelativePath) == "" {
		return "", "", errors.New("ShowRuntimeControlledLaunch requires --evidence-id or --evidence-relative-path")
	}
	return *evidenceID, *evidenceRelativePath, nil
}

func runtimeControlledLaunchOwnerConfigFromEnv() (runtimeControlledLaunchOwnerConfig, error) {
	config := runtimeControlledLaunchOwnerConfig{
		StateRoot:    strings.TrimSpace(os.Getenv("XNIX_RUNTIME_OWNER_STATE_ROOT")),
		CacheRoot:    strings.TrimSpace(os.Getenv("XNIX_RUNTIME_OWNER_KNOWN_APP_CACHE_ROOT")),
		LauncherPath: strings.TrimSpace(os.Getenv("XNIX_RUNTIME_OWNER_MANAGED_LAUNCHER")),
		OwnerTimeout: strings.TrimSpace(os.Getenv("XNIX_RUNTIME_OWNER_TIMEOUT")),
		GuestTimeout: strings.TrimSpace(os.Getenv("XNIX_RUNTIME_OWNER_GUEST_TIMEOUT")),
		GuestKeyPath: strings.TrimSpace(os.Getenv("XNIX_RUNTIME_OWNER_GUEST_KEY")),
	}
	if config.StateRoot == "" {
		return runtimeControlledLaunchOwnerConfig{}, errors.New("ShowRuntimeControlledLaunch requires XNIX_RUNTIME_OWNER_STATE_ROOT from the Runtime owner service boundary")
	}
	if config.CacheRoot == "" {
		config.CacheRoot = winapp.DefaultKnownAppCacheRoot
	}
	if config.LauncherPath == "" {
		config.LauncherPath = "xnix-compat-launch"
	}
	if config.OwnerTimeout == "" {
		config.OwnerTimeout = "5m"
	}
	for _, value := range []string{config.StateRoot, config.CacheRoot, config.LauncherPath, config.OwnerTimeout, config.GuestTimeout, config.GuestKeyPath} {
		if strings.ContainsAny(value, "\r\n") {
			return runtimeControlledLaunchOwnerConfig{}, errors.New("ShowRuntimeControlledLaunch Runtime owner configuration requires single-line values")
		}
	}
	return config, nil
}

func runtimeControlledLaunchExtraArgs(config runtimeControlledLaunchOwnerConfig) ([]string, error) {
	var args []string
	if config.GuestKeyPath != "" {
		args = append(args, "--key", config.GuestKeyPath)
	}
	if config.GuestTimeout != "" {
		args = append(args, "--timeout", config.GuestTimeout)
	}
	for _, value := range args {
		if strings.ContainsAny(value, "\r\n") {
			return nil, errors.New("ShowRuntimeControlledLaunch extra launcher arguments must be single-line")
		}
	}
	return args, nil
}

func showRuntimeControlledLaunchResultFromOutput(plan appidentity.KnownAppKDERuntimeStatusLaunchExecutionPlan, launcherPath string, output []byte) (ShowRuntimeControlledLaunchResult, error) {
	var delegated map[string]any
	if err := json.Unmarshal(output, &delegated); err != nil {
		return ShowRuntimeControlledLaunchResult{}, errors.New("Runtime owner action managed launcher must emit JSON")
	}
	result := ShowRuntimeControlledLaunchResult{
		KnownAppKDERuntimeStatusLaunchExecutionPlan: plan,
		SchemaVersion:                                   "xnix.runtime.owner_show_runtime_controlled_launch.v1",
		OwnerRequestType:                                "runtime-owner-show-runtime-controlled-launch",
		OwnerRuntimeMethod:                              "ShowRuntimeControlledLaunch",
		OwnerServiceBoundary:                            "go-runtime-owner-in-process-service",
		DesktopCallableActionID:                         appidentity.KnownAppKDERuntimeStatusLaunchAction,
		DesktopCallableRoute:                            "kde-dbus-runtime-status-action",
		DesktopCallableRuntimeMethod:                    "ShowRuntimeControlledLaunch",
		DesktopCallableExecutionType:                    appidentity.KnownAppKDERuntimeStatusLaunchExecutionRequestType,
		DesktopEvidenceHandleForwarded:                  true,
		DesktopReceiptFieldsReconstructed:               false,
		DesktopKDEStateRootAccess:                       false,
		DesktopRuntimeOwnerAdapterUsed:                  true,
		DesktopStateRootSuppliedByRuntimeOwner:          true,
		DesktopCacheRootSuppliedByRuntimeOwner:          true,
		DesktopLauncherSuppliedByRuntimeOwner:           true,
		DesktopTimeoutSuppliedByRuntimeOwner:            true,
		RuntimeOwnerServiceActionDispatch:               true,
		RuntimeOwnerServiceCallReady:                    true,
		RuntimeOwnerServiceSuppliesOwnerInputs:          true,
		KDEForwardsOnlyEvidenceHandle:                   true,
		WriteMethodsEnabled:                             false,
		DispatchReady:                                   true,
		ManagedLauncherName:                             filepath.Base(strings.TrimSpace(launcherPath)),
		ManagedLauncherInvoked:                          true,
		ExistingManagedLauncherInvoked:                  true,
		LauncherExitCode:                                0,
		LauncherOutputJSONObserved:                      true,
		DelegatedRequestType:                            ownerStringJSONField(delegated, "request_type"),
		DelegatedStatus:                                 ownerStringJSONField(delegated, "status"),
		DelegatedSkipReason:                             ownerStringJSONField(delegated, "skip_reason"),
		DelegatedFailureReason:                          ownerStringJSONField(delegated, "failure_reason"),
		DelegatedGuestBoundary:                          ownerStringJSONField(delegated, "guest_boundary"),
		DelegatedRuntimeOwnedDispatch:                   ownerBoolJSONField(delegated, "runtime_owned_dispatch"),
		DelegatedArtifactVerified:                       ownerBoolJSONField(delegated, "artifact_verified"),
		DelegatedMarkerObserved:                         ownerBoolJSONField(delegated, "marker_observed"),
		DelegatedSmokePassed:                            ownerBoolJSONField(delegated, "smoke_passed"),
		DelegatedExecutionStarted:                       ownerBoolJSONField(delegated, "execution_started"),
		DelegatedBackendProcessStarted:                  ownerBoolJSONField(delegated, "backend_process_started"),
		DelegatedSessionGatedControlledDispatchConsumed: ownerBoolJSONField(delegated, "session_gated_controlled_dispatch_consumed"),
		DelegatedSessionGatedControlledDispatchState:    ownerStringJSONField(delegated, "session_gated_controlled_dispatch_state"),
		DelegatedSessionGatedReviewReceiptID:            ownerStringJSONField(delegated, "session_gated_review_receipt_id"),
		DelegatedLaunchAuthorizationReceiptID:           ownerStringJSONField(delegated, "launch_authorization_receipt_id"),
		DelegatedControlledExecutionSessionConsumed:     ownerBoolJSONField(delegated, "controlled_execution_session_consumed"),
		DelegatedControlledExecutionSessionID:           ownerStringJSONField(delegated, "controlled_execution_session_id"),
		DelegatedHostRootModified:                       ownerBoolJSONField(delegated, "host_root_modified"),
		DelegatedDockerSocketMounted:                    ownerBoolJSONField(delegated, "docker_socket_mounted"),
		DelegatedBroadHostMountRequired:                 ownerBoolJSONField(delegated, "broad_host_mount_required"),
		DelegatedRawCommandExposed:                      ownerBoolJSONField(delegated, "raw_command_exposed"),
		DelegatedBackendDetailsExposed:                  ownerBoolJSONField(delegated, "backend_details_exposed"),
	}
	if result.ManagedLauncherName == "" || strings.ContainsAny(result.ManagedLauncherName, "\r\n") {
		return ShowRuntimeControlledLaunchResult{}, errors.New("Runtime owner action managed launcher name is unsafe")
	}
	if result.DelegatedRequestType != winapp.KnownDispatchSmokeRequestType {
		return ShowRuntimeControlledLaunchResult{}, errors.New("Runtime owner action managed launcher returned an unexpected delegated request type")
	}
	if result.DelegatedHostRootModified || result.DelegatedDockerSocketMounted || result.DelegatedBroadHostMountRequired || result.DelegatedRawCommandExposed || result.DelegatedBackendDetailsExposed {
		return ShowRuntimeControlledLaunchResult{}, errors.New("Runtime owner action managed launcher reported an unsafe delegated result")
	}
	projection, err := appidentity.ProjectKnownAppKDERuntimeStatusLaunchDelegatedEvidence(appidentity.KnownAppKDERuntimeStatusLaunchDelegatedEvidenceRequest{
		AppID:                                  plan.AppID,
		DisplayName:                            plan.DisplayName,
		AppVersion:                             plan.AppVersion,
		RequestType:                            result.DelegatedRequestType,
		Status:                                 result.DelegatedStatus,
		SkipReason:                             result.DelegatedSkipReason,
		FailureReason:                          result.DelegatedFailureReason,
		GuestBoundary:                          result.DelegatedGuestBoundary,
		RuntimeOwnedDispatch:                   result.DelegatedRuntimeOwnedDispatch,
		ArtifactVerified:                       result.DelegatedArtifactVerified,
		MarkerObserved:                         result.DelegatedMarkerObserved,
		SessionGatedControlledDispatchConsumed: ownerBoolJSONField(delegated, "session_gated_controlled_dispatch_consumed"),
		SessionGatedControlledDispatchState:    ownerStringJSONField(delegated, "session_gated_controlled_dispatch_state"),
		SessionGatedReviewReceiptID:            ownerStringJSONField(delegated, "session_gated_review_receipt_id"),
		LaunchAuthorizationReceiptID:           ownerStringJSONField(delegated, "launch_authorization_receipt_id"),
		LaunchAuthorizationReceiptState:        "recorded",
		LaunchGateState:                        "controlled-dispatch-ready",
		LaunchGateConsumed:                     plan.LaunchGateConsumed,
		LaunchGateReceiptAccepted:              plan.LaunchGateReceiptAccepted,
		LaunchGateGuestBoundaryAccepted:        plan.LaunchGateGuestBoundaryAccepted,
		LaunchGateBlockedReason:                plan.LaunchGateBlockedReason,
		ControlledDispatchReady:                true,
		ControlledExecutionSessionConsumed:     ownerBoolJSONField(delegated, "controlled_execution_session_consumed"),
		ControlledExecutionSessionID:           ownerStringJSONField(delegated, "controlled_execution_session_id"),
		ControlledSessionDigestVerified:        ownerBoolJSONField(delegated, "controlled_session_digest_verified"),
		ControlledSessionRelativePath:          ownerStringJSONField(delegated, "controlled_session_relative_path"),
		RuntimeOwnerConsumableSession:          ownerBoolJSONField(delegated, "runtime_owner_consumable_session"),
		KDEReadModelConsumableSession:          ownerBoolJSONField(delegated, "kde_read_model_consumable_session"),
		ControlledSessionLiveStateObserved:     ownerBoolJSONField(delegated, "controlled_session_live_state_observed"),
		ControlledSessionRegistered:            ownerBoolJSONField(delegated, "controlled_session_registered"),
		ControlledSessionWindowObserved:        ownerBoolJSONField(delegated, "controlled_session_window_observed"),
		ControlledSessionHostRootModified:      ownerBoolJSONField(delegated, "controlled_session_host_root_modified"),
		ControlledSessionBackendProcessStart:   ownerBoolJSONField(delegated, "controlled_session_backend_process_start"),
		HostRootModified:                       result.DelegatedHostRootModified,
		DockerSocketMounted:                    result.DelegatedDockerSocketMounted,
		BroadHostMountRequired:                 result.DelegatedBroadHostMountRequired,
	})
	if err != nil {
		return ShowRuntimeControlledLaunchResult{}, err
	}
	result.CompatibilityCenterKnownAppEvidence = projection
	if err := validateNoBackendTerms(result, "Runtime owner show Runtime-controlled launch result"); err != nil {
		return ShowRuntimeControlledLaunchResult{}, err
	}
	return result, nil
}

func ownerStringJSONField(payload map[string]any, key string) string {
	value, _ := payload[key].(string)
	return value
}

func ownerBoolJSONField(payload map[string]any, key string) bool {
	value, _ := payload[key].(bool)
	return value
}
