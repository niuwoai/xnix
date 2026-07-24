package winapp

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	SchemaVersion                  = "xnix.runtime.windows_app_smoke.v1"
	RunnerDiagnosticsSchemaVersion = "xnix.runtime.windows_app_runner_diagnostics.v1"
	RequestType                    = "windows-app-run-smoke"
	RunnerDiagnosticsRequestType   = "windows-app-runner-diagnostics"
	RunnerEnvVar                   = "XNIX_WINDOWS_RUNNER"
	PassedStatus                   = "passed"
	FailedStatus                   = "failed"
	SkippedStatus                  = "skipped"
	SuccessModeMarker              = "marker"
	SuccessModeExitCode            = "exit-code"
	SuccessModeStartupWindow       = "startup-window"
	WorkingDirectoryModeExecutable = "executable-directory"
	WorkingDirectoryModeOperator   = "operator-supplied"
	WorkingDirectoryModeStaged     = "staged-application-workspace"
	ApplicationWorkspaceModeDirect = "direct-executable"
	ApplicationWorkspaceModeStaged = "staged-application-directory"
	DefaultMarker                  = "XNIX_WINAPP_SMOKE_OK"
)

type Request struct {
	ExecutablePath   string
	Arguments        []string
	RunnerArguments  []string
	RunnerBottle     string
	StateRoot        string
	WorkingDirectory string
	RunnerPath       string
	Timeout          time.Duration
	ExpectedMarker   string
	SuccessMode      string
	RedactOutput     bool
	SkipBootstrap    bool
	StageAppDir      bool
}

type Result struct {
	SchemaVersion               string `json:"schema_version"`
	RequestType                 string `json:"request_type"`
	Status                      string `json:"status"`
	ExecutableName              string `json:"executable_name"`
	ExecutableFormat            string `json:"executable_format"`
	WindowsExecutableSignature  bool   `json:"windows_executable_signature_observed"`
	ExecutableArchitecture      string `json:"executable_architecture"`
	ExecutableArchitectureReady bool   `json:"executable_architecture_supported"`
	WineArchitecture            string `json:"wine_architecture"`
	WinePrefixMode              string `json:"wine_prefix_mode"`
	WinePrefixPrepared          bool   `json:"wine_prefix_prepared"`
	RunnerAvailable             bool   `json:"runner_available"`
	RunnerArgumentCount         int    `json:"runner_argument_count"`
	CompatibilityLayer          string `json:"compatibility_layer"`
	WineBootstrapAttempted      bool   `json:"wine_bootstrap_attempted"`
	WineBootstrapSucceeded      bool   `json:"wine_bootstrap_succeeded"`
	WineBootstrapSkipped        bool   `json:"wine_bootstrap_skipped"`
	WineBootstrapExitCode       int    `json:"wine_bootstrap_exit_code"`
	ExpectedMarker              string `json:"expected_marker"`
	SuccessMode                 string `json:"success_mode"`
	MarkerObserved              bool   `json:"marker_observed"`
	StartupWindowObserved       bool   `json:"startup_window_observed"`
	WorkingDirectoryMode        string `json:"working_directory_mode"`
	ApplicationWorkspaceMode    string `json:"application_workspace_mode"`
	ApplicationStaged           bool   `json:"application_staged"`
	ApplicationStagedFileCount  int    `json:"application_staged_file_count"`
	ApplicationStagedBytes      int64  `json:"application_staged_bytes"`
	ExitCode                    int    `json:"exit_code"`
	DurationMillis              int64  `json:"duration_millis"`
	Stdout                      string `json:"stdout"`
	Stderr                      string `json:"stderr"`
	StdoutBytes                 int    `json:"stdout_bytes"`
	StderrBytes                 int    `json:"stderr_bytes"`
	StdoutLineCount             int    `json:"stdout_line_count"`
	StderrLineCount             int    `json:"stderr_line_count"`
	RawOutputIncluded           bool   `json:"raw_output_included"`
	RawOutputRedacted           bool   `json:"raw_output_redacted"`
	KDESafeOutputSummary        string `json:"kde_safe_output_summary"`
	SkipReason                  string `json:"skip_reason,omitempty"`
	FailureReason               string `json:"failure_reason,omitempty"`
	IsolatedStateRoot           bool   `json:"isolated_state_root"`
	HostRootModified            bool   `json:"host_root_modified"`
	PrivilegedContainerRequired bool   `json:"privileged_container_required"`
	HostNetworkingRequired      bool   `json:"host_networking_required"`
	DockerSocketMounted         bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired      bool   `json:"broad_host_mount_required"`
}

