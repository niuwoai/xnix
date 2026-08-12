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

func runExternalWinAppApplicationDetailPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.ExternalWinAppApplicationDetailRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	bundlePath := flags.String("compatibility-evidence-bundle", "", "external Windows app compatibility evidence bundle JSON")
	acceptancePath := flags.String("q4-staged-external-winapp-acceptance", "", "optional q4 staged external Windows app acceptance JSON")
	outputPath := flags.String("output", "", "optional JSON output path for the external Windows app detail")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", appidentity.ExternalWinAppApplicationDetailRequestType)
	}

	detail, err := appidentity.PreviewExternalWinAppApplicationDetail(appidentity.ExternalWinAppApplicationDetailRequest{
		CompatibilityEvidenceBundlePath:      *bundlePath,
		Q4StagedExternalWinAppAcceptancePath: *acceptancePath,
	})
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(detail); err != nil {
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
		return fmt.Errorf("prepare external Windows app detail output directory: %w", err)
	}
	if err := os.WriteFile(cleanOutput, buffer.Bytes(), 0o600); err != nil {
		return fmt.Errorf("write external Windows app detail output: %w", err)
	}
	return nil
}
