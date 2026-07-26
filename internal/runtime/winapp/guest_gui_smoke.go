package winapp

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	GuestGUISchemaVersion    = "xnix.runtime.windows_app_guest_wine_gui_smoke.v1"
	GuestGUIRequestType      = "windows-app-guest-wine-gui-smoke"
	DefaultGuestGUIApp       = "/usr/lib/wine/i386-windows/winemine.exe"
	DefaultGuestGUIDisplay   = "10.0.2.2:100"
	DefaultHostGUIDisplay    = ":100"
	GuestGUIWineDLLOVERRIDES = "winemenubuilder.exe=d,mscoree,mshtml="
	GuestGUIWineDebug        = "err+winediag"
)

type GuestGUIRequest struct {
	ExecutablePath    string
	GUIAppPath        string
	KnownAppID        string
	KnownAppName      string
	KnownAppVersion   string
	Host              string
	Port              string
	User              string
	KeyPath           string
	RemoteDir         string
	SSHPath           string
	SCPPath           string
	XWinInfoPath      string
	GuestDisplay      string
	HostDisplay       string
	FileArgumentPaths []string
	WindowMatch       string
	Timeout           time.Duration
	Wait              time.Duration
}

type GuestGUIResult struct {
	SchemaVersion                             string `json:"schema_version"`
	RequestType                               string `json:"request_type"`
	Status                                    string `json:"status"`
	KnownAppID                                string `json:"known_app_id,omitempty"`
	KnownAppName                              string `json:"known_app_name,omitempty"`
	KnownAppVersion                           string `json:"known_app_version,omitempty"`
	GUIAppName                                string `json:"gui_app_name"`
	Backend                                   string `json:"backend"`
	GuestTransport                            string `json:"guest_transport"`
	GuestReachable                            bool   `json:"guest_reachable"`
	WineAvailable                             bool   `json:"wine_available"`
	GuestX11DriverAvailable                   bool   `json:"guest_x11_driver_available"`
	WinebootInvoked                           bool   `json:"wineboot_invoked"`
	WinebootExitCode                          int    `json:"wineboot_exit_code"`
	ExecutableCopied                          bool   `json:"executable_copied"`
	FileArgumentCount                         int    `json:"file_argument_count"`
	FileArgumentCopiedCount                   int    `json:"file_argument_copied_count"`
	FileArgumentWinePathTranslated            bool   `json:"file_argument_winepath_translated"`
	FileArgumentWinePathTranslatedCount       int    `json:"file_argument_winepath_translated_count"`
	FileArgumentsPassed                       bool   `json:"file_arguments_passed"`
	LaunchAttempted                           bool   `json:"launch_attempted"`
	LaunchPIDRecorded                         bool   `json:"launch_pid_recorded"`
	XWinInfoInvoked                           bool   `json:"xwininfo_invoked"`
	XWindowObserved                           bool   `json:"x_window_observed"`
	WindowMatch                               string `json:"window_match,omitempty"`
	WindowMatchObserved                       bool   `json:"window_match_observed"`
	WindowEvidenceSummary                     string `json:"window_evidence_summary"`
	XWindowChildCount                         int    `json:"x_window_child_count"`
	XWindowObservationAttempts                int    `json:"x_window_observation_attempts"`
	XWinInfoBytes                             int    `json:"xwininfo_bytes"`
	WinebootStderrBytes                       int    `json:"wineboot_stderr_bytes"`
	GuestStderrBytes                          int    `json:"guest_stderr_bytes"`
	GuestGraphicsDriverErrorObserved          bool   `json:"guest_graphics_driver_error_observed"`
	DurationMillis                            int64  `json:"duration_millis"`
	KDESafeOutputSummary                      string `json:"kde_safe_output_summary"`
	SkipReason                                string `json:"skip_reason,omitempty"`
	FailureReason                             string `json:"failure_reason,omitempty"`
	LoopbackSSHForwardingOnly                 bool   `json:"loopback_ssh_forwarding_only"`
	QEMURequired                              bool   `json:"qemu_required"`
	XvfbRequired                              bool   `json:"xvfb_required"`
	QEMUUserNetworkRestrictDisabledForDisplay bool   `json:"qemu_user_network_restrict_disabled_for_display"`
	HostRootModified                          bool   `json:"host_root_modified"`
	PrivilegedContainerRequired               bool   `json:"privileged_container_required"`
	HostNetworkingRequired                    bool   `json:"host_networking_required"`
	DockerSocketMounted                       bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired                    bool   `json:"broad_host_mount_required"`
	RawHostPathExposed                        bool   `json:"raw_host_path_exposed"`
	RawGuestGUIAppPathExposed                 bool   `json:"raw_guest_gui_app_path_exposed"`
	RawFileArgumentPathExposed                bool   `json:"raw_file_argument_path_exposed"`
	RawCommandExposed                         bool   `json:"raw_command_exposed"`
}