type RunnerDiagnosticsResult struct {
	SchemaVersion               string                    `json:"schema_version"`
	RequestType                 string                    `json:"request_type"`
	Status                      string                    `json:"status"`
	RunnerAvailable             bool                      `json:"runner_available"`
	ExplicitRunnerSupplied      bool                      `json:"explicit_runner_supplied"`
	EnvRunnerConfigured         bool                      `json:"env_runner_configured"`
	CandidateCount              int                       `json:"candidate_count"`
	Candidates                  []RunnerCandidateEvidence `json:"candidates"`
	SelectedRunnerName          string                    `json:"selected_runner_name"`
	NextAction                  string                    `json:"next_action"`
	RunnerCommandHints          []string                  `json:"runner_command_hints"`
	RawPathExposed              bool                      `json:"raw_path_exposed"`
	HostRootModified            bool                      `json:"host_root_modified"`
	PrivilegedContainerRequired bool                      `json:"privileged_container_required"`
	HostNetworkingRequired      bool                      `json:"host_networking_required"`
	DockerSocketMounted         bool                      `json:"docker_socket_mounted"`
	BroadHostMountRequired      bool                      `json:"broad_host_mount_required"`
	DockerExecuted              bool                      `json:"docker_executed"`
	QEMUExecuted                bool                      `json:"qemu_executed"`
	ColimaExecuted              bool                      `json:"colima_executed"`
	NetworkChecksRun            bool                      `json:"network_checks_run"`
	PackageManagerInvoked       bool                      `json:"package_manager_invoked"`
}

type RunnerCandidateEvidence struct {
	ID        string `json:"id"`
	Source    string `json:"source"`
	Available bool   `json:"available"`
	Selected  bool   `json:"selected"`
	Reason    string `json:"reason"`
}

type runnerCandidate struct {
	id     string
	source string
	path   string
}

