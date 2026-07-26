package main

import (
	"errors"
	"flag"
	"io"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runQ4WinAppAcceptancePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.Q4WinAppAcceptanceRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	smokeReportPath := flags.String("q4-winapp-smoke", "", "passed q4 Windows app smoke JSON")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("q4-winapp-acceptance-preview does not accept positional arguments")
	}
	acceptance, err := appidentity.PreviewQ4WinAppAcceptance(appidentity.Q4WinAppAcceptanceRequest{
		SmokeReportPath: *smokeReportPath,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, acceptance)
}
