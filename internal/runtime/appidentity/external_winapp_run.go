package appidentity

import (
	"context"
	"fmt"
	"strings"
	"time"

	"xnix.local/xnix/internal/runtime/winapp"
)

const (
	ExternalWinAppRunSchemaVersion = "xnix.runtime.external_winapp_run.v1"
	ExternalWinAppRunRequestType   = "windows-external-app-run"
)

type ExternalWinAppRunRequest struct {
	ImportRecordPath         string
	StateRoot                string
	ExternalAppHandle        string
	ExternalDesktopArguments []string
	WindowMatch              string
	Image                    string
	Platform                 string
	DockerPath               string
	Timeout                  time.Duration
}

type ExternalWinAppRunResult struct {
	Version                          string                     `json:"version"`
	SchemaVersion                    string                     `json:"schema_version"`
	RequestType                      string                     `json:"request_type"`
	RunType                          string                     `json:"run_type"`
	Source                           string                     `json:"source"`
	RuntimeMethod                    string                     `json:"runtime_method"`
	ReadMethod                       string                     `json:"read_method"`
	Status                           string                     `json:"status"`
	ApplicationID                    string                     `json:"application_id"`
	DisplayName                      string                     `json:"display_name"`
	AppVersion                       string                     `json:"app_version"`
	ExecutableName                   string                     `json:"executable_name"`
	ExternalAppImportRecordConsumed  bool                       `json:"external_app_import_record_consumed"`
	ExternalAppHandleConsumed        bool                       `json:"external_app_handle_consumed"`
	ExternalAppHandle                string                     `json:"external_app_handle,omitempty"`
	ExternalDesktopArgumentCount     int                        `json:"external_desktop_argument_count"`
	ExternalFileURIArgumentsAccepted bool                       `json:"external_file_uri_arguments_accepted"`
	ExternalFileOpenRequested        bool                       `json:"external_file_open_requested"`
	ExternalFileBridgeMountEnabled   bool                       `json:"external_file_bridge_mount_enabled"`
	RawFileURIArgumentsExposed       bool                       `json:"raw_file_uri_arguments_exposed"`
	ImportedArtifactDigestVerified   bool                       `json:"imported_artifact_digest_verified"`
	ImportedArtifactSHA256           string                     `json:"imported_artifact_sha256"`
	RuntimeRunRequested              bool                       `json:"runtime_run_requested"`
	RuntimeRunExecuted               bool                       `json:"runtime_run_executed"`
	ExecutionStarted                 bool                       `json:"execution_started"`
	BackendProcessStarted            bool                       `json:"backend_process_started"`
	ContainerRuntimeUsed             bool                       `json:"container_runtime_used"`
	ContainerNetworkMode             string                     `json:"container_network_mode"`
	ContainerHostMountCount          int                        `json:"container_host_mount_count"`
	XServerStarted                   bool                       `json:"x_server_started"`
	WineBootstrapAttempted           bool                       `json:"wine_bootstrap_attempted"`
	XWindowObserved                  bool                       `json:"x_window_observed"`
	WindowObserved                   bool                       `json:"window_observed"`
	WindowEvidenceSummary            string                     `json:"window_evidence_summary,omitempty"`
	RuntimePayload                   winapp.ContainerXGUIResult `json:"runtime_payload"`
	RuntimeOwned                     bool                       `json:"runtime_owned"`
	GoRuntimeBacked                  bool                       `json:"go_runtime_backed"`
	KDEPolicyOwner                   bool                       `json:"kde_policy_owner"`
	DesktopLaunchEnabled             bool                       `json:"desktop_launch_enabled"`
	ActionExecutionEnabled           bool                       `json:"action_execution_enabled"`
	BackendDetailsExposed            bool                       `json:"backend_details_exposed"`
	RawImportRecordPathExposed       bool                       `json:"raw_import_record_path_exposed"`
	RawExternalAppHandlePathExposed  bool                       `json:"raw_external_app_handle_path_exposed"`
	RawStateRootPathExposed          bool                       `json:"raw_state_root_path_exposed"`
	RawExecutablePathExposed         bool                       `json:"raw_executable_path_exposed"`
	HostRootModified                 bool                       `json:"host_root_modified"`
	PrivilegedContainerRequired      bool                       `json:"privileged_container_required"`
	HostNetworkingRequired           bool                       `json:"host_networking_required"`
	DockerSocketMounted              bool                       `json:"docker_socket_mounted"`
	BroadHostMountRequired           bool                       `json:"broad_host_mount_required"`
	DesktopSafeSummary               string                     `json:"desktop_safe_summary"`
	SkipReason                       string                     `json:"skip_reason,omitempty"`
	FailureReason                    string                     `json:"failure_reason,omitempty"`
}

