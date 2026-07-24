package appidentity

import (
	"errors"
	"strings"
)

const (
	KnownAppRuntimeStatusLaunchOwnerTriggerSchemaVersion = "xnix.runtime.known_app_runtime_status_launch_owner_trigger.v1"
	KnownAppRuntimeStatusLaunchOwnerTriggerRequestType   = "known-app-runtime-status-launch-owner-trigger-preview"
)

type KnownAppRuntimeStatusLaunchOwnerTriggerRequest struct {
	StateRoot            string
	EvidenceID           string
	EvidenceRelativePath string
}

type KnownAppRuntimeStatusLaunchOwnerTriggerPreview struct {
	SchemaVersion                     string   `json:"schema_version"`
	RequestType                       string   `json:"request_type"`
	Source                            string   `json:"source"`
	RuntimeMethod                     string   `json:"runtime_method"`
	ReadMethod                        string   `json:"read_method"`
	AppID                             string   `json:"app_id"`
	DisplayName                       string   `json:"display_name"`
	AppVersion                        string   `json:"app_version"`
	EvidenceID                        string   `json:"evidence_id"`
	EvidenceRelativePath              string   `json:"evidence_relative_path"`
	EvidenceSHA256                    string   `json:"evidence_sha256"`
	EvidenceHandoffConsumed           bool     `json:"evidence_handoff_consumed"`
	EvidenceDigestVerified            bool     `json:"evidence_digest_verified"`
	DesktopTriggerReady               bool     `json:"desktop_trigger_ready"`
	DesktopCallableRoute              string   `json:"desktop_callable_route"`
	DesktopCallableRuntimeMethod      string   `json:"desktop_callable_runtime_method"`
	DesktopCallableExecutionType      string   `json:"desktop_callable_execution_type"`
	DesktopEvidenceHandleForwarded    bool     `json:"desktop_evidence_handle_forwarded"`
	DesktopDBusMethod                 string   `json:"desktop_dbus_method"`
	OwnerServiceCallReady             bool     `json:"owner_service_call_ready"`
	OwnerServiceBoundary              string   `json:"owner_service_boundary"`
	OwnerServiceMethod                string   `json:"owner_service_method"`
	OwnerServiceCallType              string   `json:"owner_service_call_type"`
	OwnerServiceCallArgs              []string `json:"owner_service_call_args"`
	OwnerServiceCLIArgs               []string `json:"owner_service_cli_args"`
	RuntimeOwnerServiceSuppliesInputs bool     `json:"runtime_owner_service_supplies_inputs"`
	RuntimeOwned                      bool     `json:"runtime_owned"`
	RuntimeOwnedDispatch              bool     `json:"runtime_owned_dispatch"`
	GoRuntimeBacked                   bool     `json:"go_runtime_backed"`
	KDEPolicyOwner                    bool     `json:"kde_policy_owner"`
	KDEForwardsOnlyEvidenceHandle     bool     `json:"kde_forwards_only_evidence_handle"`
	DesktopReceiptFieldsReconstructed bool     `json:"desktop_receipt_fields_reconstructed"`
	DesktopKDEStateRootAccess         bool     `json:"desktop_kde_state_root_access"`
	StateRootPathExposed              bool     `json:"state_root_path_exposed"`
	EvidencePathExposed               bool     `json:"evidence_path_exposed"`
	ManagedLauncherPathExposed        bool     `json:"managed_launcher_path_exposed"`
	RawLauncherOutputExposed          bool     `json:"raw_launcher_output_exposed"`
	BackendDetailsExposed             bool     `json:"backend_details_exposed"`
	HostRootModified                  bool     `json:"host_root_modified"`
	DockerSocketMounted               bool     `json:"docker_socket_mounted"`
	BroadHostMountRequired            bool     `json:"broad_host_mount_required"`
	DesktopLaunchEnabled              bool     `json:"desktop_launch_enabled"`
	BackendLaunchEnabled              bool     `json:"backend_launch_enabled"`
	ExecutionStarted                  bool     `json:"execution_started"`
	BackendProcessStarted             bool     `json:"backend_process_started"`
	DesktopSafeSummary                string   `json:"desktop_safe_summary"`
}

