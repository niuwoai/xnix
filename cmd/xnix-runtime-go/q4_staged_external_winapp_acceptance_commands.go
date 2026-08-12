package main

import (
	"errors"
	"flag"
	"io"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runQ4StagedExternalWinAppAcceptancePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.Q4StagedExternalWinAppAcceptanceRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	smokeReportPath := flags.String("q4-staged-external-winapp-smoke", "", "passed q4 staged desktop external Windows app smoke JSON")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("q4-staged-external-winapp-acceptance-preview does not accept positional arguments")
	}
	acceptance, err := appidentity.PreviewQ4StagedExternalWinAppAcceptance(appidentity.Q4StagedExternalWinAppAcceptanceRequest{
		SmokeReportPath: *smokeReportPath,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, acceptance)
}
