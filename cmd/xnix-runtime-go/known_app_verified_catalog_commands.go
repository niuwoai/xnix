package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runKnownAppVerifiedCatalogPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.KnownAppVerifiedCatalogPreviewRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	matrixEvidence := flags.String("matrix-evidence", "", "Go-owned known app matrix evidence JSON path")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", appidentity.KnownAppVerifiedCatalogPreviewRequestType)
	}

	preview, err := appidentity.PreviewKnownAppVerifiedCatalog(appidentity.KnownAppVerifiedCatalogRequest{
		MatrixEvidencePath: *matrixEvidence,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}
