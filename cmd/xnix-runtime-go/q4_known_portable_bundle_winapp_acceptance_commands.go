package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runQ4KnownPortableBundleWinAppAcceptancePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.Q4KnownPortableBundleWinAppAcceptanceRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	smokeReportPath := flags.String("q4-notepadpp-portable-winapp-smoke", "", "passed q4 Notepad++ Portable Windows app smoke JSON")
	outputPath := flags.String("output", "", "optional JSON output path for the known portable bundle Windows app acceptance")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("q4-known-portable-bundle-winapp-acceptance-preview does not accept positional arguments")
	}
	acceptance, err := appidentity.PreviewQ4KnownPortableBundleWinAppAcceptance(appidentity.Q4KnownPortableBundleWinAppAcceptanceRequest{
		SmokeReportPath: *smokeReportPath,
	})
	if err != nil {
		return err
	}
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(acceptance); err != nil {
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
		return fmt.Errorf("prepare q4 known portable bundle Windows app acceptance output directory: %w", err)
	}
	if err := os.WriteFile(cleanOutput, buffer.Bytes(), 0o600); err != nil {
		return fmt.Errorf("write q4 known portable bundle Windows app acceptance output: %w", err)
	}
	return nil
}
