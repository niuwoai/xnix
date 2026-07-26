package main

import (
	"errors"
	"flag"
	"io"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runQ4SampleNotepadAcceptancePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.Q4SampleNotepadAcceptanceRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	smokeReportPath := flags.String("q4-sample-notepad-smoke", "", "passed q4 Sample Notepad smoke JSON")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("q4-sample-notepad-acceptance-preview does not accept positional arguments")
	}
	acceptance, err := appidentity.PreviewQ4SampleNotepadAcceptance(appidentity.Q4SampleNotepadAcceptanceRequest{
		SmokeReportPath: *smokeReportPath,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, acceptance)
}
