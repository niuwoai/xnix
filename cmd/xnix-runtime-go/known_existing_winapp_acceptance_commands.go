package main

import (
	"errors"
	"flag"
	"io"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runKnownExistingWinAppAcceptancePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.KnownExistingWinAppAcceptanceRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	runReportPath := flags.String("known-winapp-run", "", "passed known Windows app run JSON")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("known-existing-winapp-acceptance-preview does not accept positional arguments")
	}
	acceptance, err := appidentity.PreviewKnownExistingWinAppAcceptance(appidentity.KnownExistingWinAppAcceptanceRequest{
		RunReportPath: *runReportPath,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, acceptance)
}
