package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"time"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/winapp"
)

func runWindowsCompatibilityWorkstreamsPreview(args []string, stdout io.Writer) error {
	if len(args) != 0 {
		return fmt.Errorf("%s does not accept arguments", "windows-compatibility-workstreams-preview")
	}
	preview, err := appidentity.NewWindowsCompatibilityWorkstreamsPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runWindowsAppRunSmoke(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("windows-app-run-smoke", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var exePath string
	var stateRoot string
	var workingDir string
	var runnerPath string
	var timeoutText string
	var expectedMarker string
	var successMode string
	var runnerBottle string
	var redactOutput bool
	flags.StringVar(&exePath, "exe", "", "Windows executable path")
	flags.StringVar(&stateRoot, "state-root", "", "isolated Runtime state root")
	flags.StringVar(&workingDir, "working-dir", "", "working directory for the compatibility runner; defaults to the executable directory")
	flags.StringVar(&runnerPath, "runner", "", "explicit compatibility runner path")
	flags.StringVar(&runnerBottle, "runner-bottle", "", "compatibility runner bottle name passed before the executable path")
	flags.StringVar(&timeoutText, "timeout", "20s", "execution timeout")
	flags.StringVar(&expectedMarker, "expected-marker", winapp.DefaultMarker, "expected stdout marker")
	flags.StringVar(&successMode, "success-mode", winapp.SuccessModeMarker, "success mode: marker, exit-code, or startup-window")
	flags.BoolVar(&redactOutput, "redact-output", false, "omit raw stdout and stderr while preserving safe output evidence")

	var appArgs repeatedStringFlag
	var runnerArgs repeatedStringFlag
	flags.Var(&appArgs, "arg", "argument passed to the Windows executable")
	flags.Var(&runnerArgs, "runner-arg", "argument passed to the compatibility runner before the executable path")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", "windows-app-run-smoke")
	}

	timeout, err := time.ParseDuration(timeoutText)
	if err != nil {
		return fmt.Errorf("parse timeout: %w", err)
	}

	result, err := winapp.RunSmoke(context.Background(), winapp.Request{
		ExecutablePath:   exePath,
		Arguments:        []string(appArgs),
		RunnerArguments:  []string(runnerArgs),
		RunnerBottle:     runnerBottle,
		StateRoot:        stateRoot,
		WorkingDirectory: workingDir,
		RunnerPath:       runnerPath,
		Timeout:          timeout,
		ExpectedMarker:   expectedMarker,
		SuccessMode:      successMode,
		RedactOutput:     redactOutput,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

func runWindowsAppRunnerDiagnostics(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("windows-app-runner-diagnostics", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var runnerPath string
	flags.StringVar(&runnerPath, "runner", "", "explicit compatibility runner path")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", "windows-app-runner-diagnostics")
	}

	result := winapp.RunnerDiagnostics(runnerPath)
	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

func runWindowsAppContainerRunSmoke(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("windows-app-container-run-smoke", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var exePath string
	var stateRoot string
	var image string
	var platform string
	var dockerPath string
	var timeoutText string
	var bootstrapTimeoutText string
	flags.StringVar(&exePath, "exe", "", "Windows executable path")
	flags.StringVar(&stateRoot, "state-root", "", "isolated Runtime state root")
	flags.StringVar(&image, "image", winapp.DefaultContainerImage, "local Wine container image")
	flags.StringVar(&platform, "platform", winapp.DefaultWinePlatform, "container platform")
	flags.StringVar(&dockerPath, "docker", "", "explicit docker runner path")
	flags.StringVar(&timeoutText, "timeout", "45s", "execution timeout")
	flags.StringVar(&bootstrapTimeoutText, "bootstrap-timeout", winapp.DefaultWineBootstrapTimeout.String(), "Wine prefix bootstrap timeout")

	var appArgs repeatedStringFlag
	flags.Var(&appArgs, "arg", "argument passed to the Windows executable")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", "windows-app-container-run-smoke")
	}

	timeout, err := time.ParseDuration(timeoutText)
	if err != nil {
		return fmt.Errorf("parse timeout: %w", err)
	}
	bootstrapTimeout, err := time.ParseDuration(bootstrapTimeoutText)
	if err != nil {
		return fmt.Errorf("parse bootstrap timeout: %w", err)
	}

	result, err := winapp.RunContainerSmoke(context.Background(), winapp.ContainerRequest{
		ExecutablePath:   exePath,
		Arguments:        []string(appArgs),
		StateRoot:        stateRoot,
		Image:            image,
		Platform:         platform,
		DockerPath:       dockerPath,
		Timeout:          timeout,
		BootstrapTimeout: bootstrapTimeout,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

func runWindowsAppGuestWineSmoke(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("windows-app-guest-wine-smoke", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var exePath string
	var host string
	var port string
	var user string
	var keyPath string
	var remoteDir string
	var sshPath string
	var scpPath string
	var timeoutText string
	var expectedMarker string
	flags.StringVar(&exePath, "exe", "", "Windows executable path")
	flags.StringVar(&host, "host", winapp.DefaultGuestHost, "guest SSH host")
	flags.StringVar(&port, "port", winapp.DefaultGuestPort, "guest SSH port")
	flags.StringVar(&user, "user", winapp.DefaultGuestUser, "guest SSH user")
	flags.StringVar(&keyPath, "key", "", "guest SSH private key path")
	flags.StringVar(&remoteDir, "remote-dir", winapp.DefaultRemoteDir, "guest remote smoke directory")
	flags.StringVar(&sshPath, "ssh", "", "explicit ssh client path")
	flags.StringVar(&scpPath, "scp", "", "explicit scp client path")
	flags.StringVar(&timeoutText, "timeout", "60s", "guest execution timeout")
	flags.StringVar(&expectedMarker, "expected-marker", winapp.DefaultMarker, "expected stdout marker")

	var appArgs repeatedStringFlag
	flags.Var(&appArgs, "arg", "argument passed to the Windows executable")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", "windows-app-guest-wine-smoke")
	}

	timeout, err := time.ParseDuration(timeoutText)
	if err != nil {
		return fmt.Errorf("parse timeout: %w", err)
	}
	parsedPort, err := winapp.ParseGuestPort(port)
	if err != nil {
		return err
	}

	result, err := winapp.RunGuestSmoke(context.Background(), winapp.GuestRequest{
		ExecutablePath: exePath,
		Arguments:      []string(appArgs),
		Host:           host,
		Port:           parsedPort,
		User:           user,
		KeyPath:        keyPath,
		RemoteDir:      remoteDir,
		SSHPath:        sshPath,
		SCPPath:        scpPath,
		Timeout:        timeout,
		ExpectedMarker: expectedMarker,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

func runWindowsKnownAppFetch(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("windows-known-app-fetch", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var appID string
	var cacheRoot string
	var timeoutText string
	flags.StringVar(&appID, "app", winapp.DefaultKnownAppID, "known Windows app id")
	flags.StringVar(&cacheRoot, "cache-root", winapp.DefaultKnownAppCacheRoot, "managed known Windows app cache root")
	flags.StringVar(&timeoutText, "timeout", winapp.DefaultKnownAppFetchTimeout.String(), "download timeout")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", "windows-known-app-fetch")
	}

	timeout, err := time.ParseDuration(timeoutText)
	if err != nil {
		return fmt.Errorf("parse timeout: %w", err)
	}

	result, err := winapp.FetchKnownPortableApp(context.Background(), winapp.KnownFetchRequest{
		AppID:         appID,
		CacheRoot:     cacheRoot,
		AllowDownload: true,
		Timeout:       timeout,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

func runWindowsKnownAppGuestWineSmoke(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("windows-known-app-guest-wine-smoke", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var appID string
	var cacheRoot string
	var host string
	var port string
	var user string
	var keyPath string
	var remoteDir string
	var sshPath string
	var scpPath string
	var timeoutText string
	flags.StringVar(&appID, "app", winapp.DefaultKnownAppID, "known Windows app id")
	flags.StringVar(&cacheRoot, "cache-root", winapp.DefaultKnownAppCacheRoot, "managed known Windows app cache root")
	flags.StringVar(&host, "host", winapp.DefaultGuestHost, "guest SSH host")
	flags.StringVar(&port, "port", winapp.DefaultGuestPort, "guest SSH port")
	flags.StringVar(&user, "user", winapp.DefaultGuestUser, "guest SSH user")
	flags.StringVar(&keyPath, "key", "", "guest SSH private key path")
	flags.StringVar(&remoteDir, "remote-dir", "/tmp/xnix-known-winapp-smoke", "guest remote smoke directory")
	flags.StringVar(&sshPath, "ssh", "", "explicit ssh client path")
	flags.StringVar(&scpPath, "scp", "", "explicit scp client path")
	flags.StringVar(&timeoutText, "timeout", winapp.DefaultKnownAppGuestTimeout.String(), "guest execution timeout")

	var appArgs repeatedStringFlag
	flags.Var(&appArgs, "arg", "argument passed to the known Windows app")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", "windows-known-app-guest-wine-smoke")
	}

	timeout, err := time.ParseDuration(timeoutText)
	if err != nil {
		return fmt.Errorf("parse timeout: %w", err)
	}
	parsedPort, err := winapp.ParseGuestPort(port)
	if err != nil {
		return err
	}

	result, err := winapp.RunKnownPortableGuestSmoke(context.Background(), winapp.KnownGuestRequest{
		AppID:     appID,
		CacheRoot: cacheRoot,
		Arguments: []string(appArgs),
		Host:      host,
		Port:      parsedPort,
		User:      user,
		KeyPath:   keyPath,
		RemoteDir: remoteDir,
		SSHPath:   sshPath,
		SCPPath:   scpPath,
		Timeout:   timeout,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

func runWindowsKnownAppManagedLaunchPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("windows-known-app-managed-launch-preview", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var appID string
	var cacheRoot string
	flags.StringVar(&appID, "app", winapp.DefaultKnownAppID, "known Windows app id")
	flags.StringVar(&cacheRoot, "cache-root", winapp.DefaultKnownAppCacheRoot, "managed known Windows app cache root")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", "windows-known-app-managed-launch-preview")
	}

	result, err := winapp.PreviewKnownPortableManagedLaunch(winapp.KnownManagedLaunchRequest{
		AppID:     appID,
		CacheRoot: cacheRoot,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

func runWindowsKnownAppKDELauncherPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("windows-known-app-kde-launcher-preview", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var appID string
	var cacheRoot string
	flags.StringVar(&appID, "app", winapp.DefaultKnownAppID, "known Windows app id")
	flags.StringVar(&cacheRoot, "cache-root", winapp.DefaultKnownAppCacheRoot, "managed known Windows app cache root")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", "windows-known-app-kde-launcher-preview")
	}

	result, err := winapp.PreviewKnownPortableKDELauncher(winapp.KnownKDELauncherRequest{
		AppID:     appID,
		CacheRoot: cacheRoot,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

func runWindowsKnownAppLaunchRequestPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("windows-known-app-launch-request-preview", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var appID string
	var cacheRoot string
	flags.StringVar(&appID, "app", winapp.DefaultKnownAppID, "known Windows app id")
	flags.StringVar(&cacheRoot, "cache-root", winapp.DefaultKnownAppCacheRoot, "managed known Windows app cache root")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", "windows-known-app-launch-request-preview")
	}

	result, err := winapp.PreviewKnownPortableLaunchRequest(winapp.KnownLaunchRequestRequest{
		AppID:     appID,
		CacheRoot: cacheRoot,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

func runWindowsKnownAppDispatchPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("windows-known-app-dispatch-preview", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var appID string
	var cacheRoot string
	flags.StringVar(&appID, "app", winapp.DefaultKnownAppID, "known Windows app id")
	flags.StringVar(&cacheRoot, "cache-root", winapp.DefaultKnownAppCacheRoot, "managed known Windows app cache root")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", "windows-known-app-dispatch-preview")
	}

	result, err := winapp.PreviewKnownPortableDispatch(winapp.KnownDispatchRequest{
		AppID:     appID,
		CacheRoot: cacheRoot,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

func runWindowsKnownAppDispatchSmoke(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("windows-known-app-dispatch-smoke", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var appID string
	var cacheRoot string
	var guestBoundary string
	var host string
	var port string
	var user string
	var keyPath string
	var remoteDir string
	var sshPath string
	var scpPath string
	var timeoutText string
	flags.StringVar(&appID, "app", winapp.DefaultKnownAppID, "known Windows app id")
	flags.StringVar(&cacheRoot, "cache-root", winapp.DefaultKnownAppCacheRoot, "managed known Windows app cache root")
	flags.StringVar(&guestBoundary, "guest-boundary", "", "controlled managed guest boundary supplied by the smoke harness")
	flags.StringVar(&host, "host", winapp.DefaultGuestHost, "guest SSH host")
	flags.StringVar(&port, "port", winapp.DefaultGuestPort, "guest SSH port")
	flags.StringVar(&user, "user", winapp.DefaultGuestUser, "guest SSH user")
	flags.StringVar(&keyPath, "key", "", "guest SSH private key path")
	flags.StringVar(&remoteDir, "remote-dir", "/tmp/xnix-known-winapp-smoke", "guest remote smoke directory")
	flags.StringVar(&sshPath, "ssh", "", "explicit ssh client path")
	flags.StringVar(&scpPath, "scp", "", "explicit scp client path")
	flags.StringVar(&timeoutText, "timeout", winapp.DefaultKnownAppGuestTimeout.String(), "guest execution timeout")

	var appArgs repeatedStringFlag
	flags.Var(&appArgs, "arg", "argument passed to the known Windows app")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", "windows-known-app-dispatch-smoke")
	}

	timeout, err := time.ParseDuration(timeoutText)
	if err != nil {
		return fmt.Errorf("parse timeout: %w", err)
	}
	parsedPort, err := winapp.ParseGuestPort(port)
	if err != nil {
		return err
	}

	result, err := winapp.RunKnownPortableDispatchSmoke(context.Background(), winapp.KnownDispatchSmokeRequest{
		AppID:         appID,
		CacheRoot:     cacheRoot,
		Arguments:     []string(appArgs),
		GuestBoundary: guestBoundary,
		Host:          host,
		Port:          parsedPort,
		User:          user,
		KeyPath:       keyPath,
		RemoteDir:     remoteDir,
		SSHPath:       sshPath,
		SCPPath:       scpPath,
		Timeout:       timeout,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

func runWindowsKnownAppLaunchBridgePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("windows-known-app-launch-bridge-preview", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var appID string
	var cacheRoot string
	var launcherArgv repeatedStringFlag
	flags.StringVar(&appID, "app", winapp.DefaultKnownAppID, "known Windows app id")
	flags.StringVar(&cacheRoot, "cache-root", winapp.DefaultKnownAppCacheRoot, "managed known Windows app cache root")
	flags.Var(&launcherArgv, "launcher-arg", "argument from the managed launcher invocation")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", "windows-known-app-launch-bridge-preview")
	}

	result, err := winapp.PreviewKnownPortableLaunchBridge(winapp.KnownLaunchBridgeRequest{
		AppID:               appID,
		CacheRoot:           cacheRoot,
		ManagedLauncherArgv: []string(launcherArgv),
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

type repeatedStringFlag []string

func (flag *repeatedStringFlag) String() string {
	return fmt.Sprint([]string(*flag))
}

func (flag *repeatedStringFlag) Set(value string) error {
	*flag = append(*flag, value)
	return nil
}
