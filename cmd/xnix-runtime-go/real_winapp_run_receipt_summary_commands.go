package main

import (
	"errors"
	"flag"
	"io"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runRealWinAppRunReceiptSummaryPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.RealWinAppRunReceiptSummaryRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	reportPath := flags.String("remote-smoke-report", "", "passed remote Wine guest GUI smoke execute-result JSON report")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("real-winapp-run-receipt-summary-preview does not accept positional arguments")
	}
	summary, err := appidentity.PreviewRealWinAppRunReceiptSummary(appidentity.RealWinAppRunReceiptSummaryRequest{
		RemoteSmokeReportPath: *reportPath,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, summary)
}