func RunGuestGUISmoke(ctx context.Context, request GuestGUIRequest) (GuestGUIResult, error) {
	result := baseGuestGUIResult(request)
	executablePath := strings.TrimSpace(request.ExecutablePath)
	fileArgumentPaths, err := validateGuestGUIFileArgumentPaths(request.FileArgumentPaths)
	if err != nil {
		return result, err
	}
	if executablePath != "" {
		validatedExecutablePath, err := validateExecutable(executablePath)
		if err != nil {
			return result, err
		}
		executablePath = validatedExecutablePath
		result.GUIAppName = filepath.Base(executablePath)
	}
	sshPath, err := resolveTool(request.SSHPath, "ssh")
	if err != nil {
		result.Status = SkippedStatus
		result.SkipReason = "guest ssh transport unavailable"
		return result, nil
	}
	scpPath := ""
	if executablePath != "" || len(fileArgumentPaths) > 0 {
		scpPath, err = resolveTool(request.SCPPath, "scp")
		if err != nil {
			result.Status = SkippedStatus
			result.SkipReason = "guest scp transport unavailable"
			return result, nil
		}
	}
	xwininfoPath, err := resolveTool(request.XWinInfoPath, "xwininfo")
	if err != nil {
		result.Status = SkippedStatus
		result.SkipReason = "host xwininfo unavailable"
		return result, nil
	}

	timeout := request.Timeout
	if timeout <= 0 {
		timeout = 90 * time.Second
	}
	runCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	startedAt := time.Now()
	remoteDir := guestRemoteDir(request.RemoteDir)
	guest := guestTarget(GuestRequest{Host: request.Host, Port: request.Port, User: request.User})
	sshBase := guestSSHBaseArgs(GuestRequest{Host: request.Host, Port: request.Port, User: request.User, KeyPath: request.KeyPath})

	if err := runGuestSSH(runCtx, sshPath, sshBase, guest, "true", nil, nil); err != nil {
		result.DurationMillis = time.Since(startedAt).Milliseconds()
		result.Status = SkippedStatus
		result.SkipReason = "guest ssh endpoint unavailable"
		result.WinebootExitCode = exitCode(err)
		return result, nil
	}
	result.GuestReachable = true

	if err := runGuestSSH(runCtx, sshPath, sshBase, guest, "command -v wine >/dev/null 2>&1 && command -v wineboot >/dev/null 2>&1", nil, nil); err != nil {
		result.DurationMillis = time.Since(startedAt).Milliseconds()
		result.Status = SkippedStatus
		result.SkipReason = "guest wine GUI runner unavailable"
		result.WinebootExitCode = exitCode(err)
		return result, nil
	}
	result.WineAvailable = true

	if err := runGuestSSH(runCtx, sshPath, sshBase, guest, guestGUIX11DriverCheckCommand(), nil, nil); err != nil {
		result.DurationMillis = time.Since(startedAt).Milliseconds()
		result.Status = FailedStatus
		result.FailureReason = "guest Wine X11 graphics driver unavailable"
		result.WinebootExitCode = exitCode(err)
		result.KDESafeOutputSummary = guestGUIKDESafeOutputSummary(result)
		return result, nil
	}
	result.GuestX11DriverAvailable = true

	if err := runGuestSSH(runCtx, sshPath, sshBase, guest, "mkdir -p "+shellQuote(remoteDir)+" && chmod 700 "+shellQuote(remoteDir), nil, nil); err != nil {
		result.DurationMillis = time.Since(startedAt).Milliseconds()
		result.Status = FailedStatus
		result.FailureReason = "guest work directory preparation failed"
		result.WinebootExitCode = exitCode(err)
		return result, nil
	}

	remoteGUIApp := guiAppPath(request)
	if executablePath != "" {
		remoteGUIApp = remoteDir + "/" + filepath.Base(executablePath)
		if err := runGuestSCP(runCtx, scpPath, GuestRequest{Host: request.Host, Port: request.Port, User: request.User, KeyPath: request.KeyPath}, executablePath, guest+":"+remoteGUIApp); err != nil {
			result.DurationMillis = time.Since(startedAt).Milliseconds()
			result.Status = FailedStatus
			result.FailureReason = "guest GUI executable copy failed"
			result.WinebootExitCode = exitCode(err)
			return result, nil
		}
		result.ExecutableCopied = true
	}
	remoteFileArguments, err := copyGuestGUIFileArguments(runCtx, scpPath, GuestRequest{Host: request.Host, Port: request.Port, User: request.User, KeyPath: request.KeyPath}, remoteDir, guest, fileArgumentPaths)
	if err != nil {
		result.DurationMillis = time.Since(startedAt).Milliseconds()
		result.Status = FailedStatus
		result.FailureReason = "guest GUI file argument copy failed"
		result.WinebootExitCode = exitCode(err)
		return result, nil
	}
	result.FileArgumentCopiedCount = len(remoteFileArguments)

	winebootCommand := guestGUIWinebootCommand(remoteDir, guestDisplay(request))
	var winebootStderr bytes.Buffer
	winebootErr := runGuestSSH(runCtx, sshPath, sshBase, guest, winebootCommand, nil, &winebootStderr)
	result.WinebootInvoked = true
	result.WinebootExitCode = exitCode(winebootErr)
	winebootText := winebootStderr.String()
	result.WinebootStderrBytes = len(winebootText)
	if winebootErr != nil && runCtx.Err() == context.DeadlineExceeded {
		result.DurationMillis = time.Since(startedAt).Milliseconds()
		result.Status = FailedStatus
		result.FailureReason = "guest wineboot timed out"
		result.KDESafeOutputSummary = guestGUIKDESafeOutputSummary(result)
		return result, nil
	}

	launchCommand := guestGUIWineLaunchCommand(remoteDir, guestDisplay(request), remoteGUIApp, remoteFileArguments)
	var launchStderr bytes.Buffer
	launchErr := runGuestSSH(runCtx, sshPath, sshBase, guest, launchCommand, nil, &launchStderr)
	result.LaunchAttempted = true
	result.LaunchPIDRecorded = launchErr == nil
	result.FileArgumentWinePathTranslatedCount = readGuestIntFile(runCtx, sshPath, sshBase, guest, remoteDir+"/file_args_winepath_translated")
	result.FileArgumentWinePathTranslated = result.FileArgumentCount > 0 && result.FileArgumentWinePathTranslatedCount == result.FileArgumentCount
	result.FileArgumentsPassed = result.FileArgumentCount > 0 && readGuestIntFile(runCtx, sshPath, sshBase, guest, remoteDir+"/file_args_passed") == result.FileArgumentCount
	if launchErr != nil {
		result.DurationMillis = time.Since(startedAt).Milliseconds()
		result.Status = FailedStatus
		result.FailureReason = "guest Wine GUI launch command failed"
		result.GuestStderrBytes = launchStderr.Len()
		result.GuestGraphicsDriverErrorObserved = guiDriverErrorObserved(winebootText, launchStderr.String())
		result.KDESafeOutputSummary = guestGUIKDESafeOutputSummary(result)
		return result, nil
	}

	observeGuestGUIWindow(runCtx, &result, guestGUIObservationRequest{
		Wait:         request.Wait,
		SSHPath:      sshPath,
		SSHBase:      sshBase,
		Guest:        guest,
		XWinInfoPath: xwininfoPath,
		HostDisplay:  hostDisplay(request),
		WindowMatch:  strings.TrimSpace(request.WindowMatch),
	})
	if runCtx.Err() == context.DeadlineExceeded {
		result.DurationMillis = time.Since(startedAt).Milliseconds()
		result.Status = FailedStatus
		result.FailureReason = "guest Wine GUI observation timed out"
		result.KDESafeOutputSummary = guestGUIKDESafeOutputSummary(result)
		return result, nil
	}

	guestStderrText := readGuestFile(runCtx, sshPath, sshBase, guest, remoteDir+"/stderr.txt")
	result.GuestStderrBytes = len(guestStderrText)
	result.GuestGraphicsDriverErrorObserved = guiDriverErrorObserved(winebootText, guestStderrText)
	_ = runGuestSSH(context.Background(), sshPath, sshBase, guest, "wineserver -k 2>/dev/null || true", nil, nil)

	result.DurationMillis = time.Since(startedAt).Milliseconds()
	result.KDESafeOutputSummary = guestGUIKDESafeOutputSummary(result)
	if result.XWindowObserved {
		result.Status = PassedStatus
		result.FailureReason = ""
		return result, nil
	}
	result.Status = FailedStatus
	if strings.TrimSpace(result.FailureReason) == "" {
		result.FailureReason = "Wine GUI app did not create an X window"
	}
	if result.GuestGraphicsDriverErrorObserved {
		result.FailureReason = "guest Wine X11 graphics driver unavailable"
	}
	return result, nil
}

