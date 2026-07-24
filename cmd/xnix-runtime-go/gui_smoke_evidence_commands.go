package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runGUISmokeEvidencePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.GUISmokeEvidencePreviewRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	reportPath := flags.String("gui-smoke-report", "", "executed Wine guest GUI smoke JSON report path")
	appID := flags.String("app-id", "", "application id to attach to the GUI smoke evidence")
	displayName := flags.String("display-name", "", "display name to attach to the GUI smoke evidence")
	appVersion := flags.String("app-version", "", "application version to attach to the GUI smoke evidence")
	outputPath := flags.String("output", "", "optional JSON output path for the projected GUI smoke evidence")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", appidentity.GUISmokeEvidencePreviewRequestType)
	}

	preview, err := appidentity.PreviewGUISmokeEvidence(appidentity.GUISmokeEvidencePreviewRequest{
		ReportPath:  *reportPath,
		AppID:       *appID,
		DisplayName: *displayName,
		AppVersion:  *appVersion,
	})
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(preview); err != nil {
		return err
	}
	if _, err := stdout.Write(buffer.Bytes()); err != nil {
		return err
	}
	if *outputPath == "" {
		return nil
	}
	cleanOutput := filepath.Clean(*outputPath)
	if err := os.MkdirAll(filepath.Dir(cleanOutput), 0o700); err != nil {
		return fmt.Errorf("prepare GUI smoke evidence output directory: %w", err)
	}
	if err := os.WriteFile(cleanOutput, buffer.Bytes(), 0o600); err != nil {
		return fmt.Errorf("write GUI smoke evidence output: %w", err)
	}
	return nil
}
