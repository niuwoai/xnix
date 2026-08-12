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

func runExternalWinAppCompatibilityEvidenceBundlePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.ExternalWinAppCompatibilityEvidenceBundleRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	oneShotResultPath := flags.String("one-shot-result", "", "JSON result produced by external-winapp-import-stage-and-launch")
	desktopLaunchPacketPath := flags.String("desktop-launch-packet", "", "desktop external Windows app launch packet JSON artifact")
	runtimeGUIEvidencePacketPath := flags.String("runtime-gui-evidence-packet", "", "real Windows app GUI evidence packet JSON artifact")
	kdePagePath := flags.String("kde-page", "", "KDE external app page JSON artifact")
	outputPath := flags.String("output", "", "optional JSON output path for the compatibility evidence bundle")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", appidentity.ExternalWinAppCompatibilityEvidenceBundleRequestType)
	}

	bundle, err := appidentity.PreviewExternalWinAppCompatibilityEvidenceBundle(appidentity.ExternalWinAppCompatibilityEvidenceBundleRequest{
		OneShotResultPath:            *oneShotResultPath,
		DesktopLaunchPacketPath:      *desktopLaunchPacketPath,
		RuntimeGUIEvidencePacketPath: *runtimeGUIEvidencePacketPath,
		KDEPagePath:                  *kdePagePath,
	})
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(bundle); err != nil {
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
		return fmt.Errorf("prepare external Windows app compatibility evidence bundle output directory: %w", err)
	}
	if err := os.WriteFile(cleanOutput, buffer.Bytes(), 0o600); err != nil {
		return fmt.Errorf("write external Windows app compatibility evidence bundle output: %w", err)
	}
	return nil
}