func baseGuestGUIResult(request GuestGUIRequest) GuestGUIResult {
	return GuestGUIResult{
		SchemaVersion:             GuestGUISchemaVersion,
		RequestType:               GuestGUIRequestType,
		Status:                    FailedStatus,
		KnownAppID:                strings.TrimSpace(request.KnownAppID),
		KnownAppName:              strings.TrimSpace(request.KnownAppName),
		KnownAppVersion:           strings.TrimSpace(request.KnownAppVersion),
		GUIAppName:                filepath.Base(guiAppPath(request)),
		Backend:                   "qemu-guest-wine-x11",
		GuestTransport:            "loopback-ssh",
		WinebootExitCode:          -1,
		LoopbackSSHForwardingOnly: true,
		QEMURequired:              true,
		XvfbRequired:              true,
		QEMUUserNetworkRestrictDisabledForDisplay: true,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		HostNetworkingRequired:      false,
		DockerSocketMounted:         false,
		BroadHostMountRequired:      false,
		RawHostPathExposed:          false,
		RawGuestGUIAppPathExposed:   false,
		RawFileArgumentPathExposed:  false,
		RawCommandExposed:           false,
		FileArgumentCount:           len(request.FileArgumentPaths),
		WindowMatch:                 strings.TrimSpace(request.WindowMatch),
	}
}