func PreviewKnownAppRuntimeStatusLaunchOwnerTrigger(request KnownAppRuntimeStatusLaunchOwnerTriggerRequest) (KnownAppRuntimeStatusLaunchOwnerTriggerPreview, error) {
	evidence, err := PreviewKnownAppKDERuntimeStatusLaunchEvidence(KnownAppKDERuntimeStatusLaunchEvidencePreviewRequest{
		StateRoot:            request.StateRoot,
		EvidenceID:           request.EvidenceID,
		EvidenceRelativePath: request.EvidenceRelativePath,
	})
	if err != nil {
		return KnownAppRuntimeStatusLaunchOwnerTriggerPreview{}, err
	}
	preview := KnownAppRuntimeStatusLaunchOwnerTriggerPreview{
		SchemaVersion:                     KnownAppRuntimeStatusLaunchOwnerTriggerSchemaVersion,
		RequestType:                       KnownAppRuntimeStatusLaunchOwnerTriggerRequestType,
		Source:                            KnownAppKDERuntimeStatusLaunchEvidencePreviewRequestType + "+owner-service-desktop-trigger",
		RuntimeMethod:                     "PreviewKnownAppRuntimeStatusLaunchOwnerTrigger",
		ReadMethod:                        "GetKnownAppRuntimeStatusLaunchOwnerTrigger",
		AppID:                             evidence.AppID,
		DisplayName:                       evidence.DisplayName,
		AppVersion:                        evidence.AppVersion,
		EvidenceID:                        evidence.EvidenceID,
		EvidenceRelativePath:              evidence.EvidenceRelativePath,
		EvidenceSHA256:                    evidence.EvidenceSHA256,
		EvidenceHandoffConsumed:           evidence.EvidenceHandoffConsumed,
		EvidenceDigestVerified:            evidence.EvidenceDigestVerified,
		DesktopTriggerReady:               true,
		DesktopCallableRoute:              "kde-dbus-runtime-status-action",
		DesktopCallableRuntimeMethod:      "ShowRuntimeControlledLaunch",
		DesktopCallableExecutionType:      KnownAppKDERuntimeStatusLaunchExecutionRequestType,
		DesktopEvidenceHandleForwarded:    true,
		DesktopDBusMethod:                 "org.xnix.Compatibility1.ShowRuntimeControlledLaunch",
		OwnerServiceCallReady:             true,
		OwnerServiceBoundary:              "go-runtime-owner-in-process-service",
		OwnerServiceMethod:                "ShowRuntimeControlledLaunch",
		OwnerServiceCallType:              "desktop-action-dispatch",
		OwnerServiceCallArgs:              []string{"ShowRuntimeControlledLaunch", "evidence-relative-path", evidence.EvidenceRelativePath},
		OwnerServiceCLIArgs:               []string{"--service-call", "ShowRuntimeControlledLaunch", "evidence-relative-path", evidence.EvidenceRelativePath},
		RuntimeOwnerServiceSuppliesInputs: true,
		RuntimeOwned:                      true,
		RuntimeOwnedDispatch:              evidence.RuntimeOwnedDispatch,
		GoRuntimeBacked:                   true,
		KDEPolicyOwner:                    false,
		KDEForwardsOnlyEvidenceHandle:     true,
		DesktopReceiptFieldsReconstructed: false,
		DesktopKDEStateRootAccess:         false,
		StateRootPathExposed:              false,
		EvidencePathExposed:               false,
		ManagedLauncherPathExposed:        false,
		RawLauncherOutputExposed:          false,
		BackendDetailsExposed:             false,
		HostRootModified:                  false,
		DockerSocketMounted:               false,
		BroadHostMountRequired:            false,
		DesktopLaunchEnabled:              false,
		BackendLaunchEnabled:              false,
		ExecutionStarted:                  false,
		BackendProcessStarted:             false,
		DesktopSafeSummary:                evidence.DisplayName + " Runtime-status launch handoff can be forwarded by the desktop to the Go Runtime owner service without reconstructing launch receipts or exposing Runtime paths.",
	}
	return validateKnownAppRuntimeStatusLaunchOwnerTriggerPreview(preview)
}

