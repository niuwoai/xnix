package winapp

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	ContainerSchemaVersion      = "xnix.runtime.windows_app_container_smoke.v1"
	ContainerRequestType        = "windows-app-container-run-smoke"
	ContainerXGUISchemaVersion  = "xnix.runtime.windows_app_container_x_gui_smoke.v1"
	ContainerXGUIRequestType    = "windows-app-container-x-gui-smoke"
	DefaultContainerImage       = "xnix-wine-smoke:local"
	DefaultWinePlatform         = "linux/amd64"
	DefaultContainerGUIApp      = "notepad.exe"
	DefaultContainerWindowMatch = "notepad.exe"
	DefaultWineBootstrapTimeout = 300 * time.Second
	wineBootstrapExitMarker     = "XNIX_WINE_BOOTSTRAP_EXIT:"
)

type ContainerRequest struct {
	ExecutablePath   string
	Arguments        []string
	StateRoot        string
	Image            string
	Platform         string
	DockerPath       string
	Timeout          time.Duration
	BootstrapTimeout time.Duration
	ExpectedMarker   string
}

type ContainerXGUIRequest struct {
	ApplicationName string
	WindowMatch     string
	ApplicationID   string
	DisplayName     string
	AppVersion      string
	RecipeBacked    bool
	Image           string
	Platform        string
	DockerPath      string
	Timeout         time.Duration
}

type ContainerResult struct {
	SchemaVersion               string `json:"schema_version"`
	RequestType                 string `json:"request_type"`
	Status                      string `json:"status"`
	ExecutableName              string `json:"executable_name"`
	ContainerImage              string `json:"container_image"`
	ContainerPlatform           string `json:"container_platform"`
	ContainerStateMode          string `json:"container_state_mode"`
	PullPolicy                  string `json:"pull_policy"`
	NetworkMode                 string `json:"network_mode"`
	WineBootstrapRequired       bool   `json:"wine_bootstrap_required"`
	WineBootstrapTimedOut       bool   `json:"wine_bootstrap_timed_out"`
	WineBootstrapExitCode       int    `json:"wine_bootstrap_exit_code"`
	RunnerAvailable             bool   `json:"runner_available"`
	ImageAvailable              bool   `json:"image_available"`
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
	HostMountCount              int    `json:"host_mount_count"`
}

type ContainerXGUIResult struct {
	SchemaVersion               string `json:"schema_version"`
	RequestType                 string `json:"request_type"`
	Status                      string `json:"status"`
	ApplicationID               string `json:"application_id,omitempty"`
	DisplayName                 string `json:"display_name,omitempty"`
	AppVersion                  string `json:"app_version,omitempty"`
	RecipeBacked                bool   `json:"recipe_backed"`
	ApplicationName             string `json:"application_name"`
	WindowMatch                 string `json:"window_match"`
	ContainerImage              string `json:"container_image"`
	ContainerPlatform           string `json:"container_platform"`
	ContainerStateMode          string `json:"container_state_mode"`
	PullPolicy                  string `json:"pull_policy"`
	NetworkMode                 string `json:"network_mode"`
	DesktopDisplay              string `json:"desktop_display"`
	XServerStarted              bool   `json:"x_server_started"`
	WineBootstrapAttempted      bool   `json:"wine_bootstrap_attempted"`
	RunnerAvailable             bool   `json:"runner_available"`
	ImageAvailable              bool   `json:"image_available"`
	XWindowObserved             bool   `json:"x_window_observed"`
	WindowEvidenceSummary       string `json:"window_evidence_summary"`
	ExitCode                    int    `json:"exit_code"`
	DurationMillis              int64  `json:"duration_millis"`
	SkipReason                  string `json:"skip_reason,omitempty"`
	FailureReason               string `json:"failure_reason,omitempty"`
	HostRootModified            bool   `json:"host_root_modified"`
	PrivilegedContainerRequired bool   `json:"privileged_container_required"`
	HostNetworkingRequired      bool   `json:"host_networking_required"`
	DockerSocketMounted         bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired      bool   `json:"broad_host_mount_required"`
	HostMountCount              int    `json:"host_mount_count"`
}

