package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
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
	var profilePath string
	var runnerPath string
	var timeoutText string
	var expectedMarker string
	var successMode string
	var runnerBottle string
	var redactOutput bool
	var skipBootstrap bool
	var stageAppDir bool
	flags.StringVar(&exePath, "exe", "", "Windows executable path")
	flags.StringVar(&profilePath, "profile", "", "Windows app smoke profile JSON path")
	flags.StringVar(&stateRoot, "state-root", "", "isolated Runtime state root")
	flags.StringVar(&workingDir, "working-dir", "", "working directory for the compatibility runner; defaults to the executable directory")
	flags.StringVar(&runnerPath, "runner", "", "explicit compatibility runner path")
	flags.StringVar(&runnerBottle, "runner-bottle", "", "compatibility runner bottle name passed before the executable path")
	flags.StringVar(&timeoutText, "timeout", "20s", "execution timeout")
	flags.StringVar(&expectedMarker, "expected-marker", winapp.DefaultMarker, "expected stdout marker")
	flags.StringVar(&successMode, "success-mode", winapp.SuccessModeMarker, "success mode: marker, exit-code, or startup-window")
	flags.BoolVar(&redactOutput, "redact-output", false, "omit raw stdout and stderr while preserving safe output evidence")
	flags.BoolVar(&skipBootstrap, "skip-bootstrap", false, "skip Wine prefix bootstrap before application execution")
	flags.BoolVar(&stageAppDir, "stage-app-dir", false, "stage the application directory into the isolated Runtime state root before execution")

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
	visitedFlags := map[string]bool{}
	flags.Visit(func(flag *flag.Flag) {
		visitedFlags[flag.Name] = true
	})

	request, err := winapp.LoadSmokeProfile(profilePath)
	if err != nil {
		return err
	}
	if visitedFlags["timeout"] || request.Timeout == 0 {
		timeout, err := time.ParseDuration(timeoutText)
		if err != nil {
			return fmt.Errorf("parse timeout: %w", err)
		}
		request.Timeout = timeout
	}
	applySmokeCLIOverrides(&request, visitedFlags, exePath, []string(appArgs), []string(runnerArgs), runnerBottle, stateRoot, workingDir, runnerPath, expectedMarker, successMode, redactOutput, skipBootstrap, stageAppDir)

	result, err := winapp.RunSmoke(context.Background(), request)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