func validateKnownAppRuntimeStatusLaunchOwnerTriggerPreview(preview KnownAppRuntimeStatusLaunchOwnerTriggerPreview) (KnownAppRuntimeStatusLaunchOwnerTriggerPreview, error) {
	switch {
	case preview.SchemaVersion != KnownAppRuntimeStatusLaunchOwnerTriggerSchemaVersion || preview.RequestType != KnownAppRuntimeStatusLaunchOwnerTriggerRequestType:
		return KnownAppRuntimeStatusLaunchOwnerTriggerPreview{}, errors.New("known app Runtime-status launch owner trigger has invalid schema")
	case preview.AppID == "" || preview.DisplayName == "" || preview.AppVersion == "":
		return KnownAppRuntimeStatusLaunchOwnerTriggerPreview{}, errors.New("known app Runtime-status launch owner trigger requires known app identity")
	case preview.EvidenceRelativePath == "" || preview.EvidenceSHA256 == "" || !preview.EvidenceHandoffConsumed || !preview.EvidenceDigestVerified:
		return KnownAppRuntimeStatusLaunchOwnerTriggerPreview{}, errors.New("known app Runtime-status launch owner trigger requires verified handoff evidence")
	case !preview.DesktopTriggerReady || preview.DesktopCallableRoute != "kde-dbus-runtime-status-action" || preview.DesktopCallableRuntimeMethod != "ShowRuntimeControlledLaunch" || preview.DesktopCallableExecutionType != KnownAppKDERuntimeStatusLaunchExecutionRequestType || !preview.DesktopEvidenceHandleForwarded || preview.DesktopDBusMethod != "org.xnix.Compatibility1.ShowRuntimeControlledLaunch":
		return KnownAppRuntimeStatusLaunchOwnerTriggerPreview{}, errors.New("known app Runtime-status launch owner trigger requires a desktop-callable D-Bus trigger")
	case !preview.OwnerServiceCallReady || preview.OwnerServiceBoundary != "go-runtime-owner-in-process-service" || preview.OwnerServiceMethod != "ShowRuntimeControlledLaunch" || preview.OwnerServiceCallType != "desktop-action-dispatch" || !preview.RuntimeOwnerServiceSuppliesInputs:
		return KnownAppRuntimeStatusLaunchOwnerTriggerPreview{}, errors.New("known app Runtime-status launch owner trigger requires a Runtime owner service bridge")
	case !sameRuntimeStatusLaunchOwnerFixtureArgs(preview.OwnerServiceCallArgs, []string{"ShowRuntimeControlledLaunch", "evidence-relative-path", preview.EvidenceRelativePath}):
		return KnownAppRuntimeStatusLaunchOwnerTriggerPreview{}, errors.New("known app Runtime-status launch owner trigger requires evidence-only owner service arguments")
	case !sameRuntimeStatusLaunchOwnerFixtureArgs(preview.OwnerServiceCLIArgs, []string{"--service-call", "ShowRuntimeControlledLaunch", "evidence-relative-path", preview.EvidenceRelativePath}):
		return KnownAppRuntimeStatusLaunchOwnerTriggerPreview{}, errors.New("known app Runtime-status launch owner trigger requires evidence-only owner service CLI arguments")
	case !preview.RuntimeOwned || !preview.RuntimeOwnedDispatch || !preview.GoRuntimeBacked || preview.KDEPolicyOwner:
		return KnownAppRuntimeStatusLaunchOwnerTriggerPreview{}, errors.New("known app Runtime-status launch owner trigger must remain Runtime-owned and Go-backed")
	case !preview.KDEForwardsOnlyEvidenceHandle || preview.DesktopReceiptFieldsReconstructed || preview.DesktopKDEStateRootAccess:
		return KnownAppRuntimeStatusLaunchOwnerTriggerPreview{}, errors.New("known app Runtime-status launch owner trigger must keep KDE evidence-only")
	case preview.StateRootPathExposed || preview.EvidencePathExposed || preview.ManagedLauncherPathExposed || preview.RawLauncherOutputExposed || preview.BackendDetailsExposed:
		return KnownAppRuntimeStatusLaunchOwnerTriggerPreview{}, errors.New("known app Runtime-status launch owner trigger must not expose paths, launcher output, or backend details")
	case preview.HostRootModified || preview.DockerSocketMounted || preview.BroadHostMountRequired:
		return KnownAppRuntimeStatusLaunchOwnerTriggerPreview{}, errors.New("known app Runtime-status launch owner trigger must keep host and container boundaries closed")
	case preview.DesktopLaunchEnabled || preview.BackendLaunchEnabled || preview.ExecutionStarted || preview.BackendProcessStarted:
		return KnownAppRuntimeStatusLaunchOwnerTriggerPreview{}, errors.New("known app Runtime-status launch owner trigger must not start launch or backend processes")
	}
	for _, value := range []string{
		preview.SchemaVersion,
		preview.RequestType,
		preview.Source,
		preview.RuntimeMethod,
		preview.ReadMethod,
		preview.AppID,
		preview.DisplayName,
		preview.AppVersion,
		preview.EvidenceID,
		preview.EvidenceRelativePath,
		preview.EvidenceSHA256,
		preview.DesktopCallableRoute,
		preview.DesktopCallableRuntimeMethod,
		preview.DesktopCallableExecutionType,
		preview.DesktopDBusMethod,
		preview.OwnerServiceBoundary,
		preview.OwnerServiceMethod,
		preview.OwnerServiceCallType,
		preview.DesktopSafeSummary,
	} {
		if value != "" && !singleLine(value) {
			return KnownAppRuntimeStatusLaunchOwnerTriggerPreview{}, errors.New("known app Runtime-status launch owner trigger requires single-line fields")
		}
	}
	for _, value := range append(append([]string{}, preview.OwnerServiceCallArgs...), preview.OwnerServiceCLIArgs...) {
		if strings.TrimSpace(value) == "" || !singleLine(value) {
			return KnownAppRuntimeStatusLaunchOwnerTriggerPreview{}, errors.New("known app Runtime-status launch owner trigger requires single-line owner service arguments")
		}
	}
	if err := validateNoBackendTerms(preview, "known app Runtime-status launch owner trigger"); err != nil {
		return KnownAppRuntimeStatusLaunchOwnerTriggerPreview{}, err
	}
	return preview, nil
}
