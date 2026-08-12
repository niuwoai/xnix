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
	"strings"
	"time"

	"xnix.local/xnix/internal/runtime/activation"
	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/winapp"
	"xnix.local/xnix/internal/testversion"
)

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
