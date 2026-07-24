package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runGUISmokeEvidencePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.GUISmokeEvidencePreviewRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	reportPath := flags.String("gui-smoke-report", "", "executed Wine guest GUI smoke JSON report path")
	appID := flags.String("app-id", "", "application id to attach to the GUI smoke evidence")
	displayName := flags.String("display-name", "", "display name to attach to the GUI smoke evidence")
	appVersion := flags.String("app-version", "", "application version to attach to the GUI smoke evidence")
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

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}