func RunSmoke(ctx context.Context, request Request) (Result, error) {
	result := baseResult(request)
	successMode, err := normalizeSuccessMode(request.SuccessMode)
	if err != nil {
		return result, err
	}
	result.SuccessMode = successMode
	result.WineBootstrapSkipped = request.SkipBootstrap

	executablePath, err := validateExecutable(request.ExecutablePath)
	if err != nil {
		return result, err
	}
	result.ExecutableName = filepath.Base(executablePath)
	executableFormat, signatureObserved, executableArchitecture, architectureReady, err := inspectWindowsExecutableSignature(executablePath)
	if err != nil {
		return result, err
	}
	result.ExecutableFormat = executableFormat
	result.WindowsExecutableSignature = signatureObserved
	result.ExecutableArchitecture = executableArchitecture
	result.ExecutableArchitectureReady = architectureReady
	if !architectureReady {
		result.Status = FailedStatus
		result.FailureReason = "Windows executable architecture is not supported"
		return result, nil
	}
	wineArchitecture := wineArchitectureForExecutable(executableArchitecture)
	result.WineArchitecture = wineArchitecture

	if strings.TrimSpace(request.StateRoot) == "" {
		return result, errors.New("state root is required")
	}
	stateRoot, err := filepath.Abs(request.StateRoot)
	if err != nil {
		return result, fmt.Errorf("resolve state root: %w", err)
	}
	if err := os.MkdirAll(stateRoot, 0o700); err != nil {
		return result, fmt.Errorf("create isolated state root: %w", err)
	}
	result.IsolatedStateRoot = true
	workingDirectory := ""
	workingDirectoryMode := ""
	if request.StageAppDir {
		stagedExecutablePath, stagedWorkingDirectory, fileCount, byteCount, stageErr := stageApplicationDirectory(stateRoot, executablePath, request.WorkingDirectory)
		if stageErr != nil {
			return result, stageErr
		}
		executablePath = stagedExecutablePath
		workingDirectory = stagedWorkingDirectory
		workingDirectoryMode = WorkingDirectoryModeStaged
		result.ApplicationWorkspaceMode = ApplicationWorkspaceModeStaged
		result.ApplicationStaged = true
		result.ApplicationStagedFileCount = fileCount
		result.ApplicationStagedBytes = byteCount
	} else {
		workingDirectory, workingDirectoryMode, err = resolveWorkingDirectory(request.WorkingDirectory, executablePath)
		if err != nil {
			return result, err
		}
	}
	result.WorkingDirectoryMode = workingDirectoryMode

	runnerPath, err := resolveRunner(request.RunnerPath)
	if err != nil {
		result.Status = SkippedStatus
		result.SkipReason = "windows compatibility runner unavailable"
		return result, nil
	}
	result.RunnerAvailable = true
	winePrefix, winePrefixMode, err := prepareWinePrefix(stateRoot, wineArchitecture)
	if err != nil {
		return result, err
	}
	result.WinePrefixMode = winePrefixMode
	result.WinePrefixPrepared = true
	runnerArguments := runnerInvocationArguments(request)
	result.RunnerArgumentCount = len(runnerArguments)

	timeout := request.Timeout
	if timeout <= 0 {
		timeout = 20 * time.Second
	}
	runCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	bootstrapPath, bootstrapAvailable := resolveWineboot(runnerPath)
	if bootstrapAvailable && !request.SkipBootstrap {
		result.WineBootstrapAttempted = true
		bootstrapArgs := append([]string{}, runnerArguments...)
		bootstrapArgs = append(bootstrapArgs, "--init")
		bootstrapCommand := exec.CommandContext(runCtx, bootstrapPath, bootstrapArgs...)
		bootstrapCommand.Env = runnerEnvironment(winePrefix, wineArchitecture)
		bootstrapCommand.Dir = workingDirectory
		var bootstrapStderr bytes.Buffer
		bootstrapCommand.Stderr = &bootstrapStderr
		bootstrapErr := bootstrapCommand.Run()
		result.WineBootstrapExitCode = exitCode(bootstrapErr)
		if runCtx.Err() == context.DeadlineExceeded {
			result.Status = FailedStatus
			result.FailureReason = "Wine prefix bootstrap timed out"
			return result, nil
		}
		if bootstrapErr != nil {
			result.Status = FailedStatus
			result.FailureReason = "Wine prefix bootstrap failed"
			result.StderrBytes = len(bootstrapStderr.String())
			result.StderrLineCount = lineCount(bootstrapStderr.String())
			result.KDESafeOutputSummary = kdeSafeOutputSummary(result)
			return result, nil
		}
		result.WineBootstrapSucceeded = true
	}

	args := append([]string{}, runnerArguments...)
	args = append(args, executablePath)
	args = append(args, request.Arguments...)
	command := exec.CommandContext(runCtx, runnerPath, args...)
	command.Env = runnerEnvironment(winePrefix, wineArchitecture)
	command.Dir = workingDirectory

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	command.Stdout = &stdout
	command.Stderr = &stderr

	startedAt := time.Now()
	err = command.Run()
	result.DurationMillis = time.Since(startedAt).Milliseconds()
	stdoutText := stdout.String()
	stderrText := stderr.String()
	result.StdoutBytes = len(stdoutText)
	result.StderrBytes = len(stderrText)
	result.StdoutLineCount = lineCount(stdoutText)
	result.StderrLineCount = lineCount(stderrText)
	result.RawOutputIncluded = !request.RedactOutput
	result.RawOutputRedacted = request.RedactOutput
	result.MarkerObserved = strings.Contains(stdoutText, result.ExpectedMarker)
	result.KDESafeOutputSummary = kdeSafeOutputSummary(result)
	if request.RedactOutput {
		result.Stdout = ""
		result.Stderr = ""
	} else {
		result.Stdout = stdoutText
		result.Stderr = stderrText
	}
	result.ExitCode = exitCode(err)

	if runCtx.Err() == context.DeadlineExceeded {
		if successMode == SuccessModeStartupWindow {
			result.Status = PassedStatus
			result.StartupWindowObserved = true
			return result, nil
		}
		result.Status = FailedStatus
		result.FailureReason = "execution timed out"
		return result, nil
	}
	if err != nil {
		result.Status = FailedStatus
		result.FailureReason = "runner returned a non-zero exit status"
		return result, nil
	}
	if successMode == SuccessModeExitCode {
		result.Status = PassedStatus
		return result, nil
	}
	if successMode == SuccessModeStartupWindow {
		result.Status = FailedStatus
		result.FailureReason = "process exited before startup window elapsed"
		return result, nil
	}
	if !result.MarkerObserved {
		result.Status = FailedStatus
		result.FailureReason = "expected smoke marker was not observed"
		return result, nil
	}

	result.Status = PassedStatus
	return result, nil
}

