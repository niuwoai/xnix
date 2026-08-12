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
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"xnix.local/xnix/internal/runtime/activation"
	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/winapp"
	"xnix.local/xnix/internal/testversion"
)

const externalWinAppImportStageLaunchSchemaVersion = "xnix.runtime.external_winapp_import_stage_launch.v1"
const externalWinAppImportStageLaunchRequestType = "external-winapp-import-stage-and-launch"

type externalWinAppImportStageLaunchResult struct {
	SchemaVersion                    string                              `json:"schema_version"`
	RequestType                      string                              `json:"request_type"`
	Status                           string                              `json:"status"`
	ApplicationID                    string                              `json:"application_id"`
	DisplayName                      string                              `json:"display_name"`
	ExternalAppHandle                string                              `json:"external_app_handle"`
	ImportRecorded                   bool                                `json:"import_recorded"`
	DesktopActivationStaged          bool                                `json:"desktop_activation_staged"`
	StagedLauncherInvoked            bool                                `json:"staged_launcher_invoked"`
	StagedLauncherFromActivationRoot bool                                `json:"staged_launcher_from_activation_root"`
	ManagedLauncherExecutableStaged  bool                                `json:"managed_launcher_executable_staged"`
	DesktopExecUsesExternalAppHandle bool                                `json:"desktop_exec_uses_external_app_handle"`
	ExternalAppDesktopHandleReady    bool                                `json:"external_app_desktop_handle_ready"`
	DesktopLaunchPacketRequested     bool                                `json:"desktop_launch_packet_requested"`
	DesktopLaunchPacketWritten       bool                                `json:"desktop_launch_packet_written"`
	LauncherRequestType              string                              `json:"launcher_request_type"`
	LauncherStatus                   string                              `json:"launcher_status"`
	ExternalAppImportRecordConsumed  bool                                `json:"external_app_import_record_consumed"`
	ExternalAppHandleConsumed        bool                                `json:"external_app_handle_consumed"`
	ExternalDesktopArgumentCount     int                                 `json:"external_desktop_argument_count"`
	ExternalFileURIArgumentsAccepted bool                                `json:"external_file_uri_arguments_accepted"`
	ExternalFileBridgeReady          bool                                `json:"external_file_bridge_ready"`
	ImportedArtifactDigestVerified   bool                                `json:"imported_artifact_digest_verified"`
	RuntimeLaunchExecuted            bool                                `json:"runtime_launch_executed"`
	WindowObserved                   bool                                `json:"window_observed"`
	XWindowObserved                  bool                                `json:"x_window_observed"`
	RuntimeOwned                     bool                                `json:"runtime_owned"`
	GoRuntimeBacked                  bool                                `json:"go_runtime_backed"`
	KDEPolicyOwner                   bool                                `json:"kde_policy_owner"`
	LaunchEnabled                    bool                                `json:"launch_enabled"`
	BackendLaunchEnabled             bool                                `json:"backend_launch_enabled"`
	ExecutionStarted                 bool                                `json:"execution_started"`
	HostRootModified                 bool                                `json:"host_root_modified"`
	PrivilegedContainerRequired      bool                                `json:"privileged_container_required"`
	HostNetworkingRequired           bool                                `json:"host_networking_required"`
	DockerSocketMounted              bool                                `json:"docker_socket_mounted"`
	BroadHostMountRequired           bool                                `json:"broad_host_mount_required"`
	RawImportRecordPathExposed       bool                                `json:"raw_import_record_path_exposed"`
	RawStateRootPathExposed          bool                                `json:"raw_state_root_path_exposed"`
	RawExecutablePathExposed         bool                                `json:"raw_executable_path_exposed"`
	RawLauncherPathExposed           bool                                `json:"raw_launcher_path_exposed"`
	RawLauncherOutputExposed         bool                                `json:"raw_launcher_output_exposed"`
	StageResult                      activation.StageResult              `json:"stage_result"`
	LauncherResult                   appidentity.ExternalWinAppRunResult `json:"launcher_result"`
	DesktopSafeSummary               string                              `json:"desktop_safe_summary"`
}

