package winapp

import "context"

const (
	LaunchProfileSchemaVersion = "xnix.runtime.windows_app_launch_profile.v1"
	LaunchProfileRequestType   = "windows-app-launch-profile"
)

type LaunchProfileRequest struct {
	ProfilePath string
}

type LaunchProfileResult struct {
	SchemaVersion               string                      `json:"schema_version"`
	RequestType                 string                      `json:"request_type"`
	Status                      string                      `json:"status"`
	ProfileSupplied             bool                        `json:"profile_supplied"`
	PreflightStatus             string                      `json:"preflight_status"`
	LaunchAttempted             bool                        `json:"launch_attempted"`
	ExecutableName              string                      `json:"executable_name"`
	ExecutableFormat            string                      `json:"executable_format"`
	WindowsExecutableSignature  bool                        `json:"windows_executable_signature_observed"`
	ExecutableArchitecture      string                      `json:"executable_architecture"`
	ExecutableArchitectureReady bool                        `json:"executable_architecture_supported"`
	WineArchitecture            string                      `json:"wine_architecture"`
	ApplicationWorkspaceMode    string                      `json:"application_workspace_mode"`
	RunnerAvailable             bool                        `json:"runner_available"`
	RawOutputRedacted           bool                        `json:"raw_output_redacted"`
	PreflightPayload            SmokeProfilePreflightResult `json:"preflight_payload"`
	RuntimePayload              *Result                     `json:"runtime_payload,omitempty"`
	NextAction                  string                      `json:"next_action"`
	SkipReason                  string                      `json:"skip_reason,omitempty"`
	FailureReason               string                      `json:"failure_reason,omitempty"`
	RawProfilePathExposed       bool                        `json:"raw_profile_path_exposed"`
	RawExecutablePathExposed    bool                        `json:"raw_executable_path_exposed"`
	RawRunnerPathExposed        bool                        `json:"raw_runner_path_exposed"`
	HostRootModified            bool                        `json:"host_root_modified"`
	PrivilegedContainerRequired bool                        `json:"privileged_container_required"`
	HostNetworkingRequired      bool                        `json:"host_networking_required"`
	DockerSocketMounted         bool                        `json:"docker_socket_mounted"`
	BroadHostMountRequired      bool                        `json:"broad_host_mount_required"`
	DockerExecuted              bool                        `json:"docker_executed"`
	QEMUExecuted                bool                        `json:"qemu_executed"`
	WineExecuted                bool                        `json:"wine_executed"`
	ColimaExecuted              bool                        `json:"colima_executed"`
	NetworkChecksRun            bool                        `json:"network_checks_run"`
	PackageManagerInvoked       bool                        `json:"package_manager_invoked"`
}

func LaunchProfile(ctx context.Context, request LaunchProfileRequest) (LaunchProfileResult, error) {
	result := baseLaunchProfileResult(request)
	preflight, err := PreflightSmokeProfile(request.ProfilePath)
	if err != nil {
		return result, err
	}
	result.PreflightPayload = preflight
	result.ProfileSupplied = preflight.ProfileSupplied
	result.PreflightStatus = preflight.Status
	result.ExecutableName = preflight.ExecutableName
	result.ExecutableFormat = preflight.ExecutableFormat
	result.WindowsExecutableSignature = preflight.WindowsExecutableSignature
	result.ExecutableArchitecture = preflight.ExecutableArchitecture
	result.ExecutableArchitectureReady = preflight.ExecutableArchitectureReady
	result.WineArchitecture = preflight.WineArchitecture
	result.ApplicationWorkspaceMode = preflight.ApplicationWorkspaceMode
	result.RunnerAvailable = preflight.RunnerAvailable
	result.NextAction = preflight.NextAction
	result.SkipReason = preflight.SkipReason
	result.FailureReason = preflight.FailureReason

	if preflight.Status != ProfileReadyStatus {
		result.Status = ProfileBlockedStatus
		return result, nil
	}

	smokeRequest, err := LoadSmokeProfile(request.ProfilePath)
	if err != nil {
		result.Status = FailedStatus
		result.FailureReason = "profile is unreadable or invalid"
		result.NextAction = "Fix the profile before launching the Windows application."
		return result, nil
	}
	smokeRequest.RedactOutput = true
	smokeResult, err := RunSmoke(ctx, smokeRequest)
	if err != nil {
		return result, err
	}
	result.LaunchAttempted = true
	result.RuntimePayload = &smokeResult
	result.Status = smokeResult.Status
	result.RawOutputRedacted = smokeResult.RawOutputRedacted
	result.NextAction = ""
	result.SkipReason = smokeResult.SkipReason
	result.FailureReason = smokeResult.FailureReason
	result.HostRootModified = smokeResult.HostRootModified
	result.PrivilegedContainerRequired = smokeResult.PrivilegedContainerRequired
	result.HostNetworkingRequired = smokeResult.HostNetworkingRequired
	result.DockerSocketMounted = smokeResult.DockerSocketMounted
	result.BroadHostMountRequired = smokeResult.BroadHostMountRequired
	result.WineExecuted = smokeResult.RunnerAvailable
	return result, nil
}

func baseLaunchProfileResult(request LaunchProfileRequest) LaunchProfileResult {
	return LaunchProfileResult{
		SchemaVersion:               LaunchProfileSchemaVersion,
		RequestType:                 LaunchProfileRequestType,
		Status:                      ProfileBlockedStatus,
		ProfileSupplied:             request.ProfilePath != "",
		PreflightStatus:             ProfileBlockedStatus,
		ExecutableFormat:            "unknown",
		ExecutableArchitecture:      "unknown",
		WineArchitecture:            "unknown",
		ApplicationWorkspaceMode:    ApplicationWorkspaceModeDirect,
		RawOutputRedacted:           true,
		RawProfilePathExposed:       false,
		RawExecutablePathExposed:    false,
		RawRunnerPathExposed:        false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		HostNetworkingRequired:      false,
		DockerSocketMounted:         false,
		BroadHostMountRequired:      false,
		DockerExecuted:              false,
		QEMUExecuted:                false,
		WineExecuted:                false,
		ColimaExecuted:              false,
		NetworkChecksRun:            false,
		PackageManagerInvoked:       false,
	}
}
