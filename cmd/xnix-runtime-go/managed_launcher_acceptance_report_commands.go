package main

import (
	"errors"
	"flag"
	"io"
	"strings"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runManagedLauncherAcceptanceReportPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.ManagedLauncherAcceptanceReportRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	stateRoot := flags.String("state-root", "", "Runtime state root containing the Runtime-status launch evidence handoff")
	evidenceID := flags.String("evidence-id", "", "opaque Runtime-status launch evidence handoff id")
	evidenceRelativePath := flags.String("evidence-relative-path", "", "relative Runtime-status launch evidence handoff path")
	expectedEvidenceSHA256 := flags.String("expected-evidence-sha256", "", "expected Runtime-status launch evidence digest")
	fullCheckpointPromoted := flags.Bool("full-checkpoint-promoted", false, "treat the formal full checkpoint as promoted")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("managed-launcher-acceptance-report-preview does not accept positional arguments")
	}
	if strings.TrimSpace(*stateRoot) == "" {
		return errors.New("managed-launcher-acceptance-report-preview requires --state-root")
	}
	preview, err := appidentity.PreviewManagedLauncherAcceptanceReport(appidentity.ManagedLauncherAcceptanceReportRequest{
		StateRoot:              *stateRoot,
		EvidenceID:             *evidenceID,
		EvidenceRelativePath:   *evidenceRelativePath,
		ExpectedEvidenceSHA256: *expectedEvidenceSHA256,
		FullCheckpointPromoted: *fullCheckpointPromoted,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}
