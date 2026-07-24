package winapp

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	SmokeProfilePreflightSchemaVersion = "xnix.runtime.windows_app_smoke_profile_preflight.v1"
	SmokeProfilePreflightRequestType   = "windows-app-smoke-profile-preflight"
	ProfileReadyStatus                 = "ready"
	ProfileBlockedStatus               = "blocked"
)

type SmokeProfilePreflightResult struct {
	SchemaVersion               string                  `json:"schema_version"`
	RequestType                 string                  `json:"request_type"`
	Status                      string                  `json:"status"`
	ProfileSupplied             bool                    `json:"profile_supplied"`
	ExecutableName              string                  `json:"executable_name"`
	ExecutableExists            bool                    `json:"executable_exists"`
	ExecutableFormat            string                  `json:"executable_format"`
	WindowsExecutableSignature  bool                    `json:"windows_executable_signature_observed"`
	WorkingDirectoryMode        string                  `json:"working_directory_mode"`
	WorkingDirectoryValid       bool                    `json:"working_directory_valid"`
	StateRootConfigured         bool                    `json:"state_root_configured"`
	RunnerAvailable             bool                    `json:"runner_available"`
	RunnerArgumentCount         int                     `json:"runner_argument_count"`
	AppArgumentCount            int                     `json:"app_argument_count"`
	SuccessMode                 string                  `json:"success_mode"`
	TimeoutConfigured           bool                    `json:"timeout_configured"`
	ExpectedMarkerConfigured    bool                    `json:"expected_marker_configured"`
	RunnerDiagnosticsPayload    RunnerDiagnosticsResult `json:"runner_diagnostics_payload"`
	NextAction                  string                  `json:"next_action"`
	FailureReason               string                  `json:"failure_reason,omitempty"`
	SkipReason                  string                  `json:"skip_reason,omitempty"`
	RawProfilePathExposed       bool                    `json:"raw_profile_path_exposed"`
	RawExecutablePathExposed    bool                    `json:"raw_executable_path_exposed"`
	RawRunnerPathExposed        bool                    `json:"raw_runner_path_exposed"`
	RawWorkingDirectoryExposed  bool                    `json:"raw_working_directory_exposed"`
	RawRunnerArgumentsExposed   bool                    `json:"raw_runner_arguments_exposed"`
	HostRootModified            bool                    `json:"host_root_modified"`
	PrivilegedContainerRequired bool                    `json:"privileged_container_required"`
	HostNetworkingRequired      bool                    `json:"host_networking_required"`
	DockerSocketMounted         bool                    `json:"docker_socket_mounted"`
	BroadHostMountRequired      bool                    `json:"broad_host_mount_required"`
	DockerExecuted              bool                    `json:"docker_executed"`
	QEMUExecuted                bool                    `json:"qemu_executed"`
	WineExecuted                bool                    `json:"wine_executed"`
	ColimaExecuted              bool                    `json:"colima_executed"`
	NetworkChecksRun            bool                    `json:"network_checks_run"`
	PackageManagerInvoked       bool                    `json:"package_manager_invoked"`
}