func runWindowsAppLaunchProfile(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("windows-app-launch-profile", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var profilePath string
	flags.StringVar(&profilePath, "profile", "", "Windows app smoke profile JSON path")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", "windows-app-launch-profile")
	}

	result, err := winapp.LaunchProfile(context.Background(), winapp.LaunchProfileRequest{
		ProfilePath: profilePath,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

func runWindowsAppSmokeProfileRender(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("windows-app-smoke-profile-render", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var exePath string
	var stateRoot string
	var workingDir string
	var profilePath string
	var runnerPath string
	var timeoutText string
	var expectedMarker string
	var successMode string
	var runnerBottle string
	var redactOutput bool
	var skipBootstrap bool
	var stageAppDir bool
	flags.StringVar(&exePath, "exe", "", "Windows executable path")
	flags.StringVar(&profilePath, "profile", "", "base Windows app smoke profile JSON path")
	flags.StringVar(&stateRoot, "state-root", "", "isolated Runtime state root")
	flags.StringVar(&workingDir, "working-dir", "", "working directory for the compatibility runner")
	flags.StringVar(&runnerPath, "runner", "", "explicit compatibility runner path")
	flags.StringVar(&runnerBottle, "runner-bottle", "", "compatibility runner bottle name")
	flags.StringVar(&timeoutText, "timeout", "30s", "execution timeout")
	flags.StringVar(&expectedMarker, "expected-marker", winapp.DefaultMarker, "expected stdout marker")
	flags.StringVar(&successMode, "success-mode", winapp.SuccessModeMarker, "success mode: marker, exit-code, or startup-window")
	flags.BoolVar(&redactOutput, "redact-output", false, "request redacted Runtime smoke output")
	flags.BoolVar(&skipBootstrap, "skip-bootstrap", false, "skip Wine prefix bootstrap before application execution")
	flags.BoolVar(&stageAppDir, "stage-app-dir", false, "stage the application directory into the isolated Runtime state root before execution")

	var appArgs repeatedStringFlag
	var runnerArgs repeatedStringFlag
	flags.Var(&appArgs, "arg", "argument passed to the Windows executable")
	flags.Var(&runnerArgs, "runner-arg", "argument passed to the compatibility runner before the executable path")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", "windows-app-smoke-profile-render")
	}
	visitedFlags := map[string]bool{}
	flags.Visit(func(flag *flag.Flag) {
		visitedFlags[flag.Name] = true
	})

	request, err := winapp.LoadSmokeProfile(profilePath)
	if err != nil {
		return err
	}
	if visitedFlags["timeout"] || request.Timeout == 0 {
		timeout, err := time.ParseDuration(timeoutText)
		if err != nil {
			return fmt.Errorf("parse timeout: %w", err)
		}
		request.Timeout = timeout
	}
	applySmokeCLIOverrides(&request, visitedFlags, exePath, []string(appArgs), []string(runnerArgs), runnerBottle, stateRoot, workingDir, runnerPath, expectedMarker, successMode, redactOutput, skipBootstrap, stageAppDir)

	profile := winapp.SmokeProfileFromRequest(request)
	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(profile)
}

func runWindowsAppLauncherBundleRecord(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("windows-app-launcher-bundle-record", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var profilePath string
	var appID string
	var name string
	var runtimeBinary string
	var launcherMode string
	flags.StringVar(&profilePath, "profile", "", "Windows app smoke profile JSON path")
	flags.StringVar(&appID, "app-id", "", "desktop-safe application id")
	flags.StringVar(&name, "name", "", "desktop display name")
	flags.StringVar(&runtimeBinary, "runtime-bin", "", "runtime binary used by the managed launcher")
	flags.StringVar(&launcherMode, "launcher-mode", "", "launcher mode: launch, execute, or preflight")
	var runtimeArgs repeatedStringFlag
	flags.Var(&runtimeArgs, "runtime-arg", "argument passed to the runtime binary before windows-app-run-smoke")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", "windows-app-launcher-bundle-record")
	}

	record, err := winapp.RecordLauncherBundle(winapp.LauncherBundleRequest{
		ProfilePath:      profilePath,
		ApplicationID:    appID,
		DisplayName:      name,
		RuntimeBinary:    runtimeBinary,
		RuntimeArguments: []string(runtimeArgs),
		LauncherMode:     launcherMode,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(record)
}

func applySmokeCLIOverrides(request *winapp.Request, visitedFlags map[string]bool, exePath string, appArgs []string, runnerArgs []string, runnerBottle string, stateRoot string, workingDir string, runnerPath string, expectedMarker string, successMode string, redactOutput bool, skipBootstrap bool, stageAppDir bool) {
	if strings.TrimSpace(exePath) != "" {
		request.ExecutablePath = exePath
	}
	if len(appArgs) > 0 {
		request.Arguments = append(request.Arguments, appArgs...)
	}
	if len(runnerArgs) > 0 {
		request.RunnerArguments = append(request.RunnerArguments, runnerArgs...)
	}
	if strings.TrimSpace(runnerBottle) != "" {
		request.RunnerBottle = runnerBottle
	}
	if strings.TrimSpace(stateRoot) != "" {
		request.StateRoot = stateRoot
	}
	if strings.TrimSpace(workingDir) != "" {
		request.WorkingDirectory = workingDir
	}
	if strings.TrimSpace(runnerPath) != "" {
		request.RunnerPath = runnerPath
	}
	if visitedFlags["expected-marker"] || strings.TrimSpace(request.ExpectedMarker) == "" {
		request.ExpectedMarker = expectedMarker
	}
	if visitedFlags["success-mode"] || strings.TrimSpace(request.SuccessMode) == "" {
		request.SuccessMode = successMode
	}
	if redactOutput {
		request.RedactOutput = true
	}
	if skipBootstrap {
		request.SkipBootstrap = true
	}
	if stageAppDir {
		request.StageAppDir = true
	}
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

func runWindowsAppSmokeProfilePreflight(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("windows-app-smoke-profile-preflight", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var profilePath string
	flags.StringVar(&profilePath, "profile", "", "Windows app smoke profile JSON path")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", "windows-app-smoke-profile-preflight")
	}

	result, err := winapp.PreflightSmokeProfile(profilePath)
	if err != nil {
		return err
	}

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

func runWindowsAppContainerXGUISmoke(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("windows-app-container-x-gui-smoke", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var appName string
	var windowMatch string
	var executablePath string
	var importRecordPath string
	var appID string
	var displayName string
	var appVersion string
	var recipeApp string
	var registryPath string
	var image string
	var platform string
	var dockerPath string
	var timeoutText string
	var outputPath string
	importRecordConsumed := false
	importedArtifactDigestVerified := false
	importedArtifactSHA256 := ""
	flags.StringVar(&appName, "app", winapp.DefaultContainerGUIApp, "Windows GUI application name available in the Wine image")
	flags.StringVar(&windowMatch, "window-match", winapp.DefaultContainerWindowMatch, "case-insensitive X window match text")
	flags.StringVar(&executablePath, "executable", "", "optional local Windows GUI .exe to copy into the isolated container")
	flags.StringVar(&importRecordPath, "external-app-import-record", "", "optional Runtime import record for an external Windows GUI .exe")
	flags.StringVar(&appID, "app-id", "", "optional application id for ad-hoc container GUI evidence")
	flags.StringVar(&displayName, "display-name", "", "optional display name for ad-hoc container GUI evidence")
	flags.StringVar(&appVersion, "app-version", "", "optional application version for ad-hoc container GUI evidence")
	flags.StringVar(&recipeApp, "recipe-app", "", "registered application id with container GUI smoke hints")
	flags.StringVar(&registryPath, "registry", "", "recipe registry path used with --recipe-app")
	flags.StringVar(&image, "image", winapp.DefaultContainerImage, "local Wine X GUI container image")
	flags.StringVar(&platform, "platform", "", "container platform, or empty to use the local image platform")
	flags.StringVar(&dockerPath, "docker", "", "explicit docker runner path")
	flags.StringVar(&timeoutText, "timeout", "90s", "execution timeout")
	flags.StringVar(&outputPath, "output", "", "optional JSON output path")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", "windows-app-container-x-gui-smoke")
	}
	appFlagProvided := false
	flags.Visit(func(visited *flag.Flag) {
		if visited.Name == "app" {
			appFlagProvided = true
		}
	})
	if strings.TrimSpace(importRecordPath) != "" {
		if strings.TrimSpace(executablePath) != "" ||
			strings.TrimSpace(recipeApp) != "" ||
			strings.TrimSpace(appID) != "" ||
			strings.TrimSpace(displayName) != "" ||
			strings.TrimSpace(appVersion) != "" {
			return errors.New("windows-app-container-x-gui-smoke --external-app-import-record cannot be combined with --executable, --recipe-app, or ad-hoc identity flags")
		}
		record, importedExecutablePath, err := appidentity.ResolveExternalWinAppImportedArtifact(importRecordPath)
		if err != nil {
			return err
		}
		executablePath = importedExecutablePath
		appID = record.ApplicationID
		displayName = record.DisplayName
		appVersion = record.AppVersion
		importRecordConsumed = true
		importedArtifactDigestVerified = true
		importedArtifactSHA256 = record.ArtifactSHA256
		if !appFlagProvided {
			appName = "/" + record.ExecutableName
		}
	}
	if strings.TrimSpace(executablePath) != "" && !appFlagProvided {
		appName = "/" + filepath.Base(executablePath)
	}

	timeout, err := time.ParseDuration(timeoutText)
	if err != nil {
		return fmt.Errorf("parse timeout: %w", err)
	}
	request := winapp.ContainerXGUIRequest{
		ExecutablePath:                  strings.TrimSpace(executablePath),
		ApplicationName:                 appName,
		WindowMatch:                     windowMatch,
		ApplicationID:                   strings.TrimSpace(appID),
		DisplayName:                     strings.TrimSpace(displayName),
		AppVersion:                      strings.TrimSpace(appVersion),
		ExternalAppImportRecordConsumed: importRecordConsumed,
		ImportedArtifactDigestVerified:  importedArtifactDigestVerified,
		ImportedArtifactSHA256:          importedArtifactSHA256,
		Image:                           image,
		Platform:                        platform,
		DockerPath:                      dockerPath,
		Timeout:                         timeout,
	}
	for _, value := range []string{request.ApplicationID, request.DisplayName, request.AppVersion} {
		if strings.ContainsAny(value, "\r\n") {
			return errors.New("windows-app-container-x-gui-smoke identity fields must be single-line values")
		}
	}
	if strings.TrimSpace(recipeApp) != "" {
		if strings.TrimSpace(registryPath) == "" {
			return errors.New("windows-app-container-x-gui-smoke --recipe-app requires --registry")
		}
		recipe, _, err := appidentity.LoadRecipeFromRegistry(registryPath, "", recipeApp)
		if err != nil {
			return err
		}
		if err := recipe.Validate(); err != nil {
			return err
		}
		if recipe.ContainerGUISmoke == (appidentity.ContainerGUISmokeHints{}) {
			return fmt.Errorf("recipe %s does not define container_gui_smoke hints", recipe.ID)
		}
		request.ApplicationName = recipe.ContainerGUISmoke.App
		request.WindowMatch = recipe.ContainerGUISmoke.WindowMatch
		request.ApplicationID = recipe.ID
		request.DisplayName = recipe.Name
		request.AppVersion = recipe.Version
		request.RecipeBacked = true
	}

	result, err := winapp.RunContainerXGUISmoke(context.Background(), request)
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(result); err != nil {
		return err
	}
	if _, err := stdout.Write(buffer.Bytes()); err != nil {
		return err
	}
	if strings.TrimSpace(outputPath) == "" {
		return nil
	}
	cleanOutput := filepath.Clean(outputPath)
	if err := os.MkdirAll(filepath.Dir(cleanOutput), 0o700); err != nil {
		return fmt.Errorf("prepare container X GUI smoke output directory: %w", err)
	}
	if err := os.WriteFile(cleanOutput, buffer.Bytes(), 0o600); err != nil {
		return fmt.Errorf("write container X GUI smoke output: %w", err)
	}
	return nil
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

func runWindowsAppGuestWineGUISmoke(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("windows-app-guest-wine-gui-smoke", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var guiAppPath string
	var appID string
	var executablePath string
	var host string
	var port string
	var user string
	var keyPath string
	var remoteDir string
	var sshPath string
	var scpPath string
	var xwininfoPath string
	var guestDisplay string
	var hostDisplay string
	var windowMatch string
	var timeoutText string
	var waitText string
	var fileArguments repeatedStringFlag
	flags.StringVar(&appID, "app", "", "known Windows GUI app id")
	flags.StringVar(&guiAppPath, "gui-app", winapp.DefaultGuestGUIApp, "Windows GUI app path inside the Wine guest")
	flags.StringVar(&executablePath, "executable", "", "local Windows GUI .exe to copy into the guest before launch")
	flags.StringVar(&host, "host", winapp.DefaultGuestHost, "guest SSH host")
	flags.StringVar(&port, "port", winapp.DefaultGuestPort, "guest SSH port")
	flags.StringVar(&user, "user", winapp.DefaultGuestUser, "guest SSH user")
	flags.StringVar(&keyPath, "key", "", "guest SSH private key path")
	flags.StringVar(&remoteDir, "remote-dir", "/tmp/xnix-wine-guest-gui-smoke", "guest remote GUI smoke directory")
	flags.StringVar(&sshPath, "ssh", "", "explicit ssh client path")
	flags.StringVar(&scpPath, "scp", "", "explicit scp client path")
	flags.StringVar(&xwininfoPath, "xwininfo", "", "explicit xwininfo path")
	flags.StringVar(&guestDisplay, "guest-display", winapp.DefaultGuestGUIDisplay, "guest-visible X11 display")
	flags.StringVar(&hostDisplay, "host-display", winapp.DefaultHostGUIDisplay, "host X11 display inspected by xwininfo")
	flags.StringVar(&windowMatch, "window-match", "", "case-insensitive X window title/text required for GUI observation")
	flags.StringVar(&timeoutText, "timeout", "90s", "guest GUI execution timeout")
	flags.StringVar(&waitText, "wait", "10s", "GUI window observation wait")
	flags.Var(&fileArguments, "file-argument", "local file copied into the guest and passed to the Windows GUI app; may be repeated")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", "windows-app-guest-wine-gui-smoke")
	}
	guiAppExplicit := false
	flags.Visit(func(flag *flag.Flag) {
		if flag.Name == "gui-app" {
			guiAppExplicit = true
		}
	})
	var knownApp winapp.KnownPortableApp
	if strings.TrimSpace(appID) != "" {
		var lookupErr error
		knownApp, lookupErr = winapp.LookupKnownPortableApp(appID)
		if lookupErr != nil {
			return lookupErr
		}
		if !knownApp.GuestBuiltinGUI {
			return fmt.Errorf("known app %s is not a guest GUI app", knownApp.ID)
		}
		if !guiAppExplicit && strings.TrimSpace(knownApp.GuestGUIAppPath) != "" {
			guiAppPath = knownApp.GuestGUIAppPath
		}
		if strings.TrimSpace(knownApp.GuestGUIAppPath) == "" && !guiAppExplicit && strings.TrimSpace(executablePath) == "" {
			return fmt.Errorf("known GUI app %s requires --gui-app or --executable", knownApp.ID)
		}
	}
	timeout, err := time.ParseDuration(timeoutText)
	if err != nil {
		return fmt.Errorf("parse timeout: %w", err)
	}
	wait, err := time.ParseDuration(waitText)
	if err != nil {
		return fmt.Errorf("parse wait: %w", err)
	}
	parsedPort, err := winapp.ParseGuestPort(port)
	if err != nil {
		return err
	}

	result, err := winapp.RunGuestGUISmoke(context.Background(), winapp.GuestGUIRequest{
		ExecutablePath:    executablePath,
		GUIAppPath:        guiAppPath,
		KnownAppID:        knownApp.ID,
		KnownAppName:      knownApp.DisplayName,
		KnownAppVersion:   knownApp.Version,
		Host:              host,
		Port:              parsedPort,
		User:              user,
		KeyPath:           keyPath,
		RemoteDir:         remoteDir,
		SSHPath:           sshPath,
		SCPPath:           scpPath,
		XWinInfoPath:      xwininfoPath,
		GuestDisplay:      guestDisplay,
		HostDisplay:       hostDisplay,
		FileArgumentPaths: []string(fileArguments),
		WindowMatch:       windowMatch,
		Timeout:           timeout,
		Wait:              wait,
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

func runWindowsKnownAppLaunchProfileMaterialize(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("windows-known-app-launch-profile-materialize", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var appID string
	var cacheRoot string
	var stateRoot string
	var profileOutput string
	var applicationID string
	var name string
	var runtimeBinary string
	var runnerPath string
	var runnerBottle string
	var skipBootstrap bool
	flags.StringVar(&appID, "app", winapp.DefaultKnownAppID, "known Windows app id")
	flags.StringVar(&cacheRoot, "cache-root", winapp.DefaultKnownAppCacheRoot, "known app cache root")
	flags.StringVar(&stateRoot, "state-root", "", "isolated Runtime state root for the launch profile")
	flags.StringVar(&profileOutput, "profile-output", "", "profile output path; defaults under the state root")
	flags.StringVar(&applicationID, "app-id", "", "desktop-safe application id")
	flags.StringVar(&name, "name", "", "desktop display name")
	flags.StringVar(&runtimeBinary, "runtime-bin", "", "runtime binary used by the managed launcher")
	flags.StringVar(&runnerPath, "runner", "", "Windows compatibility runner path stored in the launch profile")
	flags.StringVar(&runnerBottle, "runner-bottle", "", "runner bottle name stored in the launch profile")
	flags.BoolVar(&skipBootstrap, "skip-bootstrap", false, "store a launch profile that skips runner bootstrap")
	var runtimeArgs repeatedStringFlag
	var runnerArgs repeatedStringFlag
	flags.Var(&runtimeArgs, "runtime-arg", "argument passed to the runtime binary before windows-app-launch-profile")
	flags.Var(&runnerArgs, "runner-arg", "argument passed to the compatibility runner when launching the app")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", "windows-known-app-launch-profile-materialize")
	}

	result, err := winapp.MaterializeKnownPortableLaunchProfile(winapp.KnownLaunchProfileMaterializeRequest{
		AppID:            appID,
		CacheRoot:        cacheRoot,
		StateRoot:        stateRoot,
		ProfileOutput:    profileOutput,
		ApplicationID:    applicationID,
		DisplayName:      name,
		RuntimeBinary:    runtimeBinary,
		RuntimeArguments: []string(runtimeArgs),
		RunnerPath:       runnerPath,
		RunnerBottle:     runnerBottle,
		RunnerArguments:  []string(runnerArgs),
		SkipBootstrap:    skipBootstrap,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

func runWindowsKnownAppPrepareLaunchProfile(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("windows-known-app-prepare-launch-profile", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var appID string
	var cacheRoot string
	var stateRoot string
	var profileOutput string
	var applicationID string
	var name string
	var runtimeBinary string
	var runnerPath string
	var runnerBottle string
	var timeoutText string
	var allowDownload bool
	var skipBootstrap bool
	flags.StringVar(&appID, "app", winapp.DefaultKnownAppID, "known Windows app id")
	flags.StringVar(&cacheRoot, "cache-root", winapp.DefaultKnownAppCacheRoot, "known app cache root")
	flags.StringVar(&stateRoot, "state-root", "", "isolated Runtime state root for the launch profile")
	flags.StringVar(&profileOutput, "profile-output", "", "profile output path; defaults under the state root")
	flags.StringVar(&applicationID, "app-id", "", "desktop-safe application id")
	flags.StringVar(&name, "name", "", "desktop display name")
	flags.StringVar(&runtimeBinary, "runtime-bin", "", "runtime binary used by the managed launcher")
	flags.StringVar(&runnerPath, "runner", "", "Windows compatibility runner path stored in the launch profile")
	flags.StringVar(&runnerBottle, "runner-bottle", "", "runner bottle name stored in the launch profile")
	flags.StringVar(&timeoutText, "timeout", winapp.DefaultKnownAppFetchTimeout.String(), "known app fetch timeout when downloads are allowed")
	flags.BoolVar(&allowDownload, "allow-download", false, "download the known app artifact when it is missing")
	flags.BoolVar(&skipBootstrap, "skip-bootstrap", false, "store a launch profile that skips runner bootstrap")
	var runtimeArgs repeatedStringFlag
	var runnerArgs repeatedStringFlag
	flags.Var(&runtimeArgs, "runtime-arg", "argument passed to the runtime binary before windows-app-launch-profile")
	flags.Var(&runnerArgs, "runner-arg", "argument passed to the compatibility runner when launching the app")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", "windows-known-app-prepare-launch-profile")
	}
	timeout, err := time.ParseDuration(timeoutText)
	if err != nil {
		return fmt.Errorf("parse timeout: %w", err)
	}

	result, err := winapp.PrepareKnownPortableLaunchProfile(context.Background(), winapp.KnownPrepareLaunchProfileRequest{
		AppID:            appID,
		CacheRoot:        cacheRoot,
		StateRoot:        stateRoot,
		ProfileOutput:    profileOutput,
		ApplicationID:    applicationID,
		DisplayName:      name,
		RuntimeBinary:    runtimeBinary,
		RuntimeArguments: []string(runtimeArgs),
		RunnerPath:       runnerPath,
		RunnerBottle:     runnerBottle,
		RunnerArguments:  []string(runnerArgs),
		SkipBootstrap:    skipBootstrap,
		AllowDownload:    allowDownload,
		Timeout:          timeout,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

func runWindowsKnownAppPrepareAndLaunchProfile(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("windows-known-app-prepare-and-launch-profile", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var appID string
	var cacheRoot string
	var stateRoot string
	var profileOutput string
	var applicationID string
	var name string
	var runtimeBinary string
	var runnerPath string
	var runnerBottle string
	var timeoutText string
	var allowDownload bool
	var skipBootstrap bool
	flags.StringVar(&appID, "app", winapp.DefaultKnownAppID, "known Windows app id")
	flags.StringVar(&cacheRoot, "cache-root", winapp.DefaultKnownAppCacheRoot, "known app cache root")
	flags.StringVar(&stateRoot, "state-root", "", "isolated Runtime state root for the launch profile")
	flags.StringVar(&profileOutput, "profile-output", "", "profile output path; defaults under the state root")
	flags.StringVar(&applicationID, "app-id", "", "desktop-safe application id")
	flags.StringVar(&name, "name", "", "desktop display name")
	flags.StringVar(&runtimeBinary, "runtime-bin", "", "runtime binary used by the managed launcher")
	flags.StringVar(&runnerPath, "runner", "", "Windows compatibility runner path stored in the launch profile")
	flags.StringVar(&runnerBottle, "runner-bottle", "", "runner bottle name stored in the launch profile")
	flags.StringVar(&timeoutText, "timeout", winapp.DefaultKnownAppFetchTimeout.String(), "known app fetch timeout when downloads are allowed")
	flags.BoolVar(&allowDownload, "allow-download", false, "download the known app artifact when it is missing")
	flags.BoolVar(&skipBootstrap, "skip-bootstrap", false, "store a launch profile that skips runner bootstrap")
	var runtimeArgs repeatedStringFlag
	var runnerArgs repeatedStringFlag
	flags.Var(&runtimeArgs, "runtime-arg", "argument passed to the runtime binary before windows-app-launch-profile")
	flags.Var(&runnerArgs, "runner-arg", "argument passed to the compatibility runner when launching the app")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", "windows-known-app-prepare-and-launch-profile")
	}
	timeout, err := time.ParseDuration(timeoutText)
	if err != nil {
		return fmt.Errorf("parse timeout: %w", err)
	}

	result, err := winapp.PrepareAndLaunchKnownPortableProfile(context.Background(), winapp.KnownPrepareAndLaunchProfileRequest{
		AppID:            appID,
		CacheRoot:        cacheRoot,
		StateRoot:        stateRoot,
		ProfileOutput:    profileOutput,
		ApplicationID:    applicationID,
		DisplayName:      name,
		RuntimeBinary:    runtimeBinary,
		RuntimeArguments: []string(runtimeArgs),
		RunnerPath:       runnerPath,
		RunnerBottle:     runnerBottle,
		RunnerArguments:  []string(runnerArgs),
		SkipBootstrap:    skipBootstrap,
		AllowDownload:    allowDownload,
		Timeout:          timeout,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

func runWindowsKnownAppRun(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("windows-known-app-run", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var appID string
	var backend string
	var cacheRoot string
	var stateRoot string
	var profileOutput string
	var applicationID string
	var name string
	var runtimeBinary string
	var runnerPath string
	var runnerBottle string
	var host string
	var port string
	var user string
	var keyPath string
	var remoteDir string
	var sshPath string
	var scpPath string
	var timeoutText string
	var qemuBinary string
	var qemuKernel string
	var qemuMemory string
	var qemuCPUs string
	var qemuCPU string
	var qemuBootTimeoutText string
	var qemuSerialLog string
	var reportOutput string
	var allowDownload bool
	var skipBootstrap bool
	var startQEMU bool
	var redactOutput bool
	flags.StringVar(&appID, "app", winapp.DefaultKnownAppID, "known Windows app id")
	flags.StringVar(&backend, "backend", winapp.KnownRunBackendLocal, "known app backend: local or guest-wine")
	flags.StringVar(&cacheRoot, "cache-root", winapp.DefaultKnownAppCacheRoot, "known app cache root")
	flags.StringVar(&stateRoot, "state-root", "", "isolated Runtime state root for local launch profiles")
	flags.StringVar(&profileOutput, "profile-output", "", "profile output path; defaults under the state root")
	flags.StringVar(&applicationID, "app-id", "", "desktop-safe application id")
	flags.StringVar(&name, "name", "", "desktop display name")
	flags.StringVar(&runtimeBinary, "runtime-bin", "", "runtime binary used by the managed launcher")
	flags.StringVar(&runnerPath, "runner", "", "Windows compatibility runner path stored in the local launch profile")
	flags.StringVar(&runnerBottle, "runner-bottle", "", "runner bottle name stored in the local launch profile")
	flags.StringVar(&host, "host", winapp.DefaultGuestHost, "guest SSH host for guest-wine backend")
	flags.StringVar(&port, "port", winapp.DefaultGuestPort, "guest SSH port for guest-wine backend; use auto with --start-qemu")
	flags.StringVar(&user, "user", winapp.DefaultGuestUser, "guest SSH user for guest-wine backend")
	flags.StringVar(&keyPath, "key", "", "guest SSH private key path for guest-wine backend")
	flags.StringVar(&remoteDir, "remote-dir", "/tmp/xnix-known-winapp-smoke", "guest remote smoke directory for guest-wine backend")
	flags.StringVar(&sshPath, "ssh", "", "explicit ssh client path for guest-wine backend")
	flags.StringVar(&scpPath, "scp", "", "explicit scp client path for guest-wine backend")
	flags.StringVar(&timeoutText, "timeout", winapp.DefaultKnownAppFetchTimeout.String(), "known app operation timeout")
	flags.StringVar(&qemuBinary, "qemu-binary", winapp.DefaultQEMUBinary, "QEMU binary used only when --start-qemu is set")
	flags.StringVar(&qemuKernel, "qemu-kernel", "", "guest kernel image used only when --start-qemu is set")
	flags.StringVar(&qemuMemory, "qemu-memory", winapp.DefaultQEMUMemory, "QEMU memory used only when --start-qemu is set")
	flags.StringVar(&qemuCPUs, "qemu-cpus", winapp.DefaultQEMUCPUCount, "QEMU CPU count used only when --start-qemu is set")
	flags.StringVar(&qemuCPU, "qemu-cpu", winapp.DefaultQEMUCPUModel, "QEMU CPU model used only when --start-qemu is set")
	flags.StringVar(&qemuBootTimeoutText, "qemu-boot-timeout", winapp.DefaultQEMUBootTimeout.String(), "QEMU guest SSH boot timeout used only when --start-qemu is set")
	flags.StringVar(&qemuSerialLog, "qemu-serial-log", "", "serial log path used only when --start-qemu is set")
	flags.StringVar(&reportOutput, "report-output", "", "optional JSON report output path for operator-controlled evidence capture")
	flags.BoolVar(&allowDownload, "allow-download", false, "download the known app artifact when it is missing")
	flags.BoolVar(&skipBootstrap, "skip-bootstrap", false, "store a local launch profile that skips runner bootstrap")
	flags.BoolVar(&startQEMU, "start-qemu", false, "start and stop a loopback-only QEMU guest for the guest-wine backend")
	flags.BoolVar(&redactOutput, "redact-output", false, "redact raw known app stdout and stderr while preserving output evidence")
	var runtimeArgs repeatedStringFlag
	var runnerArgs repeatedStringFlag
	var appArgs repeatedStringFlag
	flags.Var(&runtimeArgs, "runtime-arg", "argument passed to the runtime binary before windows-app-launch-profile")
	flags.Var(&runnerArgs, "runner-arg", "argument passed to the compatibility runner when launching the app locally")
	flags.Var(&appArgs, "arg", "argument passed to the known Windows app by guest-wine backend")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", "windows-known-app-run")
	}
	timeout, err := time.ParseDuration(timeoutText)
	if err != nil {
		return fmt.Errorf("parse timeout: %w", err)
	}
	qemuBootTimeout, err := time.ParseDuration(qemuBootTimeoutText)
	if err != nil {
		return fmt.Errorf("parse qemu boot timeout: %w", err)
	}
	parsedPort, err := winapp.ParseGuestPortForQEMU(port, startQEMU)
	if err != nil {
		return err
	}

	result, err := winapp.RunKnownPortableApp(context.Background(), winapp.KnownRunRequest{
		AppID:            appID,
		RuntimeVersion:   readRuntimeGoProjectVersion(),
		Backend:          backend,
		CacheRoot:        cacheRoot,
		StateRoot:        stateRoot,
		ProfileOutput:    profileOutput,
		ApplicationID:    applicationID,
		DisplayName:      name,
		RuntimeBinary:    runtimeBinary,
		RuntimeArguments: []string(runtimeArgs),
		RunnerPath:       runnerPath,
		RunnerBottle:     runnerBottle,
		RunnerArguments:  []string(runnerArgs),
		SkipBootstrap:    skipBootstrap,
		AllowDownload:    allowDownload,
		Arguments:        []string(appArgs),
		Host:             host,
		Port:             parsedPort,
		User:             user,
		KeyPath:          keyPath,
		RemoteDir:        remoteDir,
		SSHPath:          sshPath,
		SCPPath:          scpPath,
		Timeout:          timeout,
		RedactOutput:     redactOutput,
		StartQEMU:        startQEMU,
		QEMUBinary:       qemuBinary,
		QEMUKernelImage:  qemuKernel,
		QEMUMemory:       qemuMemory,
		QEMUCPUCount:     qemuCPUs,
		QEMUCPUModel:     qemuCPU,
		QEMUBootTimeout:  qemuBootTimeout,
		QEMUSerialLog:    qemuSerialLog,
	})
	if err != nil {
		return err
	}

	return writeJSONResponse(stdout, reportOutput, result)
}

func readRuntimeGoProjectVersion() string {
	workingDirectory, err := os.Getwd()
	if err != nil {
		return ""
	}
	for {
		versionPath := filepath.Join(workingDirectory, "VERSION")
		content, err := os.ReadFile(versionPath)
		if err == nil {
			return strings.TrimSpace(string(content))
		}
		parent := filepath.Dir(workingDirectory)
		if parent == workingDirectory {
			return ""
		}
		workingDirectory = parent
	}
}

func writeJSONResponse(stdout io.Writer, reportOutput string, value any) error {
	payload, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	payload = append(payload, '\n')

	if strings.TrimSpace(reportOutput) != "" {
		reportPath, err := filepath.Abs(reportOutput)
		if err != nil {
			return fmt.Errorf("resolve report output path: %w", err)
		}
		if err := os.MkdirAll(filepath.Dir(reportPath), 0o700); err != nil {
			return fmt.Errorf("create report output directory: %w", err)
		}
		if err := os.WriteFile(reportPath, payload, 0o600); err != nil {
			return fmt.Errorf("write report output: %w", err)
		}
	}

	_, err = stdout.Write(payload)
	return err
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
