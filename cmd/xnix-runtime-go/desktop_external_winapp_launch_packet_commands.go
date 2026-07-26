package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runDesktopExternalWinAppLaunchPacketPreview(args []string, stdout io.Writer) error {
	plan, activationRoot, runRecordPath, mode, outputPath, err := parseDesktopExternalWinAppLaunchPacketPreviewSource(args)
	if err != nil {
		return err
	}
	packet, err := plan.DesktopExternalWinAppLaunchPacketPreview(activationRoot, runRecordPath, mode)
	if err != nil {
		return err
	}
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(packet); err != nil {
		return err
	}
	if _, err := stdout.Write(buffer.Bytes()); err != nil {
		return err
	}
	if strings.TrimSpace(outputPath) == "" {
		return nil
	}
	if err := os.WriteFile(outputPath, buffer.Bytes(), 0o600); err != nil {
		return fmt.Errorf("write desktop external Windows app launch packet: %w", err)
	}
	return nil
}

func parseDesktopExternalWinAppLaunchPacketPreviewSource(args []string) (appidentity.Plan, string, string, string, string, error) {
	flags := flag.NewFlagSet(appidentity.DesktopExternalWinAppLaunchPacketRequestType, flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	externalAppImportRecord := flags.String("external-app-import-record", "", "Runtime external Windows app import record")
	activationRoot := flags.String("activation-root", "", "root containing staged desktop activation receipt evidence")
	runRecordPath := flags.String("run-record", "", "JSON record produced by windows-external-app-run or xnix-compat-launch")
	mode := flags.String("mode", "development", "desktop launch packet mode: production or development")
	outputPath := flags.String("output", "", "optional JSON output path")
	if err := flags.Parse(args); err != nil {
		return appidentity.Plan{}, "", "", "", "", err
	}
	if strings.TrimSpace(*externalAppImportRecord) == "" {
		return appidentity.Plan{}, "", "", "", "", errors.New("desktop-external-winapp-launch-packet-preview requires --external-app-import-record")
	}
	if strings.TrimSpace(*activationRoot) == "" {
		return appidentity.Plan{}, "", "", "", "", errors.New("desktop-external-winapp-launch-packet-preview requires --activation-root")
	}
	if strings.TrimSpace(*runRecordPath) == "" {
		return appidentity.Plan{}, "", "", "", "", errors.New("desktop-external-winapp-launch-packet-preview requires --run-record")
	}
	if flags.NArg() != 0 {
		return appidentity.Plan{}, "", "", "", "", errors.New("desktop-external-winapp-launch-packet-preview does not accept positional arguments")
	}
	record, err := appidentity.LoadExternalWinAppImportRecord(*externalAppImportRecord)
	if err != nil {
		return appidentity.Plan{}, "", "", "", "", err
	}
	recipe, provenance, err := appidentity.ExternalAppRecipeFromImportRecord(record)
	if err != nil {
		return appidentity.Plan{}, "", "", "", "", err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return appidentity.Plan{}, "", "", "", "", err
	}
	return plan, *activationRoot, *runRecordPath, *mode, *outputPath, nil
}
