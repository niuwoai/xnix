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

func runRealWinAppGUIEvidencePacketPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.RealWinAppGUIEvidencePacketRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	reportPath := flags.String("gui-smoke-report", "", "executed Windows GUI smoke JSON report path")
	appID := flags.String("app-id", "", "application id to attach to the real GUI evidence packet")
	displayName := flags.String("display-name", "", "display name to attach to the real GUI evidence packet")
	appVersion := flags.String("app-version", "", "application version to attach to the real GUI evidence packet")
	outputPath := flags.String("output", "", "optional JSON output path for the real GUI evidence packet")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", appidentity.RealWinAppGUIEvidencePacketRequestType)
	}

	packet, err := appidentity.PreviewRealWinAppGUIEvidencePacket(appidentity.RealWinAppGUIEvidencePacketRequest{
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
	if err := encoder.Encode(packet); err != nil {
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
		return fmt.Errorf("prepare real GUI evidence packet output directory: %w", err)
	}
	if err := os.WriteFile(cleanOutput, buffer.Bytes(), 0o600); err != nil {
		return fmt.Errorf("write real GUI evidence packet output: %w", err)
	}
	return nil
}