func resolveWorkingDirectory(path string, executablePath string) (string, string, error) {
	path = strings.TrimSpace(path)
	mode := WorkingDirectoryModeOperator
	if path == "" {
		path = filepath.Dir(executablePath)
		mode = WorkingDirectoryModeExecutable
	}
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return "", "", fmt.Errorf("resolve working directory: %w", err)
	}
	info, err := os.Stat(absolutePath)
	if err != nil {
		return "", "", fmt.Errorf("inspect working directory: %w", err)
	}
	if !info.IsDir() {
		return "", "", errors.New("working directory must be a directory")
	}
	return absolutePath, mode, nil
}

func normalizeSuccessMode(mode string) (string, error) {
	mode = strings.TrimSpace(mode)
	if mode == "" {
		return SuccessModeMarker, nil
	}
	switch mode {
	case SuccessModeMarker, SuccessModeExitCode, SuccessModeStartupWindow:
		return mode, nil
	default:
		return "", fmt.Errorf("unsupported success mode %q", mode)
	}
}

func RunnerDiagnostics(explicitRunner string) RunnerDiagnosticsResult {
	explicitRunner = strings.TrimSpace(explicitRunner)
	envRunner := strings.TrimSpace(os.Getenv(RunnerEnvVar))
	candidates := configuredRunnerCandidates(explicitRunner, envRunner)
	if explicitRunner != "" {
		envRunner = ""
	}

	result := RunnerDiagnosticsResult{
		SchemaVersion:               RunnerDiagnosticsSchemaVersion,
		RequestType:                 RunnerDiagnosticsRequestType,
		Status:                      SkippedStatus,
		ExplicitRunnerSupplied:      explicitRunner != "",
		EnvRunnerConfigured:         envRunner != "",
		CandidateCount:              len(candidates),
		RawPathExposed:              false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		HostNetworkingRequired:      false,
		DockerSocketMounted:         false,
		BroadHostMountRequired:      false,
		DockerExecuted:              false,
		QEMUExecuted:                false,
		ColimaExecuted:              false,
		NetworkChecksRun:            false,
		PackageManagerInvoked:       false,
	}

	for _, candidate := range candidates {
		evidence, selectedName := probeRunnerCandidate(candidate)
		if selectedName != "" && !result.RunnerAvailable {
			evidence.Selected = true
			result.Status = PassedStatus
			result.RunnerAvailable = true
			result.SelectedRunnerName = selectedName
		}
		result.Candidates = append(result.Candidates, evidence)
	}

	if result.RunnerAvailable {
		result.NextAction = "Run `windows-app-run-smoke` with the selected runner or omit `--runner` when it is discoverable from the managed environment."
		result.RunnerCommandHints = runnerCommandHints(true)
		return result
	}
	if explicitRunner != "" {
		result.NextAction = "Provide an existing Wine-compatible runner file through `--runner PATH` or install Wine so `wine` or `wine64` is discoverable."
	} else if envRunner != "" {
		result.NextAction = "Fix the XNIX_WINDOWS_RUNNER value so it points to an existing Wine-compatible runner file, or unset it and install Wine so `wine` or `wine64` is discoverable."
	} else {
		result.NextAction = "Install or provide a Wine-compatible runner, then rerun `windows-app-run-smoke`; no Docker, QEMU, Colima, network, or package-manager action was attempted by this diagnostic."
	}
	result.RunnerCommandHints = runnerCommandHints(false)
	return result
}

