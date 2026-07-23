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
	GuestSchemaVersion = "xnix.runtime.windows_app_guest_wine_smoke.v1"
	GuestRequestType   = "windows-app-guest-wine-smoke"
	DefaultGuestHost   = "127.0.0.1"
	DefaultGuestPort   = "2222"
	DefaultGuestUser   = "root"
	DefaultRemoteDir   = "/tmp/xnix-winapp-smoke"
)

type GuestRequest struct {
	ExecutablePath string
	Arguments      []string
	Host           string
	Port           string
	User           string
	KeyPath        string
	RemoteDir      string
	SSHPath        string
	SCPPath        string
	Timeout        time.Duration
	ExpectedMarker string
}

type GuestResult struct {
	SchemaVersion               string `json:"schema_version"`
	RequestType                 string `json:"request_type"`
	Status                      string `json:"status"`
	ExecutableName              string `json:"executable_name"`
	GuestTransport              string `json:"guest_transport"`
	GuestReachable              bool   `json:"guest_reachable"`
	WineAvailable               bool   `json:"wine_available"`
	ExecutableCopied            bool   `json:"executable_copied"`
	ExpectedMarker              string `json:"expected_marker"`
	MarkerObserved              bool   `json:"marker_observed"`
	ExitCode                    int    `json:"exit_code"`
	DurationMillis              int64  `json:"duration_millis"`
	Stdout                      string `json:"stdout"`
	Stderr                      string `json:"stderr"`
	SkipReason                  string `json:"skip_reason,omitempty"`
	FailureReason               string `json:"failure_reason,omitempty"`
	LoopbackOnlyNetworking      bool   `json:"loopback_only_networking"`
	QEMURequired                bool   `json:"qemu_required"`
	HostRootModified            bool   `json:"host_root_modified"`
	PrivilegedContainerRequired bool   `json:"privileged_container_required"`
	HostNetworkingRequired      bool   `json:"host_networking_required"`
	DockerSocketMounted         bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired      bool   `json:"broad_host_mount_required"`
	RawHostPathExposed          bool   `json:"raw_host_path_exposed"`
}

func RunGuestSmoke(ctx context.Context, request GuestRequest) (GuestResult, error) {
	result := baseGuestResult(request)

	executablePath, err := validateExecutable(request.ExecutablePath)
	if err != nil {
		return result, err
	}
	result.ExecutableName = filepath.Base(executablePath)

	sshPath, err := resolveTool(request.SSHPath, "ssh")
	if err != nil {
		result.Status = SkippedStatus
		result.SkipReason = "guest ssh transport unavailable"
		return result, nil
	}
	scpPath, err := resolveTool(request.SCPPath, "scp")
	if err != nil {
		result.Status = SkippedStatus
		result.SkipReason = "guest scp transport unavailable"
		return result, nil
	}

	timeout := request.Timeout
	if timeout <= 0 {
		timeout = 60 * time.Second
	}
	runCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	startedAt := time.Now()
	remoteDir := guestRemoteDir(request.RemoteDir)
	remoteExe := remoteDir + "/" + filepath.Base(executablePath)
	guest := guestTarget(request)
	sshBase := guestSSHBaseArgs(request)

	if err := runGuestSSH(runCtx, sshPath, sshBase, guest, "true", nil, nil); err != nil {
		result.DurationMillis = time.Since(startedAt).Milliseconds()
		result.Status = SkippedStatus
		result.SkipReason = "guest ssh endpoint unavailable"
		result.ExitCode = exitCode(err)
		return result, nil
	}
	result.GuestReachable = true

	if err := runGuestSSH(runCtx, sshPath, sshBase, guest, "command -v wine >/dev/null 2>&1", nil, nil); err != nil {
		result.DurationMillis = time.Since(startedAt).Milliseconds()
		result.Status = SkippedStatus
		result.SkipReason = "guest wine runner unavailable"
		result.ExitCode = exitCode(err)
		return result, nil
	}
	result.WineAvailable = true

	if err := runGuestSSH(runCtx, sshPath, sshBase, guest, "mkdir -p "+shellQuote(remoteDir)+" && chmod 700 "+shellQuote(remoteDir), nil, nil); err != nil {
		result.DurationMillis = time.Since(startedAt).Milliseconds()
		result.Status = FailedStatus
		result.FailureReason = "guest work directory preparation failed"
		result.ExitCode = exitCode(err)
		return result, nil
	}

	if err := runGuestSCP(runCtx, scpPath, request, executablePath, guest+":"+remoteExe); err != nil {
		result.DurationMillis = time.Since(startedAt).Milliseconds()
		result.Status = FailedStatus
		result.FailureReason = "guest executable copy failed"
		result.ExitCode = exitCode(err)
		return result, nil
	}
	result.ExecutableCopied = true

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	remoteCommand := guestWineCommand(remoteDir, remoteExe, request.Arguments)
	err = runGuestSSH(runCtx, sshPath, sshBase, guest, remoteCommand, &stdout, &stderr)
	result.DurationMillis = time.Since(startedAt).Milliseconds()
	result.Stdout = stdout.String()
	result.Stderr = stderr.String()
	result.MarkerObserved = strings.Contains(result.Stdout, result.ExpectedMarker)
	result.ExitCode = exitCode(err)

	if runCtx.Err() == context.DeadlineExceeded {
		result.Status = FailedStatus
		result.FailureReason = "guest wine execution timed out"
		return result, nil
	}
	if err != nil {
		result.Status = FailedStatus
		result.FailureReason = "guest wine runner returned a non-zero exit status"
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

func baseGuestResult(request GuestRequest) GuestResult {
	marker := request.ExpectedMarker
	if strings.TrimSpace(marker) == "" {
		marker = DefaultMarker
	}
	return GuestResult{
		SchemaVersion:               GuestSchemaVersion,
		RequestType:                 GuestRequestType,
		Status:                      FailedStatus,
		GuestTransport:              "loopback-ssh",
		ExpectedMarker:              marker,
		ExitCode:                    -1,
		LoopbackOnlyNetworking:      true,
		QEMURequired:                true,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		HostNetworkingRequired:      false,
		DockerSocketMounted:         false,
		BroadHostMountRequired:      false,
		RawHostPathExposed:          false,
	}
}

func resolveTool(path string, name string) (string, error) {
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
			return "", fmt.Errorf("%s path must be a file", name)
		}
		return absolutePath, nil
	}
	return exec.LookPath(name)
}