func PreflightSmokeProfile(profilePath string) (SmokeProfilePreflightResult, error) {
	result := baseSmokeProfilePreflightResult()
	if strings.TrimSpace(profilePath) == "" {
		result.FailureReason = "profile path is required"
		result.NextAction = "Provide a Windows app smoke profile JSON path."
		return result, nil
	}
	request, err := LoadSmokeProfile(profilePath)
	if err != nil {
		result.FailureReason = "profile is unreadable or invalid"
		result.NextAction = "Fix the profile path, JSON syntax, schema version, and timeout before preflight."
		return result, nil
	}
	result.ProfileSupplied = true

	successMode, err := normalizeSuccessMode(request.SuccessMode)
	if err != nil {
		result.FailureReason = "unsupported success mode"
		result.NextAction = "Use marker, exit-code, or startup-window in the smoke profile."
		return result, nil
	}
	result.SuccessMode = successMode
	result.TimeoutConfigured = request.Timeout > 0
	result.ExpectedMarkerConfigured = strings.TrimSpace(request.ExpectedMarker) != ""
	result.RunnerArgumentCount = len(runnerInvocationArguments(request))
	result.AppArgumentCount = len(request.Arguments)
	result.StateRootConfigured = strings.TrimSpace(request.StateRoot) != ""

	executablePath, executableErr := inspectProfileExecutable(request.ExecutablePath)
	if executableErr != nil {
		result.FailureReason = executableErr.Error()
		result.NextAction = "Fix executable_path in the smoke profile before running the app."
		return result, nil
	}
	result.ExecutableExists = true
	result.ExecutableName = filepath.Base(executablePath)
	executableFormat, signatureObserved, signatureErr := inspectWindowsExecutableSignature(executablePath)
	if signatureErr != nil {
		result.FailureReason = signatureErr.Error()
		result.NextAction = "Use a Windows PE executable with an MZ header in executable_path before running the app."
		return result, nil
	}
	result.ExecutableFormat = executableFormat
	result.WindowsExecutableSignature = signatureObserved

	workingDirectoryMode, workingDirectoryErr := inspectProfileWorkingDirectory(request.WorkingDirectory, executablePath)
	if workingDirectoryErr != nil {
		result.FailureReason = workingDirectoryErr.Error()
		result.NextAction = "Fix working_directory in the smoke profile or omit it to use the executable directory."
		return result, nil
	}
	result.WorkingDirectoryMode = workingDirectoryMode
	result.WorkingDirectoryValid = true

	if !result.StateRootConfigured {
		result.FailureReason = "state root is required"
		result.NextAction = "Set state_root in the smoke profile before running the app."
		return result, nil
	}

	diagnostics := RunnerDiagnostics(request.RunnerPath)
	result.RunnerDiagnosticsPayload = diagnostics
	result.RunnerAvailable = diagnostics.RunnerAvailable
	if !diagnostics.RunnerAvailable {
		result.SkipReason = "windows compatibility runner unavailable"
		result.NextAction = diagnostics.NextAction
		return result, nil
	}

	result.Status = ProfileReadyStatus
	result.NextAction = "Run `ruby scripts/winapp_smoke.rb --format json --profile <profile>` or `go run ./cmd/xnix-runtime-go windows-app-run-smoke --profile <profile>`."
	return result, nil
}

func baseSmokeProfilePreflightResult() SmokeProfilePreflightResult {
	return SmokeProfilePreflightResult{
		SchemaVersion:               SmokeProfilePreflightSchemaVersion,
		RequestType:                 SmokeProfilePreflightRequestType,
		Status:                      ProfileBlockedStatus,
		ExecutableFormat:            "unknown",
		WorkingDirectoryMode:        WorkingDirectoryModeExecutable,
		SuccessMode:                 SuccessModeMarker,
		RawProfilePathExposed:       false,
		RawExecutablePathExposed:    false,
		RawRunnerPathExposed:        false,
		RawWorkingDirectoryExposed:  false,
		RawRunnerArgumentsExposed:   false,
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

func inspectProfileExecutable(path string) (string, error) {
	path = strings.TrimSpace(path)
	if path == "" {
		return "", fmt.Errorf("executable path is required")
	}
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("resolve executable path")
	}
	info, err := os.Stat(absolutePath)
	if err != nil {
		return "", fmt.Errorf("executable path is not readable")
	}
	if info.IsDir() {
		return "", fmt.Errorf("executable path must be a file")
	}
	return absolutePath, nil
}

func inspectProfileWorkingDirectory(path string, executablePath string) (string, error) {
	path = strings.TrimSpace(path)
	mode := WorkingDirectoryModeOperator
	if path == "" {
		path = filepath.Dir(executablePath)
		mode = WorkingDirectoryModeExecutable
	}
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("resolve working directory")
	}
	info, err := os.Stat(absolutePath)
	if err != nil {
		return "", fmt.Errorf("working directory is not readable")
	}
	if !info.IsDir() {
		return "", fmt.Errorf("working directory must be a directory")
	}
	return mode, nil
}

func inspectWindowsExecutableSignature(path string) (string, bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return "unknown", false, fmt.Errorf("executable signature is not readable")
	}
	defer func() {
		_ = file.Close()
	}()
	header := make([]byte, 2)
	read, err := file.Read(header)
	if err != nil || read < len(header) {
		return "unknown", false, fmt.Errorf("executable signature is not readable")
	}
	if string(header) != "MZ" {
		return "unknown", false, fmt.Errorf("executable is not a Windows PE file")
	}
	return "pe-mz", true, nil
}
