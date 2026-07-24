package winapp

import (
	"bytes"
	"context"
	"errors"
	"fmt"
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
	DefaultMarker                  = "XNIX_WINAPP_SMOKE_OK"
)

type Request struct {
	ExecutablePath  string
	Arguments       []string
	RunnerArguments []string
	RunnerBottle    string
	StateRoot       string
	RunnerPath      string
	Timeout         time.Duration
	ExpectedMarker  string
	RedactOutput    bool
}

type Result struct {
	SchemaVersion               string `json:"schema_version"`
	RequestType                 string `json:"request_type"`
	Status                      string `json:"status"`
	ExecutableName              string `json:"executable_name"`
	RunnerAvailable             bool   `json:"runner_available"`
	RunnerArgumentCount         int    `json:"runner_argument_count"`
	CompatibilityLayer          string `json:"compatibility_layer"`
	WineBootstrapAttempted      bool   `json:"wine_bootstrap_attempted"`
	WineBootstrapSucceeded      bool   `json:"wine_bootstrap_succeeded"`
	WineBootstrapExitCode       int    `json:"wine_bootstrap_exit_code"`
	ExpectedMarker              string `json:"expected_marker"`
	MarkerObserved              bool   `json:"marker_observed"`
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

	executablePath, err := validateExecutable(request.ExecutablePath)
	if err != nil {
		return result, err
	}
	result.ExecutableName = filepath.Base(executablePath)

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

	runnerPath, err := resolveRunner(request.RunnerPath)
	if err != nil {
		result.Status = SkippedStatus
		result.SkipReason = "windows compatibility runner unavailable"
		return result, nil
	}
	result.RunnerAvailable = true
	runnerArguments := runnerInvocationArguments(request)
	result.RunnerArgumentCount = len(runnerArguments)

	timeout := request.Timeout
	if timeout <= 0 {
		timeout = 20 * time.Second
	}
	runCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	bootstrapPath, bootstrapAvailable := resolveWineboot(runnerPath)
	if bootstrapAvailable {
		result.WineBootstrapAttempted = true
		bootstrapArgs := append([]string{}, runnerArguments...)
		bootstrapArgs = append(bootstrapArgs, "--init")
		bootstrapCommand := exec.CommandContext(runCtx, bootstrapPath, bootstrapArgs...)
		bootstrapCommand.Env = runnerEnvironment(stateRoot)
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
	command.Env = runnerEnvironment(stateRoot)

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
		result.Status = FailedStatus
		result.FailureReason = "execution timed out"
		return result, nil
	}
	if err != nil {
		result.Status = FailedStatus
		result.FailureReason = "runner returned a non-zero exit status"
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
		CompatibilityLayer:          "windows-compatibility-layer",
		WineBootstrapExitCode:       -1,
		ExpectedMarker:              marker,
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

func runnerEnvironment(stateRoot string) []string {
	return append(
		os.Environ(),
		"WINEPREFIX="+stateRoot,
		"WINEARCH=win64",
		"WINEDEBUG=-all",
		"WINEDLLOVERRIDES=winemenubuilder.exe=d,mscoree=d,mshtml=d",
	)
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
