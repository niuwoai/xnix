package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

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
	windowMatch := flags.String("window-match", "", "case-insensitive X window match text; defaults to the imported executable name")
	image := flags.String("image", winapp.DefaultContainerImage, "local Wine X GUI container image")
	platform := flags.String("platform", "", "container platform, or empty to use the local image platform")
	dockerPath := flags.String("docker", "", "explicit docker runner path")
	timeoutText := flags.String("timeout", "90s", "execution timeout")
	outputPath := flags.String("output", "", "optional JSON output path")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", appidentity.ExternalWinAppRunRequestType)
	}
	timeout, err := time.ParseDuration(*timeoutText)
	if err != nil {
		return fmt.Errorf("parse timeout: %w", err)
	}
	result, err := appidentity.RunExternalWinApp(context.Background(), appidentity.ExternalWinAppRunRequest{
		ImportRecordPath: *importRecordPath,
		WindowMatch:      *windowMatch,
		Image:            *image,
		Platform:         *platform,
		DockerPath:       *dockerPath,
		Timeout:          timeout,
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