func runExternalWinAppImportRecord(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.ExternalWinAppImportRecordRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	stateRoot := flags.String("state-root", "", "controlled Runtime state root for imported external Windows apps")
	executablePath := flags.String("executable", "", "local Windows .exe file to import")
	appID := flags.String("app-id", "", "application id for the imported external Windows app")
	displayName := flags.String("display-name", "", "display name for the imported external Windows app")
	appVersion := flags.String("app-version", "", "optional imported app version; defaults to the project version")
	version := flags.String("version", "", "optional record version; defaults to the project version")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", appidentity.ExternalWinAppImportRecordRequestType)
	}
	recordVersion := *version
	if recordVersion == "" {
		current, err := testversion.Current()
		if err != nil {
			return err
		}
		recordVersion = current
	}
	record, err := appidentity.RecordExternalWinAppImport(appidentity.ExternalWinAppImportRequest{
		Version:        recordVersion,
		StateRoot:      *stateRoot,
		ExecutablePath: *executablePath,
		AppID:          *appID,
		DisplayName:    *displayName,
		AppVersion:     *appVersion,
	})
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(record)
}

func runExternalWinAppBundleImportRecord(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.ExternalWinAppBundleImportRecordRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	stateRoot := flags.String("state-root", "", "controlled Runtime state root for imported external Windows app bundles")
	bundleRoot := flags.String("bundle-root", "", "local portable Windows app directory to import")
	executableRelativePath := flags.String("executable-relative-path", "", "portable bundle relative path to the Windows .exe")
	appID := flags.String("app-id", "", "application id for the imported external Windows app bundle")
	displayName := flags.String("display-name", "", "display name for the imported external Windows app bundle")
	appVersion := flags.String("app-version", "", "optional imported app version; defaults to the project version")
	version := flags.String("version", "", "optional record version; defaults to the project version")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", appidentity.ExternalWinAppBundleImportRecordRequestType)
	}
	recordVersion := *version
	if recordVersion == "" {
		current, err := testversion.Current()
		if err != nil {
			return err
		}
		recordVersion = current
	}
	record, err := appidentity.RecordExternalWinAppBundleImport(appidentity.ExternalWinAppBundleImportRequest{
		Version:                recordVersion,
		StateRoot:              *stateRoot,
		BundleRoot:             *bundleRoot,
		ExecutableRelativePath: *executableRelativePath,
		AppID:                  *appID,
		DisplayName:            *displayName,
		AppVersion:             *appVersion,
	})
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(record)
}

func runExternalWinAppRun(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.ExternalWinAppRunRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	importRecordPath := flags.String("external-app-import-record", "", "Runtime external Windows app import record")
	stateRoot := flags.String("state-root", "", "controlled Runtime state root for resolving an external Windows app handle")
	externalAppHandle := flags.String("external-app-handle", "", "opaque external Windows app handle; currently the imported reverse-DNS app id")
	windowMatch := flags.String("window-match", "", "case-insensitive X window match text; defaults to the imported executable name")
	image := flags.String("image", winapp.DefaultContainerImage, "local Wine X GUI container image")
	platform := flags.String("platform", "", "container platform, or empty to use the local image platform")
	dockerPath := flags.String("docker", "", "explicit docker runner path")
	timeoutText := flags.String("timeout", "90s", "execution timeout")
	outputPath := flags.String("output", "", "optional JSON output path")
	if err := flags.Parse(args); err != nil {
		return err
	}
	desktopArguments := flags.Args()
	timeout, err := time.ParseDuration(*timeoutText)
	if err != nil {
		return fmt.Errorf("parse timeout: %w", err)
	}
	result, err := appidentity.RunExternalWinApp(context.Background(), appidentity.ExternalWinAppRunRequest{
		ImportRecordPath:         *importRecordPath,
		StateRoot:                *stateRoot,
		ExternalAppHandle:        *externalAppHandle,
		ExternalDesktopArguments: desktopArguments,
		WindowMatch:              *windowMatch,
		Image:                    *image,
		Platform:                 *platform,
		DockerPath:               *dockerPath,
		Timeout:                  timeout,
	})
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
	if strings.TrimSpace(*outputPath) == "" {
		return nil
	}
	if err := os.WriteFile(*outputPath, buffer.Bytes(), 0o600); err != nil {
		return fmt.Errorf("write external Windows app run output: %w", err)
	}
	return nil
}