func runnerCommandHints(runnerAvailable bool) []string {
	if runnerAvailable {
		return []string{
			"ruby scripts/winapp_smoke.rb --exe path/to/app.exe --format json",
			"ruby scripts/winapp_smoke.rb --exe path/to/app.exe --runner path/to/wine --format json",
			"ruby scripts/winapp_smoke.rb --exe path/to/app.exe --runner path/to/wine --runner-bottle bottle-name --format json",
		}
	}
	return []string{
		"XNIX_WINDOWS_RUNNER=path/to/wine ruby scripts/winapp_smoke.rb --exe path/to/app.exe --format json",
		"ruby scripts/winapp_smoke.rb --exe path/to/app.exe --runner path/to/wine --format json",
		"ruby scripts/winapp_smoke.rb --exe path/to/app.exe --runner path/to/wine --runner-bottle bottle-name --format json",
		"ruby scripts/winapp_smoke.rb --format json",
	}
}

func runnerInvocationArguments(request Request) []string {
	args := []string{}
	if bottle := strings.TrimSpace(request.RunnerBottle); bottle != "" {
		args = append(args, "--bottle", bottle)
	}
	args = append(args, request.RunnerArguments...)
	return args
}

func baseResult(request Request) Result {
	marker := request.ExpectedMarker
	if strings.TrimSpace(marker) == "" {
		marker = DefaultMarker
	}
	return Result{
		SchemaVersion:               SchemaVersion,
		RequestType:                 RequestType,
		Status:                      FailedStatus,
		ExecutableFormat:            "unknown",
		ExecutableArchitecture:      "unknown",
		WineArchitecture:            "unknown",
		WinePrefixMode:              "unknown",
		CompatibilityLayer:          "windows-compatibility-layer",
		WineBootstrapExitCode:       -1,
		ExpectedMarker:              marker,
		SuccessMode:                 SuccessModeMarker,
		WorkingDirectoryMode:        WorkingDirectoryModeExecutable,
		ApplicationWorkspaceMode:    ApplicationWorkspaceModeDirect,
		ExitCode:                    -1,
		RawOutputIncluded:           !request.RedactOutput,
		RawOutputRedacted:           request.RedactOutput,
		KDESafeOutputSummary:        "execution has not started",
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		HostNetworkingRequired:      false,
		DockerSocketMounted:         false,
		BroadHostMountRequired:      false,
	}
}

