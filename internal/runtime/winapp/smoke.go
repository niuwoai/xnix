package winapp

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
	result.Stdout = stdout.String()
	result.Stderr = stderr.String()
	result.MarkerObserved = strings.Contains(result.Stdout, result.ExpectedMarker)
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
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		HostNetworkingRequired:      false,
		DockerSocketMounted:         false,
		BroadHostMountRequired:      false,
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
	if strings.TrimSpace(path) != "" {
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
	return exec.LookPath("wine")
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