func runExternalWinAppImportAndRun(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("external-winapp-import-and-run", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	stateRoot := flags.String("state-root", "", "controlled Runtime state root for imported external Windows apps")
	executablePath := flags.String("executable", "", "local Windows .exe file to import and run")
	appID := flags.String("app-id", "", "application id for the imported external Windows app")
	displayName := flags.String("display-name", "", "display name for the imported external Windows app")
	appVersion := flags.String("app-version", "", "optional imported app version; defaults to the project version")
	recordVersion := flags.String("version", "", "optional record version; defaults to the project version")
	windowMatch := flags.String("window-match", "", "case-insensitive X window match text; defaults to the opened document name or imported executable name")
	image := flags.String("image", winapp.DefaultContainerImage, "local Wine X GUI container image")
	platform := flags.String("platform", "", "container platform, or empty to use the local image platform")
	dockerPath := flags.String("docker", "", "explicit docker runner path")
	timeoutText := flags.String("timeout", "90s", "execution timeout")
	outputPath := flags.String("output", "", "optional JSON output path")
	if err := flags.Parse(args); err != nil {
		return err
	}
	timeout, err := time.ParseDuration(*timeoutText)
	if err != nil {
		return fmt.Errorf("parse timeout: %w", err)
	}
	version := *recordVersion
	if version == "" {
		current, err := testversion.Current()
		if err != nil {
			return err
		}
		version = current
	}
	record, err := appidentity.RecordExternalWinAppImport(appidentity.ExternalWinAppImportRequest{
		Version:        version,
		StateRoot:      *stateRoot,
		ExecutablePath: *executablePath,
		AppID:          *appID,
		DisplayName:    *displayName,
		AppVersion:     *appVersion,
	})
	if err != nil {
		return err
	}
	runResult, err := appidentity.RunExternalWinApp(context.Background(), appidentity.ExternalWinAppRunRequest{
		StateRoot:                *stateRoot,
		ExternalAppHandle:        record.ApplicationID,
		ExternalDesktopArguments: flags.Args(),
		WindowMatch:              *windowMatch,
		Image:                    *image,
		Platform:                 *platform,
		DockerPath:               *dockerPath,
		Timeout:                  timeout,
	})
	if err != nil {
		return err
	}
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(runResult); err != nil {
		return err
	}
	if _, err := stdout.Write(buffer.Bytes()); err != nil {
		return err
	}
	if strings.TrimSpace(*outputPath) == "" {
		return nil
	}
	if err := os.WriteFile(*outputPath, buffer.Bytes(), 0o600); err != nil {
		return fmt.Errorf("write external Windows app import-and-run output: %w", err)
	}
	return nil
}

func runExternalWinAppImportAndStage(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("external-winapp-import-and-stage", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	stateRoot := flags.String("state-root", "", "controlled Runtime state root for imported external Windows apps")
	executablePath := flags.String("executable", "", "local Windows .exe file to import and stage for desktop activation")
	appID := flags.String("app-id", "", "application id for the imported external Windows app")
	displayName := flags.String("display-name", "", "display name for the imported external Windows app")
	appVersion := flags.String("app-version", "", "optional imported app version; defaults to the project version")
	recordVersion := flags.String("version", "", "optional record version; defaults to the project version")
	mode := flags.String("mode", "development", "activation staging mode: production or development")
	stagingRoot := flags.String("staging-root", "", "test root where desktop activation files may be staged")
	managedLauncherBinary := flags.String("managed-launcher-bin", "", "optional path to a prebuilt xnix-compat-launch binary to copy into the staging root")
	outputPath := flags.String("output", "", "optional JSON output path")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("external-winapp-import-and-stage does not accept positional arguments")
	}
	version := *recordVersion
	if version == "" {
		current, err := testversion.Current()
		if err != nil {
			return err
		}
		version = current
	}
	record, err := appidentity.RecordExternalWinAppImport(appidentity.ExternalWinAppImportRequest{
		Version:        version,
		StateRoot:      *stateRoot,
		ExecutablePath: *executablePath,
		AppID:          *appID,
		DisplayName:    *displayName,
		AppVersion:     *appVersion,
	})
	if err != nil {
		return err
	}
	recipe, provenance, err := appidentity.ExternalAppRecipeFromImportRecord(record)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	result, err := activation.Stage(activation.StageRequest{
		Root:                  *stagingRoot,
		Mode:                  *mode,
		Plan:                  plan,
		ManagedLauncherBinary: *managedLauncherBinary,
	})
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
	if strings.TrimSpace(*outputPath) == "" {
		return nil
	}
	if err := os.WriteFile(*outputPath, buffer.Bytes(), 0o600); err != nil {
		return fmt.Errorf("write external Windows app import-and-stage output: %w", err)
	}
	return nil
}

func runExternalWinAppImportStageAndLaunch(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(externalWinAppImportStageLaunchRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	stateRoot := flags.String("state-root", "", "controlled Runtime state root for imported external Windows apps")
	executablePath := flags.String("executable", "", "local Windows .exe file to import, stage, and launch")
	bundleRoot := flags.String("bundle-root", "", "local portable Windows app directory to import, stage, and launch")
	executableRelativePath := flags.String("executable-relative-path", "", "portable bundle relative path to the Windows .exe")
	appID := flags.String("app-id", "", "application id for the imported external Windows app")
	displayName := flags.String("display-name", "", "display name for the imported external Windows app")
	appVersion := flags.String("app-version", "", "optional imported app version; defaults to the project version")
	recordVersion := flags.String("version", "", "optional record version; defaults to the project version")
	mode := flags.String("mode", "development", "activation staging mode: production or development")
	stagingRoot := flags.String("staging-root", "", "test root where desktop activation files may be staged")
	managedLauncherBinary := flags.String("managed-launcher-bin", "", "optional path to a prebuilt xnix-compat-launch binary to copy into the staging root")
	launcherBinary := flags.String("launcher-bin", "", "optional xnix-compat-launch binary to invoke; defaults to the staged launcher when available")
	desktopLaunchPacketOutput := flags.String("desktop-launch-packet-output", "", "optional KDE-safe external app desktop launch packet sidecar output path")
	windowMatch := flags.String("window-match", "", "case-insensitive X window match text forwarded to the managed launcher")
	image := flags.String("image", winapp.DefaultContainerImage, "local Wine X GUI container image")
	platform := flags.String("platform", "", "container platform, or empty to use the local image platform")
	dockerPath := flags.String("docker", "", "explicit docker runner path")
	timeoutText := flags.String("timeout", "90s", "execution timeout")
	outputPath := flags.String("output", "", "optional JSON output path")
	if err := flags.Parse(args); err != nil {
		return err
	}
	version := *recordVersion
	if version == "" {
		current, err := testversion.Current()
		if err != nil {
			return err
		}
		version = current
	}
	timeout, err := time.ParseDuration(*timeoutText)
	if err != nil {
		return fmt.Errorf("parse timeout: %w", err)
	}
	executableSupplied := strings.TrimSpace(*executablePath) != ""
	bundleSupplied := strings.TrimSpace(*bundleRoot) != "" || strings.TrimSpace(*executableRelativePath) != ""
	if executableSupplied && bundleSupplied {
		return errors.New("external-winapp-import-stage-and-launch accepts either --executable or --bundle-root with --executable-relative-path, not both")
	}
	if !executableSupplied && !bundleSupplied {
		return errors.New("external-winapp-import-stage-and-launch requires --executable or --bundle-root with --executable-relative-path")
	}
	var record appidentity.ExternalWinAppImportRecord
	if bundleSupplied {
		record, err = appidentity.RecordExternalWinAppBundleImport(appidentity.ExternalWinAppBundleImportRequest{
			Version:                version,
			StateRoot:              *stateRoot,
			BundleRoot:             *bundleRoot,
			ExecutableRelativePath: *executableRelativePath,
			AppID:                  *appID,
			DisplayName:            *displayName,
			AppVersion:             *appVersion,
		})
	} else {
		record, err = appidentity.RecordExternalWinAppImport(appidentity.ExternalWinAppImportRequest{
			Version:        version,
			StateRoot:      *stateRoot,
			ExecutablePath: *executablePath,
			AppID:          *appID,
			DisplayName:    *displayName,
			AppVersion:     *appVersion,
		})
	}
	if err != nil {
		return err
	}
	recipe, provenance, err := appidentity.ExternalAppRecipeFromImportRecord(record)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	stageResult, err := activation.Stage(activation.StageRequest{
		Root:                  *stagingRoot,
		Mode:                  *mode,
		Plan:                  plan,
		ManagedLauncherBinary: *managedLauncherBinary,
	})
	if err != nil {
		return err
	}
	launcherPath, launcherFromStage := stagedLauncherPath(*launcherBinary, *managedLauncherBinary, *stagingRoot)
	launcherResult, err := invokeExternalWinAppStagedLauncher(invokeExternalWinAppStagedLauncherRequest{
		LauncherPath:              launcherPath,
		StateRoot:                 *stateRoot,
		ExternalAppHandle:         record.ApplicationID,
		DesktopArguments:          flags.Args(),
		ActivationRoot:            *stagingRoot,
		DesktopLaunchPacketOutput: *desktopLaunchPacketOutput,
		DesktopLaunchPacketMode:   *mode,
		WindowMatch:               *windowMatch,
		Image:                     *image,
		Platform:                  *platform,
		DockerPath:                *dockerPath,
		Timeout:                   timeout,
	})
	if err != nil {
		return err
	}
	desktopLaunchPacketWritten := false
	if strings.TrimSpace(*desktopLaunchPacketOutput) != "" {
		if info, statErr := os.Stat(*desktopLaunchPacketOutput); statErr == nil && !info.IsDir() {
			desktopLaunchPacketWritten = true
		}
	}
	result := externalWinAppImportStageLaunchResult{
		SchemaVersion:                    externalWinAppImportStageLaunchSchemaVersion,
		RequestType:                      externalWinAppImportStageLaunchRequestType,
		Status:                           launcherResult.Status,
		ApplicationID:                    record.ApplicationID,
		DisplayName:                      record.DisplayName,
		ExternalAppHandle:                record.ApplicationID,
		ImportRecorded:                   record.ImportRecorded,
		DesktopActivationStaged:          stageResult.FileWritesPerformed,
		StagedLauncherInvoked:            true,
		StagedLauncherFromActivationRoot: launcherFromStage,
		ManagedLauncherExecutableStaged:  containsStageFileID(stageResult.WrittenFileIDs, "managed-launcher-executable"),
		DesktopExecUsesExternalAppHandle: stageResult.DesktopExecUsesExternalAppHandle,
		ExternalAppDesktopHandleReady:    stageResult.ExternalAppDesktopHandleReady,
		DesktopLaunchPacketRequested:     strings.TrimSpace(*desktopLaunchPacketOutput) != "",
		DesktopLaunchPacketWritten:       desktopLaunchPacketWritten,
		LauncherRequestType:              launcherResult.RequestType,
		LauncherStatus:                   launcherResult.Status,
		ExternalAppImportRecordConsumed:  launcherResult.ExternalAppImportRecordConsumed,
		ExternalAppHandleConsumed:        launcherResult.ExternalAppHandleConsumed,
		ExternalDesktopArgumentCount:     launcherResult.ExternalDesktopArgumentCount,
		ExternalFileURIArgumentsAccepted: launcherResult.ExternalFileURIArgumentsAccepted,
		ExternalFileBridgeReady:          launcherResult.ExternalFileBridgeReady,
		ImportedArtifactDigestVerified:   launcherResult.ImportedArtifactDigestVerified,
		RuntimeLaunchExecuted:            launcherResult.RuntimeRunExecuted,
		WindowObserved:                   launcherResult.WindowObserved,
		XWindowObserved:                  launcherResult.XWindowObserved,
		RuntimeOwned:                     true,
		GoRuntimeBacked:                  true,
		KDEPolicyOwner:                   false,
		LaunchEnabled:                    stageResult.LaunchEnabled,
		BackendLaunchEnabled:             stageResult.BackendLaunchEnabled,
		ExecutionStarted:                 launcherResult.ExecutionStarted,
		HostRootModified:                 stageResult.HostRootModified || launcherResult.HostRootModified,
		PrivilegedContainerRequired:      launcherResult.PrivilegedContainerRequired,
		HostNetworkingRequired:           launcherResult.HostNetworkingRequired,
		DockerSocketMounted:              launcherResult.DockerSocketMounted,
		BroadHostMountRequired:           launcherResult.BroadHostMountRequired,
		RawImportRecordPathExposed:       launcherResult.RawImportRecordPathExposed,
		RawStateRootPathExposed:          launcherResult.RawStateRootPathExposed,
		RawExecutablePathExposed:         launcherResult.RawExecutablePathExposed,
		RawLauncherPathExposed:           false,
		RawLauncherOutputExposed:         false,
		StageResult:                      stageResult,
		LauncherResult:                   launcherResult,
		DesktopSafeSummary:               record.DisplayName + " was imported, staged as a KDE desktop activation entry, and launched through the Runtime-managed desktop launcher without exposing local paths.",
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
	if strings.TrimSpace(*outputPath) == "" {
		return nil
	}
	if err := os.WriteFile(*outputPath, buffer.Bytes(), 0o600); err != nil {
		return fmt.Errorf("write external Windows app import-stage-and-launch output: %w", err)
	}
	return nil
}

type invokeExternalWinAppStagedLauncherRequest struct {
	LauncherPath              string
	StateRoot                 string
	ExternalAppHandle         string
	DesktopArguments          []string
	ActivationRoot            string
	DesktopLaunchPacketOutput string
	DesktopLaunchPacketMode   string
	WindowMatch               string
	Image                     string
	Platform                  string
	DockerPath                string
	Timeout                   time.Duration
}

func stagedLauncherPath(explicitLauncherPath string, managedLauncherBinary string, stagingRoot string) (string, bool) {
	if strings.TrimSpace(explicitLauncherPath) != "" {
		return explicitLauncherPath, false
	}
	if strings.TrimSpace(managedLauncherBinary) != "" && strings.TrimSpace(stagingRoot) != "" {
		return filepath.Join(stagingRoot, "usr/local/bin/xnix-compat-launch"), true
	}
	return "xnix-compat-launch", false
}

func invokeExternalWinAppStagedLauncher(request invokeExternalWinAppStagedLauncherRequest) (appidentity.ExternalWinAppRunResult, error) {
	args := []string{
		"--external-app-handle", request.ExternalAppHandle,
		"--timeout", request.Timeout.String(),
	}
	if strings.TrimSpace(request.WindowMatch) != "" {
		args = append(args, "--external-window-match", request.WindowMatch)
	}
	args = append(args, request.DesktopArguments...)
	ctx, cancel := context.WithTimeout(context.Background(), request.Timeout+5*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, request.LauncherPath, args...)
	cmd.Env = append(os.Environ(),
		appidentity.ExternalWinAppStateRootEnv+"="+request.StateRoot,
		"XNIX_WINE_IMAGE="+request.Image,
		"XNIX_CONTAINER_PLATFORM="+request.Platform,
		"XNIX_DOCKER_BIN="+request.DockerPath,
		"XNIX_COMPAT_LAUNCH_TIMEOUT="+request.Timeout.String(),
	)
	if strings.TrimSpace(request.DesktopLaunchPacketOutput) != "" {
		cmd.Env = append(cmd.Env,
			"XNIX_EXTERNAL_APP_DESKTOP_ACTIVATION_ROOT="+request.ActivationRoot,
			"XNIX_EXTERNAL_APP_DESKTOP_LAUNCH_PACKET_OUTPUT="+request.DesktopLaunchPacketOutput,
			"XNIX_EXTERNAL_APP_DESKTOP_LAUNCH_PACKET_MODE="+request.DesktopLaunchPacketMode,
		)
	}
	output, err := cmd.Output()
	if err != nil {
		return appidentity.ExternalWinAppRunResult{}, fmt.Errorf("external Windows app staged launcher failed: %w", err)
	}
	var result appidentity.ExternalWinAppRunResult
	if err := json.Unmarshal(output, &result); err != nil {
		return appidentity.ExternalWinAppRunResult{}, fmt.Errorf("parse external Windows app staged launcher JSON: %w", err)
	}
	if result.SchemaVersion != appidentity.ExternalWinAppRunSchemaVersion ||
		result.RequestType != appidentity.ExternalWinAppRunRequestType {
		return appidentity.ExternalWinAppRunResult{}, errors.New("external Windows app staged launcher returned unsupported JSON")
	}
	return result, nil
}

func containsStageFileID(values []string, expected string) bool {
	for _, value := range values {
		if value == expected {
			return true
		}
	}
	return false
}