func guiAppPath(request GuestGUIRequest) string {
	if strings.TrimSpace(request.GUIAppPath) == "" {
		return DefaultGuestGUIApp
	}
	return request.GUIAppPath
}

func guestDisplay(request GuestGUIRequest) string {
	if strings.TrimSpace(request.GuestDisplay) == "" {
		return DefaultGuestGUIDisplay
	}
	return request.GuestDisplay
}

func hostDisplay(request GuestGUIRequest) string {
	if strings.TrimSpace(request.HostDisplay) == "" {
		return DefaultHostGUIDisplay
	}
	return request.HostDisplay
}

func guestGUIWinebootCommand(remoteDir string, display string) string {
	return strings.Join([]string{
		"(",
		"DISPLAY=" + shellQuote(display),
		"WINEPREFIX=" + shellQuote(remoteDir+"/wineprefix"),
		"WINEDEBUG=" + shellQuote(GuestGUIWineDebug),
		"WINEDLLOVERRIDES=" + shellQuote(GuestGUIWineDLLOVERRIDES),
		"wineboot",
		"--init",
		"&",
		"wineboot_pid=$!",
		";",
		"while kill -0 \"$wineboot_pid\" 2>/dev/null; do",
		guestGUIWineInstallerSuppressCommand(),
		"sleep 1",
		";",
		"done",
		";",
		"wait \"$wineboot_pid\"",
		")",
	}, " ")
}

