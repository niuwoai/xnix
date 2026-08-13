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

func runQ4KnownPortableWinAppRunAcceptancePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.Q4KnownPortableWinAppRunAcceptanceRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	runReportPath := flags.String("q4-known-portable-winapp-run", "", "passed q4 known portable Windows app run JSON")
	outputPath := flags.String("output", "", "optional JSON output path for the known portable Windows app operator run acceptance")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("q4-known-portable-winapp-run-acceptance-preview does not accept positional arguments")
	}

	acceptance, err := appidentity.PreviewQ4KnownPortableWinAppRunAcceptance(appidentity.Q4KnownPortableWinAppRunAcceptanceRequest{
		RunReportPath: *runReportPath,
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
	if trimmed := *outputPath; trimmed != "" {
		if err := os.MkdirAll(filepath.Dir(trimmed), 0o755); err != nil {
			return fmt.Errorf("prepare q4 known portable Windows app run acceptance output directory: %w", err)
		}
		if err := os.WriteFile(trimmed, buffer.Bytes(), 0o644); err != nil {
			return fmt.Errorf("write q4 known portable Windows app run acceptance output: %w", err)
		}
	}
	return nil
}