func stageApplicationDirectory(stateRoot string, executablePath string, workingDirectory string) (string, string, int, int64, error) {
	sourceDir := strings.TrimSpace(workingDirectory)
	if sourceDir == "" {
		sourceDir = filepath.Dir(executablePath)
	}
	sourceDir, err := filepath.Abs(sourceDir)
	if err != nil {
		return "", "", 0, 0, fmt.Errorf("resolve application source directory: %w", err)
	}
	info, err := os.Stat(sourceDir)
	if err != nil {
		return "", "", 0, 0, fmt.Errorf("inspect application source directory: %w", err)
	}
	if !info.IsDir() {
		return "", "", 0, 0, errors.New("application source directory must be a directory")
	}
	relativeExecutablePath, err := filepath.Rel(sourceDir, executablePath)
	if err != nil {
		return "", "", 0, 0, fmt.Errorf("resolve executable path relative to application source directory: %w", err)
	}
	if relativeExecutablePath == "." || strings.HasPrefix(relativeExecutablePath, ".."+string(os.PathSeparator)) || relativeExecutablePath == ".." || filepath.IsAbs(relativeExecutablePath) {
		return "", "", 0, 0, errors.New("executable must be inside the staged application source directory")
	}

	workspaceRoot := filepath.Join(stateRoot, "app-workspace")
	nextWorkspaceRoot := filepath.Join(stateRoot, "app-workspace.next")
	if err := os.RemoveAll(nextWorkspaceRoot); err != nil {
		return "", "", 0, 0, fmt.Errorf("clear pending application workspace: %w", err)
	}
	if err := os.MkdirAll(nextWorkspaceRoot, 0o700); err != nil {
		return "", "", 0, 0, fmt.Errorf("create pending application workspace: %w", err)
	}

	stateRootClean := filepath.Clean(stateRoot)
	fileCount := 0
	var byteCount int64
	walkErr := filepath.WalkDir(sourceDir, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		cleanPath := filepath.Clean(path)
		if cleanPath == stateRootClean {
			if entry.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		relativePath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}
		if relativePath == "." {
			return nil
		}
		destinationPath := filepath.Join(nextWorkspaceRoot, relativePath)
		info, err := entry.Info()
		if err != nil {
			return err
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return nil
		}
		if entry.IsDir() {
			return os.MkdirAll(destinationPath, 0o700)
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		if err := copyRegularFile(path, destinationPath, info.Mode().Perm()); err != nil {
			return err
		}
		fileCount++
		byteCount += info.Size()
		return nil
	})
	if walkErr != nil {
		return "", "", 0, 0, fmt.Errorf("stage application directory: %w", walkErr)
	}
	if err := os.RemoveAll(workspaceRoot); err != nil {
		return "", "", 0, 0, fmt.Errorf("clear previous application workspace: %w", err)
	}
	if err := os.Rename(nextWorkspaceRoot, workspaceRoot); err != nil {
		return "", "", 0, 0, fmt.Errorf("publish application workspace: %w", err)
	}
	stagedExecutablePath := filepath.Join(workspaceRoot, relativeExecutablePath)
	return stagedExecutablePath, workspaceRoot, fileCount, byteCount, nil
}

func copyRegularFile(sourcePath string, destinationPath string, mode os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(destinationPath), 0o700); err != nil {
		return err
	}
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer func() {
		_ = sourceFile.Close()
	}()
	destinationFile, err := os.OpenFile(destinationPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	if _, err := io.Copy(destinationFile, sourceFile); err != nil {
		_ = destinationFile.Close()
		return err
	}
	if err := destinationFile.Close(); err != nil {
		return err
	}
	return nil
}

func lineCount(text string) int {
	if text == "" {
		return 0
	}
	count := strings.Count(text, "\n")
	if !strings.HasSuffix(text, "\n") {
		count++
	}
	return count
}

func kdeSafeOutputSummary(result Result) string {
	streamSummary := fmt.Sprintf("stdout_bytes=%d stderr_bytes=%d stdout_lines=%d stderr_lines=%d", result.StdoutBytes, result.StderrBytes, result.StdoutLineCount, result.StderrLineCount)
	if result.MarkerObserved {
		return "expected smoke marker observed; " + streamSummary
	}
	return "expected smoke marker not observed; " + streamSummary
}

func runnerEnvironment(winePrefix string, wineArchitecture string) []string {
	if strings.TrimSpace(wineArchitecture) == "" {
		wineArchitecture = "unknown"
	}
	return append(
		os.Environ(),
		"WINEPREFIX="+winePrefix,
		"WINEARCH="+wineArchitecture,
		"WINEDEBUG=-all",
		"WINEDLLOVERRIDES=winemenubuilder.exe=d,mscoree=d,mshtml=d",
	)
}