func RunContainerSmoke(ctx context.Context, request ContainerRequest) (ContainerResult, error) {
	result := baseContainerResult(request)

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

	dockerPath, err := resolveDocker(request.DockerPath)
	if err != nil {
		result.Status = SkippedStatus
		result.SkipReason = "docker runner unavailable"
		return result, nil
	}
	result.RunnerAvailable = true

	imageCtx, imageCancel := context.WithTimeout(ctx, 10*time.Second)
	defer imageCancel()
	if err := inspectLocalImage(imageCtx, dockerPath, result.ContainerImage); err != nil {
		result.Status = SkippedStatus
		result.SkipReason = "local wine container image unavailable"
		return result, nil
	}
	result.ImageAvailable = true

	timeout := request.Timeout
	if timeout <= 0 {
		timeout = 45 * time.Second
	}
	runCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	bootstrapTimeout := request.BootstrapTimeout
	if bootstrapTimeout <= 0 {
		bootstrapTimeout = DefaultWineBootstrapTimeout
	}
	args := restrictedDockerRunArgs(
		executablePath,
		result.ContainerImage,
		result.ContainerPlatform,
		bootstrapTimeout,
		request.Arguments,
	)
	command := exec.CommandContext(runCtx, dockerPath, args...)

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
	result.WineBootstrapExitCode = parseWineBootstrapExitCode(result.Stderr)
	result.WineBootstrapTimedOut = result.WineBootstrapExitCode == 124

	if runCtx.Err() == context.DeadlineExceeded {
		result.Status = FailedStatus
		result.FailureReason = "container execution timed out"
		return result, nil
	}
	if result.WineBootstrapExitCode != -1 {
		result.Status = FailedStatus
		if result.WineBootstrapTimedOut {
			result.FailureReason = "wine bootstrap timed out"
		} else {
			result.FailureReason = "wine bootstrap failed"
		}
		return result, nil
	}
	if err != nil {
		result.Status = FailedStatus
		result.FailureReason = "container runner returned a non-zero exit status"
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

func RunContainerXGUISmoke(ctx context.Context, request ContainerXGUIRequest) (ContainerXGUIResult, error) {
	result := baseContainerXGUIResult(request)

	dockerPath, err := resolveDocker(request.DockerPath)
	if err != nil {
		result.Status = SkippedStatus
		result.SkipReason = "docker runner unavailable"
		return result, nil
	}
	result.RunnerAvailable = true

	imageCtx, imageCancel := context.WithTimeout(ctx, 10*time.Second)
	defer imageCancel()
	if err := inspectLocalImage(imageCtx, dockerPath, result.ContainerImage); err != nil {
		result.Status = SkippedStatus
		result.SkipReason = "local Wine X GUI container image unavailable"
		return result, nil
	}
	result.ImageAvailable = true

	timeout := request.Timeout
	if timeout <= 0 {
		timeout = 90 * time.Second
	}
	runCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	args := restrictedDockerXGUIRunArgs(result.ApplicationName, result.WindowMatch, result.ContainerImage, result.ContainerPlatform)
	command := exec.CommandContext(runCtx, dockerPath, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	command.Stdout = &stdout
	command.Stderr = &stderr

	startedAt := time.Now()
	err = command.Run()
	result.DurationMillis = time.Since(startedAt).Milliseconds()
	output := stdout.String() + "\n" + stderr.String()
	result.ExitCode = exitCode(err)
	result.XServerStarted = strings.Contains(output, "XNIX_X_GUI_XSERVER_STARTED=true")
	result.WineBootstrapAttempted = strings.Contains(output, "XNIX_X_GUI_WINE_BOOTSTRAP_ATTEMPTED=true")
	result.XWindowObserved = strings.Contains(output, "XNIX_X_GUI_WINDOW_OBSERVED=true") ||
		strings.Contains(strings.ToLower(output), strings.ToLower(result.WindowMatch))
	result.WindowEvidenceSummary = summarizeXWindowEvidence(output, result.WindowMatch)

	if runCtx.Err() == context.DeadlineExceeded {
		result.Status = FailedStatus
		result.FailureReason = "container X GUI execution timed out"
		return result, nil
	}
	if err != nil {
		result.Status = FailedStatus
		result.FailureReason = "container X GUI runner returned a non-zero exit status"
		return result, nil
	}
	if !result.XWindowObserved {
		result.Status = FailedStatus
		result.FailureReason = "expected X window was not observed"
		return result, nil
	}

	result.Status = PassedStatus
	return result, nil
}

func baseContainerResult(request ContainerRequest) ContainerResult {
	marker := request.ExpectedMarker
	if strings.TrimSpace(marker) == "" {
		marker = DefaultMarker
	}
	image := request.Image
	if strings.TrimSpace(image) == "" {
		image = DefaultContainerImage
	}
	platform := request.Platform
	if strings.TrimSpace(platform) == "" {
		platform = DefaultWinePlatform
	}
	return ContainerResult{
		SchemaVersion:               ContainerSchemaVersion,
		RequestType:                 ContainerRequestType,
		Status:                      FailedStatus,
		ContainerImage:              image,
		ContainerPlatform:           platform,
		ContainerStateMode:          "tmpfs",
		PullPolicy:                  "never",
		NetworkMode:                 "none",
		WineBootstrapRequired:       true,
		WineBootstrapTimedOut:       false,
		WineBootstrapExitCode:       -1,
		CompatibilityLayer:          "containerized-windows-compatibility-layer",
		ExpectedMarker:              marker,
		ExitCode:                    -1,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		HostNetworkingRequired:      false,
		DockerSocketMounted:         false,
		BroadHostMountRequired:      false,
		HostMountCount:              1,
	}
}

func baseContainerXGUIResult(request ContainerXGUIRequest) ContainerXGUIResult {
	appName := strings.TrimSpace(request.ApplicationName)
	if appName == "" {
		appName = DefaultContainerGUIApp
	}
	windowMatch := strings.TrimSpace(request.WindowMatch)
	if windowMatch == "" {
		windowMatch = DefaultContainerWindowMatch
	}
	image := request.Image
	if strings.TrimSpace(image) == "" {
		image = DefaultContainerImage
	}
	platform := request.Platform
	if strings.TrimSpace(platform) == "" {
		platform = DefaultWinePlatform
	}
	return ContainerXGUIResult{
		SchemaVersion:               ContainerXGUISchemaVersion,
		RequestType:                 ContainerXGUIRequestType,
		Status:                      FailedStatus,
		ApplicationID:               strings.TrimSpace(request.ApplicationID),
		DisplayName:                 strings.TrimSpace(request.DisplayName),
		AppVersion:                  strings.TrimSpace(request.AppVersion),
		RecipeBacked:                request.RecipeBacked,
		ApplicationName:             appName,
		WindowMatch:                 windowMatch,
		ContainerImage:              image,
		ContainerPlatform:           platform,
		ContainerStateMode:          "tmpfs",
		PullPolicy:                  "never",
		NetworkMode:                 "none",
		DesktopDisplay:              "Xvfb",
		ExitCode:                    -1,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		HostNetworkingRequired:      false,
		DockerSocketMounted:         false,
		BroadHostMountRequired:      false,
		HostMountCount:              0,
	}
}

func resolveDocker(path string) (string, error) {
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
			return "", errors.New("docker path must be a file")
		}
		return absolutePath, nil
	}
	return exec.LookPath("docker")
}

func inspectLocalImage(ctx context.Context, dockerPath string, image string) error {
	command := exec.CommandContext(ctx, dockerPath, "image", "inspect", image)
	return command.Run()
}

func restrictedDockerXGUIRunArgs(appName string, windowMatch string, image string, platform string) []string {
	launcher := strings.Join([]string{
		"set -eu",
		"Xvfb \"$DISPLAY\" -screen 0 1024x768x24 -ac >/tmp/xvfb.log 2>&1 &",
		"xvfb_pid=$!",
		"cleanup() { kill \"$xvfb_pid\" >/dev/null 2>&1 || true; wineserver -k >/dev/null 2>&1 || true; }",
		"trap cleanup EXIT",
		"sleep 1",
		"xdpyinfo >/dev/null",
		"printf 'XNIX_X_GUI_XSERVER_STARTED=true\\n'",
		"wineboot --init >/tmp/wineboot.log 2>&1 || true",
		"printf 'XNIX_X_GUI_WINE_BOOTSTRAP_ATTEMPTED=true\\n'",
		"wine \"$XNIX_GUI_APP\" >/tmp/wine-gui.log 2>&1 &",
		"app_pid=$!",
		"observed=0",
		"for _attempt in $(seq 1 20); do",
		"  if xwininfo -root -tree 2>/dev/null | grep -i -- \"$XNIX_WINDOW_MATCH\" >/dev/null; then observed=1; break; fi",
		"  sleep 1",
		"done",
		"xwininfo -root -tree | sed -n '1,80p'",
		"kill \"$app_pid\" >/dev/null 2>&1 || true",
		"if [ \"$observed\" != \"1\" ]; then printf 'XNIX_X_GUI_WINDOW_OBSERVED=false\\n'; exit 11; fi",
		"printf 'XNIX_X_GUI_WINDOW_OBSERVED=true\\n'",
	}, "\n")
	args := []string{
		"run",
		"--rm",
		"--pull", "never",
		"--network", "none",
		"--cpus", "2",
		"--memory", "3g",
		"--pids-limit", "512",
		"--security-opt", "no-new-privileges",
		"--cap-drop", "ALL",
		"--tmpfs", "/tmp:rw,nosuid,nodev,size=256m",
		"--tmpfs", "/state:rw,nosuid,nodev,size=1g",
		"--env", "WINEPREFIX=/state/wineprefix",
		"--env", "HOME=/state/home",
		"--env", "WINEDEBUG=-all",
		"--env", "WINEDLLOVERRIDES=winemenubuilder.exe=d,mscoree=d,mshtml=d",
		"--env", "DISPLAY=:99",
		"--env", "XNIX_GUI_APP=" + appName,
		"--env", "XNIX_WINDOW_MATCH=" + windowMatch,
	}
	if strings.TrimSpace(platform) != "" {
		args = append(args, "--platform", platform)
	}
	args = append(args, image, "sh", "-lc", launcher)
	return args
}

func restrictedDockerRunArgs(executablePath string, image string, platform string, bootstrapTimeout time.Duration, appArgs []string) []string {
	workRoot := filepath.Dir(executablePath)
	executableName := filepath.Base(executablePath)
	bootstrapSeconds := int(bootstrapTimeout.Round(time.Second).Seconds())
	if bootstrapSeconds < 1 {
		bootstrapSeconds = 1
	}
	launcher := strings.Join([]string{
		"timeout \"${XNIX_WINE_BOOTSTRAP_TIMEOUT_SECONDS}s\" wineboot --init",
		"bootstrap_status=$?",
		"if [ \"$bootstrap_status\" -ne 0 ]; then",
		"printf '" + wineBootstrapExitMarker + "%s\\n' \"$bootstrap_status\" >&2",
		"exit \"$bootstrap_status\"",
		"fi",
		"exec wine \"$@\"",
	}, "\n")
	args := []string{
		"run",
		"--rm",
		"--platform", platform,
		"--pull", "never",
		"--network", "none",
		"--cpus", "2",
		"--memory", "2g",
		"--pids-limit", "256",
		"--security-opt", "no-new-privileges",
		"--cap-drop", "ALL",
		"--tmpfs", "/tmp:rw,nosuid,nodev,size=128m",
		"--tmpfs", "/state:rw,nosuid,nodev,size=768m",
		"--volume", workRoot + ":/work:ro",
		"--env", "WINEPREFIX=/state/wineprefix",
		"--env", "WINEARCH=win64",
		"--env", "HOME=/state/home",
		"--env", "WINEDEBUG=-all",
		"--env", "WINEDLLOVERRIDES=winemenubuilder.exe=d,mscoree=d,mshtml=d",
		"--env", "XNIX_WINE_BOOTSTRAP_TIMEOUT_SECONDS=" + strconv.Itoa(bootstrapSeconds),
		"--workdir", "/work",
		image,
		"sh",
		"-lc",
		launcher,
		"xnix-wine-smoke",
		"/work/" + executableName,
	}
	return append(args, appArgs...)
}

func summarizeXWindowEvidence(output string, windowMatch string) string {
	match := strings.ToLower(windowMatch)
	for _, line := range strings.Split(output, "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		lowered := strings.ToLower(trimmed)
		if strings.Contains(lowered, match) || strings.Contains(trimmed, "XNIX_X_GUI_WINDOW_OBSERVED=true") {
			return trimmed
		}
	}
	return ""
}

func parseWineBootstrapExitCode(stderr string) int {
	index := strings.LastIndex(stderr, wineBootstrapExitMarker)
	if index == -1 {
		return -1
	}
	start := index + len(wineBootstrapExitMarker)
	end := start
	for end < len(stderr) && stderr[end] >= '0' && stderr[end] <= '9' {
		end++
	}
	if end == start {
		return -1
	}
	code, err := strconv.Atoi(stderr[start:end])
	if err != nil {
		return -1
	}
	return code
}
