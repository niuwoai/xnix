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
	SchemaVersion = "xnix.runtime.windows_app_smoke.v1"
	RequestType   = "windows-app-run-smoke"
	PassedStatus  = "passed"
	FailedStatus  = "failed"
	SkippedStatus = "skipped"
	DefaultMarker = "XNIX_WINAPP_SMOKE_OK"
)

type Request struct {
	ExecutablePath string
	Arguments      []string
	StateRoot      string
	RunnerPath     string
	Timeout        time.Duration
	ExpectedMarker string
	RedactOutput   bool
}

type Result struct {
	SchemaVersion               string `json:"schema_version"`
	RequestType                 string `json:"request_type"`
	Status                      string `json:"status"`
	ExecutableName              string `json:"executable_name"`
	RunnerAvailable             bool   `json:"runner_available"`
	CompatibilityLayer          string `json:"compatibility_layer"`
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

	timeout := request.Timeout
	if timeout <= 0 {
		timeout = 20 * time.Second
	}
	runCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	args := append([]string{executablePath}, request.Arguments...)
	command := exec.CommandContext(runCtx, runnerPath, args...)
	command.Env = append(os.Environ(), "WINEPREFIX="+stateRoot)

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
	if strings.TrimSpace(path) != "" {
		return validateRunnerPath(path)
	}
	for _, candidate := range runnerCandidates() {
		if strings.ContainsRune(candidate, os.PathSeparator) {
			runnerPath, err := validateRunnerPath(candidate)
			if err == nil {
				return runnerPath, nil
			}
			continue
		}
		runnerPath, err := exec.LookPath(candidate)
		if err == nil {
			return runnerPath, nil
		}
	}
	return "", errors.New("windows compatibility runner unavailable")
}

func runnerCandidates() []string {
	candidates := []string{"wine", "wine64"}
	if runtime.GOOS == "darwin" {
		candidates = append(candidates,
			"/opt/homebrew/bin/wine",
			"/opt/homebrew/bin/wine64",
			"/usr/local/bin/wine",
			"/usr/local/bin/wine64",
			"/Applications/Wine Stable.app/Contents/Resources/wine/bin/wine",
			"/Applications/Wine Stable.app/Contents/Resources/wine/bin/wine64",
			"/Applications/Wine Devel.app/Contents/Resources/wine/bin/wine",
			"/Applications/Wine Devel.app/Contents/Resources/wine/bin/wine64",
			"/Applications/Wine Staging.app/Contents/Resources/wine/bin/wine",
			"/Applications/Wine Staging.app/Contents/Resources/wine/bin/wine64",
		)
	}
	return candidates
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