func guestGUIWineLaunchCommand(remoteDir string, display string, guiApp string, fileArguments []string) string {
	command := []string{
		"DISPLAY=" + shellQuote(display),
		"WINEPREFIX=" + shellQuote(remoteDir+"/wineprefix"),
		"WINEDEBUG=" + shellQuote(GuestGUIWineDebug),
		"WINEDLLOVERRIDES=" + shellQuote(GuestGUIWineDLLOVERRIDES),
		";",
		"export DISPLAY WINEPREFIX WINEDEBUG WINEDLLOVERRIDES",
		";",
		"_xnix_file_args_winepath_translated=0",
		";",
		"set --",
	}
	for _, path := range fileArguments {
		command = append(command,
			";",
			"_xnix_file_arg=$(winepath -w "+shellQuote(path)+" 2>/dev/null || true)",
			";",
			"if [ -n \"$_xnix_file_arg\" ]; then",
			"_xnix_file_args_winepath_translated=$((_xnix_file_args_winepath_translated + 1))",
			";",
			"set -- \"$@\" \"$_xnix_file_arg\"",
			";",
			"else",
			"set -- \"$@\" "+shellQuote(path),
			";",
			"fi",
		)
	}
	command = append(command,
		";",
		"printf '%s\\n' \"$#\" >"+shellQuote(remoteDir+"/file_args_passed"),
		";",
		"printf '%s\\n' \"$_xnix_file_args_winepath_translated\" >"+shellQuote(remoteDir+"/file_args_winepath_translated"),
		";",
		"wine",
		shellQuote(guiApp),
		"\"$@\"",
		">"+shellQuote(remoteDir+"/stdout.txt"),
		"2>"+shellQuote(remoteDir+"/stderr.txt"),
		"&",
		"printf",
		"'%s\\n'",
		"\"$!\"",
		">"+shellQuote(remoteDir+"/pid"),
	)
	return strings.Join(command, " ")
}

func guestGUIX11DriverCheckCommand() string {
	return "test -e /usr/lib/wine/i386-unix/winex11.so || test -e /usr/lib/wine/i386-unix/winex11.drv.so"
}

func guestGUIWineInstallerSuppressCommand() string {
	return "ps w | grep 'appwiz[.]cpl install_mono' | while read pid rest; do kill \"$pid\" 2>/dev/null || true; done;"
}

type guestGUIObservationRequest struct {
	Wait         time.Duration
	SSHPath      string
	SSHBase      []string
	Guest        string
	XWinInfoPath string
	HostDisplay  string
	WindowMatch  string
}

func observeGuestGUIWindow(ctx context.Context, result *GuestGUIResult, request guestGUIObservationRequest) {
	wait := request.Wait
	if wait <= 0 {
		wait = 10 * time.Second
	}
	deadline := time.Now().Add(wait)
	for {
		result.XWindowObservationAttempts++
		_ = runGuestSSH(ctx, request.SSHPath, request.SSHBase, request.Guest, guestGUIWineInstallerSuppressCommand(), nil, nil)
		xwininfoText, xwininfoStderr, xwininfoErr := runXWinInfo(ctx, request.XWinInfoPath, request.HostDisplay)
		result.XWinInfoInvoked = true
		result.XWinInfoBytes = len(xwininfoText)
		result.XWindowChildCount = countXWindowChildren(xwininfoText)
		result.WindowMatchObserved = guestXWindowMatchObserved(xwininfoText, request.WindowMatch)
		result.WindowEvidenceSummary = summarizeGuestXWindowEvidence(xwininfoText, request.WindowMatch)
		if strings.TrimSpace(request.WindowMatch) == "" {
			result.XWindowObserved = result.XWindowChildCount > 0
		} else {
			result.XWindowObserved = result.WindowMatchObserved
		}
		result.FailureReason = xwininfoStderr
		if result.XWindowObserved || ctx.Err() != nil {
			return
		}
		if xwininfoErr != nil && strings.TrimSpace(xwininfoStderr) != "" {
			return
		}
		remaining := time.Until(deadline)
		if remaining <= 0 {
			return
		}
		sleepFor := time.Second
		if remaining < sleepFor {
			sleepFor = remaining
		}
		timer := time.NewTimer(sleepFor)
		select {
		case <-ctx.Done():
			timer.Stop()
			return
		case <-timer.C:
		}
	}
}