func prepareWinePrefix(stateRoot string, wineArchitecture string) (string, string, error) {
	var prefixName string
	switch wineArchitecture {
	case "win64":
		prefixName = "wineprefix-win64"
	case "win32":
		prefixName = "wineprefix-win32"
	default:
		return "", "", fmt.Errorf("unsupported Wine architecture %q", wineArchitecture)
	}
	winePrefix := filepath.Join(stateRoot, prefixName)
	if err := os.MkdirAll(winePrefix, 0o700); err != nil {
		return "", "", fmt.Errorf("create architecture-scoped Wine prefix: %w", err)
	}
	return winePrefix, "architecture-scoped", nil
}

func wineArchitectureForExecutable(executableArchitecture string) string {
	switch executableArchitecture {
	case "x86_64":
		return "win64"
	case "x86":
		return "win32"
	default:
		return "unknown"
	}
}

func validateExecutable(path string) (string, error) {
	if strings.TrimSpace(path) == "" {
		return "", errors.New("executable path is required")
	}
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("resolve executable path: %w", err)
	}
	info, err := os.Stat(absolutePath)
	if err != nil {
		return "", fmt.Errorf("inspect executable: %w", err)
	}
	if info.IsDir() {
		return "", errors.New("executable path must be a file")
	}
	if strings.ToLower(filepath.Ext(absolutePath)) != ".exe" {
		return "", errors.New("executable path must point to a Windows .exe file")
	}
	return absolutePath, nil
}

func resolveRunner(path string) (string, error) {
	for _, candidate := range configuredRunnerCandidates(path, os.Getenv(RunnerEnvVar)) {
		if strings.ContainsRune(candidate.path, os.PathSeparator) {
			runnerPath, err := validateRunnerPath(candidate.path)
			if err == nil {
				return runnerPath, nil
			}
			continue
		}
		runnerPath, err := exec.LookPath(candidate.path)
		if err == nil {
			return runnerPath, nil
		}
	}
	return "", errors.New("windows compatibility runner unavailable")
}

func configuredRunnerCandidates(explicitRunner string, envRunner string) []runnerCandidate {
	explicitRunner = strings.TrimSpace(explicitRunner)
	if explicitRunner != "" {
		return []runnerCandidate{{
			id:     "explicit-runner",
			source: "operator-supplied-runner",
			path:   explicitRunner,
		}}
	}
	envRunner = strings.TrimSpace(envRunner)
	if envRunner != "" {
		return []runnerCandidate{{
			id:     "env-runner",
			source: "env-configured-runner",
			path:   envRunner,
		}}
	}
	return runnerCandidates()
}

func runnerCandidates() []runnerCandidate {
	candidates := []runnerCandidate{
		{id: "path-wine", source: "path-command", path: "wine"},
		{id: "path-wine64", source: "path-command", path: "wine64"},
	}
	if runtime.GOOS == "darwin" {
		candidates = append(candidates, darwinRunnerCandidates()...)
	}
	return candidates
}