func guestRemoteDir(remoteDir string) string {
	if strings.TrimSpace(remoteDir) == "" {
		return DefaultRemoteDir
	}
	return remoteDir
}

func guestTarget(request GuestRequest) string {
	user := request.User
	if strings.TrimSpace(user) == "" {
		user = DefaultGuestUser
	}
	host := request.Host
	if strings.TrimSpace(host) == "" {
		host = DefaultGuestHost
	}
	return user + "@" + host
}

func guestPort(request GuestRequest) string {
	if strings.TrimSpace(request.Port) == "" {
		return DefaultGuestPort
	}
	return request.Port
}

func guestSSHBaseArgs(request GuestRequest) []string {
	args := []string{
		"-p", guestPort(request),
		"-o", "BatchMode=yes",
		"-o", "ConnectTimeout=5",
		"-o", "IdentitiesOnly=yes",
		"-o", "StrictHostKeyChecking=no",
		"-o", "UserKnownHostsFile=/dev/null",
	}
	if strings.TrimSpace(request.KeyPath) != "" {
		args = append(args, "-i", request.KeyPath)
	}
	return args
}

func guestSCPBaseArgs(request GuestRequest) []string {
	args := []string{
		"-P", guestPort(request),
		"-o", "BatchMode=yes",
		"-o", "ConnectTimeout=5",
		"-o", "IdentitiesOnly=yes",
		"-o", "StrictHostKeyChecking=no",
		"-o", "UserKnownHostsFile=/dev/null",
	}
	if strings.TrimSpace(request.KeyPath) != "" {
		args = append(args, "-i", request.KeyPath)
	}
	return args
}

func runGuestSSH(ctx context.Context, sshPath string, baseArgs []string, guest string, remoteCommand string, stdout *bytes.Buffer, stderr *bytes.Buffer) error {
	args := append([]string{}, baseArgs...)
	args = append(args, guest, remoteCommand)
	command := exec.CommandContext(ctx, sshPath, args...)
	if stdout != nil {
		command.Stdout = stdout
	}
	if stderr != nil {
		command.Stderr = stderr
	}
	return command.Run()
}

func runGuestSCP(ctx context.Context, scpPath string, request GuestRequest, source string, destination string) error {
	args := guestSCPBaseArgs(request)
	args = append(args, source, destination)
	command := exec.CommandContext(ctx, scpPath, args...)
	return command.Run()
}

func guestWineCommand(remoteDir string, remoteExe string, appArgs []string) string {
	parts := []string{
		"WINEPREFIX=" + shellQuote(remoteDir+"/wineprefix"),
		"WINEDEBUG=-all",
		"wine",
		shellQuote(remoteExe),
	}
	for _, arg := range appArgs {
		parts = append(parts, shellQuote(arg))
	}
	return strings.Join(parts, " ")
}

func shellQuote(value string) string {
	if value == "" {
		return "''"
	}
	return "'" + strings.ReplaceAll(value, "'", "'\\''") + "'"
}

func ParseGuestPort(value string) (string, error) {
	if strings.TrimSpace(value) == "" {
		return DefaultGuestPort, nil
	}
	port, err := strconv.Atoi(value)
	if err != nil || port < 1 || port > 65535 {
		return "", errors.New("guest port must be between 1 and 65535")
	}
	return value, nil
}