func validateGuestGUIFileArgumentPaths(paths []string) ([]string, error) {
	validated := make([]string, 0, len(paths))
	for _, path := range paths {
		trimmed := strings.TrimSpace(path)
		if trimmed == "" {
			return nil, fmt.Errorf("guest GUI file argument path is required")
		}
		info, err := os.Stat(trimmed)
		if err != nil {
			return nil, fmt.Errorf("validate guest GUI file argument: %w", err)
		}
		if info.IsDir() {
			return nil, fmt.Errorf("guest GUI file argument must be a file")
		}
		validated = append(validated, trimmed)
	}
	return validated, nil
}

func copyGuestGUIFileArguments(ctx context.Context, scpPath string, guestRequest GuestRequest, remoteDir string, guest string, paths []string) ([]string, error) {
	if len(paths) == 0 {
		return nil, nil
	}
	if strings.TrimSpace(scpPath) == "" {
		return nil, fmt.Errorf("guest scp transport unavailable")
	}
	remotePaths := make([]string, 0, len(paths))
	for index, path := range paths {
		remotePath := fmt.Sprintf("%s/file-%d-%s", remoteDir, index+1, filepath.Base(path))
		if err := runGuestSCP(ctx, scpPath, guestRequest, path, guest+":"+remotePath); err != nil {
			return nil, err
		}
		remotePaths = append(remotePaths, remotePath)
	}
	return remotePaths, nil
}

func readGuestIntFile(ctx context.Context, sshPath string, sshBase []string, guest string, path string) int {
	text := strings.TrimSpace(readGuestFile(ctx, sshPath, sshBase, guest, path))
	if text == "" {
		return 0
	}
	var value int
	if _, err := fmt.Sscanf(text, "%d", &value); err != nil {
		return 0
	}
	return value
}

func guestXWindowMatchObserved(text string, match string) bool {
	trimmed := strings.TrimSpace(match)
	if trimmed == "" {
		return false
	}
	return strings.Contains(strings.ToLower(text), strings.ToLower(trimmed))
}

func summarizeGuestXWindowEvidence(text string, match string) string {
	lines := strings.Split(text, "\n")
	trimmedMatch := strings.ToLower(strings.TrimSpace(match))
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			continue
		}
		if trimmedMatch != "" && strings.Contains(strings.ToLower(trimmedLine), trimmedMatch) {
			return trimmedLine
		}
	}
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if strings.HasPrefix(trimmedLine, "0x") {
			return trimmedLine
		}
	}
	return ""
}

func runXWinInfo(ctx context.Context, xwininfoPath string, display string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	command := exec.CommandContext(ctx, xwininfoPath, "-root", "-tree")
	command.Env = append(os.Environ(), "DISPLAY="+display)
	command.Stdout = &stdout
	command.Stderr = &stderr
	err := command.Run()
	return stdout.String(), stderr.String(), err
}

func countXWindowChildren(text string) int {
	count := 0
	for _, line := range strings.Split(text, "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), "0x") {
			count++
		}
	}
	return count
}

func readGuestFile(ctx context.Context, sshPath string, sshBase []string, guest string, path string) string {
	var stdout bytes.Buffer
	_ = runGuestSSH(ctx, sshPath, sshBase, guest, "cat "+shellQuote(path)+" 2>/dev/null || true", &stdout, nil)
	return stdout.String()
}

func guiDriverErrorObserved(values ...string) bool {
	for _, value := range values {
		lower := strings.ToLower(value)
		if strings.Contains(lower, "graphics driver is missing") || strings.Contains(lower, "no driver could be loaded") {
			return true
		}
	}
	return false
}

func guestGUIKDESafeOutputSummary(result GuestGUIResult) string {
	if result.XWindowObserved {
		return fmt.Sprintf("Wine GUI window observed; child_windows=%d xwininfo_bytes=%d guest_stderr_bytes=%d", result.XWindowChildCount, result.XWinInfoBytes, result.GuestStderrBytes)
	}
	return fmt.Sprintf("Wine GUI window not observed; child_windows=%d xwininfo_bytes=%d wineboot_stderr_bytes=%d guest_stderr_bytes=%d", result.XWindowChildCount, result.XWinInfoBytes, result.WinebootStderrBytes, result.GuestStderrBytes)
}