func darwinRunnerCandidates() []runnerCandidate {
	candidates := []runnerCandidate{
		{id: "homebrew-wine", source: "macos-common-path", path: "/opt/homebrew/bin/wine"},
		{id: "homebrew-wine64", source: "macos-common-path", path: "/opt/homebrew/bin/wine64"},
		{id: "usr-local-wine", source: "macos-common-path", path: "/usr/local/bin/wine"},
		{id: "usr-local-wine64", source: "macos-common-path", path: "/usr/local/bin/wine64"},
		{id: "crossover-system-wine", source: "macos-crossover-bundle", path: "/Applications/CrossOver.app/Contents/SharedSupport/CrossOver/bin/wine"},
		{id: "crossover-system-wine64", source: "macos-crossover-bundle", path: "/Applications/CrossOver.app/Contents/SharedSupport/CrossOver/bin/wine64"},
		{id: "wine-stable-app-wine", source: "macos-app-bundle", path: "/Applications/Wine Stable.app/Contents/Resources/wine/bin/wine"},
		{id: "wine-stable-app-wine64", source: "macos-app-bundle", path: "/Applications/Wine Stable.app/Contents/Resources/wine/bin/wine64"},
		{id: "wine-devel-app-wine", source: "macos-app-bundle", path: "/Applications/Wine Devel.app/Contents/Resources/wine/bin/wine"},
		{id: "wine-devel-app-wine64", source: "macos-app-bundle", path: "/Applications/Wine Devel.app/Contents/Resources/wine/bin/wine64"},
		{id: "wine-staging-app-wine", source: "macos-app-bundle", path: "/Applications/Wine Staging.app/Contents/Resources/wine/bin/wine"},
		{id: "wine-staging-app-wine64", source: "macos-app-bundle", path: "/Applications/Wine Staging.app/Contents/Resources/wine/bin/wine64"},
	}
	homeDir, err := os.UserHomeDir()
	if err != nil || strings.TrimSpace(homeDir) == "" {
		return candidates
	}
	return append(candidates,
		runnerCandidate{id: "crossover-user-wine", source: "macos-crossover-user-bundle", path: filepath.Join(homeDir, "CrossOver.app", "Contents", "SharedSupport", "CrossOver", "bin", "wine")},
		runnerCandidate{id: "crossover-user-wine64", source: "macos-crossover-user-bundle", path: filepath.Join(homeDir, "CrossOver.app", "Contents", "SharedSupport", "CrossOver", "bin", "wine64")},
		runnerCandidate{id: "whisky-user-wine", source: "macos-whisky-library", path: filepath.Join(homeDir, "Library", "Application Support", "com.isaacmarovitz.Whisky", "Libraries", "Wine", "bin", "wine")},
		runnerCandidate{id: "whisky-user-wine64", source: "macos-whisky-library", path: filepath.Join(homeDir, "Library", "Application Support", "com.isaacmarovitz.Whisky", "Libraries", "Wine", "bin", "wine64")},
	)
}

func probeRunnerCandidate(candidate runnerCandidate) (RunnerCandidateEvidence, string) {
	evidence := RunnerCandidateEvidence{
		ID:     candidate.id,
		Source: candidate.source,
		Reason: "not-found",
	}
	var runnerPath string
	var err error
	if strings.ContainsRune(candidate.path, os.PathSeparator) {
		runnerPath, err = validateRunnerPath(candidate.path)
	} else {
		runnerPath, err = exec.LookPath(candidate.path)
	}
	if err != nil {
		if strings.Contains(err.Error(), "runner path must be a file") {
			evidence.Reason = "is-directory"
		}
		return evidence, ""
	}
	evidence.Available = true
	evidence.Reason = "available"
	return evidence, filepath.Base(runnerPath)
}

func resolveWineboot(runnerPath string) (string, bool) {
	runnerPath = strings.TrimSpace(runnerPath)
	if runnerPath == "" {
		return "", false
	}
	runnerDir := filepath.Dir(runnerPath)
	for _, candidate := range []string{
		filepath.Join(runnerDir, "wineboot"),
		filepath.Join(runnerDir, "wineboot64"),
	} {
		if path, err := validateRunnerPath(candidate); err == nil {
			return path, true
		}
	}
	if path, err := exec.LookPath("wineboot"); err == nil {
		return path, true
	}
	return "", false
}

func validateRunnerPath(path string) (string, error) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	info, err := os.Stat(absolutePath)
	if err != nil {
		return "", err
	}
	if info.IsDir() {
		return "", errors.New("runner path must be a file")
	}
	return absolutePath, nil
}

func exitCode(err error) int {
	if err == nil {
		return 0
	}
	var exitError *exec.ExitError
	if errors.As(err, &exitError) {
		return exitError.ExitCode()
	}
	return -1
}
