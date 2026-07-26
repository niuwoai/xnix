package main

import (
	"errors"
	"flag"
	"io"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runRealWinAppRunAcceptancePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.RealWinAppRunAcceptanceRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	executeResultPath := flags.String("remote-execute-result", "", "passed remote GUI smoke execute-result JSON")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("real-winapp-run-acceptance-preview does not accept positional arguments")
	}
	acceptance, err := appidentity.PreviewRealWinAppRunAcceptance(appidentity.RealWinAppRunAcceptanceRequest{
		RemoteExecuteResultPath: *executeResultPath,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, acceptance)
}
