package main

import (
	"errors"
	"flag"
	"io"
	"os"
	"strings"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runKDEControlledLaunchActionSurfaceAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.KDEControlledLaunchActionSurfaceAuditRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	desktopEntryFile := flags.String("desktop-entry-file", "", "KDE controlled-launch desktop entry metadata file")
	evidenceHandle := flags.String("evidence-handle", "", "safe Runtime-status launch evidence handle")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("kde-controlled-launch-action-surface-audit-preview does not accept positional arguments")
	}
	if strings.TrimSpace(*desktopEntryFile) == "" {
		return errors.New("kde-controlled-launch-action-surface-audit-preview requires --desktop-entry-file")
	}
	contents, err := os.ReadFile(*desktopEntryFile)
	if err != nil {
		return errors.New("kde-controlled-launch-action-surface-audit-preview could not read desktop entry metadata")
	}
	preview, err := appidentity.PreviewKDEControlledLaunchActionSurfaceAudit(appidentity.KDEControlledLaunchActionSurfaceAuditRequest{
		DesktopEntryContent: string(contents),
		EvidenceHandle:      *evidenceHandle,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}