func RunExternalWinApp(ctx context.Context, request ExternalWinAppRunRequest) (ExternalWinAppRunResult, error) {
	importRecordPath := strings.TrimSpace(request.ImportRecordPath)
	stateRoot := strings.TrimSpace(request.StateRoot)
	externalAppHandle := strings.TrimSpace(request.ExternalAppHandle)
	if importRecordPath != "" && (stateRoot != "" || externalAppHandle != "") {
		return ExternalWinAppRunResult{}, fmt.Errorf("external Windows app run accepts either --external-app-import-record or --state-root with --external-app-handle")
	}
	if importRecordPath == "" && externalAppHandle == "" {
		return ExternalWinAppRunResult{}, fmt.Errorf("external Windows app run requires --external-app-import-record or --external-app-handle")
	}
	if importRecordPath == "" && stateRoot == "" {
		stateRoot = DefaultExternalWinAppRuntimeStateRoot()
	}
	argumentCount, fileOpenRequested, err := validateExternalDesktopArguments(request.ExternalDesktopArguments)
	if err != nil {
		return ExternalWinAppRunResult{}, err
	}
	var record ExternalWinAppImportRecord
	var importedExecutablePath string
	if importRecordPath != "" {
		record, importedExecutablePath, err = ResolveExternalWinAppImportedArtifact(importRecordPath)
	} else {
		record, importedExecutablePath, err = ResolveExternalWinAppImportedArtifactByHandle(stateRoot, externalAppHandle)
	}
	if err != nil {
		return ExternalWinAppRunResult{}, err
	}
	windowMatch := strings.TrimSpace(request.WindowMatch)
	if windowMatch == "" {
		windowMatch = record.ExecutableName
	}
	runtimePayload, err := winapp.RunContainerXGUISmoke(ctx, winapp.ContainerXGUIRequest{
		ExecutablePath:                  importedExecutablePath,
		ApplicationName:                 "/" + record.ExecutableName,
		WindowMatch:                     windowMatch,
		ApplicationID:                   record.ApplicationID,
		DisplayName:                     record.DisplayName,
		AppVersion:                      record.AppVersion,
		ExternalAppImportRecordConsumed: true,
		ImportedArtifactDigestVerified:  true,
		ImportedArtifactSHA256:          record.ArtifactSHA256,
		Image:                           request.Image,
		Platform:                        request.Platform,
		DockerPath:                      request.DockerPath,
		Timeout:                         request.Timeout,
	})
	if err != nil {
		return ExternalWinAppRunResult{}, err
	}
	result := ExternalWinAppRunResult{
		Version:                          record.Version,
		SchemaVersion:                    ExternalWinAppRunSchemaVersion,
		RequestType:                      ExternalWinAppRunRequestType,
		RunType:                          "external-windows-app-container-gui-run",
		Source:                           "go-runtime-external-winapp-import+container-x-gui-run",
		RuntimeMethod:                    "RunExternalWinApp",
		ReadMethod:                       "GetExternalWinAppRunResult",
		Status:                           runtimePayload.Status,
		ApplicationID:                    record.ApplicationID,
		DisplayName:                      record.DisplayName,
		AppVersion:                       record.AppVersion,
		ExecutableName:                   record.ExecutableName,
		ExternalAppImportRecordConsumed:  true,
		ExternalAppHandleConsumed:        externalAppHandle != "",
		ExternalAppHandle:                externalAppHandle,
		ExternalDesktopArgumentCount:     argumentCount,
		ExternalFileURIArgumentsAccepted: argumentCount > 0,
		ExternalFileOpenRequested:        fileOpenRequested,
		ExternalFileBridgeMountEnabled:   false,
		RawFileURIArgumentsExposed:       false,
		ImportedArtifactDigestVerified:   runtimePayload.ImportedArtifactDigestVerified,
		ImportedArtifactSHA256:           record.ArtifactSHA256,
		RuntimeRunRequested:              true,
		RuntimeRunExecuted:               runtimePayload.Status == winapp.PassedStatus,
		ExecutionStarted:                 runtimePayload.XServerStarted,
		BackendProcessStarted:            runtimePayload.XServerStarted,
		ContainerRuntimeUsed:             true,
		ContainerNetworkMode:             runtimePayload.NetworkMode,
		ContainerHostMountCount:          runtimePayload.HostMountCount,
		XServerStarted:                   runtimePayload.XServerStarted,
		WineBootstrapAttempted:           runtimePayload.WineBootstrapAttempted,
		XWindowObserved:                  runtimePayload.XWindowObserved,
		WindowObserved:                   runtimePayload.XWindowObserved,
		WindowEvidenceSummary:            runtimePayload.WindowEvidenceSummary,
		RuntimePayload:                   runtimePayload,
		RuntimeOwned:                     true,
		GoRuntimeBacked:                  true,
		KDEPolicyOwner:                   false,
		DesktopLaunchEnabled:             false,
		ActionExecutionEnabled:           false,
		BackendDetailsExposed:            false,
		RawImportRecordPathExposed:       false,
		RawExternalAppHandlePathExposed:  false,
		RawStateRootPathExposed:          false,
		RawExecutablePathExposed:         false,
		HostRootModified:                 runtimePayload.HostRootModified,
		PrivilegedContainerRequired:      runtimePayload.PrivilegedContainerRequired,
		HostNetworkingRequired:           runtimePayload.HostNetworkingRequired,
		DockerSocketMounted:              runtimePayload.DockerSocketMounted,
		BroadHostMountRequired:           runtimePayload.BroadHostMountRequired,
		DesktopSafeSummary:               fmt.Sprintf("%s was run by the Runtime from a digest-verified imported Windows app artifact in an isolated container GUI session.", record.DisplayName),
		SkipReason:                       runtimePayload.SkipReason,
		FailureReason:                    runtimePayload.FailureReason,
	}
	if result.Status != winapp.PassedStatus {
		result.RuntimeRunExecuted = false
		result.DesktopSafeSummary = fmt.Sprintf("%s Runtime imported Windows app run did not complete successfully; unsafe desktop and backend details remain hidden.", record.DisplayName)
	}
	return result, nil
}

func validateExternalDesktopArguments(values []string) (int, bool, error) {
	count := 0
	fileOpenRequested := false
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed == "" || trimmed == "%U" || trimmed == "%u" {
			continue
		}
		if !singleLine(trimmed) {
			return 0, false, fmt.Errorf("external Windows app desktop argument must be single-line")
		}
		if !strings.HasPrefix(strings.ToLower(trimmed), "file://") {
			return 0, false, fmt.Errorf("external Windows app desktop argument must be a file URI")
		}
		count++
		fileOpenRequested = true
	}
	return count, fileOpenRequested, nil
}
