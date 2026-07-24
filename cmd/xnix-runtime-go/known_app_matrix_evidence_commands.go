package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runKnownAppMatrixEvidencePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.KnownAppMatrixEvidencePreviewRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	matrixReport := flags.String("matrix-report", "", "aggregate remote known-app matrix JSON report path")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", appidentity.KnownAppMatrixEvidencePreviewRequestType)
	}

	preview, err := appidentity.PreviewKnownAppMatrixEvidence(appidentity.KnownAppMatrixEvidencePreviewRequest{
		MatrixReportPath: *matrixReport,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}
